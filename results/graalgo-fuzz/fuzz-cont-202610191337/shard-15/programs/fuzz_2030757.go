package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 string; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int64, p1 int16) int8 {
    v6 := GenS1{F0: 609, F1: 0}
    for fi7 := 0; fi7 < 20; fi7++ {
        p0 += p0
        fmt.Println(p0)
        _ = fi7
    }
    p0 -= p0
    var v8 uint8 = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 0
}
func genFn1(p0 uint64, p1 int16, p2 uint) uint32 {
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := []int{}
            _ = len(v11)
            _ = fi10
            _ = v11
        }
        v12 := []int{fi9, (fi9 >> ((p2) & 63))}
        _ = fi9
        _ = v12
    }
    v13 := []int{630, 413, -938}
    var v14 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[string]uint32)
    v4 := []int{-2147483648}
    var v5 int8 = -65
    fmt.Println(v5)
    if (995 < 796) {
        for fi15 := 0; fi15 < 4; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := GenS1{F0: 610, F1: int64(v5)}
                _ = len(v3)
                v18 := GenS2{F0: 0}
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int32 = 1
            _ = fi15
            _ = v19
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        fmt.Println(v5)
        fmt.Println(v5)
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 int16 = 32767
                _ = fi21
                _ = v22
            }
            var v23 uint = 42
            _ = fi20
            _ = v23
        }
    }
    v5 = (v5 + 1)
    var v24 float64 = 3.14159
    var v25 uint64 = 2
    var v26 int16 = -795
    var v27 uint8 = 14
    var v28 float64 = (float64((((float32(v24) + 3.14159) / (3.14159 + 0.0)) * (float32(0.0) + (1e10 / 3.14159)))) * 3.14159)
    v5 -= (v5 % ((v5) | 1))
    fmt.Println(v5)
    var v29 uint = 42
    switch "a" {
    case "":
        var v30 float32 = (3.14159 - (float32((float64(3.14159) * (v24 + (v24 * 1.5)))) + 1e10))
        _ = v30
    case "foo bar":
        var v31 uint32 = 881
        _ = v31
    }
    if false {
        v5 = 127
    } else {
        v3["\n\t"] = 65535
        for fi32 := 0; fi32 < 20; fi32++ {
            if len(v4) > 0 {
                v4[0] = (fi32 | fi32)
            }
            var v33 int8 = v5
            _ = fi32
            _ = v33
        }
    }
    v34 := GenS3{F0: (("hello" + "a") + ("hello" + ("" + "\n\t"))), F1: 796}
    if true {
        switch 32767 {
        case 127:
            var v35 float64 = (v28 - v24)
            _ = v35
        default:
            v29 = v29
        }
        if (!(-2.25 != 1e10)) {
            v24 += -2.25
        } else {
            _ = len(v4)
            for fi36 := 0; fi36 < 1; fi36++ {
                if (!false) {
                    v37 := GenS4{F0: ("" + ("hello" + ("\n\t" + "foo bar"))), F1: 127, F2: 2}
                    v38 := GenS2{F0: 286}
                    v5 -= (-(v5))
                    _ = v37
                    _ = v38
                } else {
                    v39 := GenS1{F0: (v25 / (((^(v25))) | 1)), F1: 1}
                    _ = v39
                }
                _ = fi36
            }
        }
    } else {
        fmt.Println(v27, v29, v24)
        v40 := GenS5{F0: v27}
        _ = v40
    }
    v41 := GenS5{F0: v27}
    switch "hello" {
    case "a":
        var v42 float64 = float64(1.5)
        _ = v42
    case "":
        v2["\n\t"] = -1
    default:
        switch "" {
        case "hello":
            v43 := GenS4{F0: "a", F1: v27, F2: 65535}
            _ = v43
        case "":
            for fi44 := 0; fi44 < 15; fi44++ {
                v5 -= 127
                _ = fi44
            }
        default:
            switch "foo bar" {
            case "":
                _ = len(v4)
            }
        }
    }
    var v45 uint = uint(v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v45
    _ = v5
}
