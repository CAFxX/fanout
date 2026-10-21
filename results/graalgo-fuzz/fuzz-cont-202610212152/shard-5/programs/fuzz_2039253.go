package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint16; F1 uint64; F2 int }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8; F1 int32 }
func genFn0() uint64 {
    var v5 int32 = 886
    v6 := []int{276, -1}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int16 = 16
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 139
    v3 := []uint32{1, 2147483648, 42, 65535}
    var v4 float32 = (-2.25 * ((0.0 * 1e10) - 0.0))
    if true {
        var v9 float32 = v4
        var v10 uint16 = 22
        _ = v10
        _ = v9
    } else {
        if true {
            var v11 float32 = 1e10
            for fi12 := 0; fi12 < 12; fi12++ {
                var v13 float64 = float64(v11)
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 10; fi14++ {
                if len(v3) > 0 {
                    v3[0] = 0
                }
                var v15 uint16 = 65535
                _ = fi14
                _ = v15
            }
            _ = v11
        } else {
            fmt.Println(v4, v2)
            fmt.Println(v4, v4)
        }
    }
    v2 = (v2 + v2)
    v2 = 0
    if (!true) {
        v4 -= v4
        var v16 uint = 1
        if (95 <= 2) {
            var v17 int16 = v2
            var v18 string = "foo bar"
            for fi19 := 0; fi19 < 5; fi19++ {
                v18 = v18
                _ = fi19
            }
            _ = v17
            _ = v18
        }
        _ = v16
    } else {
        var v20 bool = true
        fmt.Println(v20, v4)
        _ = v20
    }
    fmt.Println(v4, v4)
    v2 = (v2 % ((v2) | 1))
    fmt.Println(v2, v4)
    switch "foo bar" {
    case "hello":
        v21 := GenS1{F0: 1, F1: 2, F2: -1}
        _ = v21
    case "\n\t":
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS2{F0: 232, F1: -40, F2: (v2 - (-(v2)))}
            if true {
                v24 := GenS3{F0: 1, F1: ("a" + "a")}
                v25 := GenS3{F0: 1, F1: "\n\t"}
                for fi26 := 0; fi26 < 13; fi26++ {
                    if (((!(!false)) && (198 == 575)) && (false || (2147483648 <= 1))) {
                        var v27 uint8 = 42
                        delete(v1, 42)
                        v2 = (v2 % (((v2 + 42)) | 1))
                        var v28 int8 = -1
                        _ = v27
                        _ = v28
                    }
                    switch "hello" {
                    case "foo bar":
                        if true {
                            if ((-(v2)) > (v2 % ((v2) | 1))) {
                                var v29 string = "hello"
                                var v30 uint = 418
                                switch fi22 {
                                case -31:
                                    switch fi22 {
                                    case 2147483647:
                                        fmt.Println(v4, fi26, v4)
                                    }
                                case 32767:
                                    v3 = append(v3, 42)
                                case -668:
                                    var v31 string = v29
                                    _ = v31
                                }
                                _ = v29
                                _ = v30
                            }
                            if true {
                                var v32 uint = 598
                                v4 += v4
                                v2 = v2
                                _ = v32
                            } else {
                                v33 := GenS3{F0: 255, F1: (("\n\t" + "a") + "\n\t")}
                                if true {
                                    fmt.Println(fi22)
                                    v4 += v4
                                    v34 := GenS4{F0: 564, F1: 4294967295, F2: fi26}
                                    _ = v34
                                } else {
                                    fmt.Println(v4)
                                    if true {
                                        v4 = v4
                                        v4 += float32(1e10)
                                    }
                                    v35 := v1[2147483647]
                                    var v36 int8 = 1
                                    v4 -= 1.5
                                    _ = v35
                                    _ = v36
                                }
                                if (!(3.14159 <= float64(v4))) {
                                    switch "a" {
                                    case "":
                                        switch "hello" {
                                        case "a":
                                            v37 := GenS3{F0: 1, F1: "hello"}
                                            _ = v37
                                        case "foo bar":
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(fi22, fi26)
                                            }
                                            delete(v1, 83)
                                        case "\n\t":
                                            if true {
                                                if ((-(v2)) <= v2) {
                                                    var v38 float64 = 3.14159
                                                    v39 := GenS1{F0: -39, F1: 42, F2: fi22}
                                                    v40 := GenS2{F0: 0, F1: 1, F2: (v2 >> ((4294967295) & 63))}
                                                    _ = v38
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    fmt.Println(fi26)
                                                    var v41 float32 = (v4 * v4)
                                                    v2 = v2
                                                    v4 = ((1e10 / v41) * ((((v41 / v41) * 3.14159) / v41) + (v4 - v41)))
                                                    _ = v41
                                                }
                                            }
                                        }
                                    }
                                    if true {
                                        if true {
                                            if false {
                                                var v42 int8 = -34
                                                v2 = v2
                                                if (859 >= 993) {
                                                    v2 += (v2 << ((604) & 63))
                                                    v42 -= v42
                                                }
                                                _ = v42
                                            } else {
                                                fmt.Println(fi22, fi22)
                                            }
                                        }
                                        switch "hello" {
                                        case "hello":
                                            var v43 int = 128
                                            _ = v43
                                        }
                                    }
                                    if true {
                                        v2 -= (v2 + v2)
                                        delete(v1, 0)
                                        var v44 int64 = 1
                                        v45 := v1[2]
                                        fmt.Println(v45)
                                        _ = v44
                                        _ = v45
                                    }
                                } else {
                                    var v46 int64 = 661
                                    _ = v46
                                }
                                v47 := GenS1{F0: 100, F1: 1, F2: 4294967295}
                                _ = v33
                                _ = v47
                            }
                            var v48 bool = false
                            var v49 string = (("\n\t" + "a") + ((("a" + "foo bar") + "") + "hello"))
                            _ = v48
                            _ = v49
                        } else {
                            var v50 uint8 = 158
                            v2 = v2
                            var v51 int8 = -98
                            delete(v1, 65535)
                            fmt.Println(v50, v2, v50)
                            _ = v50
                            _ = v51
                        }
                    case "hello":
                        switch fi26 {
                        case 32767:
                            var v52 int = (fi22 % ((-1) | 1))
                            _ = v52
                        case -9223372036854775808:
                            v4 += v4
                            var v53 uint64 = 1
                            _ = v53
                        }
                    case "":
                        v2 = 126
                    }
                    v3 = append(v3, 42)
                    _ = fi26
                }
                _ = v24
                _ = v25
            } else {
                for fi54 := 0; fi54 < 16; fi54++ {
                    v2 -= v2
                    var v55 float64 = -2.25
                    _ = fi54
                    _ = v55
                }
                for fi56 := 0; fi56 < 5; fi56++ {
                    fmt.Println(v4)
                    v2 -= (v2 + v2)
                    var v57 uint16 = 1
                    var v58 uint16 = v57
                    _ = fi56
                    _ = v57
                    _ = v58
                }
            }
            _ = fi22
            _ = v23
        }
    case "foo bar":
        v59 := GenS5{F0: 252, F1: (!(true || true))}
        _ = v59
    }
    if ("hello" >= "hello") {
        var v60 int = 2147483647
        var v61 uint64 = 2
        v3 = append(v3, 2)
        _ = v60
        _ = v61
    } else {
        v4 = ((0.0 - ((v4 * -2.25) / v4)) / (v4 * v4))
    }
    if true {
        v4 -= 1e10
    } else {
        v62 := GenS6{F0: -2147483648, F1: int8(v2)}
        for fi63 := 0; fi63 < 12; fi63++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi64 := 0; fi64 < 16; fi64++ {
                if len(v3) > 0 {
                    v3[0] = 40
                }
                fmt.Println(v2)
                v65 := GenS1{F0: 1, F1: 32767, F2: fi63}
                _ = fi64
                _ = v65
            }
            for fi66 := 0; fi66 < 7; fi66++ {
                if true {
                    if (false || (false && (!false))) {
                        v4 -= (1e10 - v4)
                        if true {
                            switch "" {
                            case "\n\t":
                                if false {
                                    fmt.Println(v4, v4, fi66)
                                    switch 709 {
                                    case -367:
                                        v67 := GenS1{F0: 100, F1: 65535, F2: (fi63 & (-(fi63)))}
                                        _ = v67
                                    case -1:
                                        v2 += (v2 / (((v2 | v2)) | 1))
                                    case 532:
                                        var v68 uint = 326
                                        _ = v68
                                    default:
                                        switch (("\n\t" + ("a" + "")) + "") {
                                        case "a":
                                            var v69 int32 = -518
                                            _ = v69
                                        case "hello":
                                            var v70 int8 = 100
                                            _ = v70
                                        }
                                    }
                                    if false {
                                        v1[2147483647] = (fi66 % (((fi66 / ((fi63) | 1))) | 1))
                                        v4 = float32((((0.1 / (3.14159 / -2.25)) / -2.25) - 1e10))
                                    } else {
                                        fmt.Println(fi63)
                                        if true {
                                            var v71 int64 = 0
                                            var v72 float32 = v4
                                            _ = v71
                                            _ = v72
                                        }
                                    }
                                    v73 := GenS7{F0: 65535}
                                    v4 += ((-2.25 + float32(3.14159)) / v4)
                                    _ = len(v3)
                                    _ = v73
                                }
                            case "":
                                if true {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = len(v1)
                                    v2 += v2
                                }
                            case "hello":
                                if true {
                                    v4 -= v4
                                    v2 += (v2 ^ (-(v2)))
                                    _ = len(v1)
                                } else {
                                    v74 := GenS2{F0: 255, F1: -128, F2: (-(v2))}
                                    if (!true) {
                                        switch (fi63 & (fi63 / (((fi63 / (((fi66 << ((65535) & 63))) | 1))) | 1))) {
                                        case -368:
                                            var v75 uint64 = 1
                                            _ = v75
                                        case 32768:
                                            v76 := GenS4{F0: 42, F1: 0, F2: -2147483648}
                                            _ = v76
                                        }
                                        v1[2147483648] = (fi66 >> ((4294967295) & 63))
                                    }
                                    _ = v74
                                }
                                v4 = ((1e10 * v4) + (0.0 + v4))
                                v2 = (v2 - (v2 * v2))
                            default:
                                var v77 int8 = -49
                                _ = v77
                            }
                            switch "hello" {
                            case "hello":
                                var v78 int16 = v2
                                _ = v78
                            }
                            switch ((("hello" + "\n\t") + "foo bar") + ("a" + ("a" + ("foo bar" + "a")))) {
                            case "\n\t":
                                v4 = v4
                            case "a":
                                var v79 float32 = v4
                                _ = v79
                            case "":
                                v4 += 0.1
                            }
                        } else {
                            if false {
                                v80 := GenS4{F0: 42, F1: 333, F2: (fi66 ^ fi66)}
                                _ = v80
                            } else {
                                switch "" {
                                case "hello":
                                    var v81 float64 = 0.0
                                    _ = v81
                                case "a":
                                    v2 -= 32767
                                default:
                                    v2 = v2
                                }
                            }
                        }
                        v82 := GenS8{F0: 8, F1: 1}
                        _ = v82
                    } else {
                        var v83 string = "a"
                        v2 = (v2 | -1)
                        delete(v1, 2147483648)
                        _ = v83
                    }
                }
                var v84 int8 = 0
                var v85 float64 = 3.14159
                v4 = 1e10
                _ = fi66
                _ = v84
                _ = v85
            }
            _ = fi63
        }
        v2 -= 1
        _ = v62
    }
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
