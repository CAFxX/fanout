package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint16; F1 string; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint32; F1 uint16; F2 uint }
type GenS8 struct { F0 int32; F1 int8 }
type GenS9 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int32, p2 string) uint {
    v5 := []int{42, -55, 32767}
    var v6 int = 96
    var v7 uint16 = 414
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 736
}
func genFn1(p0 uint, p1 uint8, p2 uint64) int64 {
    if false {
        p1 = p1
    }
    v8 := []int{-9223372036854775808, 32767, int(p0)}
    p0 = genFn0(743, -8, "")
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 627
    v3 := make(map[int]int)
    var v4 int8 = 100
    v9 := v3[2147483647]
    v10 := []int{(v9 & (v9 * 135)), (v9 | 0)}
    _ = len(v1)
    switch (v9 ^ (v9 / (((^(v9))) | 1))) {
    case -9223372036854775808:
        switch "hello" {
        case "hello":
            switch -338 {
            case -9223372036854775808:
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 string = ("" + ("hello" + ("foo bar" + (("" + "\n\t") + "a"))))
                    _ = fi11
                    _ = v12
                }
            case 9223372036854775807:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            case 1:
                if true {
                    v9 -= (v9 % (((v9 + v9)) | 1))
                    if false {
                        v13 := GenS1{F0: true}
                        _ = v13
                    }
                    if (uint16(v2) > 255) {
                        v4 -= (v4 - (-(v4)))
                        switch (("foo bar" + "a") + ("a" + "")) {
                        case "a":
                            v14 := v1[(v9 + v9)]
                            _ = v14
                        default:
                            v15 := GenS1{F0: ((v2 == v2) && false)}
                            _ = v15
                        }
                    }
                }
            }
        }
    case 4294967295:
        switch "hello" {
        case "foo bar":
            v9 -= (v9 ^ (v9 - (^(v9))))
        default:
            switch (v9 + v9) {
            case 32768:
                v9 = v9
            case 9223372036854775807:
                v16 := GenS2{F0: v9}
                _ = v16
            }
        }
    }
    switch (v9 - (^(v9))) {
    case -778:
        var v17 int64 = -296
        _ = v17
    default:
        if false {
            var v18 uint16 = 702
            var v19 uint8 = 127
            _ = v18
            _ = v19
        } else {
            switch "\n\t" {
            case "a":
                var v20 bool = (false && true)
                _ = v20
            case "":
                v2 -= (v2 * v2)
            case "\n\t":
                if len(v10) > 0 {
                    v10[0] = v9
                }
            default:
                v2 -= v2
            }
            for fi21 := 0; fi21 < 8; fi21++ {
                fmt.Println(v4, v9)
                _ = fi21
            }
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v4, v9)
            case "":
                fmt.Println(v4, v9, v2)
            case "foo bar":
                v2 = 0
            }
        }
    }
    switch (v9 >> ((255) & 63)) {
    case 953:
        v22 := GenS3{F0: v9, F1: 42, F2: 1000}
        _ = v22
    case 4294967295:
        if (!false) {
            switch 919 {
            case 128:
                switch (("a" + "foo bar") + "") {
                case "\n\t":
                    if (6 < 185) {
                        fmt.Println(v4)
                        v4 = (v4 % (((v4 | 82)) | 1))
                        if (false && (255 == 909)) {
                            var v23 bool = true
                            delete(v1, (v9 + (v9 + v9)))
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v4, v23, v23)
                            case "a":
                                fmt.Println(v23, v9, v4)
                            }
                            _ = v23
                        }
                    }
                case "hello":
                    for fi24 := 0; fi24 < 3; fi24++ {
                        switch "foo bar" {
                        case "foo bar":
                            var v25 int8 = v4
                            _ = v25
                        case "a":
                            delete(v1, v9)
                        case "":
                            v9 += 32768
                        }
                        switch "" {
                        case "":
                            if false {
                                fmt.Println(v4, fi24)
                                var v26 string = (("foo bar" + ("" + (("foo bar" + "a") + "hello"))) + "a")
                                _ = v26
                            }
                        case "\n\t":
                            for fi27 := 0; fi27 < 17; fi27++ {
                                v28 := GenS4{F0: 32767, F1: "", F2: (v2 + v2)}
                                _ = fi27
                                _ = v28
                            }
                        default:
                            if (true && ((!(2147483647 == -270)) && ((((((("a" + "\n\t") + "hello") + ("hello" + "\n\t")) + "hello") >= "hello") && ("a" >= (("\n\t" + ("\n\t" + "")) + "hello"))) && true))) {
                                if (0.1 < (((1e10 / 1.0) - (3.14159 * 1.5)) * 1e10)) {
                                    delete(v3, (^(fi24)))
                                } else {
                                    v9 = (fi24 * v9)
                                    for fi29 := 0; fi29 < 8; fi29++ {
                                        var v30 int = fi24
                                        v30 -= fi29
                                        _ = fi29
                                        _ = v30
                                    }
                                    if (!(!false)) {
                                        v31 := GenS5{F0: 544, F1: 255}
                                        _ = v31
                                    } else {
                                        for fi32 := 0; fi32 < 20; fi32++ {
                                            if true {
                                                v2 += v2
                                            } else {
                                                var v33 int32 = 2147483647
                                                _ = v33
                                            }
                                            v2 += v2
                                            if false {
                                                var v34 bool = (!true)
                                                _ = v34
                                            } else {
                                                if (true || false) {
                                                    v2 -= v2
                                                    v35 := GenS1{F0: false}
                                                    v36 := GenS5{F0: -239, F1: 2}
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    fmt.Println(v2)
                                                }
                                                var v37 uint32 = (v2 - 2)
                                                v10 = append(v10, 2147483647)
                                                _ = v37
                                            }
                                            if true {
                                                var v38 float64 = ((float64((3.14159 + float32((1.5 * 1.5)))) / 3.14159) * (1e10 - 3.14159))
                                                v10 = append(v10, 1)
                                                _ = v38
                                            } else {
                                                v39 := GenS6{F0: 0, F1: (v4 % (((v4 ^ v4)) | 1))}
                                                v4 += v4
                                                _ = v39
                                            }
                                            _ = len(v10)
                                            _ = fi32
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v9, v4)
                                v4 += -62
                            }
                        }
                        var v40 int8 = v4
                        _ = fi24
                        _ = v40
                    }
                }
            case 0:
                v4 -= v4
            case 752:
                _ = len(v1)
            default:
                if ((1 >= 1000) && false) {
                    if (!((0.0 - 3.14159) <= (float64(1.5) - 3.14159))) {
                        for fi41 := 0; fi41 < 17; fi41++ {
                            var v42 float64 = (1.5 - 1.5)
                            var v43 uint = 255
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 19; fi44++ {
                            switch "hello" {
                            case "a":
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            case "":
                                for fi45 := 0; fi45 < 13; fi45++ {
                                    var v46 int16 = 1
                                    if (204 >= 0) {
                                        v4 = v4
                                        v10 = append(v10, -358)
                                        if true {
                                            v9 = -450
                                            v47 := GenS4{F0: 255, F1: "foo bar", F2: v2}
                                            if false {
                                                switch ((("a" + ("a" + ("a" + ""))) + ("foo bar" + "hello")) + ("foo bar" + "")) {
                                                case "":
                                                    var v48 int16 = (v46 - int16(v4))
                                                    _ = v48
                                                case "a":
                                                    var v49 float32 = (float32((-2.25 * -2.25)) + (3.14159 + 0.0))
                                                    _ = v49
                                                case "foo bar":
                                                    v4 = 125
                                                default:
                                                    v50 := GenS7{F0: 366, F1: 388, F2: 789}
                                                    _ = v50
                                                }
                                                var v51 uint = genFn0(42, -1, "hello")
                                                var v52 float64 = (3.14159 - -2.25)
                                                _ = v51
                                                _ = v52
                                            }
                                            _ = v47
                                        }
                                    }
                                    var v53 int8 = v4
                                    _ = fi45
                                    _ = v46
                                    _ = v53
                                }
                            case "\n\t":
                                for fi54 := 0; fi54 < 18; fi54++ {
                                    v55 := GenS7{F0: (^(v2)), F1: 32767, F2: genFn0(2147483648, 350, "foo bar")}
                                    v56 := GenS1{F0: false}
                                    _ = fi54
                                    _ = v55
                                    _ = v56
                                }
                            default:
                                for fi57 := 0; fi57 < 4; fi57++ {
                                    var v58 uint32 = v2
                                    v59 := v1[(fi44 | (v9 ^ int(v4)))]
                                    v60 := GenS8{F0: -2147483648, F1: 49}
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                            }
                            v9 = (v9 << ((42) & 63))
                            var v61 uint = 4294967295
                            _ = fi44
                            _ = v61
                        }
                        v62 := GenS3{F0: v9, F1: 255, F2: int16(v9)}
                        _ = v62
                    }
                    var v63 int = (-(v9))
                    _ = v63
                }
            }
        }
    default:
        var v64 int64 = -9223372036854775808
        _ = v64
    }
    v65 := GenS9{F0: v2}
    if len(v10) > 0 {
        v10[0] = (v9 ^ (v9 ^ v9))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v9
}
