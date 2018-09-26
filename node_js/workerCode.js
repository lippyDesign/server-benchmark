const randomstring = require("randomstring");
const { parentPort } = require('worker_threads');

const Person = require('./Person');

const persons = [];

// you can receive messages from the main thread this way:
parentPort.on('message', () => {
  for (let i = 0; i < 101; i++) {
    const person = new Person(
      randomstring.generate({ length: 7, charset: 'alphabetic'}), // first name
      randomstring.generate({ length: 7, charset: 'alphabetic'}), // lastName
      randomstring.generate({ length: 15, charset: 'alphabetic'}), // id
      randomstring.generate({ length: 10, charset: 'alphabetic'}), // phone
      randomstring.generate({ length: 15, charset: 'alphabetic'}), // street
      randomstring.generate({ length: 10, charset: 'alphabetic'}), // city
      randomstring.generate({ length: 2, charset: 'alphabetic'}), // state
      randomstring.generate({ length: 5, charset: 'alphabetic'}) // zip
    );
    persons.push(person)
  }
  parentPort.postMessage(persons);
});