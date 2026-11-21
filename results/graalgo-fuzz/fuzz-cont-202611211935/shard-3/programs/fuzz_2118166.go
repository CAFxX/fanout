package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int16; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int8; F2 string }
type GenS6 struct { F0 uint8; F1 bool; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    v2 += 127
    var v4 uint64 = 705
    var v5 int8 = 42
    v6 := v1[127]
    switch -241 {
    case 32768:
        v2 = 163
    case -132:
        v4 -= 18446744073709551615
    default:
        for fi7 := 0; fi7 < 4; fi7++ {
            switch fi7 {
            case 475:
                var v8 uint = 255
                _ = v8
            case 42:
                if ((!((fi7 < (fi7 % ((fi7) | 1))) || (((v6 + ("" + v6)) + "hello") >= v6))) || true) {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        switch fi7 {
                        case 42:
                            v5 -= v5
                        case 1:
                            var v10 uint8 = (v2 + v2)
                            _ = v10
                        default:
                            v4 = v4
                        }
                        var v11 uint16 = 255
                        fmt.Println(v11, fi9, v11)
                        if false {
                            fmt.Println(fi9)
                        } else {
                            fmt.Println(fi9)
                        }
                        _ = fi9
                        _ = v11
                    }
                    var v12 uint32 = 601
                    v13 := GenS1{F0: v4, F1: (580 == 351)}
                    _ = v12
                    _ = v13
                }
            case -1:
                var v14 int = (fi7 & (fi7 >> ((190) & 63)))
                _ = v14
            default:
                switch v6 {
                case "\n\t":
                    v15 := []int{}
                    _ = v15
                case "a":
                    v4 = (v4 >> ((42) & 63))
                case "hello":
                    var v16 int8 = v5
                    _ = v16
                }
            }
            v17 := GenS1{F0: (v4 << ((2) & 63)), F1: (true || ((4294967295 <= 65535) || false))}
            fmt.Println(v4)
            _ = fi7
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        v5 = 0
        _ = fi18
    }
    if false {
        v5 = (v5 >> ((694) & 63))
        v6 = v6
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        if true {
            for fi20 := 0; fi20 < 11; fi20++ {
                fmt.Println(v5, fi20, fi20)
                switch fi20 {
                case 0:
                    var v21 uint64 = v4
                    _ = v21
                case 32768:
                    v5 -= v5
                case -1:
                    if false {
                        if (65535 < 255) {
                            v22 := GenS1{F0: v4, F1: true}
                            _ = v22
                        } else {
                            if (fi19 > 4294967295) {
                                var v23 bool = false
                                _ = v23
                            } else {
                                if ("hello" == "hello") {
                                    var v24 int8 = 100
                                    _ = v24
                                }
                                var v25 string = (v6 + ("foo bar" + "foo bar"))
                                _ = v25
                            }
                            delete(v1, fi20)
                        }
                        fmt.Println(v4, fi20)
                        v5 = -9
                        v26 := GenS2{F0: 1000, F1: (1 > -1), F2: (v4 / ((2) | 1))}
                        _ = v26
                    } else {
                        v2 -= 127
                        v4 += v4
                        if true {
                            switch v6 {
                            case "\n\t":
                                if false {
                                    if ((3.14159 < float64(((0.0 * (1.5 * 0.1)) / 1.0))) || (3.14159 <= 1.5)) {
                                        if (-412 != -621) {
                                            var v27 uint16 = 427
                                            var v28 uint8 = (v2 << ((18446744073709551615) & 63))
                                            fmt.Println(v2)
                                            var v29 float32 = 3.14159
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        } else {
                                            if ((!(!(-32768 > -849))) || ((!false) && false)) {
                                                v30 := GenS3{F0: 2, F1: v2}
                                                v4 = 2
                                                v5 = (v5 % ((v5) | 1))
                                                _ = v30
                                            } else {
                                                var v31 uint8 = 253
                                                v2 = (v2 / ((1) | 1))
                                                _ = v31
                                            }
                                            var v32 uint = 42
                                            _ = v32
                                        }
                                        switch (fi20 << ((975) & 63)) {
                                        case 1:
                                            v2 -= v2
                                        }
                                        if true {
                                            v2 = v2
                                        } else {
                                            var v33 uint16 = 42
                                            if ((!true) || (1000 <= 0)) {
                                                v2 = (^(v2))
                                                v34 := GenS3{F0: v33, F1: (v2 ^ 128)}
                                                var v35 float64 = (1.5 / (0.0 - float64(3.14159)))
                                                v36 := GenS4{F0: ((-2.25 * float32((float64(1e10) / float64(1e10)))) >= (((1e10 / 1.5) - float32((v35 / v35))) * (1.5 * 1.5))), F1: 1}
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                            } else {
                                                var v37 uint8 = v2
                                                _ = v37
                                            }
                                            _ = v33
                                        }
                                    } else {
                                        v5 = (v5 ^ -96)
                                    }
                                }
                                var v38 int = (^(fi20))
                                v39 := v1[583]
                                _ = v38
                                _ = v39
                            default:
                                switch v6 {
                                case "foo bar":
                                    switch "foo bar" {
                                    case "":
                                        if true {
                                            v4 += v4
                                        } else {
                                            var v40 float64 = 3.14159
                                            v41 := GenS5{F0: 2, F1: -1, F2: v6}
                                            var v42 int16 = 381
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                        }
                                    case "\n\t":
                                        v43 := GenS6{F0: v2, F1: true, F2: 42}
                                        _ = v43
                                    case "hello":
                                        var v44 float32 = (1e10 * 0.1)
                                        v45 := []int{}
                                        v2 = (v2 >> ((18446744073709551615) & 63))
                                        _ = v44
                                        _ = v45
                                    default:
                                        if (983 != 355) {
                                            v46 := []int{}
                                            v4 += 34
                                            v2 = (v2 << ((18446744073709551615) & 63))
                                            _ = v46
                                        } else {
                                            switch v6 {
                                            case "\n\t":
                                                var v47 uint16 = 65535
                                                _ = v47
                                            case "hello":
                                                if false {
                                                    if (true && false) {
                                                        fmt.Println(v5, v5)
                                                        v48 := GenS5{F0: 826, F1: v5, F2: v6}
                                                        _ = v48
                                                    }
                                                } else {
                                                    v4 = v4
                                                    v49 := GenS2{F0: 32767, F1: true, F2: v4}
                                                    _ = v49
                                                }
                                            case "a":
                                                var v50 uint32 = 797
                                                switch (v6 + ("a" + v6)) {
                                                case "a":
                                                    delete(v1, fi19)
                                                    if true {
                                                        fmt.Println(fi19)
                                                        v51 := GenS4{F0: (!(625 < v50)), F1: 2}
                                                        if false {
                                                            fmt.Println(v5, v6)
                                                            var v52 string = ("a" + v6)
                                                            var v53 uint64 = (^(v4))
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                        _ = v51
                                                    } else {
                                                        v50 = (^(v50))
                                                    }
                                                case "hello":
                                                    delete(v3, v50)
                                                case "":
                                                    if ((9223372036854775807 != (fi20 % ((128) | 1))) && false) {
                                                        if true {
                                                            v4 = v4
                                                            if (int32(fi20) > 911) {
                                                                var v54 int16 = 0
                                                                v55 := GenS5{F0: 2, F1: v5, F2: v6}
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                        } else {
                                                            var v56 int32 = 2147483647
                                                            v57 := v3[(v50 ^ (^(v50)))]
                                                            fmt.Println(fi19)
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                    }
                                                }
                                                switch (v6 + (v6 + "a")) {
                                                case "\n\t":
                                                    v58 := []int{(fi19 ^ fi19), fi20, -9223372036854775808, 9223372036854775807}
                                                    _ = v58
                                                }
                                                _ = v50
                                            default:
                                                v59 := []int{128, fi19, (fi19 - (fi19 << ((42) & 63)))}
                                                _ = v59
                                            }
                                            v60 := GenS2{F0: -394, F1: true, F2: 210}
                                            _ = v60
                                        }
                                    }
                                case "":
                                    _ = len(v3)
                                case "\n\t":
                                    v1[(fi19 + (fi19 >> ((18446744073709551615) & 63)))] = ("a" + "\n\t")
                                default:
                                    if true {
                                        var v61 bool = (128 < v2)
                                        _ = v61
                                    }
                                }
                            }
                        } else {
                            var v62 string = "\n\t"
                            v63 := GenS3{F0: uint16(v2), F1: (v2 + 128)}
                            v64 := v1[0]
                            _ = v62
                            _ = v63
                            _ = v64
                        }
                    }
                    var v65 uint32 = 2
                    v65 -= (v65 | (v65 - v65))
                    delete(v3, (v65 & (v65 + (v65 - (^(v65))))))
                    _ = v65
                }
                _ = fi20
            }
            v66 := []int{}
            _ = v66
        }
        switch v6 {
        case "a":
            var v67 int8 = 42
            _ = v67
        case "foo bar":
            fmt.Println(v5, v6, v4)
        default:
            v68 := []int{(^(fi19)), fi19, (fi19 - (fi19 - fi19)), 1}
            _ = v68
        }
        _ = fi19
    }
    delete(v1, -1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
