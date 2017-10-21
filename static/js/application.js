function calculate_item_quantity(id){
  quantity = $('#bill_items_'+id+'_quantity').val()
  price_per_unit = $('#bill_items_'+id+'_price_per_unit').val()
  item_qty_price = quantity * price_per_unit
  $('#bill_items_'+id+'_item_qty_price').val(parseFloat((item_qty_price).toFixed(2))).trigger("change");
}

function calculate_sub_total(){
  sub_total = 0.0
  $('.item_qty_price').each(function(){
    if($(this).parents('tr').find('input[type="hidden"]').val() == "false"){
      item_qty_price = $(this).val()
      sub_total = sub_total + parseFloat(item_qty_price)
    }
  })
  $('#sub_total').val(parseFloat((sub_total).toFixed(2))).trigger("change");
}

function calculate_grand_total(){
  sub_total = parseFloat($('#sub_total').val())
  cgst = ((sub_total * 14.5)/100)
  sgst = ((sub_total * 14.5)/100)
  grand_total = (sub_total + cgst + sgst)
  grand_total = parseFloat((grand_total).toFixed(2))
  $('#cgst').val(cgst)
  $('#sgst').val(sgst)
  $('#grand_total').val(grand_total)
}

function add_item(){
  var new_id = new Date().getTime();
  string = "<tr>"
    string += "<td class='col-md-5'>"
      string += "<input type='text' class='form-control' name='bill_items["+new_id+"][particulars]' id='bill_items_"+new_id+"_particulars'>"
    string += "</td>"
    string +="<td class='col-md-1'>"
      string +="<input type='text' class='form-control' name='bill_items["+new_id+"][quantity]' onchange='calculate_item_quantity("+new_id+");' id='bill_items_"+new_id+"_quantity'>"
    string += "</td>"
    string += "<td class='col-md-2'>"
      string += "<input type='text' class='form-control' name='bill_items["+new_id+"][price_per_unit]'' onchange='calculate_item_quantity("+new_id+");' id='bill_items_"+new_id+"_price_per_unit'>"
    string +="</td>"
    string += "<td class='col-md-2'>"
      string += "<input type='text' class='form-control item_qty_price' name='bill_items["+new_id+"][item_qty_price]' id='bill_items_"+new_id+"_item_qty_price' onchange='calculate_sub_total();' value='0' readonly>"
    string += "</td>"
    string +="<td><a href='#' class='btn btn-info' onclick='removeItem(this); return false;'><span class='glyphicon glyphicon-remove'></span></a><input type='hidden' name='bill_items["+new_id+"][destroy]' value='false'></td>"
  string +="</tr>"
  $('.bill_items tr:last').after(string);
}

function removeItem(current_link){
  $(current_link).parents('tr').hide();
  $(current_link).parent().find('input[type="hidden"]').val("true") 
  calculate_sub_total();
}
