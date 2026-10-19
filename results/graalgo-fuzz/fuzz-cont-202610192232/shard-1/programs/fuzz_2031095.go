package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int32; F2 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 uint8 }
func genFn0(p0 int64, p1 int64, p2 int8) int8 {
    p1 += (p1 ^ (^(p1)))
    if true {
        p0 = (p0 ^ (p1 | 9223372036854775807))
        var v6 uint = 4294967295
        _ = v6
    } else {
        v7 := []int{32767, 2147483647}
        v7 = append(v7, 576)
        var v8 uint16 = 0
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{(-(fi9))}
        _ = fi9
        _ = v10
    }
    var v11 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -128
}
func genFn1(p0 uint32, p1 int16, p2 bool) int8 {
    switch (("foo bar" + "") + "hello") {
    case "hello":
        if p2 {
            if true {
                var v12 float64 = (1e10 * (0.0 / -2.25))
                _ = v12
            }
        }
    }
    switch "" {
    case "hello":
        var v13 uint32 = (^(p0))
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: -2147483648}
            for fi16 := 0; fi16 < 12; fi16++ {
                p0 += p0
                switch ("foo bar" + ("foo bar" + "")) {
                case "hello":
                    var v17 float64 = (0.1 / -2.25)
                    _ = v17
                case "\n\t":
                    p2 = true
                case "a":
                    switch "hello" {
                    case "":
                        var v18 float64 = (((3.14159 * float64(0.0)) + ((((float64((3.14159 * 1e10)) + 3.14159) * float64(3.14159)) * float64((((3.14159 * (float32(1e10) - float32(-2.25))) * 1e10) * 1e10))) / 1.0)) + (0.0 / 0.1))
                        _ = v18
                    case "\n\t":
                        v19 := GenS2{F0: 65535}
                        _ = v19
                    default:
                        fmt.Println(p0, p1)
                    }
                    if false {
                        p2 = p2
                        var v20 uint64 = 556
                        _ = v20
                    } else {
                        var v21 int32 = 0
                        _ = v21
                    }
                }
                p2 = ("foo bar" != "hello")
                _ = fi16
            }
            p2 = (0.0 != float32(0.0))
            _ = fi14
            _ = v15
        }
    default:
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 int64 = 42
            _ = fi22
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 100
}

func main() {
    var v1 uint = 4294967295
    var v2 uint64 = 2
    v3 := make(map[string]bool)
    var v4 int = -821
    v5 := make(map[int]int)
    v24 := []int{v4, v4}
    switch ("hello" + "foo bar") {
    case "a":
        if false {
            var v25 uint = v1
            _ = v25
        } else {
            fmt.Println(v4, v2)
            for fi26 := 0; fi26 < 13; fi26++ {
                if (v1 >= 795) {
                    var v27 uint64 = 1
                    _ = v27
                }
                switch (fi26 >> ((v1) & 63)) {
                case 32767:
                    switch 864 {
                    case 127:
                        fmt.Println(fi26)
                    default:
                        switch v4 {
                        case 2147483647:
                            switch 1 {
                            case -33:
                                v2 = (v2 | (v2 >> ((v1) & 63)))
                            case -2147483648:
                                v28 := GenS1{F0: -1}
                                _ = v28
                            }
                        case 127:
                            v29 := GenS3{F0: -226}
                            _ = v29
                        case -2147483648:
                            v30 := GenS1{F0: 0}
                            _ = v30
                        }
                    }
                case -2147483648:
                    fmt.Println(v2)
                default:
                    v31 := GenS2{F0: 2}
                    _ = v31
                }
                v1 += v1
                _ = fi26
            }
            if false {
                v32 := GenS3{F0: -9223372036854775808}
                _ = v32
            }
        }
    case "hello":
        switch "" {
        case "":
            var v33 string = "foo bar"
            _ = v33
        case "hello":
            v4 += (-(v4))
        case "a":
            for fi34 := 0; fi34 < 6; fi34++ {
                var v35 float64 = 3.14159
                switch (fi34 & -58) {
                case 127:
                    v3[""] = (1 != 221)
                case 0:
                    var v36 float64 = 1e10
                    _ = v36
                case 32767:
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                }
                var v37 int16 = -302
                _ = fi34
                _ = v35
                _ = v37
            }
        }
    case "":
        v38 := GenS2{F0: 1}
        _ = v38
    default:
        var v39 int = -1
        _ = v39
    }
    var v40 uint8 = 127
    var v41 float64 = 0.0
    v42 := GenS3{F0: 851}
    v41 = (v41 - float64(0.0))
    if len(v24) > 0 {
        v24[0] = (-(v4))
    }
    switch (v4 % (((v4 << (((v1 % ((v1) | 1))) & 63))) | 1)) {
    case 885:
        var v43 float32 = 1e10
        _ = v43
    case -1:
        fmt.Println(v4)
    }
    if (127 == 127) {
        fmt.Println(v40, v40)
        for fi44 := 0; fi44 < 3; fi44++ {
            if (!(uint8(v1) == (v40 + v40))) {
                v41 = 0.1
            } else {
                for fi45 := 0; fi45 < 16; fi45++ {
                    if true {
                        if (true || true) {
                            var v46 int32 = -1
                            _ = v46
                        }
                        v4 -= fi45
                        fmt.Println(v1)
                    } else {
                        if true {
                            v47 := GenS4{F0: v1, F1: v2}
                            _ = v47
                        } else {
                            v48 := GenS5{F0: (v40 - v40), F1: 42, F2: -358}
                            _ = v48
                        }
                    }
                    var v49 uint = v1
                    _ = fi45
                    _ = v49
                }
            }
            if (66 >= -128) {
                v50 := GenS6{F0: 2147483647}
                for fi51 := 0; fi51 < 16; fi51++ {
                    v2 -= (v2 ^ 1)
                    _ = fi51
                }
                var v52 bool = (-1 == 1)
                _ = v50
                _ = v52
            }
            if false {
                switch "hello" {
                case "hello":
                    v41 -= 1.5
                case "a":
                    var v53 bool = true
                    _ = v53
                case "\n\t":
                    fmt.Println(v1, fi44, fi44)
                default:
                    if len(v24) > 0 {
                        v24[0] = (fi44 | v4)
                    }
                }
            } else {
                if len(v24) > 0 {
                    v24[0] = (-(v4))
                }
            }
            _ = fi44
        }
    } else {
        fmt.Println(v4)
        var v54 float64 = (v41 + float64(0.1))
        var v55 bool = true
        _ = v54
        _ = v55
    }
    v56 := GenS7{F0: "\n\t", F1: 199}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v56
}
