package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int64 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 int; F2 int64 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 bool; F1 uint; F2 uint32 }
type GenS7 struct { F0 int32; F1 int64; F2 int16 }
func genFn0(p0 uint, p1 uint) int {
    if true {
        v4 := GenS1{F0: (p0 & p1), F1: -32768, F2: -610}
        v5 := []int{42}
        var v6 uint64 = 42
        _ = v4
        _ = v5
        _ = v6
    }
    if false {
        v7 := GenS1{F0: p0, F1: -32768, F2: 982}
        _ = v7
    } else {
        switch ("\n\t" + "foo bar") {
        case "":
            var v8 int32 = 246
            _ = v8
        case "a":
            var v9 float64 = 0.0
            _ = v9
        case "foo bar":
            p1 = 255
        }
        v10 := GenS1{F0: (p0 << (((p0 >> (((p1 >> (((^(p1))) & 63))) & 63))) & 63)), F1: 0, F2: -586}
        _ = v10
    }
    _ = p0
    _ = p1
    return -2147483648
}
func genFn1(p0 uint, p1 int8) string {
    p0 += 42
    v11 := []int{}
    _ = len(v11)
    v11 = append(v11, -9223372036854775808)
    _ = p0
    _ = p1
    _ = v11
    return ("a" + "a")
}

func main() {
    var v1 bool = (-1 > -9223372036854775808)
    v2 := make(map[int]uint32)
    var v3 uint32 = 0
    switch 1 {
    case 32768:
        v12 := GenS2{F0: 1000, F1: -1}
        _ = v12
    case -2147483648:
        v13 := GenS1{F0: 65535, F1: -1, F2: 747}
        _ = v13
    case 819:
        switch genFn0(13, 1) {
        case 0:
            v14 := []int{32768, 995, -557, 9223372036854775807}
            _ = v14
        case -177:
            v15 := GenS1{F0: 0, F1: 0, F2: 42}
            _ = v15
        default:
            var v16 uint8 = uint8(v3)
            _ = v16
        }
    default:
        v17 := GenS1{F0: 65535, F1: 1000, F2: 749}
        _ = v17
    }
    v18 := GenS3{F0: 255}
    switch (("foo bar" + "hello") + "a") {
    case "foo bar":
        v19 := GenS2{F0: 343, F1: 100}
        _ = v19
    case "hello":
        v3 = (v3 >> ((1) & 63))
    default:
        switch -1 {
        case 32768:
            switch 42 {
            case 32767:
                var v20 int = -274
                _ = v20
            case 9223372036854775807:
                if (32767 != 621) {
                    fmt.Println(v1, v3)
                }
            case 42:
                var v21 int16 = -1
                _ = v21
            }
        }
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        v1 = ("" < "a")
        _ = fi22
    }
    var v23 float32 = 1.5
    var v24 float64 = 0.0
    v25 := GenS1{F0: 2, F1: 1000, F2: 0}
    var v26 float32 = v23
    switch "hello" {
    case "\n\t":
        var v27 float32 = ((-2.25 + (1e10 / v26)) / (1.5 / float32(v24)))
        _ = v27
    default:
        switch 274 {
        case 32767:
            var v28 uint16 = 32767
            _ = v28
        case 108:
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 int64 = -847
                _ = fi29
                _ = v30
            }
        }
    }
    var v31 uint8 = 42
    v31 += 51
    v32 := []int{genFn0(18446744073709551615, 776), 633, -9223372036854775808, -1}
    var v33 uint = 1
    if v1 {
        _ = len(v2)
        var v34 string = "hello"
        _ = v34
    } else {
        if (544 <= 0) {
            v35 := GenS4{F0: "hello", F1: 32767, F2: 634}
            v36 := GenS4{F0: "a", F1: -2147483648, F2: 483}
            for fi37 := 0; fi37 < 12; fi37++ {
                if v1 {
                    v32 = append(v32, fi37)
                    v2[2] = 65535
                }
                _ = fi37
            }
            _ = v35
            _ = v36
        }
        v38 := GenS1{F0: v33, F1: 0, F2: -973}
        delete(v2, -1)
        _ = v38
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        fmt.Println(v33, v26)
        switch fi39 {
        case 32767:
            var v40 uint64 = 307
            _ = v40
        case 42:
            if v1 {
                fmt.Println(v3)
                v41 := GenS2{F0: 641, F1: int8(fi39)}
                v42 := GenS5{F0: -427, F1: false}
                _ = v41
                _ = v42
            } else {
                var v43 bool = (v24 < v24)
                v44 := GenS6{F0: ((fi39 >> ((18446744073709551615) & 63)) <= -158), F1: v33, F2: v3}
                v45 := GenS6{F0: ((0.1 != (v24 - v24)) || v1), F1: (v33 & 567), F2: (^(v3))}
                _ = v43
                _ = v44
                _ = v45
            }
        default:
            var v46 int8 = 42
            _ = v46
        }
        _ = fi39
    }
    for fi47 := 0; fi47 < 13; fi47++ {
        if v1 {
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
            v31 -= uint8(v33)
        }
        v33 += uint(v31)
        v48 := GenS7{F0: 0, F1: -1, F2: 0}
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
