package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0(p0 uint16) int8 {
    fmt.Println(p0)
    var v5 bool = (false || (!(0 <= -1)))
    _ = p0
    _ = v5
    return 111
}
func genFn1(p0 int8, p1 uint32) uint32 {
    v6 := []int{0, 142, -9223372036854775808}
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 float32 = (-2.25 * 1e10)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v10 float32 = v9
            _ = fi8
            _ = v10
            _ = v9
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
    }
    fmt.Println(p0)
    var v11 float32 = 0.1
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return (p1 >> ((4294967295) & 63))
}

func main() {
    var v1 float64 = 1.5
    var v2 uint16 = 42
    v3 := []uint32{4294967295, 2147483647, 681}
    var v4 int16 = 0
    var v12 uint32 = 2147483648
    switch (("" + ("" + ("a" + "\n\t"))) + "a") {
    case "\n\t":
        var v13 int64 = 902
        _ = v13
    case "":
        var v14 int16 = (v4 ^ v4)
        _ = v14
    }
    var v15 int64 = 65
    v16 := GenS1{F0: v4}
    var v17 uint16 = (^(v2))
    var v18 bool = false
    v4 = (v4 >> ((18446744073709551615) & 63))
    var v19 bool = false
    v20 := GenS1{F0: (v4 << ((255) & 63))}
    switch -353 {
    case -216:
        v21 := GenS2{F0: 0, F1: 9223372036854775807}
        _ = v21
    case 0:
        fmt.Println(v17, v4)
    default:
        var v22 uint = 0
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
