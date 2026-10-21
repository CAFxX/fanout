package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int64 }
func genFn0(p0 int16, p1 int8) uint8 {
    if ((!true) || (!false)) {
        p1 += (p1 | (p1 + -1))
    } else {
        if true {
            v5 := GenS1{F0: 127}
            var v6 uint = 1
            v6 += v6
            _ = v5
            _ = v6
        }
        switch (("hello" + ("a" + "foo bar")) + (("\n\t" + "hello") + "foo bar")) {
        case "":
            v7 := []int{2147483647, 2147483647, -946, 242}
            _ = v7
        case "hello":
            fmt.Println(p0, p0)
        case "a":
            if false {
                p1 += p1
                for fi8 := 0; fi8 < 11; fi8++ {
                    if true {
                        v9 := []int{}
                        _ = v9
                    }
                    _ = fi8
                }
            }
        default:
            var v10 uint64 = 42
            _ = v10
        }
        fmt.Println(p1, p1)
    }
    if (38 == 1) {
        for fi11 := 0; fi11 < 13; fi11++ {
            switch ((((("foo bar" + ("foo bar" + "hello")) + "") + ("\n\t" + (("a" + "\n\t") + ("foo bar" + ("foo bar" + "foo bar"))))) + ("\n\t" + ("a" + "hello"))) + ("a" + (("a" + "") + "hello"))) {
            case "foo bar":
                fmt.Println(fi11, p0)
            case "hello":
                v12 := GenS1{F0: 1}
                _ = v12
            case "":
                var v13 uint8 = 255
                _ = v13
            default:
                v14 := []int{fi11, fi11, fi11, (^(fi11))}
                _ = v14
            }
            _ = fi11
        }
        var v15 float32 = (0.0 + 0.1)
        _ = v15
    } else {
        p1 += 0
        if (-1 < 2147483647) {
            var v16 string = "foo bar"
            v17 := GenS1{F0: 27}
            _ = v16
            _ = v17
        }
    }
    var v18 int32 = 0
    switch "\n\t" {
    case "foo bar":
        if true {
            v19 := []int{9223372036854775807, -1}
            _ = v19
        } else {
            if true {
                var v20 float64 = 0.1
                for fi21 := 0; fi21 < 2; fi21++ {
                    if (p1 >= -1) {
                        if (42 >= 392) {
                            p0 = -32768
                            v22 := []int{0, -358, (fi21 << ((0) & 63)), (^(fi21))}
                            if (0.1 != (1e10 - 1.5)) {
                                var v23 int16 = -581
                                _ = v23
                            } else {
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            }
                            _ = v22
                        }
                        p1 -= p1
                        if true {
                            v24 := GenS2{F0: ("foo bar" + "hello")}
                            p1 = -67
                            fmt.Println(v18, p0)
                            _ = v24
                        }
                    }
                    _ = fi21
                }
                v25 := []int{199, -863}
                _ = v20
                _ = v25
            } else {
                switch -1 {
                case 1:
                    if (1 > 0) {
                        fmt.Println(v18, v18)
                    }
                }
            }
            v26 := []int{127}
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 string = ("hello" + ("foo bar" + "foo bar"))
                _ = fi27
                _ = v28
            }
            _ = v26
        }
    default:
        v29 := []int{450}
        _ = v29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 uint8 = 1
            _ = fi31
            _ = v32
        }
        switch ("a" + "foo bar") {
        case "foo bar":
            if (1 != -9223372036854775808) {
                v33 := GenS3{F0: 2, F1: 0}
                _ = v33
            }
        case "hello":
            var v34 string = "a"
            _ = v34
        }
        _ = fi30
    }
    _ = p0
    _ = p1
    _ = v18
    return 42
}

func main() {
    v1 := []int{11}
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    v35 := GenS2{F0: "a"}
    switch ("\n\t" + "") {
    case "hello":
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 uint = 42
            fmt.Println(fi36, v2)
            fmt.Println(v2, v37)
            _ = fi36
            _ = v37
        }
    case "":
        if true {
            var v38 float32 = (0.1 / ((((3.14159 / 1.0) - 1e10) * 1e10) / 1e10))
            if true {
                var v39 float64 = 1e10
                if (857 <= 42) {
                    for fi40 := 0; fi40 < 3; fi40++ {
                        var v41 float32 = v38
                        _ = fi40
                        _ = v41
                    }
                }
                if (((65535 != 319) || (!(34 > 146))) && (false && false)) {
                    v38 = (v38 - 0.0)
                    var v42 uint16 = 1
                    _ = v42
                } else {
                    fmt.Println(v2)
                }
                _ = v39
            } else {
                v1 = append(v1, 127)
                v43 := v4[-971]
                for fi44 := 0; fi44 < 9; fi44++ {
                    for fi45 := 0; fi45 < 9; fi45++ {
                        v46 := GenS4{F0: 18446744073709551615, F1: 971, F2: v2}
                        _ = len(v4)
                        _ = fi45
                        _ = v46
                    }
                    _ = fi44
                }
                _ = v43
            }
            _ = v38
        } else {
            if (255 == 0) {
                for fi47 := 0; fi47 < 18; fi47++ {
                    for fi48 := 0; fi48 < 7; fi48++ {
                        var v49 string = "foo bar"
                        if false {
                            v49 = v49
                        }
                        var v50 uint16 = 285
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    for fi51 := 0; fi51 < 9; fi51++ {
                        var v52 int64 = v2
                        if false {
                            var v53 uint16 = 1
                            v1 = append(v1, (-(fi51)))
                            _ = v53
                        }
                        _ = fi51
                        _ = v52
                    }
                    _ = fi47
                }
                for fi54 := 0; fi54 < 4; fi54++ {
                    for fi55 := 0; fi55 < 17; fi55++ {
                        v56 := GenS3{F0: 65535, F1: 9223372036854775807}
                        var v57 float64 = 0.0
                        v57 -= (((v57 - v57) - v57) / 1.0)
                        _ = fi55
                        _ = v56
                        _ = v57
                    }
                    _ = fi54
                }
            } else {
                var v58 int = 0
                switch "" {
                case "hello":
                    if len(v1) > 0 {
                        v1[0] = (v58 - 9223372036854775807)
                    }
                case "foo bar":
                    var v59 int16 = 1
                    _ = v59
                case "":
                    fmt.Println(v58)
                default:
                    for fi60 := 0; fi60 < 13; fi60++ {
                        if false {
                            v58 += (fi60 + (fi60 ^ 42))
                            if len(v1) > 0 {
                                v1[0] = fi60
                            }
                            v2 = -1
                        }
                        for fi61 := 0; fi61 < 13; fi61++ {
                            v2 = v2
                            _ = fi61
                        }
                        var v62 string = "foo bar"
                        _ = fi60
                        _ = v62
                    }
                }
                v58 = (v58 % ((v58) | 1))
                _ = v58
            }
            if ((907 == 1) || (1 <= 128)) {
                fmt.Println(v2)
                v63 := GenS3{F0: 65535, F1: 42}
                _ = v63
            }
            if false {
                if true {
                    v2 -= 1
                    v64 := GenS1{F0: 112}
                    v1 = append(v1, 32768)
                    _ = v64
                } else {
                    fmt.Println(v2)
                    v65 := GenS4{F0: 42, F1: 32767, F2: v2}
                    _ = v65
                }
            }
        }
    case "\n\t":
        v2 = v2
    }
    v2 = (v2 / (((v2 ^ (v2 | v2))) | 1))
    v66 := GenS2{F0: ""}
    v2 = (v2 << ((255) & 63))
    fmt.Println(v2)
    v2 -= 0
    delete(v3, 4294967295)
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    if true {
        for fi67 := 0; fi67 < 5; fi67++ {
            for fi68 := 0; fi68 < 16; fi68++ {
                switch "hello" {
                case "\n\t":
                    if false {
                        v2 += 1
                        var v69 float32 = 1e10
                        _ = v69
                    } else {
                        fmt.Println(fi67, fi68)
                        _ = len(v4)
                    }
                }
                v70 := GenS1{F0: 2}
                switch "hello" {
                case "":
                    fmt.Println(fi68, fi68)
                case "\n\t":
                    v2 = v2
                default:
                    _ = len(v4)
                }
                _ = fi68
                _ = v70
            }
            v1 = append(v1, (fi67 << ((1) & 63)))
            _ = fi67
        }
        _ = len(v1)
        v71 := GenS3{F0: 42, F1: v2}
        _ = v71
    }
    v1 = append(v1, 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v66
}
