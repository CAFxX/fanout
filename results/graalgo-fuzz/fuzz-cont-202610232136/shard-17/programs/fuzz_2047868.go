package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int32; F2 int16 }
type GenS4 struct { F0 int64; F1 string; F2 int64 }
func genFn0(p0 string) uint32 {
    if (-9223372036854775808 == -1) {
        v7 := []int{4294967295}
        _ = v7
    } else {
        if (false || (0 > 1)) {
            v8 := []int{-103, 0, -163}
            var v9 int16 = 1000
            _ = v8
            _ = v9
        } else {
            var v10 int32 = 0
            _ = v10
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := []int{(fi11 >> ((4294967295) & 63)), fi11, fi11, fi11}
            p0 = p0
            _ = fi11
            _ = v12
        }
    }
    if (109 != 100) {
        for fi13 := 0; fi13 < 16; fi13++ {
            switch (p0 + p0) {
            case "foo bar":
                v14 := []int{9223372036854775807, -1, -18}
                _ = v14
            case "":
                for fi15 := 0; fi15 < 4; fi15++ {
                    if false {
                        v16 := []int{-323, (fi15 ^ 32768)}
                        p0 = p0
                        _ = v16
                    }
                    if false {
                        var v17 int32 = 42
                        var v18 uint16 = 255
                        _ = v17
                        _ = v18
                    } else {
                        v19 := GenS1{F0: 0, F1: 1}
                        p0 = ("hello" + p0)
                        var v20 uint16 = 1
                        _ = v19
                        _ = v20
                    }
                    var v21 float64 = 1.5
                    _ = fi15
                    _ = v21
                }
            default:
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 float64 = 3.14159
                    v24 := []int{fi13, (^(fi13)), (fi13 % ((0) | 1))}
                    p0 = ((("a" + ("hello" + p0)) + ((("a" + "a") + "foo bar") + ("" + p0))) + p0)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            var v25 string = ((p0 + ("hello" + ((p0 + "") + "a"))) + "\n\t")
            var v26 uint64 = 2
            _ = fi13
            _ = v25
            _ = v26
        }
        v27 := []int{32768, 1, 4294967295, 866}
        _ = v27
    } else {
        if true {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 uint = 42
                v30 := []int{fi28, fi28, (fi28 & (fi28 + fi28))}
                if len(v30) > 0 {
                    v30[0] = fi28
                }
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            if (false || (!true)) {
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 bool = ((127 < 42) || false)
                    v33 := GenS1{F0: 2147483647, F1: 466}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                v34 := GenS1{F0: 2147483648, F1: 1}
                _ = v34
            } else {
                var v35 uint64 = 18446744073709551615
                _ = v35
            }
        }
        var v36 float32 = ((1.5 * float32(((0.1 + 1e10) - 1e10))) + (float32(((0.1 / float64(float32(3.14159))) * (0.1 + 0.1))) + (1e10 * 3.14159)))
        _ = v36
    }
    var v37 uint64 = 42
    var v38 bool = (-2147483648 < 224)
    _ = p0
    _ = v37
    _ = v38
    return 2
}
func genFn1(p0 uint32) string {
    p0 = (p0 - 4294967295)
    var v39 int16 = 42
    for fi40 := 0; fi40 < 11; fi40++ {
        fmt.Println(p0, fi40, v39)
        v39 -= (v39 * v39)
        var v41 string = "\n\t"
        _ = fi40
        _ = v41
    }
    _ = p0
    _ = v39
    return ("hello" + "")
}

func main() {
    var v1 float64 = float64((-2.25 - (float32(-2.25) - 1.5)))
    var v2 uint16 = 1
    var v3 int8 = 20
    var v4 uint = 255
    var v5 float64 = ((-2.25 * (1.5 / (1e10 + 3.14159))) + v1)
    v6 := make(map[uint32]int)
    v2 += 1
    var v42 uint16 = (v2 % (((v2 % ((290) | 1))) | 1))
    v43 := GenS1{F0: 816, F1: 4294967295}
    for fi44 := 0; fi44 < 6; fi44++ {
        delete(v6, uint32(v4))
        for fi45 := 0; fi45 < 1; fi45++ {
            var v46 int64 = -9223372036854775808
            if ((v3 % ((v3) | 1)) > -1) {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v4)
                default:
                    v47 := []int{int(v46), fi44, fi45, (fi44 >> (((v4 / (((v4 & (v4 | v4))) | 1))) & 63))}
                    _ = v47
                }
                fmt.Println(v42, v46, v42)
                v2 -= v42
            } else {
                var v48 uint16 = (^(v42))
                v6[2] = fi45
                var v49 uint64 = 24
                v50 := GenS2{F0: "foo bar"}
                v51 := v6[2147483647]
                v46 += (v46 * (v46 / ((-9223372036854775808) | 1)))
                _ = v48
                _ = v49
                _ = v50
                _ = v51
            }
            fmt.Println(v2, v4)
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    v52 := GenS2{F0: ""}
    v53 := GenS3{F0: (!true), F1: 42, F2: 1000}
    v54 := GenS2{F0: ("\n\t" + "\n\t")}
    for fi55 := 0; fi55 < 19; fi55++ {
        for fi56 := 0; fi56 < 4; fi56++ {
            fmt.Println(v1, fi55, v5)
            v57 := GenS3{F0: (-2.25 <= (-2.25 / 0.1)), F1: 2147483647, F2: -649}
            _ = fi56
            _ = v57
        }
        _ = fi55
    }
    if true {
        var v58 int32 = 11
        v6[42] = 4294967295
        _ = v58
    } else {
        if (42 > 459) {
            var v59 uint = (v4 >> (((v4 >> (((v4 >> ((2) & 63))) & 63))) & 63))
            if false {
                var v60 uint32 = 4294967295
                _ = v60
            }
            _ = v59
        }
        var v61 uint = (v4 >> ((v4) & 63))
        _ = v61
    }
    switch 32767 {
    case 584:
        switch 343 {
        case 4294967295:
            v42 += (v2 | 42)
        case 2147483647:
            for fi62 := 0; fi62 < 11; fi62++ {
                v63 := GenS4{F0: 1, F1: "foo bar", F2: 42}
                v64 := []int{fi62}
                fmt.Println(v3)
                _ = fi62
                _ = v63
                _ = v64
            }
        case -9223372036854775808:
            delete(v6, 136)
        default:
            switch "\n\t" {
            case "hello":
                var v65 bool = (1e10 <= 0.1)
                _ = v65
            default:
                v4 = v4
            }
        }
    case -9223372036854775808:
        fmt.Println(v1, v2, v4)
    default:
        var v66 int = -1
        _ = v66
    }
    for fi67 := 0; fi67 < 10; fi67++ {
        var v68 float64 = v1
        _ = fi67
        _ = v68
    }
    fmt.Println(v3, v1, v5)
    var v69 int8 = v3
    v6[308] = -524
    v42 -= v42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v6
    _ = v69
}
