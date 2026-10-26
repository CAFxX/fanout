package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint32; F2 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS7 struct { F0 uint32; F1 uint64; F2 int64 }
func genFn0() uint16 {
    var v4 bool = true
    var v5 uint8 = 1
    v6 := GenS1{F0: ((("" + "\n\t") + ("a" + (("foo bar" + "foo bar") + "a"))) + "foo bar"), F1: 255, F2: v4}
    if (false && false) {
        v7 := GenS2{F0: 4294967295}
        if ((float32(1.5) - ((1.5 - (float32(1e10) / 1.0)) + float32((0.1 * 1.5)))) == 0.1) {
            switch -351 {
            case 2147483647:
                v4 = (32767 < 174)
            case -27:
                v4 = v4
            case 32768:
                fmt.Println(v4, v4)
            }
            fmt.Println(v5)
            v5 += v5
        }
        v4 = (!(uint(v5) != 1))
        _ = v7
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(v5, v5, v5)
            v9 := GenS1{F0: "hello", F1: 2, F2: v4}
            _ = fi8
            _ = v9
        }
    }
    var v10 uint32 = 373
    _ = v10
    _ = v4
    _ = v5
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -19
    v3 := []uint32{}
    v2 = v2
    v2 -= 100
    switch "hello" {
    case "hello":
        v11 := GenS1{F0: "foo bar", F1: 255, F2: true}
        _ = v11
    case "":
        var v12 uint8 = 139
        _ = v12
    case "a":
        var v13 uint64 = 480
        _ = v13
    default:
        v14 := GenS3{F0: 255, F1: 142, F2: 9223372036854775807}
        _ = v14
    }
    fmt.Println(v2)
    var v15 string = "\n\t"
    v16 := GenS4{F0: v15}
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := GenS5{F0: v2, F1: (fi18 & 42), F2: 0}
            _ = fi18
            _ = v19
        }
        var v20 float64 = 3.14159
        var v21 int64 = -9223372036854775808
        _ = fi17
        _ = v20
        _ = v21
    }
    switch 9223372036854775807 {
    case 42:
        v22 := GenS4{F0: v15}
        _ = v22
    case -144:
        v3 = append(v3, 1)
    }
    switch ((v15 + "") + (v15 + v15)) {
    case "hello":
        var v23 uint8 = 42
        _ = v23
    case "\n\t":
        if false {
            v1[(v15 + (((v15 + "foo bar") + v15) + v15))] = 32767
            v2 = v2
        }
    default:
        v2 = v2
    }
    switch -2147483648 {
    case 913:
        if true {
            v24 := GenS3{F0: 18446744073709551615, F1: 0, F2: 9223372036854775807}
            switch v15 {
            case "\n\t":
                fmt.Println(v2)
            case "foo bar":
                v2 -= -1
            }
            _ = v24
        } else {
            var v25 uint64 = 2
            _ = v25
        }
    case 42:
        for fi26 := 0; fi26 < 18; fi26++ {
            fmt.Println(v15, fi26)
            _ = fi26
        }
    case 734:
        switch "\n\t" {
        case "hello":
            switch "foo bar" {
            case "hello":
                if (!(2147483647 < -1)) {
                    if true {
                        var v27 uint32 = 2147483648
                        var v28 uint32 = (v27 | v27)
                        _ = v27
                        _ = v28
                    } else {
                        fmt.Println(v15, v15)
                    }
                    v2 = v2
                    switch 886 {
                    case 2147483647:
                        v2 += (v2 ^ (v2 * v2))
                    }
                }
            case "\n\t":
                var v29 uint8 = 70
                _ = v29
            default:
                v1[(v15 + "hello")] = 32767
            }
        }
    }
    v30 := GenS4{F0: (("" + v15) + "\n\t")}
    v15 = ((("" + "foo bar") + "a") + (v15 + (v15 + ("\n\t" + "a"))))
    v3 = append(v3, 775)
    if (255 >= 18446744073709551615) {
        for fi31 := 0; fi31 < 8; fi31++ {
            v32 := GenS5{F0: (v2 | v2), F1: (fi31 / ((fi31) | 1)), F2: 206}
            _ = fi31
            _ = v32
        }
        var v33 uint16 = 2
        for fi34 := 0; fi34 < 10; fi34++ {
            fmt.Println(v15, v2)
            for fi35 := 0; fi35 < 12; fi35++ {
                var v36 uint32 = 604
                var v37 uint64 = 1
                if true {
                    v15 = ("\n\t" + v15)
                } else {
                    v3 = append(v3, 1)
                }
                _ = fi35
                _ = v36
                _ = v37
            }
            var v38 int16 = 32767
            _ = fi34
            _ = v38
        }
        _ = v33
    }
    v39 := GenS6{F0: 121, F1: 65535, F2: 1}
    delete(v1, "hello")
    v40 := GenS7{F0: 2147483648, F1: 142, F2: 826}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v39
    _ = v40
}
