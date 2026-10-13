package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS7 struct { F0 uint32 }
func genFn0() int32 {
    var v7 uint64 = 42
    var v8 float32 = (3.14159 * -2.25)
    for fi9 := 0; fi9 < 13; fi9++ {
        if false {
            v7 = v7
            v10 := GenS1{F0: v7}
            v11 := []int{fi9, -515, (fi9 | fi9), (fi9 * fi9)}
            _ = v10
            _ = v11
        } else {
            v12 := GenS2{F0: false}
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int64 = 9223372036854775807
                v15 := []int{fi13, fi9, fi13}
                v14 -= (v14 / ((-9223372036854775808) | 1))
                _ = fi13
                _ = v14
                _ = v15
            }
            fmt.Println(v7, v8)
            _ = v12
        }
        _ = fi9
    }
    var v16 int32 = -306
    if false {
        var v17 int8 = -1
        _ = v17
    }
    _ = v16
    _ = v7
    _ = v8
    return v16
}
func genFn1(p0 uint, p1 string, p2 uint64) uint16 {
    v18 := GenS3{F0: ((p2 == p2) && (1 == 1)), F1: (!false)}
    switch ("" + p1) {
    case "\n\t":
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 int16 = 1
            _ = fi19
            _ = v20
        }
    case "hello":
        v21 := GenS2{F0: ((p0 - p0) >= p0)}
        _ = v21
    default:
        fmt.Println(p2)
    }
    fmt.Println(p1, p2, p1)
    v22 := []int{-520}
    if false {
        p0 = 269
        var v23 int = -1
        p2 = 18446744073709551615
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v22
    return 323
}

func main() {
    var v1 uint64 = 0
    v2 := []uint32{2, 200, 989}
    var v3 float32 = -2.25
    var v4 string = ("hello" + "a")
    var v5 int16 = 42
    var v6 uint8 = 134
    switch -9223372036854775808 {
    case 2147483647:
        var v24 int = 4294967295
        _ = v24
    case 988:
        switch (("foo bar" + v4) + (v4 + ((("" + ("a" + v4)) + ("foo bar" + "foo bar")) + "hello"))) {
        case "":
            var v25 float32 = -2.25
            _ = v25
        case "foo bar":
            v4 = (v4 + "\n\t")
        }
    default:
        v26 := GenS4{F0: 42, F1: false}
        _ = v26
    }
    fmt.Println(v5, v3, v6)
    if true {
        v1 = 171
        switch "foo bar" {
        case "foo bar":
            v2 = append(v2, 938)
        }
    } else {
        if ((v3 + v3) != 0.0) {
            v2 = append(v2, 42)
            fmt.Println(v3, v5)
            switch -759 {
            case 4294967295:
                if (65535 <= 124) {
                    var v27 uint32 = 2147483647
                    for fi28 := 0; fi28 < 20; fi28++ {
                        v6 = v6
                        _ = fi28
                    }
                    if (true || false) {
                        if len(v2) > 0 {
                            v2[0] = v27
                        }
                        v5 += (v5 - (v5 | (-(v5))))
                        v4 = (("a" + (v4 + v4)) + ("a" + "foo bar"))
                    }
                    _ = v27
                }
            case 0:
                fmt.Println(v5)
            }
        } else {
            var v29 int32 = -987
            v3 += (v3 / -2.25)
            v29 -= (v29 >> ((513) & 63))
            _ = v29
        }
        for fi30 := 0; fi30 < 7; fi30++ {
            if true {
                if (73 <= 4294967295) {
                    if (false || false) {
                        var v31 float64 = ((-2.25 - (0.0 + 3.14159)) - 3.14159)
                        for fi32 := 0; fi32 < 12; fi32++ {
                            v6 = v6
                            if false {
                                var v33 int32 = 42
                                fmt.Println(v5)
                                _ = v33
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                                v6 = v6
                            }
                            _ = fi32
                        }
                        _ = v31
                    } else {
                        v2 = append(v2, 65535)
                    }
                    if (v4 != (v4 + v4)) {
                        for fi34 := 0; fi34 < 3; fi34++ {
                            v1 = (v1 >> ((42) & 63))
                            _ = fi34
                        }
                        v4 = v4
                        var v35 uint32 = 1
                        _ = v35
                    }
                }
                fmt.Println(v3)
                if len(v2) > 0 {
                    v2[0] = 2
                }
            } else {
                v36 := GenS3{F0: true, F1: ((-1 >= 1) && (-2.25 == (1.5 - (3.14159 + ((0.1 / -2.25) - (-2.25 - 1e10))))))}
                var v37 bool = false
                v3 -= (0.0 / ((3.14159 * (float32((1.5 + -2.25)) - (v3 + 0.1))) * 0.0))
                _ = v36
                _ = v37
            }
            _ = fi30
        }
    }
    v6 -= 128
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 float32 = v3
        _ = fi38
        _ = v39
    }
    v3 += (-2.25 - (v3 - (float32((1e10 / 3.14159)) * (-2.25 / (((float32((3.14159 + 0.0)) - (0.1 + v3)) / (v3 * float32(((3.14159 + 1.5) / ((1.5 + 0.0) + (1e10 - (float64(float32(0.0)) + 1e10))))))) + v3)))))
    if len(v2) > 0 {
        v2[0] = 1
    }
    var v40 int = -1
    switch "a" {
    case "\n\t":
        v3 += (((v3 + v3) - 0.0) / 3.14159)
    default:
        switch (v40 * v40) {
        case 681:
            if false {
                for fi41 := 0; fi41 < 7; fi41++ {
                    switch v4 {
                    case "foo bar":
                        v40 -= fi41
                    case "a":
                        v42 := GenS4{F0: uint16(v1), F1: (42 > 0)}
                        _ = v42
                    case "hello":
                        if true {
                            _ = len(v2)
                            var v43 float64 = (0.0 * (1.5 * 1.5))
                            _ = v43
                        } else {
                            var v44 int8 = 100
                            for fi45 := 0; fi45 < 19; fi45++ {
                                v2 = append(v2, 349)
                                fmt.Println(v40, fi41)
                                _ = fi45
                            }
                            _ = v44
                        }
                    }
                    _ = fi41
                }
                v46 := GenS5{F0: ((true && false) || true), F1: 1}
                _ = v46
            } else {
                v6 = v6
                var v47 int16 = 528
                _ = v47
            }
        }
    }
    _ = len(v2)
    fmt.Println(v4)
    for fi48 := 0; fi48 < 12; fi48++ {
        switch "\n\t" {
        case "":
            switch fi48 {
            case 0:
                var v49 uint64 = (v1 << ((2) & 63))
                _ = v49
            case -9223372036854775808:
                fmt.Println(fi48, v4, v4)
            default:
                if len(v2) > 0 {
                    v2[0] = 0
                }
            }
        case "\n\t":
            var v50 uint64 = (^(v1))
            _ = v50
        case "a":
            var v51 uint8 = 2
            _ = v51
        default:
            v52 := GenS6{F0: 42, F1: v1, F2: genFn0()}
            _ = v52
        }
        if (!false) {
            var v53 uint = 1
            _ = v53
        }
        for fi54 := 0; fi54 < 3; fi54++ {
            v6 -= v6
            _ = fi54
        }
        _ = fi48
    }
    v55 := GenS3{F0: true, F1: (false && (0 >= 42))}
    for fi56 := 0; fi56 < 8; fi56++ {
        var v57 int16 = v5
        _ = fi56
        _ = v57
    }
    var v58 int8 = 100
    for fi59 := 0; fi59 < 16; fi59++ {
        if (!true) {
            v60 := GenS6{F0: 0, F1: 4294967295, F2: genFn0()}
            _ = v60
        } else {
            var v61 string = v4
            for fi62 := 0; fi62 < 15; fi62++ {
                v63 := GenS1{F0: v1}
                v64 := GenS7{F0: 319}
                fmt.Println(v61)
                _ = fi62
                _ = v63
                _ = v64
            }
            switch "a" {
            case "a":
                _ = len(v2)
            case "":
                var v65 uint8 = v6
                _ = v65
            case "hello":
                var v66 int8 = 1
                _ = v66
            default:
                var v67 string = "\n\t"
                _ = v67
            }
            _ = v61
        }
        _ = fi59
    }
    if (false || false) {
        v4 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v55
    _ = v58
    _ = v6
}
