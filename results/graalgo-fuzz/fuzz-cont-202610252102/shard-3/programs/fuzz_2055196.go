package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int32 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint; F1 string; F2 string }
type GenS4 struct { F0 bool; F1 bool; F2 int }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint16; F1 bool; F2 uint }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 bool; F1 uint64 }
func genFn0(p0 int32, p1 string, p2 uint) int16 {
    var v6 int = 9223372036854775807
    fmt.Println(p0)
    var v7 uint16 = 960
    if false {
        v8 := GenS1{F0: "", F1: (p0 | p0), F2: p0}
        _ = v8
    } else {
        var v9 uint64 = 42
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int8 = 34
            switch p1 {
            case "":
                var v12 int8 = (v11 * (v11 * v11))
                _ = v12
            case "a":
                switch ((p1 + p1) + (("a" + (p1 + (p1 + "hello"))) + (p1 + "\n\t"))) {
                case "\n\t":
                    p1 = p1
                case "a":
                    if false {
                        v9 -= (v9 >> (((^(p2))) & 63))
                        p1 = p1
                    }
                }
            default:
                fmt.Println(v11, p0, v6)
            }
            _ = fi10
            _ = v11
        }
        if false {
            var v13 uint16 = uint16(p2)
            _ = v13
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                switch (v6 * v6) {
                case 32767:
                    v15 := []int{fi14, 992}
                    _ = v15
                case 128:
                    var v16 uint = (p2 & (p2 ^ 4294967295))
                    _ = v16
                }
                var v17 int16 = int16(p0)
                var v18 int8 = -128
                _ = fi14
                _ = v17
                _ = v18
            }
            var v19 string = p1
            _ = v19
        }
        _ = v9
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(p2, p2, p1)
        switch p1 {
        case "hello":
            var v21 int32 = 42
            _ = v21
        }
        switch fi20 {
        case 42:
            v22 := []int{(fi20 >> (((p2 >> ((1) & 63))) & 63)), (fi20 % ((4294967295) | 1)), (-(v6)), (fi20 >> ((p2) & 63))}
            _ = v22
        }
        _ = fi20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -490
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[uint32]int)
    var v3 float64 = 0.0
    v4 := make(map[uint32]bool)
    var v5 bool = (1e10 < 3.14159)
    v23 := []int{127, 9223372036854775807, int(v1), 668}
    if ((0.0 + ((float32(-2.25) - 0.0) - 1e10)) <= -2.25) {
        v3 += v3
        var v24 uint32 = 933
        fmt.Println(v1, v1, v1)
        _ = v24
    } else {
        fmt.Println(v1, v3)
        var v25 uint64 = v1
        if (42 == 9223372036854775807) {
            v1 = (v25 - (v1 ^ 176))
            for fi26 := 0; fi26 < 18; fi26++ {
                switch "" {
                case "a":
                    if true {
                        for fi27 := 0; fi27 < 11; fi27++ {
                            if ((0.1 - v3) <= 1.5) {
                                v28 := GenS1{F0: "\n\t", F1: int32(v1), F2: 16}
                                v29 := GenS2{F0: 9223372036854775807, F1: 824}
                                _ = v28
                                _ = v29
                            }
                            _ = fi27
                        }
                        v30 := GenS3{F0: 2, F1: "", F2: "\n\t"}
                        _ = v30
                    }
                case "hello":
                    for fi31 := 0; fi31 < 6; fi31++ {
                        var v32 float64 = 0.1
                        var v33 int8 = 100
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                case "foo bar":
                    v34 := GenS1{F0: "foo bar", F1: -912, F2: -1}
                    _ = v34
                default:
                    fmt.Println(fi26, v25)
                }
                var v35 uint16 = 32767
                v36 := GenS2{F0: -185, F1: 792}
                _ = fi26
                _ = v35
                _ = v36
            }
            v1 = 4294967295
        }
        _ = v25
    }
    if false {
        var v37 uint8 = 127
        _ = v37
    } else {
        for fi38 := 0; fi38 < 9; fi38++ {
            v39 := GenS4{F0: (!true), F1: (127 < 2), F2: 32768}
            if v5 {
                v40 := GenS5{F0: -1}
                for fi41 := 0; fi41 < 19; fi41++ {
                    var v42 uint32 = 662
                    v43 := GenS6{F0: -32768}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                var v44 int16 = -255
                _ = v40
                _ = v44
            } else {
                v5 = v5
                v45 := GenS2{F0: 616, F1: 4294967295}
                _ = v45
            }
            _ = fi38
            _ = v39
        }
        v3 = v3
        v46 := GenS7{F0: 152}
        _ = v46
    }
    v47 := GenS5{F0: 1}
    var v48 int8 = -1
    v49 := GenS8{F0: 1}
    v4[1] = (!(v5 || (65535 == 0)))
    if (-1 != 42) {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        if (1 > 125) {
            switch 4294967295 {
            case 4294967295:
                v50 := GenS9{F0: 42, F1: v5, F2: 1}
                _ = v50
            case 1:
                var v51 uint16 = 42
                _ = v51
            case 285:
                v4[382] = true
            default:
                var v52 uint32 = 2147483648
                _ = v52
            }
        } else {
            var v53 string = ((("foo bar" + ("" + "hello")) + (((("foo bar" + "\n\t") + "") + ("hello" + ("a" + "a"))) + "\n\t")) + "\n\t")
            switch 32768 {
            case 32768:
                v1 -= (v1 % (((v1 & 2)) | 1))
            case 1:
                var v54 uint16 = 2
                _ = v54
            case 366:
                switch v53 {
                case "a":
                    v53 = "hello"
                default:
                    if (2147483647 <= 0) {
                        if (-2147483648 != 42) {
                            for fi55 := 0; fi55 < 20; fi55++ {
                                var v56 uint64 = (v1 * v1)
                                var v57 uint16 = 42
                                _ = fi55
                                _ = v56
                                _ = v57
                            }
                            var v58 float32 = 1e10
                            v53 = v53
                            _ = v58
                        } else {
                            for fi59 := 0; fi59 < 18; fi59++ {
                                v4[627] = (0 <= 619)
                                v1 = (v1 >> ((65535) & 63))
                                switch (v53 + "hello") {
                                case "hello":
                                    for fi60 := 0; fi60 < 2; fi60++ {
                                        fmt.Println(v3)
                                        v61 := GenS6{F0: -157}
                                        _ = fi60
                                        _ = v61
                                    }
                                }
                                _ = fi59
                            }
                            fmt.Println(v1)
                        }
                        for fi62 := 0; fi62 < 5; fi62++ {
                            var v63 int16 = genFn0(1, (v53 + "hello"), 255)
                            v64 := GenS7{F0: 0}
                            v5 = (1 < 2)
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    }
                }
            default:
                v65 := GenS5{F0: -2147483648}
                _ = v65
            }
            if v5 {
                if false {
                    var v66 int64 = 0
                    if ((((127 < 4294967295) || false) || true) || (-69 > int8(v1))) {
                        var v67 int16 = genFn0(0, "\n\t", 255)
                        _ = len(v2)
                        _ = v67
                    }
                    _ = v66
                } else {
                    v3 = (-2.25 / ((float64(float32(v3)) - (0.0 - 1.5)) * (v3 + v3)))
                }
            } else {
                for fi68 := 0; fi68 < 10; fi68++ {
                    switch v53 {
                    case "a":
                        for fi69 := 0; fi69 < 15; fi69++ {
                            _ = len(v23)
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            _ = fi69
                        }
                    case "hello":
                        v23 = append(v23, fi68)
                    }
                    var v70 int64 = 1
                    _ = fi68
                    _ = v70
                }
                for fi71 := 0; fi71 < 7; fi71++ {
                    if (!(!(("foo bar" + "hello") <= v53))) {
                        for fi72 := 0; fi72 < 11; fi72++ {
                            v53 = v53
                            v73 := GenS10{F0: (fi71 + fi71), F1: v1}
                            v1 -= v1
                            _ = fi72
                            _ = v73
                        }
                    } else {
                        fmt.Println(fi71, v3, v3)
                        var v74 int64 = 42
                        var v75 float32 = ((0.0 / (float32((float64(1e10) - (3.14159 * 0.1))) * float32(v3))) - 3.14159)
                        _ = v74
                        _ = v75
                    }
                    if (fi71 == -715) {
                        for fi76 := 0; fi76 < 11; fi76++ {
                            var v77 uint32 = 4294967295
                            if ("" != ("\n\t" + v53)) {
                                v4[4294967295] = ((-(v48)) != v48)
                                if len(v23) > 0 {
                                    v23[0] = fi71
                                }
                            }
                            v78 := GenS11{F0: ((v77 % (((v77 * (v77 - v77))) | 1)) <= (^(v77))), F1: v1}
                            _ = fi76
                            _ = v77
                            _ = v78
                        }
                        var v79 bool = (237 < 231)
                        _ = len(v2)
                        _ = v79
                    } else {
                        var v80 int8 = -128
                        for fi81 := 0; fi81 < 4; fi81++ {
                            if (!((42 < 0) && ((1 != 42) || v5))) {
                                v82 := GenS10{F0: fi71, F1: 2}
                                var v83 float32 = 0.1
                                fmt.Println(fi81, v48)
                                _ = v82
                                _ = v83
                            }
                            _ = fi81
                        }
                        _ = v80
                    }
                    _ = len(v2)
                    _ = fi71
                }
                v5 = (2 <= 1)
            }
            _ = v53
        }
    } else {
        if v5 {
            var v84 uint32 = 2
            _ = v84
        } else {
            v85 := GenS4{F0: v5, F1: (v1 < v1), F2: -161}
            _ = len(v23)
            switch "" {
            case "\n\t":
                if (1 > 451) {
                    fmt.Println(v1, v1)
                    v1 = 2
                }
            case "a":
                switch -2147483648 {
                case 42:
                    var v86 int32 = 818
                    _ = v86
                case 2147483647:
                    v87 := GenS11{F0: (4294967295 != 255), F1: v1}
                    _ = v87
                case 32767:
                    var v88 int8 = (v48 >> ((1) & 63))
                    _ = v88
                default:
                    var v89 uint32 = 42
                    _ = v89
                }
            }
            _ = v85
        }
    }
    v90 := GenS11{F0: v5, F1: (^(v1))}
    var v91 bool = v5
    for fi92 := 0; fi92 < 4; fi92++ {
        v91 = (!true)
        _ = fi92
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v90
    _ = v91
}
