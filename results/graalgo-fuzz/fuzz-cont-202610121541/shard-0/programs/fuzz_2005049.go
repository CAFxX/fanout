package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 bool; F2 int }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 bool; F1 uint; F2 int }
type GenS6 struct { F0 bool; F1 uint; F2 int }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 string; F1 int64; F2 uint16 }
func genFn0(p0 uint16, p1 int32, p2 int64) uint {
    v6 := GenS1{F0: -2147483648}
    var v7 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint16 = 886
    var v2 uint8 = 189
    var v3 int64 = -914
    var v4 int8 = 100
    var v5 int64 = v3
    var v8 string = "\n\t"
    var v9 int16 = -338
    v5 = 669
    switch -887 {
    case 128:
        var v10 uint8 = v2
        _ = v10
    case 0:
        v11 := GenS1{F0: 1}
        _ = v11
    default:
        fmt.Println(v9, v3, v4)
    }
    var v12 int64 = 42
    for fi13 := 0; fi13 < 20; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            if (v9 > v9) {
                if ((v2 + (v2 << ((320) & 63))) > (v2 & v2)) {
                    v4 = v4
                    if (v9 > v9) {
                        v2 = v2
                        var v15 float32 = 0.1
                        v16 := GenS2{F0: 1, F1: false, F2: (fi13 & (fi13 >> ((255) & 63)))}
                        _ = v15
                        _ = v16
                    }
                    v1 = v1
                } else {
                    var v17 uint64 = 18446744073709551615
                    v4 = (v4 % (((v4 >> ((4294967295) & 63))) | 1))
                    fmt.Println(v9)
                    _ = v17
                }
                var v18 uint16 = uint16(v2)
                var v19 uint16 = v1
                _ = v18
                _ = v19
            } else {
                fmt.Println(v4, v3)
                v20 := []int{-107, 127}
                var v21 bool = false
                _ = v20
                _ = v21
            }
            v22 := []int{1, fi13, (fi13 % ((fi13) | 1)), fi13}
            v23 := GenS3{F0: (v2 % ((255) | 1)), F1: (v3 - 429)}
            _ = fi14
            _ = v22
            _ = v23
        }
        switch "\n\t" {
        case "\n\t":
            v24 := GenS3{F0: (^(v2)), F1: -944}
            _ = v24
        }
        _ = fi13
    }
    var v25 string = v8
    v1 += (v1 / ((65535) | 1))
    if ((0.1 + ((-2.25 + 0.0) + float64(0.1))) < 0.1) {
        v12 -= v5
        v8 = "hello"
    } else {
        v9 = (v9 - (v9 % ((v9) | 1)))
        v5 = (^(v3))
        if true {
            var v26 int64 = (v5 - (v5 & (-(v3))))
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 uint = 2
        var v29 bool = false
        if (!true) {
            var v30 int = fi27
            _ = v30
        } else {
            for fi31 := 0; fi31 < 13; fi31++ {
                fmt.Println(v25, v3, v28)
                if (!(1.5 >= (1e10 / 1.5))) {
                    fmt.Println(v8, v25, fi31)
                    var v32 string = ""
                    v32 = v32
                    _ = v32
                } else {
                    switch (fi27 + fi31) {
                    case -865:
                        var v33 int = fi31
                        var v34 uint32 = 0
                        if v29 {
                            v12 = v3
                        }
                        if (32767 > v1) {
                            switch v8 {
                            case "foo bar":
                                v28 -= v28
                            case "a":
                                v35 := GenS4{F0: fi31, F1: (v9 + v9)}
                                _ = v35
                            case "hello":
                                v33 -= 948
                            default:
                                var v36 int = (^(v33))
                                _ = v36
                            }
                            switch (fi31 & (fi31 / (((fi31 | 768)) | 1))) {
                            case -9223372036854775808:
                                v37 := []int{v33, (fi31 << ((uint(v34)) & 63))}
                                _ = v37
                            }
                        }
                        switch (fi31 * 32767) {
                        case 9223372036854775807:
                            switch fi27 {
                            case 2147483647:
                                var v38 float64 = float64(0.1)
                                _ = v38
                            case 127:
                                var v39 float64 = 0.1
                                _ = v39
                            case -2147483648:
                                fmt.Println(v25)
                            }
                        case -820:
                            var v40 uint64 = 1
                            _ = v40
                        }
                        v29 = v29
                        _ = v33
                        _ = v34
                    case 9223372036854775807:
                        if false {
                            v41 := GenS5{F0: (1 <= 42), F1: 65535, F2: (fi27 / (((fi31 * (fi27 + fi27))) | 1))}
                            var v42 uint32 = 519
                            v42 = 0
                            if v29 {
                                var v43 int32 = 0
                                v1 += 42
                                fmt.Println(v8, v12)
                                _ = v43
                            }
                            _ = v41
                            _ = v42
                        } else {
                            var v44 int16 = 32767
                            v45 := GenS4{F0: fi27, F1: v9}
                            v8 = v8
                            var v46 float64 = 3.14159
                            var v47 float64 = 1e10
                            _ = v44
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                    case 0:
                        v48 := []int{(fi31 * fi27), fi27, (fi27 / ((fi31) | 1)), (fi27 - (fi31 << (((v28 + 255)) & 63)))}
                        _ = v48
                    default:
                        v49 := GenS6{F0: false, F1: v28, F2: fi31}
                        _ = v49
                    }
                    var v50 float32 = 0.0
                    _ = v50
                }
                _ = fi31
            }
            var v51 uint = genFn0((v1 % (((v1 | 32767)) | 1)), -2147483648, v5)
            _ = v51
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    var v52 float64 = 1.5
    v25 = "hello"
    for fi53 := 0; fi53 < 8; fi53++ {
        v4 += 79
        if (true && true) {
            v54 := GenS7{F0: 65535, F1: 42}
            for fi55 := 0; fi55 < 10; fi55++ {
                v56 := GenS8{F0: 22}
                v57 := []int{(^(fi53)), fi53, fi53, -2147483648}
                _ = fi55
                _ = v56
                _ = v57
            }
            switch ("a" + v8) {
            case "foo bar":
                var v58 int = (fi53 + (-(fi53)))
                _ = v58
            }
            _ = v54
        } else {
            var v59 uint = 4294967295
            for fi60 := 0; fi60 < 9; fi60++ {
                var v61 int16 = (v9 % ((-268) | 1))
                var v62 float64 = (v52 + float64(3.14159))
                _ = fi60
                _ = v61
                _ = v62
            }
            _ = v59
        }
        if (2147483647 == -2147483648) {
            v63 := GenS4{F0: (fi53 - -2147483648), F1: (v9 & (v9 << ((255) & 63)))}
            _ = v63
        }
        _ = fi53
    }
    if true {
        switch "" {
        case "a":
            v8 = ("hello" + v8)
        default:
            for fi64 := 0; fi64 < 15; fi64++ {
                for fi65 := 0; fi65 < 20; fi65++ {
                    if (-2147483648 < -365) {
                        var v66 int = fi65
                        v2 += 0
                        var v67 int = -9223372036854775808
                        v52 += float64((1.5 + float32((v52 * 0.1))))
                        v2 = 2
                        _ = v66
                        _ = v67
                    }
                    if false {
                        v68 := GenS9{F0: (v8 + ((v8 + v8) + v25)), F1: v3, F2: v1}
                        var v69 float32 = (0.0 / 1.0)
                        _ = v68
                        _ = v69
                    }
                    if true {
                        v2 -= 154
                        if true {
                            v70 := GenS3{F0: (v2 ^ (v2 >> ((65535) & 63))), F1: v5}
                            v9 += 42
                            var v71 int16 = (v9 / ((v9) | 1))
                            _ = v70
                            _ = v71
                        }
                        v1 -= (v1 / ((v1) | 1))
                    } else {
                        if ((v1 + v1) != (^(v1))) {
                            if (true && true) {
                                v9 = 42
                                v8 = "a"
                            }
                            if true {
                                v25 = v8
                                fmt.Println(v52, v3, v52)
                                if true {
                                    if true {
                                        var v72 int64 = v3
                                        _ = v72
                                    } else {
                                        var v73 int = fi64
                                        _ = v73
                                    }
                                    var v74 string = "a"
                                    v75 := GenS7{F0: 42, F1: 1}
                                    _ = v74
                                    _ = v75
                                }
                            }
                            var v76 int16 = (v9 ^ 367)
                            v52 -= (v52 + v52)
                            _ = v76
                        } else {
                            var v77 int = -1
                            var v78 uint32 = 2147483647
                            if true {
                                if false {
                                    fmt.Println(v1)
                                    v4 += (v4 / ((42) | 1))
                                    var v79 int = 456
                                    v1 = (v1 / (((^(v1))) | 1))
                                    _ = v79
                                }
                                var v80 int32 = -428
                                var v81 int16 = 1000
                                v82 := []int{}
                                var v83 int32 = (v80 | int32(v4))
                                _ = v80
                                _ = v81
                                _ = v82
                                _ = v83
                            } else {
                                v84 := []int{(fi65 % ((v77) | 1)), fi65, fi65, fi65}
                                _ = v84
                            }
                            _ = v77
                            _ = v78
                        }
                        var v85 uint16 = v1
                        _ = v85
                    }
                    _ = fi65
                }
                _ = fi64
            }
        }
        if false {
            fmt.Println(v3, v3)
            var v86 uint = 18446744073709551615
            if (!(!(!(v86 >= genFn0(v1, 2147483647, 9223372036854775807))))) {
                switch v8 {
                case "\n\t":
                    for fi87 := 0; fi87 < 15; fi87++ {
                        v88 := GenS7{F0: 0, F1: 966}
                        v89 := []int{fi87, fi87, (fi87 << (((v86 >> ((genFn0(v1, -100, (v5 - (v12 / ((v5) | 1))))) & 63))) & 63)), 966}
                        v90 := GenS3{F0: 179, F1: 0}
                        _ = fi87
                        _ = v88
                        _ = v89
                        _ = v90
                    }
                default:
                    if false {
                        var v91 string = "foo bar"
                        _ = v91
                    }
                }
            }
            _ = v86
        } else {
            v92 := GenS9{F0: ("\n\t" + (v25 + "")), F1: (-(v5)), F2: v1}
            v4 -= (v4 % (((v4 ^ (-(v4)))) | 1))
            _ = v92
        }
        fmt.Println(v4, v52)
    }
    v93 := []int{-2147483648, 9223372036854775807}
    if (65535 >= 298) {
        v4 = (v4 << ((255) & 63))
        v2 += v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    fmt.Println("slicelen v93:", len(v93), v93)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v8
    _ = v9
    _ = v93
}
