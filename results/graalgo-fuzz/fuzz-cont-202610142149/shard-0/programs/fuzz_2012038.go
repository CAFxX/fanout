package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 uint64; F1 uint; F2 int }
type GenS6 struct { F0 bool; F1 bool; F2 uint64 }
type GenS7 struct { F0 int32; F1 uint16; F2 int64 }
func genFn0(p0 uint64, p1 int8, p2 int) uint16 {
    p1 -= p1
    switch p2 {
    case -256:
        var v5 bool = (-1 != 427)
        _ = v5
    case 32767:
        v6 := []int{(-(p2)), (-(p2)), p2, p2}
        _ = v6
    }
    v7 := GenS1{F0: 2}
    v8 := GenS1{F0: 37}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 276
}
func genFn1() int8 {
    v9 := GenS2{F0: 3, F1: -2147483648, F2: 542}
    var v10 float32 = ((1.5 - 0.1) * 0.1)
    var v11 uint64 = 0
    _ = v10
    _ = v11
    _ = v9
    return -128
}

func main() {
    var v1 int32 = -926
    v2 := make(map[uint32]string)
    v3 := []uint32{42, 649, 65535, 333}
    var v4 uint16 = 2
    v12 := GenS3{F0: 42}
    if (false || false) {
        _ = len(v3)
        if false {
            v4 -= 65535
        } else {
            var v13 uint = 255
            v14 := GenS4{F0: (v13 & (^(v13))), F1: 0}
            _ = v13
            _ = v14
        }
    } else {
        var v15 int64 = 9223372036854775807
        if false {
            switch "foo bar" {
            case "\n\t":
                delete(v2, 1)
            case "foo bar":
                v16 := GenS2{F0: v4, F1: -1, F2: 501}
                _ = v16
            }
            v17 := GenS2{F0: v4, F1: v1, F2: 1}
            _ = v17
        }
        _ = v15
    }
    var v18 uint32 = 178
    v19 := GenS5{F0: 4294967295, F1: 1, F2: -1}
    for fi20 := 0; fi20 < 13; fi20++ {
        v21 := GenS6{F0: true, F1: true, F2: 2}
        _ = fi20
        _ = v21
    }
    v18 = v18
    switch 32767 {
    case 123:
        v18 += 42
    case -550:
        if ((^(v18)) >= (v18 >> ((18446744073709551615) & 63))) {
            fmt.Println(v4, v1, v18)
            switch "a" {
            case "a":
                v22 := v2[4294967295]
                _ = v22
            case "\n\t":
                if true {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v4 = v4
                        _ = fi23
                    }
                    v24 := GenS7{F0: v1, F1: (v4 / ((v4) | 1)), F2: 42}
                    _ = v24
                }
            case "":
                for fi25 := 0; fi25 < 15; fi25++ {
                    if len(v3) > 0 {
                        v3[0] = v18
                    }
                    _ = fi25
                }
            default:
                var v26 int8 = genFn1()
                _ = v26
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v18 -= (v18 << ((486) & 63))
            var v27 string = ""
            _ = v27
        }
    }
    v4 -= (v4 << ((851) & 63))
    for fi28 := 0; fi28 < 11; fi28++ {
        if true {
            v29 := GenS7{F0: (v1 << ((0) & 63)), F1: 32767, F2: 1}
            fmt.Println(v18, v1, fi28)
            for fi30 := 0; fi30 < 7; fi30++ {
                v18 -= (v18 + uint32(v4))
                v1 -= 2147483647
                v1 = (v1 + (v1 * v1))
                v1 = v1
                _ = fi30
            }
            _ = v29
        }
        _ = fi28
    }
    if true {
        var v31 float64 = (((1.5 + 3.14159) + 1.5) + ((1.5 + ((0.1 * (0.0 + float64(1e10))) + 1e10)) * float64(0.1)))
        switch 1 {
        case 2147483647:
            var v32 string = ("\n\t" + (("hello" + "\n\t") + ("foo bar" + ((("foo bar" + "") + "foo bar") + "\n\t"))))
            _ = v32
        case 32768:
            if (!false) {
                v4 -= 2
            } else {
                v18 -= v18
                var v33 int32 = v1
                _ = v33
            }
        case 42:
            if true {
                if ("a" >= "") {
                    for fi34 := 0; fi34 < 15; fi34++ {
                        if (!((true && true) && (("" + "") == "a"))) {
                            delete(v2, (v18 + (v18 << ((451) & 63))))
                            if (42 <= 42) {
                                var v35 string = "a"
                                if true {
                                    v1 = 42
                                    v36 := GenS6{F0: true, F1: true, F2: 736}
                                    v37 := v2[(v18 + v18)]
                                    _ = v36
                                    _ = v37
                                } else {
                                    if (127 != 255) {
                                        delete(v2, v18)
                                        v18 -= (v18 - 0)
                                        var v38 int64 = 0
                                        _ = v38
                                    } else {
                                        var v39 uint32 = (^(v18))
                                        _ = v39
                                    }
                                    if false {
                                        for fi40 := 0; fi40 < 7; fi40++ {
                                            v31 -= (v31 / (v31 + (v31 + ((1.5 * (-2.25 + ((((v31 + (0.0 * (v31 / float64(float32((((v31 / v31) / v31) * (v31 - v31))))))) / (v31 + 3.14159)) + (v31 / v31)) * v31))) + ((v31 + (((float64((float32(1e10) / (0.1 + (-2.25 / 1e10)))) * v31) * (0.0 - v31)) - v31)) - v31)))))
                                            _ = fi40
                                        }
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    v41 := GenS4{F0: 18446744073709551615, F1: 1}
                                    _ = v41
                                }
                                _ = v35
                            } else {
                                if ((fi34 % ((-330) | 1)) <= (fi34 ^ fi34)) {
                                    var v42 bool = (!((!true) || ((1e10 + 0.1) <= 0.0)))
                                    _ = v42
                                }
                                v43 := GenS5{F0: 906, F1: 1, F2: fi34}
                                _ = v43
                            }
                        }
                        v4 += (^(v4))
                        _ = fi34
                    }
                    var v44 int = 128
                    v44 -= 292
                    _ = v44
                } else {
                    fmt.Println(v1)
                    v45 := GenS5{F0: 18446744073709551615, F1: 940, F2: 42}
                    _ = v45
                }
                v18 = v18
            }
        }
        _ = v31
    } else {
        var v46 uint16 = v4
        if false {
            var v47 uint64 = uint64(v46)
            _ = v47
        }
        v2[v18] = "a"
        _ = v46
    }
    var v48 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v48
}
