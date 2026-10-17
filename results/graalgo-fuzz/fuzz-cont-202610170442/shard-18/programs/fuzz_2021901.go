package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint32; F2 int32 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int32; F1 int32 }
func genFn0(p0 uint, p1 uint32) uint {
    var v7 int64 = 9223372036854775807
    p0 = p0
    v8 := GenS1{F0: 952}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 bool, p1 int8) int {
    p0 = p0
    for fi9 := 0; fi9 < 4; fi9++ {
        p0 = (1 == 2)
        v10 := GenS2{F0: 1, F1: 2147483648, F2: -548}
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 749
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = v11
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 bool = (!(247 != 65535))
    var v4 int8 = 1
    v5 := []uint32{1, 123, 814}
    var v6 uint = 255
    fmt.Println(v3, v6, v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (-353 > 272) {
        switch "\n\t" {
        case "hello":
            if (!(1 <= 127)) {
                for fi12 := 0; fi12 < 10; fi12++ {
                    delete(v2, (-(fi12)))
                    switch genFn1((fi12 < genFn1((v3 || false), v4)), v4) {
                    case 2147483647:
                        v3 = (928 != 632)
                    case 4294967295:
                        fmt.Println(fi12, v3, v3)
                    case -9223372036854775808:
                        var v13 bool = (!v3)
                        _ = v13
                    }
                    var v14 uint64 = 42
                    _ = fi12
                    _ = v14
                }
                for fi15 := 0; fi15 < 16; fi15++ {
                    v6 = 0
                    _ = fi15
                }
            }
        default:
            fmt.Println(v4)
        }
        v4 += (v4 * int8(v6))
        v16 := v2[205]
        _ = v16
    } else {
        var v17 uint16 = 26
        _ = v17
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v18 int64 = -701
    v3 = false
    switch -748 {
    case 42:
        var v19 uint64 = 896
        _ = v19
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            if len(v5) > 0 {
                v5[0] = 15
            }
            _ = fi21
        }
        v22 := v2[fi20]
        if v3 {
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS3{F0: v18, F1: (fi23 & (-(v22)))}
                var v25 bool = ((0 >= 1) && (777 <= 1))
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        _ = fi20
        _ = v22
    }
    v3 = (false || v3)
    var v26 float32 = 0.1
    v4 = (^(v4))
    v27 := GenS4{F0: 1, F1: 21}
    _ = len(v5)
    for fi28 := 0; fi28 < 4; fi28++ {
        if (((!((("foo bar" + "\n\t") >= ("foo bar" + (((("a" + ("foo bar" + (("" + "foo bar") + "hello"))) + ("\n\t" + "foo bar")) + "") + "hello"))) || v3)) || (151 == 2)) && v3) {
            if v3 {
                v29 := GenS1{F0: 2147483647}
                if (v4 >= v4) {
                    var v30 uint16 = 32767
                    v6 = (v6 / (((^(v6))) | 1))
                    _ = v30
                }
                _ = v29
            } else {
                v3 = (false && (65535 >= 32767))
            }
            delete(v1, "foo bar")
            var v31 uint = 65535
            _ = v31
        }
        v26 = v26
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
