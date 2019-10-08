use ornet
var d = new Date(ISODate().getTime() - 1000 * 86400 * 3)
db.cases.find({ case_date: { $lt: d }}).forEach(function(d) {
    db.cases_historical.insert(d)
})
db.cases.remove({ case_date: { $lt: d }})
