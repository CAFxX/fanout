package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int; F1 int; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0() int8 {
    switch 2147483647 {
    case 32767:
        v5 := GenS1{F0: -103, F1: 128}
        _ = v5
    }
    if (-32768 >= 0) {
        var v6 int32 = -448
        v6 = (v6 ^ v6)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            if (390 < 0) {
                v8 := []int{}
                var v9 int16 = -1
                _ = v8
                _ = v9
            }
            v10 := GenS2{F0: "foo bar"}
            _ = fi7
            _ = v10
        }
        var v11 float32 = 0.0
        for fi12 := 0; fi12 < 2; fi12++ {
            v11 -= (float32(float64(1e10)) * (((1e10 + v11) + v11) + v11))
            _ = fi12
        }
        _ = v11
    }
    var v13 uint32 = 508
    v13 = v13
    fmt.Println(v13)
    _ = v13
    return 127
}

func main() {
    var v1 uint16 = 2
    var v2 uint = 755
    v3 := make(map[string]uint32)
    var v4 bool = (!true)
    switch "a" {
    case "":
        for fi14 := 0; fi14 < 11; fi14++ {
            if true {
                if v4 {
                    var v15 float32 = (((0.1 + 0.1) + 0.0) + (-2.25 * -2.25))
                    v16 := []int{(^(fi14)), int(v2)}
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        if (false && v4) {
                            for fi17 := 0; fi17 < 12; fi17++ {
                                if len(v16) > 0 {
                                    v16[0] = -83
                                }
                                _ = len(v3)
                                _ = fi17
                            }
                            _ = len(v16)
                        } else {
                            if v4 {
                                v4 = (1 < -867)
                            }
                            v18 := GenS1{F0: 110, F1: fi14}
                            var v19 string = ""
                            _ = v18
                            _ = v19
                        }
                    default:
                        var v20 float64 = 0.0
                        _ = v20
                    }
                    _ = v15
                    _ = v16
                } else {
                    var v21 bool = false
                    var v22 uint = v2
                    switch (fi14 | fi14) {
                    case -2147483648:
                        for fi23 := 0; fi23 < 12; fi23++ {
                            v22 = 761
                            _ = fi23
                        }
                    default:
                        switch "" {
                        case "a":
                            switch 0 {
                            case -2147483648:
                                var v24 bool = false
                                _ = v24
                            case 2147483647:
                                switch "hello" {
                                case "hello":
                                    var v25 int16 = 32767
                                    _ = v25
                                case "\n\t":
                                    if v4 {
                                        v26 := GenS3{F0: 42, F1: (v22 * (v2 % (((^(v2))) | 1)))}
                                        var v27 int = (fi14 & (fi14 << ((v2) & 63)))
                                        var v28 uint64 = 784
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                case "":
                                    if false {
                                        switch "hello" {
                                        case "foo bar":
                                            v29 := []int{}
                                            _ = v29
                                        case "hello":
                                            fmt.Println(v21, v21)
                                        case "":
                                            var v30 int16 = 0
                                            _ = v30
                                        default:
                                            var v31 bool = (589 <= 539)
                                            _ = v31
                                        }
                                    } else {
                                        var v32 float64 = (0.1 / 0.1)
                                        _ = v32
                                    }
                                }
                            case 42:
                                if (!(0 >= -128)) {
                                    v33 := GenS3{F0: 109, F1: 241}
                                    fmt.Println(v21, v2, v1)
                                    _ = v33
                                } else {
                                    v2 += v22
                                }
                            default:
                                var v34 uint64 = 18446744073709551615
                                _ = v34
                            }
                        case "foo bar":
                            var v35 int8 = -128
                            _ = v35
                        default:
                            var v36 int8 = genFn0()
                            _ = v36
                        }
                    }
                    _ = v21
                    _ = v22
                }
            } else {
                v4 = v4
                if (538 == 411) {
                    v2 = uint(v1)
                    var v37 uint32 = 686
                    v38 := GenS4{F0: fi14, F1: fi14, F2: uint64(v1)}
                    _ = v37
                    _ = v38
                }
            }
            v4 = v4
            v39 := GenS3{F0: genFn0(), F1: v2}
            _ = fi14
            _ = v39
        }
    default:
        var v40 string = ("hello" + "hello")
        _ = v40
    }
    v2 = v2
    v4 = (1 < 81)
    var v41 int = 531
    var v42 int64 = -15
    v1 = (v1 >> (((v2 / ((v2) | 1))) & 63))
    v43 := v3["\n\t"]
    if (uint64(v2) == 67) {
        v41 = v41
        v44 := GenS5{F0: 128, F1: 23}
        _ = v44
    }
    v43 = v43
    v45 := GenS4{F0: 32767, F1: (v41 & 2147483647), F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v45
}
