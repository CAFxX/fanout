package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint; F1 uint8 }
func genFn0(p0 int64, p1 uint, p2 uint16) uint64 {
    p1 = 0
    var v6 int16 = int16(p0)
    switch "\n\t" {
    case "foo bar":
        p1 -= p1
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 138, F1: (fi7 & (fi7 ^ (-(fi7)))), F2: 2}
        for fi9 := 0; fi9 < 10; fi9++ {
            if true {
                var v10 float32 = 1.5
                _ = v10
            } else {
                var v11 int8 = 42
                if false {
                    if true {
                        var v12 uint16 = p2
                        if ((42 == 182) || false) {
                            var v13 uint16 = p2
                            v6 = int16(fi7)
                            _ = v13
                        }
                        _ = v12
                    } else {
                        v14 := []int{(fi9 >> (((^(p1))) & 63)), (fi7 >> ((uint(p2)) & 63))}
                        v15 := GenS1{F0: 219, F1: (fi9 - fi7), F2: 2}
                        _ = v14
                        _ = v15
                    }
                    p2 = (p2 >> ((694) & 63))
                }
                v16 := []int{(fi7 / (((fi7 | (fi9 % ((int(v6)) | 1)))) | 1)), fi7, -9223372036854775808, fi9}
                _ = v11
                _ = v16
            }
            var v17 int16 = v6
            _ = fi9
            _ = v17
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1(p0 int16) int {
    switch "foo bar" {
    case "":
        p0 += p0
    case "hello":
        switch -348 {
        case 941:
            for fi18 := 0; fi18 < 5; fi18++ {
                p0 = p0
                _ = fi18
            }
        }
    case "\n\t":
        for fi19 := 0; fi19 < 5; fi19++ {
            if false {
                p0 = (p0 | p0)
                var v20 uint16 = 2
                v21 := GenS2{F0: 0, F1: 4294967295, F2: (^(v20))}
                _ = v20
                _ = v21
            }
            if (4294967295 <= 42) {
                fmt.Println(p0)
                v22 := GenS3{F0: 57}
                fmt.Println(p0, p0)
                _ = v22
            }
            _ = fi19
        }
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        fmt.Println(p0)
        for fi24 := 0; fi24 < 10; fi24++ {
            p0 -= -32768
            v25 := GenS2{F0: 800, F1: 2, F2: 220}
            p0 -= p0
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v26 string = ("foo bar" + "a")
    v27 := GenS2{F0: 2147483647, F1: 570, F2: 1}
    if false {
        v28 := []int{2147483647, 0, 127}
        _ = v28
    } else {
        v26 = v26
        for fi29 := 0; fi29 < 14; fi29++ {
            v30 := GenS4{F0: -953, F1: fi29}
            _ = fi29
            _ = v30
        }
    }
    _ = p0
    _ = v26
    _ = v27
    return 128
}

func main() {
    var v1 uint = 4294967295
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]int)
    var v4 uint64 = 762
    v5 := make(map[uint32]bool)
    v31 := GenS2{F0: -617, F1: 890, F2: 298}
    _ = len(v5)
    switch 792 {
    case -1:
        _ = len(v3)
    case 1:
        var v32 int64 = 42
        _ = v32
    }
    fmt.Println(v2, v4, v1)
    for fi33 := 0; fi33 < 5; fi33++ {
        switch -2147483648 {
        case 9223372036854775807:
            for fi34 := 0; fi34 < 6; fi34++ {
                switch "" {
                case "foo bar":
                    v35 := GenS5{F0: 259, F1: 255}
                    _ = v35
                case "":
                    var v36 int8 = -1
                    _ = v36
                case "\n\t":
                    var v37 int16 = 1
                    _ = v37
                default:
                    v4 = v4
                }
                v2 = (v2 / (((v2 << (((v1 | 255)) & 63))) | 1))
                v1 -= 1
                if true {
                    switch (("foo bar" + "") + "hello") {
                    case "\n\t":
                        v38 := GenS1{F0: 1, F1: (fi33 << (((v1 & v1)) & 63)), F2: v4}
                        _ = v38
                    case "hello":
                        switch -180 {
                        case -404:
                            if (18446744073709551615 != 4294967295) {
                                var v39 uint8 = 2
                                _ = v39
                            } else {
                                v40 := GenS2{F0: -704, F1: 556, F2: 255}
                                _ = v40
                            }
                            var v41 uint64 = (v4 % (((v4 / ((922) | 1))) | 1))
                            switch fi34 {
                            case -596:
                                fmt.Println(fi33)
                            case 4294967295:
                                var v42 int8 = 42
                                _ = v42
                            }
                            _ = v41
                        case 32767:
                            v1 -= (v1 & (v1 + (v1 & (^(v1)))))
                        default:
                            var v43 uint16 = uint16(v4)
                            v2 += 9223372036854775807
                            _ = v43
                        }
                    case "foo bar":
                        var v44 float32 = (0.0 * 3.14159)
                        _ = v44
                    }
                    v4 = (v4 & (^(v4)))
                    if true {
                        v2 = (-(v2))
                        var v45 int16 = 122
                        if (-1 == 2147483647) {
                            v5[2147483647] = ((float32(1e10) + (float32(3.14159) - 1e10)) <= (1e10 - (3.14159 + 1.5)))
                        }
                        _ = v45
                    }
                }
                _ = fi34
            }
        case 32767:
            v2 -= v2
        }
        var v46 int32 = 0
        _ = fi33
        _ = v46
    }
    v47 := []int{}
    v2 = -9223372036854775808
    if false {
        v48 := GenS6{F0: v1, F1: 186}
        if false {
            var v49 uint = (^(v1))
            fmt.Println(v49, v2, v1)
            _ = len(v47)
            _ = v49
        } else {
            var v50 bool = (!(127 >= 0))
            _ = v50
        }
        var v51 uint16 = 1
        _ = v48
        _ = v51
    }
    var v52 uint = 65535
    var v53 uint64 = (v4 | 1)
    for fi54 := 0; fi54 < 10; fi54++ {
        switch ("a" + "") {
        case "\n\t":
            v3[""] = genFn1(-32768)
        case "hello":
            v55 := v3[""]
            _ = v55
        default:
            var v56 uint64 = (v53 & v53)
            _ = v56
        }
        _ = fi54
    }
    if (true || (false || (true || (32767 < 32767)))) {
        v53 += (^(v53))
    } else {
        if true {
            if true {
                var v57 uint32 = 2147483648
                for fi58 := 0; fi58 < 20; fi58++ {
                    for fi59 := 0; fi59 < 15; fi59++ {
                        if false {
                            var v60 uint64 = (v4 >> ((1) & 63))
                            v53 -= (v60 ^ v53)
                            var v61 uint = (v1 << (((v1 | v52)) & 63))
                            _ = v60
                            _ = v61
                        } else {
                            if false {
                                switch -9223372036854775808 {
                                case 127:
                                    v62 := GenS4{F0: -2147483648, F1: (^(fi59))}
                                    _ = v62
                                case 128:
                                    var v63 int16 = 611
                                    _ = v63
                                }
                                v64 := GenS5{F0: -82, F1: 1}
                                _ = v64
                            }
                        }
                        _ = fi59
                    }
                    var v65 int32 = 2147483647
                    _ = fi58
                    _ = v65
                }
                fmt.Println(v2, v53)
                _ = v57
            }
        }
        v1 -= v1
        if len(v47) > 0 {
            v47[0] = -9223372036854775808
        }
    }
    v4 += (v4 | (v53 >> (((v52 * (^(v1)))) & 63)))
    var v66 int = genFn1(1)
    v66 = genFn1(42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v47
    _ = v5
    _ = v52
    _ = v53
    _ = v66
}
