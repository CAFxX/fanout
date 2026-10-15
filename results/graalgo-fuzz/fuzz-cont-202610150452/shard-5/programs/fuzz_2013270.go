package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 int8 }
type GenS3 struct { F0 int8; F1 int32; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 bool; F2 uint32 }
type GenS6 struct { F0 int8; F1 string; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint8 }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 8; fi5++ {
        if true {
            v6 := []int{(fi5 << ((0) & 63))}
            if ((fi5 << ((255) & 63)) != (-(fi5))) {
                var v7 uint8 = 255
                _ = v7
            }
            fmt.Println(fi5)
            _ = v6
        }
        var v8 uint8 = 127
        _ = fi5
        _ = v8
    }
    var v9 uint8 = 1
    for fi10 := 0; fi10 < 7; fi10++ {
        v9 = v9
        v11 := []int{(fi10 - fi10), (^(fi10)), fi10}
        if (-1 <= 30) {
            for fi12 := 0; fi12 < 10; fi12++ {
                v9 = (v9 & v9)
                switch fi12 {
                case -232:
                    switch 1 {
                    case -405:
                        v13 := GenS1{F0: 42, F1: 42}
                        _ = v13
                    case 102:
                        var v14 uint8 = 128
                        var v15 uint = 0
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        _ = v14
                        _ = v15
                    case 239:
                        v16 := GenS2{F0: 4294967295, F1: 1000, F2: int8(fi12)}
                        _ = v16
                    default:
                        switch fi12 {
                        case -427:
                            var v17 bool = (!true)
                            _ = v17
                        case 32767:
                            v9 += v9
                        case -451:
                            var v18 int32 = 0
                            _ = v18
                        default:
                            if len(v11) > 0 {
                                v11[0] = fi10
                            }
                        }
                    }
                case 963:
                    var v19 uint16 = 42
                    _ = v19
                case 433:
                    if (878 != 9223372036854775807) {
                        fmt.Println(fi10, fi10)
                    }
                }
                if true {
                    _ = len(v11)
                    if true {
                        v9 -= (v9 << ((65535) & 63))
                        v9 = v9
                    }
                    var v20 float32 = (float32(0.0) / 1.0)
                    if (false || (42 < 1)) {
                        if (120 == (^(v9))) {
                            if true {
                                v20 += -2.25
                                var v21 uint16 = 42
                                fmt.Println(v9, v20, v20)
                                v21 += (v21 ^ (v21 - (v21 ^ 828)))
                                _ = v21
                            }
                            if (!false) {
                                var v22 int = fi10
                                var v23 string = ("hello" + "hello")
                                _ = v22
                                _ = v23
                            } else {
                                v20 = 3.14159
                                if false {
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    if false {
                                        var v24 string = ("hello" + ("\n\t" + ("a" + "a")))
                                        var v25 float32 = v20
                                        if true {
                                            v26 := GenS3{F0: 41, F1: 1, F2: 2}
                                            _ = v26
                                        } else {
                                            fmt.Println(fi12, v9, v24)
                                            switch 2147483647 {
                                            case -1:
                                                var v27 float64 = ((-2.25 + (1.5 / 0.1)) - -2.25)
                                                _ = v27
                                            case -530:
                                                v9 = v9
                                            case 32768:
                                                v28 := GenS4{F0: ((((v24 + v24) + v24) + ((v24 + v24) + "a")) + v24)}
                                                _ = v28
                                            }
                                        }
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v20 = float32(0.1)
                                    }
                                    switch ("" + "") {
                                    case "\n\t":
                                        if false {
                                            if (false && false) {
                                                v9 = v9
                                            } else {
                                                switch "\n\t" {
                                                case "hello":
                                                    v9 = 0
                                                case "a":
                                                    _ = len(v11)
                                                case "foo bar":
                                                    var v29 uint = 0
                                                    _ = v29
                                                }
                                                v9 = (v9 >> ((4294967295) & 63))
                                                if false {
                                                    fmt.Println(fi10, v9, v20)
                                                } else {
                                                    fmt.Println(fi10)
                                                    var v30 int16 = -32768
                                                    v30 = (v30 ^ v30)
                                                    _ = v30
                                                }
                                                if false {
                                                    v31 := GenS5{F0: 1, F1: true, F2: 105}
                                                    _ = v31
                                                } else {
                                                    var v32 int32 = -781
                                                    v33 := GenS6{F0: 61, F1: ("a" + "hello"), F2: 2}
                                                    v34 := GenS7{F0: 32767, F1: v9}
                                                    switch 160 {
                                                    case 4294967295:
                                                        fmt.Println(v32, fi10, v20)
                                                    }
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                }
                                                var v35 uint64 = 18446744073709551615
                                                _ = v35
                                            }
                                            v11 = append(v11, 127)
                                        }
                                    case "":
                                        if (!(false || true)) {
                                            var v36 float32 = 3.14159
                                            v11 = append(v11, fi10)
                                            if (-550 >= (fi12 << ((0) & 63))) {
                                                var v37 float32 = (1.5 / v36)
                                                v11 = append(v11, -2147483648)
                                                _ = v37
                                            }
                                            _ = v36
                                        } else {
                                            var v38 uint8 = (^(v9))
                                            _ = v38
                                        }
                                    }
                                    fmt.Println(v20)
                                }
                                var v39 bool = true
                                v39 = ((false || (1000 < 0)) && v39)
                                _ = v39
                            }
                            _ = len(v11)
                        }
                        switch fi12 {
                        case 4294967295:
                            switch "hello" {
                            case "":
                                switch "a" {
                                case "foo bar":
                                    var v40 uint8 = (v9 / ((v9) | 1))
                                    _ = v40
                                case "\n\t":
                                    if (!(false || true)) {
                                        v20 = v20
                                        if (((v20 * v20) != (((v20 * v20) * float32(3.14159)) / (3.14159 * v20))) && ((true && (42 == -1)) || false)) {
                                            var v41 bool = false
                                            _ = v41
                                        }
                                    }
                                }
                            case "a":
                                var v42 int = fi10
                                _ = v42
                            case "hello":
                                if (2147483647 >= uint32(v9)) {
                                    var v43 int8 = 127
                                    fmt.Println(v43, v43)
                                    v44 := GenS6{F0: 100, F1: "\n\t", F2: 4294967295}
                                    _ = v43
                                    _ = v44
                                }
                            default:
                                var v45 uint64 = 4294967295
                                _ = v45
                            }
                        default:
                            fmt.Println(v20, fi12, v9)
                        }
                    } else {
                        var v46 int16 = 0
                        _ = v46
                    }
                    var v47 uint32 = 0
                    _ = v20
                    _ = v47
                } else {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    if len(v11) > 0 {
                        v11[0] = 127
                    }
                    fmt.Println(v9, fi12)
                }
                _ = fi12
            }
            v9 = v9
        }
        _ = fi10
        _ = v11
    }
    if (true && (255 == (v9 - (v9 ^ v9)))) {
        v48 := GenS6{F0: -53, F1: "a", F2: 2147483648}
        v9 -= v9
        var v49 string = "a"
        _ = v48
        _ = v49
    }
    _ = v9
    return 2
}
func genFn1(p0 int16, p1 uint64, p2 int32) int64 {
    fmt.Println(p0, p2)
    if true {
        p2 = (p2 & -60)
        for fi50 := 0; fi50 < 19; fi50++ {
            if (162 >= 0) {
                switch fi50 {
                case 128:
                    p1 += p1
                case 246:
                    for fi51 := 0; fi51 < 12; fi51++ {
                        var v52 int16 = p0
                        switch 282 {
                        case 1:
                            v52 = p0
                        case 945:
                            v53 := []int{(fi50 - (^(fi50))), (fi51 + (fi50 * -2147483648))}
                            _ = v53
                        case -2147483648:
                            v52 = 32767
                        default:
                            p1 -= p1
                        }
                        p1 -= p1
                        _ = fi51
                        _ = v52
                    }
                case 0:
                    p2 = p2
                default:
                    var v54 float64 = 3.14159
                    _ = v54
                }
            }
            for fi55 := 0; fi55 < 8; fi55++ {
                v56 := GenS2{F0: 65535, F1: (p0 + p0), F2: -1}
                fmt.Println(fi50)
                _ = fi55
                _ = v56
            }
            for fi57 := 0; fi57 < 17; fi57++ {
                var v58 uint16 = 915
                _ = fi57
                _ = v58
            }
            _ = fi50
        }
        p1 = p1
    } else {
        p0 = (p0 + p0)
        if (!false) {
            v59 := GenS3{F0: 0, F1: (p2 >> ((0) & 63)), F2: 255}
            for fi60 := 0; fi60 < 3; fi60++ {
                v61 := GenS5{F0: 1000, F1: ((0.0 - 0.0) == 1.5), F2: 2}
                var v62 float32 = (0.1 + -2.25)
                for fi63 := 0; fi63 < 12; fi63++ {
                    var v64 float64 = 1e10
                    _ = fi63
                    _ = v64
                }
                _ = fi60
                _ = v61
                _ = v62
            }
            _ = v59
        } else {
            p0 = p0
            for fi65 := 0; fi65 < 8; fi65++ {
                if true {
                    p0 += p0
                    fmt.Println(fi65, p2, p2)
                    p0 = (p0 & 808)
                } else {
                    p1 += p1
                    v66 := []int{(fi65 - fi65)}
                    var v67 int32 = -1
                    _ = v66
                    _ = v67
                }
                v68 := GenS1{F0: 39, F1: -9223372036854775808}
                var v69 bool = false
                _ = fi65
                _ = v68
                _ = v69
            }
        }
        if ((false && (!false)) && (!false)) {
            v70 := []int{127}
            var v71 uint64 = p1
            _ = v70
            _ = v71
        } else {
            v72 := []int{127}
            var v73 uint = 986
            _ = v72
            _ = v73
        }
    }
    if (2147483647 != 0) {
        for fi74 := 0; fi74 < 4; fi74++ {
            for fi75 := 0; fi75 < 11; fi75++ {
                var v76 uint32 = 489
                v77 := GenS6{F0: 1, F1: ("\n\t" + ("hello" + ("foo bar" + ((("a" + "\n\t") + "\n\t") + "a")))), F2: 2147483647}
                _ = fi75
                _ = v76
                _ = v77
            }
            _ = fi74
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -973
}

func main() {
    v1 := []string{}
    var v2 int = 1
    v3 := make(map[int]string)
    var v4 string = ("" + ("foo bar" + "foo bar"))
    var v78 int8 = 81
    var v79 int16 = -412
    v2 = -2147483648
    v80 := GenS2{F0: 2147483647, F1: 1, F2: v78}
    for fi81 := 0; fi81 < 2; fi81++ {
        fmt.Println(v78, v2)
        var v82 int8 = v78
        _ = fi81
        _ = v82
    }
    v83 := GenS7{F0: v79, F1: 0}
    v84 := GenS7{F0: (-(v79)), F1: 97}
    v2 = (v2 ^ v2)
    v85 := GenS5{F0: v79, F1: (202 < 4294967295), F2: 2}
    for fi86 := 0; fi86 < 12; fi86++ {
        fmt.Println(v2, v78, v4)
        v78 = v78
        v79 += (v79 >> ((194) & 63))
        _ = fi86
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v78
    _ = v79
    _ = v80
    _ = v83
    _ = v84
    _ = v85
}
