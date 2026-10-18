package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint32 }
type GenS5 struct { F0 int64; F1 bool; F2 uint }
type GenS6 struct { F0 uint64; F1 string; F2 int32 }

func main() {
    var v1 int32 = -887
    var v2 uint = 2
    v3 := make(map[string]string)
    fmt.Println(v1, v2)
    if true {
        v4 := GenS1{F0: uint64(v2), F1: -25, F2: 1}
        _ = v4
    }
    var v5 uint8 = 223
    v6 := GenS1{F0: 42, F1: -68, F2: 0}
    _ = len(v3)
    v5 = 0
    if false {
        v3[("" + "\n\t")] = "foo bar"
    }
    var v7 string = "a"
    v2 += (v2 | 2)
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float64 = 1e10
        var v10 uint = 4294967295
        v11 := []int{128, (fi8 + fi8), (^(fi8))}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    switch v7 {
    case "hello":
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 float32 = 3.14159
            v14 := GenS2{F0: v7}
            _ = fi12
            _ = v13
            _ = v14
        }
    default:
        v15 := GenS3{F0: 4294967295, F1: 535, F2: 2}
        _ = v15
    }
    var v16 int64 = 42
    switch -2147483648 {
    case 9223372036854775807:
        if true {
            switch "foo bar" {
            case "hello":
                v17 := []int{-65, 32767}
                _ = v17
            case "":
                var v18 int8 = -1
                _ = v18
            default:
                switch ((v7 + "foo bar") + v7) {
                case "\n\t":
                    v7 = ((v7 + v7) + v7)
                case "hello":
                    v16 -= v16
                case "":
                    v19 := GenS3{F0: 42, F1: 65535, F2: 1}
                    _ = v19
                }
            }
            for fi20 := 0; fi20 < 9; fi20++ {
                for fi21 := 0; fi21 < 10; fi21++ {
                    v7 = v7
                    switch v7 {
                    case "foo bar":
                        v22 := []int{}
                        _ = v22
                    case "":
                        if (!true) {
                            v23 := GenS1{F0: 11, F1: -10, F2: 84}
                            v5 += v5
                            _ = v23
                        }
                    case "\n\t":
                        if (-673 <= 32767) {
                            v7 = "a"
                        }
                    default:
                        switch fi21 {
                        case -530:
                            switch 127 {
                            case 127:
                                switch "" {
                                case "foo bar":
                                    v2 = v2
                                case "a":
                                    fmt.Println(v16, fi20, v16)
                                case "hello":
                                    v24 := GenS4{F0: v1, F1: 2147483647}
                                    _ = v24
                                }
                            case 32768:
                                if (223 > v2) {
                                    if false {
                                        var v25 uint16 = 978
                                        if true {
                                            _ = len(v3)
                                            v2 = v2
                                            v26 := GenS5{F0: 804, F1: false, F2: 65535}
                                            v27 := []int{}
                                            _ = v26
                                            _ = v27
                                        }
                                        v28 := GenS4{F0: (v1 / ((v1) | 1)), F1: 545}
                                        _ = v25
                                        _ = v28
                                    } else {
                                        if (183 > 56) {
                                            if (true && ((v2 << (((v2 % ((v2) | 1))) & 63)) == v2)) {
                                                var v29 int64 = v16
                                                var v30 uint8 = 39
                                                _ = v29
                                                _ = v30
                                            }
                                        } else {
                                            var v31 float64 = 1.5
                                            var v32 uint64 = 1
                                            fmt.Println(fi20, v16, fi20)
                                            switch (fi21 % ((32768) | 1)) {
                                            case -1:
                                                var v33 bool = ((v32 & v32) == (v32 & (v32 ^ v32)))
                                                _ = v33
                                            default:
                                                delete(v3, (v7 + v7))
                                            }
                                            _ = v31
                                            _ = v32
                                        }
                                        v1 = (v1 >> ((623) & 63))
                                    }
                                }
                            default:
                                v16 += (v16 % (((v16 - -9223372036854775808)) | 1))
                            }
                        case -10:
                            var v34 uint = 2
                            _ = v34
                        default:
                            fmt.Println(v16)
                        }
                    }
                    var v35 int32 = int32(v2)
                    _ = fi21
                    _ = v35
                }
                v5 -= (v5 - 1)
                v36 := GenS1{F0: 1, F1: 0, F2: 0}
                _ = fi20
                _ = v36
            }
            var v37 int8 = 127
            _ = v37
        }
    default:
        if (false || false) {
            fmt.Println(v5, v1)
            var v38 int32 = (v1 ^ 42)
            _ = len(v3)
            _ = v38
        } else {
            if (!true) {
                v39 := []int{-420, 127, 127, -2147483648}
                for fi40 := 0; fi40 < 11; fi40++ {
                    if len(v39) >= 2 {
                        v39 = v39[1:len(v39)]
                    }
                    _ = len(v39)
                    fmt.Println(v7, v2)
                    _ = fi40
                }
                _ = v39
            } else {
                for fi41 := 0; fi41 < 6; fi41++ {
                    for fi42 := 0; fi42 < 19; fi42++ {
                        var v43 float64 = 1.5
                        var v44 float64 = ((((((v43 - v43) - ((v43 - float64(1.5)) * (v43 - -2.25))) * ((1.5 * (-2.25 * (v43 + float64((0.0 + float32(0.0)))))) - v43)) - v43) + v43) - v43)
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                    var v45 uint = 18446744073709551615
                    _ = fi41
                    _ = v45
                }
                v46 := GenS6{F0: 0, F1: v7, F2: (v1 - v1)}
                delete(v3, "")
                _ = v46
            }
            for fi47 := 0; fi47 < 1; fi47++ {
                v48 := GenS2{F0: v7}
                v49 := []int{fi47}
                v5 = 1
                _ = fi47
                _ = v48
                _ = v49
            }
        }
    }
    var v50 int16 = -591
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
