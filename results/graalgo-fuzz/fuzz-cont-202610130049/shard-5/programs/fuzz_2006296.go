package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 int }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint16; F2 string }
type GenS6 struct { F0 int8; F1 uint16; F2 string }
func genFn0() int64 {
    v7 := GenS1{F0: (!false)}
    var v8 int8 = 39
    var v9 bool = false
    switch 708 {
    case -932:
        v9 = true
    case 1:
        v10 := GenS2{F0: 1, F1: v8, F2: 32768}
        _ = v10
    case -9223372036854775808:
        v11 := []int{9223372036854775807, 0, 0}
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 501
}

func main() {
    v1 := []uint32{}
    var v2 int8 = 116
    var v3 uint = 65535
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    var v6 int8 = v2
    v3 = 1
    var v12 bool = (1 > 0)
    for fi13 := 0; fi13 < 16; fi13++ {
        v6 = 42
        _ = fi13
    }
    v4[828] = (false || (-32768 != -184))
    v6 += (v6 + 71)
    fmt.Println(v12, v12, v3)
    switch -314 {
    case 289:
        v12 = false
    case -9223372036854775808:
        switch -9223372036854775808 {
        case 2147483647:
            if (!(!(true || (42 > 0)))) {
                v14 := v5[0]
                switch "\n\t" {
                case "":
                    v6 -= v2
                case "foo bar":
                    v6 = v2
                case "hello":
                    if (3.14159 != (3.14159 * 0.0)) {
                        var v15 uint16 = 250
                        _ = v15
                    } else {
                        v16 := GenS3{F0: 1, F1: v6, F2: v6}
                        var v17 uint16 = 32767
                        _ = v16
                        _ = v17
                    }
                default:
                    for fi18 := 0; fi18 < 18; fi18++ {
                        for fi19 := 0; fi19 < 20; fi19++ {
                            var v20 float64 = ((1e10 / 1.0) + (0.1 * float64(1.5)))
                            _ = len(v4)
                            _ = fi19
                            _ = v20
                        }
                        _ = fi18
                    }
                }
                fmt.Println(v14)
                _ = v14
            }
        case -258:
            var v21 uint16 = 65535
            _ = v21
        default:
            v22 := GenS4{F0: 32767}
            _ = v22
        }
    default:
        var v23 int = 2147483647
        _ = v23
    }
    if v12 {
        v3 = (^(v3))
        var v24 uint16 = 1
        _ = v24
    }
    _ = len(v4)
    if (1 > -1) {
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
        fmt.Println(v12)
        if true {
            if (255 >= 32767) {
                if (0.0 > 0.0) {
                    for fi25 := 0; fi25 < 5; fi25++ {
                        var v26 uint16 = 1
                        var v27 float32 = (0.0 - (float32(float64((-2.25 - 3.14159))) + -2.25))
                        if v12 {
                            switch "hello" {
                            case "foo bar":
                                _ = len(v5)
                            }
                            v1 = append(v1, 4294967295)
                            v26 -= (v26 | (v26 & (^(v26))))
                        }
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                } else {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        v2 = v6
                        for fi29 := 0; fi29 < 20; fi29++ {
                            if v12 {
                                v30 := GenS3{F0: 32767, F1: 30, F2: (v6 << ((291) & 63))}
                                var v31 uint16 = 1
                                delete(v4, (fi29 | fi29))
                                _ = v30
                                _ = v31
                            } else {
                                if false {
                                    if v12 {
                                        v32 := GenS5{F0: genFn0(), F1: 1, F2: (((("hello" + "foo bar") + ("hello" + ("" + (("" + "a") + "")))) + "hello") + "")}
                                        _ = v32
                                    } else {
                                        switch ("hello" + "") {
                                        case "":
                                            _ = len(v4)
                                            v33 := GenS6{F0: (v2 & (^(v6))), F1: 1, F2: "foo bar"}
                                            _ = v33
                                        default:
                                            var v34 int8 = 100
                                            _ = v34
                                        }
                                        v3 += (v3 | v3)
                                    }
                                    fmt.Println(fi29, v6)
                                    var v35 uint64 = 42
                                    _ = v35
                                }
                            }
                            fmt.Println(fi28, fi29)
                            _ = fi29
                        }
                        fmt.Println(v6, v3)
                        _ = fi28
                    }
                }
                var v36 string = "hello"
                v2 -= (v6 & -1)
                _ = v36
            } else {
                fmt.Println(v6, v6)
                var v37 uint8 = 2
                _ = v37
            }
            v38 := GenS1{F0: v12}
            var v39 string = (("foo bar" + ("" + "a")) + "\n\t")
            _ = v38
            _ = v39
        }
    } else {
        v40 := v4[-1]
        fmt.Println(v2, v6, v2)
        var v41 uint = (^(v3))
        _ = v40
        _ = v41
    }
    if v12 {
        fmt.Println(v6)
        for fi42 := 0; fi42 < 15; fi42++ {
            v6 -= (v6 - (v2 & (v6 / (((v6 | v2)) | 1))))
            switch (^(fi42)) {
            case 819:
                v3 -= 1
            case 9223372036854775807:
                v2 = v2
            }
            _ = fi42
        }
        if true {
            var v43 int16 = 1
            _ = v43
        } else {
            if v12 {
                fmt.Println(v6)
            }
        }
    }
    v3 -= 65535
    for fi44 := 0; fi44 < 17; fi44++ {
        if ((2 <= 18446744073709551615) || false) {
            fmt.Println(v12, fi44)
        } else {
            if ((false || (true || v12)) || (float64(0.1) > 1e10)) {
                v12 = (true || v12)
                v1 = append(v1, uint32(v3))
            }
        }
        _ = fi44
    }
    fmt.Println(v2)
    v45 := v4[0]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
