package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int16; F2 int64 }
type GenS6 struct { F0 int; F1 int32; F2 int16 }
type GenS7 struct { F0 int8; F1 int32 }
func genFn0(p0 int32, p1 int16) int32 {
    v4 := GenS1{F0: -9223372036854775808, F1: 1, F2: p1}
    v5 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p0 * -339)
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 968
    var v3 uint16 = 1
    v2 = v2
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{(fi6 ^ 1), 9223372036854775807}
        switch (^(fi6)) {
        case 0:
            var v8 uint = (v2 - 0)
            _ = v8
        case 1:
            var v9 bool = true
            _ = v9
        default:
            fmt.Println(v2, fi6, fi6)
        }
        _ = fi6
        _ = v7
    }
    v10 := GenS2{F0: 1}
    fmt.Println(v1, v3)
    fmt.Println(v1)
    var v11 int16 = -1
    switch "hello" {
    case "foo bar":
        v12 := []int{2147483647, 933, 32767, 263}
        _ = v12
    case "\n\t":
        if (-1 != -1) {
            switch "hello" {
            case "a":
                for fi13 := 0; fi13 < 6; fi13++ {
                    if true {
                        for fi14 := 0; fi14 < 18; fi14++ {
                            if true {
                                var v15 float64 = 3.14159
                                switch "" {
                                case "a":
                                    v15 = 1.5
                                default:
                                    var v16 float64 = v15
                                    fmt.Println(fi14, v3, v2)
                                    v11 = v11
                                    _ = v16
                                }
                                _ = v15
                            } else {
                                var v17 bool = (!false)
                                _ = v17
                            }
                            if (!false) {
                                if true {
                                    var v18 uint16 = v3
                                    v19 := []int{9223372036854775807}
                                    _ = v18
                                    _ = v19
                                } else {
                                    if ("a" == "\n\t") {
                                        if false {
                                            fmt.Println(fi14, v2)
                                            var v20 int8 = 42
                                            v2 = v2
                                            _ = v20
                                        } else {
                                            v21 := []int{fi14, (fi13 ^ (fi14 + (fi14 | fi13))), fi14}
                                            var v22 float64 = 1e10
                                            if len(v21) > 0 {
                                                v21[0] = (fi14 / ((fi13) | 1))
                                            }
                                            _ = v21
                                            _ = v22
                                        }
                                        v2 = v2
                                        v23 := []int{fi14, fi13}
                                        v24 := GenS3{F0: (v1 * (v1 >> ((99) & 63))), F1: 42}
                                        _ = v23
                                        _ = v24
                                    }
                                    var v25 uint32 = 6
                                    v26 := GenS4{F0: 0, F1: 2}
                                    _ = v25
                                    _ = v26
                                }
                                if (!false) {
                                    switch -9223372036854775808 {
                                    case -892:
                                        v27 := []int{(fi13 - fi13), fi14}
                                        _ = v27
                                    default:
                                        v28 := GenS3{F0: 992, F1: 127}
                                        _ = v28
                                    }
                                    v11 -= -1
                                } else {
                                    var v29 bool = true
                                    _ = v29
                                }
                            } else {
                                var v30 int32 = -1
                                fmt.Println(fi13, v30)
                                switch fi14 {
                                case 32767:
                                    v31 := []int{(fi13 << ((v2) & 63)), fi14}
                                    _ = v31
                                case -1:
                                    var v32 int64 = -9223372036854775808
                                    _ = v32
                                }
                                _ = v30
                            }
                            v33 := GenS5{F0: (v1 | v1), F1: v11, F2: -9223372036854775808}
                            fmt.Println(fi14)
                            v11 += (v11 ^ v11)
                            _ = fi14
                            _ = v33
                        }
                        if (-113 != -128) {
                            v34 := GenS4{F0: 0, F1: 159}
                            var v35 int = fi13
                            _ = v34
                            _ = v35
                        }
                    } else {
                        var v36 string = "foo bar"
                        for fi37 := 0; fi37 < 16; fi37++ {
                            switch fi37 {
                            case -318:
                                v38 := []int{fi13, int(v2), fi13}
                                _ = v38
                            case 4294967295:
                                if true {
                                    var v39 int16 = (v11 + (v11 / ((v11) | 1)))
                                    fmt.Println(v11, v36, v39)
                                    _ = v39
                                } else {
                                    var v40 int8 = 3
                                    if (true || (fi13 != (fi13 | (fi37 + fi13)))) {
                                        v11 += (v11 % (((v11 << (((^(v2))) & 63))) | 1))
                                    } else {
                                        v41 := []int{-2147483648, fi37, (fi13 / ((4294967295) | 1)), 127}
                                        fmt.Println(v40, v3, v40)
                                        _ = v41
                                    }
                                    v36 = v36
                                    _ = v40
                                }
                            default:
                                v42 := []int{32768}
                                _ = v42
                            }
                            _ = fi37
                        }
                        _ = v36
                    }
                    var v43 bool = ("a" == "a")
                    v44 := GenS6{F0: (fi13 % ((-2147483648) | 1)), F1: 2147483647, F2: -294}
                    _ = fi13
                    _ = v43
                    _ = v44
                }
            case "":
                v2 = (^(v2))
            }
            v45 := GenS4{F0: (^(v11)), F1: 195}
            v46 := GenS4{F0: v11, F1: 127}
            _ = v45
            _ = v46
        }
    case "hello":
        fmt.Println(v2, v1, v1)
    default:
        for fi47 := 0; fi47 < 12; fi47++ {
            if false {
                v3 -= 65535
            } else {
                if false {
                    var v48 int8 = 127
                    _ = v48
                }
                var v49 uint = (^(v2))
                _ = v49
            }
            v50 := []int{42, fi47}
            v3 -= (v3 & 2)
            _ = fi47
            _ = v50
        }
    }
    if ((false || (0 <= 1)) && true) {
        if (true && ((v2 / (((v2 + v2)) | 1)) < v2)) {
            fmt.Println(v2, v3, v11)
        } else {
            v51 := GenS2{F0: 987}
            v52 := GenS7{F0: 127, F1: -2147483648}
            _ = v51
            _ = v52
        }
        if true {
            if (887 < 2147483647) {
                if true {
                    if (!false) {
                        v3 = 0
                    } else {
                        var v53 uint16 = (v3 ^ (v3 - v3))
                        _ = v53
                    }
                } else {
                    var v54 int = -1
                    _ = v54
                }
                fmt.Println(v11)
            }
            for fi55 := 0; fi55 < 19; fi55++ {
                v2 = 0
                _ = fi55
            }
            for fi56 := 0; fi56 < 6; fi56++ {
                var v57 float32 = -2.25
                for fi58 := 0; fi58 < 3; fi58++ {
                    if true {
                        switch 312 {
                        case -400:
                            var v59 string = ("" + "foo bar")
                            _ = v59
                        default:
                            if (v2 <= (v2 | v2)) {
                                var v60 int = (fi58 >> ((v2) & 63))
                                _ = v60
                            }
                        }
                    }
                    _ = fi58
                }
                _ = fi56
                _ = v57
            }
        }
    }
    for fi61 := 0; fi61 < 14; fi61++ {
        switch ("\n\t" + ("hello" + "")) {
        case "foo bar":
            v1 += 42
        case "\n\t":
            var v62 uint64 = v1
            _ = v62
        default:
            for fi63 := 0; fi63 < 17; fi63++ {
                fmt.Println(v2)
                var v64 int64 = -568
                fmt.Println(fi61)
                _ = fi63
                _ = v64
            }
        }
        if (((42 <= genFn0(genFn0(genFn0(1, 904), -83), 1000)) || true) || (!true)) {
            fmt.Println(v3)
            fmt.Println(fi61)
        }
        _ = fi61
    }
    for fi65 := 0; fi65 < 12; fi65++ {
        var v66 bool = (int64(fi65) == 42)
        v11 -= v11
        if v66 {
            v67 := []int{(fi65 * 262), (fi65 + (-(fi65)))}
            switch "\n\t" {
            case "\n\t":
                if v66 {
                    v2 += v2
                    var v68 float32 = 3.14159
                    if len(v67) >= 2 {
                        v67 = v67[1:len(v67)]
                    }
                    _ = v68
                } else {
                    v3 += 501
                }
            case "foo bar":
                if len(v67) > 0 {
                    v67[0] = (fi65 % (((fi65 / ((fi65) | 1))) | 1))
                }
            }
            _ = v67
        }
        _ = fi65
        _ = v66
    }
    var v69 uint16 = (v3 & v3)
    v69 = 32767
    v69 = (v3 & v69)
    v69 = (v3 / ((1) | 1))
    for fi70 := 0; fi70 < 17; fi70++ {
        var v71 int8 = -1
        _ = fi70
        _ = v71
    }
    v72 := []int{-1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v72:", len(v72), v72)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v69
    _ = v72
}
