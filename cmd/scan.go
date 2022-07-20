package cmd

import (
	"Seeyoner/core"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "漏洞检测",
	Long: `漏洞检测功能
`,
	Run: func(cmd *cobra.Command, args []string) {
		factory := new(core.IFactory)
		iScan := factory.NewFactory(vulnId)
		iScan.Scan(url)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&url, "targetUrl", "u", "", "targetUrl")
	scanCmd.Flags().IntVarP(&vulnId, "vulnId", "i", 0, "vulnId")
	scanCmd.MarkFlagRequired("targetUrl")
	scanCmd.MarkFlagRequired("vulnId")
}
