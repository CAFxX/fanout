package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint16; F2 int }
type GenS4 struct { F0 uint8; F1 int64 }
func genFn0(p0 int32, p1 int) int8 {
    var v4 int64 = 129
    v5 := GenS1{F0: -128, F1: 2}
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int64 = (v4 % ((450) | 1))
        v8 := GenS1{F0: 127, F1: 2147483648}
        if ((true || true) && false) {
            var v9 int = fi6
            var v10 uint32 = 2147483647
            _ = v10
            _ = v9
        } else {
            fmt.Println(v4, p1, p1)
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 127
}
func genFn1(p0 int, p1 uint, p2 uint32) uint32 {
    if false {
        p2 -= (p2 ^ (p2 % ((0) | 1)))
        v11 := []int{(p0 % (((p0 / (((p0 << (((p1 >> ((p1) & 63))) & 63))) | 1))) | 1)), (p0 / ((p0) | 1)), 953}
        p2 = (p2 >> ((2) & 63))
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        if true {
            v13 := GenS2{F0: -9223372036854775808, F1: 760}
            if (42 == genFn0(-142, 9223372036854775807)) {
                for fi14 := 0; fi14 < 17; fi14++ {
                    p1 = (^(p1))
                    _ = fi14
                }
            } else {
                var v15 uint64 = 0
                if false {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v15 = (v15 ^ v15)
                        _ = fi16
                    }
                    switch "\n\t" {
                    case "a":
                        var v17 uint64 = v15
                        _ = v17
                    case "foo bar":
                        p2 = uint32(v15)
                    default:
                        fmt.Println(v15, fi12)
                    }
                }
                _ = v15
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                v19 := []int{(p0 >> (((p1 + p1)) & 63)), (-(p0)), (-(fi12))}
                v20 := GenS1{F0: 87, F1: (p2 >> (((^(p1))) & 63))}
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v13
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"a", ("" + "\n\t"), ("\n\t" + (("foo bar" + ("" + ("a" + "a"))) + "\n\t")), "\n\t"}
    var v3 int8 = 42
    v1["hello"] = false
    var v21 string = ("a" + "hello")
    v21 = ("a" + ("foo bar" + (v21 + ("foo bar" + (v21 + ("" + (v21 + v21)))))))
    switch "hello" {
    case "":
        v22 := GenS3{F0: 2147483647, F1: 987, F2: -1}
        _ = v22
    default:
        _ = len(v1)
    }
    if true {
        if (!(897 == 2147483647)) {
            for fi23 := 0; fi23 < 12; fi23++ {
                v24 := GenS4{F0: 128, F1: -995}
                v3 = -1
                _ = fi23
                _ = v24
            }
        }
        v3 += v3
        for fi25 := 0; fi25 < 11; fi25++ {
            v3 -= genFn0(-219, (^(fi25)))
            var v26 float64 = float64(0.0)
            delete(v1, "")
            _ = fi25
            _ = v26
        }
    }
    if (true || (true || false)) {
        if (!(!(v21 < ""))) {
            var v27 int = 274
            _ = v27
        } else {
            v21 = v21
        }
    }
    v28 := GenS1{F0: (v3 ^ 100), F1: genFn1(745, 0, 635)}
    fmt.Println(v21)
    var v29 string = "foo bar"
    switch v29 {
    case "":
        v30 := GenS3{F0: 2147483647, F1: 32767, F2: -2147483648}
        _ = v30
    case "\n\t":
        v3 -= (v3 << ((1) & 63))
    case "foo bar":
        var v31 int32 = 2147483647
        _ = v31
    default:
        v3 = (v3 / (((^(v3))) | 1))
    }
    v2 = append(v2, "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
}
