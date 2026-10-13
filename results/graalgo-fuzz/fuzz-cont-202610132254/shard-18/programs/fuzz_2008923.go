package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 string }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 string; F1 int32; F2 int64 }
func genFn0(p0 int, p1 int64) uint16 {
    v7 := GenS1{F0: 32767, F1: 214, F2: "foo bar"}
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(p0, p0, p0)
        _ = fi8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        p0 -= p0
        _ = fi9
    }
    v10 := GenS1{F0: 1000, F1: 255, F2: ""}
    p1 = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 42
}
func genFn1(p0 int64) int16 {
    p0 = -9223372036854775808
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS2{F0: 32768, F1: 483}
        v13 := []int{(fi11 | (fi11 >> ((569) & 63))), -2147483648, -675, fi11}
        p0 -= (-(p0))
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    return -1
}

func main() {
    var v1 uint16 = 0
    var v2 float64 = (3.14159 * float64(1.5))
    var v3 int8 = 1
    var v4 int16 = -1
    v5 := []uint32{4294967295, 0, 4294967295}
    var v6 uint32 = 0
    if (("foo bar" + "hello") == ("" + "foo bar")) {
        if (!false) {
            v5 = append(v5, (v6 ^ v6))
            v4 = v4
            var v14 int = 32768
            _ = v14
        }
        if ((-2.25 / 1.0) > 1e10) {
            fmt.Println(v6, v6, v1)
            switch -557 {
            case 127:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v1 += 679
        }
    }
    var v15 int = -2147483648
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 bool = true
        var v18 uint = 42
        _ = fi16
        _ = v17
        _ = v18
    }
    v6 += v6
    v3 += -1
    var v19 int64 = -638
    if false {
        v20 := GenS2{F0: -113, F1: 65535}
        v21 := GenS3{F0: ("foo bar" + "\n\t"), F1: -670, F2: -1}
        fmt.Println(v2, v4, v6)
        _ = v20
        _ = v21
    } else {
        var v22 string = "foo bar"
        _ = v22
    }
    switch "hello" {
    case "\n\t":
        var v23 int = int(v1)
        _ = v23
    default:
        v24 := GenS3{F0: "a", F1: 1, F2: 42}
        _ = v24
    }
    if true {
        if len(v5) > 0 {
            v5[0] = v6
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
