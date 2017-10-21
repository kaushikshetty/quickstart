package helpers
import (
  "net/http"
  //"strconv"
  "regexp"
  models "quickstart/models"
  "fmt"
)
func ParseFormCollection(r *http.Request)  []map[string]string {
  // reading nested attributes
  //form := c.Ctx.Input.Context.Request.Form
  form := r.Form
  var index []string
  for key, _ := range form {
    re := regexp.MustCompile("bill_items" + "\\[([0-9]+)\\]\\[([a-zA-Z_]+)\\]")
    matches := re.FindStringSubmatch(key) // this will return [bill_items[0][sl_no], 0, sl_no]
    if len(matches) >= 3 {
       i := matches[1]
       if (!contains(index,i)){
         index = append(index, i)
         fmt.Println(index)
       }
    }
  }
  var bill_items []map[string]string
  for i :=0 ; i < len(index); i++{
    bill_item := make(map[string]string)
    bill_item["particulars"] = r.PostFormValue("bill_items["+ index[i] +"][particulars]")
    bill_item["item_qty_price"] = r.PostFormValue("bill_items["+ index[i] +"][item_qty_price]")
    bill_item["quantity"] = r.PostFormValue("bill_items["+ index[i] +"][quantity]")
    bill_item["price_per_unit"] = r.PostFormValue("bill_items["+ index[i] +"][price_per_unit]")
    bill_item["destroy"] = r.PostFormValue("bill_items["+ index[i] +"][destroy]")
    bill_item["id"] = index[i]
    bill_items = append(bill_items, bill_item)
  }
  //fmt.Println(bill_items)
  return bill_items
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func BillPrintHtmlCode(bill models.Bill, billitems []*models.Billitem) string {
  str := ""
  return str
}
