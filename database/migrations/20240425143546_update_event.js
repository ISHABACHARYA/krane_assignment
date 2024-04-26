/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.schema.alterTable("Event", function (table) {
    table.dropColumn("sessions");
  });
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function (knex) {
  return knex.schema.alterTable("Event", function (table) {
    table.specificType("sessions", "TEXT[]"); // Recreate the sessions field as it was before
  });
};
