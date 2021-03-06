<a href="/bills" class="btn btn-primary margin_bottom_10px pull-right">Back</a>
<form class="form-horizontal" action="/bills/create" method="POST">
  <table class="table table-bordered bill_items">
    <tr>
      <th>
        Bill No:
      </th>
      <td>
        <input type="text" class="form-control" name="bill_no" id="bill_no" value={{.bill_no}}>
      </td>
      <th>
        Date:
      </th> 
      <td colspan="2">
        {{.bill_date.Format "02/01/2006 03:04:05 PM"}}
      </td>
    </tr>
    <tr>
      <td colspan="5">
        <b>To:</b><input type="text" class="form-control" name="bill_to" id="bill_to"> 
      </td>
    </tr>
    <tr>
      <th>PARTICULARS</th>
      <th>Qty</th>
      <th>Price Per Unit</th>
      <th>Amount</th>
      <th></th>
    </tr>
    <tr>
      <td class="col-md-5">
        <input type="text" class="form-control" name="bill_items[0][particulars]" id="bill_items_0_particulars">
      </td>
      <td class="col-md-1">
        <input type="text" class="form-control" name="bill_items[0][quantity]" onchange="calculate_item_quantity(0);" id="bill_items_0_quantity">
      </td>
      <td class="col-md-2">
        <input type="text" class="form-control" name="bill_items[0][price_per_unit]" onchange="calculate_item_quantity(0);"  id="bill_items_0_price_per_unit">
      </td>
      <td class="col-md-2">
        <input type="text" class="form-control item_qty_price" name="bill_items[0][item_qty_price]" id="bill_items_0_item_qty_price" onchange="calculate_sub_total();" value="0" readonly>
      </td>
      <td>
        <a href="#" class="btn btn-info" onclick="removeItem(this); return false;">
          <span class="glyphicon glyphicon-remove"></span>
        </a>
        <input type="hidden" name="bill_items[0][destroy]" value="false">
      </td>
    </tr>
  </table>
  <a href="#" class="btn btn-info" onclick="add_item();">Add Item</a>
  <div class="col-md-12">
    <div class="col-md-10">
      <div class="pull-right h4">Sub Total:</div>
    </div> 
    <div class="col-md-2"> 
      <input type="text" class="form-control" id="sub_total" name="sub_total" onchange="calculate_grand_total();" value="0.0" readonly>
    </div>  
  </div>
  <div class="col-md-12">
    <div class="col-md-10">
      <div class="pull-right h4">CGST:</div>
    </div>  
    <div class="col-md-2">
      <input type="text" class="form-control" id="cgst" name="cgst" value="0.0" readonly>
    </div>  
  </div>
  <div class="col-md-12">
    <div class="col-md-10">
      <div class="pull-right h4">SGST:</div>
    </div> 
    <div class="col-md-2"> 
      <input type="text" class="form-control" id="sgst" name="sgst" value="0.0" readonly>
    </div>  
  </div>
  <div class="col-md-12">
    <div class="col-md-5">
      <div class="col-md-6">
		    <div class="pull-right h4">Vehicle No:</div>
      </div>
      <div class="col-md-6">
		    <input type="text" class="form-control" id="vehicle_no" name="vehicle_no">
      </div>   
		</div>
    <div class="col-md-7 padding_0px">
      <div class="col-md-6">
        <div class="pull-right h4">Total:</div>
      </div>
      <div class="col-md-6">
        <input type="text" class="form-control" id="grand_total" name="grand_total" value="0.0" readonly>
      </div>  
    </div>
  </div>
  <div class="clearfix"></div>
  <div class="pull-left">
    <button type="submit" class="btn btn-success">Save</button>
  </div>  
</form>  
<script type="text/javascript">
  $(document).ready(function(){
        //$('.datetimepicker').datetimepicker();
  })
</script>
