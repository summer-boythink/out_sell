const mysql = require('mysql')

let db = mysql.createConnection({host:"121.5.118.127",user:"root",password:'root',port:3311,database:"out_sell"})

module.exports = db
