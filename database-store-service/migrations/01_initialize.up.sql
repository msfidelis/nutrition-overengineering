CREATE TABLE IF NOT EXISTS health_report (
    id VARCHAR ( 255 ) PRIMARY KEY,
    imc REAL NOT NULL,
    imc_class VARCHAR (255) NOT NULL,
    bmr  REAL NOT NULL,
    bmr_unity VARCHAR (20) NOT NULL DEFAULT 'kcal',
    necessity REAL NOT NULL,
    necessity_unity VARCHAR(20) NOT NULL DEFAULT 'kcal',
    age INT NOT NULL, 
    weight REAL NOT NULL, 
    height REAL NOT NULL, 
    gender VARCHAR (1),
    activity_insensity VARCHAR (20) NOT NULL,
    recommendation_protein REAL NOT NULL,
    recommendation_protein_unit VARCHAR (20) NOT NULL DEFAULT 'kcal', 
    recommendation_water REAL NOT NULL,
    recommendation_water_unit VARCHAR (20) NOT NULL DEFAULT 'ml', 
    recommendation_calores_maintain REAL NOT NULL,
    recommendation_calores_maintain_unit VARCHAR (20) NOT NULL DEFAULT 'kcal',
    recommendation_calores_loss REAL NOT NULL,
    recommendation_calores_loss_unit VARCHAR (20) NOT NULL DEFAULT 'kcal',    
    recommendation_calores_gain REAL NOT NULL,
    recommendation_calores_gain_unit VARCHAR (20) NOT NULL DEFAULT 'kcal'             
);