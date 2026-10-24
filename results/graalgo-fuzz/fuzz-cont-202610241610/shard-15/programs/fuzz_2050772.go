package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int64; F1 int64; F2 uint }
type GenS6 struct { F0 uint32; F1 uint8; F2 string }
type GenS7 struct { F0 string; F1 int64 }
func genFn0(p0 int32) uint {
    v7 := GenS1{F0: 0}
    v8 := GenS2{F0: true, F1: 42}
    p0 -= (p0 + -2147483648)
    _ = p0
    _ = v7
    _ = v8
    return 104
}
func genFn1() int8 {
    var v9 int32 = -2147483648
    v9 = v9
    if (-311 <= 9223372036854775807) {
        v9 += v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        v9 = v9
        v9 += v9
        var v11 string = ("hello" + "hello")
        _ = fi10
        _ = v11
    }
    v9 = 2147483647
    _ = v9
    return -73
}

func main() {
    var v1 string = ("hello" + ("a" + ("\n\t" + "")))
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    var v4 int64 = 9223372036854775807
    v5 := []string{("" + (((("foo bar" + "hello") + v1) + "hello") + "\n\t")), "hello"}
    var v6 uint64 = 4294967295
    v1 = v1
    fmt.Println(v1)
    for fi12 := 0; fi12 < 10; fi12++ {
        if true {
            v4 -= (v4 & (v4 | 792))
            if true {
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := GenS3{F0: 255}
                    _ = fi13
                    _ = v14
                }
            } else {
                for fi15 := 0; fi15 < 1; fi15++ {
                    _ = len(v3)
                    var v16 float64 = float64(float32(1e10))
                    v16 = float64((((0.0 - (-2.25 + 1e10)) + float32(v16)) + (0.0 + -2.25)))
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 8; fi17++ {
                    v18 := GenS2{F0: true, F1: fi17}
                    _ = fi17
                    _ = v18
                }
            }
        } else {
            v19 := GenS3{F0: 1}
            fmt.Println(fi12, v1, fi12)
            _ = v19
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            switch "hello" {
            case "":
                v21 := GenS3{F0: 999}
                _ = v21
            }
            var v22 uint32 = 2147483648
            if (true && (!((-(fi20)) <= 0))) {
                var v23 float32 = ((1e10 + ((-2.25 + float32(3.14159)) / -2.25)) + 1e10)
                var v24 string = "hello"
                if false {
                    v4 = (v4 ^ 9223372036854775807)
                    v22 += (v22 & v22)
                    v25 := GenS3{F0: 151}
                    _ = v25
                } else {
                    if true {
                        v23 = (((1.5 / v23) / v23) - float32(0.1))
                        var v26 string = (v24 + (v1 + v1))
                        _ = v26
                    }
                }
                if false {
                    delete(v2, v22)
                    if true {
                        fmt.Println(v4, v22, v22)
                        var v27 uint16 = 42
                        _ = v27
                    } else {
                        v23 = v23
                        v22 -= v22
                        if (!true) {
                            v28 := GenS1{F0: (^(v22))}
                            v24 = ""
                            _ = v28
                        }
                    }
                    v1 = v1
                    var v29 int8 = -1
                    _ = v29
                } else {
                    if true {
                        fmt.Println(v1)
                        v4 += int64(v6)
                        var v30 int64 = -858
                        _ = v30
                    } else {
                        var v31 int8 = 1
                        v3[fi20] = false
                        _ = v31
                    }
                    var v32 uint = uint(v6)
                    if ((160 >= -492) && false) {
                        if true {
                            if (("\n\t" + v24) >= "") {
                                if false {
                                    if len(v5) > 0 {
                                        v5[0] = v1
                                    }
                                    switch ((v24 + v24) + "\n\t") {
                                    case "hello":
                                        fmt.Println(v23, v23)
                                    case "":
                                        v33 := GenS4{F0: (-9223372036854775808 != 9223372036854775807), F1: v24}
                                        fmt.Println(v22, v4)
                                        v34 := GenS2{F0: false, F1: (fi20 + (fi12 - (fi12 + fi20)))}
                                        _ = v33
                                        _ = v34
                                    case "foo bar":
                                        var v35 uint64 = (v6 ^ 4294967295)
                                        _ = v35
                                    }
                                    var v36 uint = v32
                                    var v37 uint = genFn0(1)
                                    _ = len(v3)
                                    var v38 uint64 = v6
                                    v38 += v6
                                    v39 := GenS2{F0: false, F1: fi12}
                                    var v40 string = v24
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                            } else {
                                v23 -= ((v23 + ((v23 * v23) * 0.1)) - float32(((1.5 * ((float64(0.1) + (float64(float32(float64(v23))) / 1.0)) + (-2.25 - (0.0 / 1.0)))) * 0.1)))
                                switch 616 {
                                case 128:
                                    switch 32768 {
                                    case 0:
                                        v1 = (((v1 + v1) + ("" + (((("hello" + "") + (v1 + "foo bar")) + v1) + v1))) + v1)
                                        if len(v5) > 0 {
                                            v5[0] = "foo bar"
                                        }
                                        switch fi12 {
                                        case 176:
                                            v41 := v3[128]
                                            _ = v41
                                        case -410:
                                            switch ((v24 + v1) + v1) {
                                            case "\n\t":
                                                if (2147483647 <= 2147483647) {
                                                    var v42 bool = true
                                                    if v42 {
                                                        v43 := GenS1{F0: 413}
                                                        fmt.Println(v1, v6, v42)
                                                        _ = v43
                                                    } else {
                                                        var v44 uint32 = (v22 + (v22 | uint32(v32)))
                                                        var v45 float64 = ((0.0 - 1.5) / (((((float64(((float32((((0.1 - 1.5) + (0.1 / (0.1 + (0.0 * -2.25)))) / ((3.14159 * (float64((-2.25 / v23)) / 1e10)) + -2.25))) * (((1e10 + (3.14159 * ((0.1 + 3.14159) + v23))) + 3.14159) + v23)) + (float32(3.14159) - ((v23 / (1.5 / (v23 * 1e10))) - (v23 + v23))))) - 0.1) + float64(v23)) / ((0.1 / (1e10 * 1.5)) - 0.0)) - 0.1) + 0.0))
                                                        v24 = (v24 + "")
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    _ = v42
                                                } else {
                                                    v4 += v4
                                                }
                                            case "foo bar":
                                                v46 := GenS4{F0: ((!(!false)) && false), F1: v24}
                                                _ = v46
                                            case "a":
                                                if true {
                                                    v23 -= v23
                                                    v32 += uint(v22)
                                                    if len(v5) > 0 {
                                                        v5[0] = v1
                                                    }
                                                }
                                            default:
                                                fmt.Println(fi12, v4, v22)
                                            }
                                        case 251:
                                            if (1 == 42) {
                                                var v47 bool = true
                                                var v48 int64 = (v4 << ((v32) & 63))
                                                _ = v47
                                                _ = v48
                                            } else {
                                                v23 = (v23 / v23)
                                                v32 += 65535
                                            }
                                        default:
                                            var v49 uint32 = 2147483647
                                            _ = v49
                                        }
                                    }
                                case 158:
                                    v1 = "a"
                                case 42:
                                    _ = len(v2)
                                }
                                v50 := GenS5{F0: 9223372036854775807, F1: (v4 + 42), F2: v32}
                                _ = v50
                            }
                            var v51 uint16 = 65535
                            fmt.Println(fi12, v24, v24)
                            _ = v51
                        }
                    }
                    _ = v32
                }
                _ = v23
                _ = v24
            }
            _ = fi20
            _ = v22
        }
        var v52 uint32 = 42
        _ = fi12
        _ = v52
    }
    var v53 uint = 18446744073709551615
    switch -9223372036854775808 {
    case 9223372036854775807:
        if true {
            v4 -= v4
            for fi54 := 0; fi54 < 9; fi54++ {
                v55 := GenS4{F0: false, F1: v1}
                _ = fi54
                _ = v55
            }
        } else {
            if false {
                fmt.Println(v6, v4)
            } else {
                v56 := GenS6{F0: 0, F1: 2, F2: v1}
                _ = v56
            }
        }
    case 321:
        if (-1 > 127) {
            fmt.Println(v4, v4)
            if (!(false || (v1 <= v1))) {
                v6 += v6
                switch (((v1 + ("" + v1)) + v1) + v1) {
                case "\n\t":
                    v57 := GenS1{F0: 714}
                    _ = v57
                }
                v58 := GenS7{F0: v1, F1: (v4 / ((-9223372036854775808) | 1))}
                _ = v58
            } else {
                switch 42 {
                case 2147483647:
                    fmt.Println(v6)
                case -849:
                    v6 = v6
                }
                v5 = append(v5, v1)
                v59 := GenS7{F0: v1, F1: (v4 ^ -9223372036854775808)}
                _ = v59
            }
        }
    case 0:
        v53 = 648
    default:
        delete(v3, 406)
    }
    v4 -= 79
    var v60 float32 = (0.1 / (0.1 - 1.5))
    fmt.Println(v4, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v60
}
