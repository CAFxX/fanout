package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string }
func genFn0(p0 uint) uint16 {
    p0 += 2
    v6 := GenS1{F0: 42, F1: (p0 << ((42) & 63)), F2: 42}
    _ = p0
    _ = v6
    return 1
}
func genFn1() int {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int8 = -1
        switch (fi7 ^ (-(fi7))) {
        case 4294967295:
            v8 += (v8 << ((0) & 63))
        }
        for fi9 := 0; fi9 < 13; fi9++ {
            switch "a" {
            case "\n\t":
                fmt.Println(fi7, fi9, fi9)
            case "":
                switch "hello" {
                case "\n\t":
                    if false {
                        if false {
                            fmt.Println(v8)
                            if (127 >= 212) {
                                if ((v8 ^ (v8 | v8)) != v8) {
                                    var v10 uint8 = 255
                                    v11 := GenS1{F0: v8, F1: 856, F2: -458}
                                    v10 = (v10 ^ 0)
                                    _ = v10
                                    _ = v11
                                }
                            } else {
                                v12 := GenS1{F0: v8, F1: 18446744073709551615, F2: -83}
                                _ = v12
                            }
                        } else {
                            if false {
                                var v13 uint8 = 255
                                v14 := GenS2{F0: 2147483647, F1: 18446744073709551615}
                                v13 = (^(v13))
                                _ = v13
                                _ = v14
                            } else {
                                v15 := GenS1{F0: -101, F1: 0, F2: 42}
                                var v16 int8 = v8
                                v16 = (v8 | (-(v8)))
                                if false {
                                    v17 := GenS1{F0: (v8 ^ -53), F1: 845, F2: -32768}
                                    if true {
                                        var v18 float32 = 0.1
                                        _ = v18
                                    }
                                    v19 := []int{(fi9 % ((fi9) | 1))}
                                    if false {
                                        if len(v19) > 0 {
                                            v19[0] = (fi7 - fi7)
                                        }
                                    } else {
                                        fmt.Println(fi7)
                                        v16 -= 100
                                    }
                                    fmt.Println(fi9, fi7, fi7)
                                    _ = v17
                                    _ = v19
                                } else {
                                    switch (fi9 + (fi7 % (((fi9 >> ((658) & 63))) | 1))) {
                                    case 1:
                                        fmt.Println(v16)
                                    case 90:
                                        v20 := GenS1{F0: v8, F1: 65535, F2: 396}
                                        _ = v20
                                    case -9223372036854775808:
                                        v8 = -1
                                    }
                                }
                                v8 -= v8
                                _ = v15
                                _ = v16
                            }
                            v8 += (v8 * (-(v8)))
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                var v21 float64 = ((-2.25 * 0.1) / 1.0)
                                _ = v21
                            }
                            v22 := GenS3{F0: "a", F1: 100}
                            switch ("hello" + ("hello" + "\n\t")) {
                            case "":
                                var v23 int32 = 0
                                _ = v23
                            case "hello":
                                var v24 int64 = 938
                                _ = v24
                            default:
                                v8 += 100
                            }
                            v25 := []int{(fi7 & fi9)}
                            v8 = v8
                            _ = v22
                            _ = v25
                        }
                    } else {
                        v26 := GenS4{F0: (-(v8))}
                        _ = v26
                    }
                case "foo bar":
                    var v27 uint = 18446744073709551615
                    _ = v27
                }
                v8 = v8
            default:
                var v28 uint8 = 2
                _ = v28
            }
            var v29 uint = 18446744073709551615
            if ((!(!(genFn0(v29) >= 649))) && (!true)) {
                v30 := GenS5{F0: ("" + ""), F1: 131, F2: -32768}
                var v31 string = "foo bar"
                _ = v30
                _ = v31
            }
            _ = fi9
            _ = v29
        }
        _ = fi7
        _ = v8
    }
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 uint8 = 255
        _ = fi32
        _ = v33
    }
    v34 := []int{32768}
    _ = v34
    return 32768
}

func main() {
    var v1 int = 4294967295
    var v2 uint = 2
    v3 := []int{v1, v1, v1}
    var v4 uint32 = 2
    var v5 float32 = ((float32(0.0) / -2.25) - -2.25)
    fmt.Println(v4, v4, v5)
    fmt.Println(v1, v4, v2)
    if true {
        if ("a" < "hello") {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v5 = v5
        _ = len(v3)
    } else {
        for fi35 := 0; fi35 < 20; fi35++ {
            var v36 int = (fi35 - v1)
            var v37 uint8 = 210
            var v38 uint32 = (v4 >> ((v2) & 63))
            _ = fi35
            _ = v36
            _ = v37
            _ = v38
        }
    }
    v39 := GenS3{F0: "hello", F1: -97}
    v40 := GenS1{F0: -1, F1: 65535, F2: -949}
    var v41 uint32 = (v4 % (((v4 ^ 65535)) | 1))
    var v42 uint64 = 402
    if (true || ((!(1 <= 1)) || ((v5 * (v5 + float32(-2.25))) <= v5))) {
        switch ("foo bar" + "foo bar") {
        case "a":
            v43 := GenS3{F0: "hello", F1: -57}
            _ = v43
        default:
            if (0.0 != 1e10) {
                v4 += (v41 * 886)
                v4 += 4294967295
                v41 -= (v4 | 940)
            } else {
                fmt.Println(v42)
            }
        }
        fmt.Println(v4, v41, v2)
    } else {
        if (v5 == ((1.5 * (v5 + 1.5)) + float32((-2.25 - 0.0)))) {
            v4 -= (v41 << (((v2 ^ v2)) & 63))
            v41 = 42
        } else {
            for fi44 := 0; fi44 < 8; fi44++ {
                if (0.0 < v5) {
                    switch "\n\t" {
                    case "foo bar":
                        for fi45 := 0; fi45 < 3; fi45++ {
                            v46 := GenS4{F0: 48}
                            v3 = append(v3, (fi45 << (((v2 * (v2 / ((v2) | 1)))) & 63)))
                            _ = fi45
                            _ = v46
                        }
                    case "hello":
                        fmt.Println(v4, v42)
                    case "a":
                        fmt.Println(v1)
                    }
                    var v47 uint64 = 315
                    var v48 bool = ((v47 * v47) < 2)
                    _ = v47
                    _ = v48
                } else {
                    var v49 uint64 = (v42 & 1)
                    if true {
                        if (0.0 != (3.14159 / v5)) {
                            v50 := GenS6{F0: v4}
                            for fi51 := 0; fi51 < 20; fi51++ {
                                v52 := GenS7{F0: "foo bar"}
                                _ = fi51
                                _ = v52
                            }
                            _ = v50
                        } else {
                            v41 = 1
                            fmt.Println(v1, v41, v42)
                            var v53 float64 = ((0.0 + 0.1) * (((0.0 * 1e10) + (0.1 / 1e10)) + -2.25))
                            _ = v53
                        }
                        _ = len(v3)
                    }
                    v2 = v2
                    _ = v49
                }
                v5 += v5
                for fi54 := 0; fi54 < 20; fi54++ {
                    v55 := GenS7{F0: ""}
                    _ = fi54
                    _ = v55
                }
                _ = fi44
            }
            var v56 int16 = 1
            v2 = (v2 | uint(v4))
            _ = v56
        }
        switch v1 {
        case 32767:
            if false {
                if (false && true) {
                    v3 = append(v3, -100)
                }
                var v57 bool = (!true)
                var v58 uint16 = 16
                _ = v57
                _ = v58
            } else {
                for fi59 := 0; fi59 < 20; fi59++ {
                    v5 = ((v5 + 0.0) + 0.0)
                    _ = fi59
                }
                v5 = float32(-2.25)
                if len(v3) > 0 {
                    v3[0] = (^(v1))
                }
            }
        default:
            var v60 int8 = 79
            _ = v60
        }
        var v61 float32 = -2.25
        _ = v61
    }
    if len(v3) > 0 {
        v3[0] = genFn1()
    }
    v1 -= (^(v1))
    v5 = float32(3.14159)
    for fi62 := 0; fi62 < 16; fi62++ {
        switch 127 {
        case 32767:
            if false {
                if len(v3) > 0 {
                    v3[0] = genFn1()
                }
                fmt.Println(v4, v1, v2)
                v2 = 0
            }
        case 32768:
            fmt.Println(v2)
        case 602:
            var v63 uint = (v2 << (((v2 | (v2 ^ (v2 % (((v2 ^ uint(v42))) | 1))))) & 63))
            _ = v63
        default:
            for fi64 := 0; fi64 < 3; fi64++ {
                var v65 int16 = 1000
                _ = fi64
                _ = v65
            }
        }
        _ = fi62
    }
    v1 += (v1 - genFn1())
    switch "" {
    case "hello":
        var v66 uint8 = 17
        _ = v66
    case "\n\t":
        switch (v1 >> ((1) & 63)) {
        case -9223372036854775808:
            if true {
                v42 = v42
            }
        }
    case "foo bar":
        var v67 uint16 = 247
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
