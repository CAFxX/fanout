package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int16 }

func main() {
    var v1 float64 = (1e10 * float64(3.14159))
    var v2 float64 = v1
    var v3 int64 = -935
    var v4 int = -9223372036854775808
    v1 = 0.0
    switch 1 {
    case 960:
        var v5 uint16 = 0
        _ = v5
    case 2147483647:
        var v6 uint64 = 42
        _ = v6
    default:
        fmt.Println(v1, v4)
    }
    v1 += v1
    if true {
        if true {
            v7 := []int{(v4 << ((191) & 63))}
            _ = v7
        } else {
            v8 := GenS1{F0: -2147483648, F1: 65535}
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := []int{(v4 + 127)}
                if true {
                    v10 = append(v10, 1)
                    fmt.Println(v4, v2)
                }
                _ = fi9
                _ = v10
            }
            v11 := []int{32767, v4, 0, (v4 | (-(v4)))}
            _ = v11
            _ = v8
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            v3 = 9223372036854775807
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS2{F0: (("hello" + ("a" + "")) + ("a" + ("hello" + ""))), F1: 609}
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    } else {
        switch v4 {
        case -935:
            if ((!(!true)) || true) {
                var v15 uint16 = 32767
                _ = v15
            } else {
                switch v4 {
                case -9223372036854775808:
                    for fi16 := 0; fi16 < 5; fi16++ {
                        switch 127 {
                        case 181:
                            for fi17 := 0; fi17 < 19; fi17++ {
                                v18 := []int{}
                                v3 = 9223372036854775807
                                var v19 float64 = ((float64(1e10) - (v2 - v2)) / v2)
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        case -607:
                            if (1000 >= -1) {
                                v20 := []int{}
                                _ = v20
                            }
                        case -2147483648:
                            switch "hello" {
                            case "\n\t":
                                var v21 uint64 = 42
                                _ = v21
                            case "a":
                                if (!false) {
                                    for fi22 := 0; fi22 < 4; fi22++ {
                                        v3 = v3
                                        fmt.Println(v4, fi16, v3)
                                        _ = fi22
                                    }
                                    fmt.Println(v1, fi16)
                                }
                            default:
                                if false {
                                    v23 := []int{v4, (-(fi16)), v4}
                                    if (32767 != 255) {
                                        _ = len(v23)
                                        for fi24 := 0; fi24 < 8; fi24++ {
                                            switch "a" {
                                            case "foo bar":
                                                v25 := GenS3{F0: 0, F1: -32768, F2: 18446744073709551615}
                                                _ = v25
                                            case "\n\t":
                                                if false {
                                                    v26 := GenS4{F0: 0, F1: 2147483648, F2: -32768}
                                                    _ = v26
                                                }
                                            case "":
                                                switch v4 {
                                                case 574:
                                                    if false {
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                        var v27 uint64 = 18446744073709551615
                                                        v27 += 1
                                                        _ = v27
                                                    } else {
                                                        var v28 float32 = (1.5 + (1e10 - float32(-2.25)))
                                                        v2 -= (float64(3.14159) * ((((3.14159 * (3.14159 * 0.1)) / (v2 / 3.14159)) + v1) + ((v1 - 0.0) - v1)))
                                                        switch (v4 - 2147483647) {
                                                        case 42:
                                                            switch -1 {
                                                            case -1:
                                                                if len(v23) >= 2 {
                                                                    v23 = v23[1:len(v23)]
                                                                }
                                                            case -72:
                                                                fmt.Println(v1)
                                                            default:
                                                                var v29 int32 = 0
                                                                _ = v29
                                                            }
                                                        default:
                                                            switch fi24 {
                                                            case 327:
                                                                v30 := GenS4{F0: 4294967295, F1: 2147483648, F2: -784}
                                                                _ = v30
                                                            }
                                                        }
                                                        v2 = (0.1 / v1)
                                                        var v31 uint64 = 1
                                                        var v32 int16 = 42
                                                        fmt.Println(v1, v4)
                                                        _ = v28
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                case 440:
                                                    if false {
                                                        fmt.Println(v3, v4, fi16)
                                                    }
                                                }
                                            default:
                                                v33 := GenS1{F0: -2147483648, F1: 259}
                                                _ = v33
                                            }
                                            fmt.Println(v2, fi24, fi24)
                                            switch fi24 {
                                            case 4294967295:
                                                if len(v23) > 0 {
                                                    v23[0] = (-(fi16))
                                                }
                                            }
                                            _ = fi24
                                        }
                                    }
                                    var v34 string = "hello"
                                    _ = v23
                                    _ = v34
                                }
                            }
                        default:
                            v35 := GenS3{F0: 255, F1: -1, F2: 18446744073709551615}
                            _ = v35
                        }
                        if (-1 >= (v3 * v3)) {
                            var v36 float32 = (-2.25 / -2.25)
                            v37 := GenS2{F0: "\n\t", F1: 1}
                            v3 -= v3
                            _ = v36
                            _ = v37
                        }
                        _ = fi16
                    }
                default:
                    var v38 int16 = 42
                    _ = v38
                }
            }
        case 0:
            v39 := []int{(v4 / ((9223372036854775807) | 1)), -2147483648}
            _ = v39
        case -2147483648:
            if true {
                if true {
                    v2 -= float64(3.14159)
                    var v40 uint64 = 372
                    _ = v40
                }
            }
        default:
            fmt.Println(v4, v1, v4)
        }
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        v3 = v3
        _ = fi41
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        v43 := []int{}
        _ = fi42
        _ = v43
    }
    v44 := GenS2{F0: ("foo bar" + ("hello" + "foo bar")), F1: 2}
    fmt.Println(v4, v2, v3)
    for fi45 := 0; fi45 < 19; fi45++ {
        v1 = (v2 / float64((float32((v1 * (float64(1e10) * float64(-2.25)))) / 0.1)))
        v3 -= (^(v3))
        var v46 uint = 255
        _ = fi45
        _ = v46
    }
    var v47 uint16 = 32767
    for fi48 := 0; fi48 < 12; fi48++ {
        fmt.Println(fi48, v2, v1)
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
}
