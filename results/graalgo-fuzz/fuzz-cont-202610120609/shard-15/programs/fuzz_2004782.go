package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 uint; F2 int }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int, p1 int64) uint {
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: p0}
        v8 := GenS1{F0: fi6}
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(fi9)
            switch (fi9 % (((fi9 >> ((2) & 63))) | 1)) {
            case 2147483647:
                if ((p1 ^ (p1 + -1)) == p1) {
                    fmt.Println(p1, p0, fi6)
                } else {
                    if true {
                        p0 = 780
                        var v10 uint = 65535
                        if false {
                            var v11 int16 = -733
                            p1 -= p1
                            _ = v11
                        }
                        _ = v10
                    }
                    if false {
                        var v12 bool = (!true)
                        var v13 uint = 304
                        v14 := []int{(-(p0))}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    p0 -= (fi9 | fi6)
                }
            case 42:
                var v15 string = "a"
                _ = v15
            default:
                p0 += 1
            }
            var v16 int16 = -1
            _ = fi9
            _ = v16
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if ((!(1 < 42)) || true) {
        p0 = p0
    } else {
        var v17 int8 = 42
        _ = v17
    }
    v18 := GenS2{F0: true}
    _ = p0
    _ = p1
    _ = v18
    return 255
}
func genFn1(p0 bool, p1 uint8, p2 int16) uint16 {
    var v19 int16 = (p2 * int16(p1))
    switch -9223372036854775808 {
    case 9223372036854775807:
        p0 = p0
    case 1:
        v19 = (p2 - (p2 * p2))
    default:
        if false {
            p2 = p2
            v19 = v19
            switch "\n\t" {
            case "foo bar":
                fmt.Println(p0, p1)
            case "a":
                fmt.Println(v19, p1)
            case "hello":
                for fi20 := 0; fi20 < 14; fi20++ {
                    fmt.Println(v19)
                    _ = fi20
                }
            default:
                for fi21 := 0; fi21 < 4; fi21++ {
                    fmt.Println(p1)
                    fmt.Println(p1)
                    p1 += 42
                    _ = fi21
                }
            }
        }
    }
    if p0 {
        switch ("" + "") {
        case "foo bar":
            if true {
                if (!false) {
                    if p0 {
                        p2 -= v19
                        var v22 float32 = -2.25
                        _ = v22
                    }
                    var v23 float64 = (1.5 + 0.0)
                    for fi24 := 0; fi24 < 11; fi24++ {
                        p0 = p0
                        _ = fi24
                    }
                    _ = v23
                } else {
                    var v25 uint8 = p1
                    for fi26 := 0; fi26 < 4; fi26++ {
                        v27 := GenS1{F0: 0}
                        _ = fi26
                        _ = v27
                    }
                    v28 := []int{9223372036854775807, -556, 32768, 930}
                    _ = v25
                    _ = v28
                }
                v29 := []int{128, -1, 4294967295, 143}
                _ = v29
            } else {
                p2 = p2
                if true {
                    if true {
                        if false {
                            var v30 int = 9223372036854775807
                            for fi31 := 0; fi31 < 10; fi31++ {
                                for fi32 := 0; fi32 < 3; fi32++ {
                                    v33 := []int{fi31}
                                    p2 += (v19 * -32768)
                                    _ = fi32
                                    _ = v33
                                }
                                _ = fi31
                            }
                            _ = v30
                        }
                    } else {
                        v34 := GenS1{F0: 4294967295}
                        _ = v34
                    }
                    var v35 int = 2147483647
                    _ = v35
                }
                var v36 int16 = (p2 % ((42) | 1))
                _ = v36
            }
        }
        v19 += (v19 & 1)
        if true {
            v37 := GenS3{F0: 0}
            var v38 bool = (2147483647 <= 42)
            v39 := GenS2{F0: p0}
            _ = v37
            _ = v38
            _ = v39
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 65535
}

func main() {
    var v1 uint32 = 1
    var v2 float32 = 1e10
    var v3 uint16 = 2
    v4 := []uint32{42, (v1 % (((v1 ^ uint32(v3))) | 1)), (v1 / (((v1 << ((uint(v1)) & 63))) | 1)), (v1 | v1)}
    var v5 uint64 = 1
    var v40 float64 = -2.25
    v40 = (v40 * float64((((v2 / 3.14159) * (v2 + 3.14159)) * (((v2 - 0.1) - 1.5) / float32(v40)))))
    v41 := GenS1{F0: -9223372036854775808}
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 uint32 = v1
        var v44 uint64 = (v5 ^ 0)
        _ = fi42
        _ = v43
        _ = v44
    }
    if (!true) {
        switch 127 {
        case 953:
            _ = len(v4)
        case 0:
            v3 = 2
        }
        switch "hello" {
        case "hello":
            fmt.Println(v40, v5, v1)
        case "foo bar":
            fmt.Println(v5, v5)
        default:
            var v45 float64 = v40
            _ = v45
        }
    }
    v1 -= v1
    v46 := GenS4{F0: false, F1: 87, F2: 582}
    v3 = v3
    v4 = append(v4, 2147483648)
    for fi47 := 0; fi47 < 9; fi47++ {
        if (0 <= -128) {
            v3 = 65535
            var v48 int16 = -32768
            v49 := GenS5{F0: 32767, F1: 1}
            _ = v48
            _ = v49
        } else {
            v1 = (v1 % (((v1 * v1)) | 1))
        }
        var v50 float64 = v40
        _ = len(v4)
        _ = fi47
        _ = v50
    }
    for fi51 := 0; fi51 < 14; fi51++ {
        if true {
            var v52 uint32 = 973
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v52
        }
        v2 = (v2 / (v2 - ((3.14159 - (-2.25 / v2)) / ((0.0 / v2) / 1.0))))
        fmt.Println(v40, v2, fi51)
        _ = fi51
    }
    v40 = v40
    var v53 int = 127
    var v54 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v46
    _ = v5
    _ = v53
    _ = v54
}
