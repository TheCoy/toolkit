package oop

import "testing"

func TestStruct(t *testing.T) {
    var newcar = Car{name: "macan", wheelCount: 23}
    newcar.start()
    newcar.stop()

    var mycar Car
    mycar.start()
    mycar.stop()

    var merkel = Mercedes{Car: newcar}
    merkel.start()
    merkel.stop()
    merkel.sayHiToMerkel()


}
