package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint; F1 int16; F2 int64 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int16; F1 string; F2 uint32 }
type GenS7 struct { F0 uint32; F1 uint; F2 string }
func genFn0() uint {
    if false {
        switch -2147483648 {
        case 32767:
            switch "hello" {
            case "a":
                var v5 int32 = 2147483647
                _ = v5
            default:
                for fi6 := 0; fi6 < 10; fi6++ {
                    if true {
                        if false {
                            var v7 uint32 = 42
                            _ = v7
                        } else {
                            var v8 uint64 = 2
                            v9 := []int{fi6, fi6, -9223372036854775808, (^(fi6))}
                            v8 += 18446744073709551615
                            _ = v8
                            _ = v9
                        }
                    } else {
                        v10 := GenS1{F0: 255, F1: "foo bar"}
                        if false {
                            fmt.Println(fi6)
                            var v11 float32 = 1.5
                            _ = v11
                        } else {
                            var v12 uint = 56
                            var v13 bool = (100 != -1)
                            if (!(!true)) {
                                v14 := []int{fi6, (fi6 >> (((v12 * (v12 ^ 65535))) & 63))}
                                v15 := GenS1{F0: 411, F1: ("\n\t" + "hello")}
                                _ = v14
                                _ = v15
                            }
                            _ = v12
                            _ = v13
                        }
                        var v16 int = (fi6 - (fi6 << ((0) & 63)))
                        _ = v10
                        _ = v16
                    }
                    v17 := GenS2{F0: 415, F1: 1000, F2: -294}
                    _ = fi6
                    _ = v17
                }
            }
        case 32768:
            var v18 bool = false
            _ = v18
        }
        var v19 int32 = 1
        v19 = 753
        _ = v19
    } else {
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 float64 = 3.14159
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        if true {
            v23 := GenS1{F0: 362, F1: "a"}
            _ = v23
        } else {
            switch "foo bar" {
            case "\n\t":
                switch 128 {
                case 257:
                    var v24 uint64 = 0
                    _ = v24
                case 527:
                    var v25 int64 = 1
                    _ = v25
                case 42:
                    if (18446744073709551615 > 42) {
                        if false {
                            var v26 uint16 = 255
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 18; fi27++ {
                            if ((0.1 <= 0.1) && (-2.25 < 1e10)) {
                                v28 := []int{}
                                for fi29 := 0; fi29 < 15; fi29++ {
                                    v30 := GenS2{F0: 42, F1: 32767, F2: -1}
                                    _ = fi29
                                    _ = v30
                                }
                                switch "hello" {
                                case "foo bar":
                                    var v31 float64 = 0.1
                                    _ = v31
                                }
                                _ = v28
                            }
                            var v32 uint = 255
                            var v33 uint64 = 42
                            _ = fi27
                            _ = v32
                            _ = v33
                        }
                    }
                default:
                    v34 := []int{}
                    _ = v34
                }
            case "hello":
                var v35 bool = false
                _ = v35
            }
        }
        v36 := GenS3{F0: true, F1: 1}
        _ = v36
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 int = -9223372036854775808
        for fi39 := 0; fi39 < 13; fi39++ {
            if ((true || false) || false) {
                v38 = fi39
            } else {
                var v40 bool = false
                switch "\n\t" {
                case "\n\t":
                    switch 32768 {
                    case 127:
                        v41 := GenS4{F0: 2}
                        _ = v41
                    case 1:
                        if (32767 >= 2) {
                            switch ("foo bar" + ("foo bar" + ((("\n\t" + "") + "") + ""))) {
                            case "foo bar":
                                var v42 uint16 = 802
                                _ = v42
                            case "hello":
                                if v40 {
                                    v43 := GenS5{F0: (v38 % ((fi39) | 1)), F1: false}
                                    var v44 uint8 = 2
                                    _ = v43
                                    _ = v44
                                }
                            case "\n\t":
                                v38 += (v38 / (((fi37 ^ (v38 - (fi39 - fi39)))) | 1))
                            }
                            var v45 int8 = 47
                            fmt.Println(fi37, fi37, fi39)
                            _ = v45
                        } else {
                            var v46 string = "hello"
                            fmt.Println(v40, v46, v46)
                            _ = v46
                        }
                    default:
                        fmt.Println(fi37, v40)
                    }
                }
                var v47 int64 = -582
                _ = v40
                _ = v47
            }
            v38 -= fi37
            v38 -= fi37
            _ = fi39
        }
        _ = fi37
        _ = v38
    }
    return 0
}
func genFn1(p0 int8, p1 uint32) string {
    v48 := []int{32767, 32767}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v48
    return "foo bar"
}

func main() {
    var v1 uint16 = 255
    var v2 uint16 = v1
    v3 := make(map[string]string)
    v4 := []uint32{42, 875, 136}
    for fi49 := 0; fi49 < 5; fi49++ {
        for fi50 := 0; fi50 < 20; fi50++ {
            var v51 int64 = 42
            fmt.Println(v51)
            _ = fi50
            _ = v51
        }
        _ = fi49
    }
    var v52 int16 = 42
    if false {
        v53 := GenS6{F0: v52, F1: ("a" + "foo bar"), F2: 2147483648}
        var v54 uint = 415
        _ = v53
        _ = v54
    }
    var v55 int32 = -1
    if (-945 <= -835) {
        if ((-816 > 32768) || false) {
            v56 := GenS4{F0: 127}
            for fi57 := 0; fi57 < 5; fi57++ {
                v58 := GenS7{F0: 940, F1: 0, F2: "a"}
                if (18446744073709551615 < 1) {
                    if true {
                        v52 -= (v52 ^ v52)
                    }
                    switch fi57 {
                    case 32767:
                        for fi59 := 0; fi59 < 5; fi59++ {
                            fmt.Println(fi59, v2, v1)
                            if false {
                                var v60 uint64 = 948
                                v1 = v2
                                var v61 int64 = 0
                                _ = v60
                                _ = v61
                            } else {
                                delete(v3, "a")
                            }
                            _ = len(v4)
                            _ = fi59
                        }
                    case 47:
                        switch "a" {
                        case "a":
                            v62 := GenS4{F0: 2}
                            _ = v62
                        }
                    }
                }
                v55 = 424
                _ = fi57
                _ = v58
            }
            var v63 uint64 = 42
            _ = v56
            _ = v63
        }
        fmt.Println(v2, v1)
    } else {
        fmt.Println(v52, v55)
    }
    if ((true || (-1 >= -290)) && false) {
        var v64 uint16 = v2
        v2 -= v1
        _ = v64
    }
    var v65 int8 = 100
    if (false && (!(("a" + (("a" + "a") + "a")) == "\n\t"))) {
        v65 = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v55
    _ = v65
}
