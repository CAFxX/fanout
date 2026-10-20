package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS4 struct { F0 int8; F1 string; F2 int }
func genFn0(p0 uint64) int32 {
    v6 := []int{-722, 1, 32767, 32767}
    for fi7 := 0; fi7 < 2; fi7++ {
        p0 += p0
        var v8 uint16 = 255
        _ = fi7
        _ = v8
    }
    if (0 != 65535) {
        fmt.Println(p0)
    }
    p0 += p0
    v6 = append(v6, 32767)
    _ = p0
    _ = v6
    return 0
}
func genFn1() uint16 {
    var v9 uint = 271
    v10 := GenS1{F0: 0, F1: v9, F2: 18446744073709551615}
    _ = v10
    _ = v9
    return 259
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int = (v1 & (v1 + v1))
    v3 := make(map[string]string)
    var v4 uint8 = 1
    v5 := make(map[uint32]string)
    fmt.Println(v1, v1)
    var v11 uint64 = 18446744073709551615
    v11 = (v11 * (v11 & (v11 | 42)))
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(v4, v2)
            _ = fi13
        }
        fmt.Println(v11, v2)
        _ = fi12
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if (v4 >= (^(v4))) {
            v15 := GenS1{F0: 4294967295, F1: 255, F2: 1}
            if true {
                v16 := GenS1{F0: 42, F1: 787, F2: 433}
                v17 := GenS2{F0: -32768, F1: 255}
                _ = v16
                _ = v17
            }
            _ = v15
        }
        _ = fi14
    }
    var v18 float64 = (0.1 - 1e10)
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 bool = (4294967295 < 1)
        switch v2 {
        case 0:
            if true {
                for fi21 := 0; fi21 < 5; fi21++ {
                    if false {
                        switch (v1 & fi21) {
                        case 127:
                            var v22 float32 = 0.0
                            _ = v22
                        case 42:
                            var v23 int8 = 16
                            _ = v23
                        case 370:
                            v1 = -670
                        }
                    } else {
                        v4 += (v4 & v4)
                        fmt.Println(v18, v2)
                    }
                    switch "" {
                    case "hello":
                        v2 -= fi21
                    case "a":
                        var v24 int32 = 42
                        _ = v24
                    case "\n\t":
                        v25 := GenS3{F0: 112, F1: 162, F2: (v4 / (((^(v4))) | 1))}
                        _ = v25
                    }
                    v11 += (v11 - v11)
                    _ = fi21
                }
            }
        default:
            var v26 float32 = ((0.1 - (-2.25 / 1e10)) / ((((1e10 * 1e10) - 0.1) * float32(v18)) - (0.0 * (3.14159 * 0.1))))
            _ = v26
        }
        _ = fi19
        _ = v20
    }
    if true {
        for fi27 := 0; fi27 < 6; fi27++ {
            var v28 bool = true
            v28 = v28
            v11 = v11
            _ = fi27
            _ = v28
        }
        switch v2 {
        case 0:
            for fi29 := 0; fi29 < 20; fi29++ {
                fmt.Println(fi29, v1, fi29)
                _ = fi29
            }
        case -167:
            v3["hello"] = (("\n\t" + "hello") + "hello")
        case -1:
            for fi30 := 0; fi30 < 11; fi30++ {
                for fi31 := 0; fi31 < 8; fi31++ {
                    v1 = (^(fi31))
                    if (false || ("foo bar" > "foo bar")) {
                        var v32 float64 = ((v18 / v18) * -2.25)
                        _ = v32
                    } else {
                        v33 := v3[""]
                        if false {
                            var v34 float64 = v18
                            v35 := v5[uint32(v4)]
                            _ = v34
                            _ = v35
                        } else {
                            var v36 uint = 18446744073709551615
                            if (!(!(1 > 0))) {
                                v1 = (-(v2))
                                if false {
                                    v37 := []int{(-(fi30)), (fi30 | (v1 & fi31)), 32768, 506}
                                    _ = v37
                                } else {
                                    v38 := []int{(fi30 & fi30), v1, 1}
                                    fmt.Println(v4, v33)
                                    v18 = (0.1 + ((v18 / v18) * (v18 + v18)))
                                    v38 = append(v38, fi30)
                                    fmt.Println(v4, v11)
                                    _ = v38
                                }
                                fmt.Println(v4, v33)
                            } else {
                                switch v2 {
                                case 32767:
                                    var v39 uint8 = 128
                                    _ = v39
                                case 467:
                                    var v40 int = v1
                                    _ = v40
                                default:
                                    v4 += v4
                                }
                                v18 = 1.5
                                var v41 uint32 = 42
                                _ = v41
                            }
                            v11 = 812
                            _ = v36
                        }
                        _ = v33
                    }
                    _ = fi31
                }
                v2 += -926
                v18 += 1e10
                _ = fi30
            }
        }
        _ = len(v5)
    } else {
        var v42 int16 = 1
        var v43 string = ("" + "foo bar")
        _ = v42
        _ = v43
    }
    v44 := GenS4{F0: 1, F1: "a", F2: (v2 >> ((65535) & 63))}
    for fi45 := 0; fi45 < 19; fi45++ {
        var v46 uint16 = 255
        _ = fi45
        _ = v46
    }
    switch ("hello" + (("foo bar" + ("" + "foo bar")) + "foo bar")) {
    case "":
        v47 := v3["a"]
        _ = v47
    default:
        for fi48 := 0; fi48 < 1; fi48++ {
            v49 := GenS4{F0: -1, F1: "\n\t", F2: 128}
            v18 = (-2.25 + (v18 * float64(1.5)))
            v1 = (fi48 - -1)
            _ = fi48
            _ = v49
        }
    }
    switch "" {
    case "foo bar":
        fmt.Println(v11, v2)
    case "":
        v50 := v5[2147483648]
        _ = v50
    case "\n\t":
        delete(v5, 2147483647)
    }
    v51 := []int{}
    v2 = v1
    switch "foo bar" {
    case "foo bar":
        for fi52 := 0; fi52 < 17; fi52++ {
            fmt.Println(fi52)
            for fi53 := 0; fi53 < 3; fi53++ {
                delete(v5, 2147483647)
                var v54 int64 = -1
                fmt.Println(v1)
                fmt.Println(fi52)
                fmt.Println(fi52, v2, v54)
                _ = fi53
                _ = v54
            }
            fmt.Println(v18, v4)
            _ = fi52
        }
    case "\n\t":
        var v55 uint64 = v11
        _ = v55
    default:
        v56 := v5[1]
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v51
}
