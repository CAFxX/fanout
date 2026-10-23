package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int32; F1 int16; F2 int8 }
type GenS6 struct { F0 uint; F1 int16; F2 uint }
type GenS7 struct { F0 int }

func main() {
    var v1 int64 = -1
    var v2 int64 = -786
    var v3 uint16 = 0
    var v4 int8 = 100
    v5 := make(map[int]bool)
    var v6 uint32 = uint32(v3)
    v7 := v5[32768]
    switch "a" {
    case "":
        v8 := []int{}
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int = (^(fi9))
            v5[(fi9 + v10)] = (1 != 1)
            v10 += (v10 | 9223372036854775807)
            _ = fi9
            _ = v10
        }
    case "a":
        v11 := GenS1{F0: (v2 ^ v1), F1: 65535, F2: int32(v1)}
        _ = v11
    default:
        var v12 uint8 = 72
        _ = v12
    }
    switch 689 {
    case 4294967295:
        v2 -= -1
    default:
        if false {
            if true {
                if v7 {
                    if false {
                        if false {
                            if false {
                                var v13 uint32 = v6
                                v14 := GenS1{F0: v1, F1: 156, F2: -1}
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    var v16 float32 = (float32(float64(1e10)) - 0.1)
                                    v1 = 785
                                    v17 := GenS2{F0: (v3 ^ v3)}
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                                _ = v13
                                _ = v14
                            } else {
                                var v18 uint8 = 138
                                var v19 string = ("foo bar" + "\n\t")
                                v20 := GenS3{F0: 1, F1: (v6 / ((208) | 1)), F2: 925}
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            if (!(!v7)) {
                                var v21 int16 = 32767
                                v21 = (v21 * 1000)
                                _ = v21
                            }
                        } else {
                            for fi22 := 0; fi22 < 8; fi22++ {
                                for fi23 := 0; fi23 < 12; fi23++ {
                                    var v24 float64 = 0.1
                                    v25 := []int{fi23, (^(fi23))}
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                var v26 int16 = 108
                                var v27 int = fi22
                                _ = fi22
                                _ = v26
                                _ = v27
                            }
                            var v28 int64 = (v1 >> ((418) & 63))
                            v7 = v7
                            _ = v28
                        }
                        if v7 {
                            switch "a" {
                            case "a":
                                v1 = (v1 << ((4294967295) & 63))
                            case "":
                                if ("hello" <= ((("" + "a") + "hello") + "foo bar")) {
                                    var v29 uint32 = (^(v6))
                                    var v30 bool = v7
                                    _ = v29
                                    _ = v30
                                } else {
                                    fmt.Println(v6, v6)
                                }
                            case "foo bar":
                                if v7 {
                                    var v31 float64 = 1.5
                                    _ = v31
                                }
                            default:
                                v32 := GenS2{F0: v3}
                                _ = v32
                            }
                            for fi33 := 0; fi33 < 19; fi33++ {
                                var v34 uint = 65535
                                _ = fi33
                                _ = v34
                            }
                        } else {
                            if v7 {
                                var v35 string = "\n\t"
                                _ = v35
                            } else {
                                fmt.Println(v3, v7)
                                fmt.Println(v4, v4)
                            }
                            if (uint(v6) >= 18446744073709551615) {
                                v1 += (v1 << ((18446744073709551615) & 63))
                                switch "a" {
                                case "foo bar":
                                    if true {
                                        var v36 uint16 = v3
                                        _ = v36
                                    }
                                }
                                v4 = (v4 % (((v4 | -1)) | 1))
                            }
                            v37 := GenS4{F0: v6, F1: ""}
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 17; fi38++ {
                            if (-32768 == 1) {
                                v39 := GenS5{F0: -585, F1: 42, F2: v4}
                                var v40 string = "foo bar"
                                var v41 float32 = 1e10
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                            _ = fi38
                        }
                    } else {
                        var v42 int16 = -546
                        var v43 uint = 18446744073709551615
                        var v44 uint8 = 127
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                    v2 -= -9223372036854775808
                }
            } else {
                if (3.14159 <= (((0.1 + (0.0 * (float32(1.5) / 1e10))) + ((0.1 - float32((((float64(3.14159) + (3.14159 * 3.14159)) + float64(0.0)) / ((((0.0 * (-2.25 * float64(3.14159))) + 0.1) / float64(-2.25)) / 1.5)))) + -2.25)) + 1e10)) {
                    v2 += v2
                    var v45 int64 = (v2 | v1)
                    _ = v45
                } else {
                    var v46 uint32 = (^(v6))
                    v6 = v46
                    _ = v46
                }
                v3 -= (^(v3))
            }
            v47 := GenS1{F0: v2, F1: 255, F2: 2147483647}
            for fi48 := 0; fi48 < 5; fi48++ {
                switch (("hello" + ((("hello" + ((("hello" + "foo bar") + "") + "hello")) + ("hello" + "\n\t")) + ("hello" + "a"))) + "") {
                case "":
                    v3 = 255
                }
                v49 := GenS4{F0: (v6 ^ (v6 & v6)), F1: ""}
                v1 = (-(v2))
                _ = fi48
                _ = v49
            }
            _ = v47
        } else {
            fmt.Println(v4, v6, v7)
        }
    }
    if false {
        if v7 {
            switch 63 {
            case 398:
                fmt.Println(v6, v2)
            case 127:
                if (!false) {
                    switch (("a" + ("\n\t" + "foo bar")) + ("a" + ("foo bar" + "a"))) {
                    case "a":
                        var v50 int16 = 467
                        _ = v50
                    case "foo bar":
                        v7 = (111 > 127)
                    }
                    v7 = v7
                    var v51 uint32 = 758
                    _ = v51
                } else {
                    v3 = 0
                    v7 = v7
                }
            }
            if (v2 <= int64(v4)) {
                switch "" {
                case "":
                    v52 := []int{}
                    _ = v52
                }
            } else {
                v53 := []int{-820}
                var v54 int = 2147483647
                _ = v53
                _ = v54
            }
        }
        var v55 int = -9223372036854775808
        for fi56 := 0; fi56 < 20; fi56++ {
            v55 -= v55
            for fi57 := 0; fi57 < 16; fi57++ {
                fmt.Println(v2)
                var v58 uint16 = v3
                _ = fi57
                _ = v58
            }
            _ = fi56
        }
        _ = v55
    }
    var v59 uint16 = v3
    v5[-2147483648] = (9223372036854775807 < 32768)
    var v60 uint64 = 1
    v61 := GenS1{F0: (^(v2)), F1: 255, F2: 1}
    fmt.Println(v3, v60, v60)
    v62 := GenS4{F0: v6, F1: ("a" + "a")}
    for fi63 := 0; fi63 < 14; fi63++ {
        switch (-(fi63)) {
        case -316:
            for fi64 := 0; fi64 < 17; fi64++ {
                if (v59 > v59) {
                    switch 2147483647 {
                    case -9223372036854775808:
                        if false {
                            var v65 int16 = 0
                            v2 += 1
                            v66 := []int{}
                            _ = v65
                            _ = v66
                        }
                    case 812:
                        delete(v5, fi64)
                    case 42:
                        v60 -= 1
                    }
                    v7 = (v59 != (v3 ^ 1))
                    v67 := GenS1{F0: (v2 * 0), F1: uint(v60), F2: -1}
                    _ = v67
                } else {
                    if v7 {
                        if (fi63 > int(v60)) {
                            _ = len(v5)
                            if v7 {
                                if (127 == 76) {
                                    v7 = v7
                                    var v68 int32 = 42
                                    if ((!v7) || v7) {
                                        var v69 uint32 = (v6 * (v6 << ((475) & 63)))
                                        _ = v69
                                    }
                                    _ = v68
                                }
                                v70 := GenS6{F0: 18446744073709551615, F1: -1, F2: 42}
                                delete(v5, fi63)
                                _ = v70
                            } else {
                                v4 = -128
                                var v71 int64 = -1
                                _ = v71
                            }
                            var v72 uint32 = v6
                            _ = v72
                        } else {
                            if v7 {
                                v60 = 42
                            }
                            v4 -= (v4 ^ 42)
                        }
                        v59 += uint16(v60)
                    } else {
                        fmt.Println(v60)
                    }
                    var v73 bool = (21 != v4)
                    _ = v73
                }
                fmt.Println(v7, v2)
                if (!v7) {
                    v2 = 703
                    fmt.Println(fi63, v2, v1)
                } else {
                    v74 := GenS3{F0: 2, F1: 65535, F2: 4294967295}
                    _ = len(v5)
                    v3 += (v59 << ((631) & 63))
                    _ = v74
                }
                _ = fi64
            }
        }
        if false {
            if v7 {
                var v75 bool = (true && (!true))
                _ = v75
            } else {
                v76 := GenS2{F0: (^(v3))}
                for fi77 := 0; fi77 < 13; fi77++ {
                    switch 253 {
                    case -71:
                        v78 := GenS7{F0: int(v2)}
                        _ = v78
                    default:
                        fmt.Println(v6)
                    }
                    v7 = v7
                    _ = fi77
                }
                _ = v76
            }
            var v79 string = "foo bar"
            _ = v79
        }
        _ = fi63
    }
    v60 = v60
    fmt.Println(v6, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
    _ = v7
}
