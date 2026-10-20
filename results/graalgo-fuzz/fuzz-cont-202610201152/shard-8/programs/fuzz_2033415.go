package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 int; F1 uint8; F2 int8 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 int }
func genFn0() int16 {
    v7 := []int{0, 9223372036854775807, 1}
    if (!(false && true)) {
        for fi8 := 0; fi8 < 6; fi8++ {
            _ = len(v7)
            _ = fi8
        }
        switch -1 {
        case 4294967295:
            v9 := GenS1{F0: -30}
            _ = v9
        case 9223372036854775807:
            var v10 uint8 = 167
            _ = v10
        case 2147483647:
            var v11 int32 = 0
            _ = v11
        default:
            if (0 < 432) {
                var v12 uint32 = 153
                if (89 == 2147483647) {
                    var v13 uint64 = 42
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        if (0 >= v12) {
                            if true {
                                var v15 float32 = 1.5
                                _ = v15
                            } else {
                                var v16 uint64 = 4294967295
                                v16 += (v16 - (v16 & 4294967295))
                                _ = v16
                            }
                            for fi17 := 0; fi17 < 18; fi17++ {
                                var v18 string = ""
                                v18 = ("foo bar" + v18)
                                if (1 == -691) {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    switch "\n\t" {
                                    case "\n\t":
                                        if true {
                                            v19 := GenS1{F0: 0}
                                            _ = v19
                                        } else {
                                            v18 = (("\n\t" + "") + v18)
                                        }
                                    default:
                                        switch v18 {
                                        case "\n\t":
                                            var v20 int8 = -110
                                            _ = v20
                                        }
                                    }
                                } else {
                                    v21 := GenS1{F0: 1000}
                                    if false {
                                        if false {
                                            var v22 float32 = 1.5
                                            _ = v22
                                        } else {
                                            v23 := GenS2{F0: 209, F1: 42, F2: 647}
                                            var v24 float64 = float64((-2.25 - 3.14159))
                                            v25 := GenS3{F0: v18, F1: 42}
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        if (!true) {
                                            switch v18 {
                                            case "hello":
                                                if len(v7) > 0 {
                                                    v7[0] = fi17
                                                }
                                            default:
                                                var v26 float32 = (3.14159 * ((float32(-2.25) + 1.5) * ((0.0 + float32(3.14159)) - (1.5 - (1.5 / 3.14159)))))
                                                _ = v26
                                            }
                                            fmt.Println(v18)
                                        } else {
                                            fmt.Println(fi14)
                                            v12 = v12
                                            v12 -= v12
                                        }
                                    }
                                    _ = v21
                                }
                                _ = fi17
                                _ = v18
                            }
                        } else {
                            v27 := GenS4{F0: 1, F1: v12}
                            if false {
                                v28 := GenS2{F0: 0, F1: 42, F2: 4294967295}
                                _ = v28
                            }
                            switch fi14 {
                            case 32767:
                                var v29 int16 = 32767
                                _ = v29
                            default:
                                for fi30 := 0; fi30 < 9; fi30++ {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    _ = fi30
                                }
                            }
                            _ = v27
                        }
                        if false {
                            var v31 float32 = float32((((float64(3.14159) - float64((3.14159 - (0.0 / 1e10)))) / 0.1) + 0.0))
                            v12 = v12
                            _ = v31
                        } else {
                            v7 = append(v7, (fi14 - fi14))
                            v12 -= v12
                        }
                        v12 -= (^(v12))
                        _ = fi14
                    }
                }
                var v32 float32 = float32(0.1)
                _ = v12
                _ = v32
            } else {
                var v33 uint32 = 2147483648
                _ = v33
            }
        }
        v34 := GenS5{F0: 127, F1: 128, F2: 127}
        _ = v34
    }
    if len(v7) > 0 {
        v7[0] = -26
    }
    _ = v7
    return 32767
}
func genFn1(p0 bool, p1 uint, p2 int64) bool {
    v35 := []int{128, 9223372036854775807, 127}
    if true {
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 uint8 = 255
            _ = fi36
            _ = v37
        }
    }
    if (679 != 1) {
        for fi38 := 0; fi38 < 5; fi38++ {
            p2 = (p2 | 0)
            p2 += -1
            fmt.Println(fi38, p0, fi38)
            _ = fi38
        }
        v39 := GenS6{F0: (p2 * p2), F1: 0}
        p0 = (4294967295 == 0)
        _ = v39
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    return p0
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (0.1 - (0.0 + 0.1))
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v5 string = "\n\t"
    v6 := make(map[int]uint32)
    v40 := []int{105, 1, 9223372036854775807, 32767}
    v40 = append(v40, 479)
    var v41 uint64 = 744
    var v42 float32 = (1e10 - 1e10)
    var v43 string = v5
    if true {
        switch -902 {
        case 42:
            var v44 uint = 463
            _ = v44
        case 9223372036854775807:
            v45 := GenS7{F0: ("foo bar" + (((((v5 + ((v5 + "hello") + "a")) + v43) + v5) + ("hello" + v5)) + v43))}
            _ = v45
        case -681:
            var v46 float64 = float64(v42)
            _ = v46
        }
    } else {
        switch "a" {
        case "":
            var v47 string = (("\n\t" + (((v43 + v43) + v43) + "")) + "foo bar")
            _ = v47
        case "foo bar":
            switch "\n\t" {
            case "foo bar":
                v5 = v5
            }
        default:
            if (!false) {
                var v48 uint64 = v41
                switch (v43 + (("\n\t" + (((v5 + "") + v43) + ((v5 + "hello") + ((v43 + v5) + v43)))) + "")) {
                case "a":
                    var v49 int32 = 1
                    _ = v49
                case "":
                    if false {
                        v48 += (^(v41))
                        var v50 int = 659
                        v5 = (v5 + "")
                        _ = v50
                    } else {
                        for fi51 := 0; fi51 < 2; fi51++ {
                            for fi52 := 0; fi52 < 9; fi52++ {
                                var v53 int = (fi51 >> ((260) & 63))
                                v54 := GenS8{F0: 1, F1: 55}
                                v43 = v5
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                            var v55 float64 = float64(0.0)
                            if false {
                                var v56 bool = (!false)
                                fmt.Println(v43, v2, v41)
                                _ = v56
                            } else {
                                fmt.Println(v43, v48)
                                var v57 int32 = 0
                                var v58 uint8 = 42
                                _ = v57
                                _ = v58
                            }
                            _ = fi51
                            _ = v55
                        }
                        if false {
                            if true {
                                var v59 string = (v5 + "foo bar")
                                _ = v59
                            } else {
                                var v60 uint8 = 180
                                if ((v60 / ((v60) | 1)) == 128) {
                                    v43 = v43
                                } else {
                                    v40 = append(v40, -889)
                                }
                                _ = v60
                            }
                            var v61 uint = 309
                            v2 = -2.25
                            _ = v61
                        } else {
                            v62 := GenS9{F0: -2147483648}
                            var v63 uint16 = 2
                            _ = v62
                            _ = v63
                        }
                    }
                }
                if (2 >= 42) {
                    var v64 string = v43
                    _ = v64
                }
                _ = v48
            }
        }
    }
    var v65 string = v5
    v2 -= 3.14159
    for fi66 := 0; fi66 < 15; fi66++ {
        _ = len(v40)
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
        var v67 int16 = 774
        _ = fi66
        _ = v67
    }
    v42 -= v42
    var v68 uint = 1
    v2 = (0.0 + (float64(v42) - v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v65
    _ = v68
}
