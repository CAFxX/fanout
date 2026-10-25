package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 bool; F1 int32; F2 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint8; F1 int; F2 int64 }
type GenS12 struct { F0 uint16 }
func genFn0() uint32 {
    var v5 int8 = 1
    var v6 uint = 4294967295
    v7 := []int{128}
    var v8 uint64 = 2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 868
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 32767
    var v3 uint32 = 1
    var v4 string = "a"
    if true {
        v4 = "a"
        var v9 uint64 = 500
        v10 := GenS1{F0: -257}
        _ = v10
        _ = v9
    } else {
        var v11 uint = 253
        var v12 int = 1
        _ = v11
        _ = v12
    }
    v2 = -32768
    v13 := []int{128}
    fmt.Println(v2)
    switch -540 {
    case 487:
        var v14 int8 = -128
        _ = v14
    case 950:
        v15 := GenS1{F0: -161}
        _ = v15
    case 113:
        v4 = (((((v4 + v4) + (v4 + "")) + v4) + v4) + v4)
    }
    v2 -= 32767
    v4 = v4
    if false {
        var v16 string = (v4 + v4)
        _ = v16
    } else {
        v4 = v4
        v4 = v4
        v17 := GenS2{F0: -1, F1: 127}
        _ = v17
    }
    var v18 bool = (!(!true))
    fmt.Println(v4)
    for fi19 := 0; fi19 < 17; fi19++ {
        switch "" {
        case "hello":
            if (0 != 156) {
                var v20 float64 = (1.5 - 1.5)
                _ = v20
            } else {
                switch ("\n\t" + v4) {
                case "\n\t":
                    v21 := GenS1{F0: 1}
                    _ = v21
                default:
                    if v18 {
                        for fi22 := 0; fi22 < 19; fi22++ {
                            v23 := GenS3{F0: (v4 + v4), F1: (v18 && (165 < 0)), F2: 2}
                            v4 = "hello"
                            _ = fi22
                            _ = v23
                        }
                        var v24 float64 = 0.1
                        switch "a" {
                        case "":
                            var v25 uint = 18446744073709551615
                            _ = v25
                        }
                        _ = v24
                    }
                }
            }
        case "foo bar":
            if (!v18) {
                v2 = (^(v2))
            } else {
                var v26 bool = true
                v18 = (42 != -122)
                _ = v26
            }
        case "a":
            if (!v18) {
                v4 = ("\n\t" + v4)
                v4 = v4
                v27 := GenS3{F0: ((v4 + (v4 + "hello")) + v4), F1: v18, F2: 2}
                _ = v27
            } else {
                if (45 <= 1) {
                    var v28 int32 = -2147483648
                    fmt.Println(v2)
                    _ = len(v13)
                    _ = v28
                } else {
                    fmt.Println(v4)
                    var v29 int = -9223372036854775808
                    _ = v29
                }
                v30 := GenS2{F0: 42, F1: 206}
                _ = v30
            }
        default:
            switch fi19 {
            case -9223372036854775808:
                if v18 {
                    v31 := GenS4{F0: (^(v3)), F1: 42}
                    var v32 int8 = -30
                    v32 -= v32
                    _ = v31
                    _ = v32
                } else {
                    var v33 uint32 = v3
                    v4 = v4
                    switch "\n\t" {
                    case "a":
                        v34 := GenS5{F0: (-1 > 1), F1: 0, F2: (v4 + v4)}
                        _ = v34
                    case "hello":
                        var v35 int16 = v2
                        _ = v35
                    case "":
                        var v36 int8 = -128
                        _ = v36
                    }
                    _ = v33
                }
            case 42:
                for fi37 := 0; fi37 < 4; fi37++ {
                    switch (fi37 / ((fi19) | 1)) {
                    case 669:
                        v2 += v2
                        var v38 uint8 = 1
                        v39 := GenS6{F0: fi37}
                        _ = v38
                        _ = v39
                    default:
                        var v40 int64 = -9223372036854775808
                        _ = v40
                    }
                    _ = fi37
                }
            case 518:
                v18 = (!(1 < 1))
            default:
                _ = len(v13)
            }
        }
        if (2147483647 < 42) {
            var v41 uint64 = 414
            var v42 uint8 = 255
            var v43 uint64 = (v41 & 18446744073709551615)
            _ = v41
            _ = v42
            _ = v43
        } else {
            var v44 int8 = 127
            for fi45 := 0; fi45 < 2; fi45++ {
                v2 = int16(v3)
                v18 = (!(!false))
                if (0 == 100) {
                    v13 = append(v13, fi19)
                }
                v46 := GenS7{F0: -960}
                switch (-(fi19)) {
                case -2147483648:
                    v47 := GenS8{F0: 1, F1: v2}
                    _ = v47
                case 32767:
                    switch ("a" + (v4 + v4)) {
                    case "hello":
                        v48 := GenS3{F0: v4, F1: (42 != 64), F2: 18446744073709551615}
                        _ = v48
                    case "\n\t":
                        if true {
                            v49 := GenS6{F0: 0}
                            v13 = append(v13, (fi45 % ((fi19) | 1)))
                            _ = v49
                        } else {
                            var v50 uint64 = 0
                            var v51 uint8 = 0
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            _ = v50
                            _ = v51
                        }
                    }
                }
                switch int(v2) {
                case 32768:
                    fmt.Println(v3, v18, fi45)
                    var v52 int64 = 0
                    if len(v13) > 0 {
                        v13[0] = (fi19 | fi45)
                    }
                    var v53 bool = (!v18)
                    _ = v52
                    _ = v53
                case 9223372036854775807:
                    var v54 uint = 0
                    _ = v54
                default:
                    delete(v1, (v3 | v3))
                }
                _ = fi45
                _ = v46
            }
            v4 = (("\n\t" + ("\n\t" + "")) + ((v4 + ("" + v4)) + v4))
            _ = v44
        }
        _ = fi19
    }
    if len(v13) > 0 {
        v13[0] = 458
    }
    if (false || (-520 > 42)) {
        var v55 int = 25
        v56 := GenS9{F0: 128, F1: 127}
        _ = v55
        _ = v56
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v57 uint16 = 683
    for fi58 := 0; fi58 < 7; fi58++ {
        switch fi58 {
        case -1:
            fmt.Println(v4)
        }
        if false {
            v59 := GenS8{F0: 57, F1: v2}
            _ = v59
        } else {
            if (0.1 == 1e10) {
                v2 = v2
            }
            if (!(v57 <= (v57 ^ v57))) {
                var v60 uint64 = 42
                _ = v60
            } else {
                var v61 string = "a"
                if v18 {
                    v62 := GenS7{F0: -916}
                    var v63 int32 = 0
                    _ = v62
                    _ = v63
                }
                switch fi58 {
                case -298:
                    v64 := GenS10{F0: "a"}
                    _ = v64
                case 881:
                    v65 := GenS5{F0: (true && v18), F1: 2147483647, F2: v61}
                    _ = v65
                default:
                    if len(v13) > 0 {
                        v13[0] = fi58
                    }
                }
                _ = v61
            }
        }
        switch (v4 + "") {
        case "a":
            v18 = (!(1 != 100))
        case "\n\t":
            for fi66 := 0; fi66 < 18; fi66++ {
                if false {
                    v67 := GenS3{F0: v4, F1: false, F2: 42}
                    v4 = v4
                    v18 = (v18 || (88 >= 42))
                    _ = v67
                } else {
                    if v18 {
                        v68 := GenS11{F0: 2, F1: 1, F2: 21}
                        _ = v68
                    } else {
                        if false {
                            switch (fi66 >> ((18446744073709551615) & 63)) {
                            case -720:
                                if len(v13) > 0 {
                                    v13[0] = fi58
                                }
                            case 754:
                                var v69 bool = (!true)
                                _ = v69
                            }
                            v70 := GenS12{F0: v57}
                            _ = v70
                        }
                        v71 := GenS2{F0: int64(fi58), F1: 210}
                        _ = v71
                    }
                }
                switch fi66 {
                case 1:
                    var v72 uint64 = 4294967295
                    _ = v72
                case 127:
                    if (int(v3) != fi58) {
                        fmt.Println(v2, v18)
                    }
                }
                v1[v3] = (!(1 == 108))
                _ = fi66
            }
        case "hello":
            for fi73 := 0; fi73 < 14; fi73++ {
                fmt.Println(v4, v3, v18)
                var v74 uint32 = (v3 % (((^(v3))) | 1))
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                v3 += v3
                switch v4 {
                case "hello":
                    v57 = 32767
                default:
                    if (((3.14159 - -2.25) / (1e10 / 1e10)) > 1.5) {
                        switch fi58 {
                        case 128:
                            v1[v74] = (!(0 == -859))
                        case 1:
                            var v75 uint64 = 389
                            _ = v75
                        }
                        _ = len(v1)
                        v18 = ((!(!v18)) || v18)
                        if (int64(v3) >= 9223372036854775807) {
                            v2 = v2
                            if (!true) {
                                fmt.Println(v18, v4, v18)
                                if v18 {
                                    var v76 int8 = 26
                                    _ = v76
                                } else {
                                    v57 += 315
                                }
                                v3 -= (v74 - 967)
                            } else {
                                var v77 int64 = int64(v2)
                                if v18 {
                                    if v18 {
                                        if len(v13) > 0 {
                                            v13[0] = 314
                                        }
                                        switch (fi73 + (fi58 + (fi73 << ((895) & 63)))) {
                                        case -797:
                                            fmt.Println(v74)
                                        default:
                                            var v78 int32 = 2147483647
                                            _ = v78
                                        }
                                        var v79 uint8 = 232
                                        _ = v79
                                    }
                                    if (79 != 8) {
                                        v74 = v74
                                        v80 := GenS7{F0: v77}
                                        _ = v80
                                    }
                                    v2 -= v2
                                } else {
                                    fmt.Println(v4)
                                }
                                var v81 uint32 = v3
                                _ = v77
                                _ = v81
                            }
                        }
                        v18 = ((1e10 >= 3.14159) && true)
                    }
                }
                _ = fi73
                _ = v74
            }
        default:
            v82 := GenS8{F0: 1, F1: (v2 - v2)}
            _ = v82
        }
        _ = fi58
    }
    v83 := GenS8{F0: 42, F1: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v83
}
