package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64, p1 int, p2 int64) int16 {
    switch p1 {
    case -168:
        var v4 string = ("" + "foo bar")
        _ = v4
    case 9223372036854775807:
        var v5 uint16 = 1
        _ = v5
    case 1:
        p0 = (p0 * (p0 & (p0 & 0)))
    }
    var v6 int32 = 0
    p0 = (p0 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 int = 32767
    v2 := make(map[int]string)
    var v3 uint8 = 127
    v1 = v1
    if (2 == 255) {
        v7 := GenS1{F0: 608, F1: (v3 << ((2) & 63))}
        for fi8 := 0; fi8 < 11; fi8++ {
            switch "a" {
            case "hello":
                var v9 int16 = 467
                _ = v9
            }
            v3 += (v3 + 2)
            _ = fi8
        }
        v1 = (v1 % ((v1) | 1))
        _ = v7
    }
    v10 := GenS2{F0: 0, F1: 0}
    if false {
        if (-128 > 42) {
            var v11 uint8 = v3
            fmt.Println(v11)
            _ = v11
        }
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint32 = 65535
            _ = fi12
            _ = v13
        }
        switch "\n\t" {
        case "foo bar":
            var v14 uint = 65535
            _ = v14
        case "\n\t":
            if (128 < v3) {
                if true {
                    v15 := GenS2{F0: (v3 - v3), F1: 42}
                    for fi16 := 0; fi16 < 17; fi16++ {
                        fmt.Println(fi16, fi16, v1)
                        if true {
                            v17 := GenS1{F0: 4294967295, F1: 61}
                            switch 626 {
                            case 2147483647:
                                v1 -= (^(fi16))
                            case 471:
                                if false {
                                    v18 := GenS3{F0: -9223372036854775808}
                                    var v19 int = v1
                                    _ = v18
                                    _ = v19
                                } else {
                                    v3 -= (v3 << ((621) & 63))
                                    fmt.Println(v3, v3, v1)
                                }
                            }
                            _ = v17
                        }
                        switch fi16 {
                        case 1:
                            var v20 uint = 0
                            _ = v20
                        case 0:
                            if (0 > 127) {
                                fmt.Println(fi16, fi16, v3)
                                var v21 float32 = 0.1
                                v1 = fi16
                                _ = v21
                            }
                        }
                        _ = fi16
                    }
                    fmt.Println(v1, v3)
                    _ = v15
                }
            }
        }
    } else {
        if (127 <= v1) {
            v22 := GenS1{F0: 2147483648, F1: v3}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 int64 = 210
                v1 += (fi23 - fi23)
                _ = fi23
                _ = v24
            }
        }
        delete(v2, (v1 + (v1 ^ (v1 | v1))))
    }
    var v25 float32 = 1e10
    v26 := []int{(v1 % ((v1) | 1))}
    var v27 bool = (!(-2147483648 != 987))
    v28 := GenS1{F0: 486, F1: v3}
    switch "a" {
    case "":
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    case "a":
        v29 := GenS2{F0: v3, F1: 42}
        _ = v29
    default:
        var v30 int = -9223372036854775808
        _ = v30
    }
    fmt.Println(v27, v27)
    v3 = v3
    v27 = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
