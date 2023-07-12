package player

import (
	"context"
	"github.com/ejagombar/CLSpotify/authstore"
	"github.com/ejagombar/CLSpotify/prechecks"
	"github.com/spf13/cobra"
)

// pauseCmd represents the pause command
var PauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pauses the current song",
	Long: `Use this command to pause the current song playing
    If no song is playing, the command will not do anything.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := authStore.GetClient()
		prechecks.DeviceAvailable(client)
		cobra.CheckErr(err)
		client.Pause(context.Background())
	},
}

func init() {
}
