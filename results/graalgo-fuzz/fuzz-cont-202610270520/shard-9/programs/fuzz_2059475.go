package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int64; F1 int8; F2 uint }
func genFn0(p0 uint8, p1 uint8) int64 {
    p0 += 127
    p1 -= p0
    switch 1 {
    case 32768:
        fmt.Println(p0, p0)
    default:
        fmt.Println(p0, p0)
    }
    switch 4294967295 {
    case 127:
        p0 = (p0 & (p1 ^ (^(p1))))
    case 2147483647:
        switch ("a" + "") {
        case "foo bar":
            var v7 uint64 = 1
            _ = v7
        case "":
            for fi8 := 0; fi8 < 16; fi8++ {
                for fi9 := 0; fi9 < 3; fi9++ {
                    v10 := []int{42, (fi8 + fi9)}
                    _ = fi9
                    _ = v10
                }
                p1 = p1
                switch "" {
                case "\n\t":
                    var v11 string = "\n\t"
                    _ = v11
                }
                _ = fi8
            }
        default:
            v12 := GenS1{F0: 768, F1: 65535}
            _ = v12
        }
    default:
        if true {
            for fi13 := 0; fi13 < 1; fi13++ {
                p1 += 42
                _ = fi13
            }
        }
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}
func genFn1(p0 int16, p1 uint64) uint32 {
    var v14 uint = 2
    var v15 uint64 = (^(p1))
    fmt.Println(v14, v15, p1)
    v16 := GenS1{F0: -2147483648, F1: 4294967295}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 96
    v3 := make(map[uint32]string)
    v4 := []string{}
    var v5 uint32 = 4294967295
    v6 := make(map[string]bool)
    var v17 int64 = 1
    var v18 uint64 = 725
    v3[genFn1((v1 * (v1 << ((4294967295) & 63))), (v18 ^ 239))] = ""
    var v19 int16 = (v1 << ((811) & 63))
    var v20 int16 = (v1 << ((255) & 63))
    if (!false) {
        if false {
            switch ("hello" + "\n\t") {
            case "hello":
                for fi21 := 0; fi21 < 13; fi21++ {
                    var v22 int16 = int16(fi21)
                    v23 := GenS2{F0: (fi21 - fi21)}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            default:
                v24 := GenS3{F0: 32767, F1: v5}
                _ = v24
            }
            v25 := v6[""]
            v6["foo bar"] = v25
            _ = v25
        }
    } else {
        var v26 int16 = 42
        v18 += (^(v18))
        _ = v26
    }
    if (float32(1e10) >= 3.14159) {
        v27 := GenS4{F0: false, F1: 128}
        _ = v27
    } else {
        switch "a" {
        case "a":
            v2 = (v2 | (v2 ^ uint16(v18)))
        case "":
            if false {
                var v28 uint64 = (v18 << ((4294967295) & 63))
                _ = v28
            }
        case "\n\t":
            v18 = 2
        }
        v29 := GenS3{F0: 284, F1: (v5 << ((1) & 63))}
        v30 := v6[("a" + "")]
        _ = v29
        _ = v30
    }
    var v31 int16 = (v19 >> ((42) & 63))
    v5 = (v5 & v5)
    switch "\n\t" {
    case "\n\t":
        v32 := v3[v5]
        _ = v32
    }
    v3[(v5 * (v5 - v5))] = "foo bar"
    switch 32767 {
    case 9223372036854775807:
        switch 753 {
        case 806:
            v33 := GenS5{F0: 0, F1: 84, F2: 4294967295}
            _ = v33
        default:
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 uint = 0
                _ = fi34
                _ = v35
            }
        }
    case 166:
        fmt.Println(v1)
    case -215:
        var v36 float64 = float64(0.1)
        _ = v36
    default:
        _ = len(v3)
    }
    for fi37 := 0; fi37 < 14; fi37++ {
        for fi38 := 0; fi38 < 20; fi38++ {
            fmt.Println(v1)
            v31 -= (-(v19))
            if false {
                var v39 int32 = 2147483647
                _ = v39
            }
            _ = fi38
        }
        v1 = (v1 >> ((1) & 63))
        fmt.Println(v19)
        _ = fi37
    }
    v5 -= v5
    v19 -= v19
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
