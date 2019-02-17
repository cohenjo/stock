package main

import (
	"github.com/cohenjo/stock/pkg/config"
	"github.com/cohenjo/stock/pkg/scanner"
	"github.com/cohenjo/stock/pkg/trader"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// need to read more on this https://github.com/spf13/viper
var v *viper.Viper

func main() {

	var cmdScan = &cobra.Command{
		Use:   "scan [string to print]",
		Short: "Scan stocks",
		Long:  `do something super complex which finds stocks to buy.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   scan,
	}
	var cmdTrade = &cobra.Command{
		Use:   "trade [string to print]",
		Short: "Trade stocks",
		Long:  `do something super complex which buy stocks.`,
		Args:  cobra.MinimumNArgs(1),
		Run:   trade,
	}

	var rootCmd = &cobra.Command{Use: "stock"}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stock.yaml)")

	cmdScan.Flags().Int("port", 1138, "Port to run Application server on")
	cmdScan.Flags().String("name", "AAPL", "Port to run Application server on")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(cmdScan)
	rootCmd.AddCommand(cmdTrade)
	rootCmd.Execute()
}

func scan(cmd *cobra.Command, args []string) {
	// fmt.Println("Print: " + strings.Join(args, " "))
	v = newViper()
	v.BindPFlags(cmd.Flags())
	v.Unmarshal(&config.C)
	scanner.Scan()
}

func trade(cmd *cobra.Command, args []string) {
	// fmt.Println("Print: " + strings.Join(args, " "))
	v = newViper()
	v.BindPFlags(cmd.Flags())
	v.Unmarshal(&config.C)
	trader.Trade()
}

func newViper() *viper.Viper {
	v := viper.New()
	return v
}
