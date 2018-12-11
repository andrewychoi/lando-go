The Lando system is built up of a number of primitives:

A **User** is the person who uses the system. One has:

- a name,
- an email,
- and other contact information.

A **Person** is a person that a user would like to keep in contact with. One has:

- a name,
- a frequency with which they would like to be contacted.
- a User that they are associated to,
- notes

A **Meeting** is a time at which a User interacted with a Person. A meeting has:

- a date and time,
- a User,
- a Person,
- a location,
- a note.

Locations of meetings are just text fields, but can be filled using the Google Maps API.

A **Note** is a record that someone wants to keep on a person. They have:

- a text field,
- a User,
- a Person,
- a date and time.
