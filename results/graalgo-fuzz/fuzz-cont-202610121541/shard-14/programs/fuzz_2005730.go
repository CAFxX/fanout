package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16; F1 int32; F2 int16 }
type GenS4 struct { F0 int16; F1 int64; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 uint8; F2 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16; F1 bool }

func main() {
    var v1 float32 = 1.5
    var v2 bool = false
    v3 := make(map[uint32]string)
    var v4 uint16 = 255
    var v5 bool = (197 <= 105)
    for fi6 := 0; fi6 < 17; fi6++ {
        v1 = ((float32(0.1) * (3.14159 / (0.0 / v1))) * 0.1)
        var v7 float64 = (3.14159 / 1.0)
        _ = fi6
        _ = v7
    }
    if (3.14159 > 1.5) {
        delete(v3, 552)
        switch "a" {
        case "":
            v4 = (v4 + v4)
        default:
            if (4294967295 >= 18446744073709551615) {
                fmt.Println(v2)
            } else {
                var v8 string = ("hello" + "foo bar")
                switch 4294967295 {
                case 42:
                    v9 := []int{}
                    _ = v9
                case 333:
                    v10 := GenS1{F0: -1, F1: 42}
                    _ = v10
                }
                _ = v8
            }
        }
        var v11 uint8 = 255
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v4 += (v4 >> ((1) & 63))
        _ = fi12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        if (true || false) {
            if v5 {
                v5 = (0 == 1)
                v14 := []int{-9223372036854775808, 351, fi13}
                _ = v14
            } else {
                var v15 uint64 = 182
                _ = v15
            }
            v16 := GenS2{F0: 0, F1: 255}
            switch ("hello" + (("" + "hello") + ("foo bar" + "hello"))) {
            case "foo bar":
                switch ("a" + ("hello" + "a")) {
                case "hello":
                    var v17 float32 = (3.14159 + v1)
                    _ = v17
                case "\n\t":
                    if v5 {
                        if false {
                            if false {
                                v18 := GenS1{F0: 42, F1: 809}
                                v5 = v5
                                if true {
                                    v2 = (!v2)
                                    v4 = (v4 << ((42) & 63))
                                } else {
                                    if false {
                                        v19 := []int{(fi13 - 32767), fi13, (fi13 >> ((18446744073709551615) & 63))}
                                        _ = v19
                                    }
                                }
                                _ = v18
                            } else {
                                v2 = v2
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    var v21 int16 = 0
                                    _ = fi20
                                    _ = v21
                                }
                                v4 += (v4 * (v4 & (^(v4))))
                            }
                            for fi22 := 0; fi22 < 8; fi22++ {
                                if v5 {
                                    if v2 {
                                        var v23 uint16 = v4
                                        _ = v23
                                    }
                                    if v2 {
                                        v1 = -2.25
                                        if ((v4 != (^(v4))) || (!((1e10 != (0.1 - -2.25)) && true))) {
                                            fmt.Println(v5)
                                            v24 := GenS3{F0: 1, F1: -1, F2: 32767}
                                            delete(v3, 65535)
                                            if v5 {
                                                switch 1 {
                                                case 9223372036854775807:
                                                    v1 = float32(1.5)
                                                    switch ("foo bar" + ("a" + "hello")) {
                                                    case "":
                                                        var v25 uint32 = 2147483648
                                                        _ = v25
                                                    case "hello":
                                                        var v26 uint8 = 2
                                                        _ = v26
                                                    default:
                                                        var v27 uint32 = 65535
                                                        switch ("\n\t" + ("" + "\n\t")) {
                                                        case "hello":
                                                            v1 += v1
                                                            if ((1.5 / (float64(1.5) / 1.0)) < 0.1) {
                                                                v28 := GenS1{F0: 0, F1: 42}
                                                                var v29 uint16 = v4
                                                                fmt.Println(v29)
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                        }
                                                        _ = v27
                                                    }
                                                    v30 := []int{-2147483648, 899}
                                                    fmt.Println(fi22, fi22, fi13)
                                                    _ = v30
                                                }
                                            } else {
                                                v31 := GenS1{F0: -32768, F1: 0}
                                                if (9223372036854775807 <= -1) {
                                                    v32 := GenS4{F0: 111, F1: 9223372036854775807, F2: 4294967295}
                                                    v2 = ((v2 && v2) || (42 >= 1))
                                                    _ = v32
                                                }
                                                if v2 {
                                                    v33 := GenS2{F0: 497, F1: 46}
                                                    var v34 int = fi22
                                                    _ = v33
                                                    _ = v34
                                                } else {
                                                    var v35 int8 = 42
                                                    _ = v35
                                                }
                                                _ = v31
                                            }
                                            _ = v24
                                        } else {
                                            v5 = v2
                                            var v36 int32 = 2147483647
                                            var v37 uint = 4294967295
                                            v1 = float32((1e10 - 1.5))
                                            _ = v36
                                            _ = v37
                                        }
                                        v38 := []int{}
                                        _ = v38
                                    } else {
                                        v39 := GenS5{F0: fi13}
                                        v40 := []int{fi13, fi13, fi22, fi13}
                                        switch "hello" {
                                        case "foo bar":
                                            var v41 uint = 197
                                            _ = v41
                                        }
                                        _ = v39
                                        _ = v40
                                    }
                                    if v2 {
                                        v4 += 487
                                        v4 -= 2
                                        v4 = 225
                                        if (!true) {
                                            if false {
                                                var v42 string = "hello"
                                                if v2 {
                                                    v4 -= 1
                                                    v3[2147483647] = "\n\t"
                                                    v43 := []int{}
                                                    _ = v43
                                                } else {
                                                    v2 = v5
                                                }
                                                _ = v42
                                            } else {
                                                _ = len(v3)
                                                switch "a" {
                                                case "foo bar":
                                                    switch (fi13 & (fi22 + (fi13 << ((1) & 63)))) {
                                                    case -639:
                                                        var v44 uint16 = v4
                                                        _ = v44
                                                    default:
                                                        v2 = true
                                                    }
                                                case "a":
                                                    v45 := GenS6{F0: 0, F1: 2, F2: 619}
                                                    _ = v45
                                                case "":
                                                    v1 = v1
                                                default:
                                                    if ((fi13 % ((fi22) | 1)) == (fi13 / (((-(fi13))) | 1))) {
                                                        fmt.Println(fi13, v4)
                                                        v1 = v1
                                                        var v46 int16 = 672
                                                        v5 = v5
                                                        var v47 float64 = 1.5
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        v48 := GenS7{F0: (240 >= 56)}
                                                        var v49 int16 = -32768
                                                        fmt.Println(fi22)
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                }
                                            }
                                        }
                                        var v50 int16 = 1000
                                        _ = v50
                                    }
                                    switch 127 {
                                    case -474:
                                        v1 = float32(-2.25)
                                    case 702:
                                        fmt.Println(v1, v5)
                                    case 763:
                                        v4 = 32767
                                    }
                                } else {
                                    v4 = v4
                                    var v51 float32 = (v1 / v1)
                                    var v52 uint32 = 177
                                    _ = v51
                                    _ = v52
                                }
                                _ = fi22
                            }
                        } else {
                            v1 = v1
                            fmt.Println(v4)
                            for fi53 := 0; fi53 < 12; fi53++ {
                                if (0 > 241) {
                                    v54 := []int{(fi13 * 127), (fi13 ^ fi13), (fi53 % ((fi53) | 1)), (fi53 | fi53)}
                                    v4 += v4
                                    _ = v54
                                }
                                _ = fi53
                            }
                        }
                        var v55 int8 = -1
                        _ = v55
                    }
                }
            case "a":
                for fi56 := 0; fi56 < 19; fi56++ {
                    switch fi13 {
                    case 127:
                        v57 := v3[279]
                        _ = v57
                    }
                    v2 = (-9223372036854775808 <= -1)
                    switch 789 {
                    case -2147483648:
                        switch (("hello" + "") + "foo bar") {
                        case "\n\t":
                            v58 := GenS6{F0: -9223372036854775808, F1: 2, F2: 42}
                            _ = v58
                        }
                    case 2147483647:
                        v59 := v3[2147483647]
                        _ = v59
                    }
                    _ = fi56
                }
            case "\n\t":
                if (true || (42 < 118)) {
                    v5 = (1 == 2147483647)
                    var v60 uint32 = 936
                    _ = v60
                }
            }
            _ = v16
        }
        v5 = (((9223372036854775807 < -573) && v2) || (!((!(("\n\t" + "") != "foo bar")) && (!((-1 == 1) || ((3.14159 - 1e10) < (0.0 + float64(((1.5 + float32(float64((3.14159 * 0.1)))) / 1.5)))))))))
        v1 += 1e10
        _ = fi13
    }
    var v61 float32 = float32((0.0 - 3.14159))
    var v62 uint8 = 128
    v63 := GenS8{F0: 801, F1: false}
    v5 = true
    v64 := []int{2147483647, -9223372036854775808, -1, 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v64:", len(v64), v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
