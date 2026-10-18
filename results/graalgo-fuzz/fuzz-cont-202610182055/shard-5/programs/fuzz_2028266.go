package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint64 }
type GenS2 struct { F0 bool; F1 int32; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint8; F2 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32 }
func genFn0() int64 {
    switch "hello" {
    case "a":
        var v7 string = "hello"
        _ = v7
    case "foo bar":
        var v8 int8 = -27
        _ = v8
    case "":
        var v9 int64 = 9223372036854775807
        _ = v9
    default:
        var v10 int64 = -435
        _ = v10
    }
    var v11 float32 = 3.14159
    v11 = float32((float64((1.5 * (-2.25 - v11))) * float64(float32(0.0))))
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: -534, F1: "foo bar", F2: 1}
        _ = fi12
        _ = v13
    }
    _ = v11
    return -1
}

func main() {
    var v1 uint64 = 2
    var v2 int64 = 42
    v3 := make(map[string]string)
    var v4 int8 = -1
    var v5 uint16 = 1
    v6 := []int{9223372036854775807}
    v1 += v1
    v3[((((("\n\t" + "foo bar") + "\n\t") + "a") + "") + (("a" + ("hello" + "\n\t")) + "\n\t"))] = "hello"
    if (65535 <= uint(v5)) {
        if true {
            var v14 uint = 0
            if (("hello" + "hello") == "foo bar") {
                v15 := GenS2{F0: (false && false), F1: -1, F2: -32768}
                _ = v15
            } else {
                var v16 string = (("a" + "foo bar") + ("" + ((("" + "hello") + "foo bar") + ("a" + (((("\n\t" + ("a" + "a")) + "hello") + "foo bar") + "hello")))))
                _ = len(v3)
                _ = v16
            }
            switch 2147483647 {
            case 1:
                if false {
                    v14 += (v14 << ((255) & 63))
                } else {
                    switch ("foo bar" + "") {
                    case "hello":
                        if (32767 == 678) {
                            if (v4 <= (-(v4))) {
                                if false {
                                    if true {
                                        for fi17 := 0; fi17 < 11; fi17++ {
                                            var v18 int64 = 9223372036854775807
                                            v18 = (v2 >> ((v14) & 63))
                                            var v19 int32 = 42
                                            _ = fi17
                                            _ = v18
                                            _ = v19
                                        }
                                        v14 -= v14
                                    }
                                    v20 := GenS3{F0: uint16(v1), F1: (v5 * v5), F2: (v1 + 4294967295)}
                                    _ = v20
                                } else {
                                    v21 := GenS1{F0: 42, F1: ("a" + ("" + "")), F2: v1}
                                    _ = v21
                                }
                            }
                            if (1e10 == 1.5) {
                                for fi22 := 0; fi22 < 20; fi22++ {
                                    var v23 float32 = 3.14159
                                    var v24 int32 = 324
                                    var v25 float64 = -2.25
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                v2 -= genFn0()
                            } else {
                                var v26 int32 = 2147483647
                                var v27 uint32 = 65535
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    v3[""] = "hello"
                                    var v29 int32 = (v26 - 1)
                                    _ = fi28
                                    _ = v29
                                }
                                _ = v26
                                _ = v27
                            }
                        } else {
                            if false {
                                v2 += v2
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            } else {
                                if false {
                                    var v30 float64 = 1e10
                                    _ = v30
                                }
                            }
                            var v31 int8 = -1
                            _ = v31
                        }
                    case "a":
                        var v32 int8 = (v4 + 42)
                        _ = v32
                    }
                }
            case -9223372036854775808:
                for fi33 := 0; fi33 < 2; fi33++ {
                    if (-1 != fi33) {
                        v14 = v14
                    }
                    var v34 uint16 = 255
                    _ = fi33
                    _ = v34
                }
            case 0:
                v14 += v14
            default:
                v2 += v2
            }
            _ = v14
        }
        var v35 int16 = 0
        var v36 int16 = 1
        _ = v35
        _ = v36
    } else {
        if (42 <= 820) {
            if (42 >= v1) {
                v37 := v3[("hello" + ("foo bar" + ("\n\t" + "foo bar")))]
                _ = v37
            } else {
                v1 = (v1 ^ 18446744073709551615)
            }
        }
    }
    if len(v6) > 0 {
        v6[0] = 42
    }
    v38 := GenS3{F0: v5, F1: 288, F2: (^(v1))}
    v39 := GenS4{F0: 65535, F1: 34, F2: (v1 * v1)}
    fmt.Println(v1, v5, v2)
    var v40 int32 = 956
    if true {
        if false {
            v3["a"] = "hello"
        }
    } else {
        if false {
            for fi41 := 0; fi41 < 5; fi41++ {
                if true {
                    switch "hello" {
                    case "":
                        var v42 int8 = v4
                        _ = v42
                    default:
                        if (!(-1 >= fi41)) {
                            var v43 int64 = 9223372036854775807
                            var v44 uint16 = (v5 * 2)
                            if (float32(-2.25) < 1e10) {
                                if len(v6) > 0 {
                                    v6[0] = (-(fi41))
                                }
                                _ = len(v3)
                                var v45 uint64 = 2
                                _ = v45
                            }
                            _ = v43
                            _ = v44
                        }
                    }
                    var v46 int8 = 43
                    var v47 bool = ((^(v40)) != v40)
                    _ = v46
                    _ = v47
                } else {
                    for fi48 := 0; fi48 < 3; fi48++ {
                        var v49 int64 = v2
                        v3[("" + "\n\t")] = ("\n\t" + (("foo bar" + "") + (("foo bar" + "a") + (("\n\t" + "a") + ""))))
                        _ = fi48
                        _ = v49
                    }
                    v40 += v40
                    var v50 string = "a"
                    _ = v50
                }
                v51 := GenS2{F0: true, F1: (v40 + 0), F2: 1000}
                v5 = v5
                _ = fi41
                _ = v51
            }
            v2 = v2
            var v52 float64 = 1.5
            _ = v52
        } else {
            v1 = 4294967295
            var v53 int32 = 1
            _ = v53
        }
    }
    if (true && true) {
        v54 := GenS3{F0: 818, F1: v5, F2: v1}
        switch int(v1) {
        case 2147483647:
            v55 := GenS5{F0: "a"}
            _ = v55
        case 32767:
            switch 1 {
            case 42:
                for fi56 := 0; fi56 < 20; fi56++ {
                    v1 = 1
                    var v57 int8 = -19
                    if false {
                        var v58 string = "foo bar"
                        v5 = v5
                        _ = v58
                    } else {
                        if (v5 < v5) {
                            var v59 int16 = 1
                            _ = v59
                        }
                        var v60 float32 = -2.25
                        _ = v60
                    }
                    _ = fi56
                    _ = v57
                }
            case 32767:
                v61 := GenS2{F0: false, F1: v40, F2: -299}
                _ = v61
            }
        case 42:
            var v62 float64 = 0.0
            _ = v62
        }
        _ = v54
    }
    switch 9223372036854775807 {
    case 127:
        switch (("a" + "\n\t") + "hello") {
        case "":
            for fi63 := 0; fi63 < 14; fi63++ {
                for fi64 := 0; fi64 < 9; fi64++ {
                    fmt.Println(v5, v4)
                    _ = fi64
                }
                for fi65 := 0; fi65 < 19; fi65++ {
                    v2 -= (v2 / (((v2 & genFn0())) | 1))
                    _ = fi65
                }
                var v66 bool = ((433 != 0) || (!false))
                _ = fi63
                _ = v66
            }
        case "\n\t":
            delete(v3, "a")
        default:
            for fi67 := 0; fi67 < 14; fi67++ {
                v68 := GenS6{F0: (0.0 > 1e10), F1: 0, F2: -44}
                if (false && false) {
                    if (32767 > 1) {
                        v5 = v5
                        _ = len(v6)
                    } else {
                        switch (fi67 | (fi67 >> ((2) & 63))) {
                        case 520:
                            v4 = v4
                        }
                    }
                }
                for fi69 := 0; fi69 < 6; fi69++ {
                    _ = len(v3)
                    var v70 uint = 18446744073709551615
                    _ = len(v3)
                    _ = fi69
                    _ = v70
                }
                _ = fi67
                _ = v68
            }
        }
    case 32767:
        fmt.Println(v40, v4, v1)
    case 9223372036854775807:
        if true {
            v71 := GenS2{F0: true, F1: v40, F2: 653}
            _ = v71
        } else {
            if (false && true) {
                v2 -= (v2 / ((v2) | 1))
            }
        }
    }
    v72 := GenS7{F0: false}
    v73 := GenS8{F0: (v40 % ((1) | 1))}
    fmt.Println(v5, v4, v40)
    if len(v6) > 0 {
        v6[0] = 32767
    }
    delete(v3, "foo bar")
    switch "\n\t" {
    case "a":
        if (!false) {
            var v74 uint32 = 558
            var v75 uint32 = v74
            v74 = (^(v74))
            _ = v74
            _ = v75
        }
    case "hello":
        var v76 uint = 386
        _ = v76
    case "":
        var v77 int16 = -887
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v72
    _ = v73
}
