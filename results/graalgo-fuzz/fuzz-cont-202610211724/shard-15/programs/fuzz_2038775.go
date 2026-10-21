package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint64 }
func genFn0() uint {
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{(fi5 | 42), (fi5 & (-(fi5))), fi5, (fi5 | fi5)}
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    var v7 uint = 794
    for fi8 := 0; fi8 < 15; fi8++ {
        switch (fi8 & 4294967295) {
        case -2147483648:
            switch int(v7) {
            case -9223372036854775808:
                var v9 int8 = -72
                _ = v9
            case 486:
                if false {
                    v10 := []int{-9223372036854775808, int(v7), fi8}
                    _ = v10
                }
            }
        default:
            for fi11 := 0; fi11 < 20; fi11++ {
                v7 = (v7 ^ 1)
                v7 += 746
                _ = fi11
            }
        }
        v12 := []int{42, 42, fi8}
        _ = fi8
        _ = v12
    }
    _ = v7
    return 42
}
func genFn1(p0 int, p1 uint32, p2 string) uint64 {
    v13 := []int{p0, 165, p0, (p0 - p0)}
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: -1, F1: p2, F2: 42}
        _ = fi14
        _ = v15
    }
    p2 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 1
}

func main() {
    var v1 float64 = 1e10
    v2 := []uint32{}
    var v3 int64 = 1
    var v4 uint64 = 802
    switch 42 {
    case 552:
        v16 := GenS1{F0: (v3 & (v3 & v3)), F1: (("" + ("" + "hello")) + "hello"), F2: v4}
        _ = v16
    case 4294967295:
        var v17 uint32 = 663
        _ = v17
    case 42:
        v1 -= (v1 - 0.1)
    default:
        var v18 uint32 = 65535
        _ = v18
    }
    v3 = (^(v3))
    v19 := GenS1{F0: 0, F1: "\n\t", F2: (v4 & genFn1(2147483647, 1, (("foo bar" + "") + "\n\t")))}
    if (!(-85 <= 70)) {
        var v20 int32 = 1
        _ = v20
    } else {
        switch "a" {
        case "hello":
            v4 = v4
        case "\n\t":
            v3 += (v3 & (v3 | 9223372036854775807))
        }
        if false {
            for fi21 := 0; fi21 < 5; fi21++ {
                v1 += float64(0.1)
                _ = fi21
            }
            fmt.Println(v1, v1, v3)
        }
    }
    v3 = v3
    fmt.Println(v3, v3, v3)
    v22 := GenS1{F0: v3, F1: "", F2: 1}
    v4 += v4
    var v23 int = -2147483648
    fmt.Println(v3, v4)
    var v24 float32 = 0.1
    switch v23 {
    case 0:
        _ = len(v2)
    default:
        var v25 int8 = 0
        _ = v25
    }
    v2 = append(v2, 4294967295)
    if false {
        v26 := GenS1{F0: (v3 ^ v3), F1: "\n\t", F2: 18446744073709551615}
        fmt.Println(v24, v4, v1)
        for fi27 := 0; fi27 < 3; fi27++ {
            v3 = -664
            _ = fi27
        }
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
