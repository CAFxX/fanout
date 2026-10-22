package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int64; F1 int16; F2 uint64 }
type GenS5 struct { F0 int; F1 bool }

func main() {
    var v1 uint16 = 2
    var v2 int16 = 1000
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    v5 := make(map[int]int)
    var v6 int64 = -9223372036854775808
    var v7 float64 = ((3.14159 - -2.25) * (1e10 - 0.0))
    v8 := GenS1{F0: false, F1: "foo bar"}
    _ = len(v3)
    if (2147483647 <= 65535) {
        var v9 int16 = (v2 / (((v2 ^ (v2 + 1))) | 1))
        fmt.Println(v1, v9, v9)
        var v10 uint32 = 65535
        _ = v10
        _ = v9
    } else {
        v2 -= (^(v2))
    }
    if (true && true) {
        var v11 uint64 = 463
        switch 394 {
        case 317:
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 int8 = 127
                switch -2147483648 {
                case -9223372036854775808:
                    if (false && true) {
                        if true {
                            _ = len(v4)
                            switch -2147483648 {
                            case -1:
                                for fi14 := 0; fi14 < 16; fi14++ {
                                    if false {
                                        v1 += v1
                                        v15 := GenS2{F0: 1}
                                        var v16 float64 = ((0.1 * 0.0) - v7)
                                        _ = v15
                                        _ = v16
                                    }
                                    if ((true && (v2 == v2)) || (!false)) {
                                        fmt.Println(v2, v13)
                                    } else {
                                        var v17 uint64 = (^(v11))
                                        var v18 int16 = (v2 * v2)
                                        if (1.5 == 0.0) {
                                            var v19 uint32 = 65535
                                            var v20 uint32 = 714
                                            if false {
                                                v17 = 42
                                                fmt.Println(v17, v2, v19)
                                                v21 := []int{fi12, 32767}
                                                _ = v21
                                            } else {
                                                delete(v5, -9223372036854775808)
                                                v22 := GenS2{F0: 711}
                                                if true {
                                                    v23 := []int{}
                                                    v24 := GenS2{F0: 0}
                                                    if (!false) {
                                                        fmt.Println(v13, v17, fi14)
                                                    }
                                                    _ = v23
                                                    _ = v24
                                                }
                                                _ = v22
                                            }
                                            v11 = (v17 - v11)
                                            switch ("a" + "foo bar") {
                                            case "foo bar":
                                                var v25 int8 = v13
                                                _ = v25
                                            case "":
                                                delete(v5, fi14)
                                            }
                                            _ = v19
                                            _ = v20
                                        } else {
                                            if false {
                                                if false {
                                                    if (881 < 42) {
                                                        if false {
                                                            var v26 string = "\n\t"
                                                            var v27 float32 = (3.14159 / ((0.1 / (3.14159 * (-2.25 - ((((1e10 - 0.0) - (3.14159 * (float32(v7) - (0.0 * (-2.25 * 3.14159))))) + -2.25) - (float32(((0.0 * v7) * 1.5)) * (0.0 * 0.0)))))) + -2.25))
                                                            _ = v26
                                                            _ = v27
                                                        } else {
                                                            var v28 uint32 = 65535
                                                            if (3.14159 >= 0.1) {
                                                                if false {
                                                                    v18 += (v2 % ((-1) | 1))
                                                                    v13 -= (v13 / (((v13 >> ((65535) & 63))) | 1))
                                                                    fmt.Println(v1)
                                                                } else {
                                                                    switch fi12 {
                                                                    case 0:
                                                                        v29 := GenS1{F0: (0 != 0), F1: ""}
                                                                        _ = v29
                                                                    }
                                                                    v30 := GenS2{F0: 42}
                                                                    v18 = v18
                                                                    _ = v30
                                                                }
                                                                var v31 int8 = int8(v6)
                                                                _ = v31
                                                            } else {
                                                                var v32 int16 = v18
                                                                _ = v32
                                                            }
                                                            v33 := []int{}
                                                            if (false && false) {
                                                                if true {
                                                                    v34 := GenS2{F0: 42}
                                                                    _ = v34
                                                                }
                                                                v6 = (-(v6))
                                                            }
                                                            var v35 int = (fi14 / ((fi12) | 1))
                                                            _ = v28
                                                            _ = v33
                                                            _ = v35
                                                        }
                                                        v4[617] = "hello"
                                                    } else {
                                                        if (("foo bar" < ((("foo bar" + "a") + "") + "foo bar")) || true) {
                                                            if (false && (!(true || (false && (181 <= 42))))) {
                                                                var v36 int32 = 590
                                                                var v37 int8 = v13
                                                                if (!(!true)) {
                                                                    v18 -= int16(v17)
                                                                } else {
                                                                    _ = len(v4)
                                                                }
                                                                fmt.Println(v11)
                                                                _ = v36
                                                                _ = v37
                                                            } else {
                                                                var v38 float64 = v7
                                                                _ = v38
                                                            }
                                                            var v39 int64 = (v6 & v6)
                                                            if ((fi14 | -9223372036854775808) > fi14) {
                                                                v40 := GenS3{F0: -167, F1: (v1 >> ((2) & 63))}
                                                                _ = v40
                                                            } else {
                                                                v6 += v39
                                                                if false {
                                                                    v41 := v3[(("a" + "foo bar") + "foo bar")]
                                                                    delete(v5, fi14)
                                                                    fmt.Println(v39)
                                                                    _ = v41
                                                                }
                                                            }
                                                            _ = v39
                                                        } else {
                                                            v1 = 1
                                                        }
                                                    }
                                                    if (!true) {
                                                        fmt.Println(v7, fi14)
                                                    }
                                                    v42 := GenS1{F0: (!(!true)), F1: "\n\t"}
                                                    _ = v42
                                                }
                                            }
                                        }
                                        _ = v17
                                        _ = v18
                                    }
                                    _ = fi14
                                }
                            case 885:
                                v1 = (v1 - 901)
                            case 32768:
                                var v43 uint16 = 2
                                _ = v43
                            }
                        } else {
                            for fi44 := 0; fi44 < 13; fi44++ {
                                v45 := GenS3{F0: (v6 | v6), F1: v1}
                                var v46 int8 = v13
                                _ = fi44
                                _ = v45
                                _ = v46
                            }
                        }
                    } else {
                        v47 := []int{fi12, 9223372036854775807, fi12}
                        _ = v47
                    }
                }
                v7 += ((1e10 - (((v7 - v7) * (v7 + v7)) * ((float64(3.14159) + v7) + v7))) + v7)
                _ = fi12
                _ = v13
            }
        }
        _ = v11
    }
    v48 := GenS4{F0: 0, F1: (v2 * -443), F2: 42}
    if (!false) {
        v49 := GenS5{F0: 926, F1: (42 != 976)}
        v50 := []int{-9223372036854775808, -463}
        if false {
            var v51 uint8 = 1
            var v52 float64 = v7
            var v53 uint32 = 2
            _ = v51
            _ = v52
            _ = v53
        } else {
            fmt.Println(v6, v6, v1)
            if (55 != (v2 >> ((226) & 63))) {
                v54 := GenS1{F0: true, F1: "foo bar"}
                v2 += (v2 % ((v2) | 1))
                v6 = v6
                _ = v54
            } else {
                fmt.Println(v7)
                switch -2147483648 {
                case -9223372036854775808:
                    delete(v4, 2147483647)
                default:
                    fmt.Println(v6)
                }
            }
            var v55 int8 = 16
            _ = v55
        }
        _ = v49
        _ = v50
    }
    var v56 uint = 756
    for fi57 := 0; fi57 < 7; fi57++ {
        fmt.Println(v2)
        for fi58 := 0; fi58 < 19; fi58++ {
            fmt.Println(v7, v56)
            fmt.Println(v7)
            v2 = (^(v2))
            _ = fi58
        }
        _ = fi57
    }
    fmt.Println(v6, v1, v6)
    for fi59 := 0; fi59 < 11; fi59++ {
        v60 := GenS3{F0: v6, F1: uint16(v56)}
        _ = fi59
        _ = v60
    }
    switch "foo bar" {
    case "":
        if false {
            _ = len(v4)
        }
    case "a":
        fmt.Println(v2, v6, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v56
    _ = v6
    _ = v7
    _ = v8
}
