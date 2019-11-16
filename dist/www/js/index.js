let collectionList = NewList({
    id:"group-list",
    text:"name",
    children: "children"
});
collectionList.setData([
        {
            name: 'API管理',
            code: "api",
        },
        {
            name: '用户管理',
            code: "user",
        }
 ]);

collectionList.onClick = function (data) {
    console.log(data);
    history.pushState(data, data.name, data.page)
};

history.QOnChange = function () {
    console.log(history);
    let body_body = Q("#body-body");
    let data = history.QData;
    body_body.setAttribute("src", "/modules/api/index.html");
    Q("#body-head").innerHTML = data.name
};

history.QInit(collectionList.data[0]);
function createGroup() {
    
}