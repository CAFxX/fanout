package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int16; F1 uint8; F2 int }
type GenS6 struct { F0 int16; F1 string; F2 uint16 }
func genFn0(p0 uint64) uint16 {
    v4 := GenS1{F0: 2, F1: 2, F2: -32768}
    switch 653 {
    case 64:
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := []int{127}
            _ = fi5
            _ = v6
        }
    case 32768:
        for fi7 := 0; fi7 < 5; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := []int{fi7, (-(fi8))}
                var v10 uint8 = 239
                _ = fi8
                _ = v10
                _ = v9
            }
            p0 -= (p0 - (p0 * (p0 / ((4294967295) | 1))))
            _ = fi7
        }
    default:
        var v11 int16 = -447
        _ = v11
    }
    _ = p0
    _ = v4
    return 828
}
func genFn1() uint8 {
    var v12 float32 = ((float32(1.5) / (3.14159 + 0.1)) + (0.1 - ((((3.14159 + float32(0.0)) * ((-2.25 * -2.25) / (-2.25 - 1e10))) / (((float32(1.5) - -2.25) * (float32(0.0) - 1e10)) - (0.0 / float32(1.5)))) - (float32(float64(3.14159)) * 0.1))))
    v13 := GenS2{F0: 467}
    if (!true) {
        v12 -= (v12 - v12)
        fmt.Println(v12)
    } else {
        fmt.Println(v12)
        fmt.Println(v12)
    }
    switch 397 {
    case 193:
        if true {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 float32 = float32(0.0)
                var v16 int8 = 82
                _ = fi14
                _ = v15
                _ = v16
            }
            v12 = (float32(3.14159) - (0.1 - v12))
            for fi17 := 0; fi17 < 20; fi17++ {
                if false {
                    v18 := GenS1{F0: 0, F1: 2, F2: 0}
                    v19 := []int{}
                    if ("\n\t" == "hello") {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            fmt.Println(v12, fi20, v12)
                            if (!false) {
                                switch fi17 {
                                case 4294967295:
                                    switch "" {
                                    case "\n\t":
                                        v12 -= 1e10
                                    case "":
                                        v12 += 1e10
                                    default:
                                        var v21 float32 = v12
                                        _ = v21
                                    }
                                case 42:
                                    fmt.Println(fi20, fi17)
                                case -9223372036854775808:
                                    _ = len(v19)
                                }
                                if false {
                                    switch (fi17 / ((32767) | 1)) {
                                    case -2147483648:
                                        v22 := GenS2{F0: 85}
                                        _ = v22
                                    case -1:
                                        if false {
                                            if true {
                                                v23 := GenS1{F0: 42, F1: 1, F2: 32767}
                                                var v24 int64 = 42
                                                _ = v23
                                                _ = v24
                                            }
                                            var v25 float64 = 0.0
                                            var v26 int8 = -100
                                            _ = v25
                                            _ = v26
                                        }
                                    }
                                    fmt.Println(fi20)
                                    var v27 uint8 = 127
                                    _ = v27
                                } else {
                                    fmt.Println(fi17, v12)
                                    switch (fi17 << ((65535) & 63)) {
                                    case 763:
                                        v12 += 0.0
                                    case 0:
                                        v19 = append(v19, (fi17 % ((fi17) | 1)))
                                    default:
                                        var v28 float32 = v12
                                        _ = v28
                                    }
                                }
                                fmt.Println(fi17)
                            }
                            if (2 >= 0) {
                                fmt.Println(fi17)
                                var v29 uint16 = genFn0(1)
                                var v30 int8 = 1
                                _ = v29
                                _ = v30
                            } else {
                                var v31 uint = 2
                                v31 = (v31 >> ((781) & 63))
                                v32 := GenS3{F0: (fi17 & (fi17 ^ (-(fi20))))}
                                _ = v31
                                _ = v32
                            }
                            _ = fi20
                        }
                    }
                    _ = v18
                    _ = v19
                } else {
                    switch ("foo bar" + "") {
                    case "\n\t":
                        var v33 int64 = -908
                        _ = v33
                    case "foo bar":
                        v34 := []int{fi17}
                        _ = v34
                    }
                }
                v35 := GenS4{F0: 255, F1: (!(false || (fi17 <= fi17)))}
                _ = fi17
                _ = v35
            }
        }
    default:
        v12 = v12
    }
    _ = v12
    _ = v13
    return 229
}

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 0
    var v3 float32 = (1e10 + float32(1.5))
    v36 := GenS1{F0: 575, F1: (v2 >> ((42) & 63)), F2: 182}
    if true {
        var v37 float64 = 3.14159
        _ = v37
    } else {
        var v38 int8 = 1
        _ = v38
    }
    for fi39 := 0; fi39 < 7; fi39++ {
        v40 := GenS5{F0: 1000, F1: genFn1(), F2: (-(fi39))}
        _ = fi39
        _ = v40
    }
    v3 += ((v3 + 0.0) / 0.1)
    fmt.Println(v1, v3, v2)
    v41 := []int{-645, 349}
    for fi42 := 0; fi42 < 17; fi42++ {
        v43 := GenS6{F0: -1, F1: "hello", F2: (v2 + v2)}
        _ = fi42
        _ = v43
    }
    var v44 int64 = -1
    v45 := GenS5{F0: -32768, F1: genFn1(), F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v41
    _ = v44
    _ = v45
}
