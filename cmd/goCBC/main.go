package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmbracelet/log"
	"github.com/sethll/goCBC/pkg/chems"
	"github.com/sethll/goCBC/pkg/progmeta"
	"github.com/sethll/goCBC/pkg/progutils"
	"github.com/spf13/cobra"
)

var (
	verbosity int
	chem      string
	chemMHL   float64
	listChems bool
)

func main() {
	progutils.PrintProgHeader()

	rootCmd := &cobra.Command{
		Use:     progmeta.Usage,
		Short:   progmeta.ShortDesc,
		Long:    progmeta.LongDesc,
		Example: progmeta.UsageExample,
		Args: func(cmd *cobra.Command, args []string) error {
			// Skip argument validation if --list-chems flag is used
			if listChems {
				return nil
			}
			// Must have at least 2 arguments (1 required + at least 1 remaining)
			if len(args) < 2 {
				return fmt.Errorf("requires at least 2 arguments: <target> <time_amount...>")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if listChems {
				chems.ListAvailableChems()
				return
			}
			runApp(args)
		},
	}

	rootCmd.Flags().CountVarP(&verbosity, "verbose", "v", "increase verbosity (use -v, -vv, -vvv)")
	rootCmd.Flags().StringVarP(&chem, "chem", "c", "caffeine", "choose chem")
	rootCmd.Flags().BoolVar(&listChems, "list-chems", false, "list all available chem options")
	//rootCmd.RegisterFlagCompletionFunc("chem", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	//	return []string{"caffeine", "nicotine"}, cobra.ShellCompDirectiveDefault
	//})
	rootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if halfLife, exists := chems.Available[chem]; !exists {
			return fmt.Errorf("invalid chem option '%s'", chem)
		} else {
			chemMHL = halfLife
		}
		if verbosity > 3 {
			verbosity = 3
		}
		return nil
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runApp(args []string) {
	initLogging()
	firstArg := args[0]
	remainingArgs := args[1:]

	targetAmount := progutils.StringToFloat(&firstArg)
	timesAndAmounts := progutils.GetTimesAndAmounts(&remainingArgs)
	timesAndAmounts = progutils.SortTimeEntries(timesAndAmounts)

	slog.Info("Finalized time/amount inputs", "targetAmount", targetAmount, "timesAndAmounts", timesAndAmounts)

	bodyChemContent, chemTargetReachedTime := progutils.RunHLCalculations(&timesAndAmounts, &targetAmount, &chemMHL)

	// Generate and print output
	fmt.Println(progutils.GenerateOutputTable(&bodyChemContent, &chemTargetReachedTime, &firstArg, &chem))
}

func initLogging() {
	// Replace the default slog logger with charmbracelet handler
	handler := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		Level:           progutils.LogLevelSelector[verbosity],
		TimeFormat:      "15:04:05",
	})

	// Set as default slog handler - intercepts ALL slog calls
	slog.SetDefault(slog.New(handler))
	//slog.SetLogLoggerLevel(slog.LevelDebug)
}
