package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

var (
	version  = "dev"
	revision = "none"
)

func main() {
	// Setup logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	app := &cli.Command{
		Name:  "mdstrip",
		Usage: "Strip Markdown formatting from text",
		Description: `mdstrip removes Markdown formatting from text files or stdin.
It preserves the actual content while removing formatting syntax like
headers (#), emphasis (**/__ ), code blocks, links, etc.`,
		Version: fmt.Sprintf("%s (rev: %s)", version, revision),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output file (default: stdout)",
			},
			&cli.BoolFlag{
				Name:    "keep-links",
				Aliases: []string{"l"},
				Usage:   "Keep link URLs in output",
			},
			&cli.BoolFlag{
				Name:    "keep-code",
				Aliases: []string{"c"},
				Usage:   "Keep code block markers",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"V"},
				Usage:   "Enable verbose logging",
			},
		},
		Action: handleStrip,
		Commands: []*cli.Command{
			{
				Name:   "version",
				Usage:  "Show version information",
				Action: handleVersion,
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal().Err(err).Msg("Failed to run application")
	}
}

func handleStrip(ctx context.Context, c *cli.Command) error {
	// Set log level
	if c.Bool("verbose") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Get input
	var input io.Reader = os.Stdin
	if c.Args().Len() > 0 {
		filename := c.Args().Get(0)
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", filename, err)
		}
		defer file.Close()
		input = file
		log.Debug().Str("file", filename).Msg("Reading from file")
	} else {
		log.Debug().Msg("Reading from stdin")
	}

	// Read content
	content, err := io.ReadAll(input)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	// Strip markdown
	options := StripOptions{
		KeepLinks: c.Bool("keep-links"),
		KeepCode:  c.Bool("keep-code"),
	}
	stripped := StripMarkdown(string(content), options)

	// Output
	outputFile := c.String("output")
	if outputFile != "" {
		if err := os.WriteFile(outputFile, []byte(stripped), 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
		log.Info().Str("file", outputFile).Msg("Output written to file")
	} else {
		fmt.Print(stripped)
	}

	return nil
}

func handleVersion(ctx context.Context, c *cli.Command) error {
	fmt.Printf("mdstrip version %s (revision: %s)\n", version, revision)
	fmt.Printf("Built with %s\n", getGoVersion())
	return nil
}

func getGoVersion() string {
	// Use runtime.Version() instead of environment variable
	return "Go runtime"
}