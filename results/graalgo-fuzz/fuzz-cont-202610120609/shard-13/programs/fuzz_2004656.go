package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint64; F2 bool }
func genFn0(p0 int, p1 string) uint {
    if false {
        v7 := GenS1{F0: 100, F1: 749}
        _ = v7
    }
    p0 = (^(p0))
    p0 += (p0 / ((p0) | 1))
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{(fi8 | fi8)}
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int = fi8
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 int32, p1 int64, p2 int32) int16 {
    var v12 uint8 = 128
    var v13 int64 = (p1 / ((-9223372036854775808) | 1))
    if (-1 < 123) {
        fmt.Println(p2, p2, v12)
        var v14 int8 = 100
        if true {
            fmt.Println(p0, v12)
            var v15 uint64 = 0
            var v16 int32 = -208
            _ = v15
            _ = v16
        }
        _ = v14
    }
    if true {
        if false {
            switch "hello" {
            case "hello":
                var v17 float64 = -2.25
                _ = v17
            }
            switch -2147483648 {
            case 2147483647:
                v12 -= 1
            case 9223372036854775807:
                var v18 float32 = 0.0
                _ = v18
            default:
                for fi19 := 0; fi19 < 9; fi19++ {
                    v20 := GenS1{F0: int8(p1), F1: 65535}
                    v21 := []int{fi19}
                    var v22 int8 = -128
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
        } else {
            for fi23 := 0; fi23 < 14; fi23++ {
                for fi24 := 0; fi24 < 6; fi24++ {
                    v13 += (p1 & v13)
                    switch 981 {
                    case 127:
                        v25 := []int{(-(fi23))}
                        _ = v25
                    default:
                        v13 = int64(fi23)
                    }
                    _ = fi24
                }
                switch fi23 {
                case 128:
                    v26 := GenS1{F0: -128, F1: 36}
                    _ = v26
                case -934:
                    v27 := GenS1{F0: 0, F1: 65535}
                    _ = v27
                case 306:
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        v12 = 255
                    case "hello":
                        v12 = 1
                    }
                }
                for fi28 := 0; fi28 < 11; fi28++ {
                    var v29 uint64 = 0
                    v13 = (v13 << ((448) & 63))
                    p1 = v13
                    _ = fi28
                    _ = v29
                }
                _ = fi23
            }
        }
        if ("\n\t" == "a") {
            fmt.Println(v13, v12)
            switch "\n\t" {
            case "a":
                if false {
                    var v30 float32 = 0.0
                    _ = v30
                }
            case "hello":
                p1 = -733
            default:
                v31 := []int{880, 42}
                _ = v31
            }
            p1 -= v13
        } else {
            var v32 uint64 = 18446744073709551615
            for fi33 := 0; fi33 < 4; fi33++ {
                fmt.Println(p0)
                for fi34 := 0; fi34 < 19; fi34++ {
                    v35 := []int{665, fi34, fi33, (fi34 + fi33)}
                    v36 := GenS2{F0: (^(v12))}
                    fmt.Println(v12, p2, v13)
                    v35 = append(v35, int(v13))
                    if true {
                        if true {
                            _ = len(v35)
                            var v37 uint32 = 4294967295
                            _ = v37
                        }
                    }
                    var v38 int = fi33
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v38
                }
                switch fi33 {
                case -644:
                    p0 = p2
                default:
                    var v39 float32 = -2.25
                    _ = v39
                }
                _ = fi33
            }
            _ = v32
        }
    }
    for fi40 := 0; fi40 < 5; fi40++ {
        v41 := GenS3{F0: (^(v12))}
        _ = fi40
        _ = v41
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int = 0
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    var v6 int64 = -9223372036854775808
    var v42 float32 = -2.25
    var v43 bool = true
    fmt.Println(v6, v6)
    for fi44 := 0; fi44 < 4; fi44++ {
        var v45 int32 = -520
        _ = fi44
        _ = v45
    }
    v6 -= (-(v6))
    var v46 float32 = ((v42 + 0.1) + float32(0.1))
    var v47 int8 = -16
    for fi48 := 0; fi48 < 6; fi48++ {
        v49 := GenS4{F0: 0, F1: 1, F2: (genFn1(-1, (v6 ^ 9223372036854775807), -1) == 1000)}
        var v50 uint64 = 2
        fmt.Println(v6, v1)
        _ = fi48
        _ = v49
        _ = v50
    }
    switch (-(v1)) {
    case -275:
        v2[""] = v43
    }
    if (421 >= 18446744073709551615) {
        var v51 int64 = -439
        fmt.Println(v47, v47, v43)
        _ = v51
    }
    if v43 {
        var v52 int = v1
        if ((1e10 >= (v42 + (v42 / v46))) || (!v43)) {
            if (2 == 65535) {
                var v53 uint64 = 42
                _ = v53
            }
            if true {
                delete(v4, "")
                if (!v43) {
                    fmt.Println(v42, v43, v52)
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        for fi54 := 0; fi54 < 7; fi54++ {
                            var v55 int = 127
                            v42 += v46
                            v6 = (v6 ^ (v6 ^ v6))
                            _ = fi54
                            _ = v55
                        }
                    case "foo bar":
                        var v56 uint8 = 1
                        _ = v56
                    case "":
                        var v57 uint = 236
                        _ = v57
                    }
                    var v58 int64 = (v6 * v6)
                    _ = v58
                } else {
                    if true {
                        for fi59 := 0; fi59 < 20; fi59++ {
                            v60 := v3[(-(fi59))]
                            if v43 {
                                v61 := []int{v52, 32768, (v1 << ((998) & 63))}
                                fmt.Println(v47, v6)
                                _ = v61
                            }
                            v42 -= float32(1.5)
                            _ = fi59
                            _ = v60
                        }
                    }
                    v62 := v4["hello"]
                    var v63 float64 = 3.14159
                    _ = v62
                    _ = v63
                }
                var v64 int64 = (-(v6))
                _ = v64
            }
        }
        _ = v52
    } else {
        v5[("a" + "hello")] = (v43 && v43)
    }
    v2["a"] = (!(((611 >= 198) || true) || (65535 > 42)))
    if true {
        v2["a"] = v43
        var v65 int = -2147483648
        _ = v65
    } else {
        for fi66 := 0; fi66 < 12; fi66++ {
            v6 += (-(v6))
            _ = fi66
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
