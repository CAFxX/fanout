package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int32; F1 uint; F2 uint8 }
func genFn0(p0 uint64, p1 uint8) uint16 {
    v5 := []int{-2147483648, 9223372036854775807, -618, 32767}
    var v6 int32 = 42
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 65535
}
func genFn1() int64 {
    switch 9223372036854775807 {
    case 42:
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 uint8 = 1
                _ = fi7
                _ = v8
            }
            if true {
                var v9 uint16 = 647
                for fi10 := 0; fi10 < 9; fi10++ {
                    if (("a" <= "foo bar") || true) {
                        var v11 string = "a"
                        var v12 int16 = 636
                        v13 := []int{}
                        _ = v11
                        _ = v12
                        _ = v13
                    } else {
                        var v14 bool = false
                        v14 = (v14 || true)
                        for fi15 := 0; fi15 < 12; fi15++ {
                            v16 := []int{fi10, fi15, 2147483647}
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                    fmt.Println(v9)
                    fmt.Println(fi10)
                    _ = fi10
                }
                _ = v9
            }
        }
    case -9223372036854775808:
        var v17 float64 = 1.5
        _ = v17
    }
    if (!true) {
        var v18 float64 = 1.5
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 uint16 = genFn0(432, 128)
            if true {
                var v21 int16 = -32768
                v20 -= (v20 << ((42) & 63))
                var v22 bool = (true || true)
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 2; fi23++ {
                    v24 := GenS1{F0: 42, F1: 18446744073709551615, F2: -572}
                    var v25 int = fi19
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            _ = fi19
            _ = v20
        }
        _ = v18
    } else {
        var v26 string = "\n\t"
        _ = v26
    }
    var v27 uint64 = 18446744073709551615
    var v28 uint = 2
    var v29 int16 = 815
    _ = v27
    _ = v28
    _ = v29
    return -1
}

func main() {
    v1 := []int{9223372036854775807, 127}
    var v2 string = "foo bar"
    v3 := []int{}
    v4 := make(map[uint32]bool)
    v2 = (v2 + v2)
    v30 := GenS2{F0: 778}
    v2 = ("hello" + (v2 + v2))
    var v31 int = 0
    for fi32 := 0; fi32 < 16; fi32++ {
        fmt.Println(v2, fi32, v31)
        _ = fi32
    }
    v31 = 1
    switch v31 {
    case 127:
        var v33 int8 = 100
        _ = v33
    case 0:
        fmt.Println(v2)
    }
    if (0.0 == (((1.5 + ((1.5 / 1.0) + 0.1)) * (((3.14159 - 0.1) - 0.1) - 1e10)) + float64(0.1))) {
        for fi34 := 0; fi34 < 2; fi34++ {
            if true {
                if ((false || false) || false) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    if (3.14159 >= (1e10 * 1.5)) {
                        for fi35 := 0; fi35 < 16; fi35++ {
                            v2 = v2
                            _ = fi35
                        }
                        v31 += fi34
                    } else {
                        var v36 int8 = 0
                        _ = v36
                    }
                    if false {
                        v31 = 723
                        var v37 int16 = -1
                        _ = v37
                    } else {
                        _ = len(v4)
                    }
                }
                if len(v3) > 0 {
                    v3[0] = (v31 >> ((65535) & 63))
                }
            } else {
                v38 := GenS3{F0: 42}
                _ = v38
            }
            var v39 int8 = 100
            _ = fi34
            _ = v39
        }
    }
    v2 = v2
    var v40 int16 = 32767
    fmt.Println(v2)
    switch "\n\t" {
    case "hello":
        if ((v31 == (v31 + (v31 >> ((60) & 63)))) && true) {
            v41 := GenS2{F0: v40}
            _ = v41
        } else {
            switch v2 {
            case "a":
                v40 -= (v40 ^ v40)
            case "":
                if true {
                    switch ("\n\t" + ("foo bar" + ((("\n\t" + v2) + (v2 + v2)) + v2))) {
                    case "foo bar":
                        v42 := GenS1{F0: 1, F1: 42, F2: v40}
                        _ = v42
                    case "\n\t":
                        for fi43 := 0; fi43 < 12; fi43++ {
                            v44 := GenS1{F0: 42, F1: 2, F2: -1}
                            fmt.Println(v40, v2, v40)
                            for fi45 := 0; fi45 < 4; fi45++ {
                                if (!(((-2.25 - ((float64(-2.25) * (((1e10 + 0.0) - 0.0) + 1.5)) - -2.25)) - 0.1) == 1e10)) {
                                    v46 := GenS3{F0: 32767}
                                    if (!(!(-1 >= (fi43 % (((fi45 * 42)) | 1))))) {
                                        var v47 uint16 = 757
                                        var v48 bool = ((v31 | fi45) >= (fi43 ^ 0))
                                        _ = v47
                                        _ = v48
                                    } else {
                                        switch -185 {
                                        case 127:
                                            var v49 int16 = (v40 << ((4294967295) & 63))
                                            _ = v49
                                        }
                                        v1 = append(v1, 448)
                                        fmt.Println(fi43, v31, v31)
                                        var v50 uint64 = 0
                                        var v51 uint16 = 973
                                        v50 -= (v50 & v50)
                                        fmt.Println(fi43, v31, v40)
                                        var v52 float64 = 1.5
                                        v4[2147483648] = ((fi43 / (((fi45 / (((v31 | fi45)) | 1))) | 1)) >= int(v51))
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    }
                                    v31 = (fi43 | fi45)
                                    _ = v46
                                } else {
                                    v2 = ((v2 + v2) + "")
                                }
                                var v53 bool = true
                                _ = fi45
                                _ = v53
                            }
                            _ = fi43
                            _ = v44
                        }
                    default:
                        v31 = v31
                    }
                } else {
                    switch v2 {
                    case "\n\t":
                        if true {
                            if (!(0 >= 0)) {
                                var v54 uint64 = 42
                                for fi55 := 0; fi55 < 9; fi55++ {
                                    for fi56 := 0; fi56 < 1; fi56++ {
                                        v57 := GenS2{F0: 0}
                                        _ = fi56
                                        _ = v57
                                    }
                                    if true {
                                        v58 := GenS2{F0: v40}
                                        _ = v58
                                    }
                                    _ = fi55
                                }
                                v59 := GenS3{F0: 32767}
                                _ = v54
                                _ = v59
                            } else {
                                for fi60 := 0; fi60 < 2; fi60++ {
                                    if true {
                                        switch v2 {
                                        case "\n\t":
                                            switch "a" {
                                            case "\n\t":
                                                if (false && (genFn0(42, 128) >= 65535)) {
                                                    var v61 string = ("" + v2)
                                                    var v62 int = v31
                                                    _ = v61
                                                    _ = v62
                                                }
                                            case "foo bar":
                                                var v63 int = v31
                                                _ = v63
                                            case "":
                                                _ = len(v3)
                                            default:
                                                for fi64 := 0; fi64 < 12; fi64++ {
                                                    if (false || true) {
                                                        var v65 uint = 1
                                                        var v66 int64 = genFn1()
                                                        delete(v4, 1)
                                                        _ = v65
                                                        _ = v66
                                                    } else {
                                                        v31 = fi60
                                                        v67 := GenS4{F0: v40, F1: 1, F2: 1}
                                                        _ = v67
                                                    }
                                                    if (v40 < (v40 / ((-32768) | 1))) {
                                                        if (65535 != 4294967295) {
                                                            v2 = ("\n\t" + "foo bar")
                                                        }
                                                        fmt.Println(v40)
                                                    }
                                                    _ = fi64
                                                }
                                            }
                                        default:
                                            switch (^(v31)) {
                                            case 32767:
                                                fmt.Println(fi60)
                                            case -2147483648:
                                                var v68 int = fi60
                                                _ = v68
                                            case 1:
                                                _ = len(v1)
                                            default:
                                                v40 = 887
                                            }
                                        }
                                        var v69 float64 = ((3.14159 * float64(-2.25)) / ((1.5 - 3.14159) - (-2.25 + 3.14159)))
                                        _ = v69
                                    }
                                    _ = fi60
                                }
                                for fi70 := 0; fi70 < 17; fi70++ {
                                    var v71 uint32 = 1
                                    if true {
                                        switch "a" {
                                        case "foo bar":
                                            fmt.Println(v31, v31, fi70)
                                        case "a":
                                            fmt.Println(fi70, v31, fi70)
                                        }
                                    } else {
                                        if (((v2 + "\n\t") + ("foo bar" + (v2 + ((v2 + v2) + v2)))) < "") {
                                            for fi72 := 0; fi72 < 6; fi72++ {
                                                v73 := GenS1{F0: -9223372036854775808, F1: 1, F2: 1000}
                                                _ = fi72
                                                _ = v73
                                            }
                                            for fi74 := 0; fi74 < 3; fi74++ {
                                                fmt.Println(fi74, v2)
                                                _ = len(v1)
                                                v40 = (v40 + v40)
                                                if true {
                                                    v71 -= (v71 * (^(v71)))
                                                    v40 -= v40
                                                    var v75 int8 = -1
                                                    _ = v75
                                                } else {
                                                    var v76 float64 = 3.14159
                                                    switch v31 {
                                                    case -323:
                                                        fmt.Println(fi74, v2, v40)
                                                    case 879:
                                                        v71 += v71
                                                    case 127:
                                                        if len(v1) > 0 {
                                                            v1[0] = (v31 >> ((42) & 63))
                                                        }
                                                    default:
                                                        var v77 int8 = -63
                                                        switch v2 {
                                                        case "a":
                                                            fmt.Println(v2, v76, v2)
                                                        case "":
                                                            v76 = v76
                                                            v78 := v4[4294967295]
                                                            _ = v78
                                                        }
                                                        var v79 int16 = (v40 >> ((0) & 63))
                                                        _ = v77
                                                        _ = v79
                                                    }
                                                    var v80 uint = 444
                                                    _ = v76
                                                    _ = v80
                                                }
                                                v81 := GenS4{F0: (v40 >> ((65535) & 63)), F1: v71, F2: v71}
                                                _ = fi74
                                                _ = v81
                                            }
                                        }
                                    }
                                    _ = fi70
                                    _ = v71
                                }
                                v82 := GenS5{F0: v2, F1: -114}
                                _ = v82
                            }
                        }
                    case "hello":
                        v31 = (v31 + (v31 ^ (v31 - (v31 * -187))))
                    case "":
                        for fi83 := 0; fi83 < 13; fi83++ {
                            _ = len(v1)
                            v40 -= (v40 - -32768)
                            var v84 float32 = 0.0
                            _ = fi83
                            _ = v84
                        }
                    default:
                        v85 := GenS6{F0: 42, F1: 18446744073709551615, F2: 190}
                        _ = v85
                    }
                    v86 := GenS1{F0: -1, F1: 2, F2: v40}
                    _ = v86
                }
            }
            if len(v1) > 0 {
                v1[0] = 32767
            }
            if true {
                v31 -= v31
                _ = len(v4)
                v31 -= (v31 >> ((255) & 63))
            }
        }
    case "":
        v1 = append(v1, -2147483648)
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
}
