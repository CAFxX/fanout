package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 uint8; F1 int64 }
func genFn0(p0 int32, p1 uint32) int32 {
    switch "hello" {
    case "hello":
        var v5 int = 127
        _ = v5
    case "":
        var v6 int16 = int16(p0)
        _ = v6
    case "\n\t":
        fmt.Println(p0, p0)
    }
    switch ("a" + ("foo bar" + "a")) {
    case "":
        v7 := []int{4294967295, 2147483647, 32767}
        _ = v7
    case "\n\t":
        var v8 float32 = 3.14159
        _ = v8
    case "hello":
        v9 := []int{int(p0), 897}
        _ = v9
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 335
}
func genFn1(p0 uint, p1 string) uint64 {
    p1 = ("\n\t" + p1)
    var v10 uint64 = uint64(p0)
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := []int{(-(fi11)), fi11, (fi11 * -1), (fi11 | (fi11 / ((42) | 1)))}
        v10 = v10
        _ = fi11
        _ = v12
    }
    switch p1 {
    case "hello":
        v13 := []int{574}
        _ = v13
    default:
        p0 += 65535
    }
    _ = p0
    _ = p1
    _ = v10
    return (v10 & v10)
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 1
    var v3 float64 = 1.5
    var v4 uint64 = 2
    v3 += float64((v1 / v1))
    if true {
        for fi14 := 0; fi14 < 2; fi14++ {
            v1 = 0.0
            var v15 uint16 = 0
            _ = fi14
            _ = v15
        }
        fmt.Println(v2)
        v16 := []int{-608, 32767, 32767}
        _ = v16
    } else {
        v17 := []int{}
        var v18 float32 = v1
        _ = v17
        _ = v18
    }
    if (v3 == (1e10 - 1.5)) {
        if (v1 != v1) {
            var v19 uint = 65535
            fmt.Println(v3, v19)
            _ = v19
        } else {
            var v20 float32 = float32(v3)
            _ = v20
        }
        v3 = float64(((v1 * v1) / (v1 * ((v1 + (1e10 / v1)) - 0.1))))
    } else {
        v1 += float32((v3 / (((v3 + ((v3 + v3) - ((3.14159 / v3) * (1e10 / (v3 + 3.14159))))) * v3) - (1.5 * v3))))
    }
    if false {
        var v21 float32 = v1
        _ = v21
    } else {
        switch 1 {
        case 128:
            v22 := []int{464, -1, 2147483647, 35}
            _ = v22
        case 2147483647:
            v1 -= (1e10 - (v1 / (-2.25 * -2.25)))
        }
    }
    var v23 bool = false
    fmt.Println(v1, v3, v1)
    var v24 float32 = (v1 * v1)
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := []int{fi25, (fi25 + fi25), fi25}
        switch fi25 {
        case -2147483648:
            for fi27 := 0; fi27 < 18; fi27++ {
                v26 = append(v26, (fi27 + (-(fi25))))
                _ = fi27
            }
        case 9223372036854775807:
            v2 = (^(v2))
        }
        var v28 float32 = float32(((v3 / float64(float32((1e10 * (v3 / v3))))) * ((((float64((v1 - v24)) - 3.14159) * (((v3 * -2.25) * float64(v24)) + v3)) - (-2.25 + (v3 - float64(0.0)))) + (float64(((float32(3.14159) / float32(v3)) - float32(-2.25))) / 1.0))))
        _ = fi25
        _ = v26
        _ = v28
    }
    switch 671 {
    case 32767:
        v2 = (v2 % (((v2 << ((286) & 63))) | 1))
    case 42:
        if (42 == 0) {
            var v29 int64 = 9223372036854775807
            _ = v29
        }
    default:
        if v23 {
            for fi30 := 0; fi30 < 1; fi30++ {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v32 := []int{fi31, (fi31 * -152), -72}
                    if false {
                        var v33 uint8 = 60
                        if (-1 >= -734) {
                            v2 = (v2 / (((v2 ^ 239)) | 1))
                            v34 := GenS1{F0: 1, F1: (v2 * v2)}
                            v3 += (-2.25 / ((-2.25 / v3) / (v3 - ((v3 + v3) + v3))))
                            switch ((("\n\t" + "") + "foo bar") + "a") {
                            case "foo bar":
                                if len(v32) > 0 {
                                    v32[0] = 42
                                }
                            case "hello":
                                v24 += 1.5
                            case "a":
                                switch fi31 {
                                case -954:
                                    var v35 int = fi30
                                    _ = v35
                                }
                            default:
                                switch "foo bar" {
                                case "":
                                    if len(v32) > 0 {
                                        v32[0] = fi30
                                    }
                                }
                            }
                            v2 = (v2 / ((v2) | 1))
                            _ = v34
                        }
                        fmt.Println(v33, v24)
                        _ = v33
                    }
                    if (-2.25 > v3) {
                        switch fi30 {
                        case 2147483647:
                            var v36 bool = ((fi31 >> ((0) & 63)) < (fi31 << ((4294967295) & 63)))
                            _ = v36
                        case 682:
                            var v37 int32 = -183
                            _ = v37
                        }
                        var v38 bool = true
                        _ = v38
                    } else {
                        var v39 string = "a"
                        if true {
                            var v40 float64 = (v3 / float64(v24))
                            switch ("a" + "") {
                            case "foo bar":
                                v4 += (^(v4))
                            case "a":
                                v1 = 0.0
                            case "\n\t":
                                var v41 float32 = v1
                                _ = v41
                            }
                            var v42 int = (fi31 & fi31)
                            _ = v40
                            _ = v42
                        } else {
                            var v43 int16 = -1
                            fmt.Println(fi31)
                            v32 = append(v32, fi31)
                            _ = v43
                        }
                        _ = v39
                    }
                    _ = fi31
                    _ = v32
                }
                var v44 int = -471
                _ = fi30
                _ = v44
            }
        }
    }
    var v45 int = -9223372036854775808
    if false {
        if true {
            v46 := GenS2{F0: 45, F1: v4, F2: (v4 + 42)}
            v4 = (v4 ^ (v4 >> ((42) & 63)))
            v47 := GenS3{F0: 128, F1: uint8(v4)}
            _ = v46
            _ = v47
        } else {
            switch "foo bar" {
            case "foo bar":
                v48 := []int{v45, int(v4), v45, v45}
                _ = v48
            }
            var v49 int = (v45 + v45)
            _ = v49
        }
        switch "" {
        case "":
            for fi50 := 0; fi50 < 7; fi50++ {
                v51 := GenS1{F0: 65535, F1: v2}
                var v52 string = "foo bar"
                v53 := GenS4{F0: -403, F1: -9223372036854775808}
                _ = fi50
                _ = v51
                _ = v52
                _ = v53
            }
        default:
            var v54 uint8 = 42
            _ = v54
        }
        if (v23 || false) {
            var v55 int32 = 1
            _ = v55
        } else {
            for fi56 := 0; fi56 < 7; fi56++ {
                for fi57 := 0; fi57 < 6; fi57++ {
                    var v58 uint16 = 255
                    _ = fi57
                    _ = v58
                }
                _ = fi56
            }
        }
    } else {
        v3 += ((float64(v24) - v3) + v3)
        fmt.Println(v24, v4)
        switch (("" + "foo bar") + "hello") {
        case "foo bar":
            for fi59 := 0; fi59 < 7; fi59++ {
                var v60 uint16 = (v2 & (^(v2)))
                for fi61 := 0; fi61 < 15; fi61++ {
                    var v62 uint16 = v60
                    _ = fi61
                    _ = v62
                }
                _ = fi59
                _ = v60
            }
        case "\n\t":
            if v23 {
                fmt.Println(v4)
                v63 := GenS1{F0: 1, F1: 255}
                _ = v63
            }
        }
    }
    if false {
        v64 := GenS5{F0: v4, F1: v4}
        _ = v64
    } else {
        var v65 int8 = 127
        v66 := []int{v45, (v45 << ((18446744073709551615) & 63)), (v45 / ((651) | 1))}
        _ = v65
        _ = v66
    }
    var v67 int16 = 741
    for fi68 := 0; fi68 < 9; fi68++ {
        v23 = v23
        switch (("" + ("a" + "\n\t")) + "a") {
        case "hello":
            v69 := []int{-2147483648}
            _ = v69
        case "foo bar":
            if ((!(((^(v67)) == 32767) && v23)) && true) {
                switch (-(v45)) {
                case 127:
                    if v23 {
                        v2 += 65535
                        fmt.Println(v24, v45, v1)
                    }
                case -2147483648:
                    switch (^(v45)) {
                    case 1:
                        var v70 uint = 18446744073709551615
                        _ = v70
                    case 815:
                        switch (^(v45)) {
                        case 9223372036854775807:
                            var v71 string = "hello"
                            _ = v71
                        case -9223372036854775808:
                            var v72 int64 = 0
                            _ = v72
                        case 0:
                            fmt.Println(v23, v2)
                        default:
                            for fi73 := 0; fi73 < 1; fi73++ {
                                var v74 int32 = -573
                                var v75 int32 = genFn0((^(v74)), 1)
                                if v23 {
                                    v76 := []int{(v45 ^ (fi68 % ((v45) | 1))), -9223372036854775808}
                                    _ = v76
                                }
                                _ = fi73
                                _ = v74
                                _ = v75
                            }
                        }
                    default:
                        v77 := GenS2{F0: int8(v4), F1: genFn1(uint(v4), "foo bar"), F2: 2}
                        _ = v77
                    }
                }
                v78 := []int{-2147483648}
                v67 -= -1
                _ = v78
            } else {
                switch (fi68 << ((4294967295) & 63)) {
                case 2147483647:
                    v79 := GenS6{F0: uint8(v4), F1: -9223372036854775808}
                    _ = v79
                case -974:
                    var v80 uint32 = 0
                    _ = v80
                }
                v24 += 1e10
            }
        case "a":
            var v81 uint32 = 34
            _ = v81
        }
        var v82 uint32 = 0
        _ = fi68
        _ = v82
    }
    v23 = (255 != 128)
    v83 := []int{}
    fmt.Println(v4, v3, v45)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v83:", len(v83), v83)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v67
    _ = v83
}
