package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
type GenS2 struct { F0 bool }
func genFn0() int64 {
    if false {
        if true {
            switch "foo bar" {
            case "\n\t":
                var v4 uint = 65535
                _ = v4
            }
            var v5 int = 9223372036854775807
            for fi6 := 0; fi6 < 11; fi6++ {
                v5 = (fi6 & (v5 + (fi6 | (^(fi6)))))
                v5 = 32767
                var v7 uint = 541
                _ = fi6
                _ = v7
            }
            _ = v5
        }
    }
    var v8 string = "\n\t"
    v9 := []int{-306}
    var v10 uint32 = 782
    v8 = v8
    _ = v10
    _ = v8
    _ = v9
    return -402
}
func genFn1(p0 int, p1 uint16, p2 string) uint {
    p2 = p2
    var v11 int32 = 1
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := []int{}
        fmt.Println(p2, p0, fi12)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: -9223372036854775808, F1: 31, F2: (p1 | p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v14
    return 762
}

func main() {
    v1 := []int{-572, 42}
    v2 := []int{4294967295}
    v3 := []uint32{}
    var v15 uint32 = 1
    if (0 == -2147483648) {
        for fi16 := 0; fi16 < 17; fi16++ {
            v15 -= (^(v15))
            var v17 int = 32767
            _ = fi16
            _ = v17
        }
        v18 := GenS2{F0: (1 != 0)}
        _ = v18
    } else {
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
        v19 := GenS1{F0: 910, F1: -32768, F2: 2}
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        if len(v2) > 0 {
            v2[0] = fi20
        }
        for fi21 := 0; fi21 < 20; fi21++ {
            v15 = v15
            var v22 uint16 = 2
            _ = fi21
            _ = v22
        }
        if (!true) {
            if (!true) {
                var v23 int16 = 0
                var v24 uint = 1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v23
                _ = v24
            }
            fmt.Println(v15, v15)
        } else {
            v15 = 0
        }
        _ = fi20
    }
    v15 = v15
    v15 = v15
    if (255 <= 42) {
        var v25 int = 869
        var v26 int16 = 1
        for fi27 := 0; fi27 < 20; fi27++ {
            if false {
                if false {
                    v15 += (v15 | v15)
                }
            } else {
                fmt.Println(v26)
                _ = len(v1)
            }
            for fi28 := 0; fi28 < 15; fi28++ {
                switch "foo bar" {
                case "\n\t":
                    var v29 float32 = 0.0
                    _ = v29
                case "a":
                    _ = len(v1)
                    fmt.Println(v26)
                    fmt.Println(fi27)
                default:
                    var v30 uint32 = 42
                    _ = v30
                }
                _ = fi28
            }
            _ = fi27
        }
        _ = v25
        _ = v26
    } else {
        for fi31 := 0; fi31 < 8; fi31++ {
            v15 = (v15 ^ (v15 - v15))
            v15 -= 2147483648
            var v32 int16 = -916
            _ = fi31
            _ = v32
        }
        v15 = 2147483648
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 uint8 = 58
            v2 = append(v2, fi33)
            fmt.Println(v34)
            _ = fi33
            _ = v34
        }
    }
    var v35 int = 127
    var v36 int16 = -1
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 int64 = -1
        _ = fi37
        _ = v38
    }
    var v39 uint32 = v15
    _ = len(v1)
    v2 = append(v2, (v35 / (((v35 * (-(v35)))) | 1)))
    for fi40 := 0; fi40 < 13; fi40++ {
        for fi41 := 0; fi41 < 20; fi41++ {
            fmt.Println(v39)
            var v42 int32 = 42
            var v43 int32 = (v42 + int32(v36))
            _ = fi41
            _ = v42
            _ = v43
        }
        var v44 bool = (!false)
        if len(v2) > 0 {
            v2[0] = fi40
        }
        _ = fi40
        _ = v44
    }
    var v45 int32 = -808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v45
}
