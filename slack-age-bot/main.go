package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommamdEvents(analyticChannel <-chan *slacker.CommandEvent) {
	for event := range analyticChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	slack_bot_token := goDotEnvVariable("SLACK_BOT_TOKEN")
	slack_app_token := goDotEnvVariable("SLACK_APP_TOKEN")

	fmt.Println(slack_bot_token)
	fmt.Println(slack_app_token)

	bot := slacker.NewClient(slack_bot_token, slack_app_token)

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		// Example: "my yob is 2020",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			age := 2024 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	go printCommamdEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
