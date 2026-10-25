package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint16) int8 {
    v7 := []int{-914, -9223372036854775808}
    p1 = p1
    if ((0 != -99) && true) {
        switch ("" + "hello") {
        case "":
            p0 = 4294967295
        default:
            v8 := GenS1{F0: 18446744073709551615}
            _ = v8
        }
        if false {
            _ = len(v7)
            fmt.Println(p0)
            switch -1 {
            case 4294967295:
                for fi9 := 0; fi9 < 19; fi9++ {
                    p0 += p0
                    _ = fi9
                }
            case 42:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case 629:
                _ = len(v7)
            default:
                fmt.Println(p1, p1)
            }
        } else {
            var v10 uint32 = p0
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 100
}

func main() {
    var v1 uint8 = 42
    var v2 uint64 = 2
    var v3 int16 = 1
    var v4 uint32 = 42
    var v5 uint = 884
    v6 := []int{-217}
    v11 := GenS2{F0: 100, F1: v4, F2: (v2 * uint64(v5))}
    switch ("hello" + "") {
    case "a":
        switch -9223372036854775808 {
        case -326:
            v12 := GenS2{F0: genFn0(v4, 42), F1: (v4 + (v4 % (((v4 - v4)) | 1))), F2: 1}
            _ = v12
        case 355:
            v13 := GenS2{F0: -128, F1: (v4 & v4), F2: v2}
            _ = v13
        default:
            switch "\n\t" {
            case "a":
                for fi14 := 0; fi14 < 12; fi14++ {
                    switch "a" {
                    case "hello":
                        for fi15 := 0; fi15 < 1; fi15++ {
                            v16 := GenS3{F0: "hello"}
                            _ = fi15
                            _ = v16
                        }
                    default:
                        if true {
                            v17 := GenS3{F0: "\n\t"}
                            v6 = append(v6, 1)
                            _ = v17
                        }
                    }
                    var v18 float64 = -2.25
                    var v19 uint = v5
                    _ = fi14
                    _ = v18
                    _ = v19
                }
            case "hello":
                fmt.Println(v5)
            default:
                for fi20 := 0; fi20 < 1; fi20++ {
                    if (2 != 42) {
                        v2 -= (v2 ^ (v2 - (v2 * (^(v2)))))
                        v21 := GenS4{F0: 65535, F1: (v1 / (((v1 << (((v5 ^ (v5 * (^(v5))))) & 63))) | 1))}
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v21
                    } else {
                        var v22 float32 = ((1e10 + 0.1) * 3.14159)
                        v6 = append(v6, (^(fi20)))
                        for fi23 := 0; fi23 < 4; fi23++ {
                            fmt.Println(v22)
                            v1 = v1
                            var v24 uint64 = (v2 * (^(v2)))
                            _ = fi23
                            _ = v24
                        }
                        _ = v22
                    }
                    var v25 uint32 = (^(v4))
                    v1 -= (^(v1))
                    _ = fi20
                    _ = v25
                }
            }
        }
    case "foo bar":
        var v26 uint64 = 2
        _ = v26
    default:
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 bool = ((v1 | 127) > v1)
            for fi29 := 0; fi29 < 12; fi29++ {
                fmt.Println(v4, v5, v3)
                v2 -= 4294967295
                var v30 float32 = 1.5
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
    }
    v3 = v3
    var v31 uint64 = 18446744073709551615
    fmt.Println(v4, v3, v5)
    if (v2 < v31) {
        v1 = 42
        switch 988 {
        case 92:
            if false {
                v6 = append(v6, -874)
                var v32 int8 = 127
                _ = v32
            }
        case 4294967295:
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        default:
            switch "foo bar" {
            case "hello":
                var v33 uint64 = (v31 * (v2 << ((v5) & 63)))
                _ = v33
            default:
                fmt.Println(v2, v5, v31)
            }
        }
    }
    v34 := GenS4{F0: 42, F1: (v1 << ((v5) & 63))}
    var v35 float64 = (-2.25 + float64(-2.25))
    v5 += (^(v5))
    var v36 uint32 = v4
    var v37 bool = (false || (!true))
    v38 := GenS5{F0: (v3 - (^(v3)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
