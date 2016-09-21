var urlStr = location.href;

$("#side-nav a").each(function () {
  if(urlStr.lastIndexOf($(this).attr('href')) > -1 && $(this).attr('href') != '') {
    $(this).parent().addClass('active');
  }else {
    $(this).parent().removeClass('active');
  }
})
