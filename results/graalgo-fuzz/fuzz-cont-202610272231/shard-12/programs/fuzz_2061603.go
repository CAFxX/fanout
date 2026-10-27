package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 string; F1 int8; F2 int64 }
type GenS6 struct { F0 int64; F1 bool }

func main() {
    var v1 int8 = 127
    var v2 string = "foo bar"
    var v3 uint64 = 480
    switch v2 {
    case "":
        for fi4 := 0; fi4 < 12; fi4++ {
            for fi5 := 0; fi5 < 5; fi5++ {
                var v6 uint16 = 32767
                v6 = uint16(v3)
                v7 := GenS1{F0: 135, F1: (fi5 ^ (fi4 ^ (fi5 >> ((65535) & 63))))}
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = fi4
        }
    case "hello":
        v1 = (^(v1))
    case "\n\t":
        for fi8 := 0; fi8 < 12; fi8++ {
            if false {
                if false {
                    fmt.Println(v3, fi8)
                    switch ((v2 + "foo bar") + (v2 + v2)) {
                    case "":
                        if true {
                            v2 = v2
                            v3 += v3
                        }
                    case "a":
                        v2 = "a"
                    case "\n\t":
                        fmt.Println(v2, v2)
                    }
                } else {
                    for fi9 := 0; fi9 < 13; fi9++ {
                        v3 += v3
                        v3 = v3
                        v1 += v1
                        _ = fi9
                    }
                    v10 := GenS2{F0: 100, F1: 326}
                    _ = v10
                }
                if false {
                    v11 := []int{128, (fi8 ^ (fi8 >> ((1) & 63))), (fi8 >> ((65535) & 63)), fi8}
                    _ = v11
                } else {
                    fmt.Println(fi8, v3)
                    fmt.Println(fi8)
                    v12 := []int{}
                    _ = v12
                }
                for fi13 := 0; fi13 < 1; fi13++ {
                    if (false && (!(!false))) {
                        var v14 uint16 = 597
                        fmt.Println(v14)
                        v3 = (v3 | v3)
                        v1 = (v1 & (v1 / ((v1) | 1)))
                        v14 = (v14 / ((v14) | 1))
                        _ = v14
                    } else {
                        var v15 int16 = 1
                        var v16 uint32 = 345
                        if true {
                            v17 := GenS1{F0: 127, F1: -533}
                            _ = v17
                        } else {
                            fmt.Println(v15, fi13, v16)
                        }
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(fi8, v3, fi13)
                    v18 := GenS3{F0: 2147483647, F1: v3}
                    _ = fi13
                    _ = v18
                }
            }
            v19 := []int{fi8, fi8, (-(fi8)), 977}
            _ = fi8
            _ = v19
        }
    default:
        v20 := GenS2{F0: 1, F1: 2}
        _ = v20
    }
    v21 := []int{-856, 9223372036854775807}
    v1 = 1
    var v22 float32 = float32(0.1)
    for fi23 := 0; fi23 < 7; fi23++ {
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 uint8 = 42
            _ = fi24
            _ = v25
        }
        fmt.Println(fi23, fi23)
        _ = fi23
    }
    v26 := GenS2{F0: 42, F1: 2147483647}
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    var v27 bool = true
    if (42 != 0) {
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := GenS4{F0: 42, F1: 1000}
            switch "a" {
            case "hello":
                v30 := GenS1{F0: 126, F1: 128}
                _ = v30
            }
            _ = fi28
            _ = v29
        }
    }
    fmt.Println(v3, v1)
    v31 := GenS1{F0: 42, F1: 931}
    fmt.Println(v27, v3, v3)
    var v32 int64 = -9223372036854775808
    switch v2 {
    case "foo bar":
        if (!(!v27)) {
            if v27 {
                v27 = (0 < 125)
                for fi33 := 0; fi33 < 11; fi33++ {
                    if true {
                        switch (fi33 ^ -709) {
                        case 128:
                            v2 = ("" + ((v2 + v2) + ((v2 + (v2 + (("\n\t" + v2) + ("" + ("" + v2))))) + "\n\t")))
                        case 0:
                            v34 := GenS1{F0: 111, F1: 1}
                            _ = v34
                        default:
                            var v35 uint32 = 2
                            _ = v35
                        }
                    } else {
                        for fi36 := 0; fi36 < 19; fi36++ {
                            v37 := GenS5{F0: v2, F1: (v1 | 127), F2: (v32 / ((0) | 1))}
                            switch ((v2 + (v2 + v2)) + v2) {
                            case "hello":
                                if true {
                                    v38 := GenS4{F0: 464, F1: 1000}
                                    fmt.Println(v3, v27, fi36)
                                    _ = v38
                                }
                            }
                            v39 := GenS1{F0: 255, F1: 9223372036854775807}
                            if v27 {
                                var v40 int32 = int32(v1)
                                _ = v40
                            }
                            v41 := GenS1{F0: 255, F1: (fi36 | fi36)}
                            v42 := GenS2{F0: 1, F1: 746}
                            _ = fi36
                            _ = v37
                            _ = v39
                            _ = v41
                            _ = v42
                        }
                    }
                    switch "\n\t" {
                    case "foo bar":
                        var v43 int64 = (-(v32))
                        _ = v43
                    case "hello":
                        if (uint(v3) >= 4294967295) {
                            if (-1 >= 42) {
                                fmt.Println(v27, v2, v3)
                            } else {
                                fmt.Println(fi33, v32)
                            }
                            v22 = (float32(0.1) + (1e10 + 1e10))
                            switch "foo bar" {
                            case "foo bar":
                                v21 = append(v21, int(v32))
                            }
                        } else {
                            for fi44 := 0; fi44 < 15; fi44++ {
                                v32 += v32
                                v45 := GenS2{F0: 42, F1: 65535}
                                if v27 {
                                    var v46 int32 = 1
                                    v22 = (float32((0.1 - 0.0)) - v22)
                                    _ = v46
                                }
                                _ = fi44
                                _ = v45
                            }
                        }
                    case "":
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    }
                    var v47 int64 = v32
                    _ = fi33
                    _ = v47
                }
                var v48 string = v2
                _ = v48
            } else {
                for fi49 := 0; fi49 < 11; fi49++ {
                    _ = len(v21)
                    if (fi49 < -686) {
                        if (false || v27) {
                            for fi50 := 0; fi50 < 6; fi50++ {
                                v21 = append(v21, -104)
                                _ = fi50
                            }
                            if len(v21) > 0 {
                                v21[0] = -1
                            }
                            v2 = (((((v2 + (v2 + v2)) + v2) + v2) + ("a" + v2)) + (((v2 + "hello") + v2) + "hello"))
                        }
                        var v51 int16 = -893
                        v52 := GenS5{F0: v2, F1: v1, F2: 9223372036854775807}
                        _ = v51
                        _ = v52
                    }
                    _ = fi49
                }
                var v53 int32 = -321
                _ = v53
            }
        }
    case "":
        for fi54 := 0; fi54 < 19; fi54++ {
            v3 = v3
            _ = fi54
        }
    case "hello":
        _ = len(v21)
    }
    v32 -= 1
    for fi55 := 0; fi55 < 6; fi55++ {
        v56 := GenS6{F0: v32, F1: (!false)}
        switch fi55 {
        case 9223372036854775807:
            v22 = 1e10
        default:
            var v57 uint16 = 255
            _ = v57
        }
        for fi58 := 0; fi58 < 6; fi58++ {
            v1 += v1
            v1 -= v1
            _ = len(v21)
            _ = fi58
        }
        _ = fi55
        _ = v56
    }
    v59 := GenS1{F0: 42, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v59
}
