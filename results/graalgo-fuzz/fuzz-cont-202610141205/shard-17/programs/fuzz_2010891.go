package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint32 }
func genFn0() int16 {
    var v5 string = "foo bar"
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{(fi6 | (fi6 - fi6)), fi6, fi6, fi6}
        _ = fi6
        _ = v7
    }
    var v8 int8 = 122
    v8 -= (v8 & v8)
    _ = v5
    _ = v8
    return -1
}

func main() {
    var v1 int16 = 1000
    var v2 uint8 = 42
    var v3 int32 = -394
    var v4 bool = (32767 >= -1)
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int32 = v3
        fmt.Println(v1)
        _ = fi9
        _ = v10
    }
    switch ("" + "foo bar") {
    case "a":
        var v11 uint8 = v2
        _ = v11
    case "hello":
        v12 := []int{335, 0}
        _ = v12
    }
    var v13 uint16 = 414
    var v14 uint = 352
    v15 := []int{943, 706, -2147483648, 32767}
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 float64 = (0.0 + float64((0.1 * ((3.14159 + -2.25) / (3.14159 / 0.1)))))
        for fi18 := 0; fi18 < 14; fi18++ {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            if v4 {
                _ = len(v15)
                v19 := GenS1{F0: 1, F1: 32768, F2: 152}
                _ = v19
            }
            var v20 float32 = ((-2.25 + -2.25) / (-2.25 - 1.5))
            _ = fi18
            _ = v20
        }
        v3 -= int32(v13)
        _ = fi16
        _ = v17
    }
    var v21 uint8 = (v2 - v2)
    var v22 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
