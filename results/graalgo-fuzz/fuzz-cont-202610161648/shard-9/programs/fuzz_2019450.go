package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int; F2 uint64 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint }
type GenS6 struct { F0 string; F1 uint8; F2 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 bool; F2 uint64 }
func genFn0(p0 int8, p1 uint8, p2 int32) int {
    for fi6 := 0; fi6 < 3; fi6++ {
        p0 += p0
        _ = fi6
    }
    fmt.Println(p2)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 string = "\n\t"
        var v9 int8 = (p0 & p0)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint16 = 255
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 32767
}
func genFn1() int32 {
    v12 := GenS1{F0: 228, F1: 42}
    if true {
        v13 := []int{42, 894}
        var v14 float32 = (0.1 + 0.1)
        v14 = v14
        _ = v13
        _ = v14
    } else {
        if false {
            var v15 uint32 = 42
            _ = v15
        } else {
            v16 := GenS1{F0: 2147483648, F1: 585}
            var v17 float64 = 1e10
            _ = v16
            _ = v17
        }
        var v18 uint16 = 2
        var v19 uint32 = 789
        _ = v18
        _ = v19
    }
    var v20 uint = 460
    var v21 int64 = 858
    _ = v12
    _ = v20
    _ = v21
    return 1
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    var v4 int32 = -2147483648
    v5 := make(map[int]int)
    var v22 uint16 = 2
    if true {
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 13; fi24++ {
                fmt.Println(v3)
                _ = len(v2)
                _ = fi24
            }
            _ = fi23
        }
    } else {
        v25 := []int{}
        switch (v3 + "\n\t") {
        case "\n\t":
            var v26 bool = false
            _ = v26
        default:
            var v27 uint = 2
            _ = v27
        }
        var v28 uint = 2
        _ = v25
        _ = v28
    }
    var v29 int = genFn0(-117, 1, 1)
    v30 := GenS1{F0: (^(v1)), F1: 1}
    for fi31 := 0; fi31 < 2; fi31++ {
        if ((true && false) || false) {
            var v32 int8 = 100
            var v33 uint8 = 2
            v4 -= (v4 + (-(v4)))
            _ = v32
            _ = v33
        } else {
            fmt.Println(v29, fi31, v4)
            for fi34 := 0; fi34 < 4; fi34++ {
                switch (fi34 - v29) {
                case 2147483647:
                    var v35 int32 = genFn1()
                    _ = v35
                case 9223372036854775807:
                    v36 := []int{-363}
                    _ = v36
                case 127:
                    v22 += 2
                }
                var v37 bool = (!false)
                _ = fi34
                _ = v37
            }
            v38 := []int{fi31}
            _ = v38
        }
        if (!(566 <= 549)) {
            fmt.Println(fi31, v29)
            v39 := []int{}
            for fi40 := 0; fi40 < 1; fi40++ {
                var v41 uint = 642
                fmt.Println(v41)
                _ = len(v2)
                _ = fi40
                _ = v41
            }
            _ = v39
        } else {
            switch (v3 + (v3 + ("hello" + v3))) {
            case "":
                var v42 uint16 = (v22 % ((0) | 1))
                _ = v42
            case "a":
                for fi43 := 0; fi43 < 9; fi43++ {
                    v29 += v29
                    v44 := GenS2{F0: v4}
                    if false {
                        v45 := GenS3{F0: 127, F1: 1}
                        v4 = (v4 << ((516) & 63))
                        _ = v45
                    } else {
                        var v46 float32 = -2.25
                        if (2147483647 != (v1 >> ((4294967295) & 63))) {
                            v1 = (v1 & v1)
                        }
                        _ = v46
                    }
                    v22 += v22
                    _ = fi43
                    _ = v44
                }
            default:
                for fi47 := 0; fi47 < 4; fi47++ {
                    if (!((!false) || (1e10 > ((-2.25 + ((((1e10 - (3.14159 + -2.25)) - (-2.25 - (float64((3.14159 * 0.0)) + -2.25))) + 0.1) + 1e10)) + 3.14159)))) {
                        v3 = (v3 + v3)
                        var v48 int16 = 1000
                        v1 = (v1 >> ((764) & 63))
                        _ = v48
                    } else {
                        if false {
                            v49 := GenS1{F0: v1, F1: 2}
                            switch (fi31 ^ (fi31 % ((128) | 1))) {
                            case -9223372036854775808:
                                v50 := GenS4{F0: (v22 - v22), F1: int(v1), F2: 18446744073709551615}
                                _ = v50
                            default:
                                if true {
                                    var v51 bool = false
                                    _ = v51
                                }
                            }
                            _ = v49
                        }
                    }
                    _ = fi47
                }
            }
        }
        _ = fi31
    }
    v22 = v22
    v52 := GenS5{F0: (^(v4)), F1: 128, F2: 2}
    if ((86 < 2) && (2 > 4294967295)) {
        if (true || (false || false)) {
            fmt.Println(v4, v3)
            fmt.Println(v3)
            if (!(!(!(42 > 1)))) {
                var v53 int8 = 21
                v53 += v53
                _ = v53
            } else {
                v54 := GenS1{F0: v1, F1: 42}
                var v55 uint8 = 239
                _ = v54
                _ = v55
            }
        } else {
            var v56 uint8 = 122
            _ = len(v2)
            _ = v56
        }
        var v57 int32 = genFn1()
        _ = v57
    } else {
        var v58 float32 = (((((1.5 / (0.0 + (-2.25 / 0.1))) * 1.5) + -2.25) - 3.14159) / 1.0)
        v3 = "a"
        _ = v58
    }
    for fi59 := 0; fi59 < 19; fi59++ {
        v60 := GenS6{F0: ((v3 + ((v3 + "") + (v3 + (v3 + v3)))) + (v3 + v3)), F1: 127, F2: (v29 / ((genFn0(127, 51, 650)) | 1))}
        _ = fi59
        _ = v60
    }
    var v61 string = (("hello" + v3) + ((v3 + (v3 + "foo bar")) + v3))
    fmt.Println(v4, v1, v29)
    fmt.Println(v1)
    for fi62 := 0; fi62 < 13; fi62++ {
        v61 = "hello"
        v63 := GenS7{F0: 0}
        _ = fi62
        _ = v63
    }
    var v64 uint64 = 18446744073709551615
    v65 := []int{(-(v29)), (v29 | (v29 - 32768)), -509, v29}
    for fi66 := 0; fi66 < 8; fi66++ {
        if false {
            for fi67 := 0; fi67 < 16; fi67++ {
                if len(v65) > 0 {
                    v65[0] = fi66
                }
                v29 += (v29 % ((v29) | 1))
                switch (v3 + ("foo bar" + v3)) {
                case "":
                    var v68 int32 = (v4 - v4)
                    _ = v68
                default:
                    fmt.Println(v3)
                }
                _ = fi67
            }
        } else {
            var v69 int64 = 0
            v70 := GenS8{F0: 32767, F1: (1.5 <= 0.1), F2: v64}
            _ = v69
            _ = v70
        }
        var v71 int64 = 9223372036854775807
        if (true && true) {
            var v72 bool = false
            if v72 {
                var v73 float32 = ((0.1 * -2.25) / 1.5)
                _ = v73
            }
            v74 := GenS2{F0: -422}
            _ = v72
            _ = v74
        } else {
            for fi75 := 0; fi75 < 5; fi75++ {
                if len(v65) > 0 {
                    v65[0] = v29
                }
                _ = fi75
            }
            delete(v2, v1)
            var v76 int = int(v1)
            _ = v76
        }
        _ = fi66
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v61:", v61)
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v52
    _ = v61
    _ = v64
    _ = v65
}
