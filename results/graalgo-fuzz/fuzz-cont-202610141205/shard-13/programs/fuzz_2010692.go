package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 bool }
type GenS2 struct { F0 uint; F1 int64; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 int32; F2 int8 }
func genFn0(p0 string) string {
    if (((0.0 > 0.0) && (-1 > -32768)) || (8 == 42)) {
        p0 = "foo bar"
        if (!(-1 >= 2147483647)) {
            p0 = p0
            if false {
                v5 := []int{-2147483648, -1}
                _ = v5
            } else {
                var v6 int = 42
                var v7 float32 = 1e10
                p0 = ("hello" + ("hello" + "foo bar"))
                _ = v6
                _ = v7
            }
        } else {
            fmt.Println(p0)
            if true {
                p0 = ((p0 + p0) + "foo bar")
            }
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            p0 = ("" + "a")
            v9 := []int{fi8, (fi8 & -9223372036854775808), 2147483647, (fi8 & fi8)}
            _ = fi8
            _ = v9
        }
    } else {
        var v10 string = p0
        _ = v10
    }
    var v11 int = -9223372036854775808
    _ = p0
    _ = v11
    return "a"
}
func genFn1(p0 int64, p1 int64) uint8 {
    switch -9223372036854775808 {
    case -25:
        var v12 int16 = 32767
        _ = v12
    case 32768:
        fmt.Println(p1, p0)
    default:
        var v13 float64 = -2.25
        _ = v13
    }
    var v14 int32 = 0
    _ = p0
    _ = p1
    _ = v14
    return 127
}

func main() {
    var v1 bool = ((true && (!(-384 >= 0))) || true)
    v2 := make(map[uint32]bool)
    var v3 int64 = 0
    var v4 string = "\n\t"
    var v15 uint64 = 1
    if v1 {
        v15 = (v15 % ((v15) | 1))
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        if ((true && ((255 != 65535) && v1)) || false) {
            if v1 {
                v17 := []int{fi16, fi16, fi16, 32767}
                var v18 uint16 = 32767
                _ = v17
                _ = v18
            } else {
                var v19 float32 = 3.14159
                if (!v1) {
                    v3 = v3
                }
                _ = v19
            }
            if (54 > -128) {
                for fi20 := 0; fi20 < 7; fi20++ {
                    if v1 {
                        v21 := []int{(-(fi16)), (fi16 | 2147483647), -217}
                        if v1 {
                            v21 = append(v21, fi16)
                            v22 := v2[65535]
                            var v23 int8 = 1
                            var v24 string = (v4 + v4)
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            var v25 int32 = -2147483648
                            _ = len(v2)
                            _ = v25
                        }
                        _ = v21
                    } else {
                        v26 := GenS1{F0: 127, F1: (18446744073709551615 != v15), F2: (-128 >= 0)}
                        _ = v26
                    }
                    v27 := v2[65535]
                    v4 = v4
                    switch (v4 + "") {
                    case "a":
                        if true {
                            var v28 uint8 = 133
                            if true {
                                fmt.Println(v15)
                            } else {
                                if true {
                                    v29 := GenS1{F0: 0, F1: v1, F2: (!(v1 && (v27 || true)))}
                                    v3 = (v3 >> ((625) & 63))
                                    _ = v29
                                }
                                v30 := GenS2{F0: 2, F1: v3, F2: -454}
                                v31 := v2[42]
                                _ = v30
                                _ = v31
                            }
                            if (v1 || true) {
                                var v32 int8 = -70
                                v28 -= 0
                                _ = v32
                            } else {
                                v28 -= v28
                                v33 := []int{fi20, (fi16 - (fi20 / ((fi16) | 1)))}
                                v2[650] = true
                                _ = v33
                            }
                            _ = v28
                        }
                    default:
                        v34 := GenS3{F0: 2}
                        _ = v34
                    }
                    v35 := GenS4{F0: 1, F1: 1000, F2: 255}
                    var v36 float32 = (0.0 - 3.14159)
                    _ = fi20
                    _ = v27
                    _ = v35
                    _ = v36
                }
            }
        } else {
            v37 := GenS5{F0: v4}
            v3 += (v3 - v3)
            switch "a" {
            case "a":
                switch (^(fi16)) {
                case 640:
                    if (fi16 > fi16) {
                        if (true && v1) {
                            switch (^(fi16)) {
                            case 127:
                                if (v1 && (fi16 < -208)) {
                                    v38 := []int{}
                                    _ = v38
                                } else {
                                    delete(v2, 2147483647)
                                }
                            case -28:
                                v1 = false
                            default:
                                switch v4 {
                                case "\n\t":
                                    v4 = "hello"
                                case "foo bar":
                                    var v39 int16 = 276
                                    _ = v39
                                case "hello":
                                    var v40 int64 = 9223372036854775807
                                    _ = v40
                                default:
                                    fmt.Println(v1, v4)
                                }
                            }
                            v41 := []int{fi16}
                            if (fi16 != (^(fi16))) {
                                v1 = v1
                                var v42 float64 = ((((0.1 - 1.5) / (1e10 / float64(-2.25))) + 1e10) * 1e10)
                                _ = v42
                            } else {
                                if (225 != genFn1(v3, v3)) {
                                    v2[2147483648] = (v3 < (v3 - v3))
                                }
                            }
                            _ = v41
                        } else {
                            switch "hello" {
                            case "":
                                v43 := GenS6{F0: fi16, F1: -2147483648, F2: -128}
                                _ = v43
                            case "foo bar":
                                v1 = v1
                            }
                        }
                        var v44 uint = 0
                        var v45 float32 = float32(1e10)
                        _ = v44
                        _ = v45
                    } else {
                        var v46 uint16 = 1
                        _ = v46
                    }
                }
            }
            _ = v37
        }
        if false {
            v1 = v1
            var v47 string = ("" + v4)
            _ = v47
        } else {
            v1 = v1
        }
        _ = fi16
    }
    var v48 int64 = v3
    v3 += v3
    if v1 {
        switch (v4 + "foo bar") {
        case "foo bar":
            var v49 uint32 = 0
            _ = v49
        default:
            v4 = "hello"
        }
        v1 = v1
        if false {
            var v50 string = "a"
            var v51 int32 = 1
            var v52 int64 = (v48 & v48)
            _ = v50
            _ = v51
            _ = v52
        }
    }
    v53 := v2[0]
    v3 = 42
    fmt.Println(v15)
    var v54 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v53
    _ = v54
}
