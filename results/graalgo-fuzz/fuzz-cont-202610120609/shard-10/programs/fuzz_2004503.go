package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint64; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 uint64; F1 uint32; F2 int64 }

func main() {
    v1 := []uint32{}
    var v2 uint64 = 776
    var v3 int64 = -750
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 int8 = -1
        v6 := GenS1{F0: 2, F1: 65535, F2: 32767}
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 bool = true
            v5 -= (v5 | (-(v5)))
            _ = fi7
            _ = v8
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    switch ("a" + "") {
    case "\n\t":
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS2{F0: 0}
            v11 := GenS2{F0: 0}
            _ = fi9
            _ = v10
            _ = v11
        }
    case "foo bar":
        v3 -= (v3 - v3)
    }
    var v12 uint16 = 1
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := GenS2{F0: 54}
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 uint8 = 128
            var v17 uint16 = (v12 << ((18446744073709551615) & 63))
            if true {
                v18 := GenS3{F0: fi13}
                _ = v18
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        v19 := GenS3{F0: (fi13 >> ((255) & 63))}
        _ = fi13
        _ = v14
        _ = v19
    }
    var v20 uint32 = 2147483648
    var v21 float32 = float32(float64(float32(0.1)))
    for fi22 := 0; fi22 < 1; fi22++ {
        switch ("" + ("\n\t" + "hello")) {
        case "foo bar":
            v3 = v3
        case "":
            if false {
                if true {
                    v23 := GenS4{F0: fi22, F1: (v2 & (v2 >> ((396) & 63))), F2: fi22}
                    var v24 int32 = -1
                    _ = v23
                    _ = v24
                }
                fmt.Println(v12, v2, fi22)
            }
        case "a":
            var v25 uint64 = (v2 / ((97) | 1))
            _ = v25
        }
        v12 = (^(v12))
        _ = fi22
    }
    fmt.Println(v2, v21)
    v12 -= (^(v12))
    if (v20 == v20) {
        var v26 int64 = (v3 / ((42) | 1))
        _ = v26
    } else {
        v27 := GenS3{F0: 42}
        var v28 float32 = (((1.5 / (0.1 + float32((1e10 + 0.1)))) * float32((float64(-2.25) - 0.0))) * ((v21 * (v21 + v21)) * v21))
        if ((!true) || false) {
            switch "hello" {
            case "":
                v2 = (v2 << ((65535) & 63))
            case "hello":
                var v29 int8 = 1
                _ = v29
            case "foo bar":
                fmt.Println(v3, v20, v2)
            default:
                var v30 uint8 = 42
                _ = v30
            }
        }
        _ = v27
        _ = v28
    }
    if (!false) {
        for fi31 := 0; fi31 < 13; fi31++ {
            v20 -= 198
            switch -1 {
            case -9223372036854775808:
                if true {
                    switch "\n\t" {
                    case "":
                        var v32 int32 = 720
                        _ = v32
                    case "\n\t":
                        v33 := GenS5{F0: 4294967295}
                        _ = v33
                    case "hello":
                        v3 = (^(v3))
                    default:
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                    }
                    var v34 uint = 2
                    _ = v34
                } else {
                    var v35 uint16 = v12
                    for fi36 := 0; fi36 < 10; fi36++ {
                        var v37 float32 = v21
                        if (v3 == (v3 ^ 1)) {
                            fmt.Println(fi31, fi31)
                            if true {
                                v38 := GenS1{F0: 600, F1: (^(v35)), F2: 118}
                                if (-128 > 0) {
                                    _ = len(v1)
                                    switch 127 {
                                    case 928:
                                        fmt.Println(v3)
                                    default:
                                        v3 = (v3 << ((4294967295) & 63))
                                    }
                                }
                                var v39 uint = 395
                                v40 := GenS4{F0: (fi31 + (-(fi31))), F1: (v2 * v2), F2: 127}
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            v41 := GenS1{F0: v2, F1: 103, F2: 1000}
                            _ = v41
                        } else {
                            fmt.Println(v37, v37, v21)
                            var v42 int64 = v3
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v43 int8 = 100
                            var v44 float32 = v37
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
            case 954:
                for fi45 := 0; fi45 < 20; fi45++ {
                    v12 = (v12 | v12)
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        v46 := GenS1{F0: 4294967295, F1: 989, F2: 0}
                        _ = v46
                    case "":
                        if true {
                            var v47 float32 = ((float32((((-2.25 / -2.25) / (0.1 - (((-2.25 / 1.5) + 0.1) / (((-2.25 + 0.0) * 0.0) + (0.1 - 0.0))))) / 1e10)) / v21) / 1.5)
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                            v21 = v21
                            _ = v47
                        }
                    case "hello":
                        switch "" {
                        case "foo bar":
                            v48 := GenS4{F0: (fi45 | fi45), F1: (v2 * (^(v2))), F2: (fi45 + (fi45 ^ fi31))}
                            _ = v48
                        }
                    default:
                        v21 -= ((float32(1e10) - ((((1e10 * (v21 * v21)) / ((0.0 / ((3.14159 / (3.14159 / v21)) - ((-2.25 / 1e10) - 1.5))) * ((0.1 + v21) / (v21 / (v21 + 1.5))))) + v21) * v21)) + v21)
                    }
                    _ = fi45
                }
            case 42:
                fmt.Println(v12)
            default:
                v49 := GenS4{F0: fi31, F1: (v2 / (((v2 - v2)) | 1)), F2: fi31}
                _ = v49
            }
            _ = fi31
        }
        for fi50 := 0; fi50 < 17; fi50++ {
            if false {
                v1 = append(v1, (v20 ^ 2))
                var v51 float32 = ((v21 / ((float32(1.5) - v21) + 0.1)) / v21)
                _ = v51
            }
            v21 -= v21
            if true {
                if len(v1) > 0 {
                    v1[0] = uint32(v2)
                }
            } else {
                v52 := GenS6{F0: 817, F1: -1}
                _ = v52
            }
            _ = fi50
        }
    } else {
        var v53 uint8 = 42
        var v54 uint8 = (v53 << ((189) & 63))
        fmt.Println(v21, v54, v20)
        _ = v53
        _ = v54
    }
    switch -2147483648 {
    case 4294967295:
        var v55 int64 = v3
        _ = v55
    default:
        fmt.Println(v3)
    }
    v3 = -1
    for fi56 := 0; fi56 < 10; fi56++ {
        switch "hello" {
        case "":
            var v57 uint16 = (v12 ^ v12)
            _ = v57
        default:
            var v58 uint32 = v20
            _ = v58
        }
        _ = fi56
    }
    v2 = v2
    if true {
        v59 := GenS7{F0: (v2 + 272), F1: 4294967295, F2: v3}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v59
    }
    v60 := GenS3{F0: 128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v60
}
