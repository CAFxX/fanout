package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int64; F2 uint }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int }
func genFn0(p0 int16, p1 int16, p2 uint16) uint64 {
    v5 := []int{608, -1, 4294967295, -9223372036854775808}
    if true {
        switch -2147483648 {
        case -540:
            switch "" {
            case "":
                var v6 string = "a"
                _ = v6
            }
        case -1:
            if (((3.14159 + 0.1) - 0.1) >= 0.1) {
                v7 := GenS1{F0: 2147483647}
                fmt.Println(p1)
                _ = v7
            }
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            if false {
                var v9 uint64 = 42
                v10 := GenS1{F0: -291}
                _ = v10
                _ = v9
            }
            var v11 float64 = ((-2.25 * (float64(3.14159) / float64(3.14159))) / -2.25)
            _ = fi8
            _ = v11
        }
    } else {
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
        var v12 uint8 = 255
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 200
    v3 := []string{}
    v4 := make(map[uint32]string)
    if (!false) {
        for fi13 := 0; fi13 < 20; fi13++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v14 int32 = (v2 & v2)
            _ = fi13
            _ = v14
        }
        var v15 string = ""
        var v16 bool = true
        _ = v15
        _ = v16
    }
    var v17 float32 = 1e10
    switch ("\n\t" + "foo bar") {
    case "a":
        var v18 string = "hello"
        _ = v18
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v17 += 0.0
    var v19 uint64 = 18446744073709551615
    v17 += ((((float32(float64(v17)) + 1e10) + ((v17 + 0.1) * v17)) - ((v17 * 0.1) + -2.25)) + (-2.25 + 3.14159))
    for fi20 := 0; fi20 < 12; fi20++ {
        if (false && true) {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 int8 = 100
                _ = fi21
                _ = v22
            }
            if ((^(fi20)) <= 213) {
                v23 := GenS2{F0: "foo bar", F1: 592, F2: 328}
                switch "foo bar" {
                case "\n\t":
                    if false {
                        for fi24 := 0; fi24 < 7; fi24++ {
                            if (430 < 65535) {
                                v19 = 4294967295
                                var v25 uint16 = 1
                                switch "" {
                                case "foo bar":
                                    v2 += (^(v2))
                                case "a":
                                    var v26 float64 = float64(v17)
                                    _ = v26
                                default:
                                    var v27 float64 = (0.0 / 0.1)
                                    _ = v27
                                }
                                _ = v25
                            }
                            switch fi24 {
                            case -1:
                                switch fi20 {
                                case 32767:
                                    _ = len(v3)
                                    if (true || true) {
                                        switch "\n\t" {
                                        case "a":
                                            switch ("\n\t" + "a") {
                                            case "a":
                                                v28 := GenS1{F0: v2}
                                                _ = v28
                                            case "hello":
                                                v29 := GenS1{F0: (^(v2))}
                                                _ = v29
                                            case "foo bar":
                                                v30 := GenS2{F0: "hello", F1: 9223372036854775807, F2: 692}
                                                _ = v30
                                            default:
                                                v31 := GenS2{F0: "\n\t", F1: 148, F2: 18446744073709551615}
                                                _ = v31
                                            }
                                        case "hello":
                                            if (false && true) {
                                                fmt.Println(v2, v19, fi20)
                                            }
                                        default:
                                            var v32 string = ""
                                            var v33 int64 = 0
                                            v32 = "a"
                                            _ = v32
                                            _ = v33
                                        }
                                        v34 := GenS2{F0: "hello", F1: 943, F2: 0}
                                        _ = v34
                                    }
                                    v19 = genFn0(689, 32767, 42)
                                }
                            default:
                                v17 += ((v17 - (3.14159 - (((v17 / ((v17 * v17) / 1.0)) - float32((0.1 - 1.5))) * (v17 + float32(0.1))))) / float32(1e10))
                            }
                            _ = fi24
                        }
                        switch "foo bar" {
                        case "":
                            v35 := GenS3{F0: 1, F1: (v2 - (v2 & v2))}
                            _ = v35
                        default:
                            _ = len(v3)
                        }
                    }
                }
                _ = v23
            }
        } else {
            fmt.Println(v2, v19, v19)
            var v36 uint = 612
            _ = v36
        }
        var v37 int = fi20
        var v38 uint = 1
        _ = fi20
        _ = v37
        _ = v38
    }
    fmt.Println(v17)
    for fi39 := 0; fi39 < 3; fi39++ {
        fmt.Println(v19, v2, v19)
        v19 -= (v19 & 1)
        _ = fi39
    }
    v40 := GenS4{F0: 0, F1: 2}
    v41 := GenS5{F0: v19, F1: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
}
