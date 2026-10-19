package main

import (
    "fmt"
)

func genFn0() int {
    var v4 uint = 541
    var v5 int16 = -1
    var v6 uint8 = 127
    var v7 uint16 = 0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 int16 = -216
    var v2 int16 = v1
    var v3 bool = false
    v8 := []int{genFn0()}
    var v9 string = ""
    fmt.Println(v2)
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int16 = v1
        v9 = "hello"
        _ = fi10
        _ = v11
    }
    v3 = (!false)
    if len(v8) > 0 {
        v8[0] = 4294967295
    }
    fmt.Println(v1, v3)
    var v12 uint8 = 0
    switch -312 {
    case -855:
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    default:
        if (-1 >= 1) {
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 uint32 = 897
                _ = fi13
                _ = v14
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
