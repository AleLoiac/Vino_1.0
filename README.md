# Ecommerce

This is the backend part of the project.

The project is organized this way: 
from the `routes` the command goes to `controllers` functions and from there to the `database` functions.

When performing queries we want to try to do as many operation as possible database side, using the built-in functions of MongoDB and avoiding working with data server side.
