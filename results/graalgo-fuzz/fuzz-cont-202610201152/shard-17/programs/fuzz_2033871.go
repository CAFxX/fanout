package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int32; F1 int; F2 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool }
func genFn0() uint16 {
    var v5 float32 = 0.0
    var v6 int = -855
    switch v6 {
    case 42:
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(fi7)
            v8 := GenS1{F0: 65535, F1: 121}
            v9 := []int{(fi7 - -527), (-(v6))}
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        var v10 int16 = -32768
        _ = v10
    }
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint = 892
    var v2 uint8 = 26
    var v3 uint = (^(v1))
    var v4 string = "a"
    var v11 float64 = -2.25
    for fi12 := 0; fi12 < 19; fi12++ {
        v11 = (v11 - (1e10 + float64(((3.14159 / (0.1 + ((3.14159 - (3.14159 * 3.14159)) / float32(((v11 - float64(1e10)) - v11))))) * 1.5))))
        _ = fi12
    }
    var v13 uint16 = genFn0()
    var v14 int = 127
    if false {
        var v15 uint32 = 42
        if (v11 >= (float64(1e10) + -2.25)) {
            v4 = (v4 + "foo bar")
            v16 := []int{}
            _ = v16
        }
        _ = v15
    }
    v17 := []int{v14, 406, v14}
    v2 -= 22
    for fi18 := 0; fi18 < 8; fi18++ {
        v4 = v4
        _ = fi18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v14 -= v14
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 bool = (((v2 | v2) < (v2 * 203)) && (!(43 == 796)))
            switch "foo bar" {
            case "hello":
                if len(v17) > 0 {
                    v17[0] = -2147483648
                }
                fmt.Println(v2)
                fmt.Println(v13, v14, v21)
            case "a":
                var v22 string = "foo bar"
                _ = v22
            }
            _ = fi20
            _ = v21
        }
        switch (("a" + v4) + "a") {
        case "":
            v14 = 9223372036854775807
        case "foo bar":
            var v23 float64 = 0.1
            _ = v23
        default:
            for fi24 := 0; fi24 < 7; fi24++ {
                if (2147483648 == 2147483648) {
                    if len(v17) > 0 {
                        v17[0] = 609
                    }
                } else {
                    v13 = (^(v13))
                }
                if len(v17) > 0 {
                    v17[0] = fi24
                }
                _ = fi24
            }
        }
        _ = fi19
    }
    fmt.Println(v1, v3, v1)
    var v25 uint32 = 388
    switch v4 {
    case "":
        if (v4 > ((v4 + v4) + ("a" + v4))) {
            if false {
                v4 = ("" + ("a" + ""))
            }
            switch 9223372036854775807 {
            case 1:
                for fi26 := 0; fi26 < 17; fi26++ {
                    if (uint64(v1) != 2) {
                        fmt.Println(v3)
                        v27 := GenS2{F0: (v2 - (v2 + (v2 - (^(v2))))), F1: 42}
                        fmt.Println(v11, v13)
                        _ = v27
                    } else {
                        switch v4 {
                        case "":
                            v4 = ("foo bar" + ("" + (v4 + v4)))
                        case "\n\t":
                            v25 = (^(v25))
                        default:
                            for fi28 := 0; fi28 < 1; fi28++ {
                                if true {
                                    fmt.Println(v1)
                                }
                                v1 += 42
                                v1 -= v3
                                v29 := GenS3{F0: false}
                                _ = fi28
                                _ = v29
                            }
                        }
                        v17 = append(v17, v14)
                        fmt.Println(v11)
                    }
                    switch (fi26 >> (((v3 * (^(v1)))) & 63)) {
                    case 841:
                        for fi30 := 0; fi30 < 12; fi30++ {
                            switch (fi30 << ((uint(v25)) & 63)) {
                            case 544:
                                fmt.Println(v2, v3)
                            case 32768:
                                var v31 int8 = 127
                                _ = v31
                            case 2147483647:
                                v32 := GenS2{F0: (v2 - (v2 * (v2 * uint8(v25)))), F1: 42}
                                _ = v32
                            default:
                                v17 = append(v17, (v14 - v14))
                            }
                            _ = fi30
                        }
                    case 1:
                        v33 := GenS1{F0: (v13 % (((v13 << ((1) & 63))) | 1)), F1: (^(v2))}
                        _ = v33
                    }
                    _ = fi26
                }
            case -837:
                v4 = ""
            case 32767:
                if ((-2147483648 == -494) || (false && ((!false) || (0.0 < -2.25)))) {
                    v17 = append(v17, (-(v14)))
                    v34 := GenS4{F0: 159, F1: (v13 * 0)}
                    _ = v34
                } else {
                    if len(v17) > 0 {
                        v17[0] = -2147483648
                    }
                }
            }
        }
    case "a":
        for fi35 := 0; fi35 < 15; fi35++ {
            var v36 float32 = 1.5
            var v37 uint64 = 2
            if (true || true) {
                v38 := GenS5{F0: 498, F1: -145, F2: v1}
                for fi39 := 0; fi39 < 17; fi39++ {
                    v3 = 255
                    if ((!true) && (false || true)) {
                        v11 = (0.1 - 1.5)
                        if ((v36 * float32(((1e10 * 0.1) * 1.5))) <= float32(-2.25)) {
                            if (fi39 != (v14 * (v14 << ((uint(v2)) & 63)))) {
                                if ((v13 * v13) <= 1) {
                                    var v40 uint8 = v2
                                    _ = v40
                                } else {
                                    if len(v17) > 0 {
                                        v17[0] = -961
                                    }
                                    fmt.Println(v2, v4)
                                }
                                if len(v17) > 0 {
                                    v17[0] = fi39
                                }
                                var v41 int32 = -1
                                _ = v41
                            } else {
                                var v42 uint = (v1 % ((v3) | 1))
                                _ = v42
                            }
                            _ = len(v17)
                        } else {
                            var v43 uint64 = 2
                            fmt.Println(fi35, v25)
                            _ = v43
                        }
                        v3 = (^(v1))
                    } else {
                        fmt.Println(v11)
                        var v44 uint = (v1 + 255)
                        if true {
                            var v45 bool = (-1 == 1)
                            var v46 uint32 = uint32(v13)
                            _ = v45
                            _ = v46
                        } else {
                            fmt.Println(v2)
                        }
                        _ = v44
                    }
                    var v47 uint16 = genFn0()
                    _ = fi39
                    _ = v47
                }
                switch -867 {
                case 1:
                    v37 = (v37 | 4294967295)
                case 32767:
                    var v48 int32 = 0
                    _ = v48
                }
                _ = v38
            } else {
                if false {
                    fmt.Println(v2)
                    switch v14 {
                    case -982:
                        var v49 uint32 = (v25 | v25)
                        _ = v49
                    case -492:
                        if false {
                            v2 = 119
                        }
                    }
                    if (!(false && (!true))) {
                        v11 = (0.0 * (((1e10 + (v11 + 3.14159)) / v11) / (v11 + v11)))
                    }
                } else {
                    fmt.Println(v3, v14, v37)
                    for fi50 := 0; fi50 < 3; fi50++ {
                        if true {
                            v11 += -2.25
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            var v51 int32 = 42
                            v1 = (v3 << ((4294967295) & 63))
                            v2 = v2
                            _ = v51
                        } else {
                            var v52 int32 = 140
                            if len(v17) > 0 {
                                v17[0] = (fi50 >> ((0) & 63))
                            }
                            v3 = v1
                            _ = v52
                        }
                        _ = fi50
                    }
                }
            }
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    v53 := GenS6{F0: (^(v2))}
    v1 -= uint(v25)
    var v54 float64 = v11
    for fi55 := 0; fi55 < 17; fi55++ {
        for fi56 := 0; fi56 < 8; fi56++ {
            v4 = (v4 + v4)
            if true {
                var v57 int = v14
                var v58 float64 = (0.0 - v54)
                switch 246 {
                case 32767:
                    var v59 uint32 = 0
                    _ = v59
                case 127:
                    var v60 uint8 = 255
                    _ = v60
                default:
                    fmt.Println(fi56, v25)
                }
                _ = v57
                _ = v58
            }
            _ = fi56
        }
        fmt.Println(v13, v13)
        var v61 int32 = 1
        _ = fi55
        _ = v61
    }
    for fi62 := 0; fi62 < 12; fi62++ {
        var v63 uint64 = 851
        if true {
            for fi64 := 0; fi64 < 1; fi64++ {
                v65 := GenS7{F0: true}
                _ = fi64
                _ = v65
            }
            switch v14 {
            case 32767:
                v54 += v54
            default:
                v3 -= 18446744073709551615
            }
            for fi66 := 0; fi66 < 5; fi66++ {
                if ((fi62 > (fi66 % ((fi66) | 1))) && true) {
                    if false {
                        if (false && (-152 == 1)) {
                            v13 -= (v13 * uint16(v25))
                            _ = len(v17)
                            var v67 uint16 = v13
                            var v68 int16 = 1
                            var v69 int32 = -2147483648
                            v14 += (fi62 / (((fi66 | -9223372036854775808)) | 1))
                            var v70 string = ""
                            _ = v67
                            _ = v68
                            _ = v69
                            _ = v70
                        }
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        v71 := GenS3{F0: false}
                        _ = v71
                    } else {
                        fmt.Println(v11, v25, v14)
                        v1 = uint(v25)
                    }
                }
                _ = fi66
            }
        }
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v53
    _ = v54
}
