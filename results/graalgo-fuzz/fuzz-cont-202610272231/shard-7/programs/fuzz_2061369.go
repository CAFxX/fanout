package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 bool }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint32, p1 uint64, p2 bool) uint {
    v4 := GenS1{F0: 128, F1: 234, F2: 42}
    for fi5 := 0; fi5 < 3; fi5++ {
        switch ((("a" + "") + "hello") + "a") {
        case "\n\t":
            fmt.Println(p2, p2, p2)
        case "foo bar":
            var v6 int16 = 42
            _ = v6
        case "hello":
            switch fi5 {
            case 894:
                if (float64(3.14159) >= 0.1) {
                    if p2 {
                        for fi7 := 0; fi7 < 10; fi7++ {
                            p1 -= (p1 + (p1 - p1))
                            fmt.Println(p0, fi7, fi5)
                            _ = fi7
                        }
                        if (0.0 > 3.14159) {
                            p2 = (!(-1 == 42))
                            if (42 == -1) {
                                p0 += (p0 / ((p0) | 1))
                                v8 := []int{(fi5 * fi5), fi5, fi5, fi5}
                                _ = v8
                            } else {
                                for fi9 := 0; fi9 < 4; fi9++ {
                                    var v10 uint = 42
                                    v11 := GenS1{F0: 127, F1: 128, F2: 128}
                                    var v12 int = (fi9 * (fi5 / (((fi5 - (fi5 & fi9))) | 1)))
                                    p1 = p1
                                    _ = fi9
                                    _ = v10
                                    _ = v11
                                    _ = v12
                                }
                            }
                        } else {
                            v13 := GenS2{F0: 42, F1: 2, F2: (-9223372036854775808 < 9223372036854775807)}
                            _ = v13
                        }
                        p1 = p1
                    } else {
                        v14 := GenS2{F0: p0, F1: (p0 / ((p0) | 1)), F2: (1e10 != (0.0 + -2.25))}
                        v15 := []int{2147483647, (fi5 ^ fi5)}
                        p0 += p0
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS3{F0: 1, F1: 42, F2: 4294967295}
                    v17 := []int{}
                    _ = v16
                    _ = v17
                }
            case -1:
                if p2 {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        fmt.Println(p1, p1)
                        v19 := GenS2{F0: (p0 + 65535), F1: 4294967295, F2: p2}
                        _ = fi18
                        _ = v19
                    }
                }
            }
        default:
            p0 += 65535
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}

func main() {
    var v1 uint16 = 255
    var v2 int64 = 1
    v3 := make(map[int]uint32)
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := []int{}
        v22 := GenS2{F0: 583, F1: 2, F2: true}
        _ = fi20
        _ = v21
        _ = v22
    }
    v2 = v2
    v1 = (v1 & v1)
    fmt.Println(v1, v2)
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := GenS2{F0: 0, F1: 0, F2: (v1 < v1)}
        for fi25 := 0; fi25 < 6; fi25++ {
            fmt.Println(fi23, fi25, v2)
            _ = fi25
        }
        v2 = v2
        _ = fi23
        _ = v24
    }
    var v26 int = 1
    v27 := []int{v26, (^(v26)), 904}
    fmt.Println(v26, v26, v26)
    if (383 == 674) {
        var v28 int16 = -839
        v29 := GenS4{F0: 0, F1: 18446744073709551615}
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
