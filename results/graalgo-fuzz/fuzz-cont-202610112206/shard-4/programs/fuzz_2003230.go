package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 bool; F2 int16 }
type GenS6 struct { F0 int16; F1 string; F2 uint }
func genFn0(p0 uint16) int {
    v6 := []int{}
    if (p0 == (p0 / ((p0) | 1))) {
        fmt.Println(p0)
    }
    _ = p0
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 int8 = 0
    var v2 float32 = (1e10 - 0.1)
    var v3 int64 = 778
    var v4 int64 = (v3 & v3)
    var v5 uint32 = 2147483647
    var v7 uint64 = 0
    v8 := GenS1{F0: v4, F1: (v4 & v4), F2: 1}
    v1 = (v1 % (((v1 ^ v1)) | 1))
    v3 += v4
    var v9 uint = 0
    fmt.Println(v2)
    switch "foo bar" {
    case "hello":
        for fi10 := 0; fi10 < 17; fi10++ {
            if (1 != -1) {
                v11 := GenS2{F0: "a"}
                _ = v11
            }
            if false {
                v12 := GenS1{F0: v3, F1: -447, F2: 947}
                v13 := []int{128, int(v1)}
                if true {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v15 := GenS2{F0: "\n\t"}
                        if (true && true) {
                            var v16 int8 = v1
                            v17 := GenS3{F0: (v16 * v1)}
                            _ = v16
                            _ = v17
                        } else {
                            var v18 uint8 = 42
                            var v19 string = "a"
                            _ = v18
                            _ = v19
                        }
                        v7 -= v7
                        _ = fi14
                        _ = v15
                    }
                    for fi20 := 0; fi20 < 20; fi20++ {
                        v9 -= v9
                        v7 = (v7 - (v7 ^ (^(v7))))
                        _ = fi20
                    }
                    for fi21 := 0; fi21 < 17; fi21++ {
                        if ((v9 ^ v9) <= (v9 ^ 4294967295)) {
                            if true {
                                var v22 uint32 = (v5 >> ((2) & 63))
                                v22 += 348
                                switch 42 {
                                case -852:
                                    switch -262 {
                                    case 32768:
                                        v13 = append(v13, fi10)
                                    case 127:
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                    default:
                                        fmt.Println(v5)
                                        if len(v13) > 0 {
                                            v13[0] = 42
                                        }
                                        v5 -= (v5 ^ v22)
                                    }
                                case 694:
                                    v13 = append(v13, (fi10 / ((32767) | 1)))
                                case 216:
                                    var v23 uint8 = uint8(v7)
                                    _ = v23
                                default:
                                    var v24 int64 = (v3 << ((v9) & 63))
                                    _ = v24
                                }
                                _ = v22
                            } else {
                                switch fi21 {
                                case 127:
                                    var v25 int16 = 32767
                                    _ = v25
                                case 4294967295:
                                    v4 -= -9223372036854775808
                                case -385:
                                    v26 := GenS3{F0: -119}
                                    _ = v26
                                default:
                                    if len(v13) > 0 {
                                        v13[0] = 32768
                                    }
                                }
                            }
                            v27 := GenS4{F0: (^(v7))}
                            v1 += v1
                            fmt.Println(v7, v9)
                            _ = v27
                        } else {
                            v2 -= (3.14159 / 1.0)
                            v28 := GenS5{F0: (v1 / ((v1) | 1)), F1: true, F2: 42}
                            fmt.Println(v9, v2)
                            _ = v28
                        }
                        var v29 float64 = 3.14159
                        _ = fi21
                        _ = v29
                    }
                }
                _ = v12
                _ = v13
            } else {
                v3 = (v4 >> ((4294967295) & 63))
                v5 = v5
                v30 := GenS6{F0: int16(v7), F1: "a", F2: (v9 / ((v9) | 1))}
                _ = v30
            }
            v3 = v3
            _ = fi10
        }
    case "":
        switch "\n\t" {
        case "\n\t":
            v31 := []int{-650, -42}
            _ = v31
        default:
            v5 = (v5 & v5)
        }
    case "\n\t":
        var v32 uint64 = v7
        _ = v32
    default:
        fmt.Println(v5, v7)
    }
    v4 = 438
    v5 = v5
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 int64 = (-(v3))
        v35 := GenS6{F0: 1, F1: "a", F2: v9}
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 float64 = ((0.1 + 1.5) - ((-2.25 * float64((-2.25 / 3.14159))) - 1e10))
    v37 := []int{genFn0(255), -706, genFn0(255), 9223372036854775807}
    v38 := GenS1{F0: 0, F1: v4, F2: 1}
    var v39 float32 = (0.0 + -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
