package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 uint64) string {
    var v4 float64 = 0.1
    if ((255 <= 4294967295) || false) {
        v4 = v4
    } else {
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: -359, F1: -2147483648}
            v7 := GenS1{F0: 42, F1: 42}
            var v8 bool = (2147483647 <= 2)
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    }
    var v9 int64 = 42
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v9
    return ""
}

func main() {
    var v1 int = -2147483648
    var v2 float32 = 1e10
    var v3 uint8 = 147
    v3 = (^(v3))
    var v10 uint64 = 905
    switch ("a" + ("" + "a")) {
    case "hello":
        if (!(true && (!(false || (733 >= 255))))) {
            switch "" {
            case "\n\t":
                v11 := []int{int(v3), 1, 2147483647}
                _ = v11
            case "":
                v12 := []int{(v1 >> ((255) & 63))}
                _ = v12
            }
        } else {
            v13 := GenS2{F0: 0, F1: 698}
            var v14 float32 = ((1.5 + -2.25) / (v2 - 0.0))
            var v15 float64 = (3.14159 * 1e10)
            _ = v13
            _ = v14
            _ = v15
        }
    default:
        var v16 uint8 = v3
        _ = v16
    }
    var v17 int64 = -259
    v18 := []int{(v1 % ((int(v17)) | 1)), (-(v1)), (^(v1)), (v1 / ((v1) | 1))}
    var v19 float32 = v2
    v20 := GenS2{F0: -2147483648, F1: 18446744073709551615}
    switch 9223372036854775807 {
    case 368:
        switch "hello" {
        case "":
            if (!(!(-697 < v1))) {
                var v21 bool = true
                if len(v18) > 0 {
                    v18[0] = v1
                }
                _ = v21
            } else {
                _ = len(v18)
                var v22 int32 = 915
                _ = v22
            }
        case "hello":
            v18 = append(v18, (v1 + v1))
        default:
            var v23 int8 = 1
            _ = v23
        }
    case 492:
        if (!true) {
            var v24 bool = true
            _ = v24
        }
    default:
        fmt.Println(v3, v3, v17)
    }
    var v25 int32 = 2147483647
    var v26 int8 = 0
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 uint8 = v3
        for fi29 := 0; fi29 < 4; fi29++ {
            v30 := GenS2{F0: v25, F1: 255}
            fmt.Println(v28, v25, v26)
            var v31 uint64 = (v10 << ((975) & 63))
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
}
