package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 int64; F1 uint8; F2 int32 }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]uint32)
    v3 := []int{346, -379, -129, 650}
    var v4 float64 = -2.25
    var v5 float32 = float32((float64(1e10) / 1.0))
    var v6 uint8 = 162
    var v7 float32 = (((v5 + (1.5 - v5)) * (v5 * float32(float64(v5)))) * (3.14159 + 1.5))
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 42}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    delete(v2, -910)
    for fi11 := 0; fi11 < 3; fi11++ {
        v3 = append(v3, (fi11 + (fi11 & (fi11 & 128))))
        switch ("foo bar" + "") {
        case "":
            fmt.Println(fi11, v1, fi11)
        case "foo bar":
            v5 = (((v7 + v7) / 1.0) * (v5 * (v5 + (((0.1 * v5) + ((v7 + (float32(3.14159) / 3.14159)) - (float32(1.5) / 1e10))) * -2.25))))
        case "hello":
            v12 := GenS1{F0: -9223372036854775808}
            _ = v12
        default:
            switch fi11 {
            case -301:
                v1 = (v4 - v1)
            case -286:
                if false {
                    v1 -= float64(v7)
                    var v13 int = 1
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        v15 := GenS1{F0: -244}
                        _ = fi14
                        _ = v15
                    }
                    var v16 float32 = v7
                    _ = v16
                }
            }
        }
        _ = fi11
    }
    v17 := v2[32767]
    var v18 float32 = 0.1
    v7 -= ((float32(((0.1 - (v1 + (v4 - 0.0))) + v4)) - 0.0) + (float32((v1 - v4)) + (v5 * 1e10)))
    var v19 int64 = 42
    _ = len(v2)
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            v19 = -1
            if false {
                if false {
                    if (!(!false)) {
                        switch "hello" {
                        case "":
                            _ = len(v3)
                        default:
                            switch 2147483647 {
                            case 1:
                                if ((v19 * v19) > v19) {
                                    if ((float64(v7) + v1) > (v1 / 1.5)) {
                                        v22 := GenS2{F0: (v19 ^ -1), F1: 0, F2: v19}
                                        v23 := GenS3{F0: -188, F1: (^(v6)), F2: -413}
                                        var v24 float32 = 1e10
                                        _ = v22
                                        _ = v23
                                        _ = v24
                                    } else {
                                        fmt.Println(v1)
                                    }
                                    v19 = v19
                                    v18 = v18
                                } else {
                                    v2[-659] = 202
                                }
                                fmt.Println(v1, v5, v17)
                                v4 += ((float64((((1.5 / float32(3.14159)) * (1.5 - (1e10 + v7))) / v7)) * v4) * v4)
                            case -1:
                                v25 := GenS3{F0: (v19 - (v19 & (v19 & 1))), F1: (v6 << ((65535) & 63)), F2: 674}
                                _ = v25
                            case 0:
                                v4 -= 1.5
                            default:
                                v26 := GenS1{F0: (v19 >> ((18446744073709551615) & 63))}
                                _ = v26
                            }
                        }
                        v3 = append(v3, (fi20 * fi21))
                    } else {
                        fmt.Println(v18, v17, v4)
                        var v27 int64 = v19
                        _ = v27
                    }
                } else {
                    var v28 int = (fi20 | (^(fi21)))
                    v29 := GenS3{F0: v19, F1: v6, F2: -1}
                    var v30 int64 = (v19 - v19)
                    _ = v28
                    _ = v29
                    _ = v30
                }
                var v31 float32 = float32(1e10)
                v1 -= (v1 / v1)
                var v32 uint32 = v17
                _ = len(v3)
                _ = v31
                _ = v32
            }
            _ = fi21
        }
        _ = fi20
    }
    _ = len(v3)
    var v33 int32 = 2147483647
    if true {
        for fi34 := 0; fi34 < 2; fi34++ {
            for fi35 := 0; fi35 < 19; fi35++ {
                v36 := v2[fi34]
                v37 := GenS2{F0: (v19 >> ((1) & 63)), F1: 1, F2: (-(v19))}
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi34
        }
        v38 := GenS3{F0: (v19 ^ int64(v6)), F1: (v6 & v6), F2: v33}
        v39 := GenS3{F0: 1, F1: (v6 - (^(v6))), F2: v33}
        _ = v38
        _ = v39
    } else {
        v5 += v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
