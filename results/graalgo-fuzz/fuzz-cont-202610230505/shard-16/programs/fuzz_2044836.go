package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int; F1 string; F2 uint16 }
func genFn0(p0 int16, p1 uint8, p2 uint64) string {
    p2 -= p2
    if (18446744073709551615 >= 65535) {
        v7 := []int{42, 504}
        switch 0 {
        case 9223372036854775807:
            var v8 string = ("" + "a")
            _ = v8
        case -9223372036854775808:
            var v9 int32 = -2147483648
            _ = v9
        default:
            fmt.Println(p0)
        }
        _ = v7
    }
    var v10 string = "foo bar"
    v11 := GenS1{F0: 127, F1: (p0 << ((65535) & 63))}
    switch 428 {
    case 4294967295:
        switch "foo bar" {
        case "\n\t":
            v10 = (v10 + v10)
        case "a":
            var v12 bool = true
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return v10
}
func genFn1(p0 uint, p1 string) int32 {
    fmt.Println(p1, p0)
    p1 = p1
    switch 127 {
    case 4294967295:
        p1 = p1
    default:
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS1{F0: 42, F1: -233}
            var v15 float64 = -2.25
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    var v16 string = "foo bar"
    switch (p1 + (((v16 + "foo bar") + v16) + p1)) {
    case "a":
        fmt.Println(p1, p0, p1)
    }
    _ = p0
    _ = p1
    _ = v16
    return -2147483648
}

func main() {
    var v1 uint64 = 659
    var v2 int32 = 1
    v3 := []uint32{65535, 802, 1, 2}
    var v4 int16 = 32767
    var v5 int16 = 1
    var v6 int32 = v2
    var v17 uint = 18446744073709551615
    fmt.Println(v1)
    if (!((false && false) && false)) {
        v18 := GenS1{F0: -83, F1: (-(v5))}
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(v6, v2)
        v20 := GenS2{F0: (fi19 << (((v17 * 313)) & 63)), F1: "foo bar", F2: 2}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 int64 = -675
        fmt.Println(v4, v2, fi21)
        _ = fi21
        _ = v22
    }
    _ = len(v3)
    var v23 uint32 = 11
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v24 float64 = (1e10 + (float64(-2.25) * -2.25))
    if false {
        v6 -= (v6 + (^(v6)))
        var v25 string = "a"
        _ = v25
    }
    if (v17 > 65535) {
        v26 := GenS2{F0: 32767, F1: "\n\t", F2: 32767}
        var v27 int8 = 127
        v4 = (v5 / ((-1) | 1))
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
