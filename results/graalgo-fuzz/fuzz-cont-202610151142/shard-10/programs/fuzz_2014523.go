package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }
func genFn0() uint {
    if (1000 > -1) {
        var v6 uint16 = 255
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 float32 = float32(1.5)
            v6 = v6
            _ = fi7
            _ = v8
        }
        _ = v6
    } else {
        var v9 float32 = 0.0
        var v10 int32 = 42
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 232, F1: -80, F2: 25}
    _ = v11
    return 255
}
func genFn1(p0 uint16) string {
    v12 := GenS1{F0: 1, F1: -64, F2: 2}
    for fi13 := 0; fi13 < 15; fi13++ {
        fmt.Println(fi13)
        var v14 bool = (1 != uint8(p0))
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        p0 -= p0
        _ = fi15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        v17 := []int{fi16, 781, fi16, (^(fi16))}
        p0 = 65535
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 uint16 = (p0 | p0)
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = v12
    return ""
}

func main() {
    var v1 float64 = -2.25
    var v2 int32 = 2147483647
    v3 := []string{"foo bar", "a", "", "a"}
    v4 := make(map[int]uint32)
    var v5 uint8 = 77
    var v20 float64 = ((0.1 - v1) / -2.25)
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    if (0.0 != float32(1.5)) {
        for fi21 := 0; fi21 < 3; fi21++ {
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 string = "hello"
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    } else {
        var v24 int32 = (-(v2))
        v25 := GenS2{F0: 1}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        switch fi26 {
        case 0:
            if (65535 != 2147483648) {
                v27 := GenS3{F0: (-(v2)), F1: fi26}
                delete(v4, 32767)
                _ = len(v3)
                _ = v27
            }
        case -830:
            switch "" {
            case "\n\t":
                var v28 uint8 = 255
                _ = v28
            case "":
                v29 := GenS1{F0: v5, F1: 1, F2: v5}
                _ = v29
            case "a":
                v20 += v20
            default:
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            }
        case 128:
            fmt.Println(v5)
        default:
            var v30 uint16 = 42
            _ = v30
        }
        var v31 float64 = v20
        var v32 bool = false
        _ = fi26
        _ = v31
        _ = v32
    }
    delete(v4, 0)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    var v33 int32 = (v2 + v2)
    v34 := GenS4{F0: 0, F1: 192, F2: -829}
    v35 := GenS5{F0: 307}
    switch "" {
    case "a":
        if (uint16(v5) > 65535) {
            switch "hello" {
            case "foo bar":
                v20 = v1
            case "hello":
                if ((0.1 - (1e10 + (-2.25 + (1e10 + (-2.25 + -2.25))))) < -2.25) {
                    var v36 uint8 = 93
                    _ = v36
                }
            default:
                fmt.Println(v5, v1, v5)
            }
            for fi37 := 0; fi37 < 19; fi37++ {
                v1 = 0.1
                for fi38 := 0; fi38 < 5; fi38++ {
                    v39 := GenS6{F0: 42}
                    var v40 int16 = 32767
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                switch "a" {
                case "":
                    v1 -= (float64(1.5) / 3.14159)
                case "hello":
                    if (955 <= 9223372036854775807) {
                        v4[(fi37 + fi37)] = 448
                        var v41 string = "\n\t"
                        fmt.Println(v41)
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 17; fi42++ {
                            v2 = (v2 + (-(v33)))
                            _ = fi42
                        }
                        for fi43 := 0; fi43 < 15; fi43++ {
                            _ = len(v4)
                            _ = fi43
                        }
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                    }
                }
                _ = fi37
            }
        }
    }
    v44 := GenS4{F0: 455, F1: 42, F2: -760}
    for fi45 := 0; fi45 < 13; fi45++ {
        var v46 uint32 = 4294967295
        _ = fi45
        _ = v46
    }
    delete(v4, -533)
    v2 = v2
    var v47 int32 = 42
    if (!true) {
        v20 = (3.14159 + v1)
        var v48 int64 = -413
        _ = v48
    }
    _ = len(v3)
    if false {
        v1 += ((1.5 - 0.1) - 0.0)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v44
    _ = v47
    _ = v5
}
