package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint64, p1 int, p2 uint64) string {
    var v4 int = p1
    var v5 int = (v4 >> ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "foo bar"
}

func main() {
    var v1 int64 = 495
    var v2 uint = 4294967295
    v3 := make(map[uint32]bool)
    if (18446744073709551615 > 221) {
        for fi6 := 0; fi6 < 13; fi6++ {
            for fi7 := 0; fi7 < 2; fi7++ {
                v2 -= (v2 & (v2 * 42))
                _ = fi7
            }
            _ = fi6
        }
        v8 := v3[2147483648]
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 1}
    v10 := []int{835, -591}
    if (1 > 202) {
        if (2147483647 > 42) {
            var v11 int16 = 42
            var v12 uint32 = 1
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 float64 = ((3.14159 + 0.1) - (float64(3.14159) / 1.5))
                _ = fi13
                _ = v14
            }
            v2 = v2
            if true {
                var v15 uint64 = 267
                if (!(276 < v15)) {
                    v16 := v3[980]
                    var v17 int16 = 143
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS2{F0: v2}
                    _ = v18
                }
                _ = v15
            }
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 uint16 = 1
            _ = fi19
            _ = v20
        }
        var v21 int64 = 9223372036854775807
        _ = v21
    } else {
        var v22 bool = (2147483648 > 2147483647)
        _ = v22
    }
    var v23 uint16 = 32767
    v24 := GenS1{F0: (v23 << ((848) & 63)), F1: 0}
    var v25 uint64 = 2
    var v26 int64 = (v1 ^ v1)
    v3[4294967295] = ((^(v23)) == v23)
    switch "hello" {
    case "\n\t":
        v27 := GenS1{F0: 2, F1: 2147483647}
        _ = v27
    case "hello":
        v3[4294967295] = ("\n\t" == ("\n\t" + ""))
    }
    fmt.Println(v26, v23)
    switch ("a" + "\n\t") {
    case "\n\t":
        var v28 int = -1
        _ = v28
    case "a":
        for fi29 := 0; fi29 < 10; fi29++ {
            v26 = -758
            v25 -= (v25 % ((v25) | 1))
            fmt.Println(v23, v2)
            _ = fi29
        }
    default:
        var v30 int16 = 32767
        _ = v30
    }
    fmt.Println(v25)
    var v31 uint8 = 78
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v9
}
