# ASCII-Art-Web
 
Ascii-art-web is a web server application that allows users to generate ASCII art using different banner styles through a web GUI (graphical user interface). Users can input text and select from various banners to create unique ASCII art. The available banner styles are:

* shadow
* standard
* thinkertoy

Authors:

* hilaromondi
* sbull

Usage:

Clone the repository:

```bash
git clone https://learn.zone01kisumu.ke/git/hilaromondi/ascii-art-web.git

cd ascii-art-web

Run the server:

go run main.go
```

Open your web browser and navigate to:

http://localhost:8080

Implementation Details: Algorithm
HTTP Endpoints

    GET /:
        Sends the main HTML page as a response.
        Uses Go templates to render and display data from the server.
        The main page includes:
            Text input for user text.
            A select dropdown to choose between banners.
            A button to send a POST request to /ascii-art.

    POST /ascii-art:
        Receives text and a selected banner style from the client.
        Generates ASCII art based on the input text and selected banner.
        Displays the result on the page.

HTTP Status Codes

    200 OK: Returned if everything went without errors.
    404 Not Found: Returned if templates or banners are not found.
    400 Bad Request: Returned for incorrect requests.
    500 Internal Server Error: Returned for unhandled errors.

## Main Page

The main page of the web application includes:

    A text input field for users to enter the text they want to convert to ASCII art.
    Radio buttons, a select dropdown, or other input elements to allow users to choose between the available banners.
    A submit button that sends a POST request to /ascii-art and displays the generated ASCII art on the same page or on a new page, based on the chosen implementation approach.

## Example Workflow

    User visits the main page:
        Enters text into the text input field.
        Selects a banner style using  a dropdown menu.
        Clicks the submit button to send a POST request to /ascii-art.

    Server processes the request:
        Receives the text and banner style.
        Generates the ASCII art using the selected banner.
        Returns the ASCII art as part of the response.

    User views the result:
        The generated ASCII art is displayed  below the textarea.
