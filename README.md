# Projeto template para APIS com Golang + postgress

```
Projeto a ser usado como base para criação de apis utilizando golang
```


### Estrutura de pastas do projeto
```
.
├── cmd                     Pasta onde está presente os main packages.
│    ├── api                  main package responsável pela api.
│    └── scheduledjobs        main package responsável pelas crons do projeto.
│ 
├── config                  Pasta que deve ser criada pelo dev. É usada para adicionar o .env .
├── core                    Package responsável por carregar o env e realizar o pool de conexões redis e banco de dados da aplicação.
├── database                Pasta utilizada para armazenar e versionar scripts de banco de dados.
│
├── docker                  Pasta responsável por armazenar arquivos referente ao docker e docker-compose.
│    ├── api                  Pasta utilizada para armazenar e versionar a imagem(Dockerfile) responsável pela api.
│    ├── scheduledjobs        Pasta utilizada para armazenar e versionar a imagem(Dockerfile) responsável pelas crons.
│    └── redis                Pasta utilizada para armazenar e versionar a imagem(Dockerfile) responsável pelo redis utilizado para desenvolvimento local.
│
├── domain                  Package responsável por TODAS as Structs do projeto(libs não incluidas), além de ser também utilizado para o carregamento de algumas variáveis de ambiente.
│
├── handlers                Package responsável pelas regras de negócio do projeto. (alternatively `controllers`)
│
├── http                    Package responsável por ações que envolvem o protocolo HTTPS.
│    ├── request              Pasta onde são armazenadas requisições para para APIs externas(Nesse projeto apenas Infobip). (alternatively `client`)
│    └── rest                 Pasta responsável por receber requisições externas ao projeto e redirecionar para os handlers.
│  
├── infrastructure          Pasta responsável por armazenar arquivos referente a terraform.
├── job                     Package responsável por realizar as conexões com redis, banco de dados e newrelic(APM) quando é usado um cronjob.
├── middleware              Package responsável por realizar panic recovers durante uma requisição serviço-externo->loki. Não funciona no scheduler.
├── repository              Package responsável por armazenar funções de banco do projeto (querys sql+go).
├── router                  Package responsável pelas rotas do projeto.
├── tmp/downloads           Pasta responsável por armazenar downloads temporários em runtime.
└── vendor                  Pasta responsável por armazenar modulos(`libs`) utilizadas no projeto.
```


# Como rodar o projeto
- Crie o arquivo '.env' na pasta config
- Configurar uma chave dev no Github para usar repositórios privados - [Tutorial](https://docs.google.com/document/d/1-3_ptGklPSj-sCgY0iTUd0q3A8Hum_mwgr-5COckZpg/edit?usp=sharing)
- Configurar aws-cli

## Executar localmente
```
cd config
# run scheduler
go run ../cmd/scheduledjobs/scheduledjobs.go
```

```
### Para rodar as schedules
```shell
./docker-redis.sh
./air-scheduledjobs.sh
```

## Usando Docker
Execute na raiz do projeto
```
sh docker-start.sh
```


## Instalação do Go
```
# download the latest version o Golang Toolchain
wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz

# install locally
sudo tar -C /usr/local -xzf go1.16.2.linux-amd64.tar.gz

# set path
export PATH=$PATH:/usr/local/go/bin

# verify go version
go version

# modify your profile to export the path for all sessions
nano ~/.profile

export PATH=$PATH:/usr/local/go/bin

# reiniciar sessão para que as alterações passem a ter efeito
```

## Instalação do client aws (aws-cli) - necessário para configurar ambiente local de desenvolvimento com sua chave
```
cd ~
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
aws --version
```

## Configurando seu acesso aws via aws-cli
```
aws configure
```

#### Example:
```
AWS Access Key ID [None]: AKIAIOSFODNN7EXAMPLE
AWS Secret Access Key [None]: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
Default region name [None]: us-east-2
Default output format [None]: json
```

