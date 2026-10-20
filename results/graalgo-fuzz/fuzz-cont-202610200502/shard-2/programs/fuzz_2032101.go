package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int8; F1 int32 }
func genFn0(p0 int16, p1 int8) uint8 {
    p1 = 3
    fmt.Println(p1)
    switch 42 {
    case -570:
        v7 := []int{9223372036854775807, 0, -953}
        _ = v7
    case 32768:
        var v8 int = 2147483647
        _ = v8
    }
    var v9 bool = (!false)
    p1 += 13
    _ = p0
    _ = p1
    _ = v9
    return 1
}
func genFn1(p0 uint16, p1 bool, p2 int) string {
    if (0 < 449) {
        v10 := []int{}
        _ = v10
    }
    if true {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint16 = 0
            _ = fi11
            _ = v12
        }
    } else {
        var v13 string = ("foo bar" + "hello")
        var v14 uint8 = genFn0(20, -1)
        _ = v13
        _ = v14
    }
    v15 := []int{p2, 565, p2, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return ""
}

func main() {
    var v1 string = ""
    v2 := []int{}
    var v3 int = 128
    v4 := []string{}
    var v5 int16 = 1000
    var v6 uint32 = 65535
    fmt.Println(v5, v3, v3)
    switch v1 {
    case "\n\t":
        switch "a" {
        case "hello":
            v5 -= v5
        default:
            v16 := GenS1{F0: 632}
            _ = v16
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if ((-1 == 42) || false) {
        v5 -= v5
    }
    if (2147483647 <= -2147483648) {
        if false {
            v17 := GenS1{F0: 0}
            var v18 int16 = v5
            switch (v1 + v1) {
            case "foo bar":
                v19 := GenS1{F0: 924}
                _ = v19
            case "":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "a":
                var v20 int32 = 42
                _ = v20
            }
            _ = v17
            _ = v18
        } else {
            v1 = "hello"
            v6 -= 422
        }
        switch "a" {
        case "a":
            v2 = append(v2, v3)
        case "hello":
            var v21 float64 = (0.1 - 1.5)
            _ = v21
        }
    } else {
        v22 := GenS1{F0: 4294967295}
        if true {
            v6 = (v6 & v6)
            v5 += -111
        }
        _ = v22
    }
    v23 := GenS2{F0: v6, F1: 73}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v4 = append(v4, "")
        v25 := GenS2{F0: (v6 - v6), F1: -1}
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: 127, F1: -2147483648}
    var v27 int16 = (v5 / ((-1) | 1))
    if (127 == 42) {
        var v28 uint64 = 42
        v6 = v6
        _ = v28
    }
    _ = len(v4)
    v5 += (-(v27))
    var v29 string = "foo bar"
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 int64 = -9223372036854775808
        _ = fi30
        _ = v31
    }
    fmt.Println(v6, v3)
    v4 = append(v4, "")
    v2 = append(v2, (v3 >> ((4294967295) & 63)))
    if ((v5 << ((18446744073709551615) & 63)) == v27) {
        v6 += v6
    } else {
        v29 = "foo bar"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
