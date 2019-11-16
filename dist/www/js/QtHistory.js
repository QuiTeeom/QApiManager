history["originalPushState"] = history.pushState;
history["pushState"] = function (data, title, url) {
    history.originalPushState(data, title, url);
    history.QData = data;
    history.QOnChange()
};
history["QData"] = null;

window.addEventListener("popstate", function (event) {
    event.preventDefault()
    history.QOnChange()
}, false);

history['QOnChange'] = function () {
    console.log("变化了");
};

history["QInit"] = function (data) {
    history.QData = data;
    history.QOnChange()
};
