package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 uint64; F2 int16 }
func genFn0(p0 uint16) int64 {
    v4 := []int{1, 185, 939, 4294967295}
    p0 -= (p0 ^ 0)
    _ = p0
    _ = v4
    return -1
}
func genFn1() int32 {
    var v5 float64 = 1.5
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: fi6, F1: 42, F2: 127}
        var v8 uint16 = 32767
        v9 := GenS1{F0: (^(fi6)), F1: genFn0((v8 % ((255) | 1))), F2: 128}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    if (65535 < 65535) {
        switch 9223372036854775807 {
        case -868:
            var v10 uint64 = 42
            _ = v10
        case 151:
            switch "\n\t" {
            case "foo bar":
                var v11 uint64 = 1
                _ = v11
            case "hello":
                v5 -= float64(1.5)
            }
        case -746:
            if true {
                for fi12 := 0; fi12 < 19; fi12++ {
                    v5 -= ((v5 - v5) + (((v5 * v5) - ((v5 / (v5 - (v5 * -2.25))) - v5)) - (float64(3.14159) / 1.0)))
                    v13 := []int{-1, (fi12 * 4294967295), -60}
                    fmt.Println(v5, v5)
                    _ = fi12
                    _ = v13
                }
            }
        }
    } else {
        v14 := []int{-9223372036854775808, 0}
        _ = v14
    }
    var v15 float64 = 3.14159
    _ = v15
    _ = v5
    return -1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 42
    var v3 uint32 = 0
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS2{F0: -36, F1: 42, F2: 888}
        v3 = (^(v3))
        var v18 string = "hello"
        _ = fi16
        _ = v17
        _ = v18
    }
    v3 = 65535
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            v3 += (v3 - v3)
            if (18446744073709551615 > 18446744073709551615) {
                var v21 float32 = -2.25
                var v22 float32 = -2.25
                _ = v21
                _ = v22
            } else {
                if ((-9223372036854775808 != 1) && (((3.14159 + -2.25) - (-2.25 / (float64(-2.25) - 1.5))) > float64((((0.0 + 0.1) * 1e10) + 1.5)))) {
                    v2 = (v2 * 42)
                } else {
                    v2 = (v2 >> ((93) & 63))
                }
                v23 := GenS1{F0: -803, F1: 9223372036854775807, F2: uint8(v3)}
                _ = v23
            }
            v2 = (v2 % ((19) | 1))
            fmt.Println(fi20, v2)
            _ = fi20
        }
        var v24 uint = 42
        v25 := GenS2{F0: -52, F1: 0, F2: (v3 * 65535)}
        _ = fi19
        _ = v24
        _ = v25
    }
    if (!true) {
        switch 2147483647 {
        case -686:
            if false {
                for fi26 := 0; fi26 < 12; fi26++ {
                    v27 := GenS2{F0: 67, F1: -1, F2: 4294967295}
                    for fi28 := 0; fi28 < 15; fi28++ {
                        fmt.Println(v3, fi28, fi28)
                        switch "" {
                        case "hello":
                            if (-2147483648 != 2147483647) {
                                var v29 int32 = 0
                                _ = v29
                            }
                        case "foo bar":
                            var v30 int64 = -624
                            _ = v30
                        }
                        v3 -= (v3 * 457)
                        _ = fi28
                    }
                    _ = fi26
                    _ = v27
                }
            }
        case 735:
            switch 2147483647 {
            case 127:
                for fi31 := 0; fi31 < 18; fi31++ {
                    v3 += (v3 + (v3 | 42))
                    v3 -= (^(v3))
                    _ = fi31
                }
            case -9223372036854775808:
                var v32 uint64 = 23
                _ = v32
            case 128:
                v3 += (^(v3))
            }
        case 434:
            delete(v1, v3)
        default:
            var v33 string = "foo bar"
            _ = v33
        }
        v34 := []int{9223372036854775807, 9223372036854775807}
        _ = v34
    }
    v2 = 8
    for fi35 := 0; fi35 < 20; fi35++ {
        if (v2 != v2) {
            var v36 int32 = 670
            _ = v36
        }
        v37 := GenS3{F0: v2, F1: 0}
        v3 -= 976
        _ = fi35
        _ = v37
    }
    v2 -= (v2 & v2)
    v2 = 193
    if true {
        switch "hello" {
        case "":
            v2 += (v2 & v2)
        case "foo bar":
            for fi38 := 0; fi38 < 20; fi38++ {
                v39 := GenS3{F0: v2, F1: 42}
                _ = fi38
                _ = v39
            }
        case "\n\t":
            for fi40 := 0; fi40 < 15; fi40++ {
                v41 := GenS4{F0: (v3 % ((v3) | 1))}
                v2 += v2
                _ = fi40
                _ = v41
            }
        }
        var v42 int = 4294967295
        _ = v42
    }
    if true {
        var v43 uint32 = 4294967295
        v44 := GenS5{F0: -128, F1: 4294967295, F2: 0}
        v3 = 1
        _ = v43
        _ = v44
    } else {
        v45 := GenS2{F0: 127, F1: -1, F2: v3}
        _ = v45
    }
    var v46 uint = 635
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
}
