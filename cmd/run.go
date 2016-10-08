package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/msmith491/gofio/iolib"
	"github.com/spf13/cobra"
)

var File string

func ParseConfig(configFile string) map[string]string {
	config := make(map[string]string)
	if configFile != "" {
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			fmt.Printf("Could not read configFile: %s, err: %s\n",
				configFile, err)
			panic(err)
		}
		datalines := strings.Split(string(data), "\n")
		for _, elem := range datalines {
			if elem == "" {
				continue
			}
			items := strings.Split(elem, "=")
			config[items[0]] = items[1]
		}
	}
	return config
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Command",
	Long: `Example:
	
	gofio run [workload_file]
	
	Example Workload File:
	
	$ cat test_workload.wlf
	io_type=random_write
	block_size=4k
	io_size=4G
	io_depth=32
	direct=true
	threads=5
	seed=someseedstring`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Printf("run Args: %s\n", args)
		if File != "" {
			fmt.Printf("%s", ParseConfig(File))
			r := iolib.GetRand(time.Now().UnixNano())
			iolib.WriteRandomBytesToDevice(r, 10)
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&File, "file", "f", "", "gofio workload file")

}
