# alf-opensearch

Alfresco OpenSearch Deployment CLI using [Docker Desktop](https://docs.docker.com/desktop/).

Additional details are available in [Alfresco Deployment](https://github.com/Alfresco/acs-deployment/tree/master/docker-compose).

## Usage

Download the binary compiled for your architecture (Linux, Windows or Mac OS) from [**Releases**](https://github.com/aborroy/alfresco-opensearch-deployer/releases).

>> You may rename the binary to `alf-opensearch`, all the following samples are using this command name by default.

Using `-h` flag provides detail on the use of the different commands available.

**Create**

`create` command produces required assets to deploy Alfresco Community in Kubernetes.

```bash
$ ./alf-opensearch create -h
Alfresco OpenSearch Deployment CLI

Usage:
  alf-opensearch [flags]
  alf-opensearch [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create assets to deploy Alfresco and OpenSearch
  help        Help about any command

Flags:
  -h, --help   help for alf-opensearch

Use "alf-opensearch [command] --help" for more information about a command.
```

### Creating a sample deployment

**Using command line parameters**

Run the command selecting the OpenSearch version to be deployed . Additional options can be set using any parameter value from *flags* list.

```bash
$ ./alf-opensearch create -i=true -v 2.13.0
```

**Replying to prompts**

Run the command using interactive mode.

```bash
$ ./alf-opensearch create
? Which OpenSearch version do you want to use? 2.13.0
```

>> Even when using interactive mode, output directory can be specified using the `-o` flag.

**Output folder**

Docker Compose assets will be produced by default in `output` folder:

```bash
$ tree output
output
├── compose.yaml
├── db
│   └── compose.yaml
├── legacy-ui
│   ├── compose.yaml
│   └── context.xml
├── messaging
│   └── compose.yaml
├── proxy
│   └── compose.yaml
├── repo
│   └── compose.yaml
├── search
│   ├── compose.yaml
│   └── reindex.prefixes-file.json
├── transform
│   └── compose.yaml
└── ui
    └── compose.yaml
```

Alfresco can be deployed using default command:

```bash
$ cd output
$ docker compose up
```

Default endpoints:

* http://localhost:8080/alfresco
* http://localhost:5601

>> Use default credentials `admin`/`admin`