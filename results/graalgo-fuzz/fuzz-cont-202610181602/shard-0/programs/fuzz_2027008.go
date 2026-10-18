package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 bool; F2 int64 }
type GenS4 struct { F0 int8; F1 uint; F2 int64 }
type GenS5 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint64, p1 int16, p2 bool) uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        fmt.Println(fi4, p2, p0)
        var v5 uint8 = 128
        switch 9223372036854775807 {
        case -186:
            var v6 bool = p2
            _ = v6
        case -2147483648:
            var v7 int32 = int32(fi4)
            _ = v7
        case 127:
            if (1 != 0) {
                v8 := []int{(fi4 & fi4), fi4, int(p1), fi4}
                fmt.Println(fi4, p1)
                var v9 uint = 18446744073709551615
                _ = v8
                _ = v9
            } else {
                for fi10 := 0; fi10 < 18; fi10++ {
                    p1 = (p1 & 1000)
                    _ = fi10
                }
            }
        default:
            var v11 float64 = 0.1
            _ = v11
        }
        _ = fi4
        _ = v5
    }
    v12 := GenS1{F0: 86}
    p2 = ((("a" + "") + "hello") > "")
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 65535
}
func genFn1(p0 int64) uint16 {
    var v13 float64 = 1e10
    var v14 int64 = p0
    _ = p0
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 float64 = (3.14159 + ((0.0 * -2.25) * 0.0))
    var v2 uint16 = 1
    var v3 float64 = float64(float32(3.14159))
    v15 := []int{}
    v3 = v1
    if true {
        fmt.Println(v1, v1, v3)
        v16 := GenS1{F0: v2}
        switch (("" + ("" + ("" + "\n\t"))) + ("foo bar" + ("\n\t" + "foo bar"))) {
        case "hello":
            v17 := GenS2{F0: 42}
            _ = v17
        case "a":
            switch "\n\t" {
            case "a":
                if true {
                    var v18 float32 = 1e10
                    _ = v18
                } else {
                    if (!(true || (255 >= 33))) {
                        v19 := GenS1{F0: v2}
                        if len(v15) > 0 {
                            v15[0] = 42
                        }
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            for fi21 := 0; fi21 < 7; fi21++ {
                                var v22 int16 = -32768
                                v23 := GenS2{F0: 0}
                                var v24 int16 = (v22 >> ((genFn0(375, (v22 + (^(v22))), false)) & 63))
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            _ = fi20
                        }
                        switch -2147483648 {
                        case 32767:
                            var v25 uint64 = 1
                            _ = v25
                        default:
                            fmt.Println(v1)
                        }
                        var v26 uint32 = 576
                        _ = v26
                    }
                    v15 = append(v15, 0)
                }
            case "":
                var v27 int64 = 0
                _ = v27
            case "\n\t":
                var v28 float64 = 1e10
                _ = v28
            default:
                if (0 != uint(v2)) {
                    fmt.Println(v1, v1, v3)
                    for fi29 := 0; fi29 < 3; fi29++ {
                        var v30 uint = 2
                        fmt.Println(v1, v30, v1)
                        _ = fi29
                        _ = v30
                    }
                    v1 -= v3
                }
            }
        case "\n\t":
            if ("\n\t" == "a") {
                if (!false) {
                    v31 := GenS2{F0: 255}
                    _ = v31
                } else {
                    if false {
                        switch "hello" {
                        case "hello":
                            var v32 string = ((("hello" + "hello") + "hello") + "")
                            _ = v32
                        case "\n\t":
                            if true {
                                if (!(0.1 > ((3.14159 + 1.5) / 3.14159))) {
                                    v33 := GenS3{F0: false, F1: true, F2: -283}
                                    if (255 == 0) {
                                        var v34 int8 = 42
                                        v35 := GenS2{F0: 0}
                                        if len(v15) > 0 {
                                            v15[0] = 4294967295
                                        }
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = v33
                                }
                                var v36 string = "a"
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                _ = v36
                            }
                        default:
                            for fi37 := 0; fi37 < 4; fi37++ {
                                switch "foo bar" {
                                case "hello":
                                    v3 = v3
                                case "a":
                                    _ = len(v15)
                                }
                                _ = fi37
                            }
                        }
                    } else {
                        for fi38 := 0; fi38 < 3; fi38++ {
                            if (!(true && (-32768 > 0))) {
                                if true {
                                    var v39 int32 = 16
                                    _ = v39
                                }
                                v3 = float64((0.1 + float32(float64((0.0 * 1e10)))))
                            }
                            for fi40 := 0; fi40 < 13; fi40++ {
                                v41 := GenS3{F0: true, F1: false, F2: 42}
                                _ = fi40
                                _ = v41
                            }
                            _ = fi38
                        }
                        fmt.Println(v2)
                        v42 := GenS3{F0: true, F1: (42 >= 0), F2: 42}
                        _ = v42
                    }
                }
            }
        }
        _ = v16
    } else {
        var v43 uint8 = 0
        _ = v43
    }
    v44 := GenS1{F0: (v2 | (v2 << ((4294967295) & 63)))}
    for fi45 := 0; fi45 < 8; fi45++ {
        var v46 int64 = -1
        _ = fi45
        _ = v46
    }
    fmt.Println(v1, v3, v3)
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    for fi47 := 0; fi47 < 6; fi47++ {
        if ((fi47 > (fi47 | -2147483648)) && true) {
            v2 -= (v2 * 32767)
            v1 += 1.5
        } else {
            switch 32768 {
            case 128:
                v1 = v1
            case 32767:
                for fi48 := 0; fi48 < 20; fi48++ {
                    v49 := GenS3{F0: (-2147483648 != -1), F1: true, F2: 0}
                    v15 = append(v15, 187)
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    var v50 float64 = 0.1
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            case 182:
                var v51 int64 = -9223372036854775808
                _ = v51
            }
            var v52 float64 = v1
            v2 -= 1
            _ = v52
        }
        _ = len(v15)
        _ = fi47
    }
    switch "foo bar" {
    case "\n\t":
        _ = len(v15)
    default:
        for fi53 := 0; fi53 < 7; fi53++ {
            if len(v15) > 0 {
                v15[0] = -954
            }
            _ = fi53
        }
    }
    switch "foo bar" {
    case "hello":
        if len(v15) > 0 {
            v15[0] = 0
        }
    case "a":
        var v54 int8 = 102
        _ = v54
    default:
        for fi55 := 0; fi55 < 4; fi55++ {
            if ((v2 % (((v2 | v2)) | 1)) < (v2 & (^(v2)))) {
                fmt.Println(fi55, v3, v3)
                for fi56 := 0; fi56 < 7; fi56++ {
                    var v57 int32 = 0
                    var v58 string = ("" + "")
                    var v59 int32 = (v57 / ((int32(v2)) | 1))
                    _ = fi56
                    _ = v57
                    _ = v58
                    _ = v59
                }
            }
            for fi60 := 0; fi60 < 12; fi60++ {
                if false {
                    var v61 int16 = -671
                    v15 = append(v15, fi60)
                    var v62 int8 = 127
                    _ = v61
                    _ = v62
                } else {
                    switch "" {
                    case "hello":
                        v1 += ((v1 + 1e10) - (float64(((3.14159 / 1.5) / 3.14159)) + float64(3.14159)))
                    case "foo bar":
                        v3 = v3
                    case "":
                        var v63 int8 = -102
                        _ = v63
                    default:
                        v1 = v1
                        v64 := GenS3{F0: true, F1: (!true), F2: 9223372036854775807}
                        if len(v15) > 0 {
                            v15[0] = fi55
                        }
                        _ = v64
                    }
                    switch (("foo bar" + "a") + "hello") {
                    case "a":
                        v2 -= v2
                    default:
                        v1 = (v3 + 1.5)
                    }
                }
                v2 = genFn1(9223372036854775807)
                _ = fi60
            }
            _ = fi55
        }
    }
    if false {
        if false {
            for fi65 := 0; fi65 < 1; fi65++ {
                switch "\n\t" {
                case "hello":
                    _ = len(v15)
                }
                _ = fi65
            }
        }
        v66 := GenS1{F0: (^(v2))}
        if true {
            switch ("a" + "\n\t") {
            case "foo bar":
                switch -895 {
                case 215:
                    switch ((("a" + ("\n\t" + "foo bar")) + "hello") + ("a" + ("\n\t" + "hello"))) {
                    case "\n\t":
                        v3 -= 3.14159
                    case "foo bar":
                        v67 := GenS4{F0: 102, F1: genFn0(4294967295, 1000, true), F2: 0}
                        _ = v67
                    case "hello":
                        v68 := GenS5{F0: 818, F1: 65535}
                        _ = v68
                    }
                case 941:
                    var v69 float64 = (v3 / (v3 - 1e10))
                    _ = v69
                case 358:
                    v70 := GenS5{F0: -32768, F1: (^(v2))}
                    _ = v70
                }
            case "\n\t":
                var v71 int64 = 9223372036854775807
                _ = v71
            case "a":
                fmt.Println(v3, v3, v3)
            default:
                if len(v15) > 0 {
                    v15[0] = 2147483647
                }
            }
            switch 2147483647 {
            case 4294967295:
                var v72 uint8 = 67
                _ = v72
            case 925:
                fmt.Println(v1, v1, v3)
            case 0:
                fmt.Println(v2, v2, v2)
            default:
                if true {
                    v2 = (v2 % ((434) | 1))
                    v2 = genFn1(9223372036854775807)
                } else {
                    if len(v15) > 0 {
                        v15[0] = 1
                    }
                }
            }
        }
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v44
}
