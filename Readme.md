# Trouble Tome

Trouble Tome is a tool designed to automate the creation of runbooks. It generates detailed documentation in Markdown or HTML format from structured JSON input, simplifying the process of documenting procedures for troubleshooting, incident management, and other operational tasks.

## Features

- **Generates Interactive runbook:** [UI mode](https://zeek-r.github.io/trouble-tome)
- **Automated Runbook Generation:** Converts JSON input into comprehensive runbooks.
- **Multiple Output Formats:** Supports both Markdown and HTML outputs for flexible documentation needs.
- **Structured JSON Input:** Easily define runbook content with a JSON structure.

## Installation
### Normal Way
> go install github.com/zeek-r/trouble-tome

### Way of Samurai
- Clone the repo
- Explore Makefile(build, run, test)

### How does it work?
- Create a runbook json with the following format.
```json
{
    "title": "Debezium CDC Connector Issue Resolution Runbook",
    "steps": [
        {
            "title": "Summary",
            "content": "What happened?"
        },
        {
            "title": "Dependencies Affected",
            "content": "What'll be affected due to this incident?"
        },
        {
            "title": "Tools Required",
            "content": "What do you need to fix this incident?"
        },
        {
            "title": "Mitigation Steps",
            "content": "How do you fix the incident?"
        },
        {
            "title": "Post Incident Steps",
            "content": "What needs to be done after the incident?"
        },
        {
            "title": "Important Resources",
            "content": "Anything more?"
        }
    ]
}
```
- Too lazy, like me, to create json? go [here](https://zeek-r.github.io/trouble-tome)

### Usage
#### Help Menu
> trouble-tome --help
