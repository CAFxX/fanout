package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int32; F2 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint32; F1 uint }

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]string)
    var v3 int16 = 1
    var v4 int32 = -2147483648
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]string)
    v4 = -813
    if (v3 < (v3 + v3)) {
        v4 -= (v4 & v4)
        v1 = v1
        if false {
            v2[""] = ""
            switch "hello" {
            case "\n\t":
                for fi7 := 0; fi7 < 17; fi7++ {
                    for fi8 := 0; fi8 < 16; fi8++ {
                        v9 := GenS1{F0: 51}
                        v1 = 1
                        v10 := GenS1{F0: 127}
                        v3 -= v3
                        var v11 uint = 4294967295
                        v1 = v1
                        _ = fi8
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                    _ = fi7
                }
            case "":
                var v12 uint32 = 42
                _ = v12
            case "hello":
                v3 = v3
            }
            if (4294967295 != v1) {
                var v13 uint64 = 42
                var v14 float32 = float32(-2.25)
                if (((18446744073709551615 <= 173) || false) || (2 > 32767)) {
                    if true {
                        for fi15 := 0; fi15 < 9; fi15++ {
                            v1 = uint32(v13)
                            _ = fi15
                        }
                        v1 -= v1
                        v1 += (v1 * 861)
                    }
                    var v16 float64 = (1e10 - 0.1)
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 15; fi17++ {
                        for fi18 := 0; fi18 < 10; fi18++ {
                            v19 := GenS1{F0: 0}
                            _ = fi18
                            _ = v19
                        }
                        var v20 uint8 = uint8(v13)
                        var v21 int = (fi17 - fi17)
                        _ = fi17
                        _ = v20
                        _ = v21
                    }
                    v5[(v1 - v1)] = (1.5 == ((v14 * v14) / 0.1))
                }
                _ = v13
                _ = v14
            }
        } else {
            var v22 int8 = 48
            _ = v22
        }
    }
    if true {
        var v23 uint64 = uint64(v1)
        _ = v23
    } else {
        v24 := GenS2{F0: (v4 + 817), F1: (-(v4)), F2: (false || (false && false))}
        var v25 float32 = 1e10
        _ = v24
        _ = v25
    }
    if true {
        if (1e10 <= -2.25) {
            v1 = v1
            for fi26 := 0; fi26 < 17; fi26++ {
                v27 := []int{fi26}
                v28 := GenS2{F0: v4, F1: -1, F2: (!false)}
                _ = fi26
                _ = v27
                _ = v28
            }
            var v29 uint64 = 863
            _ = v29
        }
    } else {
        if (true || (true && true)) {
            if true {
                var v30 float64 = 1e10
                var v31 int16 = v3
                _ = v30
                _ = v31
            } else {
                for fi32 := 0; fi32 < 7; fi32++ {
                    switch fi32 {
                    case 0:
                        v3 += v3
                    case 128:
                        for fi33 := 0; fi33 < 18; fi33++ {
                            switch "\n\t" {
                            case "hello":
                                v34 := GenS1{F0: 1}
                                _ = v34
                            default:
                                if true {
                                    fmt.Println(v3, v3)
                                    v35 := []int{(fi32 % (((fi33 * (fi33 / (((fi32 >> ((4294967295) & 63))) | 1)))) | 1)), -444, fi33}
                                    _ = v35
                                }
                            }
                            _ = fi33
                        }
                    case 382:
                        for fi36 := 0; fi36 < 10; fi36++ {
                            var v37 uint64 = 4294967295
                            var v38 string = ("" + "\n\t")
                            v3 += (v3 << ((1) & 63))
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    default:
                        for fi39 := 0; fi39 < 11; fi39++ {
                            v40 := []int{}
                            switch ("" + "a") {
                            case "\n\t":
                                v41 := GenS1{F0: 255}
                                _ = v41
                            }
                            _ = fi39
                            _ = v40
                        }
                    }
                    v42 := GenS1{F0: 204}
                    if true {
                        if false {
                            v1 = v1
                            fmt.Println(v1)
                        }
                        for fi43 := 0; fi43 < 11; fi43++ {
                            v44 := GenS2{F0: (v4 | (v4 / ((v4) | 1))), F1: (v4 % ((-1) | 1)), F2: ((v3 % ((int16(fi32)) | 1)) >= 0)}
                            if true {
                                v1 = v1
                                var v45 int64 = 327
                                v46 := []int{353, int(v1), fi32, -200}
                                _ = v45
                                _ = v46
                            }
                            v47 := GenS3{F0: 390, F1: v4, F2: 4294967295}
                            if ("\n\t" >= "a") {
                                v48 := GenS2{F0: (v4 ^ -610), F1: (^(v4)), F2: (fi43 == fi43)}
                                _ = v48
                            }
                            _ = fi43
                            _ = v44
                            _ = v47
                        }
                        v4 -= (^(v4))
                    } else {
                        switch ("foo bar" + (("a" + "") + "\n\t")) {
                        case "\n\t":
                            var v49 int8 = -128
                            _ = v49
                        default:
                            v50 := GenS1{F0: 196}
                            _ = v50
                        }
                        v3 = v3
                        var v51 int8 = -1
                        _ = v51
                    }
                    _ = fi32
                    _ = v42
                }
                if ((v1 + 892) != 670) {
                    var v52 uint16 = 430
                    _ = v52
                }
                var v53 bool = (v1 <= v1)
                _ = v53
            }
            fmt.Println(v3)
        }
        v54 := []int{42, 359, -465, -2147483648}
        _ = v54
    }
    for fi55 := 0; fi55 < 8; fi55++ {
        var v56 int8 = -32
        v57 := []int{42}
        v57 = append(v57, (fi55 & 4294967295))
        _ = fi55
        _ = v56
        _ = v57
    }
    fmt.Println(v4, v4, v4)
    var v58 int8 = 100
    if (!(false && true)) {
        if ("a" == "a") {
            if true {
                var v59 bool = true
                for fi60 := 0; fi60 < 11; fi60++ {
                    if ((-(fi60)) > fi60) {
                        v4 -= int32(v58)
                    }
                    _ = fi60
                }
                var v61 uint64 = 796
                _ = v59
                _ = v61
            }
            for fi62 := 0; fi62 < 1; fi62++ {
                v63 := GenS3{F0: 226, F1: v4, F2: 4294967295}
                _ = fi62
                _ = v63
            }
            if ((((0.0 * float64(1.5)) * -2.25) - 3.14159) < (3.14159 - -2.25)) {
                for fi64 := 0; fi64 < 19; fi64++ {
                    delete(v2, (((("\n\t" + "") + "hello") + "foo bar") + ""))
                    var v65 uint16 = 32767
                    if (2 >= 2) {
                        switch fi64 {
                        case 42:
                            switch fi64 {
                            case -70:
                                if (1 > uint64(v1)) {
                                    for fi66 := 0; fi66 < 5; fi66++ {
                                        if true {
                                            v67 := GenS4{F0: 285, F1: ""}
                                            switch (fi64 & fi64) {
                                            case -2147483648:
                                                var v68 float64 = ((3.14159 + (1.5 - float64((-2.25 + 0.0)))) - (1.5 + 1e10))
                                                _ = v68
                                            case 2147483647:
                                                v69 := []int{fi66, (fi66 | 128), (fi64 * fi66), 4294967295}
                                                v65 += (v65 | 838)
                                                if false {
                                                    v70 := GenS3{F0: (v65 | 0), F1: v4, F2: 18446744073709551615}
                                                    var v71 bool = (int64(fi66) == 42)
                                                    _ = v70
                                                    _ = v71
                                                }
                                                _ = v69
                                            case 32767:
                                                v72 := GenS1{F0: uint8(v65)}
                                                _ = v72
                                            }
                                            _ = v67
                                        }
                                        v73 := v5[v1]
                                        v74 := GenS5{F0: uint32(v65), F1: 4294967295}
                                        _ = fi66
                                        _ = v73
                                        _ = v74
                                    }
                                }
                            case 32768:
                                var v75 bool = false
                                _ = v75
                            }
                        case -848:
                            var v76 int8 = (v58 >> ((787) & 63))
                            _ = v76
                        }
                        switch fi64 {
                        case -2147483648:
                            var v77 uint64 = 4294967295
                            _ = v77
                        case -1:
                            fmt.Println(v65, v1)
                        case 933:
                            var v78 int64 = 355
                            _ = v78
                        default:
                            var v79 float32 = (1.5 + (1e10 * 3.14159))
                            _ = v79
                        }
                    } else {
                        v65 -= (v65 / (((v65 * (v65 * (v65 ^ (v65 & (v65 + 977)))))) | 1))
                        var v80 int = -2147483648
                        _ = v80
                    }
                    _ = fi64
                    _ = v65
                }
            }
        }
        var v81 string = (("foo bar" + "") + ((("\n\t" + ("\n\t" + ((("a" + ("a" + "hello")) + "") + "foo bar"))) + "hello") + "\n\t"))
        var v82 uint64 = 42
        _ = v81
        _ = v82
    }
    var v83 uint16 = 65535
    v84 := GenS3{F0: 2, F1: (v4 << ((65535) & 63)), F2: 401}
    v85 := GenS4{F0: 857, F1: "hello"}
    v86 := GenS1{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v83
    _ = v84
    _ = v85
    _ = v86
}
