$(document).ready(function () {
//1. combonun click fonksiyonu alýnacak!
  $('#get-deldata').click(function () {
    var showData = $('#del-data');

   var e = document.getElementById("billSelect");
   var strUser = e.options[e.selectedIndex].text;
    var res = String(strUser);
   console.log(res)

    $.getJSON('Account2.json', function (data) {
      console.log(data);
    var Months;

      switch (res)
      {
      case "Su":
               Months = data.Account.Bills.Water.map(function (item) {
              return  item.Month;
            });
           break;
      case "Elektrik":
               Months = data.Account.Bills.Electricity.map(function (item) {
                    return  item.Month;
                  });
            break;
       case "Telefon":
               Months = data.Account.Bills.Phone.map(function (item) {
                      return  item.Month;
                              });
            break;
        case "Diðer":
               Months = data.Account.Bills.Other.map(function (item) {
                      return  item.Month;
                              });
            break;
        case "Doðalgaz":
               Months = data.Account.Bills.Gas.map(function (item) {
                      return  item.Month;
                               });
            break;


      }
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


  $('#get-showdata').click(function () {
      var showData2 = $('#show-data');

     var e = document.getElementById("show_bill_select");
     var strUser = e.options[e.selectedIndex].text;
      var res = String(strUser);
     console.log(res)

      $.getJSON('Account2.json', function (data) {
        console.log(data);
      var Months;

        switch (res)
        {
        case "Su":
               Months = data.Account.Bills.Water.map(function (item) {
                return  item.Month;
              });
             break;
        case "Elektrik":
               Months = data.Account.Bills.Electricity.map(function (item) {
                      return  item.Month;
                    });
              break;
         case "Telefon":
                Months = data.Account.Bills.Phone.map(function (item) {
                        return  item.Month;
                                });
              break;
          case "Diðer":
                Months = data.Account.Bills.Other.map(function (item) {
                        return  item.Month;
                                });
              break;
          case "Doðalgaz":
                Months = data.Account.Bills.Gas.map(function (item) {
                        return  item.Month;
                                 });
              break;


        }
        showData2.empty();

        if (Months.length) {
          console.log(Months.length)
          console.log(data)
          var content = '<option>' + Months.join('</option><option>') + '</option>';
          var list = $('<select name="dropdown_ay" id="dropmonth"/>').html(content);
          showData2.append(list);
        }
      });

      showData2.text('Loading the JSON file.');
    });



});


