let QApi = {
    getCollections: function (body,callback) {
        QNet.get("/api-collections",body, callback)
    },
    createCollections: function (body,callback) {
        QNet.post("/api-collections", body, callback)
    },
    modifyCollections: function (body,callback) {
        QNet.put("/api-collections", body, callback)
    },
    deleteCollections: function (body,callback) {
        QNet.delete("/api-collections", body, callback)
    }
};