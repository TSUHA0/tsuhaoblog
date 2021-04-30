(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["AddArt"],{"0898":function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("h3",[t._v(t._s(t.Add?"新建文章":"编辑文章"))]),a("a-card",[a("a-form-model",{ref:"artInfoRef",attrs:{model:t.artInfo,rules:t.artInfoRules}},[a("a-row",{attrs:{gutter:16}},[a("a-col",{attrs:{span:16}},[a("a-form-model-item",{attrs:{label:"文章标题",prop:"title"}},[a("a-input",{model:{value:t.artInfo.title,callback:function(e){t.$set(t.artInfo,"title",e)},expression:"artInfo.title"}})],1)],1)],1),a("a-row",{attrs:{gutter:16}},[a("a-col",{attrs:{lg:{span:6},xs:{span:8}}},[a("a-form-model-item",{attrs:{label:"文章分类",prop:"cid"}},[a("a-select",{attrs:{placeholder:"请选择分类"},on:{change:t.CateChange},model:{value:t.artInfo.cid,callback:function(e){t.$set(t.artInfo,"cid",e)},expression:"artInfo.cid"}},t._l(t.Catelist,(function(e){return a("a-select-option",{key:e.id,attrs:{value:e.id}},[t._v(" "+t._s(e.name)+" ")])})),1)],1)],1)],1),a("a-row",{attrs:{gutter:16}},[a("a-col",{attrs:{span:16}},[a("a-form-model-item",{attrs:{label:"文章描述",prop:"desc"}},[a("a-input",{attrs:{type:"textarea"},model:{value:t.artInfo.desc,callback:function(e){t.$set(t.artInfo,"desc",e)},expression:"artInfo.desc"}})],1)],1)],1),a("a-form-model-item",{attrs:{label:"文章缩略图",prop:"img"}},[a("a-upload",{attrs:{"list-type":"picture",name:"file",action:t.upUrl,headers:t.headers},on:{change:t.upChange}},[a("a-button",[a("a-icon",{attrs:{type:"upload"}}),t._v(" Upload ")],1)],1),a("img",{attrs:{src:"artInfo.img"}})],1),a("a-form-model-item",{attrs:{label:"文章内容"}},[a("mavon-editor",{ref:"md",on:{imgAdd:t.$imgAdd},model:{value:t.artInfo.content,callback:function(e){t.$set(t.artInfo,"content",e)},expression:"artInfo.content"}})],1),a("a-form-model-item",{staticClass:"btn"},[a("a-button",{staticStyle:{"margin-right":"15px"},attrs:{type:"danger"},on:{click:t.submitClk}},[t._v("提交")])],1)],1)],1)],1)},n=[],s=a("1da1"),o=(a("96cf"),a("b64b"),a("5bf0")),i=a("bc3a"),u=a.n(i),c={data:function(){return{artInfo:{id:0,title:"",cid:void 0,desc:"",content:"",img:""},Add:!1,CurrCateName:"请选择分类",Catelist:[],headers:{},upUrl:o["a"]+"upload",token:"Bearer ".concat(window.sessionStorage.getItem("token")),artInfoRules:{title:[{required:!0,message:"请输入文章标题",trigger:"change"}],cid:[{required:!0,message:"请选择文章分类",trigger:"change"}],desc:[{required:!0,message:"请输入文章描述",trigger:"change"},{max:120,message:"描述最多可写120个字符",trigger:"change"}]}}},mounted:function(){this.getCateList(),this.Add=Object.keys(this.$route.params).indexOf("id")<0,this.headers={Authorization:this.token},this.$route.params.id&&this.getArtInfo(this.$route.params.id)},methods:{CateChange:function(t){this.artInfo.cid=t},$imgAdd:function(t,e){var a=this;return Object(s["a"])(regeneratorRuntime.mark((function r(){var n,s;return regeneratorRuntime.wrap((function(r){while(1)switch(r.prev=r.next){case 0:n=a.$refs.md,s=new FormData,s.append("file",e),u()({url:o["a"]+"upload",method:"post",data:s,headers:{"Content-Type":"multipart/form-data",Authorization:a.token}}).then((function(e){n.$img2Url(t,e.data.url)}));case 4:case"end":return r.stop()}}),r)})))()},getArtInfo:function(t){var e=this;return Object(s["a"])(regeneratorRuntime.mark((function a(){var r,n;return regeneratorRuntime.wrap((function(a){while(1)switch(a.prev=a.next){case 0:return a.next=2,e.$http.get("art/".concat(t));case 2:r=a.sent,n=r.data,200!==n.status&&(1004!==n.status&&1005!==n.status&&1006!==n.status&&1007!==n.status||(window.sessionStorage.clear(),e.$router.push("/login")),e.$message.error(n.message)),e.artInfo=n.data,e.artInfo.id=n.data.ID;case 7:case"end":return a.stop()}}),a)})))()},getCateList:function(){var t=this;return Object(s["a"])(regeneratorRuntime.mark((function e(){var a,r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,t.$http.get("cates");case 2:a=e.sent,r=a.data,200!==r.status&&(1004!==r.status&&1005!==r.status&&1006!==r.status&&1007!==r.status||(window.sessionStorage.clear(),t.$router.push("/login")),t.$message.error(r.message)),t.Catelist=r.data;case 6:case"end":return e.stop()}}),e)})))()},submitClk:function(){var t=this;this.$refs.artInfoRef.validate(function(){var e=Object(s["a"])(regeneratorRuntime.mark((function e(a){var r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(a){e.next=2;break}return e.abrupt("return",t.$message.error("参数验证未通过，请按要求录入文章内容"));case 2:if(!t.Add){e.next=8;break}return e.next=5,t.$http.post("art/add",t.artInfo);case 5:r=e.sent,e.next=11;break;case 8:return e.next=10,t.$http.put("art/".concat(t.$route.params.id),t.artInfo);case 10:r=e.sent;case 11:if(200===r.status){e.next=13;break}return e.abrupt("return",t.$message.error(r.message));case 13:t.$router.push("/admin/artlist"),t.$message.success("添加文章成功");case 15:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}())},upChange:function(t){if("done"===t.file.status){var e=t.file.response.status;if(1004===e||1005===e||1006===e||1007===e)return this.$message.error(t.file.response.message);this.$message.success("图片上传成功"),this.artInfo.img=t.file.response.url}else"error"===t.file.status&&this.$message.error("图片上传失败")}}},d=c,l=a("2877"),f=Object(l["a"])(d,r,n,!1,null,null,null);e["default"]=f.exports},b64b:function(t,e,a){var r=a("23e7"),n=a("7b0b"),s=a("df75"),o=a("d039"),i=o((function(){s(1)}));r({target:"Object",stat:!0,forced:i},{keys:function(t){return s(n(t))}})}}]);
//# sourceMappingURL=AddArt.954808a8.js.map