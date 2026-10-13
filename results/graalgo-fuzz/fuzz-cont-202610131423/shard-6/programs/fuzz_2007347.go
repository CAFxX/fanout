package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 bool, p1 int32, p2 string) int8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        p0 = (float32(((1e10 * 1e10) / 1.0)) >= float32(0.1))
        fmt.Println(fi6)
        v7 := GenS1{F0: 9, F1: (p1 | (p1 ^ (p1 + (p1 % ((p1) | 1))))), F2: 42}
        _ = fi6
        _ = v7
    }
    switch p2 {
    case "foo bar":
        p2 = "\n\t"
    case "\n\t":
        for fi8 := 0; fi8 < 14; fi8++ {
            fmt.Println(p1)
            _ = fi8
        }
    case "a":
        var v9 uint32 = 997
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 52
}
func genFn1(p0 int32) uint {
    fmt.Println(p0)
    switch "\n\t" {
    case "hello":
        v10 := GenS1{F0: 261, F1: (p0 - p0), F2: 32767}
        _ = v10
    case "foo bar":
        switch "hello" {
        case "\n\t":
            var v11 float64 = 1.5
            _ = v11
        case "hello":
            switch "hello" {
            case "\n\t":
                v12 := []int{-992, -9223372036854775808, 2147483647}
                _ = v12
            case "":
                v13 := GenS1{F0: 1, F1: (p0 >> ((65535) & 63)), F2: -226}
                _ = v13
            case "a":
                for fi14 := 0; fi14 < 16; fi14++ {
                    if (false && false) {
                        p0 -= -2147483648
                        for fi15 := 0; fi15 < 7; fi15++ {
                            var v16 bool = true
                            if v16 {
                                if v16 {
                                    v16 = (!v16)
                                    v17 := GenS2{F0: ("hello" + "a"), F1: 139}
                                    _ = v17
                                }
                                var v18 uint = 1
                                var v19 float64 = ((1e10 + 3.14159) + 0.1)
                                p0 = p0
                                v16 = ((!v16) && v16)
                                _ = v18
                                _ = v19
                            }
                            v20 := GenS2{F0: "a", F1: 42}
                            _ = fi15
                            _ = v16
                            _ = v20
                        }
                    }
                    if true {
                        var v21 int = (fi14 ^ fi14)
                        for fi22 := 0; fi22 < 16; fi22++ {
                            v23 := []int{v21, fi22, 192, (fi22 % ((v21) | 1))}
                            fmt.Println(fi14, fi22)
                            _ = fi22
                            _ = v23
                        }
                        if false {
                            v24 := []int{(v21 & v21)}
                            _ = v24
                        } else {
                            v21 += 0
                            var v25 bool = false
                            _ = v25
                        }
                        _ = v21
                    }
                    _ = fi14
                }
            default:
                var v26 int32 = p0
                _ = v26
            }
        default:
            v27 := GenS2{F0: ((("hello" + "") + "hello") + ("\n\t" + (("foo bar" + "\n\t") + "a"))), F1: 164}
            _ = v27
        }
    default:
        p0 = -2147483648
    }
    v28 := []int{32767, 0, -1}
    var v29 string = "foo bar"
    var v30 uint8 = 128
    _ = p0
    _ = v28
    _ = v29
    _ = v30
    return 154
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint8 = 127
    var v3 float32 = 0.0
    v4 := make(map[int]string)
    v5 := make(map[uint32]uint32)
    if (!false) {
        v2 = v2
        switch 128 {
        case 322:
            for fi31 := 0; fi31 < 18; fi31++ {
                if (v1 == (-(v1))) {
                    switch ("hello" + "foo bar") {
                    case "\n\t":
                        v2 = (v2 & 21)
                    case "hello":
                        v2 -= 2
                    default:
                        var v32 uint = 0
                        _ = v32
                    }
                }
                v1 = -535
                if (!false) {
                    v1 -= -744
                    v33 := GenS1{F0: 122, F1: (v1 + -1), F2: 32767}
                    _ = v33
                }
                _ = fi31
            }
        }
    }
    if (v3 <= v3) {
        v34 := GenS3{F0: 4294967295, F1: true, F2: v2}
        _ = v34
    }
    fmt.Println(v1)
    for fi35 := 0; fi35 < 2; fi35++ {
        var v36 uint64 = 0
        switch (^(fi35)) {
        case 128:
            v36 = v36
        case 32768:
            v37 := []int{}
            _ = v37
        default:
            fmt.Println(v1, fi35)
        }
        if ((v36 | (^(v36))) == (v36 ^ (v36 << ((0) & 63)))) {
            v36 = 635
            switch fi35 {
            case 1:
                fmt.Println(v2, v36)
            }
            var v38 int64 = -9223372036854775808
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    v3 -= 1e10
    switch 32767 {
    case 32767:
        v39 := GenS1{F0: 255, F1: 42, F2: 1}
        _ = v39
    case 1:
        switch "hello" {
        case "a":
            v3 = (v3 / ((float32(1.5) + v3) * (v3 * float32(((0.1 + 3.14159) + ((0.0 / -2.25) - 1.5))))))
        case "":
            var v40 uint = 241
            _ = v40
        case "foo bar":
            v41 := GenS4{F0: 18446744073709551615}
            _ = v41
        default:
            switch ("\n\t" + "\n\t") {
            case "a":
                if false {
                    if false {
                        fmt.Println(v2)
                    } else {
                        v3 -= ((v3 * (v3 / v3)) / v3)
                    }
                } else {
                    switch 0 {
                    case -92:
                        switch -1 {
                        case 42:
                            v4[-173] = "foo bar"
                        case 127:
                            v3 = ((0.0 * v3) * v3)
                        case -2147483648:
                            var v42 uint16 = 993
                            _ = v42
                        }
                    case 255:
                        v1 = 786
                    }
                    v43 := GenS3{F0: 36, F1: false, F2: v2}
                    var v44 string = ""
                    _ = v43
                    _ = v44
                }
            default:
                var v45 int32 = v1
                _ = v45
            }
        }
    }
    fmt.Println(v1, v1, v1)
    var v46 uint16 = 1
    for fi47 := 0; fi47 < 8; fi47++ {
        for fi48 := 0; fi48 < 1; fi48++ {
            v2 = 0
            _ = fi48
        }
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
