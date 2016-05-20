
$(document).ready(function () {
//1. combonun click fonksiyonu alýnacak!
  $('#get-data').click(function () {
    var showData = $('#show-data');

    $.getJSON('Account2.json', function (data) {
      console.log(data);

      var Months = data.Account.Bills.Water.map(function (item) {
        return  item.Month;
      });

      showData.empty();

      if (Months.length) {
        console.log(Months.length)
        console.log(data)
        var content = '<option>' + Months.join('</option><option>') + '</option>';
        var list = $('<select name="dropdown_ay" id="dropmonth"/>').html(content);
        showData.append(list);
      }
    });

    showData.text('Loading the JSON file.');
  });
});



/*$(document).ready(function () {
  $('#get-data').click(function () {
    var showData = $('#show-data');

    $.getJSON('example.json', function (data) {
      console.log(data);

      var items2 = data.items.map(function (item) {
        return item.key + ': ' + item.value;
      });

      showData.empty();

      if (items2.length) {
        var content = '<li>' + items2.join('</li><li>') + '</li>';
        var list = $('<ul />').html(content);
        showData.append(list);
      }
    });

    showData.text('Loading the JSON file.');
  });
});*/