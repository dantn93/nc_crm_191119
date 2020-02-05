-- DROP TABLE IF EXISTS "public"."truck_rate_card_level";
-- DROP TABLE IF EXISTS "public"."truck_level";
-- DROP TABLE IF EXISTS "public"."truck_type";
-- DROP TABLE IF EXISTS "public"."truck_rate_card";
-- DROP TABLE IF EXISTS "public"."customer_contract";
/*
 Navicat Premium Data Transfer

 Source Server         : zlocalhost
 Source Server Type    : PostgreSQL
 Source Server Version : 110005
 Source Host           : localhost:5432
 Source Catalog        : mytruck
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 110005
 File Encoding         : 65001

 Date: 08/01/2020 09:51:53
*/


-- ----------------------------
-- Table structure for customer_contract
-- ----------------------------
-- DROP TABLE IF EXISTS "public"."customer_contract";
CREATE TABLE "public"."customer_contract" (
  "id" int4 NOT NULL,
  "customer_id" int4,
  "is_active" bool,
  "created_at" timestamp(6),
  "updated_at" timestamp(6),
  "created_by" int4,
  "updated_by" int4,
  "contract_type" varchar(255) COLLATE "pg_catalog"."default",
  "expired_date" timestamp(6),
  "return_ratio" float4,
  "cod_ratio" float4,
  "stop_fee" float4,
  "paper_fee" float4,
  "lift_fee" float4,
  "check_fee" float4,
  "value_ratio" float4
)
;
ALTER TABLE "public"."customer_contract" OWNER TO "postgres";

-- ----------------------------
-- Records of customer_contract
-- ----------------------------
BEGIN;
INSERT INTO "public"."customer_contract" VALUES (1, 1, 't', '2019-10-11 21:17:50', '2019-10-11 21:17:53', 100, 100, 'Public', '2019-10-26 21:18:03', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO "public"."customer_contract" VALUES (100, 12313, 't', NULL, NULL, NULL, NULL, 'Private', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO "public"."customer_contract" VALUES (101, 12312, NULL, NULL, NULL, NULL, NULL, 'Private', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for truck_level
-- ----------------------------
-- DROP TABLE IF EXISTS "public"."truck_level";
CREATE TABLE "public"."truck_level" (
  "id" int4 NOT NULL,
  "rate" int4,
  "code" varchar(255) COLLATE "pg_catalog"."default",
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "tag" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."truck_level" OWNER TO "postgres";

-- ----------------------------
-- Records of truck_level
-- ----------------------------
BEGIN;
INSERT INTO "public"."truck_level" VALUES (1, 0, 'level1', '>0', 'public');
INSERT INTO "public"."truck_level" VALUES (2, 10, 'level2', '>10', 'public');
INSERT INTO "public"."truck_level" VALUES (3, 100, 'level3', '>100', 'public');
INSERT INTO "public"."truck_level" VALUES (4, 500, 'level4', '>500', 'public');
INSERT INTO "public"."truck_level" VALUES (5, 0, 'level1', '>0', 'customer1');
INSERT INTO "public"."truck_level" VALUES (6, 20, 'level2', '>20', 'customer1');
INSERT INTO "public"."truck_level" VALUES (7, 100, 'level3', '>100', 'customer1');
INSERT INTO "public"."truck_level" VALUES (8, 500, 'level4', '>500', 'customer1');
COMMIT;

-- ----------------------------
-- Table structure for truck_rate_card
-- ----------------------------
-- DROP TABLE IF EXISTS "public"."truck_rate_card";
CREATE TABLE "public"."truck_rate_card" (
  "id" int4 NOT NULL,
  "contract_id" int4,
  "truck_type" int4,
  "updated_at" timestamp(6) DEFAULT now(),
  "created_at" timestamp(6) DEFAULT now(),
  "created_by" int4 DEFAULT 11667
)
;
ALTER TABLE "public"."truck_rate_card" OWNER TO "postgres";

-- ----------------------------
-- Records of truck_rate_card
-- ----------------------------
BEGIN;
-- INSERT INTO "public"."truck_rate_card" VALUES (5, 100, 1500, '2019-10-11 21:24:42', '2019-10-11 21:24:42', 33454);
INSERT INTO "public"."truck_rate_card" VALUES (4, 1, 8000, '2019-10-11 21:24:42', '2019-10-11 21:24:42', 77534);
-- INSERT INTO "public"."truck_rate_card" VALUES (1, 1, 1100, '2019-10-11 21:24:42', '2019-10-11 21:24:42', 77534);
INSERT INTO "public"."truck_rate_card" VALUES (2, 1, 1500, '2019-10-11 21:24:42', '2019-10-11 21:24:42', 77534);
INSERT INTO "public"."truck_rate_card" VALUES (3, 1, 5000, '2019-10-11 21:24:42', '2019-10-11 21:24:42', 77534);
COMMIT;

-- ----------------------------
-- Table structure for truck_rate_card_level
-- ----------------------------
-- DROP TABLE IF EXISTS "public"."truck_rate_card_level";
CREATE TABLE "public"."truck_rate_card_level" (
  "rate_card_id" int4,
  "level_id" int4,
  "price" float4,
  "id" int4 NOT NULL
)
;
ALTER TABLE "public"."truck_rate_card_level" OWNER TO "postgres";

-- ----------------------------
-- Records of truck_rate_card_level
-- ----------------------------
BEGIN;
INSERT INTO "public"."truck_rate_card_level" VALUES (1, 2, 1000, 2);
INSERT INTO "public"."truck_rate_card_level" VALUES (1, 3, 900, 3);
INSERT INTO "public"."truck_rate_card_level" VALUES (1, 1, 2000, 1);
INSERT INTO "public"."truck_rate_card_level" VALUES (1, 4, 800, 4);
COMMIT;

-- ----------------------------
-- Table structure for truck_type
-- ----------------------------
-- DROP TABLE IF EXISTS "public"."truck_type";
CREATE TABLE "public"."truck_type" (
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "weight" int4 NOT NULL
)
;
ALTER TABLE "public"."truck_type" OWNER TO "postgres";

-- ----------------------------
-- Records of truck_type
-- ----------------------------
BEGIN;
INSERT INTO "public"."truck_type" VALUES ('1T5', 1500);
INSERT INTO "public"."truck_type" VALUES ('8T', 8000);
INSERT INTO "public"."truck_type" VALUES ('16T', 16000);
INSERT INTO "public"."truck_type" VALUES ('30T', 30000);
INSERT INTO "public"."truck_type" VALUES ('1T1', 1100);
INSERT INTO "public"."truck_type" VALUES ('5T', 5000);
COMMIT;

-- ----------------------------
-- Primary Key structure for table customer_contract
-- ----------------------------
ALTER TABLE "public"."customer_contract" ADD CONSTRAINT "customer_contract_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table truck_level
-- ----------------------------
ALTER TABLE "public"."truck_level" ADD CONSTRAINT "truck_level_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table truck_rate_card
-- ----------------------------
ALTER TABLE "public"."truck_rate_card" ADD CONSTRAINT "truck_rate_card_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table truck_rate_card_level
-- ----------------------------
ALTER TABLE "public"."truck_rate_card_level" ADD CONSTRAINT "truck_rate_card_level_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table truck_type
-- ----------------------------
ALTER TABLE "public"."truck_type" ADD CONSTRAINT "truck_type_pkey" PRIMARY KEY ("weight");

-- ----------------------------
-- Foreign Keys structure for table truck_rate_card
-- ----------------------------
ALTER TABLE "public"."truck_rate_card" ADD CONSTRAINT "truck_rate_card_contract_id_fkey" FOREIGN KEY ("contract_id") REFERENCES "public"."customer_contract" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."truck_rate_card" ADD CONSTRAINT "truck_rate_card_truck_type_fkey" FOREIGN KEY ("truck_type") REFERENCES "public"."truck_type" ("weight") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table truck_rate_card_level
-- ----------------------------
ALTER TABLE "public"."truck_rate_card_level" ADD CONSTRAINT "truck_rate_card_level_level_id_fkey" FOREIGN KEY ("level_id") REFERENCES "public"."truck_level" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."truck_rate_card_level" ADD CONSTRAINT "truck_rate_card_level_rate_card_id_fkey" FOREIGN KEY ("rate_card_id") REFERENCES "public"."truck_rate_card" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
