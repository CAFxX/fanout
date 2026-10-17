package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 string }
type GenS2 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 int8; F1 int32 }
type GenS7 struct { F0 bool; F1 int32; F2 int }
func genFn0(p0 string, p1 string, p2 int64) string {
    var v5 uint = 0
    var v6 uint64 = 0
    v7 := []int{-2147483648, -1, 241}
    var v8 uint32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 uint16) uint16 {
    var v9 int8 = 42
    var v10 int64 = 812
    var v11 uint = 42
    v12 := []int{42}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return (p0 ^ 32767)
}

func main() {
    var v1 uint16 = 696
    var v2 int32 = -1
    var v3 int32 = v2
    var v4 float32 = (0.0 / 3.14159)
    var v13 int32 = v2
    v14 := []int{}
    for fi15 := 0; fi15 < 14; fi15++ {
        _ = len(v14)
        _ = fi15
    }
    v14 = append(v14, 309)
    switch 2147483647 {
    case 32767:
        var v16 int64 = -505
        _ = v16
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v17 uint32 = uint32(v1)
    v18 := GenS1{F0: -83, F1: v2, F2: "\n\t"}
    for fi19 := 0; fi19 < 2; fi19++ {
        if (1 >= 42) {
            if ("" <= "\n\t") {
                if (!(0 <= 1)) {
                    if len(v14) >= 2 {
                        v14 = v14[1:len(v14)]
                    }
                    v3 = v13
                } else {
                    v17 -= v17
                    if (!false) {
                        var v20 int16 = -963
                        if (!false) {
                            v1 = genFn1(42)
                        }
                        if ((1e10 + float64(float32(float64(float32((1e10 + 0.0)))))) > float64(v4)) {
                            if (511 > 552) {
                                for fi21 := 0; fi21 < 8; fi21++ {
                                    v14 = append(v14, (-(fi21)))
                                    switch ("\n\t" + "") {
                                    case "hello":
                                        if false {
                                            v20 += (v20 + v20)
                                        } else {
                                            v17 = 2147483647
                                        }
                                    default:
                                        switch 1 {
                                        case -1:
                                            var v22 uint16 = v1
                                            _ = v22
                                        case 32768:
                                            v4 += (v4 / float32((0.1 * (-2.25 + (0.1 * (0.1 - ((0.1 * float64(v4)) - 0.1)))))))
                                        default:
                                            var v23 bool = (true && false)
                                            _ = v23
                                        }
                                    }
                                    v2 -= (-(v13))
                                    var v24 int = fi19
                                    _ = fi21
                                    _ = v24
                                }
                            } else {
                                switch ("foo bar" + (("" + "hello") + ("" + ""))) {
                                case "foo bar":
                                    if (!false) {
                                        v14 = append(v14, fi19)
                                        v25 := GenS2{F0: 289, F1: genFn1(v1), F2: int64(fi19)}
                                        _ = v25
                                    } else {
                                        v26 := GenS3{F0: "foo bar"}
                                        fmt.Println(v2, v20, v4)
                                        v27 := GenS2{F0: 0, F1: (v1 << ((714) & 63)), F2: 93}
                                        _ = v26
                                        _ = v27
                                    }
                                case "a":
                                    switch (fi19 + 685) {
                                    case 913:
                                        var v28 int8 = 127
                                        _ = v28
                                    case 0:
                                        v13 += (v3 << ((65535) & 63))
                                    }
                                case "":
                                    var v29 int16 = (-(v20))
                                    _ = v29
                                }
                            }
                            for fi30 := 0; fi30 < 13; fi30++ {
                                v4 += (v4 + v4)
                                var v31 float64 = (1.5 + 3.14159)
                                _ = fi30
                                _ = v31
                            }
                            v4 += (float32(3.14159) - v4)
                        } else {
                            if true {
                                v3 = (v3 + v3)
                            } else {
                                for fi32 := 0; fi32 < 2; fi32++ {
                                    v3 = 603
                                    if false {
                                        v4 += 0.1
                                    }
                                    _ = fi32
                                }
                                v3 -= (v13 / ((v13) | 1))
                            }
                            v2 = v13
                        }
                        _ = v20
                    } else {
                        v33 := GenS2{F0: 0, F1: (v1 % (((v1 << ((1) & 63))) | 1)), F2: -363}
                        _ = v33
                    }
                    v1 = (v1 - 65535)
                }
                fmt.Println(v17)
                var v34 float64 = -2.25
                _ = v34
            }
            if false {
                fmt.Println(v2, v17)
                var v35 float32 = 0.1
                _ = v35
            } else {
                if (21 <= 0) {
                    v36 := GenS3{F0: "\n\t"}
                    _ = v36
                } else {
                    v37 := GenS1{F0: 194, F1: (v13 >> ((987) & 63)), F2: ""}
                    v1 -= (^(v1))
                    v4 = (((-2.25 - v4) - ((float32(0.0) + v4) + (v4 + float32(1.5)))) * 3.14159)
                    _ = v37
                }
                v38 := GenS4{F0: (!(!true))}
                v2 += (v13 / ((106) | 1))
                _ = v38
            }
            v4 = (float32((float64(float32((0.0 + 3.14159))) / -2.25)) + (3.14159 + float32((3.14159 - float64((v4 + (v4 / (float32(1.5) - 3.14159))))))))
        }
        _ = fi19
    }
    switch "" {
    case "\n\t":
        fmt.Println(v3)
    case "hello":
        switch "" {
        case "foo bar":
            v39 := GenS5{F0: -618, F1: 32767}
            _ = v39
        case "\n\t":
            var v40 uint64 = 1
            _ = v40
        case "":
            v41 := GenS4{F0: (73 < 100)}
            _ = v41
        default:
            v13 -= 42
        }
    case "foo bar":
        v42 := GenS3{F0: "a"}
        _ = v42
    default:
        v43 := GenS6{F0: 47, F1: (v3 ^ v3)}
        _ = v43
    }
    var v44 int32 = v13
    var v45 int64 = 1
    var v46 string = ("" + "")
    v47 := GenS4{F0: (false || false)}
    v48 := GenS7{F0: true, F1: (v2 ^ v44), F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
}
