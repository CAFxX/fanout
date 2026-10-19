package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 bool; F2 int }
type GenS10 struct { F0 uint }
func genFn0(p0 uint8, p1 uint8) int32 {
    v4 := GenS1{F0: 0, F1: -1, F2: 446}
    var v5 uint64 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -2147483648
}
func genFn1(p0 uint16, p1 uint, p2 bool) uint {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 uint8 = 1
        var v8 float32 = 0.0
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint = p1
    if (42 < 2147483647) {
        for fi10 := 0; fi10 < 5; fi10++ {
            if p2 {
                var v11 uint16 = 32767
                if (!((float32(0.0) * 1.5) > (1.5 - 3.14159))) {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v11 = 244
                        _ = fi12
                    }
                    if true {
                        var v13 int16 = 758
                        switch 42 {
                        case 163:
                            if (!p2) {
                                var v14 int16 = (v13 + v13)
                                _ = v14
                            } else {
                                if true {
                                    v9 -= (^(p1))
                                    if (-1 < -2147483648) {
                                        v15 := []int{fi10}
                                        _ = v15
                                    } else {
                                        v16 := []int{(-(fi10)), (fi10 / ((fi10) | 1)), 128}
                                        _ = v16
                                    }
                                }
                            }
                        case -461:
                            v13 = (v13 + v13)
                        }
                        v17 := []int{(fi10 / (((fi10 % ((2147483647) | 1))) | 1)), (^(fi10)), (fi10 & (fi10 | fi10)), (fi10 | (fi10 >> ((v9) & 63)))}
                        _ = v13
                        _ = v17
                    }
                    fmt.Println(p0)
                }
                var v18 int32 = 2147483647
                _ = v11
                _ = v18
            } else {
                v19 := GenS2{F0: 1, F1: 1}
                for fi20 := 0; fi20 < 18; fi20++ {
                    v21 := []int{407, 32768}
                    fmt.Println(p0, v9, p1)
                    v22 := GenS2{F0: (fi10 % ((fi20) | 1)), F1: 780}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                p0 = p0
                _ = v19
            }
            var v23 int32 = -2147483648
            switch ((("" + "\n\t") + (("hello" + "\n\t") + ("" + "a"))) + (("foo bar" + (((("\n\t" + "hello") + "foo bar") + ("hello" + ("\n\t" + ("hello" + ("a" + "foo bar"))))) + "")) + "\n\t")) {
            case "":
                v24 := []int{(fi10 + (fi10 << ((949) & 63)))}
                _ = v24
            case "foo bar":
                if p2 {
                    var v25 bool = p2
                    fmt.Println(v23, v25)
                    var v26 uint8 = 2
                    _ = v25
                    _ = v26
                } else {
                    var v27 int64 = 9223372036854775807
                    fmt.Println(p0)
                    _ = v27
                }
            }
            _ = fi10
            _ = v23
        }
        v28 := GenS3{F0: 1}
        var v29 int64 = 396
        _ = v28
        _ = v29
    }
    fmt.Println(p0)
    p2 = (((2147483648 == 65535) && (!false)) && (p2 || (true && p2)))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p1
}

func main() {
    var v1 uint32 = 0
    var v2 string = ("hello" + "foo bar")
    v3 := []string{(v2 + v2), "hello", "foo bar"}
    v30 := GenS3{F0: 103}
    var v31 float32 = float32((0.1 * 0.0))
    var v32 int16 = 1000
    v31 -= float32(-2.25)
    var v33 uint32 = v1
    fmt.Println(v32, v2)
    switch 2147483647 {
    case -2147483648:
        var v34 int16 = 0
        _ = v34
    case 30:
        if false {
            for fi35 := 0; fi35 < 1; fi35++ {
                v36 := GenS3{F0: 128}
                v1 -= 1
                _ = fi35
                _ = v36
            }
            var v37 int32 = -1
            _ = v37
        }
    case 987:
        fmt.Println(v32)
    default:
        v38 := GenS4{F0: (v33 * (v1 << ((937) & 63))), F1: 640, F2: -128}
        _ = v38
    }
    v39 := GenS3{F0: 128}
    var v40 bool = true
    fmt.Println(v2, v2)
    switch 32767 {
    case -1:
        v31 = 3.14159
    default:
        switch 42 {
        case 4294967295:
            for fi41 := 0; fi41 < 3; fi41++ {
                fmt.Println(v33, v33, v33)
                _ = fi41
            }
        case 128:
            var v42 uint8 = 128
            _ = v42
        default:
            v43 := GenS5{F0: v2}
            _ = v43
        }
    }
    v32 -= (v32 >> ((4294967295) & 63))
    v1 -= v33
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    if ((v31 * ((0.1 / (float32(float64((v31 + float32(0.1)))) * 0.0)) + 1.5)) != v31) {
        if (!(42 >= 42)) {
            if len(v3) > 0 {
                v3[0] = v2
            }
            fmt.Println(v32, v40, v32)
            fmt.Println(v32, v2)
        } else {
            v40 = ((!(v31 < 3.14159)) || (1 > (v1 + v1)))
            v3 = append(v3, v2)
            if v40 {
                switch v2 {
                case "":
                    var v44 int = 848
                    _ = v44
                case "foo bar":
                    fmt.Println(v1, v1)
                case "a":
                    if (!true) {
                        switch 32768 {
                        case 314:
                            v45 := GenS6{F0: (v2 == "\n\t"), F1: 1}
                            _ = v45
                        case 139:
                            v46 := GenS6{F0: true, F1: 180}
                            _ = v46
                        }
                        if (!(!v40)) {
                            if ((!true) && (true && v40)) {
                                v47 := GenS7{F0: v33}
                                _ = v47
                            } else {
                                if (("" + "hello") < (v2 + v2)) {
                                    v3 = append(v3, v2)
                                    v3 = append(v3, "")
                                    v2 = v2
                                } else {
                                    if v40 {
                                        var v48 int64 = -1
                                        _ = v48
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = ("" + (v2 + ("foo bar" + v2)))
                                        }
                                        switch -9223372036854775808 {
                                        case -1:
                                            for fi49 := 0; fi49 < 17; fi49++ {
                                                v33 = v1
                                                fmt.Println(fi49, fi49, v1)
                                                var v50 uint32 = 4294967295
                                                _ = fi49
                                                _ = v50
                                            }
                                        case -236:
                                            fmt.Println(v2)
                                        }
                                        if (18446744073709551615 > 461) {
                                            var v51 int = 32768
                                            if true {
                                                var v52 uint16 = 32767
                                                v53 := GenS8{F0: (^(v33))}
                                                _ = v52
                                                _ = v53
                                            }
                                            for fi54 := 0; fi54 < 7; fi54++ {
                                                v55 := GenS9{F0: (false && false), F1: false, F2: (fi54 & (fi54 & fi54))}
                                                if v40 {
                                                    for fi56 := 0; fi56 < 12; fi56++ {
                                                        var v57 float64 = (1e10 / 0.1)
                                                        v3 = append(v3, "a")
                                                        v3 = append(v3, v2)
                                                        _ = fi56
                                                        _ = v57
                                                    }
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(fi54, v1)
                                                    }
                                                }
                                                _ = fi54
                                                _ = v55
                                            }
                                            _ = v51
                                        }
                                    }
                                    if (!v40) {
                                        _ = len(v3)
                                    } else {
                                        v40 = ((v1 + (v33 + (v33 + 0))) > v1)
                                        var v58 string = v2
                                        _ = v58
                                    }
                                }
                                if (!(genFn0(2, 175) >= -763)) {
                                    var v59 int8 = 1
                                    fmt.Println(v1, v2)
                                    _ = v59
                                }
                            }
                            fmt.Println(v1)
                            var v60 int = -2147483648
                            _ = v60
                        }
                    }
                }
                for fi61 := 0; fi61 < 8; fi61++ {
                    v33 += 4294967295
                    v62 := GenS8{F0: v33}
                    var v63 int = fi61
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                v64 := GenS10{F0: 0}
                _ = v64
            } else {
                var v65 string = v2
                switch 128 {
                case 163:
                    var v66 int16 = -501
                    _ = v66
                default:
                    v67 := GenS8{F0: v1}
                    _ = v67
                }
                v32 = (v32 % (((^(v32))) | 1))
                _ = v65
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v39
    _ = v40
}
