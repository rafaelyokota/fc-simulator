
# Simulador Posições

## Dependencias 
    - docker
    - docker-compose

## ambiente de desenvolvimento
- docker-compose up -d  #iniciar cotainer
- docker-compose ps    #ver se esta rodando
- docker exec -it simulator bash  # entra dentro do cotainer
- docker-compose up -d (.docker/kafka)
- docker-compose ps
- docker exec -it kafka_kafka_1 bash 
  - kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readtest
  - kafka-console-producer --bootstrap-server=localhost:9092 --topic=readtest
    - kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position
    - kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
      - {"clientId":"1", "routeId": "1"}
      - {"clientId":"1", "routeId": "2"}
      - {"clientId":"1", "routeId": "3"}

## Historico comandos
'#go mod init github.com/rafaelyokota/imersaofc-simulator'

## Objetivo
Recebe as rotas e comeca disparar pontos da rota