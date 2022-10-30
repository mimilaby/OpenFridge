CREATE TABLE "food" (
    "id" UUID NOT NULL DEFAULT (uuid_generate_v4()),
    "name" VARCHAR NOT NULL,
    "description" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "calories" NUMERIC NOT NULL,
    "weight" NUMERIC NOT NULL,
    "amount" NUMERIC NOT NULL,
    "weight_per_amount" NUMERIC NOT NULL,
    "photo" VARCHAR NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "food_pkey" PRIMARY KEY ("id")
);