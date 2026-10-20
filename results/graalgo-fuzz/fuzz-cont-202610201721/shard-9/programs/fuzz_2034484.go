package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint64; F2 string }
type GenS7 struct { F0 int }

func main() {
    var v1 string = "a"
    v2 := make(map[int]string)
    var v3 float64 = -2.25
    var v4 int8 = 100
    v5 := []int{}
    v6 := make(map[string]uint32)
    if false {
        var v7 float32 = (-2.25 + 0.1)
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            if true {
                v5 = append(v5, -329)
                fmt.Println(fi8)
            } else {
                var v10 int64 = 820
                _ = v10
            }
            _ = fi9
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v11 uint16 = 319
        _ = fi8
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 float32 = -2.25
        var v14 uint = 18446744073709551615
        if ((v4 >> (((v14 % ((1) | 1))) & 63)) <= v4) {
            var v15 float64 = v3
            _ = v15
        } else {
            switch fi12 {
            case 9223372036854775807:
                var v16 uint16 = 1
                _ = v16
            }
            v4 = v4
            for fi17 := 0; fi17 < 1; fi17++ {
                v3 = (v3 * v3)
                _ = fi17
            }
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 -= v3
    v1 = v1
    v18 := GenS1{F0: 0}
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 12; fi20++ {
            if len(v5) > 0 {
                v5[0] = fi19
            }
            var v21 uint16 = 65535
            _ = len(v6)
            _ = fi20
            _ = v21
        }
        fmt.Println(v3, v1)
        _ = fi19
    }
    v4 -= 42
    v1 = v1
    if (2147483648 != 685) {
        fmt.Println(v1, v3, v3)
        var v22 int32 = int32(v4)
        _ = v22
    }
    switch (v1 + (("" + (v1 + v1)) + "foo bar")) {
    case "foo bar":
        switch (v1 + (("\n\t" + "foo bar") + ((("" + "hello") + v1) + v1))) {
        case "foo bar":
            var v23 uint64 = 532
            _ = v23
        case "hello":
            v24 := GenS1{F0: 2}
            _ = v24
        case "\n\t":
            v6["\n\t"] = 2
        }
    default:
        for fi25 := 0; fi25 < 13; fi25++ {
            v3 = v3
            _ = len(v5)
            _ = fi25
        }
    }
    var v26 uint16 = 568
    if false {
        var v27 uint16 = 32767
        for fi28 := 0; fi28 < 16; fi28++ {
            if (127 > (^(v4))) {
                var v29 uint32 = 2
                for fi30 := 0; fi30 < 2; fi30++ {
                    switch 9223372036854775807 {
                    case 0:
                        v3 += (float64(((-2.25 + 0.0) / (0.1 / 1e10))) + v3)
                    case 2147483647:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        switch 0 {
                        case -553:
                            if false {
                                v5 = append(v5, (fi28 | fi28))
                                if ((1.5 * (0.0 + 0.1)) < (0.1 - -2.25)) {
                                    v26 = v27
                                    var v31 bool = (9223372036854775807 >= -484)
                                    _ = v31
                                } else {
                                    var v32 float32 = (((0.0 - 0.1) + 0.0) + (-2.25 * -2.25))
                                    if true {
                                        if ((!true) && true) {
                                            if (false && true) {
                                                switch -418 {
                                                case 4294967295:
                                                    v33 := v2[(fi30 ^ (fi28 + 127))]
                                                    _ = v33
                                                case 517:
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                                switch (-(fi30)) {
                                                case 0:
                                                    var v34 uint32 = 2147483648
                                                    _ = v34
                                                case -2147483648:
                                                    var v35 int64 = -1
                                                    fmt.Println(fi30)
                                                    _ = v35
                                                case 128:
                                                    v29 += (v29 << ((65535) & 63))
                                                default:
                                                    switch fi30 {
                                                    case 1:
                                                        var v36 bool = false
                                                        _ = v36
                                                    case 9223372036854775807:
                                                        v1 = v1
                                                    }
                                                }
                                                v3 = (v3 / v3)
                                                v27 = v27
                                                v29 = v29
                                                if false {
                                                    v37 := GenS1{F0: 0}
                                                    if (106 == 236) {
                                                        v38 := GenS1{F0: 1}
                                                        v5 = append(v5, (fi28 * fi28))
                                                        v4 += (v4 - 42)
                                                        _ = v38
                                                    }
                                                    fmt.Println(v3, v26)
                                                    _ = v37
                                                } else {
                                                    var v39 uint16 = (v26 % ((v27) | 1))
                                                    fmt.Println(v29, v26, v27)
                                                    _ = v39
                                                }
                                            }
                                            v4 += v4
                                            if (4294967295 > v29) {
                                                var v40 int32 = 2147483647
                                                fmt.Println(v27)
                                                v29 = v29
                                                _ = v40
                                            } else {
                                                switch v1 {
                                                case "\n\t":
                                                    v41 := GenS1{F0: (v29 + (v29 * (v29 % ((v29) | 1))))}
                                                    _ = v41
                                                case "foo bar":
                                                    if len(v5) > 0 {
                                                        v5[0] = fi30
                                                    }
                                                }
                                                switch fi30 {
                                                case 149:
                                                    v3 = v3
                                                case 128:
                                                    v4 = v4
                                                case 32768:
                                                    v1 = v1
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = (fi28 & fi30)
                                                }
                                                var v42 uint16 = v27
                                                _ = v42
                                            }
                                        }
                                    }
                                    _ = v32
                                }
                            } else {
                                switch "" {
                                case "a":
                                    switch fi28 {
                                    case 32767:
                                        switch 128 {
                                        case 128:
                                            if false {
                                                v43 := GenS1{F0: v29}
                                                switch int(v26) {
                                                case 127:
                                                    if len(v5) > 0 {
                                                        v5[0] = fi28
                                                    }
                                                case 128:
                                                    if (-286 >= -1) {
                                                        if true {
                                                            var v44 uint32 = (v29 | 42)
                                                            var v45 uint16 = 255
                                                            if true {
                                                                fmt.Println(fi28)
                                                                switch -2147483648 {
                                                                case -432:
                                                                    v5 = append(v5, fi30)
                                                                case -179:
                                                                    var v46 uint8 = 42
                                                                    _ = v46
                                                                default:
                                                                    v47 := v2[0]
                                                                    _ = v47
                                                                }
                                                                var v48 uint16 = v26
                                                                _ = v48
                                                            }
                                                            _ = v44
                                                            _ = v45
                                                        } else {
                                                            v49 := GenS1{F0: (v29 | v29)}
                                                            var v50 string = ("" + "hello")
                                                            if false {
                                                                var v51 float32 = 1e10
                                                                v5 = append(v5, fi30)
                                                                _ = v51
                                                            }
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        var v52 uint16 = v26
                                                        v27 = (v26 % ((2) | 1))
                                                        _ = v52
                                                    } else {
                                                        v26 = v26
                                                        v29 = 887
                                                        if false {
                                                            _ = len(v5)
                                                            v3 = v3
                                                            if false {
                                                                v3 -= ((v3 / -2.25) * 1.5)
                                                                switch v1 {
                                                                case "\n\t":
                                                                    var v53 int64 = -768
                                                                    _ = v53
                                                                case "hello":
                                                                    var v54 string = v1
                                                                    _ = v54
                                                                default:
                                                                    v1 = ""
                                                                }
                                                            }
                                                            v55 := GenS2{F0: v29, F1: -1, F2: v29}
                                                            _ = v55
                                                        } else {
                                                            v26 += v27
                                                            if true {
                                                                v26 -= (v26 / ((65535) | 1))
                                                            } else {
                                                                delete(v6, (v1 + "foo bar"))
                                                                delete(v2, 127)
                                                                fmt.Println(v1)
                                                            }
                                                            if false {
                                                                v26 = 2
                                                                var v56 float64 = v3
                                                                v57 := GenS2{F0: v29, F1: 32767, F2: (v29 % ((1) | 1))}
                                                                var v58 bool = (-2.25 != (1.5 - (0.0 - ((-2.25 * 0.1) * 3.14159))))
                                                                _ = v56
                                                                _ = v57
                                                                _ = v58
                                                            } else {
                                                                fmt.Println(v26, v29, v4)
                                                                if true {
                                                                    v59 := GenS3{F0: 18446744073709551615}
                                                                    _ = len(v2)
                                                                    fmt.Println(fi28, v1)
                                                                    _ = v59
                                                                }
                                                                var v60 int8 = v4
                                                                var v61 uint16 = v27
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                            v26 += v26
                                                            var v62 float64 = 1.5
                                                            var v63 uint8 = 97
                                                            _ = len(v5)
                                                            v64 := GenS3{F0: 4294967295}
                                                            _ = v62
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                    }
                                                case 4294967295:
                                                    fmt.Println(v27, v3, fi30)
                                                default:
                                                    v29 += 4294967295
                                                    fmt.Println(fi30)
                                                    if false {
                                                        v65 := GenS4{F0: v4}
                                                        _ = len(v5)
                                                        v2[fi30] = "foo bar"
                                                        _ = v65
                                                    }
                                                }
                                                _ = v43
                                            }
                                        case 2147483647:
                                            var v66 int8 = v4
                                            v1 = (v1 + v1)
                                            if len(v5) > 0 {
                                                v5[0] = fi28
                                            }
                                            _ = v66
                                        case 742:
                                            var v67 uint8 = 2
                                            _ = v67
                                        }
                                    }
                                case "\n\t":
                                    switch (fi30 & 9223372036854775807) {
                                    case 128:
                                        if false {
                                            var v68 int = (fi30 | fi30)
                                            v27 = v27
                                            v3 -= (0.1 + 1.5)
                                            _ = v68
                                        } else {
                                            v4 -= 0
                                            v69 := GenS3{F0: 65535}
                                            v29 = v29
                                            _ = v69
                                        }
                                    }
                                }
                                var v70 uint8 = 255
                                v3 += ((v3 - (-2.25 * (v3 - v3))) / (1.5 - ((0.0 - v3) / float64(0.1))))
                                _ = v70
                            }
                        case 1:
                            var v71 float32 = 1e10
                            _ = v71
                        case 127:
                            var v72 float64 = (3.14159 * 0.0)
                            _ = v72
                        }
                    }
                    v73 := GenS5{F0: (-(fi30))}
                    v3 += v3
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi30
                    _ = v73
                }
                var v74 int = 1
                _ = v29
                _ = v74
            } else {
                v75 := GenS1{F0: 65535}
                var v76 float32 = -2.25
                _ = v75
                _ = v76
            }
            var v77 float64 = v3
            var v78 uint32 = 821
            _ = fi28
            _ = v77
            _ = v78
        }
        _ = v27
    } else {
        v79 := GenS6{F0: 1, F1: 0, F2: v1}
        _ = v79
    }
    v5 = append(v5, int(v4))
    v4 = v4
    for fi80 := 0; fi80 < 3; fi80++ {
        if false {
            var v81 int8 = v4
            for fi82 := 0; fi82 < 15; fi82++ {
                v83 := GenS7{F0: fi80}
                v2[fi82] = "a"
                _ = fi82
                _ = v83
            }
            _ = v81
        } else {
            fmt.Println(v4, v1)
        }
        _ = fi80
    }
    for fi84 := 0; fi84 < 20; fi84++ {
        v3 -= v3
        v3 += ((-2.25 * ((((3.14159 - v3) / 1.0) - v3) + v3)) * float64((float32(0.0) - 1.5)))
        _ = fi84
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
