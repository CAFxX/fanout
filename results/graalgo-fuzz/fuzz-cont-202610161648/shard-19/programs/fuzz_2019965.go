package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 uint32; F1 int16; F2 uint32 }
func genFn0(p0 bool, p1 int64, p2 int32) string {
    p0 = (!true)
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint64 = 2
        if p0 {
            v7 -= 42
        } else {
            p0 = ((42 == 1) || p0)
            v8 := GenS1{F0: 4294967295}
            _ = v8
        }
        switch fi6 {
        case 32768:
            var v9 uint16 = 6
            _ = v9
        default:
            v10 := GenS2{F0: "a"}
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    var v11 bool = (-9223372036854775808 >= 0)
    fmt.Println(v11, p1)
    var v12 float32 = ((0.0 * 0.0) * float32(1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return ("hello" + (("" + "\n\t") + "\n\t"))
}
func genFn1(p0 uint16, p1 uint8, p2 uint) uint64 {
    p2 -= (p2 % ((18446744073709551615) | 1))
    v13 := []int{}
    if ((682 <= 2) || (1 != 42)) {
        p0 += p0
        switch "\n\t" {
        case "foo bar":
            _ = len(v13)
        default:
            var v14 int64 = int64(p1)
            _ = v14
        }
    } else {
        v15 := GenS1{F0: p2}
        _ = v15
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    switch "\n\t" {
    case "a":
        var v16 float64 = 0.0
        _ = v16
    case "":
        v17 := GenS3{F0: p2}
        _ = v17
    default:
        switch 0 {
        case 2147483647:
            var v18 int = 491
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 177
}

func main() {
    var v1 int8 = 0
    var v2 string = ""
    var v3 float32 = ((float32((1e10 * (1e10 * (float64(1e10) - float64(1.5))))) + -2.25) + 3.14159)
    var v4 int64 = 42
    v5 := make(map[string]string)
    if false {
        v19 := GenS2{F0: (v2 + (v2 + v2))}
        v1 -= (v1 << ((42) & 63))
        if false {
            v1 = 0
            switch "foo bar" {
            case "\n\t":
                if (2 != 558) {
                    if (0 == 338) {
                        v4 += (v4 - int64(v1))
                        v4 = (v4 | v4)
                        fmt.Println(v3)
                    }
                }
            default:
                v2 = ""
            }
            v20 := GenS4{F0: 42}
            _ = v20
        } else {
            var v21 uint8 = 42
            v1 += 100
            var v22 float32 = 0.0
            _ = v21
            _ = v22
        }
        _ = v19
    } else {
        fmt.Println(v2)
    }
    v23 := []int{-294}
    if false {
        var v24 int16 = -1
        v3 = v3
        var v25 uint16 = 2
        _ = v24
        _ = v25
    }
    _ = len(v23)
    if false {
        v4 += 1
        var v26 float64 = 0.1
        _ = len(v23)
        _ = v26
    }
    if false {
        v27 := GenS1{F0: 0}
        _ = v27
    } else {
        v3 = v3
    }
    if true {
        v28 := GenS5{F0: 144, F1: 18446744073709551615, F2: 1}
        _ = v28
    } else {
        v29 := GenS6{F0: 32767, F1: 1}
        v2 = v2
        fmt.Println(v4, v3, v4)
        _ = v29
    }
    v30 := GenS7{F0: 1, F1: 42, F2: 261}
    v2 = v2
    for fi31 := 0; fi31 < 17; fi31++ {
        if false {
            var v32 uint64 = 757
            var v33 string = (v2 + ("foo bar" + v2))
            _ = v32
            _ = v33
        }
        if true {
            var v34 int16 = -517
            fmt.Println(v4)
            for fi35 := 0; fi35 < 20; fi35++ {
                if len(v23) > 0 {
                    v23[0] = (fi31 >> ((0) & 63))
                }
                var v36 float64 = 1.5
                _ = fi35
                _ = v36
            }
            _ = v34
        }
        v3 = v3
        _ = fi31
    }
    v3 = -2.25
    v1 -= v1
    for fi37 := 0; fi37 < 20; fi37++ {
        fmt.Println(v1, v1)
        if (1 == 42) {
            var v38 int8 = v1
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = v38
        } else {
            var v39 uint32 = 42
            _ = v39
        }
        _ = fi37
    }
    switch v2 {
    case "foo bar":
        v3 = 3.14159
    case "hello":
        for fi40 := 0; fi40 < 15; fi40++ {
            if (!false) {
                if len(v23) > 0 {
                    v23[0] = fi40
                }
            }
            if true {
                fmt.Println(v3)
            }
            var v41 int16 = 42
            _ = fi40
            _ = v41
        }
    default:
        var v42 uint = 872
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
