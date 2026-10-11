package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint; F2 uint32 }
type GenS3 struct { F0 int16; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0(p0 bool, p1 string, p2 uint8) bool {
    var v4 string = ((p1 + "\n\t") + (p1 + p1))
    v5 := GenS1{F0: -1, F1: p2}
    v6 := GenS1{F0: 1, F1: (p2 * 1)}
    var v7 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (613 == 255)
}

func main() {
    var v1 uint32 = 2147483647
    var v2 uint8 = uint8(v1)
    v3 := make(map[int]bool)
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint32 = (v1 * v1)
        switch "" {
        case "":
            switch "\n\t" {
            case "":
                var v10 int64 = -9223372036854775808
                _ = v10
            case "a":
                v1 += (^(v9))
            }
        case "a":
            fmt.Println(fi8)
        default:
            switch 128 {
            case -546:
                var v11 uint64 = 515
                _ = v11
            default:
                v9 -= (v9 / ((68) | 1))
            }
        }
        _ = fi8
        _ = v9
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v1 = 171
        var v13 int64 = 9223372036854775807
        fmt.Println(v1, v13, v1)
        _ = fi12
        _ = v13
    }
    switch "a" {
    case "hello":
        v14 := GenS2{F0: ((^(v1)) < v1), F1: 4294967295, F2: v1}
        _ = v14
    case "":
        switch 208 {
        case 2147483647:
            var v15 uint = 2
            _ = v15
        case -2147483648:
            for fi16 := 0; fi16 < 10; fi16++ {
                if true {
                    var v17 int32 = 1
                    if ((892 == uint16(v1)) || false) {
                        v18 := GenS3{F0: 355, F1: 710, F2: 18446744073709551615}
                        if false {
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v2)
                            case "\n\t":
                                var v19 float64 = 1e10
                                _ = v19
                            default:
                                v1 -= 950
                            }
                            var v20 int32 = v17
                            v1 = 957
                            _ = v20
                        } else {
                            v1 = uint32(v2)
                            fmt.Println(v1, v2)
                        }
                        var v21 int64 = 0
                        _ = v18
                        _ = v21
                    }
                    var v22 float32 = 1.5
                    _ = v17
                    _ = v22
                } else {
                    v23 := []int{int(v1)}
                    _ = v23
                }
                v24 := GenS1{F0: -9223372036854775808, F1: v2}
                _ = fi16
                _ = v24
            }
        case 646:
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 int32 = -2147483648
                _ = fi25
                _ = v26
            }
        }
    }
    var v27 float64 = (1e10 - 1e10)
    v1 += v1
    v28 := v3[128]
    v2 = v2
    switch 128 {
    case 128:
        if (0 <= v2) {
            v29 := []int{42, -2147483648, 0}
            v29 = append(v29, 32768)
            if len(v29) > 0 {
                v29[0] = 4294967295
            }
            _ = v29
        } else {
            v30 := GenS3{F0: 42, F1: 42, F2: 1}
            _ = v30
        }
    case 850:
        if (v2 == (v2 / ((28) | 1))) {
            v31 := GenS4{F0: (v2 / (((^(v2))) | 1)), F1: true}
            _ = v31
        }
    default:
        if false {
            var v32 string = "\n\t"
            v33 := GenS1{F0: 9223372036854775807, F1: (v2 << ((42) & 63))}
            var v34 uint = 0
            _ = v32
            _ = v33
            _ = v34
        } else {
            v35 := []int{127}
            v2 = (v2 - (v2 - v2))
            var v36 bool = v28
            _ = v35
            _ = v36
        }
    }
    v37 := GenS5{F0: "foo bar"}
    delete(v3, 42)
    for fi38 := 0; fi38 < 20; fi38++ {
        v39 := GenS4{F0: v2, F1: v28}
        var v40 int = 2147483647
        _ = fi38
        _ = v39
        _ = v40
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        v42 := GenS6{F0: uint32(v2), F1: v1, F2: 568}
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v37
}
