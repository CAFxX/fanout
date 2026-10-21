package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint64; F1 int32; F2 int32 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int; F1 uint64; F2 uint8 }
type GenS8 struct { F0 bool; F1 uint; F2 int64 }
func genFn0(p0 uint16, p1 bool) int64 {
    var v6 bool = p1
    var v7 int16 = -623
    for fi8 := 0; fi8 < 3; fi8++ {
        if false {
            if p1 {
                v9 := []int{127}
                switch ("foo bar" + "a") {
                case "hello":
                    var v10 uint = 1
                    _ = v10
                case "foo bar":
                    _ = len(v9)
                }
                _ = v9
            } else {
                if (0 > 255) {
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 uint16 = 968
                        switch -407 {
                        case 128:
                            v13 := GenS1{F0: (!true)}
                            _ = v13
                        case 1:
                            var v14 float64 = (1e10 / 3.14159)
                            _ = v14
                        default:
                            var v15 uint32 = 42
                            _ = v15
                        }
                        v16 := []int{fi8, 502, (-(fi8))}
                        var v17 int16 = (v7 << ((0) & 63))
                        if true {
                            _ = len(v16)
                            v18 := GenS2{F0: (^(v7))}
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                            _ = v18
                        }
                        switch 42 {
                        case 2147483647:
                            var v19 int64 = 0
                            _ = v19
                        case -1:
                            if ((fi8 <= (fi11 / ((fi8) | 1))) || (32767 > (fi8 >> ((425) & 63)))) {
                                v16 = append(v16, 9223372036854775807)
                                if ((fi11 == 90) && (9223372036854775807 >= 9223372036854775807)) {
                                    switch "\n\t" {
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "":
                                            v7 = v7
                                        case "hello":
                                            var v20 uint64 = 847
                                            _ = v20
                                        case "foo bar":
                                            var v21 bool = true
                                            _ = v21
                                        default:
                                            var v22 int = (^(fi8))
                                            _ = v22
                                        }
                                    default:
                                        v23 := GenS3{F0: 268, F1: 1}
                                        _ = v23
                                    }
                                    var v24 string = "hello"
                                    _ = v24
                                }
                                v25 := GenS4{F0: 1, F1: -2147483648, F2: 0}
                                var v26 int16 = 32767
                                _ = v25
                                _ = v26
                            }
                        case 32767:
                            v16 = append(v16, fi8)
                        }
                        _ = fi11
                        _ = v12
                        _ = v16
                        _ = v17
                    }
                }
                var v27 int16 = int16(fi8)
                _ = v27
            }
            v6 = false
            v28 := GenS5{F0: 0, F1: 2147483647}
            _ = v28
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint8) int64 {
    v29 := GenS6{F0: 2, F1: 4294967295}
    p0 = p0
    p0 -= 168
    var v30 uint = 42
    _ = p0
    _ = v29
    _ = v30
    return -1
}

func main() {
    var v1 bool = (42 < -755)
    var v2 int8 = -121
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    v5 := make(map[string]int)
    for fi31 := 0; fi31 < 4; fi31++ {
        v32 := GenS3{F0: 753, F1: 18446744073709551615}
        _ = fi31
        _ = v32
    }
    if v1 {
        v2 -= 15
        fmt.Println(v2, v2)
        for fi33 := 0; fi33 < 11; fi33++ {
            var v34 string = "foo bar"
            v2 += (v2 % ((v2) | 1))
            _ = fi33
            _ = v34
        }
    } else {
        var v35 uint8 = 1
        v1 = (2 != 719)
        var v36 string = "\n\t"
        _ = v35
        _ = v36
    }
    v1 = (("hello" + ("\n\t" + "a")) == "foo bar")
    v37 := GenS7{F0: 9223372036854775807, F1: 0, F2: 127}
    var v38 int16 = 1
    v38 += -438
    var v39 uint = 65535
    fmt.Println(v39)
    var v40 uint = (v39 % ((1) | 1))
    var v41 int64 = -173
    var v42 uint = 255
    switch -9223372036854775808 {
    case -9223372036854775808:
        v41 = -1
    case -392:
        v43 := GenS8{F0: true, F1: v42, F2: (^(v41))}
        _ = v43
    case 9223372036854775807:
        var v44 int64 = (v41 % ((264) | 1))
        _ = v44
    }
    v45 := GenS1{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v5
}
