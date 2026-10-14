package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
func genFn0(p0 uint, p1 int64) int32 {
    switch 32767 {
    case -2147483648:
        if ("\n\t" < "foo bar") {
            p0 = p0
        }
    case 42:
        switch -2147483648 {
        case -36:
            p0 += (p0 * 1)
        }
    case 2147483647:
        for fi6 := 0; fi6 < 4; fi6++ {
            if ((((true || (!(!false))) && (!((!false) || true))) || false) || true) {
                p0 -= (^(p0))
                v7 := GenS1{F0: 1, F1: 42}
                _ = v7
            }
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS1{F0: 2147483648, F1: (p0 << (((p0 / ((42) | 1))) & 63))}
                var v10 int64 = (^(p1))
                var v11 uint16 = 0
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi6
        }
    }
    var v12 float64 = 1.5
    _ = p0
    _ = p1
    _ = v12
    return int32(p1)
}
func genFn1(p0 bool, p1 int64, p2 uint16) uint {
    var v13 int32 = -2147483648
    p0 = p0
    v14 := []int{127, 266, -9223372036854775808, 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    v3 := []uint32{728, 42, 2147483647}
    var v4 float32 = float32(0.0)
    var v5 uint32 = 946
    v5 = v5
    switch "foo bar" {
    case "\n\t":
        for fi15 := 0; fi15 < 10; fi15++ {
            v4 -= -2.25
            _ = fi15
        }
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := GenS1{F0: v5, F1: 65535}
        _ = fi16
        _ = v17
    }
    _ = len(v2)
    for fi18 := 0; fi18 < 20; fi18++ {
        v4 += (float32(1e10) - (v4 / -2.25))
        _ = len(v2)
        _ = fi18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint = 0
        var v21 uint16 = 65535
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS1{F0: (v5 | 422), F1: v20}
            var v24 int = (-(fi22))
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v5, v4)
    v5 = (v5 * v5)
    var v25 bool = (!true)
    for fi26 := 0; fi26 < 11; fi26++ {
        switch ("hello" + "hello") {
        case "a":
            v27 := GenS1{F0: 4294967295, F1: genFn1(((3.14159 / 1.5) > 0.0), -597, 238)}
            _ = v27
        case "foo bar":
            var v28 uint16 = 65535
            _ = v28
        case "":
            v25 = (!(("foo bar" >= "") && (!(float32((0.0 * 0.1)) <= v4))))
        }
        var v29 uint8 = 2
        v30 := GenS1{F0: 2, F1: 163}
        _ = fi26
        _ = v29
        _ = v30
    }
    fmt.Println(v25)
    v31 := GenS1{F0: 2, F1: 18446744073709551615}
    fmt.Println(v5, v25)
    v32 := GenS1{F0: (v5 % (((v5 / ((1) | 1))) | 1)), F1: genFn1((!v25), 1, 65535)}
    var v33 uint32 = v5
    v33 += v33
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
