$(function () {
    window.config = {
        TITLE:'SWord'
    };

    //防止23分钟不通信session失效
    ping();

    //图片禁止托拽和右键
    $('img').on('mousedown contextmenu',function (e) {
        e.preventDefault()
    });

    config_init();
});


function config_init() {
    $('.TITLE').each(function (index,obj) {
        $(obj).html(window.config.TITLE);
    });
}




//防止23分钟不通信session失效
function ping() {
    // setInterval(function () {
    //     $.ajax({
    //         url: "/base/Common/ping",
    //         success: function (res) {
    //             console.log(res)
    //         }
    //     });
    // }, 1000 * 60 * 5);
}


