# Twitter Weather updater

### Add some weather info to your Twitter profile!

This is a little Go program I built to learn more about Go's syntax.

It grabs current weather information from the https://openweathermap.org/ API and updates 
your Twitter display name with a little emoji representation of the current weather 🌦.

It's more of a proof of concept at this point, but it works well.

[<img src="screen-shot.png.png" width="250"/>](screen-shot.png))

How to run:

1. Sign up for a free account at https://openweathermap.org/ and get an AppID
2. Create a Twitter app and grab your authentication info. This might take a while since Twitter apps require review(😬)
3. Update the `.env` file with these credentials

To run:

`go run main.go`
