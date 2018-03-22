var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator.throw(value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : new P(function (resolve) { resolve(result.value); }).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments)).next());
    });
};
function getCookie(key) {
    var cookie = document.cookie;
    if (cookie == null || cookie == "")
        return "";
    var regex = new RegExp("(?:" + key + "=)([^&;]+)", "gi");
    if (regex.test(cookie))
        return (RegExp.$1);
    return '';
}
function isEmail(email) {
    var regex = /^([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|\_|\.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,3}$/;
    return regex.test(email);
}
function queryString() {
    let urlItems = window.location.search.split("?");
    let values = new Map();
    if (urlItems.length == 0) {
        return values;
    }
    let params = urlItems[urlItems.length - 1];
    let result = params.split("&").reduce((rt, item) => {
        let [key, val] = item.split("=");
        rt.set(key, decodeURIComponent(val));
        return rt;
    }, values);
    return result;
}
function ajaxAsync(url = "", data, method = "GET", dataType = "json") {
    return __awaiter(this, void 0, void 0, function* () {
        return new Promise((resolve, reject) => {
            let ajaxSetting = {};
            ajaxSetting.url = url;
            ajaxSetting.type = method;
            ajaxSetting.dataType = dataType;
            ajaxSetting.data = data;
            ajaxSetting.success = (response) => {
                resolve(response);
            };
            ajaxSetting.error = (xhr) => {
                reject("请求出错，请联系网站管理员");
            };
            $.ajax(ajaxSetting);
        });
    });
}
function mapToObject(map) {
    let result = {};
    for (let [key, val] of map.entries()) {
        result[key] = val;
    }
    return result;
}
function jquery2HtmlElements(jquery) {
    let result = [];
    jquery.each((index, item) => result.push(item));
    return result;
}
function isInteger(val) {
    return /^\-?\d+$/.test(val);
}
function syntaxHighlight(json) {
    if (typeof json != 'string') {
        json = JSON.stringify(json, undefined, 2);
    }
    json = json.replace(/&/g, '&').replace(/</g, '<').replace(/>/g, '>');
    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
        var cls = 'number';
        if (/^"/.test(match)) {
            if (/:$/.test(match)) {
                cls = 'key';
            }
            else {
                cls = 'string';
            }
        }
        else if (/true|false/.test(match)) {
            cls = 'boolean';
        }
        else if (/null/.test(match)) {
            cls = 'null';
        }
        return '<span class="' + cls + '">' + match + '</span>';
    });
}
function getLength(str) {
    var len = 0;
    var arr = str.split('');
    for (var i = 0; i < arr.length; i++) {
        var item = arr[i];
        if (/[\u4e00-\u9fa5]/.test(item)) {
            len += 2;
        }
        else {
            len += 1;
        }
    }
    return len;
}
//# sourceMappingURL=functions.js.map