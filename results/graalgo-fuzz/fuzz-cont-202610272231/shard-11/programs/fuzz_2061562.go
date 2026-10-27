package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int64; F1 uint64; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 uint }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint32) int32 {
    v6 := GenS1{F0: 0, F1: 2147483647}
    switch -2147483648 {
    case -904:
        if false {
            for fi7 := 0; fi7 < 9; fi7++ {
                for fi8 := 0; fi8 < 7; fi8++ {
                    switch (fi8 / (((fi7 | fi8)) | 1)) {
                    case 447:
                        switch "hello" {
                        case "":
                            v9 := []int{979}
                            _ = v9
                        case "hello":
                            if (9223372036854775807 != 1) {
                                p1 -= 4294967295
                                v10 := GenS2{F0: 816, F1: (p0 >> ((65535) & 63)), F2: 113}
                                _ = v10
                            } else {
                                var v11 int32 = 2147483647
                                if (0 != 117) {
                                    v12 := GenS1{F0: 213, F1: -400}
                                    var v13 int = fi8
                                    fmt.Println(v13)
                                    _ = v12
                                    _ = v13
                                } else {
                                    var v14 int = (-(fi7))
                                    _ = v14
                                }
                                var v15 float64 = ((float64(1.5) * (0.1 + 0.1)) / 0.1)
                                _ = v11
                                _ = v15
                            }
                        case "foo bar":
                            fmt.Println(fi8)
                        }
                    case 311:
                        v16 := GenS2{F0: 1, F1: 4294967295, F2: 42}
                        _ = v16
                    }
                    _ = fi8
                }
                _ = fi7
            }
            p1 += p1
        }
    default:
        if false {
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := GenS3{F0: 482}
                if (!(!false)) {
                    if false {
                        for fi19 := 0; fi19 < 11; fi19++ {
                            var v20 int32 = -1
                            switch "foo bar" {
                            case "\n\t":
                                v21 := []int{fi19}
                                _ = v21
                            default:
                                var v22 int64 = 9223372036854775807
                                _ = v22
                            }
                            v20 -= -2147483648
                            _ = fi19
                            _ = v20
                        }
                        v23 := GenS2{F0: 517, F1: p0, F2: int8(p1)}
                        fmt.Println(p0)
                        _ = v23
                    }
                    if (!(!(192 >= 42))) {
                        var v24 bool = true
                        var v25 float64 = 0.0
                        var v26 string = "hello"
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                } else {
                    v27 := GenS3{F0: -32768}
                    for fi28 := 0; fi28 < 14; fi28++ {
                        if false {
                            switch fi17 {
                            case 32767:
                                p0 = p0
                            case -1:
                                fmt.Println(fi17, fi28)
                                var v29 float32 = (-2.25 / 1e10)
                                _ = v29
                            default:
                                p1 += 386
                            }
                        }
                        _ = fi28
                    }
                    _ = v27
                }
                for fi30 := 0; fi30 < 12; fi30++ {
                    if false {
                        if true {
                            var v31 int64 = -1
                            fmt.Println(p1, p0)
                            _ = v31
                        }
                        p1 += p1
                        if true {
                            if (65535 != 435) {
                                var v32 string = "foo bar"
                                v33 := []int{}
                                var v34 bool = ((!true) || true)
                                _ = v32
                                _ = v33
                                _ = v34
                            } else {
                                var v35 uint64 = 4294967295
                                var v36 int32 = 42
                                if (("hello" + "") <= "") {
                                    var v37 uint32 = 2147483648
                                    v38 := []int{(fi17 | fi30), (fi17 / ((fi17) | 1))}
                                    _ = v37
                                    _ = v38
                                } else {
                                    v35 = (v35 - 2)
                                }
                                _ = v35
                                _ = v36
                            }
                            p0 = (^(p0))
                            switch (fi17 >> ((uint(p0)) & 63)) {
                            case 128:
                                v39 := []int{}
                                _ = len(v39)
                                _ = v39
                            }
                        }
                    }
                    var v40 bool = (0.1 != -2.25)
                    _ = fi30
                    _ = v40
                }
                _ = fi17
                _ = v18
            }
            fmt.Println(p0, p0)
            v41 := GenS2{F0: -1, F1: p0, F2: 42}
            _ = v41
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 761
}
func genFn1() uint8 {
    switch 9223372036854775807 {
    case 1:
        if false {
            var v42 int64 = 42
            switch ("a" + ("hello" + "")) {
            case "a":
                var v43 float64 = float64((-2.25 + float32(0.1)))
                _ = v43
            case "hello":
                v44 := GenS4{F0: -89, F1: 78, F2: 18446744073709551615}
                _ = v44
            }
            _ = v42
        } else {
            if true {
                for fi45 := 0; fi45 < 20; fi45++ {
                    var v46 uint64 = 42
                    v46 = 891
                    _ = fi45
                    _ = v46
                }
            } else {
                var v47 uint = 67
                _ = v47
            }
            v48 := []int{4294967295, -572, 128}
            _ = v48
        }
    default:
        var v49 float64 = (0.0 + 0.0)
        _ = v49
    }
    var v50 uint16 = 1
    _ = v50
    return 128
}

func main() {
    v1 := []uint32{65535, 1, 1, 756}
    var v2 int64 = 102
    var v3 float32 = 0.1
    var v4 float32 = (((v3 * v3) / (v3 * (-2.25 - v3))) + 0.0)
    var v5 int64 = (v2 ^ (v2 + 0))
    var v51 float32 = (v3 + (3.14159 - v4))
    v1 = append(v1, 8)
    if (2 > 2) {
        v51 -= float32(-2.25)
        var v52 string = ("" + ("a" + "foo bar"))
        _ = v52
    }
    v4 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += 9223372036854775807
    v53 := GenS5{F0: 601, F1: 602}
    if (("a" + "foo bar") != (("hello" + "foo bar") + "a")) {
        var v54 uint = 4294967295
        _ = v54
    } else {
        var v55 uint16 = 65535
        _ = v55
    }
    v2 += 9223372036854775807
    var v56 int16 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        fmt.Println(v3, v2)
        switch 32768 {
        case 4:
            var v57 uint64 = 798
            _ = v57
        case 42:
            v3 -= (v51 - ((v4 - v4) * 0.0))
        case -718:
            if (9 > 166) {
                if (false || false) {
                    v2 = (v2 / ((-368) | 1))
                    switch -381 {
                    case -9223372036854775808:
                        v1 = append(v1, 4294967295)
                    case 32768:
                        v58 := GenS6{F0: 152}
                        _ = v58
                    case 4294967295:
                        fmt.Println(v3, v2, v2)
                    }
                    v51 = (3.14159 * v3)
                }
                switch 128 {
                case 617:
                    v3 = 0.1
                }
            }
        default:
            switch 0 {
            case -680:
                var v59 int16 = v56
                _ = v59
            }
        }
    }
    switch "a" {
    case "hello":
        var v60 int8 = 127
        _ = v60
    case "\n\t":
        fmt.Println(v2, v56, v3)
    case "foo bar":
        var v61 uint8 = 255
        _ = v61
    default:
        v56 = (v56 ^ 1000)
    }
    var v62 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v56:", v56)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v53
    _ = v56
    _ = v62
}
