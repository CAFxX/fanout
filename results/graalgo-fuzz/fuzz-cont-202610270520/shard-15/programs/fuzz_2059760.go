package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint64; F2 int32 }
func genFn0(p0 uint, p1 bool, p2 bool) string {
    var v7 bool = (910 >= 255)
    v8 := GenS1{F0: p0}
    var v9 uint = (p0 - p0)
    switch ("\n\t" + "hello") {
    case "a":
        var v10 int32 = 787
        _ = v10
    case "hello":
        var v11 bool = p2
        _ = v11
    default:
        switch -1 {
        case 1:
            v12 := GenS2{F0: p0, F1: 0, F2: 2147483647}
            _ = v12
        case 42:
            p0 = 255
        case -1:
            v9 = 46
        }
    }
    if v7 {
        v7 = p1
        p1 = (((0.0 - (0.0 - -2.25)) / 3.14159) != ((-2.25 - 1e10) / 3.14159))
    } else {
        var v13 int = -2147483648
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int32 = 0
    var v2 int16 = -356
    var v3 int64 = -495
    var v4 int16 = (v2 ^ (v2 / ((v2) | 1)))
    var v5 uint = 0
    var v6 int = -2147483648
    var v14 uint64 = 98
    var v15 uint = (v5 * 4294967295)
    for fi16 := 0; fi16 < 4; fi16++ {
        v5 = v5
        _ = fi16
    }
    switch 127 {
    case 9223372036854775807:
        v17 := GenS1{F0: (^(v15))}
        _ = v17
    case 1:
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                v20 := []int{int(v5), (-(v6)), (fi19 >> (((v5 << ((v15) & 63))) & 63)), (fi18 ^ v6)}
                _ = fi19
                _ = v20
            }
            var v21 uint = 837
            var v22 int8 = -43
            _ = fi18
            _ = v21
            _ = v22
        }
    default:
        var v23 bool = (2147483647 != 42)
        _ = v23
    }
    v14 += (v14 ^ (v14 - (v14 >> ((18446744073709551615) & 63))))
    v24 := []int{(v6 << ((65535) & 63)), v6, (v6 >> ((v5) & 63)), (v6 & 128)}
    v24 = append(v24, int(v3))
    v1 = (v1 - (v1 & 2147483647))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
