let paths = location.pathname.split("/");

let projectInfo = {
    projectCode: "unknown",
    module: "unknown",
};
if (paths.length==3){
    projectInfo = {
        projectCode:paths[1],
        module:paths[2]
    }
}

history.QOnChange = function () {
    console.log(history);
    let project_body = Q("#project-body");
    let data = history.QData;
    project_body.setAttribute("src", "/modules/" + data.module + "/index.html");
};

Q("a").forEach(function (a) {
    // a.setAttribute("href","/"+paths[1]+a.getAttribute("href"))
    a.addEventListener("click",function (e) {
        e.preventDefault()
        history.pushState({
            projectCode:projectInfo.projectCode,
            module:a.QData("module")
        },null,"/"+projectInfo.projectCode+"/"+a.QData("module"))
    })
});

history.QInit(projectInfo);
