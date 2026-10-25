package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint16 = 1
    var v2 string = (("hello" + "foo bar") + "\n\t")
    var v3 uint = 42
    var v4 int8 = -128
    v5 := []int{511}
    switch 42 {
    case -1:
        fmt.Println(v3, v2)
    case 143:
        v3 = (v3 | (^(v3)))
    }
    var v6 int32 = 0
    switch "hello" {
    case "\n\t":
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := GenS1{F0: fi7}
            _ = fi7
            _ = v8
        }
    case "":
        var v9 uint = v3
        _ = v9
    case "foo bar":
        var v10 uint8 = 128
        _ = v10
    }
    var v11 uint8 = 255
    v4 = (v4 >> (((v3 + v3)) & 63))
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 float32 = 1e10
        _ = fi12
        _ = v13
    }
    var v14 bool = false
    var v15 float32 = 1.5
    var v16 string = v2
    if false {
        var v17 int64 = 42
        _ = v17
    }
    if v14 {
        v4 -= (-(v4))
    }
    var v18 bool = ((v14 && (-673 != 0)) || v14)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
