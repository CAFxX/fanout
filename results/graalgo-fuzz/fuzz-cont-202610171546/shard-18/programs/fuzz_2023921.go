package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 int32; F2 uint8 }
func genFn0(p0 uint32) bool {
    var v5 uint = 255
    if false {
        p0 = p0
        p0 -= p0
        v5 = 65535
    } else {
        var v6 float64 = (((float64(3.14159) * (3.14159 / 1e10)) * 0.1) / (1e10 / (-2.25 - 3.14159)))
        _ = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(fi7)
        v8 := []int{(fi7 << (((v5 - 1)) & 63)), (fi7 - fi7), (fi7 % ((9223372036854775807) | 1)), (-(fi7))}
        _ = fi7
        _ = v8
    }
    var v9 int8 = 0
    _ = p0
    _ = v5
    _ = v9
    return (false && false)
}
func genFn1(p0 uint8, p1 int8, p2 int32) int16 {
    p1 = (^(p1))
    var v10 bool = (0 > 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -115
}

func main() {
    var v1 uint16 = 42
    var v2 uint = 2
    var v3 int = 0
    var v4 int = v3
    v11 := []int{-990, v3, (v4 | v3), v4}
    if (1 != 122) {
        if (false && true) {
            switch "" {
            case "a":
                if true {
                    v12 := GenS1{F0: (42 >= 408), F1: 1}
                    v13 := GenS2{F0: 2147483647, F1: (-(v3))}
                    _ = v12
                    _ = v13
                } else {
                    var v14 string = "hello"
                    if false {
                        for fi15 := 0; fi15 < 7; fi15++ {
                            var v16 int64 = 0
                            for fi17 := 0; fi17 < 6; fi17++ {
                                v18 := GenS3{F0: 1, F1: 2, F2: 806}
                                _ = fi17
                                _ = v18
                            }
                            _ = fi15
                            _ = v16
                        }
                        for fi19 := 0; fi19 < 17; fi19++ {
                            v14 = v14
                            switch 2147483647 {
                            case 103:
                                for fi20 := 0; fi20 < 1; fi20++ {
                                    if false {
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                    }
                                    v2 += (v2 % ((v2) | 1))
                                    v21 := GenS4{F0: "hello"}
                                    _ = fi20
                                    _ = v21
                                }
                            case 32767:
                                switch ("a" + (((v14 + "hello") + v14) + v14)) {
                                case "\n\t":
                                    for fi22 := 0; fi22 < 4; fi22++ {
                                        switch (^(v3)) {
                                        case 653:
                                            switch (((v14 + v14) + "a") + (v14 + ((v14 + "") + "hello"))) {
                                            case "foo bar":
                                                v14 = ""
                                                v11 = append(v11, (fi22 * (^(fi22))))
                                            default:
                                                switch (fi19 + (fi19 - int(v2))) {
                                                case 4294967295:
                                                    if true {
                                                        fmt.Println(v1, fi19)
                                                    }
                                                case -227:
                                                    v2 = (v2 / ((v2) | 1))
                                                }
                                            }
                                        case -787:
                                            v11 = append(v11, (fi22 | fi19))
                                        }
                                        if true {
                                            v23 := GenS4{F0: v14}
                                            _ = v23
                                        } else {
                                            v2 += 42
                                            var v24 string = "\n\t"
                                            switch (v3 % ((fi19) | 1)) {
                                            case 9223372036854775807:
                                                var v25 uint = v2
                                                _ = v25
                                            }
                                            _ = v24
                                        }
                                        _ = fi22
                                    }
                                case "a":
                                    var v26 uint32 = 2147483647
                                    _ = v26
                                }
                            case -9223372036854775808:
                                var v27 uint = (^(v2))
                                _ = v27
                            }
                            _ = fi19
                        }
                    } else {
                        v2 = (v2 % ((v2) | 1))
                        v28 := GenS1{F0: true, F1: 255}
                        v1 = 32767
                        _ = v28
                    }
                    v14 = v14
                    _ = v14
                }
            default:
                var v29 uint8 = uint8(v2)
                _ = v29
            }
        } else {
            v30 := GenS2{F0: -484, F1: (-(v4))}
            v3 -= int(v1)
            _ = v30
        }
        var v31 int64 = 1
        _ = v31
    }
    v32 := GenS1{F0: true, F1: 255}
    if (-125 <= 1) {
        v33 := GenS2{F0: -446, F1: v4}
        v11 = append(v11, -1)
        _ = v33
    } else {
        v34 := GenS4{F0: "\n\t"}
        switch ("foo bar" + "a") {
        case "foo bar":
            v3 += v3
        default:
            fmt.Println(v4, v3, v2)
        }
        v35 := GenS5{F0: 363, F1: 0, F2: 127}
        _ = v34
        _ = v35
    }
    var v36 int = v4
    var v37 uint32 = 0
    if (false || (!((v1 & (v1 << ((v2) & 63))) <= (v1 + (v1 / (((v1 | uint16(v37))) | 1)))))) {
        if len(v11) > 0 {
            v11[0] = (v4 + 127)
        }
    } else {
        fmt.Println(v37)
    }
    v37 = (v37 | (^(v37)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
}
