function NewList(custom) {
    function NewDefaultConfig() {
        return{
            text:'text',
            children:'children'
        }
    }

    let dom = Q("#"+custom.id);
    dom.QAddClass("q-list");
    let body = QCreateElement("div",dom,null,"q-list-body");

    function addChildren(children,dom,parentKey,config,list) {
        parentKey = parentKey||"";
        children.forEach(function (node,i) {
            let nodeDiv = QCreateElement("div",dom,undefined,"node");
            let key = parentKey+"."+i
            nodeDiv.QData("key",key);
            let nodeHead = QCreateElement("div",nodeDiv,undefined,"node-head");
            nodeHead.addEventListener("click",function (event) {
                list.onclick(event)
            });
            let nodeBody = QCreateElement("div",nodeDiv,undefined,"node-body");
            nodeHead.innerHTML = node[config.text]
            if (node[config.children]&&node[config.children].length>0)
                addChildren(node[config.children],nodeBody,key,config,list)
        })
    }

    return {
        config:custom,
        data: null,
        dom: dom,
        onclick:function (event) {
            let node = event.target.QParent();
            let keys = node.QData("key").split('.');
            let cf = this.config.children;
            let data = {}
            data[cf] = this.data;
            for(let i = 1;i<keys.length;i++){
                data = data[cf][keys[i]]
            }
            this.onClick(data)
        },
        onClick:function (data) {},
        setData(data) {
            this.data = data;
            Q(".q-list-body",this.dom).innerHTML = "";
            addChildren(data,Q(".q-list-body",this.dom),null,this.config,this)
        }
    }
}
