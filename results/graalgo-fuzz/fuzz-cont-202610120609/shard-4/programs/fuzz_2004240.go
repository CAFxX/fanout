package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int64 = -1
    var v2 bool = (!false)
    var v3 uint16 = 255
    var v4 uint32 = uint32(v3)
    v5 := []int{128, 4294967295, -1, 128}
    v6 := []string{"hello", "a", "\n\t", "a"}
    if (242 != 42) {
        if (v2 && (-1 > 471)) {
            for fi7 := 0; fi7 < 12; fi7++ {
                for fi8 := 0; fi8 < 1; fi8++ {
                    _ = len(v5)
                    v1 = 42
                    v3 -= (v3 << ((648) & 63))
                    _ = fi8
                }
                switch -2147483648 {
                case 42:
                    var v9 int = fi7
                    _ = v9
                case -174:
                    v1 = v1
                default:
                    v1 -= 961
                }
                _ = fi7
            }
            var v10 uint16 = 255
            _ = v10
        }
        switch "" {
        case "hello":
            v2 = (v2 || (v2 && ((0.1 + ((float64(0.0) + ((0.1 + (1e10 / 1e10)) / (-2.25 - (0.0 + 0.0)))) * (1.5 + 0.0))) > -2.25)))
        default:
            v11 := GenS1{F0: 412, F1: 42}
            _ = v11
        }
    }
    var v12 int8 = 0
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: 42, F1: 219}
        v3 = v3
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 string = ("" + "a")
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v4 = v4
        _ = fi17
    }
    v2 = (v12 < (^(v12)))
    v4 -= 2147483648
    _ = len(v6)
    var v18 uint8 = 42
    v19 := GenS1{F0: 42, F1: v18}
    for fi20 := 0; fi20 < 9; fi20++ {
        v3 = (^(v3))
        _ = fi20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        if (false && (true && v2)) {
            fmt.Println(v4, v12, v1)
        } else {
            v18 = v18
            _ = len(v5)
        }
        v2 = (!v2)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
