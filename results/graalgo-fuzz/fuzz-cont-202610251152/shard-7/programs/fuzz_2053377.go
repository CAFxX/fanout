package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int64; F1 uint8 }

func main() {
    v1 := []uint32{2147483647, 2}
    var v2 uint = 255
    var v3 int16 = 1000
    v4 := make(map[string]uint32)
    var v5 int = -971
    var v6 int8 = -128
    fmt.Println(v3, v6, v6)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        v6 -= 0
        _ = len(v4)
        _ = fi7
    }
    v8 := GenS1{F0: -976}
    v9 := GenS1{F0: -1}
    var v10 int = v5
    switch ("\n\t" + "\n\t") {
    case "a":
        if true {
            v10 += v5
            delete(v4, "")
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                v12 := GenS2{F0: false, F1: true}
                for fi13 := 0; fi13 < 2; fi13++ {
                    var v14 int32 = 765
                    fmt.Println(v10, v5, fi13)
                    _ = fi13
                    _ = v14
                }
                _ = fi11
                _ = v12
            }
        }
    case "\n\t":
        var v15 float32 = 1e10
        _ = v15
    }
    switch -771 {
    case 32768:
        if (("hello" == "\n\t") && true) {
            var v16 int16 = v3
            _ = v16
        } else {
            v17 := GenS1{F0: 1}
            _ = v17
        }
    case -2147483648:
        v18 := GenS2{F0: (v6 != 89), F1: (3.14159 > -2.25)}
        _ = v18
    default:
        v19 := GenS3{F0: -77, F1: 233}
        _ = v19
    }
    v1 = append(v1, 2)
    if false {
        fmt.Println(v2)
    } else {
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := v4[("hello" + "\n\t")]
            var v22 uint8 = 159
            fmt.Println(v2, v6)
            _ = fi20
            _ = v21
            _ = v22
        }
        fmt.Println(v6, v10, v3)
        var v23 bool = false
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
