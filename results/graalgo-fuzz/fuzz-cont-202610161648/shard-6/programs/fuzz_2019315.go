package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int8; F2 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{42, 42, 647, -2147483648}
    var v3 bool = (!false)
    var v4 uint32 = 4294967295
    var v5 uint8 = 0
    if false {
        fmt.Println(v4, v5, v3)
    } else {
        fmt.Println(v4, v3)
        v6 := GenS1{F0: v4, F1: uint64(v5)}
        _ = v6
    }
    _ = len(v1)
    v5 -= v5
    switch "a" {
    case "":
        switch 9223372036854775807 {
        case 849:
            switch 42 {
            case 9223372036854775807:
                if (v3 && (!((v3 && v3) || (200 >= 0)))) {
                    _ = len(v2)
                    fmt.Println(v4)
                }
            case 128:
                for fi7 := 0; fi7 < 3; fi7++ {
                    for fi8 := 0; fi8 < 19; fi8++ {
                        v5 += (^(v5))
                        if v3 {
                            var v9 float32 = ((1.5 + 0.1) + 1e10)
                            delete(v1, 371)
                            _ = v9
                        } else {
                            v4 -= (v4 + (v4 & v4))
                            switch "a" {
                            case "foo bar":
                                var v10 int = -9223372036854775808
                                _ = v10
                            case "\n\t":
                                if ((!(v3 || (375 < -2147483648))) || (v3 || (1 == v5))) {
                                    v11 := GenS2{F0: (-(fi8))}
                                    v12 := GenS3{F0: 0, F1: 42}
                                    v13 := GenS2{F0: fi8}
                                    _ = v11
                                    _ = v12
                                    _ = v13
                                }
                            case "hello":
                                v3 = v3
                                var v14 float32 = (3.14159 - (3.14159 + (0.0 - float32((-2.25 / float64(((-2.25 * 1e10) / -2.25)))))))
                                _ = v14
                            }
                        }
                        v5 = (v5 << ((255) & 63))
                        _ = fi8
                    }
                    switch (fi7 * 127) {
                    case -77:
                        var v15 uint16 = 2
                        _ = v15
                    case 201:
                        var v16 float32 = 0.0
                        _ = v16
                    case 2147483647:
                        v17 := GenS1{F0: uint32(v5), F1: 0}
                        _ = v17
                    default:
                        if v3 {
                            fmt.Println(v3)
                        }
                    }
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v3 = (!(uint64(v5) == 4294967295))
                        if (42 <= 0) {
                            var v19 uint64 = uint64(v5)
                            v19 -= v19
                            _ = v19
                        } else {
                            v20 := GenS4{F0: 175, F1: 965}
                            fmt.Println(v3)
                            _ = v20
                        }
                        _ = fi18
                    }
                    _ = fi7
                }
            }
        case 553:
            switch 4294967295 {
            case 127:
                for fi21 := 0; fi21 < 8; fi21++ {
                    v3 = (342 <= (fi21 + (-(fi21))))
                    v5 = (v5 * v5)
                    _ = fi21
                }
            case -9223372036854775808:
                for fi22 := 0; fi22 < 2; fi22++ {
                    v23 := GenS4{F0: fi22, F1: 0}
                    _ = fi22
                    _ = v23
                }
            default:
                for fi24 := 0; fi24 < 7; fi24++ {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        if len(v2) > 0 {
                            v2[0] = (fi24 + fi24)
                        }
                        _ = fi25
                    }
                    var v26 string = "hello"
                    _ = fi24
                    _ = v26
                }
            }
        case -379:
            switch "a" {
            case "":
                switch "hello" {
                case "\n\t":
                    _ = len(v1)
                default:
                    v2 = append(v2, 0)
                }
            default:
                switch "a" {
                case "\n\t":
                    var v27 uint8 = (^(v5))
                    _ = v27
                case "a":
                    v5 += v5
                case "hello":
                    if (65535 < 255) {
                        v5 = (v5 & (v5 / ((158) | 1)))
                    }
                default:
                    switch 127 {
                    case -1:
                        v4 -= 65535
                    }
                }
            }
        default:
            var v28 int = 128
            _ = v28
        }
    case "a":
        fmt.Println(v5, v3, v3)
    default:
        v3 = v3
    }
    v29 := GenS4{F0: 182, F1: 266}
    var v30 uint8 = 0
    var v31 float32 = 3.14159
    v4 -= (v4 ^ v4)
    var v32 bool = (v3 || v3)
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 uint = 904
        _ = fi33
        _ = v34
    }
    v35 := GenS5{F0: 42}
    if false {
        for fi36 := 0; fi36 < 14; fi36++ {
            var v37 bool = (-881 != 42)
            _ = fi36
            _ = v37
        }
        fmt.Println(v31)
    } else {
        var v38 int = -19
        switch "hello" {
        case "hello":
            var v39 uint32 = (v4 + 564)
            _ = v39
        default:
            v40 := GenS6{F0: 255, F1: 112, F2: ""}
            _ = v40
        }
        _ = v38
    }
    if false {
        var v41 uint8 = (^(v30))
        var v42 uint = 2
        _ = v41
        _ = v42
    } else {
        v32 = (1 < -814)
        var v43 float64 = (-2.25 * float64(v31))
        _ = v43
    }
    switch 669 {
    case -469:
        v44 := v1[2147483647]
        _ = v44
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -993:
        v32 = false
    default:
        v32 = false
    }
    for fi45 := 0; fi45 < 19; fi45++ {
        v30 = 1
        _ = fi45
    }
    switch ((("foo bar" + "foo bar") + "a") + "a") {
    case "\n\t":
        var v46 uint8 = (v30 - (v30 + v5))
        _ = v46
    case "":
        v47 := GenS2{F0: 4294967295}
        _ = v47
    }
    v31 -= float32(float64(v31))
    if false {
        var v48 uint = 18446744073709551615
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
}
