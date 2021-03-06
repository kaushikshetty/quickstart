package controllers

import (
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  models "quickstart/models"
  helpers "quickstart/helpers"
  "fmt"
  "strconv"
  "bytes"
  "bufio"
  "github.com/jung-kurt/gofpdf"
  "time"
)

type BillController struct {
  beego.Controller
}

func (c *BillController) Print() {
  i, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
  bill := models.Bill{Id: i}
  var billitems []*models.Billitem
  o := orm.NewOrm()
  err := o.Read(&bill)
  if err == nil{
    o.QueryTable("billitem").Filter("Bill", bill).All(&billitems)
  }
  fmt.Println(billitems)
  pdf := gofpdf.New("P", "mm", "A4", "")
  pdf.AddPage()
  pdf.SetFont("Arial", "B", 12)
  pdf.CellFormat(70, 0, "GSTIN:29AJAPD11147N1Z8","0", 0, "", false, 0, "")
  pdf.CellFormat(40,0, "TAX INVOICE","0", 0, "C", false, 0, "")
  pdf.CellFormat(40, 0, "MOB:9343351674,9901374956", "0", 0, "", false, 0, "")
  pdf.Ln(5)
  pdf.CellFormat(180, 0, "CASH/CREDIT Bill", "0", 0, "C", false, 0, "")
  pdf.Ln(7)
  pdf.CellFormat(180, 0, "SHIVA ENTERPRISES", "0", 0, "C", false, 0, "")
  pdf.Ln(5)
  pdf.CellFormat(170, 0, "J.M Road, Kankanady Bajal Mangalore-575027", "0", 0, "C", false, 0, "")
  pdf.Ln(7)
  pdf.CellFormat(0, 0, "STATE CODE: 29","0", 0, "", false, 0, "")
  pdf.Ln(5)
  pdf.CellFormat(50, 0, "A/c No:1169261000640","0", 0, "", false, 0, "")
  pdf.CellFormat(80,0, "Canara Bank Carstreet Mangalore","0", 0, "C", false, 0, "")
  pdf.CellFormat(40, 0, "IFSC Code:CNRB0001169", "0", 0, "", false, 0, "")
  pdf.Ln(5)
  pdf.CellFormat(7, 0, "No","0", 0, "", false, 0, "")
  pdf.CellFormat(120, 0, bill.Billno, "0", 0, "", false, 0, "")
  pdf.CellFormat(10, 0, "Date:", "0", 0, "", false, 0, "")
  pdf.CellFormat(30, 0, bill.Billdate.Format("02/01/2006 03:04:05 PM"), "0", 0, "", false, 0, "")
  pdf.Ln(5)
  basicTable := func() {
    pdf.CellFormat(10, 6, "To:", "", 0, "", false, 0, "")
    pdf.CellFormat(70, 6, bill.Billto, "B", 0, "", false, 0, "")
    pdf.CellFormat(10, 6, "", "B", 0, "", false, 0, "")
    pdf.CellFormat(30, 6, "", "B", 0, "", false, 0, "")
    pdf.CellFormat(30, 6, "", "B", 0, "", false, 0, "")
    pdf.Ln(-1)
    pdf.CellFormat(10, 6, "Party's GSTIN:", "", 0, "", false, 0, "")
    pdf.CellFormat(70, 6, "", "B", 0, "", false, 0, "")
    pdf.CellFormat(10, 6, "", "B", 0, "", false, 0, "")
    pdf.CellFormat(30, 6, "", "B", 0, "", false, 0, "")
    pdf.CellFormat(30, 6, "", "B", 0, "", false, 0, "")
    pdf.Ln(-1)
    pdf.Ln(-1)
    pdf.CellFormat(10, 6, "Slno", "1", 0, "", false, 0, "")
    pdf.CellFormat(90, 6, "Particulars", "1", 0, "", false, 0, "")
    pdf.CellFormat(10, 6, "Qty", "1", 0, "", false, 0, "")
    pdf.CellFormat(30, 6, "Price Per Item", "1", 0, "", false, 0, "")
    pdf.CellFormat(30, 6, "Amount", "1", 0, "", false, 0, "")
		pdf.Ln(-1)
    sl_no := 0
		for i, c := range billitems {
      sl_no = i +1
			pdf.CellFormat(10, 6, strconv.Itoa(sl_no), "1", 0, "", false, 0, "")
			pdf.CellFormat(90, 6, c.Particulars, "1", 0, "", false, 0, "")
			pdf.CellFormat(10, 6, strconv.Itoa(c.Quantity), "1", 0, "", false, 0, "")
			pdf.CellFormat(30, 6, strconv.FormatFloat(c.Priceperunit, 'f', -1, 32), "1", 0, "", false, 0, "")
      pdf.CellFormat(30, 6, strconv.FormatFloat(c.Itemqtyprice,'f', -1, 32), "1", 0, "", false, 0, "")
			pdf.Ln(-1)
		}
	}
  basicTable()
  for i := 0; i < 25 - len(billitems); i++ {
    pdf.CellFormat(10, 6, "", "LR", 0, "", false, 0, "")
    pdf.CellFormat(90, 6, "", "LR", 0, "", false, 0, "")
    pdf.CellFormat(10, 6, "", "LR", 0, "", false, 0, "")
    pdf.CellFormat(30, 6, "", "LR", 0, "", false, 0, "")
    pdf.CellFormat(30, 6, "", "LR", 0, "", false, 0, "")
    pdf.Ln(-1)
  }
  pdf.CellFormat(10, 6, "", "LT", 0, "", false, 0, "")
  pdf.CellFormat(90, 6, "", "T", 0, "", false, 0, "")
  pdf.CellFormat(10, 6, "", "T", 0, "", false, 0, "")
  pdf.CellFormat(30, 6, "Sub Total", "LRT", 0, "", false, 0, "")
  pdf.CellFormat(30, 6, strconv.FormatFloat(bill.Subtotal,'f', -1, 32), "LRT", 0, "", false, 0, "")
  pdf.Ln(-1)
  pdf.CellFormat(10, 6, "", "L", 0, "", false, 0, "")
  pdf.CellFormat(90, 6, "", "", 0, "", false, 0, "")
  pdf.CellFormat(10, 6, "", "", 0, "", false, 0, "")
  pdf.CellFormat(30, 6, "CGST", "LR", 0, "", false, 0, "")
  pdf.CellFormat(30, 6, strconv.FormatFloat(bill.Cgst,'f', -1, 32), "LR", 0, "", false, 0, "")
  pdf.Ln(-1)
  pdf.CellFormat(10, 6, "", "L", 0, "", false, 0, "")
  pdf.CellFormat(90, 6, "", "", 0, "", false, 0, "")
  pdf.CellFormat(10, 6, "", "", 0, "", false, 0, "")
  pdf.CellFormat(30, 6, "SGST", "LRB", 0, "", false, 0, "")
  pdf.CellFormat(30, 6, strconv.FormatFloat(bill.Sgst,'f', -1, 32), "LRB", 0, "", false, 0, "")
  pdf.Ln(-1)
  pdf.CellFormat(10, 6, "", "LB", 0, "", false, 0, "")
  pdf.CellFormat(90, 6, "", "B", 0, "", false, 0, "")
  pdf.CellFormat(10, 6, "", "RB", 0, "", false, 0, "")
  pdf.CellFormat(30, 6, "Total", "RB", 0, "", false, 0, "")
  pdf.CellFormat(30, 6, strconv.FormatFloat(bill.Grandtotal,'f', -1, 32), "RB", 0, "", false, 0, "")
  pdf.Ln(14)
  pdf.CellFormat(100, 0, "Goods once sold will not be taken back", "0", 0, "L", false, 0, "")
  pdf.CellFormat(50, 0, "For SHIVA ENTERPRISES", "0", 0, "L", false, 0, "")
  pdf.Ln(16)
  pdf.CellFormat(25, 0, "Vechile No:","0", 0, "", false, 0, "")
  pdf.CellFormat(80, 0, bill.Vehicleno, "0", 0, "", false, 0, "")
  pdf.CellFormat(10, 0, "Signature", "0", 0, "", false, 0, "")
  var b bytes.Buffer
  w := bufio.NewWriter(&b)
  pdf.Output(w)
  pdf.Close()
  w.Flush()
  c.Ctx.Output.ContentType("application/pdf")
  c.Ctx.Output.Body(b.Bytes())
}

func (c *BillController) New() {
  o := orm.NewOrm()
  var bill models.Bill
  err := o.QueryTable("bill").OrderBy("-Id").One(&bill,"Id", "Billno")
  if err == orm.ErrNoRows {
     fmt.Println("No last bill")
  }else{
     fmt.Println(bill.Billno)
  }
  i, err := strconv.Atoi(bill.Billno)
  c.Data["bill_date"] = time.Now()
  c.Data["bill_no"] = i + 1
  c.Layout = "dashboard.html"
  c.TplName = "bill/new.tpl"
}

func (c *BillController) Create() {
  bill_items := helpers.ParseFormCollection(c.Ctx.Input.Context.Request)
  type Billtype struct {
    Billno  string `form:"bill_no"`
    Billto string  `form:"bill_to"`
    Subtotal float64 `form:"sub_total"`
    Cgst float64 `form:"cgst"`
    Sgst float64 `form:"sgst"`
    Grandtotal float64 `form:"grand_total"`
    Vehicleno string `form:"vehicle_no"`
  }
  billtype := Billtype{}
  if err := c.ParseForm(&billtype); err != nil {
    fmt.Printf("%v\n", err)
  }
  o := orm.NewOrm()
  o.Using("default")
  bill := new(models.Bill)
  bill.Billno = billtype.Billno
  bill.Billto = billtype.Billto
  bill.Subtotal = billtype.Subtotal
  bill.Grandtotal = billtype.Grandtotal
  bill.Cgst = billtype.Cgst
  bill.Sgst = billtype.Sgst
  bill.Vehicleno = billtype.Vehicleno
  id, err := o.Insert(bill)
  fmt.Printf("ID: %d, ERR: %v\n", id, err)
  if err == nil {
    for _, item := range bill_items {
      if (item["destroy"] == "false"){
		    billitem := new(models.Billitem)
		    //billitem.Slno = item["sl_no"]
		    billitem.Particulars = item["particulars"]
		    item_qty_price, _ := strconv.ParseFloat(item["item_qty_price"], 64)
		    billitem.Itemqtyprice = item_qty_price
		    qty, _ := strconv.Atoi(item["quantity"])
		    billitem.Quantity = qty
		    price_per_unit, _ := strconv.ParseFloat(item["price_per_unit"], 64)
		    billitem.Priceperunit = price_per_unit
		    billitem.Bill = bill
		    o.Insert(billitem)
     }
   }
 }
  url := "/bills/"+strconv.Itoa(int(id))
  c.Ctx.Redirect(302, url)
}

func (c *BillController) Show() {
  i, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
  bill := models.Bill{Id: i}
  var billitems []*models.Billitem
  o := orm.NewOrm()
  err := o.Read(&bill)
  if err == nil{
    o.QueryTable("billitem").Filter("Bill", bill).All(&billitems)
 }
  c.Data["bill"] = bill
  c.Data["billitems"] = billitems
  c.Layout = "dashboard.html"
  c.TplName = "bill/show.tpl"
}

func (c *BillController) Index() {
  var bills []*models.Bill
  o := orm.NewOrm()
  o.QueryTable("bill").OrderBy("-id").All(&bills, "Id", "Billno", "Billto", "Billdate", "Subtotal", "Cgst", "Sgst", "Grandtotal","Vehicleno")
  c.Data["bills"] = bills
  c.Layout = "dashboard.html"
  c.TplName = "bill/index.tpl"
}

func (c *BillController) Edit() {
  i, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
  bill := models.Bill{Id: i}
  var billitems []*models.Billitem
  o := orm.NewOrm()
  err := o.Read(&bill)
  if err == nil{
    o.QueryTable("billitem").Filter("Bill", bill).All(&billitems)
  }
  c.Data["bill"] = bill
  c.Data["billitems"] = billitems
  c.Layout = "dashboard.html"
  c.TplName = "bill/edit.tpl" 
}

func (c *BillController) Update() {
  bill_items := helpers.ParseFormCollection(c.Ctx.Input.Context.Request)
  //fmt.Println(bill_items)
  billid, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
  //bill := models.Bill{Id: billid}
  bill := new(models.Bill)
  bill.Id = billid
  type Billtype struct {
    Billno  string `form:"bill_no"`
    Billto string  `form:"bill_to"`
    Subtotal float64 `form:"sub_total"`
    Cgst float64 `form:"cgst"`
    Sgst float64 `form:"sgst"`
    Grandtotal float64 `form:"grand_total"`
    Vehicleno string `form:"vehicle_no"`
  }
  billtype := Billtype{}
  if err := c.ParseForm(&billtype); err != nil {
    fmt.Printf("%v\n", err)
  }
  o := orm.NewOrm()
  o.Using("default")
  o.Read(bill)
  bill.Billno = billtype.Billno
  bill.Billto = billtype.Billto
  bill.Subtotal = billtype.Subtotal
  bill.Grandtotal = billtype.Grandtotal
  bill.Cgst = billtype.Cgst
  bill.Sgst = billtype.Sgst
  bill.Vehicleno = billtype.Vehicleno
  bill.Updated = time.Now()
  _, err := o.Update(bill, "Subtotal", "Grandtotal","Cgst","Sgst","Vehicleno","Billno", "Billto","Updated")
  if err == nil {
    for _, item := range bill_items {
      fmt.Println(item)
      billitemid, _ := strconv.Atoi(item["id"])
      billitem := models.Billitem{Id: billitemid}
      err := o.Read(&billitem)
      if err == orm.ErrNoRows {
        fmt.Println("new")
        fmt.Println(billitemid)
        billitem := new(models.Billitem)
        billitem.Slno = item["sl_no"]
        billitem.Particulars = item["particulars"]
        item_qty_price, _ := strconv.ParseFloat(item["item_qty_price"], 64)
        billitem.Itemqtyprice = item_qty_price
        qty, _ := strconv.Atoi(item["quantity"])
        billitem.Quantity = qty
        price_per_unit, _ := strconv.ParseFloat(item["price_per_unit"], 64)
        billitem.Priceperunit = price_per_unit
        billitem.Bill = bill
        o.Insert(billitem)
      }else{
        fmt.Println("edit")
        fmt.Println(billitemid)
        if (item["destroy"] == "true"){
          fmt.Println("delete")
          if num, err := o.Delete(&billitem); err == nil {
            fmt.Println(num)
          }
        }else{
		      billitem.Slno = item["sl_no"]
				  billitem.Particulars = item["particulars"]
				  item_qty_price, _ := strconv.ParseFloat(item["item_qty_price"], 64)
				  billitem.Itemqtyprice = item_qty_price
				  qty, _ := strconv.Atoi(item["quantity"])
				  billitem.Quantity = qty
				  price_per_unit, _ := strconv.ParseFloat(item["price_per_unit"], 64)
				  billitem.Priceperunit = price_per_unit
		      billitem.Updated = time.Now()
			    o.Update(&billitem, "Slno", "Particulars","Itemqtyprice","Quantity","Priceperunit","Updated")
       }
      }
    }
  }
   url := "/bills/"+strconv.Itoa(int(billid))
   c.Ctx.Redirect(302, url)
}

