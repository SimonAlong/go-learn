package test

import (
	"fmt"
	"go-learn/src/util"
	"strings"
	"testing"
)

func TestString(t *testing.T) {

	index := strings.Index("sdf", "=")
	fmt.Println(index)
}

func TestToString(t *testing.T) {
	var value = "mail-template.jsonData=[\n" + "  {\n" + "    \"name\": \"发票模板\",\n" + "    \"key\": \"eInvoiceTemplate\",\n" + "    \"isInside\": \"true\",\n" + "    \"value\": \"<!DOCTYPE html><html lang=\\\\\"en\\\\\" xmlns:th=\\\\\"http://www.thymeleaf.org\\\\\"><head>    <meta charset=\\\\\"utf-8\\\\\">    <meta http-equiv=\\\\\"X-UA-Compatible\\\\\" content=\\\\\"IE=edge\\\\\">    <meta name=\\\\\"viewport\\\\\" content=\\\\\"width=device-width,initial-scale=1.0\\\\\">    <title>电子发票</title></head><body><div style=\\\\\"position: relative\\\\\"><table width=\\\\\"100%\\\\\" cellpadding=\\\\\"0\\\\\" cellspacing=\\\\\"0\\\\\" style=\\\\\"width:100%; border-collapse:collapse; border-spacing:0;font-family: 黑体\\\\\">    <tbody>    <tr>        <td border=\\\\\"border\\\\\" valign=\\\\\"top\\\\\" style=\\\\\"padding:20px; word-break:break-word\\\\\">            <table width=\\\\\"100%\\\\\" border=\\\\\"0\\\\\" cellpadding=\\\\\"0\\\\\" cellspacing=\\\\\"0\\\\\" style=\\\\\"border:0;border-collapse:collapse; border-spacing:0;text-align:left;color:#666;background:#fff;border-bottom-width:1px;border-bottom-color:#ccc;border-bottom-style:solid\\\\\">                <tbody>                <tr>                    <td style=\\\\\"padding-bottom:10px\\\\\">                        <p style=\\\\\"color: #666;font-weight: bold;font-style: normal;font-family: 黑体;text-align:right;font-size: 12px;margin:0\\\\\">联系电话：                            <span style=\\\\\"border-bottom:1px dashed #ccc;z-index:1\\\\\" t=\\\\\"7\\\\\" onclick=\\\\\"return false;\\\\\" data=\\\\\"0571-88157392<\\\\\" th:text=\\\\\"${phoneNumber}\\\\\"></span>                        </p>                        <p style=\\\\\"color: #666;font-weight: bold;font-style: normal;font-style: normal;font-family: 黑体;text-align:right;font-size: 12px;margin:0\\\\\"></p>                    </td>                </tr>                </tbody>            </table>            <table width=\\\\\"100%\\\\\" border=\\\\\"0\\\\\" cellpadding=\\\\\"0\\\\\" cellspacing=\\\\\"0\\\\\" style=\\\\\"border:0;border-collapse:collapse; border-spacing:0;text-align:left;color:#666;background:#fff;border-bottom-width:1px;border-bottom-color:#ccc;border-bottom-style:solid\\\\\">                <tbody>                <tr>                    <td style=\\\\\"text-align:center;padding-top:10px;padding-bottom:10px\\\\\">                        <p style=\\\\\"color: #000;font-weight: bold;font-style: normal;font-family: 黑体;text-align:center;font-size: 24px\\\\\">您申请的电子发票已开出</p>                        <p style=\\\\\"color: #333;font-weight: normal;font-style: normal;font-family: 黑体;text-align:center;font-size: 14px\\\\\">请点击附件查看并核对发票明细</p>                    </td>                </tr>                </tbody>            </table>            <table width=\\\\\"100%\\\\\" border=\\\\\"0\\\\\" cellpadding=\\\\\"0\\\\\" cellspacing=\\\\\"0\\\\\" style=\\\\\"border:0;border-collapse:collapse; border-spacing:0;text-align:left;color:#666;background:#fff;border-bottom-width:1px;border-bottom-color:#ccc;border-bottom-style:solid\\\\\">                <tbody>                <tr>                    <td style=\\\\\"padding-top:10px;padding-bottom:10px\\\\\">                        <div>                            <p style=\\\\\"color: #000;font-weight: normal;font-style: normal;font-family: 黑体;text-align:left;font-size: 16px\\\\\">尊敬的客户，您好:</p>                        </div>                        <div th:each=\\\\\"paragraph : ${paragraphs}\\\\\">                        <p style=\\\\\"color: #000;font-weight: normal;font-style: normal;font-family: 黑体;text-align:left;font-size: 14px\\\\\" th:text=\\\\\"${paragraph}\\\\\"/>                        </div>                    </td>                </tr>                <tr>                    <td style=\\\\\"padding-top:20px;padding-bottom:20px\\\\\">                        <p style=\\\\\"color: #666;font-weight: normal;font-style: normal;font-family: 黑体;text-align:left;font-size: 12px\\\\\">*如果您对发票和行程问题有任何疑问，请联系                            <span style=\\\\\"border-bottom:1px dashed #ccc;z-index:1\\\\\" t=\\\\\"7\\\\\" onclick=\\\\\"return false;\\\\\" data=\\\\\"0571-88157392<\\\\\" th:text=\\\\\"${phoneNumber}\\\\\"/>                        </p>                    </td>                </tr>                </tbody>            </table>        </td>    </tr>    </tbody></table></div></body></html>\"\n" + "  },\n" + "  {\n" + "    \"name\": \"报警模板\",\n" + "    \"key\": \"alarmTemplate\",\n" + "    \"isInside\": \"true\",\n" + "    \"value\": \"<!DOCTYPE html><html lang=\\\\\"en\\\\\" xmlns:th=\\\\\"http://www.thymeleaf.org\\\\\">    <head>        <meta charset=\\\\\"utf-8\\\\\">        <meta http-equiv=\\\\\"X-UA-Compatible\\\\\" content=\\\\\"IE=edge\\\\\">        <meta name=\\\\\"viewport\\\\\" content=\\\\\"width=device-width,initial-scale=1.0\\\\\">        <title>日志告警</title>    </head>    <body>        <table width=\\\\\"100%\\\\\" cellpadding=\\\\\"0\\\\\" cellspacing=\\\\\"0\\\\\"               style=\\\\\"width:100%; border-collapse:collapse; border-spacing:0;font-family: 黑体\\\\\">            <tbody>                <tr>                    <td border=\\\\\"border\\\\\" valign=\\\\\"top\\\\\" style=\\\\\"padding:20px; word-break:break-word\\\\\">                        <table width=\\\\\"100%\\\\\" border=\\\\\"0\\\\\" cellpadding=\\\\\"0\\\\\" cellspacing=\\\\\"0\\\\\"                               style=\\\\\"border:0;border-collapse:collapse; border-spacing:0;text-align:left;color:#666;background:#fff;border-bottom-width:1px;border-bottom-color:#ccc;border-bottom-style:solid\\\\\">                            <tbody>                                <tr>                                    <td style=\\\\\"padding-top:10px;padding-bottom:10px\\\\\">                                        <div>                                            <p style=\\\\\"color: #000;font-weight: normal;font-style: normal;text-align:left;font-size: 16px\\\\\">                                                亲爱的<span th:text=\\\\\"${toName}\\\\\"></span>小伙伴:</p>                                        </div>                                        <div>                                            <p style=\\\\\"color: #000; font-weight: normal; font-style: normal;text-align:left; font-size: 14px\\\\\"                                               th:text=\\\\\"${greetingMsg}\\\\\"></p>                                        </div>                                        <div th:each=\\\\\"entry : ${errMsgUrlMap.entrySet()}\\\\\">                                            <p style=\\\\\"background: cornsilk; border: solid 1px; padding: 10px; color: #000; font-weight: normal; font-style: normal;text-align:left; font-size: 14px\\\\\">                                                <span th:text=\\\\\"${entry.getKey()}\\\\\"></span>                                                <span>(<a th:href=\\\\\"${entry.getValue()[0]}\\\\\">查看详情</a> 或 <a                                                        th:href=\\\\\"${entry.getValue()[1]}\\\\\">30 分钟内不再提醒</a>)</span>                                            </p>                                        </div>                                        <div>                                            <p style=\\\\\"color: #000; font-weight: normal; font-style: normal; text-align:left; font-size: 14px\\\\\">                                                以上链接 1 天内有效，请你及时前往查看错误详情并处理相应问题，谢谢！                                            </p>                                        </div>                                    </td>                                </tr>                                <tr>                                    <td style=\\\\\"padding-top: 20px; padding-bottom: 20px\\\\\">                                        <p style=\\\\\"color: #666; font-weight: normal; font-style: normal; text-align:left; font-size: 12px\\\\\">                                            * 如需恢复此应用已被暂停的通知，请<a th:href=\\\\\"${resumeUrl}\\\\\">点击此处</a>。对此条告警信息若有任何疑问，请联系运维组成员                                        </p>                                    </td>                                </tr>                            </tbody>                        </table>                    </td>                </tr>            </tbody>        </table>    </body></html>\"\n" + "  }\n" + "]"
	fmt.Println(util.KvPropertiesToKeyValueList(value))
}
