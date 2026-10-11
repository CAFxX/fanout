package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int; F1 int; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint }
func genFn0(p0 uint32, p1 int, p2 uint8) uint {
    var v7 uint32 = p0
    p2 += (p2 % ((1) | 1))
    v7 -= 2147483648
    var v8 int8 = -128
    var v9 int = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 948
}
func genFn1() int32 {
    var v10 bool = false
    if false {
        var v11 string = ""
        _ = v11
    } else {
        v12 := GenS1{F0: 4294967295, F1: (!(!(176 != 2147483647)))}
        v10 = (-1 == -169)
        if (255 >= 160) {
            if false {
                fmt.Println(v10)
                switch ((("foo bar" + (("a" + "hello") + "a")) + ("foo bar" + "")) + "") {
                case "":
                    v10 = ((0.0 / 1.0) <= 0.1)
                case "foo bar":
                    v13 := []int{-679, 594}
                    _ = v13
                case "hello":
                    var v14 int8 = 100
                    _ = v14
                }
                v10 = (float32(1e10) <= 0.1)
            } else {
                var v15 uint32 = 4294967295
                _ = v15
            }
            switch -359 {
            case 4294967295:
                v16 := []int{32767, -112, 2147483647, -2147483648}
                _ = v16
            case 9223372036854775807:
                v17 := []int{}
                _ = v17
            case 1:
                var v18 float32 = (3.14159 - 0.0)
                _ = v18
            }
            var v19 uint32 = 65535
            _ = v19
        } else {
            var v20 int = 32767
            _ = v20
        }
        _ = v12
    }
    if v10 {
        var v21 uint32 = 36
        for fi22 := 0; fi22 < 1; fi22++ {
            switch (fi22 + -9223372036854775808) {
            case 1:
                v23 := GenS1{F0: 18446744073709551615, F1: (!false)}
                _ = v23
            case -2147483648:
                for fi24 := 0; fi24 < 14; fi24++ {
                    v21 -= (v21 ^ v21)
                    if (42 == 127) {
                        var v25 int8 = 1
                        _ = v25
                    }
                    var v26 int = 1
                    var v27 int16 = 1
                    _ = fi24
                    _ = v26
                    _ = v27
                }
            case 745:
                v10 = (-427 < 42)
            default:
                for fi28 := 0; fi28 < 15; fi28++ {
                    if false {
                        v29 := []int{739, (^(fi22)), fi22}
                        var v30 string = "foo bar"
                        if (42 >= 1) {
                            if (65535 < 255) {
                                var v31 uint32 = (v21 << ((4294967295) & 63))
                                var v32 uint64 = 0
                                v32 = v32
                                v10 = (0.1 >= 1e10)
                                _ = v31
                                _ = v32
                            }
                            var v33 int64 = 0
                            v30 = ""
                            _ = v33
                        }
                        var v34 uint64 = 0
                        _ = v29
                        _ = v30
                        _ = v34
                    } else {
                        v21 = v21
                        switch 127 {
                        case 0:
                            v21 += (v21 >> ((267) & 63))
                        }
                    }
                    v35 := GenS2{F0: (fi22 | fi22), F1: (-(fi28)), F2: -135}
                    _ = fi28
                    _ = v35
                }
            }
            _ = fi22
        }
        _ = v21
    }
    fmt.Println(v10)
    _ = v10
    return -838
}

func main() {
    var v1 int64 = 0
    v2 := make(map[int]int)
    var v3 int32 = 1
    v4 := make(map[int]uint32)
    v5 := []uint32{}
    v6 := make(map[int]int)
    switch "" {
    case "":
        switch -9223372036854775808 {
        case 25:
            if false {
                v3 = v3
                for fi36 := 0; fi36 < 4; fi36++ {
                    switch "\n\t" {
                    case "\n\t":
                        v3 -= genFn1()
                    case "foo bar":
                        var v37 bool = (false || true)
                        _ = v37
                    }
                    fmt.Println(v1, v1)
                    _ = fi36
                }
                v1 += v1
            } else {
                for fi38 := 0; fi38 < 12; fi38++ {
                    for fi39 := 0; fi39 < 6; fi39++ {
                        fmt.Println(fi39, v1, v1)
                        v3 = v3
                        _ = fi39
                    }
                    var v40 int64 = v1
                    _ = fi38
                    _ = v40
                }
                v41 := GenS1{F0: 971, F1: ((v3 & -2147483648) < v3)}
                v1 = (-(v1))
                _ = v41
            }
        case -517:
            var v42 int8 = -128
            _ = v42
        case 767:
            var v43 float64 = 0.1
            _ = v43
        }
    case "a":
        var v44 float64 = 1e10
        _ = v44
    case "foo bar":
        v4[0] = 4294967295
    }
    switch -662 {
    case -9223372036854775808:
        switch ((("hello" + "hello") + "hello") + ("\n\t" + "hello")) {
        case "\n\t":
            var v45 uint8 = 86
            _ = v45
        }
    }
    v5 = append(v5, 2147483648)
    v3 = (^(v3))
    v3 -= v3
    v1 -= v1
    var v46 int = 0
    v46 -= (v46 >> ((genFn0(4294967295, 32767, 2)) & 63))
    _ = len(v4)
    v47 := GenS2{F0: 4294967295, F1: 127, F2: (v1 ^ v1)}
    fmt.Println(v46, v3, v3)
    var v48 int64 = (-(v1))
    fmt.Println(v3, v46)
    if (0.1 >= 1e10) {
        if (v3 == genFn1()) {
            v49 := GenS1{F0: 42, F1: (3.14159 >= 1.5)}
            for fi50 := 0; fi50 < 7; fi50++ {
                var v51 int8 = 0
                if true {
                    var v52 float64 = float64(0.1)
                    for fi53 := 0; fi53 < 9; fi53++ {
                        _ = len(v2)
                        fmt.Println(v51, v3)
                        if true {
                            delete(v6, (-(v46)))
                            switch (v46 ^ 2147483647) {
                            case 9223372036854775807:
                                v3 -= 42
                            case -9223372036854775808:
                                v51 += v51
                            default:
                                var v54 int64 = v1
                                _ = v54
                            }
                        } else {
                            fmt.Println(v51, fi50)
                        }
                        _ = fi53
                    }
                    _ = v52
                }
                _ = fi50
                _ = v51
            }
            v55 := GenS3{F0: "\n\t"}
            _ = v49
            _ = v55
        } else {
            var v56 int16 = -670
            v57 := GenS4{F0: 1, F1: v56}
            var v58 int32 = (v3 ^ 42)
            _ = v56
            _ = v57
            _ = v58
        }
        if (2 <= 42) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            var v59 int64 = int64(v3)
            var v60 int = v46
            _ = v59
            _ = v60
        }
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v61 uint = 18446744073709551615
        _ = v61
    }
    if (1.5 == 1e10) {
        v6[83] = (^(v46))
    } else {
        if (!true) {
            switch "a" {
            case "foo bar":
                var v62 int8 = 88
                _ = v62
            case "hello":
                delete(v2, (^(v46)))
            case "":
                v63 := GenS3{F0: "hello"}
                _ = v63
            default:
                v48 = 345
            }
        } else {
            for fi64 := 0; fi64 < 11; fi64++ {
                for fi65 := 0; fi65 < 2; fi65++ {
                    var v66 int64 = (-(v48))
                    switch "a" {
                    case "hello":
                        v1 = (v48 + 770)
                    case "\n\t":
                        var v67 int16 = 196
                        _ = v67
                    case "a":
                        if (!true) {
                            var v68 uint = 2
                            var v69 string = ("" + "")
                            switch v69 {
                            case "hello":
                                var v70 uint = genFn0(2, 128, 1)
                                _ = v70
                            }
                            if (42 >= 851) {
                                var v71 bool = false
                                v72 := GenS4{F0: v48, F1: 0}
                                v73 := v4[(fi64 + 895)]
                                _ = v71
                                _ = v72
                                _ = v73
                            } else {
                                v74 := GenS1{F0: 747, F1: true}
                                if false {
                                    v75 := GenS1{F0: 1, F1: false}
                                    fmt.Println(v46, v68, v1)
                                    v76 := GenS4{F0: v66, F1: -32768}
                                    _ = v75
                                    _ = v76
                                }
                                _ = v74
                            }
                            _ = v68
                            _ = v69
                        }
                    default:
                        v48 += (v66 * (-(v66)))
                    }
                    _ = fi65
                    _ = v66
                }
                if false {
                    fmt.Println(fi64)
                }
                _ = fi64
            }
            switch v46 {
            case 127:
                if false {
                    var v77 int16 = 0
                    fmt.Println(v1)
                    _ = v77
                }
            case 42:
                v78 := GenS5{F0: 1000, F1: 128, F2: 74}
                _ = v78
            case 9223372036854775807:
                for fi79 := 0; fi79 < 8; fi79++ {
                    if false {
                        fmt.Println(fi79, fi79, v48)
                    } else {
                        v1 += -9223372036854775808
                        var v80 uint = 42
                        if false {
                            if (42 <= 255) {
                                v80 += v80
                                var v81 int16 = 781
                                var v82 bool = false
                                _ = v81
                                _ = v82
                            }
                            for fi83 := 0; fi83 < 6; fi83++ {
                                if true {
                                    v80 -= (^(v80))
                                }
                                _ = fi83
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        _ = v80
                    }
                    _ = fi79
                }
            }
            var v84 uint32 = 399
            _ = v84
        }
        var v85 string = "a"
        _ = v85
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
