<a href="/bills" class="btn btn-primary pull-right margin_bottom_10px">Back</a>
<form class="form-horizontal" action="/bills/{{.bill.Id}}/update" method="POST">
  <table class="table table-bordered bill_items">
    <tr>
      <th>
        Bill No:
      </th>
      <td>
        <input type="text" class="form-control" name="bill_no" id="bill_no" value={{.bill.Billno}}>
      </td>
      <th>
        Date:
      </th> 
      <td colspan="2">
        {{.bill.Billdate.Format "02/01/2006 03:04:05 PM"}}
      </td>
    </tr>
    <tr>
      <td colspan="5">
        <b>To:</b><input type="text" class="form-control" name="bill_to" id="bill_to" value={{.bill.Billto}}> 
      </td>
    </tr>
    <tr>
      <th>PARTICULARS</th>
      <th>Qty</th>
      <th>Price Per Unit</th>
      <th>Amount</th>
      <th></th>
    </tr>
    {{range $billitem := .billitems}}
	    <tr id="{{$billitem.Id}}">
	      <td class="col-md-5">
		      <input type="text" class="form-control" name="bill_items[{{$billitem.Id}}][particulars]" value={{$billitem.Particulars}} id="bill_items_{{$billitem.Id}}_particulars">
	      </td>
	      <td class="col-md-1">
		      <input type="text" class="form-control" name="bill_items[{{$billitem.Id}}][quantity]" onchange="calculate_item_quantity({{$billitem.Id}});" id="bill_items_{{$billitem.Id}}_quantity" value={{$billitem.Quantity}}>
	      </td>
	      <td class="col-md-2">
		      <input type="text" class="form-control" name="bill_items[{{$billitem.Id}}][price_per_unit]" onchange="calculate_item_quantity({{$billitem.Id}});" id="bill_items_{{$billitem.Id}}_price_per_unit" value={{$billitem.Priceperunit}}>
	      </td>
	      <td class="col-md-2">
		      <input type="text" class="form-control item_qty_price" name="bill_items[{{$billitem.Id}}][item_qty_price]" id="bill_items_{{$billitem.Id}}_item_qty_price" onchange="calculate_sub_total();" value={{$billitem.Itemqtyprice}} readonly>
	      </td>
        <td>
          <a href="#" class="btn btn-info" onclick="removeItem(this); return false;">
            <span class="glyphicon glyphicon-remove"></span>
          </a>
          <input type="hidden" name="bill_items[{{$billitem.Id}}][destroy]" value="false">
        </td>
	    </tr>
   {{end}}
  </table>
  <a href="#" class="btn btn-info" onclick="add_item();">Add Item</a>
  <div class="col-md-12">
    <div class="col-md-10">
      <div class="pull-right h4">Sub Total:</div>
    </div> 
    <div class="col-md-2"> 
      <input type="text" class="form-control" id="sub_total" name="sub_total" onchange="calculate_grand_total();" value={{.bill.Subtotal}} readonly>
    </div>  
  </div>
  <div class="col-md-12">
    <div class="col-md-10">
      <div class="pull-right h4">CGST:</div>
    </div>  
    <div class="col-md-2">
      <input type="text" class="form-control" id="cgst" name="cgst" value={{.bill.Cgst}} readonly>
    </div>  
  </div>
  <div class="col-md-12">
    <div class="col-md-10">
      <div class="pull-right h4">SGST:</div>
    </div> 
    <div class="col-md-2"> 
      <input type="text" class="form-control" id="sgst" name="sgst" value={{.bill.Sgst}} readonly>
    </div>  
  </div>
  <div class="col-md-12">
    <div class="col-md-5">
      <div class="col-md-6">
		    <div class="pull-right h4">Vehicle No:</div>
      </div>
      <div class="col-md-6">
		    <input type="text" class="form-control" id="vehicle_no" name="vehicle_no" value={{.bill.Vehicleno}}>
      </div>   
		</div>
    <div class="col-md-7 padding_0px">
      <div class="col-md-6">
        <div class="pull-right h4">Total:</div>
      </div>
      <div class="col-md-6">
        <input type="text" class="form-control" id="grand_total" name="grand_total" value={{.bill.Grandtotal}} readonly>
      </div>  
    </div>
  </div>
  <div class="clearfix"></div>
  <div class="pull-left">
    <button type="submit" class="btn btn-success margin_bottom_10px">Save</button>
  </div>  
</form>
