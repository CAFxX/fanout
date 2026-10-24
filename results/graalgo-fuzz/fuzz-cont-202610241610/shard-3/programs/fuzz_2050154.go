package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 int64; F1 uint; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int8; F1 int32; F2 bool }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 uint; F1 string }
func genFn0(p0 bool, p1 int8) uint32 {
    var v4 uint16 = 32767
    var v5 int = 0
    if false {
        fmt.Println(v5, p0, v4)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint, p1 uint8, p2 int64) int8 {
    if (!false) {
        for fi6 := 0; fi6 < 9; fi6++ {
            p1 += (^(p1))
            v7 := GenS1{F0: fi6, F1: 1, F2: 1}
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: -332, F1: 32767, F2: 160}
            var v10 string = "hello"
            v11 := []int{int(p2), (fi8 % ((818) | 1)), fi8, (-(fi8))}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    } else {
        v12 := []int{9223372036854775807, 32768, -1, -9223372036854775808}
        _ = v12
    }
    switch ("hello" + "\n\t") {
    case "":
        for fi13 := 0; fi13 < 8; fi13++ {
            p2 = 705
            if true {
                p0 = 65535
            } else {
                switch "a" {
                case "":
                    v14 := GenS2{F0: 0, F1: p0, F2: -128}
                    _ = v14
                case "hello":
                    if (2 < 142) {
                        p0 -= p0
                    } else {
                        p1 = 128
                        switch 127 {
                        case 2147483647:
                            for fi15 := 0; fi15 < 11; fi15++ {
                                if true {
                                    p2 += (-(p2))
                                    fmt.Println(p2, p2)
                                }
                                var v16 string = (("a" + (("a" + "hello") + ("hello" + "a"))) + "")
                                if (716 <= 0) {
                                    var v17 float32 = 0.0
                                    p2 += (-(p2))
                                    _ = v17
                                }
                                fmt.Println(p1, p0)
                                v18 := GenS3{F0: 42, F1: 42, F2: -1}
                                _ = fi15
                                _ = v16
                                _ = v18
                            }
                        case 424:
                            fmt.Println(p0, p2)
                        }
                    }
                case "\n\t":
                    for fi19 := 0; fi19 < 15; fi19++ {
                        var v20 int64 = p2
                        switch (-(fi19)) {
                        case 32767:
                            fmt.Println(fi13, fi13, p0)
                        case -1:
                            var v21 uint16 = 65535
                            p1 = (p1 | 127)
                            if true {
                                v22 := GenS4{F0: 165}
                                _ = v22
                            } else {
                                var v23 int = (fi19 % (((-(fi13))) | 1))
                                v24 := GenS3{F0: 42, F1: uint64(p1), F2: -1}
                                if true {
                                    if (v21 < (v21 ^ (^(v21)))) {
                                        var v25 uint32 = 1
                                        var v26 uint64 = 634
                                        var v27 uint32 = (v25 % (((v25 / ((2147483648) | 1))) | 1))
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                    v28 := GenS5{F0: 18446744073709551615, F1: fi19}
                                    _ = v28
                                }
                                _ = v23
                                _ = v24
                            }
                            _ = v21
                        }
                        fmt.Println(fi13, p0, fi13)
                        _ = fi19
                        _ = v20
                    }
                default:
                    v29 := []int{fi13, (^(fi13)), -61}
                    _ = v29
                }
            }
            for fi30 := 0; fi30 < 19; fi30++ {
                fmt.Println(fi30, p1, p2)
                _ = fi30
            }
            _ = fi13
        }
    case "hello":
        var v31 float32 = float32(0.0)
        _ = v31
    case "a":
        if (false && true) {
            var v32 bool = (false || (-120 < 1))
            var v33 uint32 = 1
            _ = v32
            _ = v33
        }
    }
    var v34 uint8 = 128
    v35 := []int{}
    if len(v35) >= 2 {
        v35 = v35[1:len(v35)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v35
    return -108
}

func main() {
    v1 := []int{410, 128, 127, 9223372036854775807}
    var v2 float32 = (0.1 / 1.0)
    var v3 float64 = 0.0
    var v36 uint8 = 128
    if false {
        v37 := GenS6{F0: -1, F1: 0, F2: false}
        _ = v37
    } else {
        fmt.Println(v3, v3, v36)
    }
    fmt.Println(v36)
    switch "foo bar" {
    case "a":
        v3 += ((((-2.25 / 1.5) + v3) / ((v3 - 1e10) - v3)) + ((3.14159 - float64(float32(3.14159))) * 1.5))
    case "foo bar":
        v38 := GenS3{F0: 0, F1: 4294967295, F2: -769}
        _ = v38
    case "":
        _ = len(v1)
    }
    v36 += v36
    for fi39 := 0; fi39 < 7; fi39++ {
        v40 := GenS7{F0: -1, F1: false}
        v3 -= v3
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 19; fi41++ {
        switch fi41 {
        case -9223372036854775808:
            _ = len(v1)
        case 1:
            var v42 uint64 = 0
            _ = v42
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        switch "foo bar" {
        case "hello":
            switch "hello" {
            case "\n\t":
                v43 := GenS3{F0: 2, F1: 2, F2: -672}
                _ = v43
            case "hello":
                for fi44 := 0; fi44 < 19; fi44++ {
                    v3 -= (((0.1 - 1.5) / 0.1) - -2.25)
                    if true {
                        var v45 int8 = 0
                        fmt.Println(fi41)
                        _ = v45
                    }
                    if (858 >= 42) {
                        var v46 int32 = -2147483648
                        if false {
                            switch "hello" {
                            case "hello":
                                v36 -= 42
                            default:
                                var v47 float32 = 0.1
                                _ = v47
                            }
                            v2 += v2
                        } else {
                            var v48 bool = (463 < 42)
                            var v49 uint16 = 604
                            var v50 float32 = (v2 / v2)
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                        fmt.Println(v46, fi44, fi44)
                        switch 128 {
                        case -1:
                            fmt.Println(v46, fi41)
                        default:
                            v51 := GenS8{F0: 110, F1: "\n\t"}
                            _ = v51
                        }
                        _ = v46
                    }
                    _ = fi44
                }
            }
        case "a":
            for fi52 := 0; fi52 < 2; fi52++ {
                _ = len(v1)
                _ = fi52
            }
        default:
            v3 += ((v3 + (0.1 - v3)) + (((((v3 * float64(1.5)) + 1.5) + float64(0.1)) * v3) * 3.14159))
        }
        var v53 uint8 = (v36 | 42)
        _ = fi41
        _ = v53
    }
    v3 = (((0.1 - v3) / -2.25) * (v3 - v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
}
