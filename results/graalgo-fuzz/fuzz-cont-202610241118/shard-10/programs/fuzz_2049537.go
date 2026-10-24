package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 uint16; F1 uint64; F2 int }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 int8 }
func genFn0() int8 {
    v6 := []int{127}
    v6 = append(v6, 658)
    v7 := GenS1{F0: ((("foo bar" + "hello") + ("a" + "")) + ""), F1: 1}
    if (0 >= 18446744073709551615) {
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                if false {
                    v10 := GenS1{F0: "", F1: 0}
                    _ = v10
                }
                v6 = append(v6, (^(fi9)))
                var v11 uint32 = 916
                if ((1 >= 712) || false) {
                    if false {
                        var v12 float64 = 1.5
                        v11 += (^(v11))
                        _ = v12
                    }
                    if (true || false) {
                        var v13 int = -757
                        v6 = append(v6, (fi9 | 32767))
                        if (2147483647 < 2147483647) {
                            v13 = (fi9 & -193)
                        } else {
                            v11 += v11
                            v6 = append(v6, (-(v13)))
                            var v14 float32 = float32((((((0.1 / 1.0) * float64(0.1)) - 1.5) * 1.5) * 1e10))
                            _ = v14
                        }
                        _ = v13
                    } else {
                        v15 := GenS1{F0: "hello", F1: -466}
                        _ = len(v6)
                        v16 := GenS2{F0: -2147483648}
                        _ = v15
                        _ = v16
                    }
                    _ = len(v6)
                }
                v17 := GenS2{F0: 0}
                _ = fi9
                _ = v11
                _ = v17
            }
            _ = fi8
        }
        switch "hello" {
        case "\n\t":
            var v18 uint16 = 255
            _ = v18
        case "hello":
            for fi19 := 0; fi19 < 16; fi19++ {
                v6 = append(v6, fi19)
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi19
            }
        case "":
            switch "" {
            case "foo bar":
                for fi20 := 0; fi20 < 3; fi20++ {
                    switch (("" + "a") + "hello") {
                    case "":
                        var v21 int64 = 1
                        _ = v21
                    case "foo bar":
                        for fi22 := 0; fi22 < 1; fi22++ {
                            switch "" {
                            case "\n\t":
                                var v23 int = -2147483648
                                _ = v23
                            case "":
                                if (0 == 2) {
                                    var v24 uint8 = 127
                                    v25 := GenS2{F0: 684}
                                    _ = v24
                                    _ = v25
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = (fi20 ^ 886)
                                    }
                                    if (!(100 == 127)) {
                                        var v26 int = (fi20 >> ((771) & 63))
                                        var v27 float32 = (0.0 + 1.5)
                                        _ = v26
                                        _ = v27
                                    } else {
                                        var v28 uint64 = 1
                                        _ = v28
                                    }
                                }
                            case "foo bar":
                                v29 := GenS2{F0: 1}
                                _ = v29
                            default:
                                if true {
                                    if false {
                                        var v30 float32 = 0.1
                                        if (1000 > 1000) {
                                            v30 += v30
                                        } else {
                                            v30 = v30
                                            v30 = v30
                                        }
                                        _ = v30
                                    } else {
                                        var v31 int16 = -1
                                        _ = v31
                                    }
                                    v32 := GenS2{F0: 1}
                                    var v33 int32 = -2147483648
                                    _ = v32
                                    _ = v33
                                }
                            }
                            fmt.Println(fi22, fi20)
                            _ = fi22
                        }
                    }
                    _ = fi20
                }
            case "\n\t":
                var v34 uint8 = 255
                _ = v34
            }
        }
        if false {
            var v35 uint8 = 0
            switch -147 {
            case 128:
                if (145 == -403) {
                    for fi36 := 0; fi36 < 3; fi36++ {
                        if len(v6) > 0 {
                            v6[0] = fi36
                        }
                        switch fi36 {
                        case -9223372036854775808:
                            v37 := GenS2{F0: 0}
                            _ = v37
                        case 9223372036854775807:
                            fmt.Println(v35, v35)
                        case 32767:
                            fmt.Println(fi36, fi36)
                        default:
                            v35 -= (v35 % (((v35 | (v35 * v35))) | 1))
                        }
                        _ = fi36
                    }
                    var v38 int8 = 119
                    _ = v38
                } else {
                    v39 := GenS3{F0: -122}
                    _ = v39
                }
            case -2147483648:
                v40 := GenS4{F0: 100, F1: ("" + "foo bar")}
                _ = v40
            case -1:
                var v41 int32 = 769
                _ = v41
            }
            _ = v35
        }
    } else {
        var v42 int = 128
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v42
    }
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []string{"\n\t", "", "foo bar"}
    var v2 int64 = 9223372036854775807
    var v3 int = -2147483648
    var v4 int8 = 100
    var v5 int8 = -82
    if true {
        if true {
            _ = len(v1)
        } else {
            var v43 int32 = -635
            v2 = v2
            if ((2147483647 < 872) || (663 == (v2 | v2))) {
                switch v3 {
                case 2147483647:
                    v44 := GenS5{F0: v4, F1: 299}
                    _ = v44
                default:
                    var v45 string = "\n\t"
                    _ = v45
                }
                v5 = genFn0()
                var v46 uint16 = 210
                _ = v46
            } else {
                for fi47 := 0; fi47 < 20; fi47++ {
                    v48 := GenS5{F0: (v5 / ((genFn0()) | 1)), F1: 252}
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    v3 += (-(v3))
                    _ = fi47
                    _ = v48
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = v43
        }
        v4 = 0
        v49 := GenS6{F0: 0, F1: 158, F2: (v3 + v3)}
        _ = v49
    }
    var v50 string = "a"
    if len(v1) > 0 {
        v1[0] = v50
    }
    fmt.Println(v4, v2)
    fmt.Println(v50, v4, v3)
    v2 += v2
    v51 := GenS5{F0: (-(v5)), F1: 18446744073709551615}
    for fi52 := 0; fi52 < 15; fi52++ {
        v53 := GenS7{F0: 42, F1: 2}
        _ = fi52
        _ = v53
    }
    v54 := GenS8{F0: v5}
    var v55 uint8 = 42
    v50 = v50
    var v56 uint64 = 18446744073709551615
    fmt.Println(v56, v3)
    v50 = v50
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v54
    _ = v55
    _ = v56
}
