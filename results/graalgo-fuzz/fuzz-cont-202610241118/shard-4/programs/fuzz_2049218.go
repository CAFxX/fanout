package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 string) uint64 {
    var v5 uint64 = 2
    var v6 int = 42
    _ = p0
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 uint64, p1 uint, p2 uint16) int32 {
    var v7 float32 = 1e10
    fmt.Println(p2, v7, p0)
    v8 := GenS1{F0: 42, F1: genFn0((("foo bar" + ("a" + (("\n\t" + "") + "a"))) + "\n\t")), F2: 128}
    if (!false) {
        if false {
            for fi9 := 0; fi9 < 3; fi9++ {
                for fi10 := 0; fi10 < 9; fi10++ {
                    p1 = p1
                    var v11 uint = (p1 >> ((p1) & 63))
                    var v12 int16 = 32767
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
        } else {
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS1{F0: 42, F1: p0, F2: uint8(p1)}
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 20; fi15++ {
                p2 -= (p2 + (p2 ^ (p2 + p2)))
                _ = fi15
            }
        }
    } else {
        v16 := GenS1{F0: 91, F1: p0, F2: 2}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -130
}

func main() {
    var v1 float32 = (1.5 - (0.0 - (3.14159 - float32(3.14159))))
    var v2 float64 = -2.25
    v3 := make(map[int]string)
    v4 := []uint32{1, 65535, 2}
    var v17 int64 = 9223372036854775807
    for fi18 := 0; fi18 < 12; fi18++ {
        switch "hello" {
        case "":
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 int16 = -417
                if (false && true) {
                    v2 = (v2 - (-2.25 / (v2 * 1.5)))
                    v17 = (v17 >> ((2) & 63))
                    fmt.Println(v17)
                    fmt.Println(v17, v20)
                } else {
                    v21 := GenS1{F0: -128, F1: 2, F2: 142}
                    v22 := GenS1{F0: 100, F1: 812, F2: 42}
                    if (!true) {
                        v17 = v17
                        switch ("a" + "foo bar") {
                        case "\n\t":
                            switch fi19 {
                            case -999:
                                v23 := GenS1{F0: 83, F1: 0, F2: 2}
                                _ = v23
                            case -2147483648:
                                v1 += (((0.1 / v1) - (v1 * float32(-2.25))) / v1)
                            default:
                                var v24 float64 = v2
                                _ = v24
                            }
                        case "":
                            if true {
                                v4 = append(v4, 4294967295)
                            }
                        }
                        _ = len(v3)
                    } else {
                        var v25 float32 = (3.14159 - v1)
                        if false {
                            v20 = v20
                            if (18446744073709551615 <= genFn0("hello")) {
                                if false {
                                    if (0 == genFn0(("" + ""))) {
                                        var v26 float64 = 1e10
                                        _ = v26
                                    }
                                }
                                v27 := v3[32767]
                                v1 -= ((v25 / v25) - float32(v2))
                                _ = len(v4)
                                _ = v27
                            } else {
                                var v28 uint8 = 0
                                _ = v28
                            }
                            var v29 uint32 = 849
                            _ = v29
                        }
                        v25 = float32(v2)
                        _ = v25
                    }
                    _ = v21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
        case "foo bar":
            for fi30 := 0; fi30 < 20; fi30++ {
                var v31 uint32 = 2147483648
                var v32 bool = false
                var v33 uint64 = 42
                var v34 int64 = 42
                v35 := GenS2{F0: "\n\t", F1: v34}
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
                _ = v34
                _ = v35
            }
        case "hello":
            switch "\n\t" {
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "foo bar":
                if true {
                    fmt.Println(fi18)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v36 uint32 = 133
                    _ = v36
                }
            case "hello":
                v1 -= 0.0
            }
        }
        var v37 float64 = float64(0.0)
        fmt.Println(v37)
        _ = fi18
        _ = v37
    }
    switch 127 {
    case 18:
        _ = len(v3)
    }
    switch -9223372036854775808 {
    case 42:
        v17 = v17
    default:
        for fi38 := 0; fi38 < 18; fi38++ {
            v39 := GenS2{F0: (("" + "hello") + "foo bar"), F1: (v17 - (^(v17)))}
            _ = fi38
            _ = v39
        }
    }
    v1 -= 1.5
    for fi40 := 0; fi40 < 4; fi40++ {
        v2 += v2
        v17 += 1
        fmt.Println(fi40, v1, v1)
        _ = fi40
    }
    v2 = v2
    fmt.Println(v1, v1, v17)
    v1 -= 3.14159
    v41 := GenS2{F0: ("\n\t" + ("" + "a")), F1: v17}
    v17 += (v17 - v17)
    v1 = v1
    var v42 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
