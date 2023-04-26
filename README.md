Создание первого пользователя:
``
curl --location 'http://localhost:8080/create' \
--header 'Content-Type: application/json' \
--data '{
"name": "Bogdan",
"age": 30
}'
``

Создание второго пользователя:
``
curl --location 'http://localhost:8080/create' \
--header 'Content-Type: application/json' \
--data '{
"name": "Sergey",
"age": 34
}'
``

Связывание пользователей:
``
curl --location 'http://localhost:8080/make_friends' \
--header 'Content-TYpe: application/json' \
--data '{
"source_id": 1,
"target_id": 2
}'
``

Удаление пользователя:
``
curl --location --request DELETE 'http://localhost:8080/user' \
--header 'Content-TYpe: application/json' \
--data '{
"target_id": 2
}'
``

Получение списка друзей пользователя:
``
curl --location 'http://localhost:8080/friends/2'
``
