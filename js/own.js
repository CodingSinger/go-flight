
var parseJson = function (data,node,showDialog) {

    var lists = data.Include


    var temp;
    node.html("");
    for(var i=0;i<lists.length;i++) {
        temp = node.html();
        flag = "";
        if(lists[i].Remainer == 0){
            flag = "disabled";
        }
        node.html(temp + "<tr class=\"gradeA\">" + "<td>" + lists[i].Fid + "</td>" + "<td>"+lists[i].Departure+"</td>" + "<td class=\"hidden-xs\">"+lists[i].Destination+"</td>" + "<td class=\"center\">"+lists[i].Time+"</td>" + "<td class=\"center hidden-xs\">"+lists[i].Remainer+"</td>" +"<td class=\"center hidden-xs\">"+lists[i].Price+"</td>"+ "<td><button data-id="+lists[i].Fid+" class=\"buy\" "+flag+">购买</button></td>"+ "</tr>");



    }
    init();
    function init() {
        var trs =document.getElementsByClassName("buy");
        for (var i = 0; i < trs.length; i++)
            trs[i].onclick = doclick;
    }//获取tr节点，并循环卫每个tr节点添加双击事件
    function doclick() {


        // $(this).css('background-color','red');
        //弹窗

        var fid = $(this).attr("data-id");

        localStorage.setItem("buyflight",fid);
        showDialog.modal("show");
    }


}

