package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int64; F1 bool; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 uint32 }
func genFn0(p0 int64) int {
    p0 = (^(p0))
    for fi5 := 0; fi5 < 14; fi5++ {
        if (!("a" < "hello")) {
            p0 = (p0 + p0)
        }
        _ = fi5
    }
    _ = p0
    return -2147483648
}
func genFn1(p0 int, p1 int) int16 {
    var v6 int16 = 0
    var v7 uint64 = 100
    switch 128 {
    case -2147483648:
        v8 := []int{(p1 / ((p1) | 1))}
        _ = v8
    case 300:
        v6 += (v6 << ((18446744073709551615) & 63))
    default:
        p0 = p1
    }
    if (float64(1e10) <= (0.0 + 0.0)) {
        var v9 int16 = 307
        _ = v9
    } else {
        v10 := []int{(p0 % ((127) | 1)), -523}
        v6 = 32767
        v10 = append(v10, (p1 & -209))
        _ = v10
    }
    p0 -= 2147483647
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := []uint32{671, 65535, 0, 65535}
    var v2 float32 = 3.14159
    v3 := []int{127, 128, 32767}
    var v4 float64 = 3.14159
    _ = len(v1)
    var v11 int32 = 42
    v12 := GenS1{F0: 42, F1: 1}
    for fi13 := 0; fi13 < 5; fi13++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    var v14 int = 9223372036854775807
    switch "a" {
    case "":
        v1 = append(v1, 728)
    case "foo bar":
        v15 := GenS2{F0: 0, F1: 0}
        _ = v15
    default:
        v3 = append(v3, (v14 ^ 4294967295))
    }
    v16 := GenS3{F0: 89, F1: true, F2: 65535}
    var v17 int64 = -1
    fmt.Println(v11, v4)
    var v18 uint16 = 646
    v19 := GenS4{F0: 2147483647}
    if false {
        v18 = (v18 % ((v18) | 1))
        if false {
            var v20 uint32 = 42
            _ = v20
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        fmt.Println(v18)
        if (3 > 255) {
            if false {
                var v21 string = ""
                _ = v21
            } else {
                v3 = append(v3, (v14 + (v14 << ((2) & 63))))
                if false {
                    if (-32768 > 32767) {
                        var v22 uint = 459
                        var v23 uint8 = 128
                        var v24 int = (v14 << ((0) & 63))
                        _ = v22
                        _ = v23
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 6; fi25++ {
                            var v26 string = ("" + ("foo bar" + ("foo bar" + "")))
                            if (181 < 2) {
                                v26 = ""
                                v18 += 0
                            }
                            switch genFn0((v17 & v17)) {
                            case -720:
                                var v27 string = (v26 + v26)
                                _ = v27
                            default:
                                for fi28 := 0; fi28 < 1; fi28++ {
                                    _ = len(v1)
                                    switch fi25 {
                                    case -9223372036854775808:
                                        if false {
                                            if (2 < 308) {
                                                var v29 int = v14
                                                v1 = append(v1, 65535)
                                                fmt.Println(v29, v29)
                                                _ = v29
                                            } else {
                                                if (!false) {
                                                    v30 := GenS1{F0: 2, F1: 32767}
                                                    fmt.Println(v2, v26)
                                                    _ = v30
                                                } else {
                                                    switch v26 {
                                                    case "foo bar":
                                                        if (!false) {
                                                            switch v26 {
                                                            case "\n\t":
                                                                if true {
                                                                    var v31 float32 = 0.0
                                                                    if true {
                                                                        if true {
                                                                            var v32 bool = (124 != uint32(v18))
                                                                            v33 := GenS5{F0: fi25, F1: 1, F2: 843}
                                                                            _ = v32
                                                                            _ = v33
                                                                        }
                                                                    }
                                                                    _ = v31
                                                                } else {
                                                                    v34 := GenS5{F0: 42, F1: v18, F2: 42}
                                                                    var v35 int16 = -1
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                            }
                                                            var v36 uint64 = 14
                                                            v37 := GenS6{F0: 0}
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            if (!(true && (v4 <= v4))) {
                                                                switch v26 {
                                                                case "\n\t":
                                                                    var v38 int32 = (v11 & v11)
                                                                    _ = v38
                                                                case "hello":
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                            } else {
                                                                _ = len(v3)
                                                                v39 := GenS7{F0: v11, F1: ("hello" + "hello")}
                                                                _ = v39
                                                            }
                                                            var v40 float64 = (v4 * (1.5 + ((1.5 / (v4 / v4)) * ((0.1 / (v4 * 0.1)) * ((((v4 + v4) - 0.1) - (0.0 * (((v4 - 1.5) / float64(float32(v4))) * v4))) / (v4 - v4))))))
                                                            v4 = v4
                                                            _ = v40
                                                        }
                                                    }
                                                    v4 += (1.5 / v4)
                                                    fmt.Println(v17)
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        } else {
                                            switch (((v26 + v26) + "") + (v26 + ("foo bar" + "\n\t"))) {
                                            case "foo bar":
                                                switch "hello" {
                                                case "a":
                                                    switch int(v17) {
                                                    case 2147483647:
                                                        fmt.Println(v18, v4)
                                                    case 9223372036854775807:
                                                        switch genFn0(v17) {
                                                        case -1:
                                                            v18 -= v18
                                                        case -222:
                                                            var v41 uint8 = 1
                                                            _ = v41
                                                        case -9223372036854775808:
                                                            var v42 int16 = 32767
                                                            _ = v42
                                                        default:
                                                            if ((((false && (v2 >= v2)) && (fi28 < -9223372036854775808)) && true) && (!true)) {
                                                                v43 := GenS4{F0: -163}
                                                                v26 = v26
                                                                _ = v43
                                                            } else {
                                                                v26 = ""
                                                            }
                                                        }
                                                    case 1:
                                                        var v44 uint32 = 280
                                                        _ = v44
                                                    }
                                                }
                                            case "a":
                                                var v45 int8 = 100
                                                v14 = genFn0(v17)
                                                v18 -= 0
                                                switch (fi28 & (v14 / ((4294967295) | 1))) {
                                                case 9223372036854775807:
                                                    v46 := GenS2{F0: 2, F1: 9223372036854775807}
                                                    _ = v46
                                                case -2147483648:
                                                    if false {
                                                        v47 := GenS4{F0: (v11 % ((0) | 1))}
                                                        if false {
                                                            v48 := GenS8{F0: 2147483648}
                                                            if len(v3) > 0 {
                                                                v3[0] = (fi25 % ((4294967295) | 1))
                                                            }
                                                            _ = v48
                                                        }
                                                        var v49 float32 = 3.14159
                                                        v17 -= (-(v17))
                                                        var v50 string = v26
                                                        _ = v47
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                case 127:
                                                    var v51 float32 = -2.25
                                                    _ = v51
                                                }
                                                switch v26 {
                                                case "hello":
                                                    if false {
                                                        v1 = append(v1, 65535)
                                                    }
                                                case "":
                                                    v4 += v4
                                                case "\n\t":
                                                    v52 := GenS5{F0: (fi28 * 42), F1: v18, F2: 991}
                                                    _ = v52
                                                default:
                                                    v53 := GenS2{F0: 42, F1: -90}
                                                    _ = v53
                                                }
                                                var v54 string = "\n\t"
                                                _ = v45
                                                _ = v54
                                            case "hello":
                                                var v55 uint64 = 0
                                                _ = v55
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v2 = (float32(0.1) / -2.25)
                                            }
                                        }
                                    case 213:
                                        var v56 uint = 65535
                                        _ = v56
                                    }
                                    _ = fi28
                                }
                            }
                            _ = fi25
                            _ = v26
                        }
                    }
                    for fi57 := 0; fi57 < 4; fi57++ {
                        v1 = append(v1, 42)
                        fmt.Println(v14)
                        _ = fi57
                    }
                    var v58 uint16 = (v18 & (v18 >> ((1) & 63)))
                    _ = v58
                } else {
                    var v59 string = "a"
                    _ = v59
                }
            }
            fmt.Println(v2, v14)
            _ = len(v1)
        }
    }
    v18 -= 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
