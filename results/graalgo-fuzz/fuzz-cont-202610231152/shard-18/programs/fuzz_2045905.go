package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 int64; F1 int; F2 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 1
    v3 := make(map[uint32]int)
    var v4 uint32 = 2147483647
    var v5 int32 = -2147483648
    v6 := make(map[string]uint32)
    var v7 float32 = (0.1 + -2.25)
    var v8 bool = true
    switch "hello" {
    case "\n\t":
        switch 4294967295 {
        case 475:
            switch -451 {
            case 42:
                if ((!(31 != 1)) || v8) {
                    v9 := GenS1{F0: 0, F1: 6, F2: (v5 - (v5 % ((v5) | 1)))}
                    if (((v7 * v7) - ((v7 + v7) - float32((3.14159 + (((3.14159 * (1.5 - ((((1.5 * -2.25) * (0.0 / float64((float32(0.1) * v7)))) / 0.1) * 1e10))) / 1.0) - -2.25))))) == v7) {
                        switch ("" + "\n\t") {
                        case "a":
                            if (993 <= 65535) {
                                v10 := GenS1{F0: 1, F1: 2147483647, F2: (v5 % ((-992) | 1))}
                                _ = v10
                            } else {
                                if (1 > 42) {
                                    v11 := GenS1{F0: -1, F1: 0, F2: 0}
                                    var v12 uint = 65535
                                    _ = v11
                                    _ = v12
                                } else {
                                    v4 -= (v4 ^ (v4 * 481))
                                }
                            }
                        default:
                            var v13 uint8 = 2
                            _ = v13
                        }
                        switch "" {
                        case "":
                            var v14 int64 = 42
                            _ = v14
                        case "foo bar":
                            if (0.1 <= -2.25) {
                                fmt.Println(v8, v2, v2)
                                switch 801 {
                                case 0:
                                    for fi15 := 0; fi15 < 5; fi15++ {
                                        var v16 uint = 0
                                        for fi17 := 0; fi17 < 1; fi17++ {
                                            v18 := GenS1{F0: 42, F1: -99, F2: (v5 << (((v16 >> (((v16 - (^(v16)))) & 63))) & 63))}
                                            _ = fi17
                                            _ = v18
                                        }
                                        _ = fi15
                                        _ = v16
                                    }
                                case -9223372036854775808:
                                    switch 1 {
                                    case -414:
                                        fmt.Println(v5, v4, v4)
                                    case 32768:
                                        var v19 int16 = v2
                                        _ = v19
                                    default:
                                        v20 := GenS1{F0: -9223372036854775808, F1: 42, F2: 1}
                                        _ = v20
                                    }
                                }
                                v21 := []int{731, -1}
                                _ = v21
                            } else {
                                v5 = (v5 >> ((343) & 63))
                            }
                        case "hello":
                            v22 := GenS1{F0: 0, F1: int(v5), F2: 0}
                            _ = v22
                        default:
                            v2 -= (-(v2))
                        }
                    } else {
                        v4 += v4
                    }
                    _ = v9
                }
            case 1:
                for fi23 := 0; fi23 < 2; fi23++ {
                    fmt.Println(v4)
                    v4 = 42
                    _ = fi23
                }
            case 2147483647:
                var v24 uint32 = v4
                _ = v24
            default:
                v4 = (v4 | (^(v4)))
            }
        }
    case "hello":
        v4 = (v4 ^ v4)
    case "a":
        for fi25 := 0; fi25 < 8; fi25++ {
            v26 := GenS2{F0: ("\n\t" + ("hello" + "a"))}
            v27 := []int{2147483647}
            _ = fi25
            _ = v26
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 5; fi28++ {
            if (!(1 <= 214)) {
                switch "a" {
                case "a":
                    var v29 int32 = (v5 * v5)
                    _ = v29
                case "hello":
                    var v30 int32 = -2147483648
                    _ = v30
                }
                v31 := GenS2{F0: "foo bar"}
                v32 := GenS1{F0: 843, F1: (fi28 - -175), F2: (v5 % ((v5) | 1))}
                _ = v31
                _ = v32
            }
            switch "foo bar" {
            case "\n\t":
                var v33 int16 = v2
                _ = v33
            case "hello":
                v34 := []int{}
                _ = v34
            default:
                if (2147483647 == (v5 | (-(v5)))) {
                    fmt.Println(v8, fi28)
                    v35 := GenS2{F0: (("a" + "") + ("hello" + "foo bar"))}
                    _ = v35
                }
            }
            _ = fi28
        }
    }
    v36 := GenS2{F0: "hello"}
    if v8 {
        v37 := []int{1, 669, 1, 748}
        _ = v37
    } else {
        if false {
            switch "" {
            case "foo bar":
                var v38 float32 = ((v7 - v7) + (((v7 * 1.5) - v7) / 0.1))
                _ = v38
            case "hello":
                v7 = ((1e10 + 1.5) * float32((0.1 / float64(v7))))
            case "a":
                fmt.Println(v2, v2, v2)
            }
            var v39 float64 = 1.5
            _ = v39
        }
        v40 := GenS3{F0: "\n\t", F1: 42}
        _ = v40
    }
    if (((!(32767 < 32767)) || ((float32(1e10) / (((float32(((0.1 * -2.25) - -2.25)) * 0.1) / 3.14159) - (3.14159 / v7))) >= float32((0.0 - (1.5 / 1.5))))) && (784 == 42)) {
        v41 := []int{-901, 720, -515}
        v2 += v2
        v7 = (float32(1e10) + (v7 / float32(3.14159)))
        _ = v41
    } else {
        for fi42 := 0; fi42 < 16; fi42++ {
            var v43 int64 = 9223372036854775807
            for fi44 := 0; fi44 < 8; fi44++ {
                switch 9223372036854775807 {
                case 32767:
                    if (!true) {
                        var v45 uint = 530
                        delete(v3, 722)
                        fmt.Println(v4)
                        _ = v45
                    } else {
                        delete(v3, v4)
                        v43 = 9223372036854775807
                        var v46 bool = (2147483647 > (v4 ^ 2))
                        v4 -= 42
                        fmt.Println(fi44, fi42, fi42)
                        _ = v46
                    }
                case 128:
                    v8 = (0 <= 2)
                case -9223372036854775808:
                    v47 := GenS4{F0: (v4 | (v4 + v4))}
                    _ = v47
                default:
                    v2 = v2
                }
                _ = fi44
            }
            _ = fi42
            _ = v43
        }
        if v8 {
            var v48 int = -230
            switch 0 {
            case 42:
                var v49 uint8 = 0
                _ = v49
            case 661:
                for fi50 := 0; fi50 < 5; fi50++ {
                    fmt.Println(v2)
                    _ = fi50
                }
            case 32768:
                for fi51 := 0; fi51 < 16; fi51++ {
                    _ = len(v3)
                    v48 -= 2147483647
                    var v52 int16 = 1
                    _ = fi51
                    _ = v52
                }
            }
            _ = v48
        }
        v3[v4] = -1
    }
    var v53 uint32 = 4294967295
    for fi54 := 0; fi54 < 9; fi54++ {
        v55 := GenS5{F0: (("foo bar" + "") + "\n\t"), F1: fi54}
        delete(v3, (v4 ^ v53))
        _ = fi54
        _ = v55
    }
    v2 = (^(v2))
    var v56 float32 = (v7 * (float32(0.1) - v7))
    var v57 uint8 = uint8(v4)
    if ((32767 > 207) && v8) {
        for fi58 := 0; fi58 < 6; fi58++ {
            v4 += v53
            v57 += 2
            var v59 float64 = 0.1
            _ = fi58
            _ = v59
        }
    }
    var v60 float32 = (3.14159 - 1.5)
    v61 := []int{0}
    v62 := GenS6{F0: -52, F1: 2147483647, F2: v5}
    v6[((("a" + "hello") + "") + "a")] = v53
    v63 := GenS4{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v53
    _ = v56
    _ = v57
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v7
    _ = v8
}
