package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint8; F1 string; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint; F2 int64 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 int16) int {
    var v4 float32 = -2.25
    if true {
        if (!(true || true)) {
            var v5 int32 = 2147483647
            _ = v5
        } else {
            v6 := GenS1{F0: 42}
            for fi7 := 0; fi7 < 9; fi7++ {
                for fi8 := 0; fi8 < 8; fi8++ {
                    p0 = p0
                    v9 := []int{int(p0), 32768, 2147483647}
                    _ = fi8
                    _ = v9
                }
                var v10 int64 = 9223372036854775807
                if false {
                    v4 -= v4
                }
                _ = fi7
                _ = v10
            }
            _ = v6
        }
        if (!false) {
            v4 -= v4
            v11 := []int{1, 32767, 0}
            _ = v11
        }
    } else {
        fmt.Println(v4, p0)
    }
    p0 += p0
    switch 32768 {
    case 127:
        if (!true) {
            var v12 int8 = -128
            p0 = (p0 & -479)
            _ = v12
        } else {
            v13 := GenS2{F0: (p0 == 1000)}
            if true {
                if false {
                    v14 := []int{127, -171, 9223372036854775807, 9223372036854775807}
                    for fi15 := 0; fi15 < 7; fi15++ {
                        for fi16 := 0; fi16 < 20; fi16++ {
                            p0 -= p0
                            p0 = p0
                            v4 += v4
                            _ = fi16
                        }
                        fmt.Println(v4)
                        _ = fi15
                    }
                    _ = v14
                }
            } else {
                switch "" {
                case "\n\t":
                    for fi17 := 0; fi17 < 1; fi17++ {
                        if (42 != 414) {
                            switch fi17 {
                            case 9223372036854775807:
                                v18 := GenS1{F0: 42}
                                _ = v18
                            case 765:
                                v19 := GenS3{F0: -35}
                                _ = v19
                            }
                        }
                        if (!true) {
                            var v20 float32 = (v4 * v4)
                            v21 := []int{}
                            if true {
                                fmt.Println(v20)
                                switch fi17 {
                                case 4294967295:
                                    p0 += 188
                                case 2147483647:
                                    if (!(!((p0 | p0) < (p0 << ((796) & 63))))) {
                                        v22 := GenS2{F0: false}
                                        for fi23 := 0; fi23 < 12; fi23++ {
                                            v20 += 1.5
                                            var v24 int32 = 2147483647
                                            _ = fi23
                                            _ = v24
                                        }
                                        var v25 uint64 = 1
                                        _ = v22
                                        _ = v25
                                    } else {
                                        for fi26 := 0; fi26 < 3; fi26++ {
                                            var v27 bool = false
                                            if len(v21) >= 2 {
                                                v21 = v21[1:len(v21)]
                                            }
                                            switch (fi17 - 2147483647) {
                                            case 4294967295:
                                                if (796 >= 4294967295) {
                                                    var v28 uint64 = 0
                                                    p0 = (p0 * p0)
                                                    _ = v28
                                                } else {
                                                    if ((-(fi17)) == fi26) {
                                                        var v29 int16 = 1000
                                                        _ = v29
                                                    }
                                                }
                                            }
                                            _ = fi26
                                            _ = v27
                                        }
                                    }
                                }
                            }
                            _ = v20
                            _ = v21
                        }
                        _ = fi17
                    }
                case "":
                    var v30 int8 = -128
                    _ = v30
                case "foo bar":
                    fmt.Println(v4)
                default:
                    var v31 uint16 = 65535
                    _ = v31
                }
                fmt.Println(p0)
            }
            v32 := []int{535, 42}
            _ = v13
            _ = v32
        }
    case 42:
        var v33 int16 = p0
        _ = v33
    }
    var v34 uint8 = 255
    _ = p0
    _ = v34
    _ = v4
    return 128
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 362
    var v3 string = "a"
    fmt.Println(v2, v3)
    if true {
        var v35 uint16 = 546
        for fi36 := 0; fi36 < 11; fi36++ {
            var v37 uint64 = 1
            var v38 int = fi36
            _ = fi36
            _ = v37
            _ = v38
        }
        v3 = "\n\t"
        _ = v35
    } else {
        for fi39 := 0; fi39 < 3; fi39++ {
            delete(v1, v3)
            _ = fi39
        }
        switch -283 {
        case 32767:
            fmt.Println(v3)
        }
        switch v3 {
        case "hello":
            switch -1 {
            case 855:
                v3 = ""
            default:
                var v40 int8 = -128
                _ = v40
            }
        case "a":
            v2 -= 2147483647
        case "foo bar":
            v3 = ((v3 + (v3 + "foo bar")) + (v3 + v3))
        default:
            for fi41 := 0; fi41 < 8; fi41++ {
                v3 = v3
                _ = fi41
            }
        }
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 uint32 = 2
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        v3 = (v3 + ("\n\t" + ((("" + (v3 + (("\n\t" + v3) + v3))) + "\n\t") + v3)))
        _ = fi44
    }
    switch (v3 + "\n\t") {
    case "hello":
        for fi45 := 0; fi45 < 10; fi45++ {
            v3 = ""
            _ = fi45
        }
    default:
        if (2 == 42) {
            switch v3 {
            case "hello":
                switch genFn0(42) {
                case 128:
                    var v46 int32 = 0
                    _ = v46
                case 4294967295:
                    v47 := []int{genFn0(-454), 461, 9223372036854775807}
                    _ = v47
                }
            case "\n\t":
                v48 := GenS2{F0: false}
                _ = v48
            default:
                switch v3 {
                case "a":
                    fmt.Println(v2, v3)
                default:
                    v49 := []int{2147483647, -806}
                    _ = v49
                }
            }
            v50 := v1[v3]
            v2 -= (v2 ^ v2)
            _ = v50
        } else {
            if false {
                var v51 int32 = 42
                _ = v51
            } else {
                var v52 int64 = -465
                v2 = (v2 / (((v2 >> ((42) & 63))) | 1))
                if true {
                    if (42 < 42) {
                        v3 = "hello"
                    }
                    v3 = v3
                    if true {
                        if (438 != 65535) {
                            fmt.Println(v52, v52, v2)
                        } else {
                            for fi53 := 0; fi53 < 13; fi53++ {
                                v54 := GenS3{F0: 56}
                                _ = fi53
                                _ = v54
                            }
                            for fi55 := 0; fi55 < 13; fi55++ {
                                switch (fi55 + 179) {
                                case -931:
                                    v56 := GenS3{F0: 42}
                                    _ = v56
                                case 4294967295:
                                    var v57 uint8 = 74
                                    _ = v57
                                case 2147483647:
                                    v58 := v1[("" + v3)]
                                    _ = v58
                                default:
                                    for fi59 := 0; fi59 < 4; fi59++ {
                                        v60 := GenS3{F0: 127}
                                        if false {
                                            v61 := GenS4{F0: (v2 % ((0) | 1)), F1: v52}
                                            if ((fi55 * (fi59 % ((42) | 1))) >= (fi55 << ((18446744073709551615) & 63))) {
                                                v3 = v3
                                            } else {
                                                var v62 int8 = 86
                                                fmt.Println(v62)
                                                _ = v62
                                            }
                                            _ = v61
                                        } else {
                                            fmt.Println(v52)
                                            var v63 uint32 = (v2 % ((v2) | 1))
                                            v64 := GenS5{F0: 42, F1: v3, F2: 0}
                                            _ = v63
                                            _ = v64
                                        }
                                        _ = fi59
                                        _ = v60
                                    }
                                }
                                v3 = "foo bar"
                                var v65 int16 = -190
                                _ = fi55
                                _ = v65
                            }
                            for fi66 := 0; fi66 < 6; fi66++ {
                                v3 = v3
                                v67 := GenS6{F0: 694, F1: 42, F2: v52}
                                _ = len(v1)
                                _ = fi66
                                _ = v67
                            }
                        }
                        v68 := v1["a"]
                        _ = v68
                    }
                }
                _ = v52
            }
            fmt.Println(v2, v2)
            var v69 int16 = 32767
            _ = v69
        }
    }
    v3 = (v3 + ((v3 + v3) + v3))
    for fi70 := 0; fi70 < 17; fi70++ {
        v71 := GenS4{F0: 4294967295, F1: 1}
        var v72 int32 = 2147483647
        _ = fi70
        _ = v71
        _ = v72
    }
    var v73 uint32 = 4294967295
    var v74 bool = (((0.1 + -2.25) / 3.14159) >= 0.1)
    v75 := GenS3{F0: 107}
    switch "hello" {
    case "\n\t":
        var v76 int32 = 2147483647
        _ = v76
    default:
        var v77 float32 = (1.5 - 1.5)
        _ = v77
    }
    var v78 uint8 = 127
    v73 = v73
    if true {
        v78 = (^(v78))
    }
    fmt.Println(v73, v73, v2)
    switch -270 {
    case 4294967295:
        fmt.Println(v74, v3)
    case 32768:
        for fi79 := 0; fi79 < 20; fi79++ {
            v2 = v73
            for fi80 := 0; fi80 < 5; fi80++ {
                if (v74 && (v74 || false)) {
                    v3 = ""
                    var v81 int16 = -32768
                    if (0 < 0) {
                        v82 := GenS6{F0: 0, F1: 18446744073709551615, F2: 41}
                        _ = v82
                    }
                    _ = v81
                }
                _ = fi80
            }
            var v83 uint32 = (v73 << ((0) & 63))
            _ = fi79
            _ = v83
        }
    case 1:
        fmt.Println(v78, v74)
    }
    v84 := GenS7{F0: 130}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v73
    _ = v74
    _ = v75
    _ = v78
    _ = v84
}
