package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 uint; F2 uint16 }
func genFn0(p0 uint16, p1 string, p2 string) uint8 {
    var v6 float32 = 1.5
    v7 := GenS1{F0: 348, F1: 1}
    if true {
        if (((!(!true)) || ((!((!(((float32(float64((((0.1 * 0.1) - 0.1) / (v6 - float32(0.0))))) / (float32((((float64(0.1) / 1e10) * (3.14159 - float64(v6))) + 1.5)) + (v6 - (1e10 / v6)))) != -2.25) && (1 > 9223372036854775807))) && true)) && (42 != 0))) || false) {
            var v8 uint16 = (p0 * (p0 & p0))
            p1 = p1
            var v9 int16 = 1
            _ = v8
            _ = v9
        }
        v10 := []int{-2147483648, -25, 0, -9223372036854775808}
        _ = v10
    } else {
        v11 := []int{2147483647, 32767, -721, 2147483647}
        _ = v11
    }
    p1 = p1
    var v12 int = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint = 0
    v2 := []string{}
    var v3 float32 = (0.1 + 1e10)
    v4 := []uint32{35, 2147483648}
    var v5 int64 = -497
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int = (fi13 % ((fi13) | 1))
        if false {
            var v15 uint8 = 253
            for fi16 := 0; fi16 < 14; fi16++ {
                v1 = v1
                fmt.Println(v3, v14, v5)
                var v17 bool = true
                var v18 float32 = v3
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        } else {
            var v19 int = -9223372036854775808
            _ = v19
        }
        var v20 uint16 = 702
        _ = fi13
        _ = v14
        _ = v20
    }
    var v21 string = ((("hello" + "hello") + "\n\t") + "\n\t")
    fmt.Println(v21)
    v22 := GenS1{F0: v5, F1: 100}
    var v23 float64 = 1.5
    var v24 bool = (v1 == (v1 - v1))
    if ((v3 + (1.5 - 0.0)) <= 1e10) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v25 int = -2147483648
        _ = v25
    }
    var v26 uint64 = 18446744073709551615
    if true {
        fmt.Println(v26, v5)
    }
    if v24 {
        switch 32768 {
        case 368:
            v23 = (v23 * (v23 * 3.14159))
        }
    } else {
        v27 := GenS1{F0: v5, F1: 42}
        v28 := GenS2{F0: 1, F1: (v1 % ((0) | 1)), F2: 42}
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
