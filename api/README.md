## API

This directory is for the first and second part of the assignment

Objective 1 :  Create a basic web server that listens on a specific port and returns “Hello, World!” when a user accesses the root endpoint.

Objective 2 : Create a basic CRUD (Create, Read, Update, Delete) API that allows a user to manage a list of items(e.g., books, movies, etc.)

### Port number used : 6004

### API Endpoints :

Home page - http://localhost:6004

CREATE - POST - http://localhost:6004/book
               
This adds a new book to the given list of books with details mentioned in the request body
An example body for this post request : 
{
  "Id": "9", 
  "Title": "New book", 
  "Author": "New guy", 
  "Price": "20.00" 
}
         
READ - GET  -  http://localhost:6004/books
               
               This returns all the books present in the list
               
READ - GET  -  http://localhost:6004/book/{id}
               
               (This returns the details of the book with the given id)
               (example : http://localhost:6004/book/3 gives the info of the book with id 3)
               
UPDATE - PUT -  http://localhost:6004/book/{id}
                
                (This lets you update the info of the book with the details mentioned in the request body)
                
                (Body should be similar to the one while creating a book. All the details should be mentioned)
                
DELETE - DELETE - http://localhost:6004/book/{id}
                  
                  (This deletes the book with the given id)
                  
                  
(Code should be run first before using the port)

               

