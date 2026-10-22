package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 string; F1 int }
func genFn0(p0 int64) string {
    var v5 int = -9223372036854775808
    for fi6 := 0; fi6 < 1; fi6++ {
        v5 -= 842
        if false {
            p0 -= p0
            fmt.Println(fi6)
            var v7 bool = (42 >= -32768)
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return ""
}
func genFn1(p0 uint16, p1 uint8) int16 {
    v8 := GenS1{F0: 4294967295}
    if false {
        v9 := []int{-9223372036854775808, 9223372036854775807, 32768}
        var v10 uint8 = (p1 << ((1) & 63))
        p1 += 255
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 846}
    if ((p1 / ((p1) | 1)) == p1) {
        p1 = (p1 >> ((409) & 63))
        var v12 uint = 1
        _ = v12
    }
    switch 697 {
    case 397:
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint32 = 81
            _ = fi13
            _ = v14
        }
    case 4294967295:
        v15 := GenS2{F0: 18446744073709551615}
        _ = v15
    default:
        v16 := GenS3{F0: -694, F1: 850}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return 1000
}

func main() {
    var v1 float64 = (0.1 + 1e10)
    var v2 uint64 = 457
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]string)
    v17 := GenS1{F0: v2}
    v18 := GenS2{F0: 101}
    var v19 int32 = -2147483648
    if (-664 >= 1) {
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 20; fi21++ {
                switch (fi21 + (fi20 * fi21)) {
                case -788:
                    fmt.Println(fi20, fi20, v2)
                default:
                    v22 := []int{fi20, fi21}
                    _ = v22
                }
                var v23 int16 = int16(fi21)
                _ = fi21
                _ = v23
            }
            v2 = (v2 << ((65535) & 63))
            var v24 bool = (1e10 < 1.5)
            _ = fi20
            _ = v24
        }
    } else {
        if false {
            var v25 int8 = 127
            if (224 > 42) {
                v4[0] = ("\n\t" + "\n\t")
                v1 -= v1
            } else {
                if true {
                    var v26 uint16 = 257
                    v2 += uint64(v26)
                    for fi27 := 0; fi27 < 5; fi27++ {
                        v28 := []int{}
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                } else {
                    v29 := GenS4{F0: 140}
                    var v30 int16 = 0
                    _ = v29
                    _ = v30
                }
                fmt.Println(v25, v2)
            }
            _ = v25
        } else {
            if (683 > 0) {
                var v31 uint32 = 2147483647
                switch "foo bar" {
                case "hello":
                    fmt.Println(v31, v2)
                }
                switch "" {
                case "":
                    if (!false) {
                        if true {
                            var v32 uint64 = (v2 * v2)
                            _ = v32
                        }
                        fmt.Println(v19)
                    } else {
                        if (-9223372036854775808 != 9223372036854775807) {
                            v33 := GenS1{F0: (v2 & (v2 - v2))}
                            var v34 uint64 = (v2 / ((509) | 1))
                            _ = v33
                            _ = v34
                        }
                        delete(v4, v31)
                        fmt.Println(v19, v2, v2)
                    }
                case "\n\t":
                    if true {
                        switch 32767 {
                        case -155:
                            if ((genFn1(65535, 1) <= 0) && true) {
                                var v35 uint = 18446744073709551615
                                for fi36 := 0; fi36 < 11; fi36++ {
                                    for fi37 := 0; fi37 < 7; fi37++ {
                                        v19 = -1
                                        _ = fi37
                                    }
                                    fmt.Println(v35, fi36)
                                    _ = fi36
                                }
                                v3[v31] = 1
                                _ = v35
                            } else {
                                v38 := GenS5{F0: 511, F1: 4294967295}
                                _ = v38
                            }
                        case -2147483648:
                            v39 := GenS5{F0: 2, F1: 1}
                            _ = v39
                        }
                        v19 = v19
                        v1 += (1.5 - (v1 * (v1 - (-2.25 * float64(-2.25)))))
                    } else {
                        var v40 string = "a"
                        var v41 float64 = ((v1 / 1.0) * (v1 - v1))
                        _ = v40
                        _ = v41
                    }
                case "foo bar":
                    v2 += 700
                }
                _ = v31
            }
        }
    }
    v19 = v19
    v19 = (v19 + 42)
    fmt.Println(v19, v2, v1)
    for fi42 := 0; fi42 < 10; fi42++ {
        switch ("\n\t" + "foo bar") {
        case "a":
            v43 := GenS4{F0: 83}
            _ = v43
        case "foo bar":
            _ = len(v3)
        case "\n\t":
            v44 := []int{fi42}
            _ = v44
        default:
            if false {
                switch fi42 {
                case -13:
                    switch "hello" {
                    case "a":
                        switch "hello" {
                        case "":
                            var v45 uint32 = 4294967295
                            _ = v45
                        }
                    default:
                        var v46 float32 = float32((0.0 - 1.5))
                        _ = v46
                    }
                case -367:
                    for fi47 := 0; fi47 < 15; fi47++ {
                        v48 := GenS3{F0: 0, F1: (v2 << ((42) & 63))}
                        delete(v4, uint32(v2))
                        switch -401 {
                        case -985:
                            v1 = v1
                        }
                        _ = fi47
                        _ = v48
                    }
                case 42:
                    if true {
                        v19 += (v19 + (v19 >> ((65535) & 63)))
                        v19 += 42
                    }
                }
                for fi49 := 0; fi49 < 18; fi49++ {
                    if false {
                        v1 = (((v1 / 1.0) - v1) / -2.25)
                        v2 += 1
                        v2 = 42
                    } else {
                        if (!true) {
                            var v50 float64 = ((v1 / 1.0) / 3.14159)
                            _ = v50
                        }
                        v19 -= 0
                    }
                    _ = fi49
                }
                if (fi42 <= (^(fi42))) {
                    switch (fi42 - (fi42 << ((1) & 63))) {
                    case -219:
                        v2 -= v2
                    case 32768:
                        delete(v4, 42)
                    case 32767:
                        fmt.Println(fi42, fi42, v1)
                    default:
                        v51 := GenS2{F0: 18446744073709551615}
                        _ = v51
                    }
                }
            }
        }
        v2 = v2
        v52 := GenS4{F0: 179}
        _ = fi42
        _ = v52
    }
    var v53 int8 = 121
    v54 := GenS2{F0: 1}
    switch "" {
    case "hello":
        v55 := GenS5{F0: 30, F1: 65535}
        _ = v55
    case "a":
        var v56 string = ""
        _ = v56
    }
    v57 := GenS6{F0: "a", F1: int(v19)}
    switch 4294967295 {
    case -713:
        fmt.Println(v19)
    }
    for fi58 := 0; fi58 < 9; fi58++ {
        switch fi58 {
        case 32768:
            v19 = (v19 + v19)
        case 42:
            var v59 uint16 = 2
            _ = v59
        case 2147483647:
            v60 := []int{fi58, (fi58 + (-(fi58)))}
            _ = v60
        }
        var v61 uint16 = 245
        _ = fi58
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v53:", v53)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v57
}
