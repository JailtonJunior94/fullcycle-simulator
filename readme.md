# Aula 1 da imersão fullcycle 8

### Kafka
1. Acesso o container kafka 
  ```
  docker exec -it kafka bash
  ```

2. Produzindo mensagem
   ```
   kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction 
   ```

3. Consumindo mensagem
   ```
   kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position
   ```