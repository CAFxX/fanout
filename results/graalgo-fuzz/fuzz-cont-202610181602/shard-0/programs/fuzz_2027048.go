package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint32; F2 bool }
func genFn0(p0 uint64, p1 int, p2 uint8) uint {
    v4 := GenS1{F0: p0, F1: 231, F2: 4294967295}
    fmt.Println(p1)
    p0 = (p0 & 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return uint(p2)
}
func genFn1(p0 uint8, p1 int) uint {
    v5 := GenS1{F0: 121, F1: 0, F2: 65535}
    fmt.Println(p1, p1)
    var v6 uint = 693
    v7 := GenS2{F0: -677}
    var v8 uint = (v6 + genFn0(1, -2147483648, (p0 / ((2) | 1))))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return genFn0(4294967295, p1, p0)
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]string)
    var v3 bool = (947 > uint(v1))
    fmt.Println(v3, v1)
    fmt.Println(v1, v3)
    var v9 uint64 = 1
    var v10 int32 = 42
    switch "\n\t" {
    case "hello":
        switch 32768 {
        case 128:
            fmt.Println(v3, v3)
        }
    case "foo bar":
        v10 -= (v10 * -1)
    default:
        switch "foo bar" {
        case "":
            switch -1 {
            case 127:
                for fi11 := 0; fi11 < 11; fi11++ {
                    v12 := GenS2{F0: -64}
                    _ = fi11
                    _ = v12
                }
            default:
                if v3 {
                    v13 := GenS2{F0: 29}
                    for fi14 := 0; fi14 < 16; fi14++ {
                        v15 := GenS3{F0: 1, F1: v3}
                        v16 := GenS1{F0: 514, F1: (v1 << ((1) & 63)), F2: 4294967295}
                        var v17 int64 = -9223372036854775808
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    var v18 int8 = 116
                    _ = v13
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        for fi20 := 0; fi20 < 13; fi20++ {
                            v1 = v1
                            if v3 {
                                _ = len(v2)
                                var v21 uint32 = 2147483648
                                _ = v21
                            } else {
                                if v3 {
                                    switch (fi19 / ((fi19) | 1)) {
                                    case 128:
                                        v22 := GenS3{F0: 0, F1: ((v1 << ((42) & 63)) < v1)}
                                        _ = v22
                                    case 32768:
                                        fmt.Println(v9, v10)
                                    }
                                    var v23 int16 = 1
                                    var v24 int16 = v23
                                    _ = v23
                                    _ = v24
                                } else {
                                    v1 = v1
                                }
                                switch 9223372036854775807 {
                                case 127:
                                    v25 := []int{428, (fi20 | (fi20 * fi20)), (fi20 >> ((42) & 63)), fi19}
                                    _ = v25
                                default:
                                    var v26 int8 = -52
                                    _ = v26
                                }
                            }
                            if true {
                                v27 := []int{fi20, (fi20 + (fi20 * -127))}
                                v28 := GenS2{F0: 32767}
                                var v29 uint16 = (v1 % ((v1) | 1))
                                _ = v27
                                _ = v28
                                _ = v29
                            } else {
                                switch fi19 {
                                case -9223372036854775808:
                                    if v3 {
                                        v10 = 42
                                    } else {
                                        v1 = 151
                                        v10 += v10
                                    }
                                case 32767:
                                    var v30 uint = genFn1(42, (fi20 ^ (fi20 / (((fi19 % ((fi20) | 1))) | 1))))
                                    _ = v30
                                case -2147483648:
                                    var v31 int64 = -9223372036854775808
                                    _ = v31
                                }
                            }
                            _ = fi20
                        }
                        var v32 bool = v3
                        switch (("" + "\n\t") + ("foo bar" + "foo bar")) {
                        case "a":
                            switch ("a" + ("hello" + "hello")) {
                            case "\n\t":
                                v33 := GenS4{F0: 4294967295}
                                _ = v33
                            case "hello":
                                fmt.Println(v1, v32)
                            default:
                                fmt.Println(fi19, v1, v32)
                            }
                        case "\n\t":
                            v9 = (^(v9))
                        default:
                            if v32 {
                                v1 -= v1
                                var v34 uint8 = 2
                                v35 := GenS5{F0: "hello", F1: v34}
                                _ = v34
                                _ = v35
                            } else {
                                var v36 uint16 = 1
                                _ = v36
                            }
                        }
                        _ = fi19
                        _ = v32
                    }
                }
            }
        default:
            v37 := []int{0, -9223372036854775808}
            _ = v37
        }
    }
    var v38 float32 = float32(0.0)
    if ((2 > (v9 + 2)) && (!v3)) {
        var v39 int = 32767
        _ = v39
    } else {
        var v40 bool = (false && (v3 && false))
        var v41 uint64 = v9
        v42 := v2[1]
        _ = v40
        _ = v41
        _ = v42
    }
    var v43 bool = true
    for fi44 := 0; fi44 < 13; fi44++ {
        switch fi44 {
        case 32768:
            if (((0.1 / ((v38 / v38) + v38)) == (3.14159 - v38)) && v3) {
                if v3 {
                    v45 := GenS6{F0: v1, F1: 2147483648, F2: v3}
                    var v46 bool = v43
                    _ = v45
                    _ = v46
                }
                v43 = (v10 < v10)
            }
        default:
            var v47 float64 = 1.5
            _ = v47
        }
        _ = fi44
    }
    for fi48 := 0; fi48 < 10; fi48++ {
        var v49 uint32 = 42
        var v50 uint32 = (v49 >> ((18446744073709551615) & 63))
        var v51 string = ("hello" + ("" + ("" + "foo bar")))
        _ = fi48
        _ = v49
        _ = v50
        _ = v51
    }
    var v52 float32 = 0.1
    switch -951 {
    case 32768:
        var v53 int = 2147483647
        _ = v53
    case -989:
        if v3 {
            v54 := []int{802, -1}
            _ = v54
        } else {
            if ((!(true && (1 >= v10))) || (!v43)) {
                v43 = ("" >= ("a" + "a"))
            }
        }
    case 1:
        if v43 {
            v55 := GenS2{F0: -912}
            if true {
                fmt.Println(v1, v3, v10)
            }
            _ = v55
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v43:", v43)
    fmt.Println("v52:", v52)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v43
    _ = v52
    _ = v9
}
