package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    v1 := []uint32{2}
    v2 := []uint32{522}
    var v3 float64 = (0.0 / 0.1)
    v4 := make(map[int]uint32)
    v5 := []string{}
    v6 := []string{"foo bar", "foo bar"}
    fmt.Println(v3)
    v3 += (1e10 * 0.1)
    if false {
        v7 := GenS1{F0: 42, F1: 2}
        v8 := GenS1{F0: 2, F1: 491}
        _ = v7
        _ = v8
    }
    switch 32767 {
    case 1:
        if (-128 <= 0) {
            var v9 int16 = 0
            _ = v9
        } else {
            var v10 int = 9223372036854775807
            delete(v4, v10)
            _ = v10
        }
    case 0:
        if true {
            v11 := GenS2{F0: -9223372036854775808, F1: 16, F2: 127}
            v12 := GenS3{F0: 189}
            _ = v11
            _ = v12
        } else {
            fmt.Println(v3)
        }
    }
    if false {
        _ = len(v4)
        var v13 uint = 18446744073709551615
        var v14 int8 = 70
        _ = v13
        _ = v14
    } else {
        v15 := GenS3{F0: 266}
        var v16 uint16 = 459
        _ = v15
        _ = v16
    }
    delete(v4, 32768)
    delete(v4, 2147483647)
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int16 = 42
        switch 4294967295 {
        case 988:
            v18 = 0
        case 32768:
            for fi19 := 0; fi19 < 5; fi19++ {
                _ = len(v4)
                v4[(fi17 ^ (fi19 ^ (-(fi17))))] = 310
                v18 = (-(v18))
                _ = fi19
            }
        default:
            var v20 uint8 = 255
            _ = v20
        }
        v21 := GenS3{F0: 2}
        _ = fi17
        _ = v18
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        v3 = (v3 / ((float64(1e10) + ((0.0 - float64(((1e10 / -2.25) / 1.0))) - -2.25)) / 1e10))
        var v23 uint = 0
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        if false {
            fmt.Println(v3)
            _ = len(v1)
            if false {
                for fi25 := 0; fi25 < 8; fi25++ {
                    switch ("a" + "\n\t") {
                    case "hello":
                        v3 = v3
                    case "\n\t":
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v3, fi25)
                        }
                    case "foo bar":
                        v26 := GenS3{F0: 0}
                        _ = v26
                    default:
                        switch -709 {
                        case 128:
                            if (!(127 <= 127)) {
                                if false {
                                    switch ("a" + "foo bar") {
                                    case "":
                                        var v27 uint16 = 554
                                        _ = v27
                                    case "a":
                                        var v28 uint32 = 501
                                        _ = v28
                                    default:
                                        v3 = v3
                                    }
                                    v3 = float64((1.5 + 0.0))
                                    _ = len(v4)
                                } else {
                                    v29 := GenS4{F0: 1, F1: 9}
                                    v3 = ((v3 / v3) * float64(1.5))
                                    _ = v29
                                }
                                v3 = 1.5
                            }
                        }
                    }
                    var v30 uint8 = 42
                    _ = fi25
                    _ = v30
                }
            }
        }
        v3 = (3.14159 - v3)
        _ = fi24
    }
    _ = len(v6)
    fmt.Println(v3)
    switch 9223372036854775807 {
    case 9223372036854775807:
        for fi31 := 0; fi31 < 15; fi31++ {
            v3 -= v3
            if (255 != 138) {
                var v32 int16 = 1
                if (!false) {
                    v33 := GenS2{F0: (fi31 >> ((76) & 63)), F1: v32, F2: 0}
                    v34 := GenS5{F0: v32, F1: (fi31 + (fi31 | fi31)), F2: 694}
                    _ = v33
                    _ = v34
                }
                for fi35 := 0; fi35 < 17; fi35++ {
                    _ = len(v4)
                    _ = fi35
                }
                _ = v32
            }
            switch fi31 {
            case 1:
                var v36 int32 = 42
                _ = v36
            case 2147483647:
                fmt.Println(v3)
            case 0:
                v3 += v3
            default:
                v37 := GenS5{F0: 252, F1: fi31, F2: 265}
                _ = v37
            }
            _ = fi31
        }
    case 4294967295:
        v38 := GenS1{F0: 42, F1: 90}
        _ = v38
    default:
        var v39 float64 = v3
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
