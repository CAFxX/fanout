package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 string }
func genFn0() int8 {
    v7 := GenS1{F0: -128, F1: "foo bar"}
    var v8 float64 = ((((3.14159 - 1.5) - (((3.14159 - (float64(1.5) + (((-2.25 * -2.25) - (0.1 * (3.14159 + ((((-2.25 * (-2.25 * (3.14159 + (1.5 - 1.5)))) * 0.1) + float64(-2.25)) + 0.0)))) + -2.25))) + float64((float32(1.5) + (((((0.1 - 3.14159) + -2.25) * float32(float64((1.5 - (1e10 - ((1e10 - 0.1) - 0.0)))))) * 0.1) / 0.1)))) + 1e10)) + -2.25) + -2.25)
    _ = v7
    _ = v8
    return 47
}

func main() {
    var v1 uint16 = 42
    var v2 int8 = 42
    v3 := []uint32{42, 379}
    v4 := make(map[uint32]bool)
    var v5 uint64 = 1
    var v6 int8 = -1
    v9 := GenS2{F0: 128, F1: v5, F2: 86}
    v10 := GenS3{F0: 1000}
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 float64 = ((1e10 / 0.1) - 0.1)
            v14 := v4[380]
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    v15 := GenS2{F0: 42, F1: v5, F2: -684}
    switch 655 {
    case 9223372036854775807:
        var v16 int8 = (v2 - genFn0())
        _ = v16
    case 42:
        var v17 int64 = 42
        _ = v17
    default:
        var v18 bool = (!false)
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 18; fi19++ {
            v1 += (v1 & 345)
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 int = 2147483647
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        v5 = (^(v5))
        if (false || ((0.1 != 3.14159) || true)) {
            for fi22 := 0; fi22 < 1; fi22++ {
                v6 -= (v2 >> ((0) & 63))
                switch fi22 {
                case 9223372036854775807:
                    v6 += (v6 % (((v6 * v2)) | 1))
                default:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        if ((199 < 1) && false) {
                            switch fi22 {
                            case -2147483648:
                                v6 += v2
                            case -565:
                                v5 -= (^(v5))
                            }
                        } else {
                            var v24 float64 = 0.0
                            switch ("\n\t" + ((("" + "") + "hello") + "\n\t")) {
                            case "\n\t":
                                fmt.Println(v24, fi22)
                            case "":
                                v24 = v24
                            case "hello":
                                if true {
                                    if ((!true) && true) {
                                        v25 := GenS3{F0: 471}
                                        var v26 int32 = -1
                                        _ = v25
                                        _ = v26
                                    } else {
                                        fmt.Println(fi23, fi23, fi22)
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v27 int = (fi23 & (^(fi23)))
                                            _ = v27
                                        case "a":
                                            v2 += (-(v2))
                                        }
                                        var v28 bool = true
                                        _ = v28
                                    }
                                    if (32767 == 42) {
                                        fmt.Println(fi22)
                                        _ = len(v3)
                                        var v29 int16 = 32767
                                        if (!true) {
                                            v30 := GenS2{F0: 2, F1: (v5 / ((4294967295) | 1)), F2: 9223372036854775807}
                                            var v31 int64 = 1
                                            if len(v3) > 0 {
                                                v3[0] = 0
                                            }
                                            _ = v30
                                            _ = v31
                                        }
                                        var v32 uint8 = uint8(v1)
                                        var v33 int32 = -300
                                        if true {
                                            switch ("foo bar" + "\n\t") {
                                            case "foo bar":
                                                var v34 float32 = -2.25
                                                _ = v34
                                            case "a":
                                                fmt.Println(v32)
                                            case "":
                                                var v35 string = "\n\t"
                                                _ = v35
                                            }
                                            v6 = 42
                                            v32 = (v32 / ((255) | 1))
                                        }
                                        _ = v29
                                        _ = v32
                                        _ = v33
                                    }
                                    v36 := GenS3{F0: -1}
                                    _ = v36
                                }
                            default:
                                v37 := GenS1{F0: v6, F1: "hello"}
                                _ = v37
                            }
                            _ = v24
                        }
                        _ = fi23
                    }
                }
                _ = fi22
            }
        } else {
            v5 += 18446744073709551615
        }
    }
    var v38 float64 = (float64(float32(-2.25)) + (1.5 * 3.14159))
    var v39 int64 = 42
    if ((((v1 / (((v1 & v1)) | 1)) <= 1) && false) || (!(1 < 416))) {
        var v40 string = "a"
        v5 += v5
        if true {
            var v41 int64 = -703
            _ = v41
        }
        _ = v40
    }
    v1 += 42
    for fi42 := 0; fi42 < 10; fi42++ {
        for fi43 := 0; fi43 < 9; fi43++ {
            v6 += genFn0()
            delete(v4, 2147483648)
            _ = fi43
        }
        _ = fi42
    }
    var v44 int = 25
    v45 := GenS4{F0: (v39 >> ((18446744073709551615) & 63)), F1: ("\n\t" + "")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v9
}
