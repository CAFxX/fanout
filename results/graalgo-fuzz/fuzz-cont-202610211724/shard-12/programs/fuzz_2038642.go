package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int; F1 uint8; F2 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint }
func genFn0() int32 {
    if (-1 < 487) {
        var v4 int8 = 42
        v4 -= v4
        var v5 uint8 = 1
        _ = v4
        _ = v5
    } else {
        v6 := GenS1{F0: (true && true), F1: ((-128 == 100) && false), F2: -2147483648}
        v7 := GenS2{F0: 19, F1: 18446744073709551615, F2: 1}
        var v8 int8 = 100
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 int64 = 729
    v9 += (v9 + v9)
    v10 := []int{-5}
    _ = v10
    _ = v9
    return 547
}
func genFn1() uint16 {
    var v11 uint8 = 128
    switch ("\n\t" + ("\n\t" + "hello")) {
    case "\n\t":
        var v12 int8 = 43
        _ = v12
    default:
        v11 = v11
    }
    v13 := GenS3{F0: 433, F1: 32767}
    for fi14 := 0; fi14 < 12; fi14++ {
        switch 594 {
        case -9223372036854775808:
            v11 -= v11
        case -1:
            for fi15 := 0; fi15 < 20; fi15++ {
                v16 := GenS2{F0: 166, F1: 520, F2: 0}
                _ = fi15
                _ = v16
            }
        case 128:
            var v17 int = -293
            _ = v17
        default:
            switch "foo bar" {
            case "a":
                v18 := GenS2{F0: uint16(v11), F1: 306, F2: 1}
                _ = v18
            case "hello":
                var v19 uint64 = 1
                _ = v19
            case "\n\t":
                v11 = v11
            }
        }
        _ = fi14
    }
    if (!(1.5 > (((float64(3.14159) - 0.0) * 0.0) - 0.1))) {
        v20 := []int{0, 699, 127}
        for fi21 := 0; fi21 < 3; fi21++ {
            v20 = append(v20, (fi21 << ((1) & 63)))
            fmt.Println(v11, fi21)
            _ = fi21
        }
        _ = v20
    } else {
        v22 := []int{2147483647, 638, 86}
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        var v23 float64 = (0.0 - ((1.5 + -2.25) - 1.5))
        _ = v22
        _ = v23
    }
    _ = v11
    _ = v13
    return 2
}

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 712
    v3 := make(map[int]string)
    v24 := GenS3{F0: -1, F1: 71}
    for fi25 := 0; fi25 < 16; fi25++ {
        if ((!false) || (4294967295 < 2)) {
            v26 := GenS1{F0: ((2 == 1) || false), F1: (65535 <= 1), F2: -2147483648}
            _ = v26
        }
        _ = fi25
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        for fi28 := 0; fi28 < 4; fi28++ {
            switch (("\n\t" + "") + ("" + "foo bar")) {
            case "":
                var v29 string = ("" + "\n\t")
                _ = v29
            case "foo bar":
                var v30 int16 = 0
                _ = v30
            case "\n\t":
                var v31 uint = 0
                _ = v31
            default:
                v1 = (v1 + v1)
            }
            _ = fi28
        }
        fmt.Println(v2)
        var v32 float32 = 0.1
        _ = fi27
        _ = v32
    }
    delete(v3, -617)
    v2 = (^(v2))
    v33 := GenS1{F0: false, F1: (42 >= 4294967295), F2: -1}
    v1 -= (3.14159 - float64(1.5))
    v34 := v3[32768]
    fmt.Println(v2, v34)
    for fi35 := 0; fi35 < 9; fi35++ {
        v34 = "\n\t"
        for fi36 := 0; fi36 < 11; fi36++ {
            if (!true) {
                v37 := GenS1{F0: (42 > 362), F1: (0 <= 18446744073709551615), F2: 42}
                v2 = 355
                _ = v37
            } else {
                v38 := GenS4{F0: fi36, F1: 42, F2: v34}
                _ = v38
            }
            if (345 > 2) {
                v34 = v34
                if true {
                    if (!false) {
                        fmt.Println(fi35, fi36, v34)
                        if false {
                            var v39 bool = false
                            v34 = ""
                            v40 := GenS4{F0: 42, F1: 127, F2: "a"}
                            var v41 uint64 = 322
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    } else {
                        switch 151 {
                        case 961:
                            v42 := GenS5{F0: 1}
                            _ = v42
                        }
                        v1 -= v1
                        v2 += (v2 & genFn1())
                    }
                    var v43 bool = ((3.14159 + float64(1.5)) < 0.1)
                    _ = v43
                } else {
                    v1 += 1.5
                }
                v44 := GenS5{F0: 68}
                _ = v44
            }
            v1 = float64((float32((float64((float32(v1) - (1e10 - (1.5 * 3.14159)))) * 3.14159)) + -2.25))
            v1 = (((v1 - ((((v1 * v1) - 3.14159) * (v1 * v1)) - v1)) - v1) / ((v1 - 1e10) - v1))
            if (true && (2 <= 4294967295)) {
                switch 1 {
                case 1:
                    fmt.Println(v34, v1)
                    var v45 uint64 = 88
                    _ = v45
                default:
                    v46 := GenS4{F0: fi35, F1: 0, F2: (("\n\t" + v34) + (v34 + ((v34 + ("a" + v34)) + "\n\t")))}
                    _ = v46
                }
                fmt.Println(fi35, v2)
            } else {
                var v47 int16 = 1000
                var v48 int = (fi36 % (((fi35 << ((18446744073709551615) & 63))) | 1))
                _ = v47
                _ = v48
            }
            fmt.Println(fi36, v1)
            v49 := []int{(-(fi36)), -9223372036854775808, -2147483648}
            _ = fi36
            _ = v49
        }
        var v50 int8 = -128
        _ = fi35
        _ = v50
    }
    for fi51 := 0; fi51 < 11; fi51++ {
        var v52 uint64 = 4294967295
        v1 += (-2.25 + (((v1 + v1) * float64(0.1)) - v1))
        v53 := GenS2{F0: (v2 - 32767), F1: 943, F2: 42}
        _ = fi51
        _ = v52
        _ = v53
    }
    delete(v3, 979)
    v2 = (v2 / ((255) | 1))
    _ = len(v3)
    switch "\n\t" {
    case "a":
        if true {
            v1 -= v1
        }
    case "hello":
        fmt.Println(v34, v34, v34)
    default:
        fmt.Println(v34, v2)
    }
    for fi54 := 0; fi54 < 14; fi54++ {
        if (118 == -68) {
            v55 := GenS3{F0: -1, F1: v2}
            _ = v55
        } else {
            if true {
                if (v34 == (v34 + v34)) {
                    fmt.Println(v1, v2, fi54)
                } else {
                    v1 -= (0.1 / v1)
                    var v56 uint32 = 42
                    v56 += (v56 - (v56 & 906))
                    _ = v56
                }
            }
            var v57 uint64 = 18446744073709551615
            _ = v57
        }
        var v58 uint8 = 128
        for fi59 := 0; fi59 < 1; fi59++ {
            fmt.Println(fi54, v2)
            if false {
                if true {
                    if (42 == v58) {
                        switch v34 {
                        case "a":
                            v1 -= v1
                        case "\n\t":
                            v34 = v34
                        default:
                            switch v34 {
                            case "foo bar":
                                var v60 int = (-(fi54))
                                _ = v60
                            case "":
                                var v61 bool = ((!((true && ((65535 == 1) && false)) && true)) && (!((true || (!(255 >= 2))) && ((!(0.1 == float32(3.14159))) || (v34 >= v34)))))
                                _ = v61
                            case "\n\t":
                                v58 -= v58
                            default:
                                if true {
                                    v62 := []int{}
                                    fmt.Println(v58)
                                    _ = v62
                                } else {
                                    _ = len(v3)
                                }
                            }
                        }
                    }
                } else {
                    fmt.Println(v58, v34, v34)
                    switch (fi59 * 2147483647) {
                    case 42:
                        v58 = (v58 + v58)
                    case 128:
                        if (997 == 32767) {
                            switch (((v34 + v34) + "hello") + "hello") {
                            case "\n\t":
                                var v63 int64 = -1
                                _ = v63
                            case "hello":
                                v3[fi54] = v34
                            case "":
                                v64 := []int{fi54, 32768, (fi59 % ((32768) | 1)), fi54}
                                _ = v64
                            }
                            var v65 uint64 = 427
                            _ = v65
                        } else {
                            v3[(fi54 << ((255) & 63))] = v34
                        }
                    }
                    v66 := GenS6{F0: 1, F1: 861}
                    _ = v66
                }
                var v67 int = (fi54 << ((2) & 63))
                v2 += (^(v2))
                v58 = v58
                _ = v67
            }
            _ = fi59
        }
        _ = fi54
        _ = v58
    }
    var v68 uint32 = 2
    var v69 int = -441
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v68
    _ = v69
}
