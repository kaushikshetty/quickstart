<table class="table table-bordered bill_items">
  <tr>
    <th>Billno</th>
    <th>Billto</th>
    <th>Billdate</th>
    <th>Subtotal</th>
    <th>Cgst</th>
    <th>Sgst</th>
    <th>Grandtotal</th>
    <th>VehicleNo</th>
    <th colspan="2"></th>
  </tr>
	{{range $bill := .bills}}
		  <tr>
		    <td>{{$bill.Billno}}</td>
		    <td>{{$bill.Billto}}</td>
		    <td>{{$bill.Billdate.Format "02/01/2006 03:04:05 PM" }}</td>
		    <td>{{$bill.Subtotal}}</td>
		    <td>{{$bill.Cgst}}</td>
		    <td>{{$bill.Sgst}}</td>
		    <td>{{printf "%.2f" $bill.Grandtotal}}</td>
                    <td>{{$bill.Vehicleno}}</td>
        <td><a href="/bills/{{$bill.Id}}/edit" class="btn btn-warning">Edit Bill</a></td>
        <td><a href="/bills/{{$bill.Id}}" class="btn btn-info">Show Bill</a></td>
		  </tr>
	{{end}}
</table>
