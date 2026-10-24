package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -1
    var v3 int = 32768
    v4 := make(map[string]int)
    v5 := make(map[string]bool)
    var v6 float32 = 3.14159
    v6 = (((v6 / v6) * float32(0.1)) - -2.25)
    fmt.Println(v3)
    v6 = 1e10
    if true {
        v2 -= v2
        if (888 <= 290) {
            fmt.Println(v2, v2, v2)
            var v7 uint32 = 2147483648
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 uint8 = 0
                switch "" {
                case "a":
                    var v10 uint16 = 0
                    _ = v10
                case "hello":
                    for fi11 := 0; fi11 < 6; fi11++ {
                        v6 += float32(0.1)
                        if false {
                            v12 := GenS1{F0: 0, F1: (((18446744073709551615 < 144) || false) || true), F2: v7}
                            _ = v12
                        }
                        _ = fi11
                    }
                }
                var v13 float32 = v6
                _ = fi8
                _ = v13
                _ = v9
            }
            _ = v7
        } else {
            v14 := GenS2{F0: 4294967295}
            delete(v1, 2)
            for fi15 := 0; fi15 < 17; fi15++ {
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := []int{(fi16 | fi15)}
                    _ = fi16
                    _ = v17
                }
                v18 := v4["\n\t"]
                _ = fi15
                _ = v18
            }
            _ = v14
        }
    } else {
        v3 = 0
        var v19 uint64 = 4294967295
        _ = v19
    }
    v2 -= (-(v2))
    delete(v5, "hello")
    v20 := []int{v3, (v3 >> ((836) & 63)), v3, (v3 - (v3 ^ v3))}
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint32 = 2147483648
        var v23 uint32 = 1
        _ = fi21
        _ = v22
        _ = v23
    }
    if (65535 > 774) {
        switch "a" {
        case "\n\t":
            v2 = (^(v2))
        case "a":
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v6, v3, v3)
            }
        case "":
            v2 += v2
        default:
            var v24 string = ""
            _ = v24
        }
        switch v3 {
        case -44:
            _ = len(v20)
        default:
            if true {
                var v25 int32 = -643
                v6 += (((v6 - v6) / v6) + (v6 + 0.0))
                _ = v25
            } else {
                var v26 uint64 = 42
                var v27 int64 = v2
                for fi28 := 0; fi28 < 18; fi28++ {
                    if (230 >= 127) {
                        if false {
                            var v29 uint = 1
                            var v30 uint64 = (^(v26))
                            _ = v29
                            _ = v30
                        }
                    } else {
                        v26 = v26
                        v3 = 0
                    }
                    _ = fi28
                }
                _ = v26
                _ = v27
            }
        }
        v3 = 722
    }
    var v31 bool = (!(42 > -518))
    var v32 uint16 = 344
    var v33 float32 = 3.14159
    if v31 {
        if v31 {
            var v34 uint64 = 175
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = v34
        }
        for fi35 := 0; fi35 < 18; fi35++ {
            for fi36 := 0; fi36 < 3; fi36++ {
                fmt.Println(fi36, v6, v31)
                v6 -= v6
                v37 := GenS3{F0: v2}
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        fmt.Println(v33)
    }
    var v38 uint8 = 155
    switch v3 {
    case 32768:
        var v39 uint8 = (^(v38))
        _ = v39
    case 0:
        var v40 uint = 4294967295
        _ = v40
    default:
        v41 := GenS4{F0: -128, F1: -1}
        _ = v41
    }
    v32 -= (v32 ^ (v32 % ((2) | 1)))
    v2 = (v2 % ((v2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
