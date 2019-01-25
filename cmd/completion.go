package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion <SHELL_TYPE>",
	Short: "Generates bash completion scripts",
	Long: `BASH:

To load completion run

    source <(hetzner-kube completion bash)

To configure your bash shell to load completions for each session add to your "~/.bashrc" file

    # ~/.bashrc or ~/.profile
    echo 'source <(hetzner-kube completion bash)\n' >> ~/.bashrc

Or you can add it to your bash_completition folder:

    hetzner-kube completion bash > /usr/local/etc/bash_completion.d/hetzner-kube

ZSH:

To configure your zsh shell to load completions run following commands:

    # On linux
    hetzner-kube completion zsh | sudo tee /usr/share/zsh/vendor-completions/_hetzner-kube
    # On OSX
    hetzner-kube completion zsh | sudo tee /usr/share/zsh/site-functions/_hetzner-kube

Than rebuild autocomplete function with:

    compinits

`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch strings.ToLower(args[0]) {
		case "bash":
			return rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		default:
			return fmt.Errorf("Unable to generate completition script for shell %q, please specify `bash` or `zsh`", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
