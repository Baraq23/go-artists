# groupie-tracker

## Objectives

Groupie Trackers is a project aimed at building a user-friendly website that interacts with a given API to display information about bands and artists. The project involves fetching and manipulating data from the API, and displaying it in various formats such as tables, lists, blocks, cards, or any other form of data visualization.

### API Structure

The provided API consists of four main parts:

1. **Artists**:
    - Contains information about bands and artists, including their names, images, the year they began their activity, the date of their first album, and their members.

2. **Locations**:
    - Contains data about their last and/or upcoming concert locations.

3. **Dates**:
    - Contains information about their last and/or upcoming concert dates.

4. **Relation**:
    - Links together the artists, locations, and dates.

### Project Requirements

- **Frontend**: A user-friendly website that displays the artist and band information through various data visualizations (e.g., blocks, cards, tables, lists, pages, graphics, etc.).
- **Backend**: The server must be written in **Go**, ensuring no crashes or downtime, and handling all errors gracefully.
- **Event/Action**: Implement a feature of your choice that triggers an action from the client-side, such as fetching additional information from the server. This should follow a [request-response](https://en.wikipedia.org/wiki/Request%E2%80%93response) model.
- **Data Visualization**: Information must be displayed in an intuitive and organized manner, enhancing user experience.
- **Client-Server Communication**: Actions or events should trigger communication between the client and server.

## Instructions

1. **Backend**:
    - Written in **Go**.
    - Must be robust and handle errors effectively.
    - Follow best coding practices.
    - Include unit tests for critical functions.

2. **Frontend**:
    - Display the data fetched from the API through different forms of data visualization.
    - Ensure smooth user interactions with responsive design and efficient loading of data.
    - Implement an action/event that triggers client-server communication (e.g., fetching additional information when a user clicks on an artist).

3. **API Integration**:
    - Interact with the API to fetch and manipulate data.
    - Display artist and concert information appropriately.

4. **Testing**:
    - Ensure that your code is thoroughly tested, especially for backend logic.
    - Include test files for unit testing the Go backend.

## Allowed Packages

- Only the standard Go packages are allowed for use in this project.

## Usage

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/egathang/groupie-tracker
    ```

2. Navigate to the project directory:

    ```bash
    cd groupie-tracker
    ```

3. Start the Go server:

    ```bash
    go run main.go
    ```

4. Open the frontend in your browser and explore the features.

## Example API

You can check an example of a RESTful API [here](https://jsonplaceholder.typicode.com/).

## Key Learning Outcomes

This project will help you gain experience in the following areas:

- Manipulating and storing data.
- Working with **JSON** files and formats.
- Designing **HTML** pages for displaying data.
- Creating and displaying events or actions.
- Implementing **client-server** interactions using Go.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


