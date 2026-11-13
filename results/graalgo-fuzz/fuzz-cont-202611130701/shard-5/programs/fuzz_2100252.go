package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int8; F2 int }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int; F1 uint8 }
func genFn0(p0 int16) uint8 {
    var v4 uint = 255
    var v5 uint16 = 945
    v6 := GenS1{F0: int64(v4)}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 199
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{2147483648, 2147483648, 0}
    var v3 uint64 = 0
    for fi7 := 0; fi7 < 7; fi7++ {
        if (true && true) {
            var v8 int8 = 100
            var v9 uint64 = (^(v3))
            _ = v8
            _ = v9
        } else {
            v3 = v3
        }
        v3 = 42
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: 99}
            v3 = 4294967295
            _ = fi10
            _ = v11
        }
        _ = fi7
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS1{F0: 0}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        v3 += (v3 & v3)
        if len(v2) > 0 {
            v2[0] = 697
        }
        _ = fi14
    }
    v15 := v1[42]
    fmt.Println(v3)
    switch (v15 + v15) {
    case "foo bar":
        v3 = v3
    case "":
        v15 = v15
    case "hello":
        v3 -= (v3 + v3)
    default:
        v16 := GenS2{F0: 229}
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 uint8 = 2
            _ = len(v2)
            var v19 int64 = 42
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        _ = len(v2)
    }
    var v20 int = 0
    var v21 float32 = float32((1.5 / 3.14159))
    v2 = append(v2, 867)
    for fi22 := 0; fi22 < 20; fi22++ {
        if ((!((32767 != 1000) || false)) || false) {
            switch v20 {
            case -975:
                var v23 uint8 = 2
                _ = v23
            case 42:
                var v24 uint8 = genFn0(int16(v3))
                _ = v24
            case 32768:
                v3 -= (v3 << ((738) & 63))
            default:
                var v25 int = (v20 % (((^(fi22))) | 1))
                _ = v25
            }
            v26 := GenS3{F0: (v3 << ((18446744073709551615) & 63)), F1: -128, F2: fi22}
            _ = v26
        }
        _ = fi22
    }
    switch v15 {
    case "":
        if (false || (true || (true || ((v15 + v15) > v15)))) {
            var v27 int8 = 100
            _ = v27
        } else {
            var v28 int16 = 721
            for fi29 := 0; fi29 < 1; fi29++ {
                v15 = v15
                _ = fi29
            }
            for fi30 := 0; fi30 < 2; fi30++ {
                switch (("a" + (v15 + ("\n\t" + (((("" + v15) + v15) + v15) + v15)))) + ("hello" + v15)) {
                case "\n\t":
                    if (((genFn0((v28 - v28)) >= 167) || true) && true) {
                        for fi31 := 0; fi31 < 13; fi31++ {
                            switch (fi30 >> ((4294967295) & 63)) {
                            case 2147483647:
                                v32 := GenS4{F0: (v3 >> ((425) & 63)), F1: 1000}
                                _ = v32
                            case 128:
                                v28 -= 1
                            }
                            v33 := GenS4{F0: (v3 << ((2) & 63)), F1: (v28 | 874)}
                            _ = fi31
                            _ = v33
                        }
                    } else {
                        if (!true) {
                            var v34 int16 = (v28 + -32768)
                            var v35 float64 = (1.5 - 0.0)
                            _ = v34
                            _ = v35
                        }
                    }
                case "":
                    for fi36 := 0; fi36 < 3; fi36++ {
                        v28 = v28
                        if (353 >= 2) {
                            var v37 uint = 0
                            v1[9223372036854775807] = "a"
                            switch 4294967295 {
                            case -9223372036854775808:
                                switch v15 {
                                case "a":
                                    if false {
                                        if (v28 > (v28 * v28)) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        v37 += 878
                                        var v38 float32 = v21
                                        _ = v38
                                    }
                                default:
                                    v20 = (fi30 - (fi30 % ((4294967295) | 1)))
                                }
                                if (2147483648 > 1) {
                                    v2 = append(v2, 4294967295)
                                    delete(v1, 127)
                                    if false {
                                        v39 := GenS2{F0: 2147483648}
                                        _ = v39
                                    } else {
                                        v40 := GenS5{F0: v20, F1: 50}
                                        _ = v40
                                    }
                                }
                            case 127:
                                if true {
                                    switch v15 {
                                    case "hello":
                                        _ = len(v2)
                                    }
                                    v41 := GenS1{F0: -846}
                                    _ = v41
                                } else {
                                    var v42 uint64 = v3
                                    var v43 int8 = 42
                                    _ = v42
                                    _ = v43
                                }
                            }
                            _ = v37
                        } else {
                            switch fi36 {
                            case 625:
                                if (651 <= 61) {
                                    fmt.Println(v3)
                                    fmt.Println(v20, fi36, v28)
                                    v28 += v28
                                } else {
                                    delete(v1, (fi36 >> ((733) & 63)))
                                }
                            case -888:
                                _ = len(v1)
                            }
                            switch v15 {
                            case "":
                                v1[(fi36 & 777)] = "hello"
                            default:
                                v44 := v1[fi30]
                                _ = v44
                            }
                            var v45 int = -9223372036854775808
                            var v46 uint32 = 4294967295
                            v47 := GenS4{F0: v3, F1: v28}
                            fmt.Println(fi30)
                            var v48 float64 = float64(v21)
                            if (3.14159 > ((v21 * (0.0 / 3.14159)) + -2.25)) {
                                var v49 bool = false
                                var v50 int = fi30
                                _ = v49
                                _ = v50
                            }
                            _ = v45
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        if len(v2) > 0 {
                            v2[0] = 238
                        }
                        var v51 int64 = 9223372036854775807
                        _ = fi36
                        _ = v51
                    }
                }
                v52 := GenS5{F0: fi30, F1: genFn0(v28)}
                _ = fi30
                _ = v52
            }
            _ = v28
        }
    case "hello":
        fmt.Println(v15, v20, v20)
    }
    switch (v20 % (((v20 * v20)) | 1)) {
    case 42:
        fmt.Println(v21, v20)
    case 4294967295:
        var v53 int32 = -62
        _ = v53
    default:
        delete(v1, (v20 % ((127) | 1)))
    }
    var v54 float64 = 1e10
    for fi55 := 0; fi55 < 1; fi55++ {
        v54 += v54
        fmt.Println(v54)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v54
}
