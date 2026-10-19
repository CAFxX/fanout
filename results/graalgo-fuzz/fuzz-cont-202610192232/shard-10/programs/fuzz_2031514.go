package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8; F1 int64; F2 int32 }

func main() {
    var v1 uint = 42
    v2 := []string{"\n\t", ("foo bar" + "foo bar"), ((("foo bar" + "hello") + ("a" + "hello")) + "hello")}
    var v3 int32 = 1
    var v4 float64 = (3.14159 / 1.0)
    var v5 int64 = -9223372036854775808
    fmt.Println(v5, v1, v5)
    if false {
        fmt.Println(v4)
        for fi6 := 0; fi6 < 17; fi6++ {
            v2 = append(v2, "")
            var v7 uint64 = 42
            var v8 int16 = int16(v3)
            _ = fi6
            _ = v7
            _ = v8
        }
        if (!false) {
            var v9 int64 = (v5 / ((1) | 1))
            var v10 int = -1
            _ = v10
            _ = v9
        } else {
            v4 -= float64(1.5)
        }
    } else {
        v3 = v3
    }
    var v11 int8 = -128
    if (true || true) {
        fmt.Println(v1)
        var v12 uint = v1
        _ = v12
    } else {
        v5 += (-(v5))
    }
    v13 := GenS1{F0: -465, F1: 18446744073709551615, F2: 100}
    var v14 string = ("foo bar" + "hello")
    var v15 uint8 = 0
    var v16 int = 426
    var v17 float64 = float64(0.1)
    for fi18 := 0; fi18 < 15; fi18++ {
        if ((!true) && false) {
            var v19 uint16 = 1
            v17 = v17
            for fi20 := 0; fi20 < 12; fi20++ {
                if false {
                    switch (v16 << ((362) & 63)) {
                    case 752:
                        var v21 uint16 = (v19 / ((v19) | 1))
                        _ = v21
                    case -2147483648:
                        v22 := GenS2{F0: (-(v5)), F1: (0 < 0)}
                        _ = v22
                    case -9223372036854775808:
                        switch 2147483647 {
                        case -2147483648:
                            if (!false) {
                                switch fi18 {
                                case 526:
                                    switch fi18 {
                                    case -9223372036854775808:
                                        var v23 uint16 = (v19 + (v19 ^ 405))
                                        _ = v23
                                    case -120:
                                        v24 := GenS1{F0: 1, F1: 0, F2: (-(v11))}
                                        _ = v24
                                    case 991:
                                        var v25 string = (v14 + ((v14 + v14) + v14))
                                        _ = v25
                                    }
                                case 32768:
                                    v2 = append(v2, ("a" + (v14 + "\n\t")))
                                default:
                                    v1 += v1
                                }
                            } else {
                                v26 := GenS2{F0: (v5 - 42), F1: ((v15 - (v15 % ((v15) | 1))) == 128)}
                                v17 -= 0.1
                                if len(v2) > 0 {
                                    v2[0] = "foo bar"
                                }
                                var v27 uint32 = 65535
                                _ = v26
                                _ = v27
                            }
                        default:
                            fmt.Println(v5)
                        }
                    default:
                        var v28 int16 = 0
                        _ = v28
                    }
                    switch "" {
                    case "hello":
                        v16 = 2147483647
                    }
                } else {
                    v11 = v11
                    v2 = append(v2, v14)
                    if ((v16 >> ((v1) & 63)) <= 32767) {
                        v1 += 969
                        if true {
                            v29 := GenS3{F0: ("foo bar" + ("foo bar" + v14))}
                            fmt.Println(v14, fi20, v1)
                            var v30 uint16 = 65535
                            v31 := GenS3{F0: v14}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if false {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v32 int8 = 100
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v32
                            } else {
                                fmt.Println(v30, fi20, v14)
                                v2 = append(v2, ((v14 + "a") + v14))
                                v2 = append(v2, v14)
                                if true {
                                    var v33 int = (-(v16))
                                    _ = v33
                                }
                                fmt.Println(v3, v15)
                                v19 = (v30 << ((uint(v15)) & 63))
                                switch v16 {
                                case 990:
                                    v34 := GenS2{F0: v5, F1: (v30 < (v19 & v30))}
                                    _ = v34
                                case 2147483647:
                                    v4 = v17
                                case 204:
                                    v30 -= 362
                                }
                                if ((v16 <= v16) && (!(true || (v1 != v1)))) {
                                    var v35 string = v14
                                    v36 := GenS4{F0: (v5 << ((769) & 63)), F1: (!false)}
                                    _ = v35
                                    _ = v36
                                }
                                v37 := GenS2{F0: (v5 >> (((v1 * v1)) & 63)), F1: ((-603 <= 32767) && (true && false))}
                                if len(v2) > 0 {
                                    v2[0] = ("" + "foo bar")
                                }
                                v14 = (v14 + "foo bar")
                                _ = v37
                            }
                            if len(v2) > 0 {
                                v2[0] = ((v14 + v14) + v14)
                            }
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    } else {
                        var v38 uint16 = 1
                        if (((v1 << (((v1 % ((v1) | 1))) & 63)) <= v1) && true) {
                            var v39 uint8 = (v15 * (v15 & v15))
                            v40 := GenS5{F0: 0}
                            _ = v39
                            _ = v40
                        } else {
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                            fmt.Println(v17, fi18)
                        }
                        var v41 string = ("a" + "\n\t")
                        _ = v38
                        _ = v41
                    }
                    if ((false || ((v1 ^ v1) < v1)) && false) {
                        var v42 uint64 = 457
                        var v43 int64 = (^(v5))
                        _ = v42
                        _ = v43
                    } else {
                        v44 := GenS6{F0: (v15 | (v15 & 128))}
                        v19 += uint16(v15)
                        v17 += -2.25
                        _ = v44
                    }
                    var v45 bool = (-1 > -32768)
                    _ = v45
                }
                switch "foo bar" {
                case "foo bar":
                    var v46 int16 = -621
                    _ = v46
                case "":
                    v16 = -9223372036854775808
                }
                v47 := GenS7{F0: v15, F1: 1, F2: v3}
                _ = fi20
                _ = v47
            }
            _ = v19
        }
        var v48 int32 = v3
        _ = len(v2)
        _ = fi18
        _ = v48
    }
    v17 -= float64(((3.14159 / 1.5) + 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
