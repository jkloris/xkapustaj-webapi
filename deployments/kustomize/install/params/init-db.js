const mongoHost = process.env.AMBULANCE_API_MONGODB_HOST
const mongoPort = process.env.AMBULANCE_API_MONGODB_PORT

const mongoUser = process.env.AMBULANCE_API_MONGODB_USERNAME
const mongoPassword = process.env.AMBULANCE_API_MONGODB_PASSWORD

const database = process.env.AMBULANCE_API_MONGODB_DATABASE
const collection = process.env.AMBULANCE_API_MONGODB_COLLECTION

const retrySeconds = parseInt(process.env.RETRY_CONNECTION_SECONDS || "5") || 5;

// try to connect to mongoDB until it is not available
let connection;
while(true) {
    try {
        connection = Mongo(`mongodb://${mongoUser}:${mongoPassword}@${mongoHost}:${mongoPort}`);
        break;
    } catch (exception) {
        print(`Cannot connect to mongoDB: ${exception}`);
        print(`Will retry after ${retrySeconds} seconds`)
        sleep(retrySeconds * 1000);
    }
}

// if database and collection exists, exit with success - already initialized
const databases = connection.getDBNames()
if (databases.includes(database)) {
    const dbInstance = connection.getDB(database)
    collections = dbInstance.getCollectionNames()
    if (collections.includes(collection)) {
       print(`Collection '${collection}' already exists in database '${database}'`)
        process.exit(0);
    }
}

// initialize
// create database and collection
const db = connection.getDB(database)
db.createCollection(collection)

// create indexes
db[collection].createIndex({ "id": 1 })

//insert sample data
let result = db[collection].insertMany([
    {
        "id" : "bobulova",
        "name" : "Hospital Bobulova",
        "address" : "90091 Limbach, se 12",
        "employees" : [
            {
                "id" : "jpk",
                "name" : "Jozko Pucik",
                "jobTitle" : "Doctor",
                "patientId" : "32ds"
            },
            {
                "id" : "jp23k",
                "name" : "Marek Sykora",
                "jobTitle" : "Doctor",
                "patientId" : "32123ds"
            },
            {
                "id" : "mm23k",
                "name" : "Marek Medzny",
                "jobTitle" : "Primar",
                "patientId" : "3w3ds"
            }
        ],
        "timesheets" : [
            {
                "id" : "12321",
                "hours" : 8,
                "description" : "Paperwork",
                "date" : "2023-12-24T10:35:00Z",
                "employeeId" : "mm23k"
            },
            {
                "id" : "12322",
                "hours" : 5,
                "description" : "operation",
                "date" : "2023-12-25T10:35:00Z",
                "employeeId" : "jpk"
            },
            {
                "id" : "12323",
                "hours" : 7,
                "description" : "operation",
                "date" : "2023-12-26T10:35:00Z",
                "employeeId" : "jpk"
            },
            {
                "id" : "12324",
                "hours" : 11,
                "description" : "Paperwork",
                "date" : "2023-12-26T11:35:00Z",
            }
        ]
    }
]);

if (result.writeError) {
    console.error(result)
    print(`Error when writing the data: ${result.errmsg}`)
}

// exit with success
process.exit(0);