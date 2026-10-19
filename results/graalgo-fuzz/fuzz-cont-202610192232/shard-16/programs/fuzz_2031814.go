package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 bool; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS8 struct { F0 uint; F1 uint32 }
func genFn0(p0 int, p1 uint16) uint16 {
    if true {
        v6 := GenS1{F0: 1, F1: 843}
        p0 = (p0 + -2147483648)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := []int{(p0 >> ((168) & 63))}
            switch "" {
            case "\n\t":
                var v9 bool = false
                _ = v9
            case "hello":
                for fi10 := 0; fi10 < 16; fi10++ {
                    var v11 int = fi7
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi7
            _ = v8
        }
        p0 += p0
        v12 := []int{(p0 % ((p0) | 1)), p0}
        _ = v12
    }
    if (!false) {
        v13 := GenS1{F0: 42, F1: 4294967295}
        _ = v13
    } else {
        fmt.Println(p1, p1)
        p1 = (p1 | (p1 * 65535))
        p0 = (p0 | (p0 % (((p0 >> ((1) & 63))) | 1)))
    }
    switch "hello" {
    case "foo bar":
        v14 := []int{(p0 / ((128) | 1)), (p0 * p0), p0}
        _ = v14
    case "hello":
        v15 := []int{(^(p0)), p0, p0}
        _ = v15
    default:
        v16 := GenS2{F0: (65535 != 42)}
        _ = v16
    }
    var v17 int64 = 432
    if true {
        for fi18 := 0; fi18 < 18; fi18++ {
            switch (p0 * fi18) {
            case 42:
                v19 := GenS3{F0: "hello", F1: (0.0 <= -2.25), F2: (-793 > 879)}
                _ = v19
            default:
                fmt.Println(p0, p0)
            }
            var v20 uint = 18446744073709551615
            _ = fi18
            _ = v20
        }
        switch 32767 {
        case 0:
            var v21 int16 = -1
            _ = v21
        case 127:
            v22 := GenS3{F0: "", F1: false, F2: (!(65535 != 18446744073709551615))}
            _ = v22
        default:
            var v23 int16 = int16(p1)
            _ = v23
        }
        v24 := []int{p0, p0, (p0 & (^(p0)))}
        _ = v24
    } else {
        v25 := GenS3{F0: (("\n\t" + "\n\t") + "\n\t"), F1: true, F2: (true || false)}
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v17
    return p1
}
func genFn1(p0 uint8, p1 int8) uint {
    v26 := []int{-1, 0, 32768}
    v27 := GenS4{F0: uint64(p0)}
    var v28 uint = 259
    var v29 bool = (!false)
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    return 967
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 uint = 42
    var v5 uint8 = 122
    for fi30 := 0; fi30 < 5; fi30++ {
        switch "a" {
        case "foo bar":
            if (true && (false && false)) {
                var v31 uint = v4
                for fi32 := 0; fi32 < 20; fi32++ {
                    fmt.Println(v1)
                    v31 -= v4
                    v31 += v4
                    _ = fi32
                }
                _ = v31
            }
        case "a":
            v33 := GenS2{F0: true}
            _ = v33
        default:
            switch "foo bar" {
            case "a":
                delete(v2, fi30)
            case "hello":
                v34 := GenS4{F0: v1}
                _ = v34
            case "":
                for fi35 := 0; fi35 < 2; fi35++ {
                    switch (fi30 % (((fi30 | fi35)) | 1)) {
                    case 128:
                        v36 := []int{}
                        _ = v36
                    case 42:
                        if true {
                            var v37 int64 = 9223372036854775807
                            var v38 uint64 = v1
                            fmt.Println(fi35)
                            _ = v37
                            _ = v38
                        }
                    case -9223372036854775808:
                        if true {
                            if false {
                                v39 := v3["a"]
                                v40 := GenS4{F0: v1}
                                if (!false) {
                                    if ((v5 + (v5 ^ (^(v5)))) != (v5 & (^(v5)))) {
                                        v39 = (-2147483648 == -1)
                                        var v41 float32 = 0.0
                                        var v42 int8 = -128
                                        v43 := GenS5{F0: 2, F1: ""}
                                        v1 = 77
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    } else {
                                        var v44 float32 = 0.1
                                        fmt.Println(fi35, fi35, v39)
                                        v39 = (799 <= 32767)
                                        _ = v44
                                    }
                                }
                                v1 += v1
                                _ = v39
                                _ = v40
                            }
                            v5 = v5
                        }
                    default:
                        var v45 float64 = (((-2.25 + 3.14159) - 3.14159) - 1e10)
                        _ = v45
                    }
                    _ = fi35
                }
            }
        }
        for fi46 := 0; fi46 < 11; fi46++ {
            v1 = 758
            v47 := GenS6{F0: -61, F1: 4294967295}
            fmt.Println(fi30, v4, v4)
            _ = fi46
            _ = v47
        }
        _ = fi30
    }
    switch int(v5) {
    case 869:
        v48 := GenS7{F0: v1, F1: 538, F2: v1}
        _ = v48
    case -9223372036854775808:
        for fi49 := 0; fi49 < 19; fi49++ {
            for fi50 := 0; fi50 < 17; fi50++ {
                if false {
                    v51 := []int{int(v5), -1, fi49, (-(fi50))}
                    fmt.Println(fi49, v5)
                    var v52 uint32 = uint32(v4)
                    switch "foo bar" {
                    case "":
                        if false {
                            fmt.Println(v1, v1, v52)
                            fmt.Println(v52, v1, fi49)
                            switch "" {
                            case "\n\t":
                                var v53 string = ("" + "foo bar")
                                _ = v53
                            case "a":
                                if len(v51) > 0 {
                                    v51[0] = -619
                                }
                            case "foo bar":
                                _ = len(v3)
                            default:
                                v4 -= v4
                            }
                            if (!false) {
                                var v54 string = "a"
                                var v55 float64 = float64(float32(3.14159))
                                _ = v54
                                _ = v55
                            } else {
                                var v56 string = "\n\t"
                                switch fi50 {
                                case -1:
                                    v1 -= v1
                                case 2147483647:
                                    v57 := GenS4{F0: (v1 & (v1 - v1))}
                                    v58 := GenS6{F0: 9223372036854775807, F1: fi49}
                                    var v59 int16 = -1
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                switch (("foo bar" + ("hello" + v56)) + (v56 + "")) {
                                case "foo bar":
                                    switch -9223372036854775808 {
                                    case 396:
                                        fmt.Println(v4)
                                    case -1:
                                        var v60 uint32 = (^(v52))
                                        _ = v60
                                    case 9223372036854775807:
                                        var v61 int32 = 2147483647
                                        _ = v61
                                    }
                                case "":
                                    v62 := GenS8{F0: genFn1((v5 ^ v5), -128), F1: 1}
                                    _ = v62
                                }
                                _ = v56
                            }
                            var v63 bool = (!false)
                            _ = v63
                        } else {
                            if false {
                                v1 = (v1 | v1)
                                v64 := GenS7{F0: 1, F1: v1, F2: v1}
                                fmt.Println(v4)
                                _ = v64
                            } else {
                                fmt.Println(fi50)
                                v52 = uint32(v5)
                            }
                        }
                    }
                    _ = v51
                    _ = v52
                } else {
                    v5 = (v5 << ((4294967295) & 63))
                    v65 := GenS2{F0: false}
                    fmt.Println(v5)
                    _ = v65
                }
                switch -150 {
                case 1:
                    v66 := []int{(^(fi49)), 464}
                    _ = v66
                default:
                    v2[(fi50 * (fi50 << ((255) & 63)))] = int(v1)
                }
                if true {
                    v5 = 252
                    var v67 uint64 = v1
                    v5 += (v5 * 128)
                    var v68 int64 = -1
                    _ = v67
                    _ = v68
                }
                _ = fi50
            }
            _ = fi49
        }
    }
    v69 := []int{-97, 1, -9223372036854775808, 127}
    _ = len(v69)
    v5 = (^(v5))
    var v70 uint64 = v1
    var v71 uint16 = 895
    var v72 int16 = 392
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
