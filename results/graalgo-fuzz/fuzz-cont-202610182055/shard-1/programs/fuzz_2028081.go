package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{}
    var v2 uint16 = 372
    var v3 uint16 = (v2 * (v2 * (v2 & v2)))
    v4 := GenS1{F0: ((!(65535 < 896)) || false), F1: 1, F2: 1}
    if (false && true) {
        switch "" {
        case "hello":
            fmt.Println(v2, v3)
        default:
            var v5 string = "foo bar"
            _ = v5
        }
    }
    if false {
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := GenS1{F0: ((v3 >= (v3 - v3)) && false), F1: int8(fi6), F2: 4294967295}
            if len(v1) > 0 {
                v1[0] = 128
            }
            _ = fi6
            _ = v7
        }
    }
    if (495 <= 2) {
        if true {
            var v8 int16 = 42
            var v9 uint8 = 1
            v3 = (v3 - v3)
            _ = v8
            _ = v9
        }
        if true {
            for fi10 := 0; fi10 < 8; fi10++ {
                fmt.Println(v3)
                _ = fi10
            }
            v11 := GenS2{F0: 0}
            _ = v11
        } else {
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
            if true {
                v2 = (^(v2))
                var v12 string = "foo bar"
                var v13 uint64 = 466
                _ = v12
                _ = v13
            }
            var v14 uint = 829
            _ = v14
        }
    }
    v15 := GenS1{F0: false, F1: 42, F2: 705}
    var v16 float64 = float64(3.14159)
    var v17 float64 = (v16 - (v16 + v16))
    if (2 <= 2147483648) {
        v3 = (^(v2))
        v1 = append(v1, 1)
    } else {
        switch 1 {
        case 1:
            var v18 bool = true
            _ = v18
        case 32767:
            v2 = (v3 / ((v2) | 1))
        default:
            var v19 float32 = (3.14159 - 1.5)
            _ = v19
        }
        _ = len(v1)
        switch 127 {
        case -2147483648:
            v16 += (float64(float32(1e10)) - ((v16 / (((float64(0.1) - 1e10) * v16) * v16)) - v16))
        case 9223372036854775807:
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 int64 = -9223372036854775808
                var v22 bool = ((((-(v21)) > int64(fi20)) || true) && true)
                _ = fi20
                _ = v21
                _ = v22
            }
        case 42:
            v16 += (float64((float32(((-2.25 * float64(float32((1.5 - 0.1)))) / 1.0)) / (-2.25 + 1e10))) - v17)
        default:
            fmt.Println(v17, v2)
        }
    }
    var v23 int64 = 9223372036854775807
    v24 := GenS1{F0: ((89 <= 219) && (3.14159 >= (-2.25 * (0.1 - float32(v16))))), F1: 127, F2: 2}
    _ = len(v1)
    var v25 string = "\n\t"
    var v26 uint8 = 127
    for fi27 := 0; fi27 < 10; fi27++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi27
    }
    var v28 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}
