package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 int16; F1 int16; F2 uint64 }
func genFn0(p0 int, p1 string, p2 int16) int64 {
    switch p1 {
    case "\n\t":
        p1 = p1
    case "":
        var v6 int = (p0 | (-(p0)))
        _ = v6
    case "foo bar":
        v7 := GenS1{F0: -2147483648}
        _ = v7
    default:
        v8 := []int{}
        _ = v8
    }
    v9 := []int{127, (p0 ^ (p0 + p0)), (p0 | (^(p0)))}
    if len(v9) > 0 {
        v9[0] = p0
    }
    var v10 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 uint64, p1 uint32, p2 uint8) int32 {
    switch "\n\t" {
    case "foo bar":
        p2 -= 255
    case "hello":
        switch 149 {
        case 32767:
            switch "foo bar" {
            case "":
                p1 = p1
            }
        }
    case "\n\t":
        var v11 int64 = 0
        _ = v11
    }
    p0 -= (p0 % ((4294967295) | 1))
    var v12 uint = 18446744073709551615
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -1
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]int)
    var v3 int16 = 32767
    var v4 string = ""
    v5 := make(map[string]bool)
    var v13 int16 = -484
    for fi14 := 0; fi14 < 1; fi14++ {
        delete(v2, v4)
        v15 := GenS1{F0: (fi14 / ((128) | 1))}
        v4 = "\n\t"
        _ = fi14
        _ = v15
    }
    fmt.Println(v4, v3)
    switch ("a" + v4) {
    case "":
        if ((v3 / (((v13 % (((-(v3))) | 1))) | 1)) != (v13 & 42)) {
            for fi16 := 0; fi16 < 14; fi16++ {
                switch fi16 {
                case -2147483648:
                    var v17 int = (fi16 & fi16)
                    _ = v17
                case 467:
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 int8 = 127
                        _ = fi18
                        _ = v19
                    }
                case 128:
                    if true {
                        var v20 bool = (!(((true && (false && ("\n\t" == v4))) && (2147483647 != 2)) && (!(false && (v1 <= v1)))))
                        _ = v20
                    } else {
                        v13 += 504
                        delete(v5, v4)
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v3 = (v13 << ((458) & 63))
                            v22 := []int{fi16, 1, (fi16 * 4294967295), (fi16 ^ -1)}
                            var v23 float64 = 1.5
                            var v24 bool = false
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                }
                v1 = 42
                _ = fi16
            }
            switch v4 {
            case "\n\t":
                var v25 int32 = 0
                _ = v25
            case "a":
                switch -2147483648 {
                case 32767:
                    v4 = (((v4 + ("" + ("a" + v4))) + "\n\t") + "\n\t")
                case 1:
                    v26 := []int{1, 32768}
                    _ = v26
                default:
                    switch "a" {
                    case "a":
                        var v27 string = v4
                        _ = v27
                    case "foo bar":
                        if (355 < 836) {
                            v4 = v4
                            v28 := v5[v4]
                            _ = v28
                        }
                    default:
                        if true {
                            _ = len(v2)
                            v29 := GenS1{F0: 4294967295}
                            _ = v29
                        }
                    }
                }
            case "foo bar":
                v30 := []int{int(v1)}
                _ = v30
            }
        }
    }
    fmt.Println(v13, v13)
    v4 = v4
    v3 = v3
    var v31 int32 = genFn1(0, 406, v1)
    v32 := []int{32768}
    if (false || false) {
        for fi33 := 0; fi33 < 9; fi33++ {
            v34 := GenS2{F0: -1}
            var v35 uint8 = v1
            if false {
                for fi36 := 0; fi36 < 20; fi36++ {
                    if len(v32) >= 2 {
                        v32 = v32[1:len(v32)]
                    }
                    v37 := GenS2{F0: (fi33 ^ (fi33 / ((69) | 1)))}
                    if len(v32) >= 2 {
                        v32 = v32[1:len(v32)]
                    }
                    _ = fi36
                    _ = v37
                }
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        switch -2147483648 {
        case 128:
            for fi38 := 0; fi38 < 12; fi38++ {
                if (!(!false)) {
                    v3 -= -32768
                    switch v4 {
                    case "foo bar":
                        var v39 uint32 = 1
                        _ = v39
                    case "\n\t":
                        v4 = (v4 + v4)
                    default:
                        if true {
                            for fi40 := 0; fi40 < 14; fi40++ {
                                v41 := GenS3{F0: fi38, F1: (v3 % ((v3) | 1))}
                                switch fi40 {
                                case 127:
                                    if false {
                                        v3 = v13
                                        var v42 uint32 = 2147483648
                                        v13 += -1
                                        _ = v42
                                    } else {
                                        delete(v5, "a")
                                        var v43 uint16 = 255
                                        _ = v43
                                    }
                                case -851:
                                    switch (fi38 | fi40) {
                                    case -9223372036854775808:
                                        fmt.Println(v3, v1)
                                    default:
                                        if (false || false) {
                                            v4 = v4
                                            switch 127 {
                                            case 32768:
                                                switch ((((v4 + "hello") + v4) + v4) + v4) {
                                                case "hello":
                                                    v44 := GenS3{F0: 128, F1: v3}
                                                    _ = v44
                                                case "\n\t":
                                                    if false {
                                                        v13 += (v3 ^ v13)
                                                        var v45 int16 = v3
                                                        var v46 int64 = 1
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        switch fi40 {
                                                        case 528:
                                                            if (!(!true)) {
                                                                switch (v4 + (v4 + v4)) {
                                                                case "hello":
                                                                    fmt.Println(v31, fi40)
                                                                }
                                                            }
                                                        case -285:
                                                            var v47 bool = false
                                                            v47 = (!v47)
                                                            if false {
                                                                var v48 int8 = 100
                                                                _ = v48
                                                            }
                                                            _ = v47
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v3, v1)
                                                }
                                            case 32767:
                                                v49 := GenS4{F0: (v1 << ((2) & 63)), F1: (-(v3))}
                                                _ = v49
                                            case -1:
                                                fmt.Println(v4, v4, v1)
                                            default:
                                                if false {
                                                    v50 := GenS5{F0: 505, F1: 1}
                                                    v3 = (v3 / ((-32768) | 1))
                                                    _ = v50
                                                } else {
                                                    if (((1e10 * float64(float32(((3.14159 * ((float64(float32(0.0)) / (((1.5 + float64((-2.25 / 1.0))) + 0.0) / (float64(-2.25) / 1.0))) + 0.1)) / (((-2.25 / 1e10) - 3.14159) / 0.1))))) >= 0.0) || false) {
                                                        fmt.Println(v3)
                                                    }
                                                }
                                            }
                                        }
                                    }
                                case 128:
                                    fmt.Println(v1, fi38)
                                default:
                                    switch fi40 {
                                    case 127:
                                        fmt.Println(v1, v3, v4)
                                    }
                                }
                                if true {
                                    v5[v4] = true
                                    var v51 string = ""
                                    _ = v51
                                } else {
                                    delete(v5, v4)
                                    v52 := GenS3{F0: fi40, F1: v3}
                                    _ = v52
                                }
                                _ = fi40
                                _ = v41
                            }
                            var v53 int = fi38
                            v54 := GenS4{F0: v1, F1: (v13 & (^(v3)))}
                            _ = v53
                            _ = v54
                        } else {
                            for fi55 := 0; fi55 < 11; fi55++ {
                                v13 = 1000
                                _ = fi55
                            }
                        }
                    }
                    for fi56 := 0; fi56 < 13; fi56++ {
                        v57 := v5[((v4 + "") + "a")]
                        v1 = v1
                        _ = fi56
                        _ = v57
                    }
                } else {
                    var v58 int = 743
                    var v59 int64 = 9223372036854775807
                    v60 := GenS5{F0: v3, F1: 4294967295}
                    _ = v58
                    _ = v59
                    _ = v60
                }
                _ = fi38
            }
        case 9223372036854775807:
            v61 := GenS3{F0: 4294967295, F1: (v13 * (v3 << ((65535) & 63)))}
            _ = v61
        }
    } else {
        var v62 bool = (!false)
        delete(v5, v4)
        var v63 int = 2147483647
        _ = v62
        _ = v63
    }
    v13 = v3
    v64 := GenS4{F0: (v1 / ((v1) | 1)), F1: 42}
    fmt.Println(v31, v1)
    v1 += (^(v1))
    var v65 int = 32768
    for fi66 := 0; fi66 < 18; fi66++ {
        if ((v31 + v31) != 2147483647) {
            switch v4 {
            case "\n\t":
                var v67 uint16 = 2
                _ = v67
            case "hello":
                v68 := v2[v4]
                _ = v68
            case "":
                _ = len(v2)
            default:
                switch (v4 + v4) {
                case "foo bar":
                    v32 = append(v32, (-(v65)))
                default:
                    v5[v4] = (false || true)
                }
            }
            v69 := GenS6{F0: v3, F1: v3, F2: 18446744073709551615}
            for fi70 := 0; fi70 < 17; fi70++ {
                var v71 uint32 = 65535
                _ = fi70
                _ = v71
            }
            _ = v69
        }
        if false {
            var v72 string = (v4 + "hello")
            delete(v2, v4)
            _ = v72
        }
        _ = fi66
    }
    for fi73 := 0; fi73 < 13; fi73++ {
        v31 += v31
        _ = fi73
    }
    var v74 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v74
}
