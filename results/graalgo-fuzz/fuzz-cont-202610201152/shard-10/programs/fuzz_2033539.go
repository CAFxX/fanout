package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint64; F1 bool; F2 bool }
type GenS4 struct { F0 int64; F1 int32; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 bool; F1 uint32; F2 int64 }

func main() {
    var v1 int8 = 42
    v2 := make(map[string]uint32)
    var v3 uint16 = 1
    v4 := make(map[int]int)
    var v5 int8 = v1
    var v6 uint16 = 1
    if (42 < 4) {
        if (79 >= 42) {
            v7 := GenS1{F0: "hello"}
            v8 := []int{}
            _ = v7
            _ = v8
        }
        var v9 int64 = -9223372036854775808
        v10 := []int{32767}
        _ = v10
        _ = v9
    } else {
        v1 -= -1
        fmt.Println(v3)
    }
    var v11 bool = (false || true)
    if true {
        var v12 float64 = 1e10
        for fi13 := 0; fi13 < 12; fi13++ {
            switch (fi13 + (fi13 ^ (fi13 % (((fi13 >> ((2) & 63))) | 1)))) {
            case -600:
                var v14 uint16 = (v3 + (v3 & v6))
                _ = v14
            case 127:
                switch (fi13 + 42) {
                case 9223372036854775807:
                    var v15 uint = 42
                    _ = v15
                case 24:
                    v11 = v11
                case 249:
                    var v16 uint32 = 1
                    _ = v16
                }
            case 128:
                v17 := GenS2{F0: (v1 | (v5 | v1)), F1: 255}
                _ = v17
            default:
                fmt.Println(v1, fi13, v11)
            }
            _ = fi13
        }
        _ = len(v2)
        _ = v12
    } else {
        v3 = 255
        var v18 uint16 = (^(v3))
        _ = v18
    }
    v11 = (!(255 > 945))
    fmt.Println(v3, v5, v5)
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 int8 = (v5 + v1)
        var v21 float64 = (0.1 + -2.25)
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println(v5, v11, v6)
    switch -722 {
    case 333:
        v22 := []int{32768, 0, -1}
        _ = v22
    case 127:
        v23 := []int{127, 0, 628, -2147483648}
        _ = v23
    }
    switch 0 {
    case -696:
        for fi24 := 0; fi24 < 3; fi24++ {
            switch 0 {
            case 127:
                if (!((!v11) && v11)) {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 uint32 = 0
                        if false {
                            if (671 <= 255) {
                                _ = len(v2)
                                v5 = (v1 << ((65535) & 63))
                                v27 := GenS3{F0: 4294967295, F1: (133 != 111), F2: v11}
                                var v28 uint8 = 42
                                _ = v27
                                _ = v28
                            } else {
                                var v29 float64 = (3.14159 / 1.0)
                                if v11 {
                                    v30 := GenS1{F0: "hello"}
                                    _ = v30
                                } else {
                                    var v31 int32 = -1
                                    var v32 int8 = (v1 - v5)
                                    switch -954 {
                                    case 2147483647:
                                        v33 := GenS3{F0: 0, F1: (!(((^(v26)) <= v26) || true)), F2: (!v11)}
                                        _ = v33
                                    default:
                                        var v34 uint = 18446744073709551615
                                        _ = v34
                                    }
                                    _ = v31
                                    _ = v32
                                }
                                v2["foo bar"] = (v26 - 2147483648)
                                _ = v29
                            }
                            fmt.Println(fi25, v1)
                        }
                        if v11 {
                            fmt.Println(v5, fi24)
                            var v35 uint32 = (v26 + (v26 << ((4294967295) & 63)))
                            v36 := GenS3{F0: 453, F1: v11, F2: (uint64(v35) == 18446744073709551615)}
                            _ = v35
                            _ = v36
                        } else {
                            if v11 {
                                var v37 float64 = 3.14159
                                _ = v37
                            }
                        }
                        v38 := []int{(-(fi24)), fi25}
                        _ = fi25
                        _ = v26
                        _ = v38
                    }
                } else {
                    v39 := GenS4{F0: 0, F1: 2147483647, F2: 1}
                    if (v5 < 70) {
                        v40 := GenS5{F0: 1}
                        if (!true) {
                            var v41 string = ""
                            for fi42 := 0; fi42 < 17; fi42++ {
                                switch -1 {
                                case 32768:
                                    var v43 float32 = (0.0 * (0.1 / float32(1.5)))
                                    _ = v43
                                case 2147483647:
                                    v2[""] = 1
                                case -9223372036854775808:
                                    fmt.Println(v5, v3)
                                }
                                _ = len(v4)
                                fmt.Println(v5)
                                _ = fi42
                            }
                            _ = v41
                        } else {
                            var v44 uint = 658
                            _ = v44
                        }
                        _ = v40
                    }
                    v3 -= v3
                    _ = v39
                }
            case 4294967295:
                v45 := GenS6{F0: (fi24 >> ((994) & 63)), F1: ""}
                _ = v45
            case 32768:
                var v46 int16 = 1
                _ = v46
            default:
                v47 := GenS1{F0: ("a" + "foo bar")}
                _ = v47
            }
            switch "\n\t" {
            case "\n\t":
                switch "a" {
                case "\n\t":
                    for fi48 := 0; fi48 < 2; fi48++ {
                        v6 = 27
                        var v49 int32 = 753
                        v50 := []int{}
                        var v51 string = "a"
                        var v52 uint32 = 368
                        _ = fi48
                        _ = v49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                case "a":
                    v4[9223372036854775807] = (fi24 / ((fi24) | 1))
                case "hello":
                    fmt.Println(v1, v1, v1)
                }
            case "a":
                v11 = v11
            }
            v11 = v11
            _ = fi24
        }
    default:
        for fi53 := 0; fi53 < 10; fi53++ {
            var v54 int32 = -2147483648
            v55 := GenS7{F0: 711, F1: false}
            for fi56 := 0; fi56 < 3; fi56++ {
                fmt.Println(v54)
                switch "hello" {
                case "a":
                    fmt.Println(v1, v11, v5)
                }
                v57 := []int{fi53}
                v3 = (v6 | v6)
                _ = fi56
                _ = v57
            }
            _ = fi53
            _ = v54
            _ = v55
        }
    }
    for fi58 := 0; fi58 < 10; fi58++ {
        switch (("hello" + "\n\t") + "\n\t") {
        case "foo bar":
            var v59 float32 = (0.1 * 0.1)
            _ = v59
        case "\n\t":
            if (0 >= -986) {
                for fi60 := 0; fi60 < 10; fi60++ {
                    if (695 > 4294967295) {
                        fmt.Println(v1, v5, v11)
                        var v61 uint64 = 208
                        fmt.Println(v61, v1, fi58)
                        _ = v61
                    }
                    v6 = (v6 - v6)
                    fmt.Println(v3, v11)
                    _ = fi60
                }
                v11 = (!(966 < -1))
            }
        case "":
            v6 = (v6 & (v6 - v6))
        }
        _ = fi58
    }
    _ = len(v2)
    if (42 != 4294967295) {
        if v11 {
            v5 = (v5 * (v1 ^ v5))
        }
        for fi62 := 0; fi62 < 3; fi62++ {
            switch "\n\t" {
            case "foo bar":
                v63 := GenS6{F0: (fi62 + fi62), F1: "foo bar"}
                _ = v63
            case "\n\t":
                fmt.Println(v5, fi62)
            default:
                var v64 uint8 = 127
                _ = v64
            }
            v65 := []int{}
            _ = fi62
            _ = v65
        }
        for fi66 := 0; fi66 < 15; fi66++ {
            switch "\n\t" {
            case "a":
                for fi67 := 0; fi67 < 5; fi67++ {
                    var v68 int32 = 42
                    v69 := GenS8{F0: v11, F1: 65535, F2: 1}
                    _ = fi67
                    _ = v68
                    _ = v69
                }
            case "":
                v6 += v6
            default:
                var v70 bool = v11
                _ = v70
            }
            v5 -= v5
            _ = fi66
        }
    } else {
        var v71 int8 = (v1 * v1)
        var v72 int16 = 42
        fmt.Println(v5, v72, v72)
        _ = v71
        _ = v72
    }
    v73 := GenS1{F0: (("" + "a") + ("a" + "a"))}
    for fi74 := 0; fi74 < 19; fi74++ {
        if ((0.1 < -2.25) || (0.1 <= float64(-2.25))) {
            fmt.Println(fi74, v6, v1)
            var v75 float64 = 1e10
            var v76 uint = 528
            _ = v75
            _ = v76
        } else {
            var v77 uint32 = 0
            _ = v77
        }
        v6 = v6
        v11 = v11
        _ = fi74
    }
    var v78 uint32 = 658
    v11 = (0.0 <= (0.1 * (1.5 / (float64(-2.25) / ((3.14159 + 0.1) + 0.1)))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v78
}
