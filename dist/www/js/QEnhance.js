function Q(selector,c) {
    c = c||document
    let doms = c.querySelectorAll(selector);
    if (!doms||doms.length==0)
        return null

    if (doms.length==1)
        return QEnhanceDom(doms[0]);
    else {
        doms.forEach(dom => QEnhanceDom)
        return doms;
    }
}

function QCreateElement(tag,parent,css,className) {
    let dom = QEnhanceDom(document.createElement(tag));

    if (parent)
        parent.appendChild(dom);
    if (css)
        dom.QSetCss(dom,css);
    if (className)
        dom.QAddClass(className);
    return dom
}

function QEnhanceDom(dom) {
    if (dom.QEnhance)
        return dom;

    dom["QEnhance"] = true;
    dom["QData"] = function (key,data) {
        if (data!=null)
            this.setAttribute("data-"+key,data);
        else
            return this.getAttribute("data-"+key)
    };
    dom["QSetCss"] = function(css) {
        this.style = css;
    };
    dom["QParent"] = function () {
        return QEnhanceDom(dom.parentNode);
    };
    dom["QAddClass"] = function (...className) {
        let c = this.getAttribute("class")||"";
        let cs = c.split("\s{1,}");
        className.forEach(function (value) {
            if (cs.indexOf(value)<0) {
                cs.push(value)
            }
        });
        this.setAttribute("class",cs.join(" "))
    };
    dom["QRemoveClass"] = function (...className) {
        let c = this.getAttribute("class")||"";
        let cs = c.split("\s{1,}");
        let fcs = [];
        cs.forEach(function (value) {
            if (className.indexOf(value)>=0) {
                fcs.push(value)
            }
        });
        this.setAttribute("class",cs.join(" "))
    };
    return dom
}