"use strict";(self["webpackChunksos_dash"]=self["webpackChunksos_dash"]||[]).push([[437],{67437:(e,t,a)=>{a.r(t),a.d(t,{default:()=>ie});var l=a(59835),s=a(86970),r=a(28339),o=a(25121),n=a(60499),u=a(94629),c=a(85118),i=a(20503);const d=(0,l.aZ)({__name:"RefreshRateSelect",setup(e){const t=(0,c.V)(),a=new i.Z,s=[5,20,40,60,120,180].map((e=>({label:e+"s",value:e}))),o=(0,l.Fl)((()=>{const e=(0,r.yj)();return"/overview"===e.path})),d=(0,l.Fl)({get:()=>t.getSelectedRefreshRate,set:e=>t.setSelectedRefreshRate(e)});return(0,l.bv)((()=>{d.value=a.getRefreshRate()})),(0,l.YP)(d,(e=>{a.setRefreshRate(e)})),(e,t)=>o.value?((0,l.wg)(),(0,l.j4)(u.Z,{key:0,modelValue:d.value,"onUpdate:modelValue":t[0]||(t[0]=e=>d.value=e),options:(0,n.SU)(s),label:e.$t("refreshRateSelect.selectRefreshRate"),class:"refresh-rate-select"},null,8,["modelValue","options","label"])):(0,l.kq)("",!0)}}),m=d,p=m,g={class:"row justify-start items-center"},v={class:"title-h4 top-bar-text"},f=(0,l.aZ)({__name:"TopBar",setup(e){const t=(0,o.QT)().t;function a(){const e=(0,r.yj)();return e.meta.title?t(e.meta.title.toString()):""}const n=(0,l.Fl)((()=>{const e=(0,r.yj)();return e.meta.icon?e.meta.icon.toString():""}));return(e,t)=>{const r=(0,l.up)("q-icon"),o=(0,l.up)("q-space"),u=(0,l.up)("q-toolbar"),c=(0,l.up)("q-header");return(0,l.wg)(),(0,l.j4)(c,{class:"bg-transparent q-px-xs q-py-sm"},{default:(0,l.w5)((()=>[(0,l.Wm)(u,null,{default:(0,l.w5)((()=>[(0,l._)("div",g,[(0,l.Wm)(r,{color:"primary",name:n.value,size:"md",class:"q-mr-sm"},null,8,["name"]),(0,l._)("div",v,(0,s.zw)(a()),1)]),(0,l.Wm)(o),(0,l.Wm)(p)])),_:1})])),_:1})}}});var h=a(16602),w=a(51663),_=a(22857),b=a(90136),y=a(69984),x=a.n(y);const U=f,q=U;x()(f,"components",{QHeader:h.Z,QToolbar:w.Z,QIcon:_.Z,QSpace:b.Z});var W=a(37747);const Z={class:"row justify-between items-center"},z=(0,l.aZ)({__name:"ProfileCard",setup(e){function t(){const e=new W.Z;e.logout()}return(e,a)=>{const r=(0,l.up)("q-btn"),o=(0,l.up)("q-tooltip");return(0,l.wg)(),(0,l.iD)("div",Z,[(0,l._)("div",null,[(0,l.Wm)(r,{dense:"",color:"primary",icon:"sym_s_settings",size:"sm",to:"/settings"}),(0,l.Wm)(o,{anchor:"bottom middle",class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,s.zw)(e.$t("profileCard.btnSettings")),1)])),_:1})]),(0,l._)("div",null,[(0,l.Wm)(r,{dense:"",color:"grey-7",icon:"sym_s_logout",size:"sm",onClick:a[0]||(a[0]=e=>t())},{default:(0,l.w5)((()=>[(0,l.Wm)(o,{anchor:"bottom middle",class:"bg-grey-7",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,s.zw)(e.$t("profileCard.btnLogout")),1)])),_:1})])),_:1})])])}}});var k=a(68879),Q=a(46858);const R=z,M=R;x()(z,"components",{QBtn:k.Z,QTooltip:Q.Z});var S=a(19302);const P={style:{"text-align":"center",padding:"0 8px"}},j=["src"],$={class:"column items-center"},F={class:"col",style:{"font-size":"12px"}},I=(0,l.aZ)({__name:"SideBarMenu",setup(e){const t=(0,o.QT)().t,a={right:"4px",borderRadius:"5px",backgroundColor:"#c97350",width:"5px",opacity:.75},u={right:"2px",borderRadius:"9px",backgroundColor:"#c97350",width:"9px",opacity:.2},c=(0,n.iH)(),i=(0,l.Fl)((()=>{const e=(0,S.Z)();return e.dark.isActive?"/_/assets/os-logo-dark.svg":"/_/assets/os-logo-light.svg"}));function d(){const e=(0,r.tv)().getRoutes();let a=[];return e.forEach((e=>{const l=e.children||[];l.forEach((e=>{var l,s,r,o,n,u,c,i,d,m;!1!==(null===(l=e.meta)||void 0===l?void 0:l.isMenuItem)&&a.push({title:null!==(r=t(`${null===(s=e.meta)||void 0===s?void 0:s.title}`))&&void 0!==r?r:"",icon:null!==(u=null===(n=null===(o=e.meta)||void 0===o?void 0:o.icon)||void 0===n?void 0:n.toString())&&void 0!==u?u:"",path:e.path,disabled:null!==(i=null===(c=e.meta)||void 0===c?void 0:c.disabled)&&void 0!==i&&i,useHref:null!==(m=null===(d=e.meta)||void 0===d?void 0:d.useHref)&&void 0!==m&&m})}))})),a}return(0,l.wF)((()=>{c.value=d()})),(e,r)=>{const o=(0,l.up)("router-link"),d=(0,l.up)("q-avatar"),m=(0,l.up)("q-tooltip"),p=(0,l.up)("q-item"),g=(0,l.up)("q-list"),v=(0,l.up)("q-scroll-area"),f=(0,l.up)("q-drawer"),h=(0,l.Q2)("ripple");return(0,l.wg)(),(0,l.j4)(f,{mini:!0,side:"left",bordered:"","show-if-above":"","mini-width":100},{default:(0,l.w5)((()=>[(0,l._)("div",P,[(0,l.Wm)(o,{to:"/"},{default:(0,l.w5)((()=>[(0,l._)("img",{src:i.value,alt:"Infinite OS",style:{"margin-top":"20px",height:"1.82rem"}},null,8,j)])),_:1})]),(0,l.Wm)(M,{class:"q-pt-md q-px-md"}),(0,l.Wm)(v,{"thumb-style":a,"bar-style":u,style:{height:"calc(100% - 100px)","margin-top":"10px"}},{default:(0,l.w5)((()=>[(0,l.Wm)(g,{style:{padding:"0 4px","text-align":"center","overflow-x":"hidden"}},{default:(0,l.w5)((()=>[((0,l.wg)(!0),(0,l.iD)(l.HY,null,(0,l.Ko)(c.value,((e,a)=>(0,l.wy)(((0,l.wg)(),(0,l.j4)(p,(0,l.dG)({key:a},e.useHref?{href:e.path}:{to:e.path},{disable:e.disabled,clickable:""}),{default:(0,l.w5)((()=>[(0,l._)("div",$,[(0,l.Wm)(d,{icon:e.icon,class:"icon-menu-bg"},null,8,["icon"]),(0,l._)("div",F,(0,s.zw)(e.title),1)]),e.disabled?((0,l.wg)(),(0,l.j4)(m,{key:0,anchor:"center end",class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,s.zw)((0,n.SU)(t)("sideBarMenu.disabled")),1)])),_:1})):(0,l.kq)("",!0)])),_:2},1040,["disable"])),[[h]]))),128))])),_:1})])),_:1})])),_:1})}}});var B=a(10906),C=a(66663),H=a(13246),T=a(490),L=a(61357),V=a(51136);const A=I,D=A;x()(I,"components",{QDrawer:B.Z,QScrollArea:C.Z,QList:H.Z,QItem:T.Z,QAvatar:L.Z,QTooltip:Q.Z}),x()(I,"directives",{Ripple:V.Z});var Y=a(61957),E=a(87178),G=a(16397);const K={class:"flex justify-end items-center"},O={class:"absolute-full flex flex-center"},J={class:"absolute-full flex flex-center"},N={class:"absolute-full flex flex-center"},X=(0,l.aZ)({__name:"FooterBar",setup(e){const t=(0,r.yj)(),a=(0,c.V)(),o=(0,G.n)(),u=new E.Z,i=(0,n.iH)(),d=(0,l.Fl)({get:()=>o.getSystemInfo,set:e=>o.setSystemInfo(e)}),m=(0,l.Fl)((()=>a.getSelectedRefreshRate)),p=(0,l.Fl)((()=>t.path));function g(e){return e<50?"green":e<80?"orange":"red"}function v(){i.value&&clearInterval(i.value),i.value=setInterval((()=>{u.getSystemInfo().then((e=>{d.value=e.data.body})).catch((e=>{console.error(e)}))}),1e3*m.value)}return(0,l.Ah)((()=>{clearInterval(i.value)})),(0,l.YP)(p,(()=>{clearInterval(i.value),"/overview"!==p.value&&v()})),(e,t)=>{const a=(0,l.up)("q-tooltip"),r=(0,l.up)("q-icon"),o=(0,l.up)("q-badge"),n=(0,l.up)("q-linear-progress"),u=(0,l.up)("q-footer");return(0,l.wg)(),(0,l.j4)(u,{bordered:"",class:"q-px-lg",style:{"min-height":"30px"}},{default:(0,l.w5)((()=>[(0,l.wy)((0,l._)("div",K,[(0,l.Wm)(r,{name:"sym_s_terminal",size:"1.618rem",class:"disabled q-mr-md"},{default:(0,l.w5)((()=>[(0,l.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,s.zw)(e.$t("footerBar.disabled")),1)])),_:1})])),_:1}),(0,l.Wm)(r,{name:"sym_s_speed",size:"sm",class:"q-mr-xs"}),(0,l.Wm)(n,{stripe:"",rounded:"",size:"20px",class:"q-mr-md",value:Math.trunc(d.value.currentUsage.cpuUsagePercent)/100,color:g(Math.trunc(d.value.currentUsage.cpuUsagePercent)),label:`${Math.trunc(d.value.currentUsage.cpuUsagePercent)}%`,style:{width:"100px"}},{default:(0,l.w5)((()=>[(0,l._)("div",O,[(0,l.Wm)(o,{color:"white","text-color":"grey-10",label:`${Math.trunc(d.value.currentUsage.cpuUsagePercent)}%`},null,8,["label"])]),(0,l.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,s.zw)(e.$t("footerBar.cpuUsage",{cpuUsage:Math.trunc(d.value.currentUsage.cpuUsagePercent)})),1)])),_:1})])),_:1},8,["value","color","label"]),(0,l.Wm)(r,{name:"sym_s_memory",size:"sm",class:"q-mr-xs"}),(0,l.Wm)(n,{stripe:"",rounded:"",class:"q-mr-md",size:"20px",value:Math.trunc(d.value.currentUsage.memUsagePercent)/100,color:g(Math.trunc(d.value.currentUsage.memUsagePercent)),label:`${Math.trunc(d.value.currentUsage.memUsagePercent)}%`,style:{width:"100px"}},{default:(0,l.w5)((()=>[(0,l._)("div",J,[(0,l.Wm)(o,{color:"white","text-color":"grey-10",label:`${Math.trunc(d.value.currentUsage.memUsagePercent)}%`},null,8,["label"])]),(0,l.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,s.zw)(e.$t("footerBar.ramUsage",{ramUsage:Math.trunc(d.value.currentUsage.memUsagePercent)})),1)])),_:1})])),_:1},8,["value","color","label"]),(0,l.Wm)(r,{name:"sym_s_sd_card",size:"sm",class:"q-mr-xs"}),(0,l.Wm)(n,{stripe:"",rounded:"",size:"20px",value:Math.trunc(d.value.currentUsage.storageUsage)/100,color:g(Math.trunc(d.value.currentUsage.storageUsage)),label:`${Math.trunc(d.value.currentUsage.storageUsage)}%`,style:{width:"100px"}},{default:(0,l.w5)((()=>[(0,l._)("div",N,[(0,l.Wm)(o,{color:"white","text-color":"grey-10",label:`${Math.trunc(d.value.currentUsage.storageUsage)}%`},null,8,["label"])]),(0,l.Wm)(a,{class:"bg-primary",style:{"font-size":"14px"}},{default:(0,l.w5)((()=>[(0,l.Uk)((0,s.zw)(e.$t("footerBar.storageUsage",{storageUsage:Math.trunc(d.value.currentUsage.storageUsage)})),1)])),_:1})])),_:1},8,["value","color","label"])],512),[[Y.F8,d.value.hostname]])])),_:1})}}});var ee=a(11639),te=a(71378),ae=a(8289),le=a(20990);const se=(0,ee.Z)(X,[["__scopeId","data-v-12090e64"]]),re=se;x()(X,"components",{QFooter:te.Z,QIcon:_.Z,QTooltip:Q.Z,QLinearProgress:ae.Z,QBadge:le.Z});const oe=(0,l.aZ)({__name:"MainLayout",setup(e){return(e,t)=>{const a=(0,l.up)("router-view"),s=(0,l.up)("q-page-container"),r=(0,l.up)("q-layout");return(0,l.wg)(),(0,l.j4)(r,{view:"lhh LpR lFf"},{default:(0,l.w5)((()=>[(0,l.Wm)(q),(0,l.Wm)(D),(0,l.Wm)(s,null,{default:(0,l.w5)((()=>[(0,l.Wm)(a)])),_:1}),(0,l.Wm)(re)])),_:1})}}});var ne=a(20249),ue=a(12133);const ce=oe,ie=ce;x()(oe,"components",{QLayout:ne.Z,QPageContainer:ue.Z})}}]);