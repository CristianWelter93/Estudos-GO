# Estudos-GO

Projeto criado durante os estudos iniciais da linguagem GO

## Instalação

É necessário baixar e instalar o GO em sua máquina para que se possa executar o projeto.
Acesse <https://go.dev/dl/> e faça o download.

Caso esteja utilizando windows é necessário editar as variáveis do sistema, adicione a GO_HOME e selecione o caminho para a pasta "...go/bin".
Para testar se está tudo ok execute no terminal

~~~bash
go version
~~~

Se exibir a versão está tudo correto.

## IDE

Eu utilizei o VSCode por ter maior familiaridade, mas qualquer IDE pode ser utilizada.
Caso deseje utilizar o VSCode faça a instalação dele, depois vá em Extensions e adicione a GO (ela é mantida e atualizada pelo próprio google).

## Arquivos existentes no projeto

### hello.go

Clássico hello world, apenas são printadas mensagens. Para executa-lo basta acessar a pasta e digitar

~~~bash
go run hello.go
~~~

Este mesmo comando pode ser utilizado para executar os demais arquivos, basta substituir o hello pelo nome do arquivo.

### variaveis.go

Nele encontrasse o conteúdo relacionado a declaração de variáveis em go.

### loop.go

É demonstrado as diferentes formas de se utilizar o for, já que não existe while e nem do while em GO.
