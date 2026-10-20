package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool }

func main() {
    var v1 int = -195
    v2 := []string{"\n\t", "\n\t", "hello"}
    var v3 int16 = -1
    v4 := []string{"hello"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v5 uint64 = 42
    for fi6 := 0; fi6 < 18; fi6++ {
        switch v1 {
        case 813:
            v7 := GenS1{F0: 255, F1: v3, F2: false}
            _ = v7
        case -523:
            switch (v1 + (fi6 << ((2) & 63))) {
            case 42:
                var v8 float32 = ((float32(1e10) - 0.1) + 1.5)
                _ = v8
            case 725:
                v5 += (v5 >> ((137) & 63))
            case -609:
                if (!(255 > 739)) {
                    fmt.Println(v5, v5)
                    v9 := GenS1{F0: 2, F1: 0, F2: false}
                    _ = v9
                } else {
                    v3 = v3
                    v1 = fi6
                }
            }
        default:
            var v10 float64 = 0.1
            _ = v10
        }
        _ = fi6
    }
    if true {
        v11 := GenS2{F0: ("a" + "hello"), F1: (!true)}
        var v12 int8 = -128
        switch "\n\t" {
        case "hello":
            var v13 int8 = (v12 | (-(v12)))
            _ = v13
        case "":
            var v14 bool = ((-123 >= -2147483648) || true)
            _ = v14
        case "a":
            var v15 float64 = (3.14159 / 1.5)
            _ = v15
        }
        _ = v11
        _ = v12
    } else {
        _ = len(v2)
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int16 = 42
        var v18 int = (v1 ^ v1)
        _ = fi16
        _ = v17
        _ = v18
    }
    if ((295 <= 2147483647) && (!(true && false))) {
        v1 = (v1 | v1)
    }
    var v19 uint = 255
    v19 += (v19 | 18446744073709551615)
    v1 = (v1 & (^(v1)))
    v20 := GenS3{F0: ((822 >= 1) || true)}
    v21 := GenS1{F0: 768, F1: (v3 - (-(v3))), F2: (false && (!true))}
    var v22 float64 = (float64(0.1) / 3.14159)
    for fi23 := 0; fi23 < 14; fi23++ {
        fmt.Println(v3, v22)
        _ = fi23
    }
    var v24 uint64 = v5
    if false {
        fmt.Println(v22, v3, v1)
        v25 := GenS2{F0: ("" + "hello"), F1: false}
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
