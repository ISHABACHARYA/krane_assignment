const knex = require("knex");
const config = require("./knexfile"); // Assuming your Knex configuration is in a separate file

// Create a Knex instance based on your configuration
const db = knex(config.development);

// Check database connection
async function checkDatabaseConnection() {
  try {
    // Execute a simple SQL query to test the connection
    const result = await db.raw("SELECT 1 + 1 AS result");
    console.log("Database connection successful:", result.rows[0].result === 2);
  } catch (error) {
    console.error("Database connection failed:", error.message);
  } finally {
    // Close the database connection
    await db.destroy();
  }
}

// Call the function to check the database connection
checkDatabaseConnection();
