package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 bool; F1 bool; F2 int }
func genFn0(p0 int16, p1 int32) int16 {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := []int{(fi4 | 4294967295)}
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 int64 = 0
            var v8 uint8 = 99
            var v9 uint32 = 2
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        if false {
            var v10 int8 = 0
            v11 := GenS1{F0: (-(p0))}
            _ = v10
            _ = v11
        } else {
            if len(v5) > 0 {
                v5[0] = (fi4 * (fi4 + fi4))
            }
        }
        _ = fi4
        _ = v5
    }
    var v12 int64 = 1
    _ = p0
    _ = p1
    _ = v12
    return (p0 * -1)
}
func genFn1(p0 int64) uint64 {
    var v13 uint32 = 42
    for fi14 := 0; fi14 < 13; fi14++ {
        v13 += 2147483647
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := GenS2{F0: uint64(v13)}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            if false {
                var v18 int16 = -512
                v13 = 42
                _ = v18
            } else {
                v19 := []int{(^(fi17)), (fi17 % ((0) | 1)), fi17, (fi17 >> ((1) & 63))}
                if ((1 != 127) && (1 <= 100)) {
                    p0 -= (p0 << ((2) & 63))
                    fmt.Println(fi17)
                }
                v20 := GenS3{F0: -12, F1: 32767, F2: 255}
                _ = v19
                _ = v20
            }
            v13 = (v13 ^ 2147483647)
            v21 := GenS1{F0: -32768}
            _ = fi17
            _ = v21
        }
        _ = fi14
    }
    v22 := []int{4294967295, 2147483647, 4294967295}
    var v23 uint16 = 42
    _ = p0
    _ = v13
    _ = v22
    _ = v23
    return 270
}

func main() {
    var v1 int8 = 127
    var v2 uint32 = 2
    var v3 int = 9223372036854775807
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 4; fi25++ {
            v3 = (fi24 | (fi25 - v3))
            v26 := GenS2{F0: 1}
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    if false {
        if true {
            var v27 int16 = 996
            switch "a" {
            case "\n\t":
                for fi28 := 0; fi28 < 1; fi28++ {
                    if true {
                        if false {
                            var v29 uint64 = 4294967295
                            _ = v29
                        }
                        switch ("" + "") {
                        case "":
                            v30 := GenS1{F0: (v27 | v27)}
                            _ = v30
                        case "hello":
                            v31 := []int{v3, v3, fi28}
                            _ = v31
                        }
                        if false {
                            if (v3 > (v3 ^ fi28)) {
                                v27 = (v27 / (((v27 & 861)) | 1))
                            } else {
                                switch "a" {
                                case "foo bar":
                                    for fi32 := 0; fi32 < 12; fi32++ {
                                        var v33 float32 = (3.14159 - 0.0)
                                        _ = fi32
                                        _ = v33
                                    }
                                }
                                v1 = -8
                            }
                            switch (^(v3)) {
                            case 127:
                                v2 = (v2 ^ (v2 << ((uint(v2)) & 63)))
                            case 312:
                                v34 := []int{1, fi28, v3, v3}
                                _ = v34
                            }
                            v27 = (v27 / (((v27 - v27)) | 1))
                        }
                    }
                    v35 := GenS1{F0: v27}
                    _ = fi28
                    _ = v35
                }
            case "hello":
                if (!(!((1 <= 321) && (-76 >= v1)))) {
                    fmt.Println(v3)
                }
            case "":
                v3 -= (v3 >> ((255) & 63))
            default:
                v27 = -32768
            }
            fmt.Println(v3, v27, v1)
            _ = v27
        } else {
            if false {
                fmt.Println(v1)
            } else {
                switch (v3 % (((^(v3))) | 1)) {
                case 726:
                    for fi36 := 0; fi36 < 8; fi36++ {
                        v2 = 2147483647
                        _ = fi36
                    }
                case 2147483647:
                    v1 = (v1 & v1)
                default:
                    var v37 float32 = (3.14159 * (0.0 / ((((1.5 / -2.25) - 1.5) / 0.1) / float32(3.14159))))
                    _ = v37
                }
                var v38 bool = false
                switch v3 {
                case 2147483647:
                    var v39 uint8 = 0
                    _ = v39
                }
                _ = v38
            }
            for fi40 := 0; fi40 < 2; fi40++ {
                switch fi40 {
                case 4294967295:
                    v41 := GenS1{F0: -32768}
                    _ = v41
                case -9223372036854775808:
                    var v42 float32 = -2.25
                    _ = v42
                default:
                    for fi43 := 0; fi43 < 3; fi43++ {
                        v2 = 42
                        v44 := []int{}
                        _ = fi43
                        _ = v44
                    }
                }
                v45 := GenS3{F0: 167, F1: 65535, F2: 603}
                fmt.Println(v2, v1)
                _ = fi40
                _ = v45
            }
            v2 = 1
        }
        if true {
            v1 = (v1 - v1)
            v46 := GenS2{F0: 2}
            for fi47 := 0; fi47 < 15; fi47++ {
                v3 = 42
                for fi48 := 0; fi48 < 16; fi48++ {
                    fmt.Println(v2)
                    switch "hello" {
                    case "":
                        v1 = (^(v1))
                    default:
                        var v49 int64 = 42
                        _ = v49
                    }
                    _ = fi48
                }
                _ = fi47
            }
            _ = v46
        } else {
            if ((0 != (v2 ^ v2)) || true) {
                var v50 int8 = (^(v1))
                _ = v50
            } else {
                for fi51 := 0; fi51 < 14; fi51++ {
                    var v52 float32 = 3.14159
                    _ = fi51
                    _ = v52
                }
                if (false || false) {
                    var v53 int16 = 42
                    var v54 uint8 = 1
                    _ = v53
                    _ = v54
                } else {
                    if (!false) {
                        for fi55 := 0; fi55 < 8; fi55++ {
                            v56 := GenS2{F0: genFn1(9223372036854775807)}
                            v57 := []int{4294967295, v3, 32768}
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                    } else {
                        if (42 >= 2) {
                            for fi58 := 0; fi58 < 8; fi58++ {
                                v1 -= v1
                                v2 = 992
                                var v59 int = (v3 << ((0) & 63))
                                _ = fi58
                                _ = v59
                            }
                            for fi60 := 0; fi60 < 18; fi60++ {
                                v61 := []int{v3, v3}
                                _ = fi60
                                _ = v61
                            }
                            v3 = v3
                        }
                        v62 := GenS1{F0: 1000}
                        _ = v62
                    }
                    v2 -= 42
                    v63 := GenS4{F0: (false || (false && ("\n\t" == "foo bar"))), F1: (0 <= -33), F2: v3}
                    _ = v63
                }
                if (!true) {
                    switch (v3 | (v3 / (((v3 - v3)) | 1))) {
                    case 32767:
                        v64 := GenS3{F0: -2147483648, F1: 255, F2: 2}
                        _ = v64
                    case 4294967295:
                        var v65 int8 = -115
                        _ = v65
                    }
                } else {
                    v1 -= (v1 - 0)
                    if (134 < 86) {
                        var v66 float32 = (3.14159 + (0.1 - 1.5))
                        if (true || (34 != 127)) {
                            var v67 uint32 = (v2 / ((65535) | 1))
                            v2 -= v67
                            _ = v67
                        } else {
                            v2 += (^(v2))
                            var v68 int32 = 1
                            _ = v68
                        }
                        v69 := GenS3{F0: 0, F1: 1, F2: 2}
                        _ = v66
                        _ = v69
                    } else {
                        var v70 float64 = float64(1e10)
                        if ((v3 << ((605) & 63)) != (v3 << ((0) & 63))) {
                            for fi71 := 0; fi71 < 13; fi71++ {
                                for fi72 := 0; fi72 < 7; fi72++ {
                                    var v73 int32 = -1
                                    v74 := []int{}
                                    _ = fi72
                                    _ = v73
                                    _ = v74
                                }
                                v75 := GenS2{F0: 42}
                                var v76 string = ("\n\t" + "")
                                _ = fi71
                                _ = v75
                                _ = v76
                            }
                            for fi77 := 0; fi77 < 19; fi77++ {
                                fmt.Println(v1, v2)
                                switch "hello" {
                                case "\n\t":
                                    if false {
                                        v70 -= v70
                                    }
                                case "foo bar":
                                    for fi78 := 0; fi78 < 17; fi78++ {
                                        v79 := []int{(fi77 & v3), fi78, (v3 % (((fi77 << ((uint(v2)) & 63))) | 1))}
                                        var v80 float64 = ((v70 - (v70 * ((v70 / (v70 - (v70 / v70))) / 1.0))) * 0.0)
                                        var v81 uint32 = (v2 + v2)
                                        _ = len(v79)
                                        if false {
                                            _ = len(v79)
                                            v2 = 1
                                        }
                                        _ = fi78
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                    }
                                case "hello":
                                    var v82 uint16 = 2
                                    _ = v82
                                }
                                if (v3 >= (fi77 + fi77)) {
                                    v83 := []int{-1, (fi77 + 9223372036854775807)}
                                    switch "" {
                                    case "\n\t":
                                        var v84 int32 = 2147483647
                                        _ = v84
                                    default:
                                        v70 -= v70
                                    }
                                    var v85 uint8 = 1
                                    _ = v83
                                    _ = v85
                                } else {
                                    fmt.Println(v2)
                                }
                                _ = fi77
                            }
                            for fi86 := 0; fi86 < 9; fi86++ {
                                if (false && (!(!true))) {
                                    var v87 uint16 = 32767
                                    for fi88 := 0; fi88 < 2; fi88++ {
                                        v1 = (v1 & v1)
                                        _ = fi88
                                    }
                                    _ = v87
                                }
                                _ = fi86
                            }
                        } else {
                            var v89 float64 = (v70 - v70)
                            _ = v89
                        }
                        _ = v70
                    }
                }
            }
            for fi90 := 0; fi90 < 19; fi90++ {
                for fi91 := 0; fi91 < 12; fi91++ {
                    var v92 bool = (false && (57 >= 2))
                    _ = fi91
                    _ = v92
                }
                switch (("a" + "\n\t") + "a") {
                case "a":
                    var v93 float32 = 3.14159
                    _ = v93
                case "":
                    v1 -= v1
                case "\n\t":
                    v2 = 2147483648
                default:
                    switch "a" {
                    case "hello":
                        v2 = 0
                    }
                }
                fmt.Println(v1, fi90)
                _ = fi90
            }
            var v94 int = v3
            _ = v94
        }
    } else {
        var v95 int32 = 42
        for fi96 := 0; fi96 < 9; fi96++ {
            var v97 int64 = int64(v2)
            for fi98 := 0; fi98 < 1; fi98++ {
                var v99 int64 = 9223372036854775807
                _ = fi98
                _ = v99
            }
            _ = fi96
            _ = v97
        }
        _ = v95
    }
    fmt.Println(v2, v2)
    v100 := []int{2147483647, v3}
    var v101 int64 = 869
    fmt.Println(v2)
    fmt.Println(v2)
    v100 = append(v100, int(v101))
    for fi102 := 0; fi102 < 3; fi102++ {
        var v103 uint32 = (v2 << ((4294967295) & 63))
        _ = fi102
        _ = v103
    }
    fmt.Println(v1, v3, v101)
    v104 := GenS4{F0: (1000 > 42), F1: (-2.25 <= float64(0.0)), F2: -194}
    var v105 uint8 = 227
    v106 := GenS3{F0: -580, F1: 0, F2: 326}
    v1 += v1
    for fi107 := 0; fi107 < 20; fi107++ {
        var v108 int = (v3 - fi107)
        if false {
            fmt.Println(v108, v3, v101)
        }
        _ = fi107
        _ = v108
    }
    _ = len(v100)
    v109 := GenS2{F0: 2}
    var v110 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v101:", v101)
    fmt.Println("v105:", v105)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v100
    _ = v101
    _ = v104
    _ = v105
    _ = v106
    _ = v109
    _ = v110
    _ = v2
    _ = v3
}
