package kalmanfilter

import(
  "testing"
)


func testNew(t *testing.T) {
  new(kalmanfilter.filterData(0,0,0,0,0))
}