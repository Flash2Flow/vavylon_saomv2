document.addEventListener('keyup', function(){
if (event.keyCode == 36){
$( "#container" ).stop().toggle();
$('#btn_one').on('click',() => {
  $('#panelo').stop().toggle();
  $('#panelt').stop().toggle();
  $('#panelth').stop().toggle();

  $('#indicator').hide();
  $('#slide').hide();
  $('#indicatorth').hide();
  $('#slideth').hide();
  $('#slideo').hide();
  $('#indicatort').hide();
  $('#container_message').hide();
  $('#accept').hide();
  $('#close').hide();
  return false;
});


$('#link_reg').on('click',() => {
  $('#block_one').show();
  $('#block_two').hide();
});

$('#link_auth').on('click',() => {
  $('#block_two').show();
    $('#block_one').hide();
});

$('#btn_two').on('click',() => {
  $('#indicator').stop().toggle();
  $('#slide').stop().toggle();

  $('#panelo').hide();
  $('#panelt').hide();
  $('#panelth').hide();
  $('#container_message').hide();
  $('#accept').hide();
  $('#close').hide();
  $('#indicatorth').hide();
  $('#slideth').hide();
  $('#slideo').hide();
  $('#indicatort').hide();
  return false;
});
$('#btn_three').on('click',() => {
  $('#container_message').stop().toggle();
  $('#accept').stop().toggle();
  $('#close').stop().toggle();

  $('#panelo').hide();
  $('#panelt').hide();
  $('#panelth').hide();
  $('#indicator').hide();
  $('#slide').hide();
  $('#slideo').hide();
  $('#indicatort').hide();
  $('#indicatorth').hide();
  $('#slideth').hide();
  return false;
});
$('#btn_four').on('click',() => {
  $('#slideo').stop().toggle();
  $('#indicatort').stop().toggle();

  $('#panelo').hide();
  $('#panelt').hide();
  $('#panelth').hide();
  $('#indicator').hide();
  $('#slide').hide();
  $('#container_message').hide();
  $('#accept').hide();
  $('#close').hide();
  $('#indicatorth').hide();
  $('#slideth').hide();
  return false;
});
$('#btn_five').on('click',() => {
  $('#indicatorth').stop().toggle();
  $('#slideth').stop().toggle();

  $('#panelo').hide();
  $('#panelt').hide();
  $('#panelth').hide();
  $('#indicator').hide();
  $('#slide').hide();
  $('#container_message').hide();
  $('#accept').hide();
  $('#close').hide();
  $('#slideo').hide();
  $('#indicatort').hide();
  return false;
});


var element = document.querySelector('.container');
window.addEventListener('mousemove', function(e) {
      var x = e.clientX / window.innerWidth;
      var y = e.clientY / window.innerHeight;
      element.style.transform = 'translate(-' + x * 100 + 'px, -' + y * 100 + 'px)';
});
}
});

$(document).ready(function(){ //дожидаемся загрузки страницы
 $('#link_reg').on("click", function(){ //вешаем событие на клик по кнопке id="btn1"
 $('#block_one').show();//включает/выключает элемент id="text"
 $('#block_two').hide();
});

$('#link_auth').on("click", function(){ //вешаем событие на клик по кнопке id="btn1"
$('#block_one').hide();//включает/выключает элемент id="text"
$('#block_two').show();
});
/*play*/
$('#launch').on("click", function(){ //вешаем событие на клик по кнопке id="btn1"
$('#start_one').show();//включает/выключает элемент id="text"
});

$('#yes_one').on("click", function(){ //вешаем событие на клик по кнопке id="btn1"
$('#start_two').show();//включает/выключает элемент id="text"
$('#start_one').hide();
});

$('#no_one').on("click", function(){ //вешаем событие на клик по кнопке id="btn1"
$('#start_two').show();//включает/выключает элемент id="text"
$('#start_one').hide();//включает/выключает элемент id="text"
});

$('#yes_two').on("click", function(){ //вешаем событие на клик по кнопке id="btn1"
$('#start_three').show();//включает/выключает элемент id="text"
$('#start_two').hide();
});

$('#no_two').on("click", function(){ //вешаем событие на клик по кнопке id="btn1"
$('#start_two').hide();//включает/выключает элемент id="text"
$('#start_three').show();//включает/выключает элемент id="text"
});

$('#close_one').on("click", function(){ //вешаем событие на клик по кнопке id="btn1"
$('#start_three').hide();//включает/выключает элемент id="text"
});

 });
