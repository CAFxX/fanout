package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS6 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    v5 := v2[""]
    var v6 string = v5
    var v7 float64 = (3.14159 * 1.5)
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := []int{(fi8 - -9223372036854775808)}
            v11 := GenS1{F0: 127, F1: 849}
            v12 := GenS2{F0: 42, F1: 1}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        switch v6 {
        case "hello":
            for fi13 := 0; fi13 < 18; fi13++ {
                v7 -= (float64(-2.25) - (v7 / v7))
                fmt.Println(v5, v7)
                switch "" {
                case "hello":
                    v6 = ((((v5 + "") + (v6 + (((v6 + (v6 + "\n\t")) + v5) + ""))) + ("" + "foo bar")) + "foo bar")
                default:
                    _ = len(v3)
                }
                _ = fi13
            }
        case "foo bar":
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := []int{0, fi14, (fi14 | fi14)}
                fmt.Println(fi8)
                _ = fi14
                _ = v15
            }
        case "":
            var v16 uint16 = 173
            _ = v16
        }
        _ = fi8
    }
    var v17 uint8 = 42
    var v18 int8 = 100
    for fi19 := 0; fi19 < 13; fi19++ {
        if (!(true && (true || (v18 >= (v18 & (v18 & (v18 ^ (v18 ^ 0)))))))) {
            var v20 int16 = -32768
            v21 := GenS3{F0: (v18 % ((v18) | 1))}
            _ = v20
            _ = v21
        }
        _ = fi19
    }
    switch (v5 + v6) {
    case "hello":
        var v22 uint16 = 563
        _ = v22
    case "a":
        if true {
            var v23 int16 = 42
            v7 += v7
            for fi24 := 0; fi24 < 18; fi24++ {
                if (v6 <= "foo bar") {
                    var v25 int32 = 1
                    _ = v25
                } else {
                    fmt.Println(v5, v17)
                    if true {
                        v26 := GenS3{F0: (v18 ^ 0)}
                        switch (^(fi24)) {
                        case 9223372036854775807:
                            fmt.Println(fi24)
                        case -9223372036854775808:
                            for fi27 := 0; fi27 < 9; fi27++ {
                                var v28 int32 = 0
                                _ = fi27
                                _ = v28
                            }
                        case 0:
                            switch v5 {
                            case "\n\t":
                                v5 = v6
                            default:
                                var v29 uint64 = 42
                                _ = v29
                            }
                        default:
                            fmt.Println(fi24, v17, v17)
                        }
                        _ = v26
                    } else {
                        var v30 uint16 = 1
                        _ = v30
                    }
                }
                var v31 string = "foo bar"
                if (207 != 0) {
                    fmt.Println(v5, fi24, v5)
                    if false {
                        v4[(fi24 >> ((136) & 63))] = (v31 + v6)
                    } else {
                        if ((v31 + (v6 + v5)) == v6) {
                            if ((fi24 * fi24) < fi24) {
                                if false {
                                    switch fi24 {
                                    case 750:
                                        var v32 uint32 = uint32(v17)
                                        _ = v32
                                    case 42:
                                        switch v31 {
                                        case "":
                                            if true {
                                                v6 = "hello"
                                                delete(v3, v6)
                                            }
                                        case "foo bar":
                                            var v33 int8 = (v18 << ((65535) & 63))
                                            _ = v33
                                        case "a":
                                            fmt.Println(v5, v31)
                                        default:
                                            v34 := v1[(fi24 - 2147483647)]
                                            _ = v34
                                        }
                                    case -691:
                                        switch (v6 + v5) {
                                        case "\n\t":
                                            var v35 uint32 = 65535
                                            _ = v35
                                        case "foo bar":
                                            var v36 int16 = (v23 << ((2) & 63))
                                            _ = v36
                                        case "":
                                            var v37 int64 = -467
                                            _ = v37
                                        default:
                                            v7 -= v7
                                        }
                                    default:
                                        var v38 float32 = ((0.1 / 1.0) * ((0.0 - 3.14159) - 3.14159))
                                        _ = v38
                                    }
                                    switch v5 {
                                    case "a":
                                        v7 += float64(((3.14159 - 1.5) * (3.14159 + 3.14159)))
                                    case "":
                                        if false {
                                            switch v6 {
                                            case "\n\t":
                                                if (true && ((float32(3.14159) / ((0.1 - (1e10 - 0.0)) * 1.5)) > (-2.25 + (float32(((((v7 * (v7 - 3.14159)) - float64(float32(v7))) * ((float64(0.1) + -2.25) * (((0.1 - v7) - (v7 / (float64((-2.25 / float32(v7))) - v7))) + float64(float32(0.1))))) + v7)) / ((1e10 - 1e10) + 1.5))))) {
                                                    var v39 bool = true
                                                    _ = v39
                                                } else {
                                                    if (!(!((!false) && (1 > (-(fi24)))))) {
                                                        v40 := []int{fi24}
                                                        v41 := GenS2{F0: 4294967295, F1: 484}
                                                        _ = len(v1)
                                                        _ = v40
                                                        _ = v41
                                                    } else {
                                                        for fi42 := 0; fi42 < 6; fi42++ {
                                                            var v43 int16 = v23
                                                            v6 = v31
                                                            _ = fi42
                                                            _ = v43
                                                        }
                                                        if ((v18 ^ 127) >= (v18 - v18)) {
                                                            var v44 int = (fi24 ^ (fi24 / ((fi24) | 1)))
                                                            v7 = v7
                                                            v17 = (v17 % (((^(v17))) | 1))
                                                            _ = v44
                                                        }
                                                        v45 := GenS3{F0: v18}
                                                        _ = v45
                                                    }
                                                    var v46 uint16 = 1
                                                    _ = len(v4)
                                                    _ = v46
                                                }
                                            }
                                            v47 := GenS4{F0: 2147483648}
                                            _ = v47
                                        }
                                    case "hello":
                                        v23 = v23
                                    }
                                    v3[(("" + (("foo bar" + ((v5 + ("\n\t" + v6)) + v6)) + v31)) + v5)] = (v6 + "")
                                } else {
                                    v48 := GenS3{F0: int8(v23)}
                                    _ = v48
                                }
                                fmt.Println(v31, v17, fi24)
                                var v49 int32 = -1
                                _ = v49
                            } else {
                                if false {
                                    v50 := []int{-9223372036854775808, (^(fi24)), 128}
                                    _ = v50
                                }
                            }
                            v7 -= ((0.0 * v7) * 1.5)
                        } else {
                            switch (fi24 + fi24) {
                            case 1:
                                var v51 uint16 = 2
                                _ = v51
                            case -16:
                                v52 := []int{fi24, fi24}
                                _ = v52
                            case -9223372036854775808:
                                v2[""] = "hello"
                            }
                            fmt.Println(v23, v6, v7)
                        }
                    }
                    var v53 uint32 = 2147483648
                    _ = v53
                }
                _ = fi24
                _ = v31
            }
            _ = v23
        } else {
            v17 += v17
            for fi54 := 0; fi54 < 9; fi54++ {
                for fi55 := 0; fi55 < 8; fi55++ {
                    if (!(!((false && false) && (true || true)))) {
                        var v56 int8 = v18
                        switch v5 {
                        case "hello":
                            var v57 uint64 = 374
                            _ = v57
                        case "":
                            if false {
                                var v58 uint16 = 0
                                var v59 float32 = -2.25
                                if false {
                                    v18 -= v18
                                    var v60 uint64 = 18446744073709551615
                                    _ = v60
                                }
                                _ = v58
                                _ = v59
                            } else {
                                if false {
                                    _ = len(v2)
                                    v61 := GenS3{F0: (v56 | v56)}
                                    v62 := GenS3{F0: v56}
                                    _ = v61
                                    _ = v62
                                }
                                v63 := []int{}
                                _ = v63
                            }
                        case "foo bar":
                            var v64 float64 = v7
                            _ = v64
                        default:
                            v65 := []int{1, fi54, (fi55 & fi55), -484}
                            _ = v65
                        }
                        _ = v56
                    } else {
                        v18 -= (v18 + (-(v18)))
                        var v66 bool = (true || true)
                        var v67 bool = (v66 && false)
                        _ = v66
                        _ = v67
                    }
                    _ = fi55
                }
                for fi68 := 0; fi68 < 11; fi68++ {
                    if (false || (!(!true))) {
                        if true {
                            if false {
                                var v69 int16 = 0
                                v17 = v17
                                switch (-(fi68)) {
                                case -190:
                                    if (42 > 910) {
                                        v70 := v3[v5]
                                        var v71 int64 = -947
                                        switch 42 {
                                        case 4294967295:
                                            v3[v6] = v70
                                        case 32767:
                                            switch (^(fi54)) {
                                            case 128:
                                                v69 = int16(fi68)
                                            case -1:
                                                v72 := GenS5{F0: 255, F1: v17, F2: 368}
                                                _ = v72
                                            }
                                        }
                                        _ = v70
                                        _ = v71
                                    }
                                default:
                                    v7 = (((((v7 - float64(1e10)) / v7) - v7) + v7) + v7)
                                }
                                var v73 bool = (728 != -632)
                                fmt.Println(v7)
                                _ = v69
                                _ = v73
                            } else {
                                v74 := []int{(fi54 % ((int(v17)) | 1)), fi68, (fi54 - (fi54 << ((1) & 63))), fi68}
                                _ = v74
                            }
                            var v75 uint64 = 4294967295
                            switch (-(fi68)) {
                            case -2147483648:
                                v18 = (v18 >> ((42) & 63))
                            case -993:
                                if (((1 == 1) && (!true)) || (false || (192 != 407))) {
                                    v76 := GenS5{F0: 255, F1: (v17 - 255), F2: v75}
                                    v77 := GenS6{F0: 964, F1: "a"}
                                    fmt.Println(v17, v17)
                                    v78 := GenS2{F0: 42, F1: 42}
                                    _ = v76
                                    _ = v77
                                    _ = v78
                                } else {
                                    v2[(v5 + ((v5 + ("a" + "hello")) + "\n\t"))] = "a"
                                }
                            }
                            _ = v75
                        }
                        v17 = 1
                        var v79 int16 = 1
                        _ = v79
                    } else {
                        fmt.Println(v17, fi54)
                    }
                    v5 = ("foo bar" + v6)
                    _ = fi68
                }
                _ = fi54
            }
        }
    case "foo bar":
        for fi80 := 0; fi80 < 2; fi80++ {
            var v81 int32 = 363
            v82 := []int{fi80, 816, fi80, (fi80 * fi80)}
            _ = fi80
            _ = v81
            _ = v82
        }
    }
    switch v6 {
    case "a":
        v5 = ((("" + (v6 + (v6 + "\n\t"))) + v6) + "foo bar")
    default:
        if false {
            var v83 int32 = 816
            _ = v83
        } else {
            fmt.Println(v5)
        }
    }
    v6 = "\n\t"
    var v84 uint64 = 42
    var v85 uint32 = 65535
    v5 = ("hello" + (((((v6 + v6) + v5) + v6) + (v5 + v6)) + v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v84
    _ = v85
}
