package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 int32; F1 int; F2 int32 }
type GenS8 struct { F0 int8; F1 string }
func genFn0() uint64 {
    var v5 uint = 0
    fmt.Println(v5)
    _ = v5
    return 171
}
func genFn1(p0 string, p1 string, p2 uint8) string {
    if false {
        v6 := GenS1{F0: 0, F1: 776}
        _ = v6
    }
    v7 := GenS1{F0: 65535, F1: 65535}
    var v8 int = 114
    var v9 uint32 = 42
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v8)
            switch fi11 {
            case -9223372036854775808:
                var v12 int = (v8 << ((478) & 63))
                _ = v12
            case 9223372036854775807:
                var v13 int16 = 32767
                _ = v13
            }
            _ = fi11
        }
        fmt.Println(v9)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 558
    var v3 float64 = 0.1
    var v4 float64 = ((1.5 - 1e10) + float64((0.1 - (((0.1 * 1e10) * 3.14159) / 1.0))))
    var v14 int32 = -117
    v15 := v1[2147483647]
    if (v4 == -2.25) {
        for fi16 := 0; fi16 < 14; fi16++ {
            switch "foo bar" {
            case "hello":
                fmt.Println(v3, v3, v4)
            default:
                v14 += v14
            }
            v17 := GenS1{F0: 32767, F1: v2}
            switch (v15 << ((4294967295) & 63)) {
            case 387:
                v2 -= 18446744073709551615
            case -9223372036854775808:
                var v18 float64 = 1e10
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        for fi19 := 0; fi19 < 13; fi19++ {
            v20 := []int{(-(v15)), v15, (^(v15)), -748}
            v21 := v1[2147483648]
            v15 = (v21 << (((v2 % ((86) | 1))) & 63))
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v14 = v14
    v22 := []int{(v15 | (v15 | -2147483648)), v15, v15}
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 uint32 = 1
        v15 = v15
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        if (0.0 >= 1e10) {
            var v26 int32 = (^(v14))
            if false {
                if len(v22) > 0 {
                    v22[0] = 128
                }
                v3 = float64(0.0)
                switch (("foo bar" + "foo bar") + ("" + "")) {
                case "foo bar":
                    v27 := v1[215]
                    _ = v27
                case "a":
                    if true {
                        v28 := GenS1{F0: 65535, F1: v2}
                        var v29 uint8 = 128
                        _ = v28
                        _ = v29
                    }
                }
            }
            fmt.Println(v4, fi25)
            _ = v26
        } else {
            var v30 uint64 = 42
            v4 = 3.14159
            _ = v30
        }
        if (genFn0() == 4294967295) {
            fmt.Println(fi25, v2)
            switch fi25 {
            case 185:
                if ((v14 >> (((v2 >> ((v2) & 63))) & 63)) <= int32(v15)) {
                    v31 := GenS1{F0: 492, F1: (v2 / (((^(v2))) | 1))}
                    _ = v31
                } else {
                    if true {
                        for fi32 := 0; fi32 < 10; fi32++ {
                            _ = len(v1)
                            var v33 uint16 = 134
                            _ = fi32
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 2; fi34++ {
                            if false {
                                v35 := GenS2{F0: int16(fi34)}
                                _ = v35
                            }
                            var v36 int32 = (v14 + (v14 / ((v14) | 1)))
                            if false {
                                switch fi34 {
                                case -827:
                                    v37 := GenS3{F0: 42, F1: uint8(v2), F2: 65535}
                                    _ = v37
                                }
                            }
                            if false {
                                var v38 uint = v2
                                v39 := GenS2{F0: -32768}
                                _ = v38
                                _ = v39
                            } else {
                                if true {
                                    var v40 uint16 = 0
                                    v14 -= v14
                                    _ = v40
                                }
                                var v41 float64 = float64(3.14159)
                                _ = v41
                            }
                            _ = fi34
                            _ = v36
                        }
                    } else {
                        v22 = append(v22, (^(fi25)))
                        v4 += 0.1
                        if (42 <= 2) {
                            if true {
                                v42 := GenS1{F0: 981, F1: (v2 & (^(v2)))}
                                for fi43 := 0; fi43 < 15; fi43++ {
                                    switch -1 {
                                    case -581:
                                        v2 = 42
                                    case -46:
                                        if ((-2.25 / ((v3 - 0.0) + v4)) <= (3.14159 / (((((v3 - v4) / v3) / v3) - 0.1) * -2.25))) {
                                            var v44 uint16 = 0
                                            var v45 uint16 = (v44 << ((v2) & 63))
                                            if (v2 != (v2 << (((v2 % ((v2) | 1))) & 63))) {
                                                v2 += v2
                                            }
                                            if ((v14 ^ 1) == int32(fi43)) {
                                                v44 -= 32767
                                                v15 = fi43
                                            } else {
                                                delete(v1, 0)
                                                if (18446744073709551615 < 2) {
                                                    fmt.Println(v44, v4, v15)
                                                    if false {
                                                        fmt.Println(fi43, v14, v3)
                                                        v45 = v44
                                                    } else {
                                                        switch (("foo bar" + "") + "\n\t") {
                                                        case "foo bar":
                                                            v46 := GenS2{F0: 1}
                                                            _ = v46
                                                        case "a":
                                                            v4 = (float64(1e10) / (v4 + v3))
                                                        }
                                                    }
                                                    var v47 int64 = 903
                                                    _ = v47
                                                }
                                                if false {
                                                    fmt.Println(v44, v15, v44)
                                                    switch (fi25 ^ fi43) {
                                                    case 1:
                                                        var v48 uint32 = 534
                                                        _ = v48
                                                    case -9223372036854775808:
                                                        v3 = ((((0.1 / (-2.25 / v3)) * 0.0) - ((v4 - ((v4 / v3) * (v3 + 1.5))) + float64((((1e10 / 0.1) - 1.5) * 0.0)))) / ((0.0 + ((v4 / 1.5) + v3)) - (float64((float32(v3) * -2.25)) / 1.0)))
                                                        v1[65535] = int(v2)
                                                    }
                                                    v49 := GenS4{F0: 32767, F1: "hello"}
                                                    _ = v49
                                                } else {
                                                    var v50 uint8 = uint8(v2)
                                                    fmt.Println(fi25)
                                                    _ = v50
                                                }
                                            }
                                            var v51 uint32 = 2147483647
                                            v52 := GenS5{F0: 255}
                                            _ = v44
                                            _ = v45
                                            _ = v51
                                            _ = v52
                                        } else {
                                            if true {
                                                v53 := GenS2{F0: 1}
                                                v15 -= 42
                                                v54 := GenS6{F0: "", F1: v2}
                                                _ = v53
                                                _ = v54
                                            } else {
                                                var v55 float32 = 1.5
                                                v56 := GenS3{F0: 2, F1: 175, F2: 792}
                                                _ = v55
                                                _ = v56
                                            }
                                            v14 = v14
                                            switch (v15 - 916) {
                                            case 465:
                                                v14 = v14
                                                delete(v1, 528)
                                            case -58:
                                                var v57 uint16 = 2
                                                _ = v57
                                            default:
                                                var v58 int32 = (v14 ^ v14)
                                                _ = v58
                                            }
                                            if len(v22) > 0 {
                                                v22[0] = (fi43 << ((v2) & 63))
                                            }
                                            fmt.Println(fi43)
                                            v4 += v4
                                            v4 += v3
                                        }
                                    case 4294967295:
                                        var v59 int32 = int32(v2)
                                        _ = v59
                                    default:
                                        v2 = (v2 << ((v2) & 63))
                                    }
                                    v60 := GenS5{F0: 2}
                                    _ = fi43
                                    _ = v60
                                }
                                v61 := v1[32]
                                _ = v42
                                _ = v61
                            }
                        }
                    }
                }
            default:
                fmt.Println(v4, v3)
            }
            fmt.Println(v4, fi25, v3)
        } else {
            switch (("hello" + "hello") + "foo bar") {
            case "hello":
                for fi62 := 0; fi62 < 20; fi62++ {
                    v63 := GenS3{F0: 749, F1: 0, F2: uint32(v2)}
                    _ = fi62
                    _ = v63
                }
            case "foo bar":
                v3 -= v4
            }
        }
        _ = fi25
    }
    v2 += v2
    if true {
        switch "" {
        case "hello":
            v64 := GenS2{F0: 62}
            _ = v64
        default:
            v15 = v15
        }
        var v65 bool = (!((v15 + -609) <= 127))
        v66 := GenS7{F0: (-(v14)), F1: (v15 + (v15 ^ v15)), F2: v14}
        _ = v65
        _ = v66
    } else {
        v2 += (v2 - 18446744073709551615)
    }
    if (-1 >= -1) {
        _ = len(v1)
    } else {
        v67 := GenS8{F0: 127, F1: ("\n\t" + "foo bar")}
        switch (v15 >> (((v2 | (v2 - 1))) & 63)) {
        case 2147483647:
            v2 = v2
        case -611:
            var v68 uint64 = 1
            _ = v68
        case 32768:
            for fi69 := 0; fi69 < 18; fi69++ {
                fmt.Println(v3)
                var v70 int32 = v14
                var v71 string = ""
                _ = fi69
                _ = v70
                _ = v71
            }
        default:
            var v72 bool = ((("\n\t" + "\n\t") + ((((("hello" + (("\n\t" + "\n\t") + "a")) + ("hello" + ("a" + ""))) + "\n\t") + "\n\t") + ((("a" + "") + "\n\t") + "foo bar"))) != "a")
            _ = v72
        }
        switch v15 {
        case 0:
            var v73 int8 = 2
            _ = v73
        case 2147483647:
            v2 = v2
        case 4294967295:
            if true {
                var v74 uint = 945
                _ = v74
            }
        default:
            fmt.Println(v15, v4)
        }
        _ = v67
    }
    var v75 float64 = v3
    fmt.Println(v2, v75, v4)
    var v76 uint32 = 0
    v22 = append(v22, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v75
    _ = v76
}
