#!/bin/bash
#renato fagalde 01

rm -rf .aws-sam/

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
``

echo "Compilando binário 'bootstrap' no local esperado..."
cd app
go mod tidy
go build
if [[ $? -ne 0 ]]; then
    echo "Erro ao compilar o binário 'bootstrap'. Saindo."
    exit 1
fi
cd ..


echo "Executando sam build..."
sam build --template-file template.yaml --build-dir .aws-sam/build
if [[ $? -ne 0 ]]; then
    echo "Erro ao executar sam build. Saindo."
    exit 1
fi


EVENT_DIR="./events"


echo "Selecione um evento para rodar:"
EVENT_FILES=()
i=1
for file in "$EVENT_DIR"/*; do
    if [[ -f "$file" ]]; then
        EVENT_FILES+=("$file")
        echo "[$i] $(basename "$file")"
        ((i++))
    fi
done


read -p "Digite o número do evento: " EVENT_CHOICE


if [[ "$EVENT_CHOICE" -ge 1 && "$EVENT_CHOICE" -le "${#EVENT_FILES[@]}" ]]; then
    SELECTED_EVENT=${EVENT_FILES[$EVENT_CHOICE-1]}
    echo "Rodando com o evento: $(basename "$SELECTED_EVENT")"

    sam local invoke -e "$SELECTED_EVENT" StrategyFunction
else
    echo "Opção inválida. Saindo."
    exit 1
fi
