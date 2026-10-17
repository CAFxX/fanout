package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int64; F2 int }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint64 }
func genFn0(p0 uint32, p1 uint32, p2 int32) int32 {
    if false {
        switch 127 {
        case 32767:
            p1 -= 4294967295
        case 368:
            if false {
                var v6 float32 = 0.0
                v7 := []int{42, 4294967295, -2147483648, 33}
                for fi8 := 0; fi8 < 3; fi8++ {
                    v9 := GenS1{F0: 1000}
                    if ((p0 % ((p1) | 1)) >= (p1 ^ 389)) {
                        v10 := GenS1{F0: 0}
                        fmt.Println(fi8)
                        var v11 int32 = (^(p2))
                        _ = v10
                        _ = v11
                    } else {
                        fmt.Println(p1, p1, p2)
                        var v12 uint = 4294967295
                        if true {
                            var v13 uint16 = 210
                            for fi14 := 0; fi14 < 20; fi14++ {
                                v15 := GenS2{F0: 100, F1: 9223372036854775807, F2: fi8}
                                _ = fi14
                                _ = v15
                            }
                            _ = v13
                        }
                        _ = v12
                    }
                    _ = fi8
                    _ = v9
                }
                _ = v6
                _ = v7
            }
        default:
            fmt.Println(p2, p2, p1)
        }
        var v16 string = ""
        var v17 int16 = 14
        _ = v16
        _ = v17
    }
    v18 := []int{-9223372036854775808, 9223372036854775807, 879, 32768}
    v19 := GenS3{F0: (p1 << ((255) & 63)), F1: 865}
    var v20 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    return 270
}
func genFn1() string {
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 int32 = 2147483647
        _ = fi21
        _ = v22
    }
    var v23 uint32 = 0
    var v24 uint = 255
    var v25 int16 = 694
    _ = v23
    _ = v24
    _ = v25
    return "a"
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 int = 223
    var v4 int64 = 1
    var v5 float32 = (-2.25 + (0.1 * float32(float64((float32((1e10 * ((((((0.1 + 3.14159) * (0.1 + 0.0)) - 0.1) / 1.0) + 1e10) * 1.5))) * 0.0)))))
    v26 := GenS2{F0: 0, F1: 9223372036854775807, F2: v3}
    v3 -= 4294967295
    if true {
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 bool = (!true)
            fmt.Println(v5, v3)
            switch 32768 {
            case 2147483647:
                var v29 uint16 = 2
                _ = v29
            case 250:
                v30 := GenS4{F0: 772}
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        switch 726 {
        case 0:
            var v31 int8 = -128
            _ = v31
        case -9223372036854775808:
            var v32 float32 = float32(-2.25)
            _ = v32
        }
    }
    var v33 uint32 = 1
    for fi34 := 0; fi34 < 1; fi34++ {
        v35 := GenS5{F0: true}
        var v36 uint8 = 74
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 float32 = (v5 + v5)
    if (!(-1 < 2147483647)) {
        var v38 float64 = 3.14159
        v39 := []int{(v3 | (v3 >> ((1) & 63))), int(v4), -774}
        _ = v38
        _ = v39
    } else {
        switch "hello" {
        case "\n\t":
            switch "a" {
            case "foo bar":
                if ((0 <= 128) && (!false)) {
                    var v40 int = 127
                    _ = v40
                } else {
                    v33 = (v33 + v33)
                }
            default:
                v41 := GenS4{F0: v33}
                _ = v41
            }
        case "foo bar":
            v5 = v37
        }
        for fi42 := 0; fi42 < 4; fi42++ {
            v43 := GenS1{F0: int16(v33)}
            var v44 int64 = v4
            _ = fi42
            _ = v43
            _ = v44
        }
        v45 := GenS6{F0: 0}
        _ = v45
    }
    v46 := []int{}
    switch v3 {
    case 0:
        _ = len(v46)
    case 564:
        switch 2147483647 {
        case 0:
            for fi47 := 0; fi47 < 16; fi47++ {
                switch "\n\t" {
                case "a":
                    v48 := GenS3{F0: 2, F1: 2147483647}
                    _ = v48
                }
                _ = fi47
            }
        default:
            if true {
                var v49 uint64 = 18446744073709551615
                v5 += (v5 * v5)
                if len(v46) >= 2 {
                    v46 = v46[1:len(v46)]
                }
                _ = v49
            } else {
                v46 = append(v46, 9223372036854775807)
            }
        }
    case 9223372036854775807:
        v3 = 409
    default:
        v4 -= -9223372036854775808
    }
    v50 := GenS7{F0: 42, F1: 4294967295, F2: 0}
    v5 = (v5 + v37)
    fmt.Println(v5)
    v3 = (-(v3))
    v5 += v5
    for fi51 := 0; fi51 < 18; fi51++ {
        fmt.Println(v4)
        var v52 int32 = 0
        _ = fi51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v46
    _ = v5
    _ = v50
}
