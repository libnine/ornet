con = new Mongo()
db = con.getDB("ornet");
cur = db.users.find();

while (cur.hasNext() ) {
    printjson(cur.next());
}

var x = new ISODate("")

db.users.insert({
    case_date: ""
})