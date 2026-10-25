package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 string; F1 uint; F2 string }
type GenS3 struct { F0 int64; F1 int32; F2 int }
type GenS4 struct { F0 uint8; F1 int64; F2 uint }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 uint }
func genFn0() int8 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint = 65535
        if false {
            v7 = (^(v7))
            v8 := []int{fi6, 32767, fi6}
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    v9 := []int{32767, 2147483647, 888}
    v10 := GenS1{F0: -108, F1: -2147483648}
    if true {
        switch 42 {
        case 1:
            v11 := GenS1{F0: -645, F1: -1}
            _ = v11
        }
        _ = len(v9)
    }
    var v12 uint8 = 128
    _ = v10
    _ = v12
    _ = v9
    return 42
}
func genFn1() int16 {
    v13 := GenS1{F0: -1, F1: 2147483647}
    v14 := []int{1, 128}
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.0
    var v3 int8 = 100
    var v4 int8 = (v3 & v3)
    var v5 string = "hello"
    v2 = (1.5 - ((((((float64((3.14159 * float32(((v2 + v2) + float64(3.14159))))) * v2) - (float64(-2.25) - v2)) * (0.0 - v2)) - -2.25) * ((v2 + 1.5) + v2)) + ((v2 - v2) - (v2 + (0.0 * float64(0.1))))))
    v15 := []int{-9223372036854775808, 42, 842, 32768}
    switch v5 {
    case "a":
        v4 = (v3 * v3)
    case "":
        for fi16 := 0; fi16 < 19; fi16++ {
            v2 = (v2 * v2)
            fmt.Println(v4, v2, fi16)
            switch (fi16 - fi16) {
            case 32767:
                if true {
                    switch ((v5 + "hello") + ((v5 + v5) + "hello")) {
                    case "\n\t":
                        for fi17 := 0; fi17 < 18; fi17++ {
                            var v18 string = (("foo bar" + "a") + ((("\n\t" + "\n\t") + (v5 + "foo bar")) + "\n\t"))
                            var v19 int = -779
                            var v20 uint = 4294967295
                            switch "foo bar" {
                            case "a":
                                var v21 uint32 = 0
                                _ = v21
                            default:
                                v22 := GenS1{F0: fi16, F1: 40}
                                _ = v22
                            }
                            _ = fi17
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    case "":
                        v23 := v1[("\n\t" + (v5 + v5))]
                        _ = v23
                    }
                } else {
                    var v24 int64 = 0
                    v5 = ("hello" + ("\n\t" + (v5 + v5)))
                    _ = v24
                }
            case 579:
                var v25 int8 = (-(v4))
                _ = v25
            case 902:
                if ((((v2 / v2) - v2) + float64(0.0)) == (v2 * v2)) {
                    v2 -= 1e10
                    v26 := GenS1{F0: 32768, F1: 0}
                    if false {
                        fmt.Println(v3, v5, v4)
                        fmt.Println(v4, v3)
                        v5 = v5
                    } else {
                        if (0 > 127) {
                            var v27 int16 = -741
                            if true {
                                if false {
                                    if ((fi16 % ((fi16) | 1)) == (fi16 >> ((496) & 63))) {
                                        var v28 float64 = float64((0.0 * 1.5))
                                        _ = v28
                                    } else {
                                        if false {
                                            fmt.Println(v5)
                                            if len(v15) > 0 {
                                                v15[0] = (fi16 | fi16)
                                            }
                                            if (1 > -531) {
                                                var v29 int = (-(fi16))
                                                _ = v29
                                            } else {
                                                v5 = v5
                                                v3 -= (v3 >> ((4294967295) & 63))
                                                for fi30 := 0; fi30 < 16; fi30++ {
                                                    if (true || true) {
                                                        fmt.Println(fi16)
                                                        v5 = ("hello" + ("a" + (v5 + "")))
                                                        if false {
                                                            v31 := GenS1{F0: fi16, F1: -267}
                                                            var v32 int8 = v4
                                                            if (false || false) {
                                                                switch fi30 {
                                                                case 128:
                                                                    v33 := GenS2{F0: ("a" + v5), F1: 65535, F2: ((v5 + (("hello" + (v5 + v5)) + (v5 + ("" + v5)))) + v5)}
                                                                    _ = v33
                                                                case 192:
                                                                    var v34 uint32 = 140
                                                                    _ = v34
                                                                }
                                                            }
                                                            _ = v31
                                                            _ = v32
                                                        }
                                                        v3 = v4
                                                    } else {
                                                        fmt.Println(v3)
                                                    }
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                    fmt.Println(v27, v5, v5)
                                                    _ = fi30
                                                }
                                            }
                                        } else {
                                            var v35 bool = (3.14159 != (float32(v2) + 1.5))
                                            switch v5 {
                                            case "hello":
                                                switch fi16 {
                                                case 32767:
                                                    var v36 int32 = 42
                                                    _ = v36
                                                default:
                                                    v37 := v1[v5]
                                                    _ = v37
                                                }
                                            case "a":
                                                if v35 {
                                                    switch ("a" + v5) {
                                                    case "hello":
                                                        var v38 float32 = float32((0.1 * 1.5))
                                                        _ = v38
                                                    default:
                                                        var v39 int16 = 464
                                                        _ = v39
                                                    }
                                                    var v40 int32 = -1
                                                    for fi41 := 0; fi41 < 12; fi41++ {
                                                        _ = len(v1)
                                                        v1[v5] = fi16
                                                        fmt.Println(v2, v40, v2)
                                                        _ = fi41
                                                    }
                                                    _ = v40
                                                } else {
                                                    var v42 uint8 = 127
                                                    if len(v15) >= 2 {
                                                        v15 = v15[1:len(v15)]
                                                    }
                                                    _ = v42
                                                }
                                            case "\n\t":
                                                var v43 string = "hello"
                                                _ = v43
                                            default:
                                                v35 = (248 < 255)
                                            }
                                            if v35 {
                                                var v44 int64 = -272
                                                var v45 int32 = -2147483648
                                                _ = v44
                                                _ = v45
                                            }
                                            _ = v35
                                        }
                                        for fi46 := 0; fi46 < 7; fi46++ {
                                            var v47 int16 = (-(v27))
                                            _ = fi46
                                            _ = v47
                                        }
                                    }
                                    v4 += v4
                                    if true {
                                        _ = len(v1)
                                    } else {
                                        for fi48 := 0; fi48 < 5; fi48++ {
                                            v49 := v1["a"]
                                            v49 -= 4294967295
                                            v50 := GenS3{F0: 42, F1: 2147483647, F2: (v49 >> ((410) & 63))}
                                            v49 = 9223372036854775807
                                            _ = fi48
                                            _ = v49
                                            _ = v50
                                        }
                                        var v51 int8 = (v3 ^ v4)
                                        _ = v51
                                    }
                                } else {
                                    var v52 int16 = -1
                                    switch "" {
                                    case "":
                                        _ = len(v15)
                                    case "\n\t":
                                        if ((v3 + (-(v3))) < genFn0()) {
                                            var v53 int32 = -977
                                            v54 := GenS2{F0: (("\n\t" + (v5 + ((v5 + "\n\t") + v5))) + v5), F1: 0, F2: v5}
                                            _ = v53
                                            _ = v54
                                        }
                                    default:
                                        for fi55 := 0; fi55 < 1; fi55++ {
                                            v52 -= (v52 + (v27 * (v27 << ((1) & 63))))
                                            v56 := GenS4{F0: 255, F1: 209, F2: 255}
                                            _ = fi55
                                            _ = v56
                                        }
                                    }
                                    _ = v52
                                }
                                if true {
                                    var v57 uint16 = 1
                                    _ = v57
                                }
                                v4 -= int8(v27)
                            }
                            _ = v27
                        }
                        if ((171 != -52) && (255 >= 161)) {
                            if (!(0.0 <= v2)) {
                                _ = len(v1)
                                var v58 int8 = 42
                                _ = v58
                            } else {
                                v59 := GenS4{F0: 42, F1: 9223372036854775807, F2: 495}
                                v3 = v3
                                for fi60 := 0; fi60 < 19; fi60++ {
                                    v3 = v3
                                    _ = len(v1)
                                    _ = len(v15)
                                    _ = fi60
                                }
                                _ = v59
                            }
                        } else {
                            var v61 int64 = 397
                            if true {
                                var v62 int8 = (v3 * v3)
                                _ = v62
                            } else {
                                var v63 float32 = 0.1
                                for fi64 := 0; fi64 < 4; fi64++ {
                                    v65 := GenS5{F0: 600, F1: 0}
                                    _ = fi64
                                    _ = v65
                                }
                                v3 = v3
                                _ = v63
                            }
                            _ = v61
                        }
                        if (255 > 42) {
                            v3 = (v3 >> ((18446744073709551615) & 63))
                        } else {
                            v66 := GenS6{F0: 0}
                            _ = v66
                        }
                    }
                    _ = v26
                } else {
                    var v67 string = v5
                    var v68 float32 = (((float32(float64(float32((v2 + float64(3.14159))))) * -2.25) + -2.25) - 3.14159)
                    v67 = "foo bar"
                    _ = v67
                    _ = v68
                }
            default:
                if (0 >= 159) {
                    v3 = (v3 ^ -66)
                    for fi69 := 0; fi69 < 19; fi69++ {
                        v3 += (v4 / ((v3) | 1))
                        _ = fi69
                    }
                }
            }
            _ = fi16
        }
    }
    fmt.Println(v5, v3)
    delete(v1, (("hello" + v5) + v5))
    if (!(!true)) {
        for fi70 := 0; fi70 < 9; fi70++ {
            if false {
                var v71 uint16 = 556
                v72 := GenS4{F0: 1, F1: 0, F2: 205}
                _ = v71
                _ = v72
            } else {
                fmt.Println(v3, v4)
                v3 += -1
                var v73 float32 = 1e10
                _ = v73
            }
            for fi74 := 0; fi74 < 20; fi74++ {
                v4 += (v4 % ((0) | 1))
                v75 := GenS7{F0: 65535}
                if (255 <= 1) {
                    var v76 int8 = 1
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    if (true || true) {
                        _ = len(v15)
                        var v77 int32 = 1
                        v15 = append(v15, (fi70 / ((fi74) | 1)))
                        _ = v77
                    } else {
                        if len(v15) > 0 {
                            v15[0] = fi74
                        }
                        v2 -= (v2 + (v2 + (v2 - ((float64((1e10 + 1e10)) * v2) + (float64(0.0) / 1e10)))))
                        v15 = append(v15, (^(fi70)))
                        fmt.Println(v4)
                    }
                    if ((fi70 / (((fi74 + fi74)) | 1)) != (fi70 ^ (fi70 - -9223372036854775808))) {
                        switch (fi70 >> ((255) & 63)) {
                        case 42:
                            v78 := GenS8{F0: 51, F1: 42}
                            _ = v78
                        case -2147483648:
                            if (!(255 < 441)) {
                                v4 -= (v76 * (v3 * (v76 << ((4294967295) & 63))))
                            }
                        case -9223372036854775808:
                            v79 := GenS5{F0: -585, F1: 42}
                            _ = v79
                        }
                    }
                    var v80 int64 = 9223372036854775807
                    v2 += 0.1
                    _ = v76
                    _ = v80
                } else {
                    v3 += v3
                    v81 := GenS2{F0: (((("hello" + ((v5 + ((("foo bar" + ("" + v5)) + "hello") + v5)) + "hello")) + (v5 + v5)) + "foo bar") + v5), F1: 424, F2: v5}
                    _ = v81
                }
                _ = fi74
                _ = v75
            }
            v3 -= v3
            _ = fi70
        }
        for fi82 := 0; fi82 < 1; fi82++ {
            delete(v1, v5)
            v5 = v5
            _ = fi82
        }
        delete(v1, (v5 + ""))
    } else {
        v5 = v5
    }
    if true {
        v3 = (v4 * 127)
    }
    for fi83 := 0; fi83 < 6; fi83++ {
        v3 = -101
        if false {
            v84 := GenS9{F0: 1}
            _ = v84
        }
        v4 -= -36
        _ = fi83
    }
    var v85 float64 = v2
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
}
