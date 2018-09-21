const express = require('express');
var randomstring = require("randomstring");
const path = require('path');
const Person = require('./Person');

// set up express
const app = express();
app.use(express.static('client/build'));

// GET 'Hello World!' will send back 'hello world' as a string
app.get('/helloworld', (req, res) => {
  res.send('Hello World!');
});

// GET json blob- array of 100 random persons
app.get('/persons', (req, res) => {
  const persons = [];
  for(let i = 0; i < 101; i++) {
    const person = new Person(
      randomstring.generate({ length: 7, charset: 'alphabetic'}), // first name
      randomstring.generate({ length: 7, charset: 'alphabetic'}), // lastName
      randomstring.generate({ length: 15, charset: 'alphabetic'}), // id
      randomstring.generate({ length: 10, charset: 'numeric'}), // phone
      randomstring.generate({ length: 15, charset: 'alphabetic'}), // street
      randomstring.generate({ length: 10, charset: 'alphabetic'}), // city
      randomstring.generate({ length: 2, charset: 'alphabetic'}), // state
      randomstring.generate({ length: 5, charset: 'numeric'}) // zip
    );
    persons.push(person);
  }
  res.send(persons);
});

// GET static files
app.get('/', (req, res) => {
  res.sendFile(path.resolve('client', 'build', 'index.html'));
});

// set up port the server will listen on
const PORT = process.env.PORT || 5000;
app.listen(PORT, () => console.log(`Listening on port`, PORT));