package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 int16; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint64; F1 int64; F2 bool }
type GenS8 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0(p0 uint64) uint8 {
    var v4 float32 = ((((3.14159 / (-2.25 * float32(-2.25))) * (0.0 * 3.14159)) + (float32(((-2.25 + -2.25) * ((-2.25 / 1.5) + 1e10))) - 1e10)) / -2.25)
    fmt.Println(v4)
    _ = p0
    _ = v4
    return 1
}
func genFn1(p0 uint16) uint {
    if true {
        v5 := GenS1{F0: "foo bar", F1: 717, F2: 728}
        _ = v5
    } else {
        switch 4294967295 {
        case 0:
            p0 = p0
        case 2147483647:
            fmt.Println(p0)
        }
        var v6 bool = (false || (!true))
        _ = v6
    }
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 10; fi8++ {
        if (true || false) {
            v7 = v7
            p0 += (p0 - (p0 >> ((1) & 63)))
            switch fi8 {
            case -1:
                fmt.Println(v7, fi8)
            case -2147483648:
                v9 := GenS1{F0: ("foo bar" + v7), F1: 0, F2: (fi8 << ((2) & 63))}
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 11; fi10++ {
                fmt.Println(p0)
                v7 = v7
                _ = fi10
            }
            switch fi8 {
            case -429:
                var v11 int = (fi8 >> ((2) & 63))
                _ = v11
            case -994:
                for fi12 := 0; fi12 < 9; fi12++ {
                    p0 = p0
                    _ = fi12
                }
            default:
                var v13 bool = (!false)
                _ = v13
            }
        }
        p0 = p0
        v14 := []int{(fi8 >> ((42) & 63)), (fi8 % ((32767) | 1)), 1}
        _ = fi8
        _ = v14
    }
    if (1 >= 333) {
        fmt.Println(p0, p0)
        fmt.Println(v7, p0)
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    v1 := []int{676, 32767, 311, 522}
    var v2 uint32 = 2147483648
    var v3 int16 = -745
    v15 := GenS2{F0: (v2 < v2), F1: -2147483648}
    v16 := GenS3{F0: 41, F1: 255, F2: -2147483648}
    var v17 int8 = 127
    switch ("foo bar" + "a") {
    case "a":
        v2 += (v2 ^ 0)
    case "\n\t":
        v18 := GenS1{F0: "\n\t", F1: -9223372036854775808, F2: -569}
        _ = v18
    case "hello":
        switch ("" + ("a" + "a")) {
        case "a":
            if true {
                v3 -= (v3 & v3)
                v19 := GenS4{F0: (^(v17)), F1: "\n\t"}
                fmt.Println(v3)
                _ = v19
            } else {
                switch 30 {
                case -100:
                    var v20 float32 = -2.25
                    _ = v20
                }
            }
        case "\n\t":
            for fi21 := 0; fi21 < 11; fi21++ {
                for fi22 := 0; fi22 < 10; fi22++ {
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        fmt.Println(fi22, v3, v3)
                    }
                    if (1 <= -312) {
                        if (-2.25 < 1e10) {
                            v23 := GenS5{F0: (true || ((977 <= 32767) || (false && (1e10 >= ((1.5 + ((1.5 / 1.0) * (1e10 * 0.1))) + (-2.25 - ((float32(((0.0 + 3.14159) + (3.14159 / 1.0))) / -2.25) * 3.14159))))))), F1: 0}
                            var v24 uint8 = 247
                            if false {
                                if true {
                                    v17 = v17
                                } else {
                                    var v25 int = fi21
                                    var v26 uint8 = (v24 & (^(v24)))
                                    _ = v25
                                    _ = v26
                                }
                                var v27 int16 = int16(v2)
                                v28 := GenS6{F0: 851, F1: v17, F2: (v2 * (v2 / ((v2) | 1)))}
                                _ = v27
                                _ = v28
                            }
                            _ = v23
                            _ = v24
                        }
                    }
                    _ = len(v1)
                    _ = fi22
                }
                if false {
                    v1 = append(v1, 222)
                    v3 = (^(v3))
                    for fi29 := 0; fi29 < 14; fi29++ {
                        var v30 bool = (v3 <= (-(v3)))
                        _ = fi29
                        _ = v30
                    }
                }
                v3 = v3
                _ = fi21
            }
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        v32 := GenS1{F0: "foo bar", F1: 0, F2: fi31}
        _ = fi31
        _ = v32
    }
    switch 446 {
    case -432:
        var v33 int8 = (v17 | 42)
        _ = v33
    case 32767:
        v17 = int8(v2)
    case -9223372036854775808:
        v2 -= 2147483647
    }
    v3 = v3
    switch "foo bar" {
    case "foo bar":
        for fi34 := 0; fi34 < 15; fi34++ {
            v2 = 65535
            for fi35 := 0; fi35 < 18; fi35++ {
                var v36 int16 = v3
                if len(v1) > 0 {
                    v1[0] = fi35
                }
                v37 := GenS2{F0: (!(v36 >= 978)), F1: -1}
                v36 = (-(v36))
                if false {
                    var v38 float32 = float32(3.14159)
                    _ = v38
                }
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = len(v1)
            _ = fi34
        }
    case "\n\t":
        v39 := GenS7{F0: 974, F1: -840, F2: false}
        _ = v39
    case "a":
        var v40 int64 = -1
        _ = v40
    default:
        var v41 int8 = (v17 | (v17 >> ((42) & 63)))
        _ = v41
    }
    var v42 int32 = 1
    v43 := GenS8{F0: 552, F1: 1, F2: (v3 + (v3 * v3))}
    v17 -= (v17 * (v17 | v17))
    fmt.Println(v42)
    if false {
        fmt.Println(v3, v3, v2)
        var v44 uint64 = 4294967295
        _ = v44
    } else {
        fmt.Println(v42, v2)
    }
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
