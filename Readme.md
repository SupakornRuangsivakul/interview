# use golang and mongodb
## run -> docker-compose up 
## มี default user กับ interviewCard

# API Login

curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "username":"User01",
    "password":"password"
}'
## Response
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMSIsImV4cCI6MTY5OTE4NzYzMX0.Ph9vPWGACSN_Kx1Dlh_9Nfy9R3ax0l40lwQ2ce8DmTM"
}

# ใช้ token สำหรับ Authorization: Bearer Token
## default user มีการกำหนด level ของแต่ละ user
username : User01 
password : password 
level 1 (limit 3 request/sec)

username : User02
password : password
level 2 (limit 5 request/sec)

username : User03
password : password
level 3 (limit 7 request/sec)

# API Fetch All Cards
curl --location 'http://localhost:8080/card/fetch' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMSIsImV4cCI6MTY5OTE3ODE0M30.m6EkQfqVcaESu2zVevitigt1PVdrdRkWP9UGwMnEbDQ'

# API View Card
curl --location 'http://localhost:8080/card/view/:cardId' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMiIsImV4cCI6MTY5OTE4NzEyOX0.1RaWqyA1V7o0MgHkDjJfXBeSt3UAam38ovug_4C2mWM'

# API View Card
curl --location 'http://localhost:8080/card/view/:cardId' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMiIsImV4cCI6MTY5OTE4NzEyOX0.1RaWqyA1V7o0MgHkDjJfXBeSt3UAam38ovug_4C2mWM'

# API Add Comment to Card
curl --location 'http://localhost:8080/card/comment/add' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMyIsImV4cCI6MTY5OTE3NDY3OX0.EtzQwRy8jJ8NRW5KUmQ1XU1W0bdQ9e3-2Tin4RCf5EI' \
--data '{
    "cardId":"654733dc16014ca9590e353d",
"comment":"TestComment#4"
}'

# API Edit Comment
curl --location 'http://localhost:8080/card/comment/update' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMSIsImV4cCI6MTY5OTE3MzEyNn0.8quFhgQROy6-TsH5iHz_b4Y3gI1rIGdvsJaFXPUjXHk' \
--data '{
"comment":"TestComment#1 updated",
"commentId":"654744f20098fd3a486aa08d"
}'
## edit ได้เฉพาะ user ที่ comment เท่านั้น

# API Remove Comment
curl --location 'http://localhost:8080/card/comment/remove' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMiIsImV4cCI6MTY5OTE3MzA5MH0.2_J6bQPOXvz57D1ikJHOja4ldJ5oxaRd_YGF8uibzNw' \
--data '{
"comment":"TestComment#1 updated",
"commentId":"654744f20098fd3a486aa08d"
}'
## remove ได้เฉพาะ user ที่ comment เท่านั้น

# API Edit Card
curl --location 'http://localhost:8080/card/update' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMyIsImV4cCI6MTY5OTE3NDY3OX0.EtzQwRy8jJ8NRW5KUmQ1XU1W0bdQ9e3-2Tin4RCf5EI' \
--data '{
    "cardId":"654733dc16014ca9590e353d",
"cardName":"นัดสัมภาษณ์งาน 9 #Edit2",
"cardDetail":"Test for update interview card 2",
"cardStatus":"In Progress"
}'
## เก็บ changelog ฟิลด์ที่มีการแก้ไข ชื่อการ์ด,รายละเอียด,สเตตัส

# API get history
curl --location 'http://localhost:8080/card/history/:cardId' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMyIsImV4cCI6MTY5OTE3NDY3OX0.EtzQwRy8jJ8NRW5KUmQ1XU1W0bdQ9e3-2Tin4RCf5EI'
## เรียกดูประวัติการแก้ไขของการ์ด

# API Keep card
curl --location 'http://localhost:8080/card/keep/:cardId' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyTmFtZSI6IlVzZXIwMyIsImV4cCI6MTY5OTE3NjM4OH0.2dN-D2eI45_eoTk6ghV4V7LJBlQUdEZdiBkgOCIHb_U'
## จัดเก็บ card จะเปลี่ยน status = Keep ตอน fetch จะไม่หยิบ status นี้ออกมา
