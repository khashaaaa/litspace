-- Confirmation Type
CREATE TYPE confirmation AS ENUM (
    'UNCERTAIN',
    'CONFIRMED',
    'PENDING',
    'REJECTED'
);

-- User Type
CREATE TYPE figure AS ENUM (
    'CONSUMER',
    'MERCHANT',
    'PROVIDER'
);

-- Product Order Status
CREATE TYPE porder_status AS ENUM (
    'PROVISIONED',
    'SHIPPED',
    'ONTRACK',
    'DELIVERED'
);

-- Service Order Status
CREATE TYPE sorder_status AS ENUM (
    'ONPROCESS',
    'ACCEPTED',
    'REJECTED',
    'EXTENDED'
);

-- Card Type
CREATE TYPE card_network AS ENUM (
    'VISA',
    'MASTERCARD',
    'CHINA UNIONPAY',
    'JCB',
    'AMERICAN EXPRESS'
);

-- Consumer
CREATE TABLE IF NOT EXISTS "consumer" (
    mark UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    last_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    mobile VARCHAR,
    origin_country VARCHAR,
    pass VARCHAR NOT NULL,
    type figure DEFAULT 'CONSUMER',
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP
);

-- Payment Method
CREATE TABLE IF NOT EXISTS "paymethod" (
    mark UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    owner UUID NOT NULL,
    holder_name VARCHAR NOT NULL,
    card_number VARCHAR NOT NULL,
    expiry VARCHAR NOT NULL,
    cvv VARCHAR NOT NULL,
    region VARCHAR NOT NULL,
    card_provider card_network NOT NULL
);

/*
    Product related fields =================================================
*/
-- Product Category
CREATE TABLE IF NOT EXISTS "pcategory" (
    mark SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP
);

-- Product Sub Category
CREATE TABLE IF NOT EXISTS "psubcategory" (
    mark SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    category INT NOT NULL,
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP,
    FOREIGN KEY (category) REFERENCES pcategory(mark) ON DELETE CASCADE
);

-- Merchant
CREATE TABLE IF NOT EXISTS "merchant" (
    mark UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    founder UUID NOT NULL,
    entity_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    mobile VARCHAR NOT NULL,
    address VARCHAR,
    origin_country VARCHAR,
    buy_dest VARCHAR,
    sell_dest VARCHAR,
    in_status confirmation DEFAULT 'UNCERTAIN',
    type figure DEFAULT 'MERCHANT',
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP,
    FOREIGN KEY(founder) REFERENCES consumer(mark)
);

-- Product
CREATE TABLE IF NOT EXISTS "product" (
    mark UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    merchant UUID NOT NULL,
    category INT NOT NULL,
    name VARCHAR NOT NULL UNIQUE,
    descr TEXT,
    price INT NOT NULL,
    stock INT NOT NULL,
    attrs JSON,
    image_paths VARCHAR[],
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP,
    FOREIGN KEY (category) REFERENCES psubcategory(mark),
    FOREIGN KEY (merchant) REFERENCES merchant(mark)
);

CREATE TABLE IF NOT EXISTS "orderitems" (
    mark SERIAL PRIMARY KEY,
    p_mark UUID NOT NULL,
    p_name VARCHAR NOT NULL,
    p_quantity INT NOT NULL,
    p_price INT NOT NULL,
    seller UUID NOT NULL,
    buyer UUID NOT NULL
);

CREATE TABLE IF NOT EXISTS "porder" (
    mark UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    items INT NOT NULL,
    sum INT NOT NULL,
    amount INT NOT NULL,
    in_status porder_status,
    issued TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (items) REFERENCES ordered_items(mark)
);

/*
    Service related fields ===================================================
*/
-- Service Category
CREATE TABLE IF NOT EXISTS "scategory" (
    mark SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP
);

-- Service Sub Category
CREATE TABLE IF NOT EXISTS "ssubcategory" (
    mark SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    category INT NOT NULL,
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP,
    FOREIGN KEY (category) REFERENCES scategory(mark) ON DELETE CASCADE
);

-- Provider
CREATE TABLE IF NOT EXISTS "provider" (
    mark UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    founder UUID NOT NULL,
    entity_name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    mobile VARCHAR NOT NULL,
    address VARCHAR,
    origin_country VARCHAR,
    in_status confirmation DEFAULT 'UNCERTAIN',
    type figure DEFAULT 'PROVIDER',
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP,
    FOREIGN KEY (founder) REFERENCES consumer(mark)
);

-- Service
CREATE TABLE IF NOT EXISTS "service" (
    mark UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    provider UUID NOT NULL,
    category INT NOT NULL,
    title VARCHAR NOT NULL,
    descr JSON NOT NULL,
    cost_from INT NOT NULL,
    cost_up INT NOT NULL,
    opened VARCHAR NOT NULL,
    closed VARCHAR NOT NULL,
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated TIMESTAMP,
    FOREIGN KEY (provider) REFERENCES provider(mark),
    FOREIGN KEY (category) REFERENCES ssubcategory(mark)
);

-- Service Order
CREATE TABLE IF NOT EXISTS "sorder" (
    mark UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    service UUID NOT NULL,
    amount INT NOT NULL,
    demandant UUID NOT NULL,
    executor UUID NOT NULL,
    demandant_type figure DEFAULT 'CONSUMER',
    executor_type figure DEFAULT 'PROVIDER',
    issued TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY (service) REFERENCES service(mark)
);

-- Review
CREATE TABLE IF NOT EXISTS "review" (
    mark SERIAL PRIMARY KEY,
    subjective UUID NOT NULL,
    objective UUID NOT NULL,
    comment TEXT,
    rate_1 INT,
    rate_2 INT,
    rate_3 INT,
    rate_4 INT,
    rate_5 INT,
    rate_6 INT,
    rate_7 INT,
    rate_8 INT,
    rate_9 INT,
    rate_10 INT,
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);