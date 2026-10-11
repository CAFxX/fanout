package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 string; F1 uint32; F2 int64 }
type GenS5 struct { F0 int32; F1 uint8; F2 int }
func genFn0(p0 uint64, p1 uint, p2 int8) uint {
    v7 := []int{-9223372036854775808, 127, -91, 127}
    p0 += 18446744073709551615
    p1 += 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 int32 = -2147483648
    v2 := []int{127, 32767, -1, 127}
    var v3 uint32 = 4294967295
    v4 := []uint32{(^(v3))}
    v5 := make(map[int]string)
    v6 := make(map[string]bool)
    switch -800 {
    case -9223372036854775808:
        v8 := GenS1{F0: 42}
        _ = v8
    default:
        v1 = (^(v1))
    }
    var v9 uint16 = 255
    v3 += (v3 - (v3 << ((uint(v9)) & 63)))
    switch "a" {
    case "foo bar":
        fmt.Println(v9, v1)
    case "\n\t":
        fmt.Println(v1)
    case "hello":
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 uint8 = 1
            for fi12 := 0; fi12 < 10; fi12++ {
                if len(v2) > 0 {
                    v2[0] = fi12
                }
                var v13 uint16 = (v9 >> ((29) & 63))
                var v14 int64 = 42
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    default:
        var v15 int32 = v1
        _ = v15
    }
    var v16 uint32 = (v3 % ((v3) | 1))
    switch 4294967295 {
    case 0:
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(v1)
            v18 := GenS1{F0: 171}
            _ = fi17
            _ = v18
        }
    case -1:
        v1 -= v1
    case -208:
        v19 := GenS1{F0: 237}
        _ = v19
    default:
        _ = len(v6)
    }
    if (("" + "foo bar") == "hello") {
        v16 += 2147483648
    } else {
        if (507 == (v16 << ((409) & 63))) {
            v3 += v3
            var v20 string = "hello"
            v21 := GenS1{F0: 128}
            _ = v20
            _ = v21
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v16 = 2
    v6[(("foo bar" + ("a" + ("foo bar" + "hello"))) + "")] = ((v9 >= v9) && (0 < 0))
    if (v16 == v3) {
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := v5[(fi22 | (-(fi22)))]
            _ = fi22
            _ = v23
        }
    } else {
        switch -2147483648 {
        case 1:
            switch 1 {
            case -439:
                for fi24 := 0; fi24 < 11; fi24++ {
                    switch fi24 {
                    case 32767:
                        v25 := GenS2{F0: 0, F1: 1}
                        _ = v25
                    default:
                        v9 += v9
                    }
                    var v26 float32 = float32(0.0)
                    _ = fi24
                    _ = v26
                }
            case -1:
                v27 := GenS3{F0: true, F1: v9}
                _ = v27
            default:
                if true {
                    for fi28 := 0; fi28 < 1; fi28++ {
                        v9 = (v9 | v9)
                        v9 += v9
                        var v29 int8 = 1
                        _ = fi28
                        _ = v29
                    }
                    if true {
                        if (!(false && true)) {
                            if (-1 >= -32768) {
                                v3 += v16
                                fmt.Println(v16)
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    for fi31 := 0; fi31 < 10; fi31++ {
                                        fmt.Println(v1, fi30)
                                        v6[("foo bar" + "\n\t")] = (!false)
                                        var v32 int = 0
                                        switch "hello" {
                                        case "hello":
                                            var v33 uint16 = 920
                                            _ = v33
                                        case "foo bar":
                                            var v34 uint16 = (v9 % ((32767) | 1))
                                            v35 := GenS4{F0: "foo bar", F1: v16, F2: 42}
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = fi31
                                        _ = v32
                                    }
                                    _ = fi30
                                }
                            }
                            if ((0.0 != (((0.0 + 3.14159) * float64((3.14159 * ((1e10 / -2.25) - 3.14159)))) + -2.25)) || true) {
                                var v36 int32 = -28
                                _ = len(v6)
                                v36 += v36
                                _ = v36
                            }
                        } else {
                            fmt.Println(v9)
                        }
                        var v37 bool = ((1.5 + ((0.0 + 3.14159) - ((3.14159 - 0.0) + 3.14159))) <= 1.5)
                        v1 = v1
                        _ = v37
                    } else {
                        v38 := GenS1{F0: 42}
                        _ = v38
                    }
                }
            }
        }
        var v39 float32 = 0.0
        _ = v39
    }
    var v40 int64 = -87
    fmt.Println(v3, v3)
    if false {
        fmt.Println(v3, v16)
        var v41 uint8 = 128
        _ = v41
    }
    v9 = v9
    if (float64((0.0 + 0.1)) == 0.1) {
        var v42 string = "a"
        switch -9223372036854775808 {
        case 0:
            for fi43 := 0; fi43 < 10; fi43++ {
                if (42 > 0) {
                    switch ((((("\n\t" + v42) + v42) + "a") + (("a" + ((v42 + v42) + "foo bar")) + v42)) + v42) {
                    case "":
                        _ = len(v4)
                    case "\n\t":
                        var v44 int64 = v40
                        _ = v44
                    case "foo bar":
                        v45 := GenS2{F0: 9223372036854775807, F1: 1}
                        _ = v45
                    default:
                        v3 = 0
                    }
                    if false {
                        v9 += (v9 & 2)
                        for fi46 := 0; fi46 < 7; fi46++ {
                            v47 := v5[(fi43 >> ((25) & 63))]
                            v47 = ("a" + (v42 + (("foo bar" + "foo bar") + v42)))
                            v47 = v47
                            _ = fi46
                            _ = v47
                        }
                        for fi48 := 0; fi48 < 19; fi48++ {
                            if true {
                                _ = len(v6)
                                if ((((float64(1e10) - 0.0) == (1e10 / 0.1)) && (2 >= 0)) || true) {
                                    var v49 uint64 = 2
                                    if (genFn0((v49 | 704), 16, -87) <= 18446744073709551615) {
                                        fmt.Println(v16)
                                        var v50 int = (fi48 + int(v40))
                                        _ = v50
                                    }
                                    _ = v49
                                }
                            }
                            v51 := GenS1{F0: 39}
                            _ = fi48
                            _ = v51
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = v16
                        }
                        for fi52 := 0; fi52 < 12; fi52++ {
                            v9 = 32767
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi52
                        }
                        var v53 int32 = -1
                        _ = v53
                    }
                    fmt.Println(fi43)
                } else {
                    v4 = append(v4, 2147483648)
                    v9 = (v9 + 1)
                }
                _ = fi43
            }
        case 128:
            if ((0.0 + (float64(1.5) * 1.5)) != -2.25) {
                if len(v4) > 0 {
                    v4[0] = (v16 - (^(v3)))
                }
                var v54 float32 = (-2.25 + (1.5 * (0.1 + 0.0)))
                if ((1e10 / v54) <= ((v54 * v54) - v54)) {
                    switch 32767 {
                    case 42:
                        v40 += 1
                    case 26:
                        v55 := GenS2{F0: -884, F1: -260}
                        _ = v55
                    case -502:
                        if true {
                            if len(v2) > 0 {
                                v2[0] = -1
                            }
                            for fi56 := 0; fi56 < 18; fi56++ {
                                v57 := GenS5{F0: v1, F1: 180, F2: (fi56 | 9223372036854775807)}
                                _ = fi56
                                _ = v57
                            }
                            var v58 uint64 = 82
                            _ = v58
                        } else {
                            var v59 int = 32767
                            _ = v59
                        }
                    default:
                        v3 -= v3
                    }
                    delete(v6, v42)
                } else {
                    v5[32767] = v42
                    v3 = v3
                    if (((v42 + "a") + "foo bar") == (v42 + v42)) {
                        var v60 uint = 1
                        var v61 uint64 = 18446744073709551615
                        v42 = "foo bar"
                        _ = v60
                        _ = v61
                    } else {
                        v2 = append(v2, 443)
                        switch 32768 {
                        case -1000:
                            switch v42 {
                            case "foo bar":
                                if false {
                                    var v62 uint8 = 1
                                    for fi63 := 0; fi63 < 19; fi63++ {
                                        v64 := v5[fi63]
                                        for fi65 := 0; fi65 < 1; fi65++ {
                                            var v66 bool = (uint(v62) != 0)
                                            if true {
                                                v4 = append(v4, (v16 >> ((42) & 63)))
                                            }
                                            _ = fi65
                                            _ = v66
                                        }
                                        _ = fi63
                                        _ = v64
                                    }
                                    _ = v62
                                } else {
                                    var v67 bool = (128 != 255)
                                    v68 := v5[-9223372036854775808]
                                    _ = v67
                                    _ = v68
                                }
                            case "":
                                var v69 bool = true
                                _ = v69
                            }
                        case 9223372036854775807:
                            v70 := GenS5{F0: v1, F1: 0, F2: -442}
                            _ = v70
                        default:
                            fmt.Println(v40)
                        }
                        if true {
                            fmt.Println(v9, v42, v40)
                        } else {
                            v3 += (v16 >> ((genFn0(2, 0, 127)) & 63))
                            var v71 int64 = 1
                            for fi72 := 0; fi72 < 5; fi72++ {
                                v1 -= v1
                                _ = fi72
                            }
                            _ = v71
                        }
                    }
                }
                _ = v54
            }
        case 478:
            v16 += 65535
        default:
            var v73 float32 = 1e10
            _ = v73
        }
        v9 -= (^(v9))
        _ = v42
    } else {
        var v74 int = 729
        fmt.Println(v16)
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
