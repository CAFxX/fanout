package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[uint32]int)
    var v3 float64 = float64((1.5 - 0.1))
    var v4 uint16 = 414
    switch int(v4) {
    case 32767:
        var v5 uint = 4294967295
        _ = v5
    case 42:
        _ = len(v2)
    }
    v6 := []int{128, -9223372036854775808}
    for fi7 := 0; fi7 < 18; fi7++ {
        switch "a" {
        case "\n\t":
            if false {
                v8 := GenS1{F0: v4, F1: 0}
                _ = v8
            }
        }
        _ = fi7
    }
    var v9 uint32 = 2147483647
    v10 := GenS1{F0: v4, F1: 291}
    for fi11 := 0; fi11 < 8; fi11++ {
        switch 1 {
        case 128:
            for fi12 := 0; fi12 < 7; fi12++ {
                switch "a" {
                case "a":
                    v13 := GenS1{F0: v4, F1: (v4 >> ((65535) & 63))}
                    switch ("hello" + "foo bar") {
                    case "":
                        fmt.Println(v3, fi11, fi12)
                    case "hello":
                        v1 = v1
                    case "foo bar":
                        if len(v6) > 0 {
                            v6[0] = (fi12 - (^(fi12)))
                        }
                    default:
                        if ("\n\t" >= ("" + "\n\t")) {
                            v2[(v9 << ((76) & 63))] = (-(fi11))
                            v14 := GenS1{F0: v4, F1: 255}
                            if true {
                                var v15 uint32 = v9
                                v1 -= v1
                                _ = v15
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    v16 := GenS1{F0: v4, F1: 1}
                                    _ = v16
                                case "\n\t":
                                    var v17 int64 = -1
                                    _ = v17
                                case "a":
                                    if (119 > -805) {
                                        var v18 uint = 255
                                        switch (^(fi12)) {
                                        case -848:
                                            switch "\n\t" {
                                            case "":
                                                var v19 bool = false
                                                _ = v19
                                            case "\n\t":
                                                fmt.Println(v4, fi12)
                                            case "foo bar":
                                                fmt.Println(v4)
                                            }
                                        case 9223372036854775807:
                                            v18 += uint(v4)
                                        case 210:
                                            switch (fi12 / (((fi11 + 9223372036854775807)) | 1)) {
                                            case 358:
                                                _ = len(v6)
                                                var v20 int = (fi12 << ((uint(v9)) & 63))
                                                fmt.Println(v20, v4)
                                                _ = v20
                                            default:
                                                if false {
                                                    v21 := GenS1{F0: 32767, F1: 1}
                                                    v3 = (v3 * 1e10)
                                                    _ = v21
                                                }
                                            }
                                        }
                                        _ = v18
                                    }
                                }
                                var v22 int64 = 201
                                if ((1 == 179) || false) {
                                    v9 -= 2147483648
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                } else {
                                    var v23 int64 = v22
                                    switch "a" {
                                    case "foo bar":
                                        var v24 float32 = float32(v3)
                                        _ = v24
                                    case "\n\t":
                                        fmt.Println(v22, fi12)
                                    case "":
                                        if true {
                                            v25 := GenS2{F0: (v4 - (^(v4)))}
                                            if (v3 == ((3.14159 / v3) / float64(0.1))) {
                                                v26 := GenS2{F0: v4}
                                                v27 := GenS2{F0: 65535}
                                                _ = v26
                                                _ = v27
                                            } else {
                                                delete(v2, (v9 ^ (v9 - 223)))
                                                v23 += (v23 << ((uint(v9)) & 63))
                                                _ = len(v6)
                                            }
                                            v22 = int64(v4)
                                            v3 = (v3 * v3)
                                            v28 := GenS3{F0: 0, F1: 4294967295, F2: false}
                                            fmt.Println(v9)
                                            var v29 int = int(v9)
                                            _ = v25
                                            _ = v28
                                            _ = v29
                                        }
                                    default:
                                        var v30 float32 = 0.0
                                        _ = v30
                                    }
                                    switch fi11 {
                                    case -799:
                                        var v31 bool = (100 != 100)
                                        _ = v31
                                    case 127:
                                        if ((!true) && false) {
                                            fmt.Println(fi11)
                                            _ = len(v6)
                                            fmt.Println(v4, fi12)
                                            _ = len(v2)
                                        }
                                    case -9223372036854775808:
                                        fmt.Println(v9, v3, v22)
                                    }
                                    _ = v23
                                }
                                var v32 int = (fi12 * fi11)
                                switch fi11 {
                                case 9223372036854775807:
                                    if false {
                                        v9 = (v9 | (^(v9)))
                                    } else {
                                        v33 := GenS2{F0: v4}
                                        var v34 uint = 18446744073709551615
                                        fmt.Println(v1)
                                        fmt.Println(v3)
                                        var v35 uint8 = 2
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                case -1:
                                    if false {
                                        if ((v4 ^ v4) < 2) {
                                            v1 += (v1 * v1)
                                        } else {
                                            var v36 uint = 2
                                            var v37 float32 = (0.0 + (((0.0 + ((((-2.25 / 1e10) - 1.5) / 0.1) - (0.0 - float32(((v3 / 1.0) / float64((1e10 - 0.0))))))) + 0.0) * -2.25))
                                            _ = len(v2)
                                            _ = v36
                                            _ = v37
                                        }
                                        v38 := GenS1{F0: (^(v4)), F1: 255}
                                        _ = v38
                                    }
                                default:
                                    switch fi12 {
                                    case -1:
                                        v39 := GenS4{F0: -1}
                                        _ = v39
                                    case 0:
                                        var v40 uint16 = 1
                                        _ = v40
                                    }
                                }
                                _ = v22
                                _ = v32
                            }
                            var v41 string = ("\n\t" + (("foo bar" + "a") + ("\n\t" + "\n\t")))
                            var v42 float32 = 1e10
                            _ = v14
                            _ = v41
                            _ = v42
                        }
                    }
                    v1 = (-(v1))
                    _ = v13
                case "hello":
                    if len(v6) > 0 {
                        v6[0] = (fi11 ^ (^(fi12)))
                    }
                case "\n\t":
                    fmt.Println(v3, v1, fi12)
                default:
                    delete(v2, (v9 + (^(v9))))
                }
                _ = fi12
            }
        case 0:
            var v43 int32 = (v1 & v1)
            _ = v43
        }
        if true {
            v4 += (v4 & (v4 + 1))
            v4 += 65535
            v3 -= (v3 / 0.1)
        } else {
            switch ("\n\t" + "foo bar") {
            case "a":
                v44 := v2[v9]
                _ = v44
            case "":
                v3 += ((1.5 * -2.25) - (v3 * v3))
            default:
                v2[(v9 << ((0) & 63))] = (fi11 | 4294967295)
            }
        }
        v9 += (v9 ^ 42)
        _ = fi11
    }
    v45 := GenS5{F0: -826}
    var v46 float32 = 1e10
    v46 = (-2.25 - float32(float64(v46)))
    v47 := v2[v9]
    v2[v9] = 0
    switch "a" {
    case "hello":
        if ((("hello" + ("foo bar" + ("hello" + "\n\t"))) != "\n\t") || (!(("" + "foo bar") > "\n\t"))) {
            if (127 < -128) {
                _ = len(v2)
            }
        }
    case "":
        var v48 uint32 = v9
        _ = v48
    }
    v49 := GenS3{F0: 623, F1: 563, F2: false}
    var v50 float64 = (0.1 - v3)
    var v51 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v49
    _ = v50
    _ = v51
    _ = v6
    _ = v9
}
