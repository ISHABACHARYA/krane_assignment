-- Create the User table
CREATE TABLE "User" (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phoneNumber VARCHAR(20),
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the Event table
CREATE TABLE "Event" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    startDate TIMESTAMP NOT NULL,
    endDate TIMESTAMP NOT NULL,
    sessions TEXT[], -- Array type for sessions
    createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create the UserEvent table
CREATE TABLE "UserEvent" (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    event_id INT NOT NULL,
    role VARCHAR(20) NOT NULL, -- Assuming role will be stored as a string
    FOREIGN KEY (user_id) REFERENCES "User"(id) ON DELETE CASCADE,
    FOREIGN KEY (event_id) REFERENCES "Event"(id) ON DELETE CASCADE
);

-- Create the Expense table
CREATE TABLE "Expense" (
    id SERIAL PRIMARY KEY,
    event_id INT NOT NULL,
    type VARCHAR(20) NOT NULL, -- Assuming type will be stored as a string
    amount INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    FOREIGN KEY (event_id) REFERENCES "Event"(id) ON DELETE CASCADE
);
