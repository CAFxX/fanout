package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
func genFn0() int32 {
    var v4 int16 = 0
    var v5 string = ("hello" + "")
    fmt.Println(v4, v4)
    v5 = v5
    var v6 float32 = 3.14159
    _ = v4
    _ = v5
    _ = v6
    return -695
}

func main() {
    var v1 int = 0
    var v2 uint32 = 65535
    var v3 bool = true
    v7 := GenS1{F0: -1, F1: 2, F2: 128}
    v8 := GenS1{F0: -1, F1: 206, F2: 222}
    v2 = 2147483647
    v9 := GenS1{F0: -1, F1: 68, F2: 1}
    v10 := GenS2{F0: -9223372036854775808, F1: 0}
    var v11 string = ""
    v11 = "a"
    var v12 int = 9223372036854775807
    v13 := []int{-1, (v12 & (-(v1))), (^(v12))}
    var v14 int8 = 1
    if (v11 <= (v11 + "")) {
        _ = len(v13)
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            fmt.Println(v3)
            v16 := GenS3{F0: -1, F1: -950}
            v13 = append(v13, v12)
            _ = fi15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        v11 = v11
        fmt.Println(v12)
        switch "" {
        case "\n\t":
            if ((-1 != 1) || v3) {
                v18 := GenS4{F0: (v1 + (fi17 & (-(v12))))}
                _ = v18
            } else {
                v14 = 1
            }
        case "a":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        default:
            v13 = append(v13, (v12 / ((fi17) | 1)))
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 int = (v1 * int(v14))
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v22 := GenS5{F0: (1e10 <= 0.0)}
    fmt.Println(v2, v11, v2)
    var v23 int8 = v14
    v24 := GenS3{F0: 0, F1: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
