package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int; F2 int16 }

func main() {
    var v1 bool = (1e10 > 3.14159)
    var v2 int = 121
    v3 := []int{v2}
    v4 := GenS1{F0: 65}
    fmt.Println(v2, v1)
    if (!(!v1)) {
        var v5 bool = (406 >= -1)
        _ = v5
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 string = "\n\t"
        _ = fi6
        _ = v7
    }
    if ((215 < 0) || ((1 != 1) && v1)) {
        var v8 uint = 2
        _ = v8
    }
    var v9 uint8 = 42
    if (42 <= 400) {
        var v10 uint32 = 2147483648
        _ = v10
    }
    var v11 int32 = 42
    _ = len(v3)
    fmt.Println(v9, v2, v11)
    v9 -= (v9 - 255)
    var v12 int16 = 1000
    switch ("" + "hello") {
    case "\n\t":
        var v13 uint8 = v9
        _ = v13
    default:
        v14 := GenS2{F0: v11, F1: (-(v2)), F2: -1}
        _ = v14
    }
    v15 := GenS2{F0: (v11 + v11), F1: (-(v2)), F2: (v12 & v12)}
    v16 := GenS1{F0: 5}
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint = 65535
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
