# Playlist Generator with Random Genre Picker
## **Overview**
Welcome to the Playlist Generator with Random Genre Picker! This application utilizes Spotify's Web API to generate playlists based on random genre selection. Users can log in using Auth0 authentication and enjoy the fun of exploring diverse music genres.

## **Features**
**Random Genre Picker**: Explore various music genres with a simple click until you find the one that suits your mood.
**Spotify Integration**: Connects to Spotify's Web API to fetch genre-specific songs.
**Auth0 Authentication**: Securely log in using Auth0 to personalize your experience and create playlists.
**Song Selection**: Before generating the playlist, users can choose specific songs from the generated list.

## How to Run
Follow these steps to run the application locally:
- Create a `.env` file in the project root directory with the following content:
```
CLIENT_ID=your_spotify_client_id
CLIENT_SECRET=your_spotify_client_secret
```
Replace your_spotify_client_id, your_spotify_client_secret, your_auth0_domain, your_auth0_client_id, and your_auth0_client_secret with your actual Spotify and Auth0 credentials.

- Open a terminal and run the following command: `go run main.go`
- Open your web browser and navigate to http://localhost:8080.
- Log in with your Auth0 credentials and start exploring and generating playlists!

## Dependencies
[Auth0](https://auth0.com/docs): For secure authentication.

[Spotify Web API](https://developer.spotify.com/documentation/web-api): To fetch music data and create playlists.

## Contributing
Contributions are welcome! If you find any issues or have suggestions, feel free to open an issue or submit a pull request.

License
This project is licensed under the MIT License. Feel free to use and modify the code for your needs.

Enjoy exploring new music with the Playlist Generator! 🎶
