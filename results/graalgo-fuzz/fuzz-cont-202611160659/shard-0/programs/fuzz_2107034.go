package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 2
    var v3 uint16 = 42
    fmt.Println(v2)
    v1 -= 0
    var v4 float32 = float32((((0.0 * float64(1.5)) + ((1.5 * 1e10) * (0.0 * 3.14159))) / 0.1))
    var v5 uint = 1
    for fi6 := 0; fi6 < 5; fi6++ {
        if ((0 >= (v3 | (^(v3)))) || false) {
            var v7 uint32 = 1
            var v8 uint16 = 19
            v9 := GenS1{F0: 292, F1: (v2 + (v2 & v2)), F2: (v7 & uint32(v5))}
            _ = v7
            _ = v8
            _ = v9
        } else {
            v5 = v5
            v5 += (v5 + v5)
            v1 = 0
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 int32 = v1
        v4 = v4
        v11 -= v11
        _ = fi10
        _ = v11
    }
    v12 := []int{728, -2147483648, 128, 32768}
    if len(v12) > 0 {
        v12[0] = 2147483647
    }
    v1 -= (v1 | 0)
    switch 732 {
    case 2147483647:
        v13 := GenS1{F0: 0, F1: v2, F2: uint32(v5)}
        _ = v13
    case 128:
        if (0 >= 127) {
            if false {
                var v14 bool = (("\n\t" + "foo bar") > "a")
                _ = v14
            }
        } else {
            v15 := GenS1{F0: 1000, F1: uint64(v3), F2: 2147483647}
            var v16 int32 = v1
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 string = "foo bar"
                v12 = append(v12, (fi17 + fi17))
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
    case 4294967295:
        v3 += v3
    default:
        var v19 int32 = v1
        _ = v19
    }
    fmt.Println(v1)
    var v20 bool = false
    if (-1 < -2147483648) {
        var v21 string = ("\n\t" + "hello")
        v2 = 42
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
