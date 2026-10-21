package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int16; F2 bool }
type GenS4 struct { F0 uint16 }
func genFn0() int16 {
    var v5 float32 = 1.5
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int64 = 666
            if (1 <= 548) {
                v9 := []int{(fi6 / (((fi6 * (fi7 + fi6))) | 1))}
                v5 += v5
                _ = v9
            } else {
                v8 -= v8
            }
            _ = fi7
            _ = v8
        }
        v5 = 1e10
        if (false || false) {
            v10 := []int{fi6, 32767}
            _ = v10
        } else {
            v11 := []int{(fi6 & 504), (fi6 & -1), fi6, fi6}
            v12 := GenS1{F0: 0, F1: 0}
            var v13 float64 = 1.5
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi6
    }
    v14 := []int{2147483647, 661, -68, 657}
    var v15 uint32 = 2147483648
    _ = v14
    _ = v15
    _ = v5
    return -1
}
func genFn1(p0 uint, p1 uint64, p2 uint32) int {
    var v16 int = -945
    var v17 float64 = float64(3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = 42
    var v3 uint32 = 1
    v4 := make(map[uint32]string)
    v2 += v2
    var v18 int32 = (v1 << ((42) & 63))
    v19 := []int{0, 42, 899, -677}
    for fi20 := 0; fi20 < 3; fi20++ {
        if false {
            if (32768 <= (fi20 + 2147483647)) {
                var v21 uint32 = v3
                _ = v21
            } else {
                var v22 float32 = 3.14159
                delete(v4, 711)
                _ = v22
            }
        }
        v23 := GenS2{F0: 42}
        switch "hello" {
        case "":
            if (188 < 139) {
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 uint32 = 1
                    var v26 uint = 18446744073709551615
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                switch "a" {
                case "\n\t":
                    _ = len(v4)
                case "foo bar":
                    v2 = 314
                case "":
                    switch 911 {
                    case -2147483648:
                        fmt.Println(fi20)
                    }
                }
                var v27 uint = 4294967295
                _ = v27
            } else {
                if (0 > v3) {
                    var v28 int8 = 0
                    _ = v28
                } else {
                    fmt.Println(v18)
                    var v29 uint8 = 85
                    _ = v29
                }
            }
        case "a":
            var v30 float64 = 1.5
            _ = v30
        }
        _ = fi20
        _ = v23
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        v1 += 0
        v18 += v18
        v18 -= v1
        _ = fi31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        if true {
            fmt.Println(fi32, v18, fi32)
            var v33 string = "foo bar"
            for fi34 := 0; fi34 < 14; fi34++ {
                v1 = v18
                switch -9223372036854775808 {
                case 4294967295:
                    v1 = v2
                case -1:
                    _ = len(v19)
                case 0:
                    var v35 int64 = -9223372036854775808
                    _ = v35
                default:
                    if (100 <= 0) {
                        fmt.Println(v33)
                    } else {
                        var v36 int16 = 32767
                        fmt.Println(fi32, v33)
                        v2 -= (-(v18))
                        _ = v36
                    }
                }
                _ = fi34
            }
            _ = v33
        }
        v19 = append(v19, (fi32 ^ genFn1(1, 494, v3)))
        _ = fi32
    }
    switch 369 {
    case 395:
        var v37 int8 = 0
        _ = v37
    }
    v38 := GenS1{F0: 65535, F1: 255}
    for fi39 := 0; fi39 < 7; fi39++ {
        var v40 uint = 2
        for fi41 := 0; fi41 < 7; fi41++ {
            fmt.Println(fi39, fi41)
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    if true {
        fmt.Println(v3)
        var v42 uint64 = 18446744073709551615
        _ = v42
    }
    fmt.Println(v1, v2)
    for fi43 := 0; fi43 < 18; fi43++ {
        fmt.Println(v2, v3, v2)
        var v44 uint16 = 255
        var v45 int = (fi43 % ((int(v44)) | 1))
        _ = fi43
        _ = v44
        _ = v45
    }
    if false {
        var v46 int8 = 127
        _ = v46
    } else {
        v47 := v4[(v3 | (v3 ^ (v3 >> ((1) & 63))))]
        v47 = "hello"
        _ = v47
    }
    v1 = 2147483647
    switch "" {
    case "":
        if false {
            var v48 uint16 = 0
            if (0 > uint(v48)) {
                var v49 int8 = 0
                _ = v49
            }
            _ = v48
        } else {
            delete(v4, v3)
            _ = len(v19)
            if true {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                if false {
                    v50 := GenS3{F0: (v3 << ((470) & 63)), F1: -32768, F2: false}
                    _ = v50
                } else {
                    v2 += v18
                    var v51 uint32 = v3
                    var v52 int16 = 617
                    _ = v51
                    _ = v52
                }
                v4[(^(v3))] = (("foo bar" + (("a" + "\n\t") + ("a" + (("foo bar" + ("hello" + "hello")) + "hello")))) + "")
            }
        }
    case "foo bar":
        for fi53 := 0; fi53 < 13; fi53++ {
            v54 := GenS2{F0: 32767}
            v55 := GenS4{F0: 255}
            switch fi53 {
            case -9223372036854775808:
                v2 = (v1 << ((475) & 63))
            case -1:
                var v56 uint64 = 1
                _ = v56
            case 42:
                for fi57 := 0; fi57 < 2; fi57++ {
                    var v58 int32 = v1
                    if true {
                        var v59 int = (fi53 >> ((0) & 63))
                        _ = v59
                    } else {
                        v58 = (v2 & v1)
                    }
                    var v60 uint16 = 2
                    _ = fi57
                    _ = v58
                    _ = v60
                }
            default:
                if ((1.5 <= float64(3.14159)) && false) {
                    switch (fi53 & (fi53 - (fi53 + fi53))) {
                    case 481:
                        var v61 float32 = (-2.25 + -2.25)
                        _ = v61
                    case 32767:
                        for fi62 := 0; fi62 < 7; fi62++ {
                            v1 = v2
                            var v63 float64 = float64((1e10 + -2.25))
                            if true {
                                var v64 bool = (-320 >= 1)
                                v1 -= (v18 * (v18 + (v1 | (v1 * v1))))
                                var v65 uint16 = 32767
                                _ = v64
                                _ = v65
                            }
                            _ = fi62
                            _ = v63
                        }
                    default:
                        fmt.Println(v3, v3)
                    }
                    _ = len(v19)
                }
            }
            _ = fi53
            _ = v54
            _ = v55
        }
    case "a":
        fmt.Println(v18)
    default:
        if (0.0 != 3.14159) {
            if len(v19) > 0 {
                v19[0] = genFn1(1, 42, 0)
            }
            v18 += v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
