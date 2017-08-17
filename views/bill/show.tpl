<div class="pull-right margin_bottom_10px">
  <a href="/bills" class="btn btn-primary">Back</a>
  <a href="/bills/{{.bill.Id}}/print" class="btn btn-default">Print</a>
</div>
<table class="table table-bordered bill_items">
  <tr>
    <td>TIN:29980756946</td>
    <td colspan="3">
      <div class="col-md-12">
        <div class="text-center">TAX INVOICE</div>
      </div>
      <div class="col-md-12">
        <div class="text-center">CASH/CREDIT Bill</div>
      </div>
    </td>
    <td>
      <div class="pull-right">MOB:9343351574</div>
    </td>
  </tr>
  <tr>
    <td colspan="5">
      <div class="col-md-12">
        <div class="text-center">
          <b>SHIVA ENTERPRISES</b>
        </div>  
      </div>
      <div class="col-md-12">
        <div class="text-center"> J.M Road, Kankanady Bajal Mangalore-575027 </div>
      </div>
    </td>
  </tr>
  <tr>
    <th colspan="4">
      No: {{.bill.Billno}}
    </th>
    <th>
      Date: {{.bill.Billdate.Format "02/01/2006 03:04:05 PM" }}
    </th>
  </tr>
  <tr>
    <th colspan="5">
      To: {{.bill.Billto}}
    </th>
  </tr>
  <tr>
    <th>Sl No.</th>
    <th>PARTICULARS</th>
    <th>Qty</th>
    <th>Price Per Unit</th>
    <th>Amount</th>
  </tr>
  {{range $billitem := .billitems}}
    <tr>
      <td>{{$billitem.Slno}}</td>
      <td>{{$billitem.Particulars}}</td>
      <td>{{$billitem.Quantity}}</td>
      <td>{{$billitem.Priceperunit}}</td>
      <td>{{$billitem.Itemqtyprice}}</td>
    </tr>
  {{end}}
  <tr>
    <th colspan="4">
      <div class="col-md-12">
        <div class="pull-right">Sub Total</div>
      </div>
    </th>
    <th>
      {{.bill.Subtotal}}
    </th>
  </tr>
  <tr>
    <th colspan="4">
      <div class="col-md-12">
        <div class="pull-right">CGST</div>
      </div>
    </th>
    <th>
      {{.bill.Cgst}}
    </th>
  </tr>
  <tr>
    <th colspan="4">
      <div class="col-md-12">
        <div class="pull-right">SGST</div>
      </div>
    </th>
    <th>{{.bill.Sgst}}</th>
  </tr>
  <tr>
    <th colspan="4">
      <div class="col-md-12">
        <div class="pull-right">Total</div>
      </div>
    </th>
    <th>{{printf "%.2f" .bill.Grandtotal}}</th>
  </tr>
</table>

