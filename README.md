# Go-AN
Golang Code Test

1. Finish TestCase1
   Answer : You can see code in file number1.go
2. Buat sebuah fungsi yang mengecek apakah sebuah string itu palindrome atau tidak
  Palindrome = sebuah string yang jika dicek sequencenya dari depan ke belakang itu sama
  Answer : You can see code in file palindrome.go
3. Diberikan sebuah Set Data yang bisa dipreview di bagian link
  https://docs.google.com/spreadsheets/d/1_rr63XNR9fUmZ6Ee2mNUG4Pq17J8G6pl7Xb9_F9R3wY/edit?
  usp=sharing

  (Notes : Untuk Studi Kasus Ini Yang saya tangkap saya melakukan Convert dari CSV ke SQL Untuk Dokumentasi Cara convert saya cantumkan di file test-case-backend-eratani.docx)
  a. Munculkan data country mana aja yang spend nya terbanyak (Query)

  Untuk Query Golang seperti dibawah ini 
  
  result := db.Table("creditcards").
		Select("credit_card_type, COUNT(*) as total").
		Group("credit_card_type").
		Order("total DESC").
		Limit(1).
		Scan(&resultHighestCreditCardType)

  Convert ke SQL

  SELECT country, SUM(REPLACE(credit_card, ',', '')) as total_spend
  FROM creditcards
  GROUP BY country
  ORDER BY total_spend DESC
  LIMIT 1;

  Atau bisa running API get di url http://localhost:8000/get-highest-credit-card-countries 


  b. Munculkan data jumlah tipe kartu kredit terbanyak (Query)

  Untuk Query Golang seperti dibawah ini 
  
 db.Table("creditcards").
		Select("credit_card_type, COUNT(*) as total").
		Group("credit_card_type").
		Order("total DESC").
		Limit(1).
		Scan(&resultHighestCreditCardType)

  Convert ke SQL

SELECT credit_card_type, COUNT(*) AS total
FROM creditcards
GROUP BY credit_card_type
ORDER BY total DESC
LIMIT 1;


  Atau bisa running API get di url http://localhost:8000/get-highest-credit-card-type
  
  c. Buatlah GET API untuk medapatkan data dari soal a dengan response menampilkan :
  a) Id
  b) Country
  c) Credit_card_type
  d) Credit_card
  e) First_name
  f) Last_name

  Answer : Anda Bisa cek di API documentation dengan klik folder apitest->eratani_test.postman_collection.json lalu import kedalam postman anda
  d. Buatlah POST API untuk mengirim data dengan body seperti dibawah ini :
  {
  "country”: " ",
  "credit_card_type”: " ",
  "credit_card”: " ",
  "first_name”: " ",
  "last_name”: " "
  }
  

  Answer : Anda Bisa cek di API documentation dengan klik folder apitest->eratani_test.postman_collection.json lalu import kedalam postman anda

 Cara test : dengan cara run di folder environtment apitest -> go mod tidy -> lalu jalankan go run main.go
 
  4. Buat sebuah fungsi yang mensortir sebuah bilangan acak yang digenerate oleh sistem (MUST USE 
ALGORITHM)
Answer : Anda Bisa cek di file number4.go.


