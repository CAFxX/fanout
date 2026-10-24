package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int32; F1 int8; F2 string }
type GenS4 struct { F0 int8; F1 int64 }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]uint32)
    var v3 int64 = 0
    v4 := make(map[uint32]string)
    var v5 float32 = float32(-2.25)
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 string = ""
        var v8 uint8 = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int32 = -1
    if false {
        v5 -= (float32((1.5 + 0.0)) * -2.25)
        switch "a" {
        case "\n\t":
            v5 = v5
        case "foo bar":
            v10 := GenS1{F0: 255, F1: true, F2: 1}
            _ = v10
        case "hello":
            switch 42 {
            case -641:
                var v11 uint = 2
                _ = v11
            case 32768:
                fmt.Println(v1)
            default:
                var v12 uint32 = 2
                _ = v12
            }
        default:
            if false {
                var v13 int = 32767
                v13 -= 42
                _ = v13
            }
        }
    }
    v14 := []int{-469, 2147483647, 321, 42}
    var v15 uint64 = 608
    for fi16 := 0; fi16 < 11; fi16++ {
        switch ("\n\t" + "a") {
        case "a":
            switch 564 {
            case -472:
                v17 := GenS2{F0: 510, F1: 2}
                _ = v17
            case -78:
                v18 := GenS2{F0: 283, F1: 2}
                _ = v18
            default:
                v19 := GenS1{F0: 255, F1: (((v5 / 3.14159) - v5) < v5), F2: 255}
                _ = v19
            }
        case "hello":
            v3 -= -185
        }
        v5 += v5
        _ = fi16
    }
    fmt.Println(v3, v9, v1)
    v20 := GenS3{F0: 428, F1: 100, F2: "a"}
    switch -9223372036854775808 {
    case 0:
        for fi21 := 0; fi21 < 7; fi21++ {
            if false {
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 int32 = -54
                    _ = fi22
                    _ = v23
                }
                var v24 float64 = ((3.14159 * 1e10) - 3.14159)
                fmt.Println(v1, v9, v5)
                _ = v24
            }
            _ = fi21
        }
    }
    if (!true) {
        var v25 uint64 = v15
        switch "" {
        case "foo bar":
            v9 = v9
        case "":
            for fi26 := 0; fi26 < 6; fi26++ {
                switch fi26 {
                case 32767:
                    switch "hello" {
                    case "\n\t":
                        for fi27 := 0; fi27 < 8; fi27++ {
                            v28 := v2[((("foo bar" + "hello") + "hello") + ("a" + ("hello" + "")))]
                            if ((fi27 - fi27) <= -9223372036854775808) {
                                var v29 int8 = -128
                                _ = v29
                            }
                            fmt.Println(fi26, v3)
                            _ = fi27
                            _ = v28
                        }
                    default:
                        for fi30 := 0; fi30 < 13; fi30++ {
                            v14 = append(v14, (fi26 | fi30))
                            var v31 uint = 782
                            var v32 int64 = (-(v3))
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    }
                case 0:
                    v25 = v25
                case 186:
                    v33 := GenS3{F0: (v1 / (((v9 ^ (v1 * (v1 * 42)))) | 1)), F1: 0, F2: ("a" + "a")}
                    _ = v33
                }
                for fi34 := 0; fi34 < 5; fi34++ {
                    var v35 float64 = 0.1
                    v36 := GenS3{F0: (v9 % ((v9) | 1)), F1: 42, F2: "a"}
                    var v37 float64 = v35
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                var v38 int32 = -2147483648
                _ = fi26
                _ = v38
            }
        case "hello":
            v15 = v25
        }
        var v39 int32 = (v1 ^ -462)
        _ = v25
        _ = v39
    } else {
        fmt.Println(v1, v5)
        if (!true) {
            for fi40 := 0; fi40 < 13; fi40++ {
                if true {
                    if (42 == 1) {
                        switch "foo bar" {
                        case "foo bar":
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        }
                    } else {
                        v41 := GenS1{F0: 1, F1: false, F2: 1}
                        v3 = 617
                        _ = v41
                    }
                    fmt.Println(v5)
                } else {
                    if true {
                        v5 = v5
                    } else {
                        v42 := GenS4{F0: 100, F1: (v3 << ((1) & 63))}
                        fmt.Println(v1)
                        fmt.Println(v3, v1)
                        _ = v42
                    }
                    if (9223372036854775807 != -673) {
                        var v43 int8 = -59
                        _ = v43
                    }
                    fmt.Println(v3)
                }
                v44 := GenS3{F0: 1, F1: 42, F2: "foo bar"}
                _ = fi40
                _ = v44
            }
            var v45 uint8 = 184
            v46 := GenS2{F0: 42, F1: (v45 & (v45 % (((v45 << ((255) & 63))) | 1)))}
            _ = v45
            _ = v46
        } else {
            var v47 uint16 = 32767
            var v48 uint8 = 161
            v49 := v2[("hello" + ("foo bar" + ("\n\t" + "foo bar")))]
            _ = v47
            _ = v48
            _ = v49
        }
    }
    var v50 int64 = (v3 - v3)
    switch "" {
    case "a":
        if (0 != -32768) {
            var v51 int = 128
            if false {
                v51 += v51
                fmt.Println(v50)
                if len(v14) > 0 {
                    v14[0] = -191
                }
            } else {
                var v52 uint16 = uint16(v15)
                v5 = float32(0.1)
                _ = v52
            }
            v53 := GenS1{F0: 0, F1: (734 == 4294967295), F2: 65535}
            _ = v51
            _ = v53
        } else {
            switch 32768 {
            case 42:
                v54 := GenS3{F0: (v1 >> ((1) & 63)), F1: -5, F2: "foo bar"}
                _ = v54
            case 9223372036854775807:
                var v55 int64 = (^(v50))
                _ = v55
            default:
                for fi56 := 0; fi56 < 15; fi56++ {
                    for fi57 := 0; fi57 < 6; fi57++ {
                        v2["foo bar"] = 2
                        v5 -= 3.14159
                        v58 := GenS3{F0: int32(fi57), F1: int8(v15), F2: ""}
                        v3 = (v50 - (v3 * 666))
                        _ = fi57
                        _ = v58
                    }
                    _ = fi56
                }
            }
        }
    default:
        var v59 uint32 = 325
        _ = v59
    }
    if ("" == "foo bar") {
        for fi60 := 0; fi60 < 10; fi60++ {
            v1 = (v9 << ((0) & 63))
            var v61 int32 = 247
            if true {
                v9 = int32(v15)
                v4[2147483648] = "foo bar"
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            _ = fi60
            _ = v61
        }
        var v62 bool = false
        v9 = 2147483647
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v9
}
