package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint16; F1 uint8; F2 bool }
type GenS7 struct { F0 uint32; F1 int; F2 uint32 }
type GenS8 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    if (false || true) {
        _ = len(v1)
        var v4 float64 = 1.5
        _ = v4
    }
    var v5 uint8 = 5
    var v6 bool = (((v5 == (v5 & v5)) || false) || false)
    var v7 uint = 120
    if v6 {
        for fi8 := 0; fi8 < 13; fi8++ {
            switch 799 {
            case 573:
                var v9 int8 = -1
                _ = v9
            case 32767:
                var v10 int16 = 1000
                _ = v10
            case 771:
                var v11 uint = v7
                _ = v11
            default:
                delete(v2, (fi8 + fi8))
            }
            if (1000 != 963) {
                if ((^(v5)) < 255) {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        fmt.Println(v6, v5, v6)
                        _ = fi12
                    }
                    fmt.Println(v5, v6)
                } else {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v14 := GenS1{F0: 42, F1: 2147483647}
                        if (2 <= 0) {
                            v7 -= (v7 ^ 812)
                            v15 := GenS2{F0: 2, F1: -32768}
                            var v16 int16 = 42
                            var v17 string = (((("a" + ("foo bar" + "foo bar")) + "hello") + "hello") + "hello")
                            _ = v15
                            _ = v16
                            _ = v17
                        } else {
                            fmt.Println(v7)
                        }
                        if v6 {
                            var v18 int32 = 165
                            var v19 float32 = 1.5
                            v20 := GenS2{F0: 1, F1: -1}
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        var v21 uint8 = (^(v5))
                        fmt.Println(v5, fi8, v7)
                        v1["hello"] = v6
                        _ = fi13
                        _ = v14
                        _ = v21
                    }
                    fmt.Println(v6, fi8, v5)
                }
            } else {
                for fi22 := 0; fi22 < 18; fi22++ {
                    v7 -= v7
                    v23 := GenS1{F0: 4294967295, F1: 1}
                    _ = fi22
                    _ = v23
                }
            }
            _ = fi8
        }
        for fi24 := 0; fi24 < 17; fi24++ {
            var v25 uint32 = 907
            for fi26 := 0; fi26 < 16; fi26++ {
                v27 := GenS1{F0: 65535, F1: uint32(v5)}
                _ = fi26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        v7 += v7
    }
    if v6 {
        for fi28 := 0; fi28 < 1; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 uint16 = 32767
                v6 = v6
                if (29 <= 979) {
                    if (255 == v30) {
                        switch "\n\t" {
                        case "\n\t":
                            var v31 uint32 = 58
                            _ = v31
                        case "a":
                            v7 = v7
                        case "foo bar":
                            if ("hello" < (("hello" + ("\n\t" + "a")) + "")) {
                                v32 := GenS1{F0: 42, F1: 2147483647}
                                _ = v32
                            } else {
                                if (1 != (v30 & v30)) {
                                    if (false || (314 <= 2147483648)) {
                                        v7 += (v7 / (((^(v7))) | 1))
                                        var v33 float64 = 1e10
                                        _ = v33
                                    } else {
                                        fmt.Println(fi29)
                                    }
                                } else {
                                    fmt.Println(v6, v30, fi28)
                                }
                                var v34 uint8 = (v5 | v5)
                                v30 = (v30 % (((v30 | (v30 << ((v7) & 63)))) | 1))
                                _ = v34
                            }
                        }
                        v6 = v6
                        v35 := GenS2{F0: (v30 << ((532) & 63)), F1: -1}
                        v1["hello"] = false
                        var v36 uint = 1
                        _ = v35
                        _ = v36
                    } else {
                        if v6 {
                            var v37 int = fi28
                            v38 := GenS1{F0: 226, F1: 0}
                            _ = v37
                            _ = v38
                        } else {
                            var v39 int64 = -9223372036854775808
                            v30 = 0
                            _ = v39
                        }
                    }
                    switch int(v7) {
                    case 32767:
                        var v40 int = 1
                        _ = v40
                    case 799:
                        v5 -= 99
                    case 9223372036854775807:
                        var v41 int32 = 0
                        _ = v41
                    }
                } else {
                    var v42 int32 = -2147483648
                    _ = len(v2)
                    v6 = (!((!(((-1 != 1) && true) && (!(-2 != v42)))) && (!((v30 > 2) && (("foo bar" + ("a" + "a")) <= "hello")))))
                    v7 += (v7 - v7)
                    _ = v42
                }
                if v6 {
                    if (-26 > -128) {
                        var v43 float32 = 3.14159
                        _ = v43
                    }
                    v44 := GenS1{F0: 65535, F1: 4294967295}
                    _ = v44
                } else {
                    switch "foo bar" {
                    case "hello":
                        if false {
                            v1[""] = v6
                            v7 = 4294967295
                        }
                    case "":
                        fmt.Println(v7, fi28)
                        if (0.0 != ((1.5 - 1.5) - 1e10)) {
                            var v45 bool = ((^(v7)) != (v7 + 1))
                            _ = v45
                        } else {
                            v5 = (v5 << (((v7 << ((255) & 63))) & 63))
                            v6 = (v7 == 4294967295)
                            if (-327 > 396) {
                                v46 := GenS3{F0: (v30 >> (((v7 ^ v7)) & 63))}
                                var v47 float32 = (((0.0 / float32(float64(-2.25))) * -2.25) / -2.25)
                                var v48 int32 = 0
                                _ = v46
                                _ = v47
                                _ = v48
                            } else {
                                if false {
                                    v49 := v1["a"]
                                    var v50 uint = v7
                                    var v51 float64 = ((0.1 + 1e10) + 0.0)
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                                v52 := GenS4{F0: -1, F1: 100}
                                _ = v52
                            }
                        }
                        v30 += (v30 + v30)
                    default:
                        v53 := GenS1{F0: 42, F1: 2}
                        _ = v53
                    }
                    if v6 {
                        v30 -= v30
                        v7 = v7
                    }
                }
                switch fi28 {
                case 42:
                    v54 := []int{}
                    _ = v54
                default:
                    _ = len(v2)
                }
                switch (fi29 - (fi28 * 0)) {
                case 128:
                    switch "" {
                    case "a":
                        if (438 >= 104) {
                            var v55 int64 = 0
                            v55 = v55
                            _ = v55
                        } else {
                            var v56 string = "foo bar"
                            v6 = true
                            v56 = v56
                            _ = v56
                        }
                        v6 = v6
                    case "foo bar":
                        v5 = (^(v5))
                    case "":
                        switch ("" + ("a" + "\n\t")) {
                        case "":
                            v57 := []int{605}
                            _ = v57
                        case "\n\t":
                            var v58 int8 = 42
                            _ = v58
                        default:
                            v6 = true
                        }
                        v30 = v30
                        v59 := GenS4{F0: -9223372036854775808, F1: 1}
                        _ = v59
                    default:
                        if (v6 || ((1 <= 606) || (-1 < -1))) {
                            v7 -= (v7 & 4294967295)
                            if true {
                                fmt.Println(v5, v6)
                                if (v5 >= v5) {
                                    fmt.Println(v5, v7)
                                    v60 := GenS1{F0: uint32(v5), F1: 997}
                                    v61 := GenS5{F0: 1, F1: 330}
                                    v6 = v6
                                    switch 0 {
                                    case 42:
                                        fmt.Println(fi29)
                                    }
                                    _ = v60
                                    _ = v61
                                } else {
                                    v6 = ((^(v7)) < (v7 * v7))
                                }
                                if (575 < (v30 + 42)) {
                                    switch fi28 {
                                    case -2147483648:
                                        fmt.Println(fi29)
                                    case -759:
                                        switch "a" {
                                        case "a":
                                            v62 := GenS5{F0: 42, F1: 4294967295}
                                            _ = v62
                                        case "":
                                            v63 := []int{9223372036854775807, fi29, int(v30), (-(fi29))}
                                            _ = v63
                                        case "\n\t":
                                            v6 = v6
                                        default:
                                            v5 = (v5 >> (((v7 | (v7 * v7))) & 63))
                                        }
                                    case -9223372036854775808:
                                        if v6 {
                                            v5 += (v5 >> (((v7 | v7)) & 63))
                                            v7 = (^(v7))
                                            v64 := []int{fi29, fi28, 4294967295}
                                            _ = v64
                                        }
                                    default:
                                        var v65 int32 = -2147483648
                                        var v66 uint64 = 1
                                        v67 := GenS6{F0: 32767, F1: 2, F2: (!true)}
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    }
                                    var v68 int64 = -1
                                    v69 := GenS7{F0: 0, F1: -761, F2: 1}
                                    _ = v68
                                    _ = v69
                                }
                            }
                        }
                        switch ((((("" + "a") + "hello") + "\n\t") + "hello") + ("\n\t" + "\n\t")) {
                        case "\n\t":
                            var v70 uint = v7
                            _ = v70
                        default:
                            if false {
                                fmt.Println(v30, fi29, fi28)
                                v71 := []int{(^(fi28))}
                                _ = v71
                            }
                        }
                    }
                case -2147483648:
                    var v72 uint64 = 2
                    _ = v72
                case 32768:
                    var v73 int16 = 0
                    _ = v73
                default:
                    v74 := GenS8{F0: 244}
                    _ = v74
                }
                v30 = 2
                v75 := []int{}
                _ = fi29
                _ = v30
                _ = v75
            }
            v76 := GenS1{F0: uint32(v5), F1: 42}
            _ = fi28
            _ = v76
        }
        for fi77 := 0; fi77 < 15; fi77++ {
            delete(v1, "foo bar")
            _ = fi77
        }
        if (!(!((273 >= 449) || v6))) {
            if false {
                v6 = (2 == 845)
            } else {
                if (!(42 != 32767)) {
                    if true {
                        fmt.Println(v6, v5)
                    } else {
                        var v78 int = -788
                        if v6 {
                            var v79 string = ""
                            v78 = (v78 % ((v78) | 1))
                            _ = v79
                        }
                        v6 = false
                        _ = v78
                    }
                }
            }
        } else {
            fmt.Println(v6, v6, v5)
        }
    }
    var v80 float64 = -2.25
    var v81 int = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v80
    _ = v81
}
