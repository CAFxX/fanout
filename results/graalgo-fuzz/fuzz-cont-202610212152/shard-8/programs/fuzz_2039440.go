package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int32; F1 uint32 }
func genFn0(p0 int16) int64 {
    v4 := []int{-307, -893, 32767, -244}
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            if (0 > 255) {
                var v7 uint32 = 65535
                _ = v7
            }
            _ = fi6
        }
        var v8 int16 = (-(p0))
        for fi9 := 0; fi9 < 5; fi9++ {
            p0 = p0
            _ = fi9
        }
        _ = fi5
        _ = v8
    }
    var v10 int = -748
    for fi11 := 0; fi11 < 2; fi11++ {
        switch "foo bar" {
        case "a":
            switch (fi11 % (((fi11 & (fi11 | fi11))) | 1)) {
            case -2147483648:
                p0 += (p0 | p0)
            case -1:
                v12 := GenS1{F0: -1, F1: 761}
                _ = v12
            case 9223372036854775807:
                if true {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 int32 = -116
                        var v15 int = v10
                        v16 := GenS1{F0: (v14 & v14), F1: 994}
                        var v17 float32 = 3.14159
                        v18 := GenS1{F0: -1, F1: 192}
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    var v19 uint8 = 42
                    var v20 int32 = 0
                    _ = v19
                    _ = v20
                } else {
                    p0 += -1
                    if false {
                        if true {
                            var v21 float64 = (0.0 - (0.1 + 1.5))
                            fmt.Println(p0)
                            _ = v21
                        } else {
                            if ("\n\t" != "foo bar") {
                                switch fi11 {
                                case 164:
                                    v22 := GenS1{F0: 0, F1: 65535}
                                    _ = v22
                                default:
                                    v4 = append(v4, (fi11 & fi11))
                                }
                            } else {
                                _ = len(v4)
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        if (!(42 != -2147483648)) {
                            v10 = (fi11 * 127)
                            switch v10 {
                            case -9223372036854775808:
                                v23 := GenS1{F0: -138, F1: 681}
                                _ = v23
                            }
                        } else {
                            p0 = (p0 % (((p0 - p0)) | 1))
                        }
                        p0 -= 1
                    }
                }
            default:
                fmt.Println(v10)
            }
        case "hello":
            for fi24 := 0; fi24 < 11; fi24++ {
                v4 = append(v4, (fi24 / ((fi11) | 1)))
                v10 = (^(fi11))
                _ = fi24
            }
        case "foo bar":
            var v25 int32 = -732
            _ = v25
        default:
            if ((p0 * (p0 * 1000)) <= 42) {
                var v26 int64 = -662
                _ = v26
            } else {
                for fi27 := 0; fi27 < 11; fi27++ {
                    v4 = append(v4, (fi11 / ((fi11) | 1)))
                    fmt.Println(p0)
                    _ = fi27
                }
                v28 := GenS2{F0: 42, F1: int32(fi11)}
                _ = v28
            }
        }
        v29 := GenS2{F0: 32767, F1: -935}
        _ = fi11
        _ = v29
    }
    _ = p0
    _ = v10
    _ = v4
    return 1
}
func genFn1(p0 uint16, p1 int) int16 {
    if (!true) {
        p1 += (-(p1))
        v30 := []int{(p1 ^ 1), (p1 >> ((42) & 63))}
        if (true || (9223372036854775807 != -9223372036854775808)) {
            for fi31 := 0; fi31 < 18; fi31++ {
                if len(v30) > 0 {
                    v30[0] = fi31
                }
                if false {
                    var v32 string = "hello"
                    _ = v32
                }
                v33 := GenS3{F0: 0, F1: 0}
                _ = fi31
                _ = v33
            }
            if len(v30) > 0 {
                v30[0] = 919
            }
            p1 = (p1 % ((p1) | 1))
        }
        _ = v30
    } else {
        v34 := []int{}
        for fi35 := 0; fi35 < 18; fi35++ {
            p0 = (p0 << ((1) & 63))
            _ = fi35
        }
        _ = v34
    }
    p1 += 4294967295
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 65535
    v3 := []int{}
    var v36 float64 = 1e10
    var v37 int = 128
    v38 := GenS1{F0: 0, F1: 150}
    v2 = (v2 - (v2 % ((v2) | 1)))
    var v39 int8 = -91
    fmt.Println(v37)
    v40 := GenS1{F0: -1, F1: v2}
    for fi41 := 0; fi41 < 12; fi41++ {
        fmt.Println(v39)
        fmt.Println(v2, v36)
        _ = fi41
    }
    v42 := v1[(v37 >> ((651) & 63))]
    var v43 float64 = float64((-2.25 * (((((1.5 - 1e10) * (float32(((((v36 * -2.25) + (((3.14159 * ((1.5 - (1.5 - (v36 + (v36 * float64((1.5 * 1.5)))))) * v36)) - (((((3.14159 / v36) - v36) + (0.1 / v36)) + v36) + 3.14159)) + -2.25)) + (1.5 * 0.0)) / v36)) * 1.5)) / -2.25) / float32((v36 * v36))) * ((0.0 + 3.14159) + 1e10))))
    _ = len(v3)
    v43 -= v36
    v2 += (v2 << ((1) & 63))
    var v44 int = (-(v37))
    fmt.Println(v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v42
    _ = v43
    _ = v44
}
