# Polls

A website where you can add and share polls with others.

## Contributing

Feel free to contribute to this project, read the [contributing guidelines](CONTRIBUTING.md).

## Deployment

**Running the server locally**

You'll need a firebase project and an [admin sdk key](https://firebase.google.com/docs/admin/setup#go)

```bash
FIREBASE_CRED=`cat your-firebase-adminsdk.json"` go run server.go
```

**Deploying**
You can deploy the project anywhere you want. Set u the the environment variable
`FIREBASE_CRED` and run the server.

The project is written to work with **Vercel**, but can also be used as a regular server.
If you do use it as a server, you can edit the way firebase is authenticated so it doesn't do that
on every request.

## Example requests

```bash
# Poll info
curl -X GET mydomain.com/api/info -d '{"pollId": "SOME_POLL_ID"}'
> {"Title":"Language","Votes":{"Rust":1,"Swift":5}}
```

## License

This project is licensed under the [MIT license](LICENSE).
