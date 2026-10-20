package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint16; F2 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint8 }
type GenS8 struct { F0 int32; F1 int32 }
type GenS9 struct { F0 uint }
func genFn0() uint8 {
    var v5 float32 = ((0.1 / 1.0) * float32(0.1))
    v5 = v5
    v6 := []int{-2147483648}
    if (0.1 == ((float64((v5 + (((v5 * (v5 - (v5 / v5))) - 3.14159) + 1e10))) + 0.1) + 0.1)) {
        if len(v6) > 0 {
            v6[0] = 9223372036854775807
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            if (!(240 >= 76)) {
                v8 := GenS1{F0: 127, F1: 1, F2: 42}
                v9 := GenS1{F0: 1, F1: 611, F2: 32767}
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v8
                _ = v9
            }
            _ = fi7
        }
    }
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 int16 = -1
    var v2 float32 = 1e10
    var v3 bool = (2 != 0)
    v4 := []uint32{126, 4294967295, 0}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v10 := GenS2{F0: 1, F1: 0, F2: 1}
    switch "" {
    case "a":
        var v11 uint = 2
        _ = v11
    case "foo bar":
        v1 = (v1 * 1000)
    default:
        _ = len(v4)
    }
    v2 += float32(3.14159)
    if (v3 || (127 < 75)) {
        for fi12 := 0; fi12 < 8; fi12++ {
            fmt.Println(v1, fi12, v3)
            if (1.5 >= ((-2.25 * ((0.1 + ((3.14159 * ((float64(((v2 - (v2 - v2)) / -2.25)) + 1e10) / float64(v2))) * -2.25)) + 0.0)) / ((3.14159 - 1.5) + 3.14159))) {
                v2 = v2
                v13 := GenS2{F0: 1, F1: 555, F2: 106}
                _ = v13
            }
            v3 = v3
            _ = fi12
        }
        v14 := GenS1{F0: 1, F1: 2, F2: 65535}
        _ = v14
    }
    var v15 float64 = (((1.5 * 1e10) + 1.5) * 0.0)
    var v16 float32 = 1e10
    var v17 int32 = -2147483648
    switch 2147483647 {
    case 42:
        for fi18 := 0; fi18 < 10; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 int64 = 42
                _ = fi19
                _ = v20
            }
            v3 = v3
            _ = len(v4)
            _ = fi18
        }
    case -193:
        v4 = append(v4, 2)
    case -985:
        if (9223372036854775807 == 42) {
            for fi21 := 0; fi21 < 19; fi21++ {
                fmt.Println(v16, fi21, v15)
                v22 := GenS3{F0: 1, F1: 99}
                _ = fi21
                _ = v22
            }
            v15 += 0.0
            if (234 == 127) {
                var v23 float64 = 0.1
                var v24 int = 2147483647
                v25 := GenS4{F0: 4294967295, F1: 540, F2: ((v1 * v1) != 720)}
                _ = v23
                _ = v24
                _ = v25
            }
        } else {
            if (42 != 681) {
                var v26 uint64 = 42
                for fi27 := 0; fi27 < 8; fi27++ {
                    if len(v4) > 0 {
                        v4[0] = 0
                    }
                    switch fi27 {
                    case -2147483648:
                        var v28 int8 = 18
                        _ = v28
                    }
                    switch ((((("foo bar" + "\n\t") + (("" + "\n\t") + "foo bar")) + "a") + "foo bar") + "a") {
                    case "a":
                        v29 := GenS1{F0: -97, F1: 65535, F2: 764}
                        _ = v29
                    case "foo bar":
                        _ = len(v4)
                    case "":
                        var v30 int8 = 127
                        _ = v30
                    default:
                        if ((!v3) || (v3 && (2147483647 == 0))) {
                            switch fi27 {
                            case 32767:
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    fmt.Println(v17, v26, v1)
                                    _ = fi31
                                }
                            case -1:
                                v15 -= (1.5 + 3.14159)
                            default:
                                fmt.Println(fi27)
                            }
                            v16 += (((-2.25 * 3.14159) * -2.25) / float32(float64((v2 / v16))))
                        }
                    }
                    _ = fi27
                }
                var v32 uint8 = 207
                _ = v26
                _ = v32
            } else {
                if v3 {
                    var v33 int32 = (v17 >> ((0) & 63))
                    if (!v3) {
                        if false {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v34 uint = 1
                            for fi35 := 0; fi35 < 11; fi35++ {
                                fmt.Println(v1, v17, v2)
                                _ = fi35
                            }
                            _ = v34
                        }
                        fmt.Println(v16)
                        fmt.Println(v17, v1, v3)
                    }
                    switch -9223372036854775808 {
                    case 32767:
                        var v36 bool = (v1 > (v1 << ((1) & 63)))
                        _ = v36
                    default:
                        var v37 int8 = -1
                        _ = v37
                    }
                    _ = v33
                } else {
                    var v38 int8 = -1
                    v39 := GenS2{F0: v17, F1: 521, F2: 42}
                    if (false || v3) {
                        v38 += (v38 / ((v38) | 1))
                    }
                    _ = v38
                    _ = v39
                }
                fmt.Println(v16, v2)
                for fi40 := 0; fi40 < 19; fi40++ {
                    for fi41 := 0; fi41 < 9; fi41++ {
                        if ((50 <= v17) && ("a" <= "")) {
                            v1 -= v1
                            _ = len(v4)
                            if ((false || ((-9223372036854775808 >= 1) || v3)) || v3) {
                                switch ("hello" + "a") {
                                case "hello":
                                    var v42 uint = 237
                                    _ = v42
                                case "foo bar":
                                    if true {
                                        v43 := GenS5{F0: 2147483648}
                                        v44 := GenS5{F0: 4294967295}
                                        _ = v43
                                        _ = v44
                                    } else {
                                        var v45 bool = v3
                                        if len(v4) > 0 {
                                            v4[0] = 593
                                        }
                                        v4 = append(v4, 902)
                                        switch "a" {
                                        case "\n\t":
                                            switch fi41 {
                                            case -1:
                                                v46 := GenS6{F0: 127}
                                                var v47 int64 = 9223372036854775807
                                                fmt.Println(v45, v16)
                                                _ = v46
                                                _ = v47
                                            }
                                        case "foo bar":
                                            v45 = v45
                                        default:
                                            v48 := GenS7{F0: int64(fi41), F1: 14}
                                            _ = v48
                                        }
                                        _ = v45
                                    }
                                default:
                                    _ = len(v4)
                                    v49 := GenS8{F0: v17, F1: (v17 * (^(v17)))}
                                    v17 += 1
                                    var v50 uint8 = 115
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v49
                                    _ = v50
                                }
                                _ = len(v4)
                            } else {
                                v16 -= float32(v15)
                                switch (fi40 | -2147483648) {
                                case 1:
                                    v3 = (1 >= (v17 * 42))
                                case 32767:
                                    switch "" {
                                    case "hello":
                                        var v51 float32 = v2
                                        _ = v51
                                    case "":
                                        v52 := GenS2{F0: (-(v17)), F1: 0, F2: 4294967295}
                                        _ = v52
                                    case "a":
                                        if v3 {
                                            v17 = (v17 ^ (v17 | v17))
                                            fmt.Println(fi40)
                                            v53 := GenS1{F0: -26, F1: 783, F2: 42}
                                            var v54 uint16 = 904
                                            v4 = append(v4, 916)
                                            _ = v53
                                            _ = v54
                                        }
                                    default:
                                        fmt.Println(v16, v2)
                                    }
                                }
                            }
                            switch (fi41 - (fi41 - (fi41 ^ fi40))) {
                            case 309:
                                if v3 {
                                    if len(v4) > 0 {
                                        v4[0] = 300
                                    }
                                }
                            case 9223372036854775807:
                                var v55 int = fi41
                                _ = v55
                            case 684:
                                if v3 {
                                    v16 += -2.25
                                    fmt.Println(v2, v3, v17)
                                } else {
                                    if (!v3) {
                                        switch (fi41 & (^(fi40))) {
                                        case -9223372036854775808:
                                            v1 = (v1 | (v1 | v1))
                                        }
                                        var v56 uint64 = 18446744073709551615
                                        v3 = v3
                                        if len(v4) > 0 {
                                            v4[0] = 2
                                        }
                                        v3 = ((fi40 > 32767) && v3)
                                        var v57 int = (fi40 ^ fi40)
                                        switch "\n\t" {
                                        case "a":
                                            _ = len(v4)
                                        case "hello":
                                            var v58 uint32 = 166
                                            _ = v58
                                        case "foo bar":
                                            if v3 {
                                                v15 = v15
                                                var v59 int = v57
                                                _ = len(v4)
                                                _ = v59
                                            } else {
                                                v15 = ((v15 / 3.14159) / -2.25)
                                                var v60 int16 = (v1 | (-(v1)))
                                                _ = v60
                                            }
                                        }
                                        var v61 float64 = 0.1
                                        _ = v56
                                        _ = v57
                                        _ = v61
                                    } else {
                                        v62 := GenS9{F0: 1}
                                        var v63 float32 = float32(1.5)
                                        _ = v62
                                        _ = v63
                                    }
                                    var v64 string = "hello"
                                    _ = v64
                                }
                                var v65 int64 = -9223372036854775808
                                _ = v65
                            default:
                                var v66 string = ("hello" + "hello")
                                var v67 int32 = (v17 * (v17 << ((110) & 63)))
                                _ = len(v4)
                                _ = v66
                                _ = v67
                            }
                        }
                        v16 += ((float32(1e10) * ((((v16 * (1.5 + v16)) - v16) + (float32(v15) + v2)) / (((((v2 - 1e10) - v2) - (v2 + -2.25)) / v2) / 0.1))) + v2)
                        _ = len(v4)
                        _ = fi41
                    }
                    _ = fi40
                }
            }
            var v68 int16 = 32767
            _ = v68
        }
    default:
        var v69 bool = (42 == 50)
        _ = v69
    }
    for fi70 := 0; fi70 < 10; fi70++ {
        switch (^(fi70)) {
        case -240:
            v71 := GenS7{F0: 232, F1: 1}
            _ = v71
        }
        _ = fi70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
