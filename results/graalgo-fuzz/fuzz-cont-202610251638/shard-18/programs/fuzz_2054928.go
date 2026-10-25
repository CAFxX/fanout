package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint }
type GenS3 struct { F0 uint8; F1 int16 }
func genFn0() uint16 {
    var v6 float64 = 0.1
    v7 := GenS1{F0: 128}
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 uint8) uint32 {
    var v8 float32 = 3.14159
    v8 += float32(3.14159)
    var v9 uint16 = 84
    switch "a" {
    case "\n\t":
        switch 42 {
        case 9223372036854775807:
            var v10 float64 = (0.1 * float64((float32((-2.25 - 1e10)) + v8)))
            _ = v10
        default:
            v9 += genFn0()
        }
    case "hello":
        var v11 uint16 = v9
        _ = v11
    }
    fmt.Println(v8, v9)
    _ = p0
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 int16 = 999
    v2 := make(map[uint32]bool)
    var v3 string = ("a" + "\n\t")
    var v4 uint = 345
    var v5 bool = (!true)
    fmt.Println(v1, v1, v3)
    v12 := GenS1{F0: 255}
    v3 = v3
    fmt.Println(v3)
    v13 := []int{-290, 128, 1, 2147483647}
    if false {
        if v5 {
            if v5 {
                var v14 int64 = 927
                for fi15 := 0; fi15 < 13; fi15++ {
                    v3 = v3
                    _ = fi15
                }
                var v16 int8 = -12
                _ = v14
                _ = v16
            } else {
                for fi17 := 0; fi17 < 3; fi17++ {
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v1 -= (^(v1))
                        _ = fi18
                    }
                    v5 = (255 < 1)
                    for fi19 := 0; fi19 < 1; fi19++ {
                        v4 += (v4 / (((^(v4))) | 1))
                        v4 = 255
                        v20 := GenS1{F0: 59}
                        switch (fi19 + (fi17 / ((fi19) | 1))) {
                        case 4294967295:
                            fmt.Println(v1, fi19)
                        case 128:
                            switch "a" {
                            case "foo bar":
                                var v21 int8 = 42
                                _ = v21
                            default:
                                fmt.Println(v3, v5)
                            }
                        default:
                            if v5 {
                                v4 = 105
                                v3 = v3
                                var v22 float64 = 0.0
                                v22 = v22
                                _ = v22
                            } else {
                                var v23 string = "foo bar"
                                v2[2] = v5
                                switch v3 {
                                case "\n\t":
                                    if len(v13) > 0 {
                                        v13[0] = -551
                                    }
                                case "":
                                    v24 := GenS2{F0: 0, F1: -9223372036854775808, F2: (v4 - v4)}
                                    _ = v24
                                default:
                                    fmt.Println(v3, v23, fi17)
                                }
                                _ = v23
                            }
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = fi17
                }
            }
            v25 := v2[uint32(v4)]
            _ = v25
        } else {
            var v26 int16 = (-(v1))
            fmt.Println(v1, v4)
            switch "a" {
            case "a":
                for fi27 := 0; fi27 < 2; fi27++ {
                    if (((v4 ^ v4) != (^(v4))) || ((v4 >> ((v4) & 63)) != v4)) {
                        fmt.Println(fi27)
                        var v28 uint32 = 800
                        fmt.Println(v5, v1)
                        _ = v28
                    } else {
                        fmt.Println(v1)
                        v5 = (!(307 != 4294967295))
                        _ = len(v13)
                    }
                    _ = fi27
                }
            case "hello":
                if (88 == 4294967295) {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        var v30 int16 = v26
                        _ = fi29
                        _ = v30
                    }
                }
            case "foo bar":
                var v31 uint16 = genFn0()
                _ = v31
            }
            _ = v26
        }
        v5 = (-9223372036854775808 == 32767)
    }
    v3 = v3
    var v32 int32 = -837
    var v33 int8 = 61
    v34 := GenS3{F0: 1, F1: v1}
    v1 += (v1 + (v1 % ((v1) | 1)))
    for fi35 := 0; fi35 < 12; fi35++ {
        v33 += v33
        _ = fi35
    }
    v13 = append(v13, 987)
    switch "foo bar" {
    case "a":
        v36 := GenS1{F0: 255}
        _ = v36
    case "\n\t":
        var v37 uint = (v4 + 369)
        _ = v37
    case "foo bar":
        v32 += (v32 ^ (v32 % (((^(v32))) | 1)))
    default:
        var v38 string = v3
        _ = v38
    }
    if (81 != 0) {
        var v39 uint = v4
        _ = v39
    }
    if (v5 || (!v5)) {
        if v5 {
            v40 := GenS1{F0: 42}
            v3 = ((("" + "foo bar") + ("hello" + v3)) + v3)
            var v41 string = "hello"
            _ = v40
            _ = v41
        } else {
            fmt.Println(v1, v3)
        }
        var v42 int64 = 42
        var v43 int16 = v1
        _ = v42
        _ = v43
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
