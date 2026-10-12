package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 int8 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 int8; F2 int16 }
func genFn0() int32 {
    v7 := []int{}
    var v8 float32 = ((float32(1e10) - float32((3.14159 * 0.1))) + (0.0 * -2.25))
    v9 := GenS1{F0: 1, F1: 92, F2: 127}
    fmt.Println(v8)
    _ = v7
    _ = v8
    _ = v9
    return -981
}
func genFn1(p0 string, p1 uint16, p2 uint64) uint32 {
    fmt.Println(p0)
    v10 := []int{2147483647, -759, 32768}
    if (!true) {
        v11 := GenS1{F0: 223, F1: 897, F2: 215}
        _ = v11
    } else {
        var v12 float32 = -2.25
        fmt.Println(p1)
        var v13 uint = 40
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 375
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 1
    v3 := make(map[int]int)
    var v4 uint = 637
    v5 := make(map[uint32]int)
    var v6 uint = 561
    fmt.Println(v4, v6)
    var v14 uint16 = 2
    var v15 int = -2147483648
    v16 := []int{(v15 | (v15 * v15)), (-(v15)), (^(v15)), (v15 | 2147483647)}
    _ = len(v16)
    switch (v15 % ((v15) | 1)) {
    case 42:
        var v17 uint16 = v14
        _ = v17
    case -9223372036854775808:
        if false {
            if len(v16) > 0 {
                v16[0] = v15
            }
            if (!(false || true)) {
                fmt.Println(v6)
                if (-2.25 >= -2.25) {
                    _ = len(v16)
                    v18 := GenS2{F0: -1, F1: v4, F2: 106}
                    v6 = 65535
                    _ = v18
                } else {
                    v6 += v4
                    var v19 string = ("" + "foo bar")
                    for fi20 := 0; fi20 < 15; fi20++ {
                        v6 = v4
                        if len(v16) > 0 {
                            v16[0] = (fi20 % (((v15 | fi20)) | 1))
                        }
                        _ = fi20
                    }
                    _ = v19
                }
            }
            v14 -= 748
        } else {
            var v21 int64 = 0
            _ = v21
        }
    }
    if (550 <= (^(v2))) {
        if false {
            v14 = (v14 + (v14 ^ (v14 ^ v14)))
            v22 := GenS3{F0: -81, F1: 42}
            _ = v22
        }
        v23 := GenS2{F0: genFn0(), F1: (v6 >> ((255) & 63)), F2: 55}
        _ = v23
    }
    v15 = (-(v15))
    switch v15 {
    case -2147483648:
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 bool = false
            _ = fi24
            _ = v25
        }
    case 9223372036854775807:
        if true {
            var v26 float64 = 0.1
            if true {
                if false {
                    fmt.Println(v2, v6, v14)
                    var v27 int16 = 1000
                    if len(v16) > 0 {
                        v16[0] = (v15 >> ((v4) & 63))
                    }
                    _ = v27
                } else {
                    var v28 int8 = -110
                    if (!true) {
                        var v29 uint = (v6 - (v6 << ((4294967295) & 63)))
                        if true {
                            switch 128 {
                            case 9223372036854775807:
                                switch 42 {
                                case 82:
                                    var v30 int32 = genFn0()
                                    _ = v30
                                default:
                                    if (493 <= -1) {
                                        if false {
                                            v6 = (v29 + (v29 / (((v4 >> ((42) & 63))) | 1)))
                                            v1[genFn1("hello", 761, uint64(v4))] = (^(v2))
                                        } else {
                                            var v31 float64 = v26
                                            _ = v31
                                        }
                                    } else {
                                        var v32 int16 = -912
                                        v26 = ((v26 * -2.25) * float64(3.14159))
                                        if len(v16) > 0 {
                                            v16[0] = v15
                                        }
                                        _ = v32
                                    }
                                }
                            case -2147483648:
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    v33 := GenS4{F0: (v14 ^ 255), F1: 2}
                                    _ = v33
                                case "":
                                    fmt.Println(v29, v4, v6)
                                case "foo bar":
                                    var v34 int8 = (v28 + 0)
                                    _ = v34
                                }
                            case 32768:
                                v15 = v15
                            default:
                                v35 := v1[(v2 + (v2 ^ 2147483647))]
                                _ = v35
                            }
                        }
                        _ = v29
                    }
                    if len(v16) > 0 {
                        v16[0] = (v15 % (((v15 & (v15 * (v15 ^ 108)))) | 1))
                    }
                    _ = v28
                }
                for fi36 := 0; fi36 < 8; fi36++ {
                    for fi37 := 0; fi37 < 7; fi37++ {
                        if ((v4 / (((v4 * v4)) | 1)) >= (v6 >> ((0) & 63))) {
                            v16 = append(v16, fi36)
                            v38 := GenS5{F0: v14}
                            var v39 uint8 = uint8(v14)
                            _ = v38
                            _ = v39
                        }
                        _ = fi37
                    }
                    _ = fi36
                }
            }
            _ = v26
        }
    case 128:
        v6 -= (^(v6))
    }
    var v40 uint = (v4 / ((v4) | 1))
    if false {
        switch "foo bar" {
        case "foo bar":
            var v41 int16 = -273
            _ = v41
        case "a":
            fmt.Println(v6, v15, v14)
        case "":
            fmt.Println(v2, v6, v6)
        }
        for fi42 := 0; fi42 < 6; fi42++ {
            fmt.Println(v2)
            var v43 uint32 = v2
            _ = fi42
            _ = v43
        }
    } else {
        v44 := GenS6{F0: (v14 | v14), F1: false}
        v45 := GenS7{F0: v15}
        var v46 string = ("hello" + ("\n\t" + "hello"))
        _ = v44
        _ = v45
        _ = v46
    }
    v47 := v3[v15]
    var v48 int32 = 42
    if ((808 > v48) && false) {
        v14 = v14
    } else {
        v49 := v3[v47]
        _ = v49
    }
    v50 := GenS8{F0: 657, F1: 42, F2: 32767}
    v51 := GenS2{F0: v48, F1: (v6 * (v6 / (((^(v4))) | 1))), F2: 42}
    switch (("foo bar" + ("\n\t" + ("foo bar" + ("a" + "a")))) + ("foo bar" + "a")) {
    case "\n\t":
        var v52 uint32 = v2
        _ = v52
    case "a":
        if true {
            var v53 float32 = (-2.25 * -2.25)
            _ = v53
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
