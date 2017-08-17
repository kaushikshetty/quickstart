package helpers
import (
  "net/http"
  "strconv"
  "regexp"
  models "quickstart/models"
  //"fmt"
)
func ParseFormCollection(r *http.Request)  []map[string]string {
  // reading nested attributes
  //form := c.Ctx.Input.Context.Request.Form
  form := r.Form
  var bill_items []map[string]string
  for key, values := range form {
    re := regexp.MustCompile("bill_items" + "\\[([0-9]+)\\]\\[([a-zA-Z_]+)\\]")
    matches := re.FindStringSubmatch(key) // this will return [bill_items[0][sl_no], 0, sl_no]
    if len(matches) >= 3 {
		  index, _ := strconv.Atoi(matches[1])
		  for ; index >= len(bill_items); {
			  bill_items = append(bill_items, map[string]string{})
		  }
      bill_items[index][matches[2]] = values[0] // this will add [bill_items[0][sl_no]=2 
    }
  
  }
  return bill_items
}

func BillPrintHtmlCode(bill models.Bill, billitems []*models.Billitem) string {
  str := ""
  return str
}
