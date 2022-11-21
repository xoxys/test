def main(ctx):
    return [{
        "kind": "pipeline",
        "type": "docker",
        "name": "test",
        "platform": {
            "os": "linux",
            "arch": "amd64",
        },
        "steps": [
            {
                "name": "deps",
                "image": "alpine",
                "commands": [
                    "echo \"foobar\" | tr -d \"",
                ],
            },
        ],
    }]
