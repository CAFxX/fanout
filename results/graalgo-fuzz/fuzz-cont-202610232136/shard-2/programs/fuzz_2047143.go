package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 string }
type GenS2 struct { F0 uint8; F1 int; F2 uint16 }
type GenS3 struct { F0 int; F1 bool; F2 int8 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int64; F1 uint8; F2 uint }
type GenS9 struct { F0 uint64 }
func genFn0(p0 int, p1 int32, p2 uint8) int {
    for fi7 := 0; fi7 < 2; fi7++ {
        switch "\n\t" {
        case "\n\t":
            var v8 string = ""
            _ = v8
        case "a":
            var v9 int64 = -1
            _ = v9
        case "hello":
            var v10 int = -9223372036854775808
            _ = v10
        }
        v11 := []int{-846, p0, (fi7 & (fi7 ^ (fi7 & 1)))}
        if true {
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := GenS1{F0: (^(p1)), F1: p2, F2: ("a" + ("foo bar" + ("\n\t" + "a")))}
                var v14 string = ""
                v15 := GenS1{F0: p1, F1: (p2 + 128), F2: "a"}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 string = "a"
            _ = v16
        }
        _ = fi7
        _ = v11
    }
    v17 := []int{32768, (p0 & (p0 + p0)), -812, -328}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return (p0 % (((p0 >> ((65535) & 63))) | 1))
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]uint32)
    var v3 bool = (!(((0 != 65535) && (-90 <= 0)) || (0 == 0)))
    var v4 int32 = 0
    v5 := make(map[string]uint32)
    var v6 int8 = 42
    v1 = 0.0
    v6 = (^(v6))
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := []int{fi18, (fi19 ^ genFn0(fi18, v4, 1)), fi18, (fi18 % ((fi18) | 1))}
            v4 += v4
            _ = fi19
            _ = v20
        }
        if (v3 || false) {
            v21 := GenS1{F0: (-(v4)), F1: 128, F2: (("a" + "foo bar") + (("hello" + "hello") + "\n\t"))}
            _ = v21
        } else {
            v22 := GenS1{F0: v4, F1: 42, F2: "foo bar"}
            v6 -= v6
            _ = v22
        }
        _ = fi18
    }
    v23 := v5["a"]
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS1{F0: 409, F1: 42, F2: ("a" + "a")}
        _ = fi24
        _ = v25
    }
    v26 := []int{2147483647, -1, 32767, -429}
    var v27 bool = (2 > 114)
    if (2147483647 <= (v4 >> ((1) & 63))) {
        v28 := GenS2{F0: 42, F1: -310, F2: 407}
        _ = v28
    } else {
        var v29 int16 = -32768
        v30 := GenS2{F0: 255, F1: -777, F2: 1}
        v3 = true
        _ = v29
        _ = v30
    }
    if v27 {
        v4 -= -1
        v31 := GenS1{F0: 0, F1: 42, F2: "hello"}
        v3 = v27
        _ = v31
    } else {
        if (-1 <= 0) {
            var v32 int32 = (v4 + (-(v4)))
            _ = v32
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            v6 -= v6
            switch genFn0(fi33, 2147483647, 128) {
            case -9223372036854775808:
                if v3 {
                    if (!(-32768 <= -140)) {
                        v34 := GenS3{F0: fi33, F1: v27, F2: (^(v6))}
                        switch fi33 {
                        case 699:
                            fmt.Println(v27)
                        default:
                            v35 := GenS1{F0: int32(v23), F1: 2, F2: ""}
                            _ = v35
                        }
                        _ = v34
                    } else {
                        for fi36 := 0; fi36 < 16; fi36++ {
                            var v37 uint32 = v23
                            v38 := GenS4{F0: 1, F1: 654}
                            v39 := GenS5{F0: -9223372036854775808, F1: (v4 | (v4 * (v4 - v4)))}
                            v27 = true
                            var v40 uint16 = 65535
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        fmt.Println(fi33, v3, v1)
                        for fi41 := 0; fi41 < 2; fi41++ {
                            if v3 {
                                v27 = v27
                            }
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(v4, v1)
                            case "a":
                                var v42 uint32 = v23
                                _ = v42
                            case "hello":
                                fmt.Println(v27, v23, v6)
                            }
                            _ = fi41
                        }
                    }
                    for fi43 := 0; fi43 < 2; fi43++ {
                        v6 -= -114
                        v44 := v5["a"]
                        _ = fi43
                        _ = v44
                    }
                    var v45 int32 = v4
                    _ = v45
                } else {
                    if v3 {
                        _ = len(v26)
                        switch (fi33 / (((fi33 ^ 32767)) | 1)) {
                        case 882:
                            fmt.Println(v27)
                        }
                        v27 = (1 > 65535)
                    }
                }
            case -2147483648:
                v27 = v3
            case -388:
                v46 := v5[""]
                _ = v46
            }
            for fi47 := 0; fi47 < 7; fi47++ {
                v4 = (v4 / ((v4) | 1))
                var v48 string = "hello"
                _ = fi47
                _ = v48
            }
            _ = fi33
        }
        fmt.Println(v1, v4, v6)
    }
    if false {
        v49 := v2[(v23 + (v23 ^ v23))]
        _ = v49
    }
    v50 := GenS4{F0: (v6 + v6), F1: 0}
    if (255 <= 2) {
        for fi51 := 0; fi51 < 10; fi51++ {
            if v3 {
                v27 = (!(v23 != 949))
                var v52 uint32 = v23
                if v27 {
                    switch "hello" {
                    case "":
                        fmt.Println(v52, v52, v23)
                    case "foo bar":
                        v26 = append(v26, genFn0(fi51, v4, 255))
                    case "hello":
                        v53 := GenS4{F0: (v6 >> ((42) & 63)), F1: 892}
                        _ = v53
                    default:
                        v54 := GenS3{F0: (-(fi51)), F1: true, F2: v6}
                        _ = v54
                    }
                    switch fi51 {
                    case 4294967295:
                        for fi55 := 0; fi55 < 14; fi55++ {
                            switch genFn0((^(fi51)), v4, 1) {
                            case 92:
                                v56 := v2[v52]
                                _ = v56
                            }
                            if v3 {
                                fmt.Println(fi55, v1, v6)
                                if v3 {
                                    var v57 int32 = -2147483648
                                    _ = v57
                                }
                            }
                            var v58 float64 = (float64(((float32(v1) * ((0.1 + 3.14159) + 0.1)) / 1.0)) / (1e10 * v1))
                            _ = fi55
                            _ = v58
                        }
                    case -1:
                        fmt.Println(fi51)
                    }
                    v59 := GenS6{F0: fi51}
                    _ = v59
                }
                _ = v52
            }
            fmt.Println(v27, v3)
            _ = fi51
        }
        var v60 int = 2147483647
        _ = v60
    }
    for fi61 := 0; fi61 < 19; fi61++ {
        var v62 float32 = 3.14159
        for fi63 := 0; fi63 < 18; fi63++ {
            fmt.Println(fi63, v6)
            _ = fi63
        }
        for fi64 := 0; fi64 < 15; fi64++ {
            fmt.Println(v6, v3)
            switch "foo bar" {
            case "hello":
                if (!v27) {
                    v3 = true
                }
            case "\n\t":
                fmt.Println(v6, v23, v23)
            case "a":
                if true {
                    v23 -= v23
                    var v65 uint32 = (v23 % ((663) | 1))
                    v6 += (^(v6))
                    _ = v65
                }
            default:
                var v66 uint64 = 0
                _ = v66
            }
            switch "a" {
            case "a":
                var v67 uint64 = 42
                _ = v67
            case "":
                fmt.Println(v62)
            case "hello":
                if false {
                    if v27 {
                        var v68 string = "hello"
                        v4 = (^(v4))
                        _ = v68
                    } else {
                        if (-9223372036854775808 > -9223372036854775808) {
                            switch "" {
                            case "":
                                _ = len(v26)
                            default:
                                _ = len(v26)
                            }
                            var v69 bool = (740 >= 813)
                            var v70 uint32 = (v23 ^ 95)
                            _ = v69
                            _ = v70
                        }
                        if v3 {
                            switch "foo bar" {
                            case "hello":
                                _ = len(v2)
                                fmt.Println(v4)
                            case "":
                                fmt.Println(v1, v6)
                            case "a":
                                if (!(4294967295 != 689)) {
                                    v71 := GenS7{F0: 18446744073709551615}
                                    _ = v71
                                } else {
                                    var v72 int8 = v6
                                    var v73 float64 = 0.0
                                    _ = v72
                                    _ = v73
                                }
                            }
                            switch (fi64 ^ fi64) {
                            case 0:
                                if true {
                                    v74 := GenS4{F0: 3, F1: 814}
                                    if v3 {
                                        v75 := GenS8{F0: int64(v23), F1: 127, F2: 2}
                                        _ = v75
                                    }
                                    switch ("\n\t" + "foo bar") {
                                    case "\n\t":
                                        v76 := GenS9{F0: 1}
                                        _ = v76
                                    }
                                    _ = v74
                                }
                            case 32767:
                                if (2 != 2) {
                                    v23 -= (v23 % ((v23) | 1))
                                    switch (fi61 | int(v6)) {
                                    case 0:
                                        switch -2147483648 {
                                        case 42:
                                            fmt.Println(v62, v3)
                                        case -2147483648:
                                            var v77 uint = 42
                                            _ = v77
                                        }
                                    }
                                } else {
                                    if true {
                                        v78 := GenS2{F0: 255, F1: (fi61 | fi64), F2: 42}
                                        _ = len(v2)
                                        _ = v78
                                    }
                                }
                            }
                            switch "hello" {
                            case "\n\t":
                                var v79 uint64 = 18446744073709551615
                                v4 += (v4 >> ((18446744073709551615) & 63))
                                v27 = (false || ((v4 - v4) >= -2147483648))
                                _ = v79
                            }
                        }
                        fmt.Println(v27)
                    }
                }
            }
            _ = fi64
        }
        _ = fi61
        _ = v62
    }
    v23 += v23
    fmt.Println(v6, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
