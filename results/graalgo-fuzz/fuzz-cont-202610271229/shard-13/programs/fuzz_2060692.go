package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int8; F1 int8; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8; F1 bool; F2 int64 }

func main() {
    var v1 float32 = (float32((-2.25 - 1e10)) + ((1e10 * 1e10) / (0.1 + float32(-2.25))))
    var v2 float32 = ((float32(-2.25) / 0.1) - float32(3.14159))
    var v3 int64 = 1
    var v4 int16 = -32768
    v5 := GenS1{F0: -14, F1: (v3 + (v3 - v3))}
    v6 := GenS1{F0: -116, F1: (v3 + v3)}
    var v7 uint16 = 337
    var v8 uint64 = 18446744073709551615
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS2{F0: 127, F1: 1, F2: -1}
        for fi11 := 0; fi11 < 10; fi11++ {
            if (!true) {
                v1 += (v2 * v1)
                v1 += 0.1
            }
            v12 := GenS1{F0: (-(fi11)), F1: v3}
            _ = fi11
            _ = v12
        }
        switch (-(fi9)) {
        case -30:
            v13 := GenS2{F0: 100, F1: 1, F2: 0}
            _ = v13
        case 257:
            switch "foo bar" {
            case "\n\t":
                var v14 uint8 = 35
                _ = v14
            case "":
                v1 -= v2
            case "a":
                var v15 int16 = 725
                _ = v15
            }
        }
        _ = fi9
        _ = v10
    }
    v16 := GenS3{F0: 665, F1: (v8 << ((255) & 63))}
    for fi17 := 0; fi17 < 10; fi17++ {
        if ((!(4294967295 <= 4294967295)) && (v7 == (v7 * (v7 % (((v7 % (((^(v7))) | 1))) | 1))))) {
            v2 += v1
        }
        if ((float64(v1) <= 3.14159) && true) {
            for fi18 := 0; fi18 < 17; fi18++ {
                fmt.Println(v7)
                _ = fi18
            }
            v19 := GenS2{F0: -1, F1: 127, F2: 0}
            v20 := []int{fi17}
            _ = v19
            _ = v20
        } else {
            v21 := GenS2{F0: -128, F1: 100, F2: -2147483648}
            for fi22 := 0; fi22 < 1; fi22++ {
                v23 := GenS4{F0: 42}
                switch "a" {
                case "":
                    v24 := []int{42, (fi17 & (fi17 * (fi17 | 128)))}
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        fmt.Println(fi17, fi17)
        _ = fi17
    }
    if false {
        v4 = v4
        if ((567 <= 2147483648) || true) {
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 bool = ((42 <= -128) && (false || false))
                switch ("" + ((("hello" + "hello") + "foo bar") + "foo bar")) {
                case "hello":
                    if (-315 >= fi25) {
                        if v26 {
                            v27 := GenS3{F0: (v7 / (((^(v7))) | 1)), F1: v8}
                            if v26 {
                                v7 = (v7 | 0)
                            } else {
                                var v28 bool = true
                                v28 = v26
                                for fi29 := 0; fi29 < 13; fi29++ {
                                    v7 = v7
                                    _ = fi29
                                }
                                _ = v28
                            }
                            v7 = 255
                            _ = v27
                        } else {
                            v3 -= v3
                            fmt.Println(v7, v8, v2)
                            var v30 float32 = ((v2 - ((v1 + 1e10) * v1)) * 0.0)
                            _ = v30
                        }
                        v31 := GenS2{F0: -1, F1: -128, F2: 42}
                        var v32 float32 = v1
                        _ = v31
                        _ = v32
                    }
                case "\n\t":
                    v2 = 0.1
                default:
                    v33 := GenS5{F0: (((3.14159 / 3.14159) != v2) && (fi25 != int(v8)))}
                    _ = v33
                }
                _ = fi25
                _ = v26
            }
            v34 := GenS1{F0: -964, F1: int64(v7)}
            if false {
                for fi35 := 0; fi35 < 14; fi35++ {
                    v1 = float32(0.1)
                    _ = fi35
                }
                if false {
                    v36 := GenS6{F0: v4}
                    switch "hello" {
                    case "\n\t":
                        for fi37 := 0; fi37 < 2; fi37++ {
                            fmt.Println(fi37)
                            for fi38 := 0; fi38 < 8; fi38++ {
                                v39 := GenS1{F0: fi37, F1: v3}
                                _ = fi38
                                _ = v39
                            }
                            _ = fi37
                        }
                    case "":
                        var v40 uint8 = 0
                        _ = v40
                    default:
                        var v41 float64 = 0.0
                        _ = v41
                    }
                    _ = v36
                }
            } else {
                for fi42 := 0; fi42 < 20; fi42++ {
                    if true {
                        v2 -= v1
                    } else {
                        switch "" {
                        case "foo bar":
                            if (!false) {
                                var v43 uint32 = 2147483647
                                v44 := GenS3{F0: (v7 * 65535), F1: v8}
                                _ = v43
                                _ = v44
                            }
                        case "":
                            var v45 int8 = 42
                            _ = v45
                        default:
                            v8 -= 18446744073709551615
                        }
                        var v46 int64 = v3
                        _ = v46
                    }
                    _ = fi42
                }
                v47 := []int{128, 4294967295}
                _ = v47
            }
            _ = v34
        } else {
            v48 := GenS2{F0: -1, F1: 42, F2: int32(v4)}
            v7 = 65535
            _ = v48
        }
    }
    if (128 < 127) {
        if (true && false) {
            v4 -= v4
            var v49 int16 = v4
            _ = v49
        } else {
            v50 := GenS1{F0: 461, F1: v3}
            var v51 string = ((("hello" + "") + "hello") + ("foo bar" + "\n\t"))
            if false {
                fmt.Println(v51, v3)
            } else {
                v4 += -32768
            }
            _ = v50
            _ = v51
        }
        switch ("foo bar" + "hello") {
        case "foo bar":
            for fi52 := 0; fi52 < 12; fi52++ {
                fmt.Println(v1, v8, v3)
                v2 += 1e10
                _ = fi52
            }
        case "\n\t":
            v4 = v4
        case "a":
            for fi53 := 0; fi53 < 19; fi53++ {
                switch int(v7) {
                case 2147483647:
                    v4 -= 0
                case 128:
                    v3 -= v3
                case 42:
                    var v54 uint16 = v7
                    _ = v54
                default:
                    v2 = ((float32((1.5 + (((-2.25 * 1e10) * 0.0) / (float64(v1) + 1e10)))) / (v2 / 1e10)) / v2)
                }
                _ = fi53
            }
        default:
            var v55 int64 = 1
            _ = v55
        }
    } else {
        var v56 uint32 = 65535
        fmt.Println(v7, v2)
        switch ((((("" + "a") + "") + (("\n\t" + "a") + "\n\t")) + "") + "\n\t") {
        case "":
            v56 -= 487
        case "hello":
            if true {
                v57 := GenS7{F0: (^(v7))}
                _ = v57
            }
        case "a":
            if (205 < 54) {
                switch ("hello" + "hello") {
                case "a":
                    v8 = v8
                case "\n\t":
                    v7 += v7
                case "foo bar":
                    for fi58 := 0; fi58 < 19; fi58++ {
                        var v59 float64 = ((float64((1.5 / ((v2 - ((1e10 / v1) - v1)) * 1.5))) + (0.0 - ((0.0 * (0.1 * 0.1)) - 3.14159))) - (((((float64(v1) * (float64((((1.5 + -2.25) / ((v2 + (v2 - v1)) - float32(((0.1 + (((0.1 + (3.14159 / 1e10)) / (1e10 / 0.1)) * 0.1)) + (0.0 - (0.1 + (((float64(0.0) / 3.14159) - (float64((float32(3.14159) * (((v2 / 1e10) + v2) * 0.1))) * 3.14159)) - (((3.14159 * 1e10) / (((0.1 - 1.5) + ((3.14159 + ((float64(v2) * ((0.0 + 1.5) / float64((v2 / v1)))) - -2.25)) * 1e10)) * (float64((v2 * 1.5)) * 0.0))) + (((1.5 - -2.25) * (float64(v2) + float64(v1))) - float64(((((v2 - (v1 + -2.25)) / ((((((v2 + ((3.14159 - (v2 - (v2 / v1))) + (((0.1 + v1) + 1e10) / 0.1))) * v1) * float32(((-2.25 * (-2.25 + 1.5)) / float64(((v1 - ((v1 * float32(float64(v1))) / v1)) + v2))))) - (v2 + v2)) * ((v1 / v2) - v1)) - 3.14159)) - (v2 - v1)) + v2))))))))))) - (float32(0.0) / float32(-2.25)))) + 3.14159)) - ((1e10 / 3.14159) + 0.0)) * 0.0) - 0.1) * 0.1))
                        _ = fi58
                        _ = v59
                    }
                }
                v60 := []int{}
                _ = v60
            } else {
                switch 447 {
                case 2147483647:
                    switch 127 {
                    case -2147483648:
                        v8 += (^(v8))
                    }
                case 0:
                    v56 -= (v56 * (v56 ^ (^(v56))))
                case 127:
                    v61 := []int{32767, -872, -748, -9223372036854775808}
                    _ = v61
                default:
                    if false {
                        v62 := []int{2147483647, 1}
                        _ = v62
                    } else {
                        v63 := GenS5{F0: ((int(v8) != 32768) && false)}
                        for fi64 := 0; fi64 < 12; fi64++ {
                            v65 := []int{4294967295}
                            _ = len(v65)
                            _ = fi64
                            _ = v65
                        }
                        _ = v63
                    }
                }
                if true {
                    var v66 string = ""
                    var v67 float64 = 0.0
                    _ = v66
                    _ = v67
                } else {
                    if ((1 <= 985) && (2 >= 65535)) {
                        v68 := GenS5{F0: ((("" != "a") && false) && (true || (false || false)))}
                        v69 := []int{}
                        if true {
                            v4 = v4
                        }
                        _ = v68
                        _ = v69
                    } else {
                        v8 += (v8 / ((42) | 1))
                        var v70 bool = (true || true)
                        fmt.Println(v56, v2, v8)
                        _ = v70
                    }
                    fmt.Println(v1, v2)
                }
            }
        }
        _ = v56
    }
    var v71 float32 = v2
    switch -1 {
    case 0:
        var v72 string = "\n\t"
        _ = v72
    }
    if (255 < 255) {
        v71 -= v2
    } else {
        v73 := []int{-364, 0, 42, -2147483648}
        for fi74 := 0; fi74 < 17; fi74++ {
            if ((4 < 1) && (v71 >= v1)) {
                var v75 bool = (("\n\t" + ("" + "\n\t")) != "foo bar")
                v3 -= -9223372036854775808
                _ = v75
            }
            if (0.1 >= 1.5) {
                for fi76 := 0; fi76 < 11; fi76++ {
                    switch (fi76 << ((255) & 63)) {
                    case 0:
                        if false {
                            v7 -= v7
                        }
                    case 2147483647:
                        fmt.Println(v7, v3, v3)
                    }
                    var v77 int64 = 9223372036854775807
                    _ = fi76
                    _ = v77
                }
                if (1 <= 83) {
                    v8 = v8
                    v71 = v71
                }
                if ((2 != 2) && false) {
                    v78 := GenS8{F0: 2, F1: ((fi74 + fi74) > 477), F2: v3}
                    _ = v78
                }
            } else {
                v73 = append(v73, fi74)
                var v79 uint64 = 4294967295
                fmt.Println(fi74)
                _ = v79
            }
            _ = fi74
        }
        for fi80 := 0; fi80 < 3; fi80++ {
            fmt.Println(fi80)
            for fi81 := 0; fi81 < 17; fi81++ {
                v7 += v7
                _ = fi81
            }
            _ = fi80
        }
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v71:", v71)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v71
    _ = v8
}
