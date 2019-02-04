
package main

import(
  "fmt"
  "github.com/bwmarrin/discordgo"
  "strings"
)

var(
  commandPrefix string
  botID string
)

func startBot(){
  discord, err := discordgo.New("Bot NTM2MjMyODEyNzI3MTA3NjY0.DyTxsA.puWGzmgEpADoiAdExoBHnS2TNAA")
  errCheck("error creating discord session", err)
  user, err := discord.User("@me")
  errCheck("error retrieving account", err)

  botID = user.ID
  discord.AddHandler(commandHandler)
  discord.AddHandler(func(discord *discordgo.Session, read *discordgo.Ready){
    err = discord.UpdateStatus(0, "A friendly helpful bot!")
    if err != nil {
      fmt.Println("Error attempting to set my status")
    }
    servers := discord.State.Guilds
    fmt.Printf("GoClashbotGo has started on %d servers", len(servers))
  })

  err = discord.Open()
  errCheck("Error opening connection to Discord", err)
  defer discord.Close()
  
  commandPrefix = "!"

  <-make(chan struct{})

}

func errCheck(msg string, err error){
  if err != nil{
    fmt.Printf("%s: %+v", msg, err)
    panic(err)
  }
}

//waits for the someone to tyoe something within the discord channel
func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate){
  user := message.Author
  if user.ID == botID || user.Bot{
    //do nothing because bot is talking
    return
  }
  if message.Content == "!ping"{
    discord.ChannelMessageSend(message.ChannelID, "Pong!")
  }
  if strings.Contains(message.Content, "!clan"){
    var clan = strings.TrimPrefix(message.Content, "!clan ")

    discord.ChannelMessageSend(message.ChannelID, ("Getting clan information for " + clan + "..."))
    discord.ChannelMessageSend(message.ChannelID, getClanInfo(clan))
  }
  //content := message.Content

  fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author)
}
