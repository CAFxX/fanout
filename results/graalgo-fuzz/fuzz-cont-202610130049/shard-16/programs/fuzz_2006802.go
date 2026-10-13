package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }

func main() {
    var v1 uint16 = 0
    var v2 int = 2147483647
    var v3 float32 = 3.14159
    v4 := []int{(v2 << ((678) & 63)), 4294967295}
    var v5 string = ("" + ("" + "hello"))
    var v6 int8 = -22
    for fi7 := 0; fi7 < 10; fi7++ {
        v1 += 437
        var v8 float64 = (((1.5 + 3.14159) * (0.1 * 0.1)) - (((-2.25 * float64(v3)) - -2.25) / (1e10 / 1.0)))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        v6 = v6
        v10 := GenS1{F0: v6, F1: 22, F2: -918}
        v11 := GenS1{F0: 0, F1: 120, F2: -826}
        _ = fi9
        _ = v10
        _ = v11
    }
    if len(v4) > 0 {
        v4[0] = v2
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            v2 += fi13
            _ = fi13
        }
        if true {
            fmt.Println(v2, fi12, v6)
            v2 = v2
        } else {
            v14 := GenS1{F0: (^(v6)), F1: 2, F2: 42}
            for fi15 := 0; fi15 < 7; fi15++ {
                switch v5 {
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = v2
                    }
                default:
                    fmt.Println(v5)
                }
                _ = fi15
            }
            _ = v14
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            _ = len(v4)
            _ = fi16
        }
        _ = fi12
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        if false {
            var v18 string = "\n\t"
            _ = v18
        } else {
            var v19 uint32 = 2147483648
            v1 = (v1 + v1)
            var v20 float64 = -2.25
            _ = v19
            _ = v20
        }
        v21 := GenS2{F0: 255, F1: uint64(v1), F2: ("\n\t" + v5)}
        _ = fi17
        _ = v21
    }
    var v22 uint = 65535
    v23 := GenS1{F0: v6, F1: 127, F2: 2147483647}
    fmt.Println(v3)
    v2 = v2
    v5 = v5
    v6 += v6
    v22 = v22
    fmt.Println(v2, v5)
    v3 = (0.0 * v3)
    v5 = (v5 + v5)
    v24 := GenS2{F0: 100, F1: 2, F2: (("\n\t" + "\n\t") + ((("foo bar" + (v5 + v5)) + ("hello" + (("\n\t" + v5) + (v5 + "a")))) + (v5 + v5)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
