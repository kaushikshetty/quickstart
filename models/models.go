package models
import (
  "time"
  "github.com/astaxie/beego/orm"
)

type Bill struct {
    Id int `orm:"auto"`
    Billno  string 
    Billto string 
    Billdate time.Time `orm:"auto_now_add;type(datetime)"`
    Subtotal float64 
    Cgst float64 
    Sgst float64 
    Grandtotal float64 
    Created time.Time `orm:"auto_now_add;type(datetime)"`
    Updated time.Time `orm:"auto_now_add;type(datetime)"`
    Billitems []*Billitem `orm:"reverse(many)"`
    Vehicleno string `orm:"null"`
}

type Billitem struct {
    Id int `orm:"auto"`
    Slno   string    
    Particulars string 
    Quantity int 
    Priceperunit float64 
    Itemqtyprice float64
    Created time.Time `orm:"auto_now_add;type(datetime)"`
    Updated time.Time `orm:"auto_now_add;type(datetime)"`
    Bill *Bill `orm:"rel(fk)"`
}

func init() {
    orm.RegisterModel(new(Bill), new(Billitem))
}
