package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 int32; F2 int8 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 string; F1 uint64; F2 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint32; F1 uint }
type GenS10 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint16 = 137
    var v2 int16 = 1000
    var v3 string = "a"
    var v4 bool = (42 <= 4294967295)
    v5 := make(map[string]bool)
    fmt.Println(v2, v4, v3)
    var v6 int32 = -211
    v6 = (v6 << ((18446744073709551615) & 63))
    if (-1 == 100) {
        v2 -= (^(v2))
        v1 = 32767
    }
    v7 := []int{}
    switch 121 {
    case 42:
        var v8 uint16 = 1
        _ = v8
    case 32768:
        if (4294967295 <= 0) {
            v3 = v3
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: 1, F1: 254, F2: (v6 - v6)}
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 bool = v4
                    switch (fi9 - fi11) {
                    case -2147483648:
                        fmt.Println(v12)
                        v2 -= v2
                        var v13 uint32 = 0
                        fmt.Println(fi9, v3)
                        _ = v13
                    case 2147483647:
                        switch 2147483647 {
                        case -9223372036854775808:
                            v14 := v5[v3]
                            delete(v5, v3)
                            _ = v14
                        case 42:
                            var v15 float64 = (-2.25 / 1e10)
                            _ = v15
                        case 4294967295:
                            v5["hello"] = (true || v4)
                        }
                    }
                    switch v3 {
                    case "\n\t":
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    case "foo bar":
                        v3 = ("hello" + v3)
                    }
                    _ = fi11
                    _ = v12
                }
                v1 = (v1 / (((v1 + v1)) | 1))
                _ = fi9
                _ = v10
            }
        } else {
            v16 := GenS1{F0: 42, F1: 18446744073709551615, F2: (v6 ^ 2147483647)}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v17 int32 = v6
            _ = v16
            _ = v17
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if v4 {
        fmt.Println(v6, v2, v6)
        switch -55 {
        case 463:
            var v18 string = v3
            _ = v18
        case 47:
            for fi19 := 0; fi19 < 14; fi19++ {
                v20 := GenS2{F0: v2, F1: 127}
                if true {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                } else {
                    fmt.Println(v3, v2)
                    v21 := GenS1{F0: 42, F1: 42, F2: (-(v6))}
                    if (v2 <= (v2 << ((42) & 63))) {
                        if true {
                            fmt.Println(v3, v1, v1)
                        } else {
                            v4 = true
                            for fi22 := 0; fi22 < 18; fi22++ {
                                if v4 {
                                    var v23 int = (fi22 & fi22)
                                    if (!((42 != 0) && v4)) {
                                        v24 := GenS3{F0: v1}
                                        var v25 float32 = (float32(((1.5 + 0.1) + 1.5)) - (0.0 - 1.5))
                                        var v26 int16 = (v2 >> ((365) & 63))
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    } else {
                                        var v27 int8 = -1
                                        _ = v27
                                    }
                                    fmt.Println(v3, v23, fi22)
                                    _ = v23
                                } else {
                                    v1 = 0
                                    v6 -= 147
                                    v28 := GenS4{F0: (fi19 ^ (fi22 >> ((687) & 63)))}
                                    if (false && (!(v4 && v4))) {
                                        var v29 float64 = float64(1.5)
                                        v1 = (v1 / (((v1 | v1)) | 1))
                                        switch ("a" + v3) {
                                        case "hello":
                                            v5[("foo bar" + "\n\t")] = ((v6 >= v6) || v4)
                                        case "":
                                            v30 := GenS5{F0: 1, F1: 167, F2: 42}
                                            _ = v30
                                        }
                                        _ = v29
                                    }
                                    _ = v28
                                }
                                _ = fi22
                            }
                            v6 -= (v6 ^ v6)
                        }
                        if v4 {
                            v1 = (^(v1))
                        }
                    }
                    _ = v21
                }
                _ = fi19
                _ = v20
            }
        case -1:
            switch "\n\t" {
            case "hello":
                v3 = v3
            case "foo bar":
                v31 := GenS3{F0: (v1 * v1)}
                _ = v31
            case "a":
                v5["hello"] = (0 >= 392)
            }
        default:
            fmt.Println(v4, v1)
        }
        var v32 uint8 = 0
        _ = v32
    } else {
        v6 = (v6 << ((18446744073709551615) & 63))
        if true {
            for fi33 := 0; fi33 < 8; fi33++ {
                switch (("hello" + v3) + "\n\t") {
                case "":
                    v4 = v4
                }
                _ = fi33
            }
            switch 742 {
            case -520:
                var v34 int8 = -1
                _ = v34
            case 32767:
                switch 1 {
                case 128:
                    var v35 uint8 = 17
                    _ = v35
                case 105:
                    switch 772 {
                    case -18:
                        v36 := GenS6{F0: (v6 >> ((18446744073709551615) & 63)), F1: true}
                        _ = v36
                    case -1:
                        var v37 bool = ((!v4) && (-17 == -1))
                        _ = v37
                    }
                }
            case 0:
                if v4 {
                    var v38 int = 1
                    _ = v38
                } else {
                    v7 = append(v7, 32767)
                    var v39 string = "a"
                    switch 387 {
                    case 0:
                        switch 128 {
                        case 940:
                            v39 = v39
                        default:
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                    case 42:
                        v40 := GenS1{F0: 127, F1: 0, F2: (v6 | (^(v6)))}
                        _ = v40
                    case 1:
                        var v41 float64 = (1.5 / 1e10)
                        _ = v41
                    default:
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    _ = v39
                }
            }
            var v42 int = 42
            _ = v42
        }
    }
    if (v6 <= v6) {
        var v43 string = v3
        switch v3 {
        case "\n\t":
            v7 = append(v7, -618)
        case "foo bar":
            var v44 int16 = v2
            _ = v44
        }
        _ = v43
    }
    switch (v3 + v3) {
    case "\n\t":
        for fi45 := 0; fi45 < 6; fi45++ {
            if (!((v1 << ((4294967295) & 63)) <= 65535)) {
                for fi46 := 0; fi46 < 18; fi46++ {
                    if (275 >= 42) {
                        fmt.Println(fi45)
                        var v47 uint = 0
                        switch v3 {
                        case "foo bar":
                            v4 = (!true)
                        case "":
                            if v4 {
                                if len(v7) > 0 {
                                    v7[0] = fi46
                                }
                            } else {
                                v48 := GenS7{F0: "a", F1: 1, F2: v1}
                                var v49 int8 = -44
                                _ = v48
                                _ = v49
                            }
                        default:
                            var v50 uint32 = 418
                            _ = v50
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        fmt.Println(v1, fi45)
                        _ = v47
                    } else {
                        if (!(1.5 > 1e10)) {
                            v51 := GenS8{F0: -128}
                            _ = v51
                        }
                        v52 := GenS1{F0: 128, F1: 53, F2: (v6 << ((0) & 63))}
                        v4 = (265 != 0)
                        _ = v52
                    }
                    switch fi46 {
                    case 32767:
                        if true {
                            var v53 int64 = -723
                            var v54 uint16 = (v1 - 255)
                            _ = v53
                            _ = v54
                        }
                    case 1:
                        v55 := GenS4{F0: 2147483647}
                        _ = v55
                    }
                    _ = fi46
                }
                for fi56 := 0; fi56 < 11; fi56++ {
                    switch fi45 {
                    case 2147483647:
                        switch "\n\t" {
                        case "a":
                            switch fi45 {
                            case 4294967295:
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            default:
                                var v57 int64 = 1
                                _ = v57
                            }
                        }
                    case -95:
                        switch v3 {
                        case "a":
                            v1 = (v1 << ((255) & 63))
                        case "\n\t":
                            v5["hello"] = (!false)
                            if v4 {
                                var v58 uint8 = 128
                                v6 += (^(v6))
                                _ = len(v5)
                                _ = v58
                            } else {
                                var v59 uint32 = 904
                                v59 -= (v59 ^ v59)
                                if (!(!false)) {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    if v4 {
                                        delete(v5, (v3 + ("" + v3)))
                                        var v60 uint16 = (v1 & v1)
                                        var v61 int16 = v2
                                        _ = v60
                                        _ = v61
                                    }
                                    if false {
                                        switch v3 {
                                        case "\n\t":
                                            _ = len(v5)
                                        case "hello":
                                            var v62 uint8 = 86
                                            var v63 int64 = 42
                                            if (!false) {
                                                if (v6 == 42) {
                                                    if (0.1 > (1.5 * float64((0.0 - (0.0 - 0.1))))) {
                                                        var v64 int8 = -114
                                                        _ = v64
                                                    } else {
                                                        var v65 float32 = 0.1
                                                        fmt.Println(v1)
                                                        if ((!v4) && (((fi45 % (((fi56 >> ((688) & 63))) | 1)) != fi56) || (!false))) {
                                                            v66 := GenS9{F0: 0, F1: 65535}
                                                            _ = len(v5)
                                                            v6 += 1
                                                            var v67 int8 = 0
                                                            if v4 {
                                                                if len(v7) > 0 {
                                                                    v7[0] = (fi56 & 0)
                                                                }
                                                                v59 -= v59
                                                            } else {
                                                                if ("a" < "\n\t") {
                                                                    v62 = v62
                                                                } else {
                                                                    if (2 >= 42) {
                                                                        v68 := GenS1{F0: 106, F1: uint(v1), F2: (v6 + 42)}
                                                                        var v69 int64 = 42
                                                                        _ = v68
                                                                        _ = v69
                                                                    }
                                                                    v70 := GenS10{F0: 42, F1: (-(v6))}
                                                                    _ = v70
                                                                }
                                                            }
                                                            v6 = (v6 + v6)
                                                            if v4 {
                                                                v71 := GenS3{F0: v1}
                                                                _ = v71
                                                            } else {
                                                                var v72 uint32 = (v59 * v59)
                                                                v3 = ("" + (v3 + v3))
                                                                if (0 == 42) {
                                                                    v73 := GenS1{F0: (v62 + v62), F1: 1, F2: v6}
                                                                    v67 -= (v67 >> ((517) & 63))
                                                                    switch ("" + v3) {
                                                                    case "\n\t":
                                                                        var v74 uint8 = (v62 & (v62 << ((255) & 63)))
                                                                        _ = v74
                                                                    default:
                                                                        _ = len(v5)
                                                                    }
                                                                    _ = v73
                                                                } else {
                                                                    var v75 uint8 = v62
                                                                    _ = v75
                                                                }
                                                                _ = v72
                                                            }
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        _ = v65
                                                    }
                                                }
                                            } else {
                                                if (v63 <= v63) {
                                                    fmt.Println(fi56, fi45)
                                                    v3 = "a"
                                                } else {
                                                    v62 -= v62
                                                }
                                                var v76 float64 = -2.25
                                                _ = v76
                                            }
                                            _ = v62
                                            _ = v63
                                        case "a":
                                            fmt.Println(v4)
                                        }
                                    }
                                    v59 = (^(v59))
                                } else {
                                    switch "\n\t" {
                                    case "a":
                                        var v77 int8 = 56
                                        _ = v77
                                    case "\n\t":
                                        var v78 int8 = 100
                                        _ = v78
                                    case "foo bar":
                                        _ = len(v7)
                                    }
                                    v59 -= (v59 >> ((uint(v1)) & 63))
                                    var v79 int = (fi45 - fi56)
                                    _ = v79
                                }
                                _ = v59
                            }
                            var v80 int32 = (v6 ^ v6)
                            _ = v80
                        case "":
                            _ = len(v7)
                        default:
                            v1 = v1
                        }
                    }
                    var v81 int8 = 42
                    _ = fi56
                    _ = v81
                }
            }
            _ = fi45
        }
    default:
        v6 = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
