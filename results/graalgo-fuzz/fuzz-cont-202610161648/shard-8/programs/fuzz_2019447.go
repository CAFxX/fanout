package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint16; F2 uint }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint64) int16 {
    v7 := GenS1{F0: (^(p0)), F1: 2147483647}
    if (-32768 > 0) {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint32 = 872
            if false {
                v10 := GenS1{F0: (p0 ^ (p0 + p0)), F1: fi8}
                _ = v10
            }
            v11 := GenS1{F0: p0, F1: -1}
            _ = fi8
            _ = v11
            _ = v9
        }
    } else {
        var v12 string = "a"
        switch v12 {
        case "a":
            v12 = (v12 + v12)
        }
        if true {
            p0 -= p0
            p0 -= 1
        } else {
            v13 := []int{-9223372036854775808, 32767}
            _ = v13
        }
        _ = v12
    }
    v14 := GenS2{F0: -9223372036854775808, F1: p0, F2: 0}
    switch 0 {
    case 42:
        if (985 <= 255) {
            v15 := []int{2147483647, 2147483647}
            _ = v15
        }
    default:
        var v16 float32 = 1e10
        _ = v16
    }
    switch "" {
    case "a":
        p0 -= (p0 * 503)
    case "":
        v17 := []int{42, 127, 127, 32768}
        _ = v17
    case "\n\t":
        var v18 uint32 = 0
        _ = v18
    }
    _ = p0
    _ = v14
    _ = v7
    return 32767
}
func genFn1(p0 int16) uint32 {
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 16; fi20++ {
            fmt.Println(fi20, fi20, p0)
            var v21 string = "foo bar"
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v22 := GenS3{F0: 1}
    fmt.Println(p0)
    v23 := GenS1{F0: 42, F1: 127}
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 int32 = 2147483647
        for fi26 := 0; fi26 < 8; fi26++ {
            p0 -= p0
            v27 := []int{0}
            _ = len(v27)
            var v28 int32 = v25
            fmt.Println(fi26, fi24, fi26)
            var v29 uint16 = 65535
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi24
        _ = v25
    }
    _ = p0
    _ = v22
    _ = v23
    return 42
}

func main() {
    var v1 uint64 = 217
    var v2 string = ""
    v3 := []uint32{707, 2147483648, 2147483648, 2}
    v4 := make(map[uint32]string)
    var v5 int16 = 1000
    var v6 int16 = v5
    v1 = v1
    fmt.Println(v2, v6, v1)
    if (2147483648 >= 2147483648) {
        v6 = 826
    }
    var v30 float32 = ((-2.25 / 0.1) - 0.1)
    for fi31 := 0; fi31 < 5; fi31++ {
        for fi32 := 0; fi32 < 14; fi32++ {
            if true {
                _ = len(v4)
                v4[817] = v2
                v2 = v2
            } else {
                var v33 int32 = -2147483648
                switch fi31 {
                case 317:
                    v6 = (v5 & (v5 % ((v5) | 1)))
                default:
                    switch "\n\t" {
                    case "foo bar":
                        var v34 string = v2
                        _ = v34
                    default:
                        fmt.Println(fi32, fi32, v33)
                    }
                }
                v2 = "hello"
                var v35 int64 = -1
                fmt.Println(v2, v6, v33)
                _ = v33
                _ = v35
            }
            if (1 <= 255) {
                var v36 int16 = (v5 | v6)
                var v37 bool = true
                _ = v36
                _ = v37
            } else {
                if false {
                    _ = len(v4)
                    var v38 uint32 = 65535
                    _ = v38
                } else {
                    var v39 int = (fi32 | (fi31 / ((int(v1)) | 1)))
                    if (v2 >= v2) {
                        v40 := GenS2{F0: -483, F1: 1, F2: v1}
                        _ = v40
                    }
                    var v41 uint = 541
                    v42 := GenS4{F0: int64(v41)}
                    _ = v39
                    _ = v41
                    _ = v42
                }
            }
            if false {
                if false {
                    v43 := GenS1{F0: (v1 & (v1 << ((0) & 63))), F1: (-(fi32))}
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v43
                }
            } else {
                v1 -= v1
                var v44 float64 = 1.5
                _ = v44
            }
            _ = fi32
        }
        v30 -= float32(0.1)
        _ = fi31
    }
    for fi45 := 0; fi45 < 10; fi45++ {
        for fi46 := 0; fi46 < 13; fi46++ {
            v5 += (v6 / (((v6 / ((v5) | 1))) | 1))
            v47 := GenS5{F0: 0, F1: 0, F2: 255}
            v30 = 1.5
            _ = fi46
            _ = v47
        }
        v48 := GenS4{F0: 42}
        v2 = "a"
        _ = fi45
        _ = v48
    }
    fmt.Println(v2, v1)
    switch v2 {
    case "\n\t":
        v4[2] = ("hello" + "a")
    default:
        var v49 int = 32768
        _ = v49
    }
    if true {
        var v50 int16 = 1000
        for fi51 := 0; fi51 < 19; fi51++ {
            if false {
                fmt.Println(v50, v50, v30)
            }
            _ = len(v4)
            v52 := v4[975]
            _ = fi51
            _ = v52
        }
        _ = v50
    } else {
        v5 -= (-(v5))
        v53 := GenS6{F0: 0}
        var v54 int8 = 0
        _ = v53
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
