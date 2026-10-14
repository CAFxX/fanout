package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int32; F2 int }
type GenS6 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 uint) uint8 {
    var v5 int32 = 885
    v5 -= (v5 * 2147483647)
    var v6 int8 = -1
    var v7 uint32 = 0
    var v8 uint = 18446744073709551615
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 198
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 int = -1
    var v4 string = (("\n\t" + "a") + "")
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int = 32767
        if false {
            _ = len(v2)
        } else {
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 float64 = (0.0 * 1e10)
                switch 228 {
                case 127:
                    switch ((((v9 + v4) + ((((v9 + v9) + (v9 + (v9 + (("hello" + "") + v9)))) + v4) + ("" + v9))) + "") + (v9 + "")) {
                    case "":
                        v14 := GenS1{F0: 626, F1: -9223372036854775808}
                        _ = v14
                    case "a":
                        if (!(!(true || ((fi10 % (((-(fi12))) | 1)) == (v11 - (v3 + fi12)))))) {
                            v3 = -9223372036854775808
                            v4 = ("" + v9)
                            fmt.Println(v11, v4)
                            v3 = (fi10 | v11)
                            var v15 uint64 = 0
                            _ = v15
                        } else {
                            fmt.Println(v3, fi10)
                            switch (v3 / ((-9223372036854775808) | 1)) {
                            case 9223372036854775807:
                                if false {
                                    if true {
                                        _ = len(v1)
                                        var v16 uint8 = 38
                                        v3 = -2147483648
                                        _ = v16
                                    }
                                    v9 = ("a" + "foo bar")
                                    fmt.Println(v4, v11)
                                }
                            case -9223372036854775808:
                                v17 := []int{}
                                _ = v17
                            default:
                                switch ("a" + (v9 + "")) {
                                case "a":
                                    v18 := GenS1{F0: 42, F1: -400}
                                    v11 = 452
                                    _ = v18
                                default:
                                    var v19 float64 = v13
                                    _ = v19
                                }
                            }
                            var v20 uint = 255
                            _ = v20
                        }
                    default:
                        var v21 uint16 = 0
                        _ = v21
                    }
                case -2147483648:
                    var v22 uint64 = 0
                    _ = v22
                default:
                    switch 32767 {
                    case 812:
                        var v23 uint8 = 2
                        _ = v23
                    default:
                        if (0 != 4294967295) {
                            v24 := v1[v11]
                            _ = v24
                        }
                    }
                }
                v1[(fi10 & v11)] = (-(v3))
                _ = fi12
                _ = v13
            }
        }
        _ = fi10
        _ = v11
    }
    v25 := []int{v3, -287, v3, v3}
    v26 := GenS1{F0: -905, F1: -1}
    if false {
        v3 = (-(v3))
    } else {
        if (333 > 0) {
            v27 := v1[v3]
            _ = v27
        }
        var v28 uint8 = 255
        _ = v28
    }
    v3 -= 9223372036854775807
    var v29 uint16 = 2
    v4 = ("" + "")
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 uint64 = 0
        var v32 bool = true
        v31 = (v31 ^ v31)
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := GenS2{F0: 1}
    if (!(v3 > v3)) {
        _ = len(v25)
        switch (^(v3)) {
        case -2147483648:
            v34 := GenS3{F0: 2, F1: -1}
            _ = v34
        case 440:
            var v35 int = (-(v3))
            _ = v35
        case 4294967295:
            switch "hello" {
            case "a":
                switch (v3 << ((255) & 63)) {
                case 127:
                    var v36 float64 = float64((0.1 / 3.14159))
                    _ = v36
                default:
                    v29 -= v29
                }
            case "hello":
                v29 += v29
            default:
                for fi37 := 0; fi37 < 11; fi37++ {
                    var v38 string = "foo bar"
                    _ = fi37
                    _ = v38
                }
            }
        }
    } else {
        if (-80 <= 0) {
            v29 += v29
            _ = len(v25)
        }
        v4 = v4
    }
    v39 := v2[42]
    var v40 float32 = (3.14159 * -2.25)
    for fi41 := 0; fi41 < 15; fi41++ {
        v25 = append(v25, 1)
        _ = len(v25)
        switch v3 {
        case 2147483647:
            v40 += v40
        case 127:
            fmt.Println(v3, fi41)
        }
        _ = fi41
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        if len(v25) > 0 {
            v25[0] = v3
        }
        fmt.Println(fi42)
        fmt.Println(v9)
        _ = fi42
    }
    if true {
        var v43 string = v9
        if (42 >= 32768) {
            var v44 bool = false
            if v44 {
                var v45 uint8 = 159
                switch v9 {
                case "":
                    v45 += (^(v45))
                case "hello":
                    v46 := GenS2{F0: -2147483648}
                    _ = v46
                case "foo bar":
                    var v47 uint8 = 255
                    _ = v47
                }
                v44 = (1e10 != 1e10)
                _ = v45
            } else {
                switch ("hello" + "foo bar") {
                case "hello":
                    v4 = v43
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        v40 = (1.5 - -2.25)
                    case "":
                        var v48 uint16 = (^(v29))
                        _ = v48
                    case "foo bar":
                        v40 = v40
                    default:
                        if v44 {
                            var v49 int = (v3 / (((v3 - (v39 >> ((255) & 63)))) | 1))
                            switch ((("foo bar" + v9) + "") + (v43 + "")) {
                            case "":
                                if ((v40 - (((v40 * 0.0) * 1.5) - (((float32(3.14159) / float32(-2.25)) - ((0.0 / (v40 + (v40 / ((v40 * v40) / 1e10)))) + (0.1 * v40))) * ((v40 - (v40 + 0.1)) + float32(1.5))))) > float32((float64((3.14159 / -2.25)) + (0.0 * 1.5)))) {
                                    var v50 int64 = 276
                                    if len(v25) > 0 {
                                        v25[0] = -2147483648
                                    }
                                    var v51 string = v4
                                    _ = v50
                                    _ = v51
                                } else {
                                    v43 = (("" + "") + v4)
                                }
                            case "hello":
                                for fi52 := 0; fi52 < 7; fi52++ {
                                    v53 := v2[2147483648]
                                    _ = fi52
                                    _ = v53
                                }
                            case "a":
                                var v54 int8 = 0
                                _ = v54
                            default:
                                v4 = v4
                            }
                            v55 := v2[906]
                            _ = v49
                            _ = v55
                        }
                    }
                }
            }
            for fi56 := 0; fi56 < 14; fi56++ {
                if true {
                    v57 := GenS4{F0: -66}
                    _ = v57
                } else {
                    v58 := GenS5{F0: -1, F1: 1, F2: v3}
                    _ = v58
                }
                fmt.Println(v40, fi56, v9)
                var v59 float64 = (-2.25 + 0.0)
                _ = fi56
                _ = v59
            }
            _ = v44
        } else {
            var v60 string = (v4 + ("a" + v9))
            v4 = (((v60 + v43) + (v9 + v60)) + (("foo bar" + ((v60 + (v4 + v43)) + v60)) + v43))
            fmt.Println(v40, v39, v4)
            _ = v60
        }
        _ = v43
    } else {
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    }
    for fi61 := 0; fi61 < 10; fi61++ {
        v4 = "foo bar"
        if false {
            switch ("foo bar" + ("\n\t" + v4)) {
            case "hello":
                v62 := GenS6{F0: 128, F1: 273, F2: -711}
                _ = v62
            case "foo bar":
                if false {
                    for fi63 := 0; fi63 < 4; fi63++ {
                        var v64 float64 = ((3.14159 - (-2.25 - (1e10 * ((3.14159 + ((0.0 / (-2.25 - 0.0)) + 0.1)) * 1e10)))) / 3.14159)
                        if (4294967295 == 0) {
                            v39 = (fi63 ^ fi63)
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            if false {
                                switch (v3 >> ((193) & 63)) {
                                case 4294967295:
                                    v2[42] = (^(fi63))
                                case 1:
                                    v25 = append(v25, fi63)
                                }
                            } else {
                                fmt.Println(v29)
                                switch 42 {
                                case 127:
                                    var v65 int = (fi61 ^ v39)
                                    _ = v65
                                case 128:
                                    var v66 int64 = 945
                                    _ = v66
                                default:
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                }
                            }
                        }
                        var v67 uint = 14
                        _ = fi63
                        _ = v64
                        _ = v67
                    }
                    v29 += v29
                }
            default:
                v68 := GenS4{F0: 114}
                _ = v68
            }
        }
        _ = fi61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v9
}
