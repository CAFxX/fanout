package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint16, p2 bool) int16 {
    p1 -= (^(p1))
    v4 := GenS1{F0: 1, F1: 4294967295, F2: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 42
    v3 := make(map[uint32]int)
    var v5 uint8 = 38
    if false {
        switch "a" {
        case "foo bar":
            fmt.Println(v2, v5)
        case "a":
            v5 = 2
        }
    } else {
        if (276 != 4294967295) {
            delete(v1, 1)
            var v6 int32 = (^(v2))
            _ = v6
        } else {
            v2 -= -1
        }
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        if true {
            var v8 int = fi7
            fmt.Println(v2, v8)
            if false {
                v2 = v2
                v9 := GenS2{F0: 0, F1: 0}
                v2 += -375
                _ = v9
            } else {
                switch "a" {
                case "\n\t":
                    var v10 float64 = (((1.5 * 1.5) - ((1.5 * 0.1) - float64((0.0 / float32((float64(0.1) * 1e10)))))) - float64(-2.25))
                    _ = v10
                case "hello":
                    v8 += 4294967295
                }
                var v11 int = fi7
                switch "hello" {
                case "\n\t":
                    for fi12 := 0; fi12 < 6; fi12++ {
                        v5 += (v5 & 42)
                        v13 := GenS1{F0: v2, F1: 33, F2: 1000}
                        v5 -= (v5 + 128)
                        fmt.Println(fi7)
                        v2 -= (^(v2))
                        _ = fi12
                        _ = v13
                    }
                }
                _ = v11
            }
            _ = v8
        } else {
            v5 += (v5 ^ v5)
            var v14 string = ""
            _ = v14
        }
        switch int(v2) {
        case -2147483648:
            v15 := GenS2{F0: 0, F1: 527}
            _ = v15
        }
        _ = fi7
    }
    var v16 int8 = -128
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 uint8 = 127
            _ = fi18
            _ = v19
        }
        if false {
            if (float64(-2.25) >= (((-2.25 + (((0.0 / 1.5) * 0.0) - 0.0)) / (0.1 - 3.14159)) - 3.14159)) {
                var v20 int16 = 1000
                v5 -= v5
                _ = v20
            } else {
                fmt.Println(v5)
                v5 = (v5 << ((0) & 63))
            }
            if (!(65535 == 2147483647)) {
                if (v2 <= (v2 * v2)) {
                    switch (fi17 / ((fi17) | 1)) {
                    case -2147483648:
                        v21 := GenS3{F0: 4294967295}
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 10; fi22++ {
                            _ = len(v1)
                            v1[42] = 926
                            _ = fi22
                        }
                    }
                    switch 32768 {
                    case 9223372036854775807:
                        if (genFn0(42, 0, true) <= 42) {
                            v23 := []int{int(v5), -554, (fi17 >> ((1) & 63)), fi17}
                            switch fi17 {
                            case 913:
                                if (1e10 == (1.5 / 3.14159)) {
                                    _ = len(v23)
                                }
                            case 32768:
                                for fi24 := 0; fi24 < 20; fi24++ {
                                    if (false && (true || true)) {
                                        switch (fi24 & 128) {
                                        case 4294967295:
                                            var v25 uint32 = 4294967295
                                            _ = v25
                                        case 32767:
                                            v26 := GenS2{F0: uint16(v5), F1: 717}
                                            if true {
                                                v5 -= 42
                                                if ((true && true) && (!true)) {
                                                    v23 = append(v23, fi24)
                                                    v27 := GenS4{F0: 42, F1: 1}
                                                    switch "hello" {
                                                    case "hello":
                                                        v5 -= 127
                                                    }
                                                    _ = v27
                                                }
                                                v5 = 42
                                            } else {
                                                v2 += v2
                                                switch fi24 {
                                                case -927:
                                                    v2 += 0
                                                    if true {
                                                        var v28 string = (("a" + "") + ("\n\t" + "\n\t"))
                                                        _ = len(v1)
                                                        v29 := v3[2]
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        if true {
                                                            var v30 uint = 706
                                                            _ = v30
                                                        }
                                                    }
                                                case 2147483647:
                                                    v31 := GenS5{F0: (v5 * (v5 >> ((255) & 63)))}
                                                    _ = v31
                                                case -137:
                                                    var v32 uint64 = 304
                                                    _ = v32
                                                }
                                                var v33 uint16 = 1
                                                _ = v33
                                            }
                                            _ = v26
                                        }
                                        v34 := GenS2{F0: 541, F1: 42}
                                        _ = v34
                                    }
                                    v35 := v3[2147483647]
                                    var v36 string = "foo bar"
                                    _ = fi24
                                    _ = v35
                                    _ = v36
                                }
                            default:
                                v5 = (v5 % (((v5 / (((v5 | v5)) | 1))) | 1))
                            }
                            switch "a" {
                            case "foo bar":
                                v2 = v2
                            default:
                                fmt.Println(fi17)
                            }
                            _ = v23
                        }
                    case -576:
                        for fi37 := 0; fi37 < 3; fi37++ {
                            var v38 int32 = (v2 << ((0) & 63))
                            if true {
                                var v39 int8 = 42
                                _ = v39
                            }
                            _ = fi37
                            _ = v38
                        }
                    case 127:
                        v5 += v5
                    default:
                        var v40 uint32 = 340
                        _ = v40
                    }
                } else {
                    if (42 >= v16) {
                        if false {
                            if true {
                                var v41 int32 = 177
                                fmt.Println(v5)
                                for fi42 := 0; fi42 < 4; fi42++ {
                                    v5 += 162
                                    _ = fi42
                                }
                                _ = v41
                            }
                        } else {
                            if false {
                                var v43 float64 = -2.25
                                var v44 bool = (!(false || (9223372036854775807 < 186)))
                                var v45 uint8 = (v5 / ((v5) | 1))
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                v5 = 1
                            }
                            var v46 float64 = -2.25
                            _ = v46
                        }
                    }
                }
                switch (fi17 / (((fi17 / (((fi17 << ((1) & 63))) | 1))) | 1)) {
                case 127:
                    var v47 bool = (127 == int(v2))
                    _ = v47
                case -52:
                    delete(v1, 0)
                }
                v48 := []int{fi17, (fi17 % ((240) | 1))}
                _ = v48
            }
        } else {
            var v49 string = "\n\t"
            _ = v49
        }
        _ = fi17
    }
    var v50 uint16 = 0
    if false {
        var v51 int64 = 9223372036854775807
        if (!(264 <= 132)) {
            fmt.Println(v16, v5, v16)
        } else {
            var v52 int32 = 2147483647
            v53 := v1[2]
            if (((((3.14159 / (float32(float64((3.14159 * 3.14159))) - 3.14159)) - (((1e10 / 1.5) / 1.0) * ((1.5 * (0.0 / (1.5 * ((((0.1 - (((3.14159 + float32(1.5)) * (1.5 - (1e10 - ((0.0 + 3.14159) + (3.14159 + 1.5))))) - (float32(0.1) + 0.0))) - 3.14159) / float32((0.1 / 1.5))) - 1e10)))) * 3.14159))) / 1.0) == 1.5) && ((false || ((v52 >= (-(v52))) || ((v51 + (^(v51))) != (v51 >> ((18446744073709551615) & 63))))) || ((1 == 0) && false))) {
                var v54 uint8 = (v5 - (v5 / ((v5) | 1)))
                v2 += v2
                fmt.Println(v54)
                _ = v54
            }
            _ = v52
            _ = v53
        }
        for fi55 := 0; fi55 < 13; fi55++ {
            v56 := []int{}
            if (!(false && (!false))) {
                v50 -= v50
                if true {
                    var v57 float32 = 3.14159
                    _ = v57
                }
            } else {
                v2 = v2
                if true {
                    v2 += 0
                    v3[65535] = 279
                } else {
                    if true {
                        var v58 int64 = (^(v51))
                        _ = v58
                    }
                    v59 := GenS5{F0: (v5 | (v5 >> ((1) & 63)))}
                    _ = v59
                }
            }
            fmt.Println(v16, v51, fi55)
            _ = fi55
            _ = v56
        }
        _ = v51
    }
    fmt.Println(v16)
    v5 -= v5
    fmt.Println(v50, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v50
}
