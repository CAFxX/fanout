package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{42, 425, -9223372036854775808, 0}
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 0, F1: 2147483647, F2: 42}
    if false {
        var v5 int64 = 996
        _ = v5
    } else {
        switch "\n\t" {
        case "hello":
            var v6 float64 = (3.14159 * ((((1.5 + float64(0.0)) * float64(((3.14159 + 1e10) * 1.5))) * 0.1) - 1e10))
            _ = v6
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v7 int8 = 100
    if (!false) {
        v8 := GenS2{F0: (-66 < 350)}
        var v9 bool = true
        switch 1 {
        case 567:
            switch -2147483648 {
            case 128:
                fmt.Println(v9, v9)
            case -561:
                delete(v3, 980)
            case -499:
                v7 -= v7
            }
        default:
            fmt.Println(v7, v7)
        }
        _ = v8
        _ = v9
    } else {
        v10 := GenS3{F0: 781, F1: (v7 - v7)}
        var v11 int32 = 969
        _ = v10
        _ = v11
    }
    var v12 uint64 = 989
    if (false || (42 == -940)) {
        if (2147483647 != 2147483648) {
            v1[42] = "a"
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v13 := GenS2{F0: true}
            _ = v13
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := GenS1{F0: 0, F1: 0, F2: -797}
            v7 += (v7 % ((127) | 1))
            _ = fi14
            _ = v15
        }
        fmt.Println(v12, v12)
    }
    var v16 float64 = 1e10
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            v12 -= (v12 % ((281) | 1))
            var v19 float32 = -2.25
            fmt.Println(v19, v19, fi17)
            _ = fi18
            _ = v19
        }
        v2 = append(v2, (-(fi17)))
        switch "\n\t" {
        case "foo bar":
            v7 -= (^(v7))
        case "":
            if false {
                v20 := GenS4{F0: 554, F1: 128, F2: (^(v7))}
                if true {
                    if ((!false) || (255 >= 1)) {
                        v21 := GenS4{F0: v12, F1: 0, F2: v7}
                        var v22 uint64 = (v12 / ((v12) | 1))
                        v22 = 42
                        _ = v21
                        _ = v22
                    }
                    var v23 float32 = 0.1
                    _ = v23
                } else {
                    var v24 float32 = 0.0
                    _ = len(v3)
                    _ = v24
                }
                v12 = v12
                _ = v20
            } else {
                if (v12 > (v12 - (v12 ^ (v12 / ((v12) | 1))))) {
                    fmt.Println(fi17, v16)
                    v2 = append(v2, (fi17 ^ fi17))
                    fmt.Println(v12)
                }
                for fi25 := 0; fi25 < 3; fi25++ {
                    switch ("a" + "\n\t") {
                    case "a":
                        v12 = (v12 & v12)
                    case "":
                        var v26 float64 = v16
                        _ = v26
                    case "hello":
                        v1[1] = (("\n\t" + "foo bar") + "foo bar")
                    default:
                        v12 -= 18446744073709551615
                    }
                    if true {
                        v1[1] = ("a" + ("\n\t" + ("a" + "a")))
                        if ((1e10 * (1e10 * 1.5)) > 0.0) {
                            if false {
                                v7 -= v7
                                if false {
                                    v16 -= -2.25
                                }
                                v27 := GenS3{F0: (v12 % (((v12 + v12)) | 1)), F1: (v7 | (v7 + (-(v7))))}
                                _ = v27
                            }
                            if false {
                                v7 -= v7
                                switch "foo bar" {
                                case "\n\t":
                                    delete(v1, 2147483648)
                                }
                            } else {
                                var v28 int64 = -9223372036854775808
                                v16 = -2.25
                                if false {
                                    if false {
                                        var v29 int = fi25
                                        _ = v29
                                    }
                                    v7 += v7
                                    v7 -= (v7 - 100)
                                    var v30 int8 = v7
                                    _ = v30
                                }
                                _ = v28
                            }
                            v2 = append(v2, (fi25 << ((65535) & 63)))
                        }
                    }
                    var v31 int8 = (v7 & (v7 >> ((0) & 63)))
                    _ = fi25
                    _ = v31
                }
            }
        case "hello":
            for fi32 := 0; fi32 < 11; fi32++ {
                var v33 float64 = v16
                var v34 uint8 = 14
                _ = len(v2)
                _ = fi32
                _ = v33
                _ = v34
            }
        }
        _ = fi17
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        v36 := GenS1{F0: 0, F1: -2147483648, F2: 9223372036854775807}
        _ = fi35
        _ = v36
    }
    fmt.Println(v16, v7)
    for fi37 := 0; fi37 < 19; fi37++ {
        if (626 >= 196) {
            fmt.Println(v16)
            var v38 uint32 = 2147483648
            var v39 int16 = -1
            _ = v38
            _ = v39
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            var v41 uint16 = 42
            _ = fi40
            _ = v41
        }
        _ = fi37
    }
    var v42 int64 = 107
    v16 -= (v16 + v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v7
}
