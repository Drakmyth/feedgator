# FeedGator

A lightweight RSS feed aggregation service. FeedGator is **NOT** a reader.

### Built With

* [![Golang][golang-shield]][golang-url]
* [![Htmx][htmx-shield]][htmx-url]
* [![TailwindCSS][tailwindcss-shield]][tailwindcss-url]

## Installation

### Binary Release

In the future binaries will be available from the [releases](https://github.com/Drakmyth/feedgator/releases) page.

### Docker Compose

The included [docker-compose file](docker-compose.yaml) currently builds the image locally. In the future it will pull from [Docker Hub](https://hub.docker.com/).

```
> docker compose up -d
```

## Development

We recommend using [Air](https://github.com/cosmtrek/air) to facilitate execution and hot-reloading of the server during development. A [config file](./.air.toml) has been provided so running the server is as simple as executing:

```
> air
```

To facilitate debugging in VSCode, a [launch configuration](./.vscode/launch.json) has been included that you can use to run the server with the attached debugger.

### Configuration

FeedGator utilizes a variety of environment variables for configuration. These can either be set directly in the execution environment, included as part of a `docker-compose.yaml` definition, or provided in a `.env` file in the execution directory. The `.env` file will always be overridden by other methods so is particularly useful for development. The service will also fall back to a default value if one is not otherwise provided.

| Variable | Description                                                    | Default |
| -------- | -------------------------------------------------------------- | ------- |
| `PORT`   | The port that hosts the configuration UI and aggregated feeds. | 80      |

### Example `.env` File

```
PORT=80
```

## Contributing

This service is open source! That means if you'd like to try your hand at fixing a bug or implementing a feature, please do so! Head over to the Issues page and look for any open issues tagged with the `Accepting PRs` label. Fork this repo, create a branch, work on it, then submit a pull request. We'll work together to iron out any concerns with your code, and then we'll merge it in and your code will become a part of FeedGator's legacy!

## License

Distributed under the MIT License. See [LICENSE.md](./LICENSE.md) for more information.


<!-- Reference Links -->
[golang-url]: https://go.dev
[golang-shield]: https://img.shields.io/badge/golang-09657c?style=for-the-badge&logo=go&logoColor=79d2fa
[htmx-url]: https://htmx.org
[htmx-shield]: https://img.shields.io/badge/htmx-262626?style=for-the-badge&logo=htmx&logoColor=4470d2
[tailwindcss-url]: https://tailwindcss.com/
[tailwindcss-shield]: https://img.shields.io/badge/Tailwind%20CSS-0b111f?style=for-the-badge&logo=tailwindcss&logoColor=26bcf5

