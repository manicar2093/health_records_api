
exports.up = function(knex) {

    const weight_clasification = () => knex.schema.createTable('WeightClasifications', t => {
        t.increments('id').primary();
        t.string('description').notNullable();
    });

    const biotest = () => knex.schema.createTable('Biotest', t => {
        t.increments('id').primary();
        t.integer('higher_muscle_density_id').notNullable().references('HigherMuscleDensity.id');
        t.integer('lower_muscle_density_id').notNullable().references('LowerMuscleDensity.id');
        t.integer('skin_folds_id').notNullable().references('SkinFolds.id');
        t.integer('weight_clasification_id').notNullable().references('WeightClasifications.id');
        t.decimal('weight').notNullable();
        t.decimal('height').notNullable();
        t.decimal('body_fat_percentage').notNullable();
        t.decimal('total_body_water').notNullable();
        t.decimal('body_mass_index').notNullable();
        t.decimal('oxygen_saturation_in_blood').notNullable();
        t.decimal('glucose').nullable();
        t.decimal('resting_heart_rate').nullable();
        t.decimal('maximum_heart_rate').nullable();
        t.string('heart_health').nullable();
        t.string('observations');
        t.string('recommendations');
        t.date('created_at').notNullable().defaultTo(knex.fn.now());
    });

    return weight_clasification()
        .then(biotest);
};

exports.down = function(knex) {
    const weight_clasification = () => knex.schema.dropTable('WeightClasifications');
    const biotest = () => knex.schema.dropTable('Biotest');

    return biotest()
        .then(weight_clasification);
};
