package cmd

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/spf13/cobra"
	"golang2019/golang_2019/homeworks/Anatolii.Mezentsev-Mezencev/homework3/model"
)

var vendors string

func createVendorMachine(vendors string) [][]string {

	re := regexp.MustCompile(`/`)
	index := re.FindAllStringSubmatchIndex(vendors, -1)

	var s = make([][]string, len(index)+1)
	numberVendor := 0
	a := []string{}

	arr := strings.Split(strings.Replace(vendors, " ", "", -1), "")
	for i := 0; i < len(arr); i++ {
		if arr[i] == "/" {
			s[numberVendor] = a
			numberVendor++
			a = []string{}
			continue
		}

		a = append(a, arr[i])
		if i+1 == len(arr) {
			s[numberVendor] = a
		}
	}
	return s
}

var countCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		vendors, _ := cmd.Flags().GetString("vendors")
		vendor := createVendorMachine(vendors)
		fmt.Println("Vandors:", vendor)
		order, _ := cmd.Flags().GetString("order")
		fmt.Println("Order:", order)
		c := model.VendingMachine{Backed: vendor}
		var wg sync.WaitGroup
		wg.Add(1)
		go c.SearchOrder(strings.Split(order, " "), &wg)
		wg.Wait()
	},
}

func init() {
	orderCmd.AddCommand(countCmd)
	orderCmd.PersistentFlags().StringVar(&vendors, "vendors", "1123/234/13/543/221", "vendors")

}
