package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 uint }

func main() {
    v1 := []string{}
    var v2 float64 = 1e10
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    var v5 float64 = v2
    for fi6 := 0; fi6 < 8; fi6++ {
        switch "" {
        case "a":
            var v7 int32 = -1
            _ = v7
        case "foo bar":
            switch 32767 {
            case -1:
                var v8 int8 = 42
                _ = v8
            case 9223372036854775807:
                v5 += (v5 * ((0.1 / float64(float32(1.5))) - (v5 - 3.14159)))
            case 836:
                v5 -= v2
            default:
                _ = len(v3)
            }
        case "hello":
            v5 -= float64((3.14159 + 3.14159))
        default:
            v2 = (0.0 / v5)
        }
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 uint16 = 2
            if (true || (65535 >= 722)) {
                var v11 uint8 = 255
                if (false || true) {
                    if true {
                        v12 := GenS1{F0: (fi9 - fi9), F1: fi9, F2: ""}
                        _ = v12
                    } else {
                        v10 = v10
                    }
                    v13 := GenS2{F0: v11, F1: 100, F2: 2147483647}
                    fmt.Println(fi9, v10, v10)
                    _ = v13
                } else {
                    var v14 uint = 42
                    v14 = (v14 << (((v14 - v14)) & 63))
                    _ = v14
                }
                var v15 uint = uint(v11)
                if true {
                    v16 := GenS1{F0: 127, F1: fi6, F2: "foo bar"}
                    _ = v16
                }
                var v17 bool = (((-675 == 0) || true) && (-1 <= 0))
                _ = v11
                _ = v15
                _ = v17
            }
            var v18 int = (^(fi6))
            _ = fi9
            _ = v10
            _ = v18
        }
        _ = fi6
    }
    v2 = 0.0
    if (((((float64(float32(float64((-2.25 / -2.25)))) * v5) - 1e10) + ((float64((1e10 - float32(v2))) - v5) / v2)) / 0.1) > -2.25) {
        v2 += (v5 - float64(float32((float64(1.5) - v2))))
        var v19 uint16 = 561
        v1 = append(v1, "hello")
        _ = v19
    }
    v20 := GenS1{F0: -1, F1: 177, F2: ""}
    v5 += v2
    v2 += 0.1
    if (89 == -1) {
        if ((3.14159 <= (((0.1 * (1e10 * float32(float64(float32((-2.25 + (1e10 / v2))))))) + (-2.25 + (((1e10 - 1.5) * 0.1) - 0.1))) - (float32(v5) - 3.14159))) || true) {
            v2 = (0.0 * (1e10 * (-2.25 * v5)))
        }
        v5 -= v5
        var v21 uint = 1
        _ = v21
    } else {
        switch -104 {
        case -807:
            switch -427 {
            case 133:
                delete(v4, 670)
            default:
                for fi22 := 0; fi22 < 18; fi22++ {
                    if false {
                        v2 -= (float64(1e10) + (v2 / v5))
                    }
                    _ = fi22
                }
            }
        default:
            v23 := GenS1{F0: 480, F1: 4294967295, F2: ""}
            _ = v23
        }
        fmt.Println(v5, v2)
    }
    v5 += v5
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            if (0.0 == (((0.0 - -2.25) + -2.25) * 0.1)) {
                if false {
                    v2 += -2.25
                } else {
                    v5 = v2
                    switch 899 {
                    case 969:
                        v26 := GenS3{F0: -714, F1: 32767}
                        _ = v26
                    case 948:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if true {
                        v27 := GenS4{F0: (317 >= 345)}
                        if true {
                            var v28 uint = 2
                            v29 := v4[389]
                            if true {
                                v28 = 0
                                switch "foo bar" {
                                case "foo bar":
                                    var v30 int32 = -2147483648
                                    _ = v30
                                case "":
                                    var v31 float64 = v5
                                    _ = v31
                                case "\n\t":
                                    var v32 int8 = 42
                                    _ = v32
                                default:
                                    var v33 int = fi24
                                    _ = v33
                                }
                            } else {
                                if false {
                                    v34 := GenS1{F0: fi24, F1: fi24, F2: v29}
                                    v5 += v2
                                    if (0.0 == 1.5) {
                                        var v35 float32 = ((1e10 + 0.1) / (-2.25 + (1.5 / 1.5)))
                                        fmt.Println(v29, v29)
                                        _ = v35
                                    } else {
                                        var v36 int32 = 2147483647
                                        _ = v36
                                    }
                                    _ = v34
                                }
                            }
                            _ = v28
                            _ = v29
                        }
                        var v37 int32 = 512
                        _ = v27
                        _ = v37
                    } else {
                        v38 := GenS4{F0: false}
                        switch fi24 {
                        case 127:
                            v2 = (v5 - float64(((1e10 / 1.5) - -2.25)))
                        case 128:
                            v39 := GenS4{F0: (!(0 < 2147483648))}
                            _ = v39
                        default:
                            v5 += 0.0
                        }
                        var v40 string = (("hello" + "foo bar") + ("hello" + "foo bar"))
                        _ = v38
                        _ = v40
                    }
                }
            } else {
                var v41 float64 = v2
                switch ("a" + ("a" + ("a" + "hello"))) {
                case "\n\t":
                    var v42 string = (("\n\t" + "\n\t") + ("foo bar" + "\n\t"))
                    if (42 < 42) {
                        v43 := GenS3{F0: 437, F1: 1}
                        _ = v43
                    } else {
                        v2 = -2.25
                        v42 = ("a" + v42)
                    }
                    v44 := GenS2{F0: 1, F1: 127, F2: 4294967295}
                    _ = v42
                    _ = v44
                case "hello":
                    var v45 int32 = 423
                    if false {
                        var v46 int32 = (v45 / ((-832) | 1))
                        v47 := GenS1{F0: 127, F1: (fi24 - fi24), F2: ("a" + "hello")}
                        _ = v46
                        _ = v47
                    }
                    _ = v45
                case "foo bar":
                    v48 := GenS1{F0: 104, F1: -144, F2: "\n\t"}
                    v2 -= (-2.25 / (0.1 - (((v5 * (v2 * 1e10)) + float64(((1.5 - -2.25) + -2.25))) + v5)))
                    _ = v48
                }
                v41 = (v41 + (((v2 + 1.5) + (v41 / 1.0)) - 0.1))
                _ = v41
            }
            var v49 float64 = 1.5
            v50 := GenS5{F0: ("hello" + ("foo bar" + "a")), F1: 2}
            _ = fi25
            _ = v49
            _ = v50
        }
        v2 += float64(1.5)
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
