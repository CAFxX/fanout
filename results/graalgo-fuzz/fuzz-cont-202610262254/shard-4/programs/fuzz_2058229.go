package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int; F1 uint }
func genFn0() int8 {
    if false {
        v7 := []int{}
        switch "" {
        case "":
            switch -9223372036854775808 {
            case 985:
                switch -1 {
                case -2147483648:
                    v8 := GenS1{F0: 0, F1: 2}
                    _ = v8
                }
            case 32767:
                var v9 uint8 = 2
                _ = v9
            case 127:
                var v10 uint = 2
                _ = v10
            default:
                v11 := GenS1{F0: -2147483648, F1: 283}
                _ = v11
            }
        case "a":
            if true {
                v7 = append(v7, -9223372036854775808)
            }
        }
        var v12 float32 = float32((0.1 * 3.14159))
        _ = v12
        _ = v7
    }
    v13 := GenS2{F0: 32767, F1: 1, F2: 2}
    var v14 uint = 773
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint = (v14 + 976)
        _ = fi15
        _ = v16
    }
    _ = v13
    _ = v14
    return 88
}
func genFn1(p0 int32, p1 int, p2 bool) int16 {
    fmt.Println(p2)
    if p2 {
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 uint = 215
            p1 -= 4294967295
            _ = fi17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -32768
}

func main() {
    var v1 uint16 = 65535
    var v2 int8 = -1
    var v3 float32 = 0.0
    var v4 int8 = v2
    var v5 uint32 = 42
    v6 := []uint32{0, 2147483647}
    v19 := GenS1{F0: 1, F1: v5}
    var v20 int = 128
    v5 += 2
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v23 := GenS3{F0: 18446744073709551615, F1: "\n\t"}
            if (((0.1 * float64(float32((0.0 / float64(v3))))) + 1.5) != 0.0) {
                if true {
                    if true {
                        v4 = (v2 << ((2) & 63))
                    }
                } else {
                    switch 42 {
                    case 9223372036854775807:
                        fmt.Println(v20)
                    case -9223372036854775808:
                        if false {
                            var v24 float64 = (1.5 + float64((((((v3 / v3) * ((v3 - float32((1e10 / 1.0))) - v3)) / v3) / -2.25) / 1.5)))
                            fmt.Println(v20, v4)
                            _ = v24
                        }
                    case -1:
                        if false {
                            if (false || false) {
                                var v25 int16 = 1
                                v1 += (v1 - v1)
                                var v26 uint8 = 161
                                _ = v25
                                _ = v26
                            } else {
                                fmt.Println(v1)
                                v27 := GenS3{F0: 108, F1: (("a" + "\n\t") + "a")}
                                if ((fi22 | (fi21 & 2147483647)) != (v20 / (((fi21 - fi21)) | 1))) {
                                    v4 -= genFn0()
                                    var v28 string = "hello"
                                    _ = v28
                                }
                                _ = v27
                            }
                            if (!false) {
                                _ = len(v6)
                                v2 = genFn0()
                            }
                        }
                    default:
                        v29 := GenS3{F0: 0, F1: "hello"}
                        _ = v29
                    }
                    var v30 int16 = int16(v20)
                    var v31 int = (v20 * fi21)
                    _ = v30
                    _ = v31
                }
                v32 := GenS4{F0: (v20 + v20), F1: 370}
                if len(v6) > 0 {
                    v6[0] = (v5 - v5)
                }
                _ = v32
            } else {
                fmt.Println(v1)
                var v33 int8 = genFn0()
                _ = v33
            }
            _ = fi22
            _ = v23
        }
        for fi34 := 0; fi34 < 19; fi34++ {
            v20 = 4294967295
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            var v35 bool = (918 > 4294967295)
            _ = fi34
            _ = v35
        }
        var v36 int32 = 1
        _ = fi21
        _ = v36
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        var v38 uint16 = v1
        _ = fi37
        _ = v38
    }
    v20 = -367
    var v39 uint64 = 2
    var v40 float32 = v3
    if len(v6) > 0 {
        v6[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
