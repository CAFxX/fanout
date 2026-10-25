package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int8 }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -2147483648
    var v3 float32 = (0.1 + 3.14159)
    v4 := GenS1{F0: false, F1: -107, F2: 60}
    var v5 uint16 = 32767
    v3 += 0.0
    if (!false) {
        var v6 uint64 = 1
        fmt.Println(v2, v5)
        v3 -= 0.0
        _ = v6
    }
    v1[("hello" + "")] = 2147483648
    switch 270 {
    case 32768:
        v7 := GenS1{F0: false, F1: -1, F2: 100}
        _ = v7
    case 127:
        var v8 float32 = v3
        _ = v8
    case -426:
        var v9 uint32 = 42
        _ = v9
    default:
        v10 := GenS2{F0: v5, F1: 72}
        _ = v10
    }
    var v11 uint16 = (^(v5))
    v3 += v3
    var v12 float64 = 0.1
    if (536 > 0) {
        v13 := []int{427, -1}
        var v14 int = 4294967295
        var v15 float64 = float64(v3)
        _ = v13
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 7; fi16++ {
            v11 -= v11
            switch "" {
            case "\n\t":
                switch ("" + "") {
                case "hello":
                    v17 := GenS1{F0: (!(!(1 > 1))), F1: -1, F2: -1}
                    _ = v17
                case "\n\t":
                    v18 := GenS2{F0: v11, F1: 1}
                    _ = v18
                case "foo bar":
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v5 = (v11 ^ v11)
                        fmt.Println(v5)
                        _ = fi19
                    }
                }
            case "":
                var v20 int8 = -128
                _ = v20
            }
            v21 := GenS2{F0: (v11 % (((v5 / (((v5 << ((0) & 63))) | 1))) | 1)), F1: 1000}
            _ = fi16
            _ = v21
        }
    }
    fmt.Println(v11, v5, v3)
    v11 = (v11 + (v5 + 32767))
    fmt.Println(v11, v11, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
