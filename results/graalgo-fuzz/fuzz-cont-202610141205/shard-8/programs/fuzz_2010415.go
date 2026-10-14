package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 uint8) uint8 {
    var v6 int = 42
    var v7 uint16 = 908
    v7 = 0
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := GenS1{F0: "hello"}
        for fi10 := 0; fi10 < 19; fi10++ {
            p0 -= (p0 / ((uint8(v7)) | 1))
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return (^(p0))
}
func genFn1() int16 {
    var v11 int = -1
    v12 := []int{v11}
    var v13 float32 = 1e10
    fmt.Println(v13)
    switch v11 {
    case -159:
        switch 2147483647 {
        case -2147483648:
            switch (v11 + (v11 >> ((65535) & 63))) {
            case 0:
                v14 := GenS1{F0: ("hello" + "hello")}
                _ = v14
            case 128:
                var v15 uint8 = 1
                _ = v15
            case -734:
                v11 += v11
            default:
                var v16 uint8 = 42
                _ = v16
            }
        case 9223372036854775807:
            _ = len(v12)
        case 819:
            if (true && (2 > 2)) {
                v17 := GenS1{F0: "a"}
                if true {
                    var v18 int = (v11 / (((v11 % ((v11) | 1))) | 1))
                    if (0 == 65535) {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            v13 -= 0.1
                            _ = len(v12)
                            switch "\n\t" {
                            case "":
                                if false {
                                    for fi20 := 0; fi20 < 9; fi20++ {
                                        if (255 > 65535) {
                                            v18 = 32768
                                        }
                                        var v21 string = ""
                                        v22 := GenS1{F0: v21}
                                        _ = fi20
                                        _ = v21
                                        _ = v22
                                    }
                                    v23 := GenS1{F0: "\n\t"}
                                    _ = v23
                                } else {
                                    for fi24 := 0; fi24 < 10; fi24++ {
                                        if false {
                                            fmt.Println(v13)
                                            v13 += -2.25
                                            var v25 uint8 = 23
                                            _ = v25
                                        } else {
                                            var v26 uint16 = 255
                                            if false {
                                                v13 = 1.5
                                            }
                                            _ = v26
                                        }
                                        _ = fi24
                                    }
                                    if (-1 < 1) {
                                        var v27 uint8 = 42
                                        _ = v27
                                    }
                                }
                            case "a":
                                var v28 uint16 = 283
                                _ = v28
                            case "hello":
                                if (v13 != (v13 * (v13 * 3.14159))) {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                }
                            }
                            _ = fi19
                        }
                        switch (v18 * v18) {
                        case -2147483648:
                            v13 += 1e10
                        case 4294967295:
                            switch "" {
                            case "\n\t":
                                for fi29 := 0; fi29 < 7; fi29++ {
                                    switch "hello" {
                                    case "\n\t":
                                        v30 := GenS2{F0: 255, F1: -2}
                                        _ = v30
                                    case "":
                                        var v31 int16 = -1
                                        _ = v31
                                    }
                                    _ = fi29
                                }
                            case "hello":
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                        case 0:
                            v18 += -331
                        }
                        fmt.Println(v18, v11, v13)
                    } else {
                        var v32 uint64 = 4294967295
                        switch (v18 & 32767) {
                        case 42:
                            v33 := GenS3{F0: ((v32 & 4294967295) >= (^(v32)))}
                            _ = v33
                        default:
                            v13 = v13
                        }
                        _ = v32
                    }
                    _ = v18
                }
                _ = v17
            } else {
                var v34 int64 = 1
                for fi35 := 0; fi35 < 3; fi35++ {
                    switch (fi35 / (((^(fi35))) | 1)) {
                    case -517:
                        v36 := GenS4{F0: "", F1: 0}
                        _ = v36
                    case 627:
                        v37 := GenS5{F0: 42, F1: 32767}
                        _ = v37
                    default:
                        for fi38 := 0; fi38 < 13; fi38++ {
                            fmt.Println(v34)
                            fmt.Println(v13, v34, v13)
                            _ = fi38
                        }
                    }
                    _ = fi35
                }
                _ = v34
            }
        default:
            _ = len(v12)
        }
    }
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := []uint32{1, 343, 2147483647}
    var v2 uint16 = 1
    var v3 int32 = -960
    v4 := []string{("foo bar" + "foo bar")}
    var v5 float64 = (((1.5 / 3.14159) - 1.5) / 1e10)
    v3 = (-(v3))
    v39 := GenS6{F0: 275, F1: v3}
    fmt.Println(v2, v2, v5)
    var v40 int64 = 177
    fmt.Println(v5, v5, v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v41 := GenS7{F0: 1}
    var v42 float32 = (3.14159 + (1.5 / (3.14159 - 0.0)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
