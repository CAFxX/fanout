package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 bool; F1 int8; F2 uint16 }
func genFn0() uint {
    if (65535 <= 42) {
        var v5 uint16 = 0
        v5 -= v5
        _ = v5
    } else {
        var v6 uint = 18446744073709551615
        _ = v6
    }
    var v7 float64 = 0.1
    v8 := GenS1{F0: 9223372036854775807, F1: "foo bar"}
    _ = v7
    _ = v8
    return 636
}

func main() {
    var v1 uint32 = 2147483648
    var v2 string = ("foo bar" + "\n\t")
    var v3 uint16 = 1
    v4 := make(map[string]string)
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int32 = -56
        _ = fi9
        _ = v10
    }
    v2 = (((v2 + v2) + v2) + "foo bar")
    var v11 int = -42
    v4[v2] = v2
    if true {
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                fmt.Println(v11, fi12, v3)
                v11 -= 32768
                v1 += v1
                if true {
                    v14 := GenS1{F0: 0, F1: "\n\t"}
                    v4[v2] = ("foo bar" + v2)
                    v15 := v4[(v2 + "hello")]
                    if ((2147483647 != 1) || ((v15 >= ((v2 + ("" + "")) + v15)) && false)) {
                        fmt.Println(v15)
                        if (-615 < -518) {
                            if (!(false || (v3 <= 0))) {
                                v16 := GenS1{F0: 0, F1: v15}
                                v11 += 827
                                v15 = (v2 + v2)
                                _ = v16
                            }
                            v3 -= 255
                            var v17 int64 = -9223372036854775808
                            _ = v17
                        } else {
                            fmt.Println(v15, v2, fi12)
                            v11 = v11
                            v2 = ("" + v15)
                        }
                        var v18 uint = 2
                        _ = v18
                    } else {
                        v19 := GenS2{F0: 616, F1: 4294967295, F2: uint8(v1)}
                        _ = v19
                    }
                    _ = len(v4)
                    _ = v14
                    _ = v15
                }
                var v20 float64 = 0.0
                _ = fi13
                _ = v20
            }
            for fi21 := 0; fi21 < 11; fi21++ {
                fmt.Println(v11, fi21)
                switch v2 {
                case "":
                    var v22 int = v11
                    _ = v22
                case "a":
                    if true {
                        v3 -= (^(v3))
                        fmt.Println(v2, v2)
                        switch 814 {
                        case 1:
                            v23 := []int{v11, -599}
                            _ = v23
                        }
                    }
                case "foo bar":
                    v2 = v2
                default:
                    v24 := GenS3{F0: int64(v1)}
                    _ = v24
                }
                switch (((v2 + v2) + "hello") + ("" + "hello")) {
                case "a":
                    fmt.Println(v1)
                case "":
                    v25 := []int{-645}
                    _ = v25
                case "foo bar":
                    switch fi12 {
                    case 0:
                        var v26 uint32 = 65535
                        _ = v26
                    case 32768:
                        if (!(-684 >= -1)) {
                            var v27 int8 = -128
                            _ = v27
                        } else {
                            v4[v2] = ""
                        }
                    }
                default:
                    v28 := GenS1{F0: 9223372036854775807, F1: v2}
                    _ = v28
                }
                _ = fi21
            }
            _ = fi12
        }
        v29 := GenS3{F0: -1}
        if false {
            if (-2147483648 <= -2147483648) {
                _ = len(v4)
                v30 := GenS4{F0: 9223372036854775807}
                v1 -= (v1 & 0)
                _ = v30
            }
            v31 := []int{}
            _ = v31
        } else {
            if false {
                if false {
                    v3 = v3
                    if false {
                        for fi32 := 0; fi32 < 13; fi32++ {
                            v11 = fi32
                            _ = fi32
                        }
                        if (42 < 868) {
                            v2 = v2
                            v33 := []int{(v11 - (-(v11))), v11}
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 9; fi34++ {
                            v35 := GenS4{F0: -749}
                            var v36 uint = 18446744073709551615
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    } else {
                        if (!false) {
                            switch (v11 / (((v11 * -891)) | 1)) {
                            case -2147483648:
                                var v37 int32 = -1
                                _ = v37
                            case 4294967295:
                                for fi38 := 0; fi38 < 14; fi38++ {
                                    v1 -= (^(v1))
                                    var v39 float32 = ((0.0 - (1.5 + float32((1e10 - 0.0)))) - -2.25)
                                    _ = fi38
                                    _ = v39
                                }
                            case 2147483647:
                                var v40 float64 = (float64(-2.25) + (3.14159 + 1e10))
                                _ = v40
                            default:
                                v41 := GenS5{F0: (v2 + v2), F1: -1}
                                _ = v41
                            }
                        }
                        for fi42 := 0; fi42 < 18; fi42++ {
                            v2 = "\n\t"
                            for fi43 := 0; fi43 < 12; fi43++ {
                                v44 := GenS1{F0: 0, F1: "a"}
                                var v45 uint = 1
                                v1 = 2
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                            _ = fi42
                        }
                        v11 = (v11 << ((genFn0()) & 63))
                    }
                }
            }
            var v46 int16 = -1
            var v47 int64 = 1
            _ = v46
            _ = v47
        }
        _ = v29
    }
    if (true && (false && (!(2147483647 == 1)))) {
        if false {
            v48 := GenS6{F0: 32767, F1: -195}
            if true {
                var v49 float64 = 1e10
                fmt.Println(v49, v2, v49)
                _ = v49
            } else {
                var v50 int = (v11 & v11)
                for fi51 := 0; fi51 < 17; fi51++ {
                    for fi52 := 0; fi52 < 3; fi52++ {
                        v1 -= 392
                        _ = fi52
                    }
                    _ = fi51
                }
                _ = v50
            }
            var v53 int = int(v3)
            _ = v48
            _ = v53
        }
        v54 := GenS4{F0: -183}
        var v55 float64 = 0.0
        _ = v54
        _ = v55
    } else {
        v1 += 38
        if true {
            if (v11 != (-(v11))) {
                fmt.Println(v3)
            }
        } else {
            switch (v11 >> ((612) & 63)) {
            case 0:
                for fi56 := 0; fi56 < 10; fi56++ {
                    v1 += v1
                    v11 += fi56
                    v3 = (v3 << ((65535) & 63))
                    _ = fi56
                }
            case -2147483648:
                v11 -= v11
            case -1:
                var v57 int8 = 100
                _ = v57
            default:
                v58 := GenS7{F0: 255, F1: v1}
                _ = v58
            }
        }
        var v59 string = v2
        _ = v59
    }
    v60 := []int{v11, v11, (v11 >> ((genFn0()) & 63)), v11}
    for fi61 := 0; fi61 < 18; fi61++ {
        if ((1 >= 0) || ((v1 - (v1 ^ 2147483648)) < v1)) {
            v3 = (v3 / ((v3) | 1))
            if false {
                var v62 bool = false
                v4["\n\t"] = v2
                v63 := GenS8{F0: (v62 && (v62 || (((false && true) || (v62 && false)) || v62))), F1: 0, F2: (v3 ^ v3)}
                _ = v62
                _ = v63
            } else {
                var v64 int8 = 127
                v11 += -1
                var v65 int32 = 0
                _ = v64
                _ = v65
            }
        }
        if false {
            v66 := GenS1{F0: 42, F1: v2}
            for fi67 := 0; fi67 < 4; fi67++ {
                v11 = fi67
                _ = fi67
            }
            _ = v66
        }
        _ = fi61
    }
    v68 := GenS7{F0: genFn0(), F1: v1}
    v2 = (((v2 + "foo bar") + "hello") + v2)
    for fi69 := 0; fi69 < 9; fi69++ {
        v70 := GenS5{F0: v2, F1: 1}
        _ = fi69
        _ = v70
    }
    switch 127 {
    case 9223372036854775807:
        for fi71 := 0; fi71 < 13; fi71++ {
            if len(v60) >= 2 {
                v60 = v60[1:len(v60)]
            }
            var v72 uint8 = 1
            var v73 uint32 = v1
            _ = fi71
            _ = v72
            _ = v73
        }
    }
    v11 = -9223372036854775808
    _ = len(v60)
    var v74 int64 = 9223372036854775807
    var v75 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v68
    _ = v74
    _ = v75
}
