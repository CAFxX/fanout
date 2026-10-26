package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 uint16; F1 int16; F2 int8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool; F1 bool; F2 bool }
func genFn0() bool {
    var v4 uint16 = 985
    v4 = (v4 | 1)
    fmt.Println(v4)
    switch 32768 {
    case 128:
        fmt.Println(v4)
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        v4 += 0
        var v6 uint = 650
        v7 := []int{fi5, fi5}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = v4
    return ((false && (624 == 2147483648)) || (!false))
}
func genFn1(p0 string, p1 uint8) uint {
    v8 := GenS1{F0: (^(p1))}
    p1 = (p1 & 2)
    _ = p0
    _ = p1
    _ = v8
    return 0
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]string)
    var v3 int8 = -14
    switch ("foo bar" + "") {
    case "hello":
        for fi9 := 0; fi9 < 15; fi9++ {
            switch ("hello" + "") {
            case "\n\t":
                if true {
                    if false {
                        v10 := []int{(fi9 + fi9), int(v3), (fi9 & 163)}
                        v1 = 42
                        if (!false) {
                            for fi11 := 0; fi11 < 3; fi11++ {
                                fmt.Println(v3, fi11, fi11)
                                if false {
                                    if (4294967295 == genFn1("hello", 42)) {
                                        switch "foo bar" {
                                        case "hello":
                                            var v12 int32 = 42
                                            _ = v12
                                        case "":
                                            v3 -= v3
                                        default:
                                            var v13 uint64 = 97
                                            _ = v13
                                        }
                                        v3 -= 127
                                        if (-122 != -2147483648) {
                                            if len(v10) > 0 {
                                                v10[0] = fi11
                                            }
                                            var v14 int = 32767
                                            if (false || (v3 <= (v3 - -90))) {
                                                if (0 <= 1) {
                                                    _ = len(v10)
                                                    switch (fi11 & (fi9 - -1)) {
                                                    case -9223372036854775808:
                                                        if (!(1000 < 734)) {
                                                            var v15 uint8 = 255
                                                            _ = v15
                                                        }
                                                    case 128:
                                                        if (!((0.0 * 3.14159) < (0.1 + 0.1))) {
                                                            var v16 int64 = int64(fi9)
                                                            var v17 int = fi9
                                                            var v18 bool = (!(-509 >= 32767))
                                                            _ = v16
                                                            _ = v17
                                                            _ = v18
                                                        }
                                                    case 0:
                                                        var v19 int64 = 0
                                                        _ = v19
                                                    }
                                                } else {
                                                    var v20 uint8 = 20
                                                    _ = v20
                                                }
                                            }
                                            v3 -= (v3 + v3)
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v1, v3, fi11)
                                            case "hello":
                                                if false {
                                                    if (3.14159 == (0.1 + 0.1)) {
                                                        delete(v2, fi11)
                                                        v14 = fi11
                                                    } else {
                                                        var v21 float32 = 0.1
                                                        var v22 uint8 = 1
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                    if true {
                                                        v23 := GenS2{F0: 2, F1: -32768}
                                                        fmt.Println(fi11, v14)
                                                        _ = v23
                                                    }
                                                } else {
                                                    var v24 int64 = 0
                                                    delete(v2, (fi9 - (fi11 << ((2) & 63))))
                                                    _ = v24
                                                }
                                            case "foo bar":
                                                if false {
                                                    if ((v1 < 559) && true) {
                                                        _ = len(v2)
                                                        if (((!false) && (!(!((v1 / ((v1) | 1)) <= (v1 & v1))))) || ("hello" <= "a")) {
                                                            v25 := GenS1{F0: 255}
                                                            var v26 float32 = 3.14159
                                                            switch "a" {
                                                            case "hello":
                                                                v27 := GenS3{F0: -32768}
                                                                _ = v27
                                                            default:
                                                                var v28 uint16 = (v1 % ((v1) | 1))
                                                                _ = v28
                                                            }
                                                            _ = v25
                                                            _ = v26
                                                        } else {
                                                            v29 := GenS4{F0: (v3 * 0)}
                                                            _ = v29
                                                        }
                                                        if len(v10) > 0 {
                                                            v10[0] = (v14 | fi9)
                                                        }
                                                    } else {
                                                        v1 += v1
                                                        var v30 string = "\n\t"
                                                        _ = v30
                                                    }
                                                    if false {
                                                        if false {
                                                            v10 = append(v10, -762)
                                                            v31 := GenS5{F0: 32767}
                                                            v2[v14] = "foo bar"
                                                            _ = v31
                                                        } else {
                                                            fmt.Println(v1, v3)
                                                            v32 := GenS6{F0: (-413 < -397)}
                                                            _ = v32
                                                        }
                                                        switch "a" {
                                                        case "foo bar":
                                                            v1 += v1
                                                        case "a":
                                                            v3 = (^(v3))
                                                        case "\n\t":
                                                            fmt.Println(fi9)
                                                            var v33 int8 = int8(v1)
                                                            if false {
                                                                v34 := GenS7{F0: (-(v3)), F1: (v1 + v1)}
                                                                switch fi11 {
                                                                case 2147483647:
                                                                    v10 = append(v10, (fi11 - 2147483647))
                                                                case -377:
                                                                    var v35 int8 = (v3 >> ((genFn1(("hello" + "a"), 1)) & 63))
                                                                    _ = v35
                                                                case -421:
                                                                    delete(v2, fi9)
                                                                default:
                                                                    if len(v10) > 0 {
                                                                        v10[0] = 9223372036854775807
                                                                    }
                                                                }
                                                                v1 -= v1
                                                                if ((2 > 42) && true) {
                                                                    fmt.Println(fi11, fi11)
                                                                    var v36 uint64 = 29
                                                                    _ = v36
                                                                }
                                                                _ = v34
                                                            } else {
                                                                _ = len(v10)
                                                            }
                                                            _ = v33
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(fi9, fi9, v3)
                                                    v37 := GenS8{F0: v3, F1: (^(v3))}
                                                    var v38 uint16 = v1
                                                    _ = v37
                                                    _ = v38
                                                }
                                            default:
                                                v14 -= (fi11 / ((int(v1)) | 1))
                                            }
                                            _ = v14
                                        }
                                    } else {
                                        v39 := GenS9{F0: (v1 & v1), F1: 0, F2: -128}
                                        v40 := GenS10{F0: 4294967295}
                                        if true {
                                            v3 += v3
                                            v3 += 13
                                        }
                                        _ = v39
                                        _ = v40
                                    }
                                    v3 += (v3 * v3)
                                } else {
                                    switch "" {
                                    case "":
                                        var v41 int16 = 42
                                        v42 := GenS7{F0: v3, F1: 255}
                                        var v43 int32 = 330
                                        v2[fi9] = "a"
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                    case "hello":
                                        switch "a" {
                                        case "a":
                                            v10 = append(v10, (fi9 ^ fi11))
                                        default:
                                            var v44 float64 = float64(0.0)
                                            _ = v44
                                        }
                                    default:
                                        fmt.Println(v1, fi11)
                                    }
                                }
                                var v45 float64 = -2.25
                                _ = fi11
                                _ = v45
                            }
                            v46 := GenS11{F0: false, F1: true, F2: true}
                            if (2147483647 != 378) {
                                v3 = (v3 >> ((173) & 63))
                            }
                            _ = v46
                        } else {
                            v1 += 65535
                        }
                        _ = v10
                    }
                    v1 += (v1 & v1)
                }
            case "a":
                for fi47 := 0; fi47 < 18; fi47++ {
                    var v48 string = "foo bar"
                    switch (fi9 + (fi9 ^ 42)) {
                    case 819:
                        v1 += 32767
                    case 19:
                        delete(v2, fi9)
                    }
                    _ = fi47
                    _ = v48
                }
            default:
                var v49 float32 = ((0.0 + 0.0) - -2.25)
                _ = v49
            }
            _ = fi9
        }
    case "a":
        if (-32768 != -32768) {
            if true {
                switch 32768 {
                case -653:
                    v50 := []int{}
                    _ = v50
                case 287:
                    v51 := []int{128, 9223372036854775807}
                    _ = v51
                case -851:
                    fmt.Println(v1, v3)
                }
                if ((2 < 2) || false) {
                    delete(v2, 487)
                    v52 := v2[42]
                    v3 -= 6
                    _ = v52
                }
            }
        }
    case "":
        v3 -= (-(v3))
    default:
        v3 += v3
    }
    var v53 float64 = (1.5 / -2.25)
    if (1000 >= 1000) {
        if ((0.0 - (0.0 - 0.1)) < (1.5 / (-2.25 * (1e10 / float32(v53))))) {
            v2[9223372036854775807] = "a"
        }
    } else {
        fmt.Println(v53, v3)
        v3 += 42
    }
    for fi54 := 0; fi54 < 2; fi54++ {
        fmt.Println(fi54)
        _ = fi54
    }
    var v55 int16 = 32767
    if (v3 > v3) {
        var v56 bool = ((18446744073709551615 > 679) || ((0.0 / v53) <= 0.1))
        if (genFn1("a", 255) > 0) {
            if true {
                v1 = (v1 & v1)
                fmt.Println(v3, v1)
            } else {
                var v57 int64 = int64(v55)
                var v58 int8 = (v3 << ((4294967295) & 63))
                v59 := GenS9{F0: 1, F1: v55, F2: v58}
                _ = v57
                _ = v58
                _ = v59
            }
        }
        _ = v56
    }
    var v60 uint8 = 255
    var v61 float64 = v53
    fmt.Println(v55, v3, v61)
    var v62 bool = (true || false)
    v55 += (v55 << ((595) & 63))
    if (v3 < (v3 >> ((42) & 63))) {
        v55 = (-(v55))
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v55
    _ = v60
    _ = v61
    _ = v62
}
