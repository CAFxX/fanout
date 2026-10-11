package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint8, p2 int8) bool {
    var v7 float64 = 1e10
    if (false || true) {
        if ((("\n\t" + "a") + "foo bar") >= "\n\t") {
            v8 := GenS1{F0: "foo bar", F1: p2}
            _ = v8
        }
    }
    var v9 float32 = 0.0
    v9 -= v9
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return true
}
func genFn1(p0 uint64) uint16 {
    switch 9223372036854775807 {
    case 2147483647:
        var v10 int32 = 42
        _ = v10
    case -9223372036854775808:
        var v11 float64 = 1.5
        _ = v11
    case -594:
        var v12 uint = 813
        _ = v12
    }
    v13 := []int{127, -9223372036854775808, 127, 0}
    var v14 uint = 478
    v13 = append(v13, 4294967295)
    _ = p0
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := []int{}
    var v2 uint32 = 4294967295
    var v3 float32 = 3.14159
    var v4 uint16 = 42
    var v5 int = -989
    var v6 bool = false
    v15 := GenS1{F0: ("a" + ("" + ("hello" + "foo bar"))), F1: 127}
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v16 uint16 = genFn1(653)
    _ = len(v1)
    v17 := GenS2{F0: 4}
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 int = (v5 % ((fi18) | 1))
        _ = fi18
        _ = v19
    }
    v5 -= (v5 << ((105) & 63))
    for fi20 := 0; fi20 < 4; fi20++ {
        fmt.Println(v16, v2)
        _ = fi20
    }
    fmt.Println(v3, v16, v16)
    var v21 float64 = 1.5
    fmt.Println(v16, v6)
    switch "foo bar" {
    case "a":
        if (!v6) {
            var v22 uint8 = 59
            _ = v22
        }
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
