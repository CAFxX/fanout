package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint64 }
func genFn0() uint16 {
    var v5 uint = 943
    v6 := []int{-1, -184, 1}
    v7 := GenS1{F0: 1, F1: 627}
    fmt.Println(v5)
    v5 = v5
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint) uint32 {
    v8 := GenS1{F0: 0, F1: -1}
    switch "a" {
    case "":
        switch 32768 {
        case 128:
            var v9 int16 = 0
            _ = v9
        case -1:
            var v10 bool = ((0 > 1) || true)
            _ = v10
        case 928:
            if false {
                for fi11 := 0; fi11 < 6; fi11++ {
                    fmt.Println(fi11, p0)
                    _ = fi11
                }
                if (((p0 << (((p0 / (((p0 & (p0 >> ((0) & 63)))) | 1))) & 63)) >= (p0 ^ (^(p0)))) && true) {
                    v12 := GenS2{F0: 18446744073709551615, F1: 42}
                    _ = v12
                }
                v13 := []int{-901, 32768}
                _ = v13
            } else {
                if false {
                    var v14 uint = p0
                    for fi15 := 0; fi15 < 16; fi15++ {
                        var v16 int = fi15
                        var v17 int8 = 100
                        switch (v16 << (((p0 - p0)) & 63)) {
                        case 84:
                            v14 = v14
                        case 127:
                            var v18 int64 = -973
                            _ = v18
                        case 0:
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v14)
                            }
                        default:
                            switch fi15 {
                            case -1:
                                var v19 uint = 65535
                                _ = v19
                            }
                        }
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = v14
                }
                for fi20 := 0; fi20 < 17; fi20++ {
                    v21 := []int{(fi20 - fi20), (^(fi20)), (fi20 >> (((p0 * 816)) & 63)), (fi20 << ((42) & 63))}
                    _ = fi20
                    _ = v21
                }
            }
        default:
            p0 -= 42
        }
    case "hello":
        p0 = (p0 * (p0 + (p0 << ((p0) & 63))))
    case "foo bar":
        v22 := []int{1, 32767}
        _ = v22
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            if (!true) {
                p0 += 4294967295
            } else {
                var v24 int32 = -1
                _ = v24
            }
            v25 := []int{(fi23 >> ((p0) & 63)), (^(fi23)), (fi23 - -9223372036854775808), (fi23 << ((4294967295) & 63))}
            _ = fi23
            _ = v25
        }
    }
    _ = p0
    _ = v8
    return 0
}

func main() {
    var v1 int32 = 0
    v2 := make(map[string]string)
    v3 := []uint32{1, 42, 4294967295}
    v4 := make(map[int]string)
    switch ("" + (("a" + (("a" + (("foo bar" + ("\n\t" + ("\n\t" + "foo bar"))) + "a")) + "foo bar")) + "hello")) {
    case "a":
        fmt.Println(v1)
    case "":
        var v26 int64 = 1
        _ = v26
    case "foo bar":
        v1 = v1
    }
    if (255 <= 194) {
        var v27 int32 = v1
        fmt.Println(v27, v27)
        _ = v27
    }
    v1 += v1
    switch "a" {
    case "foo bar":
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    case "":
        for fi28 := 0; fi28 < 1; fi28++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v1, v1)
            v29 := GenS1{F0: 760, F1: 9223372036854775807}
            _ = fi28
            _ = v29
        }
    default:
        switch 152 {
        case -9223372036854775808:
            v1 -= v1
        case 9223372036854775807:
            fmt.Println(v1)
        default:
            if true {
                _ = len(v3)
            }
        }
    }
    v1 -= (v1 << ((42) & 63))
    v30 := GenS2{F0: 42, F1: 2}
    var v31 uint = 2
    var v32 int64 = -1
    switch int(v31) {
    case -914:
        for fi33 := 0; fi33 < 9; fi33++ {
            fmt.Println(v31, v32, v1)
            v1 += v1
            var v34 uint16 = genFn0()
            _ = fi33
            _ = v34
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v1 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
