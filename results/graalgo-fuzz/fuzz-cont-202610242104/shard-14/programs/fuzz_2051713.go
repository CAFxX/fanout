package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int64; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]string)
    v3 := []string{("a" + "hello"), "a"}
    switch 42 {
    case -9223372036854775808:
        v4 := GenS1{F0: -1, F1: 4294967295}
        _ = v4
    case -136:
        switch ("a" + "hello") {
        case "\n\t":
            v5 := GenS2{F0: 1000}
            _ = v5
        }
    default:
        var v6 int16 = -1
        _ = v6
    }
    delete(v2, 2147483647)
    v1 = 969
    fmt.Println(v1)
    var v7 int16 = 393
    var v8 int64 = -686
    var v9 float32 = (float32(-2.25) - 3.14159)
    for fi10 := 0; fi10 < 5; fi10++ {
        if ((v1 > (v1 - (v1 % ((65535) | 1)))) && false) {
            _ = len(v3)
            if false {
                switch (("" + ("\n\t" + "hello")) + ("\n\t" + "")) {
                case "hello":
                    v9 -= (float32((float64(3.14159) - 0.1)) / v9)
                case "a":
                    v11 := GenS3{F0: v7, F1: (v1 * (^(v1)))}
                    _ = v11
                default:
                    v1 += (v1 + (v1 - 255))
                }
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    v12 := GenS3{F0: (v7 + (v7 * (v7 + (v7 / ((v7) | 1))))), F1: v1}
                    _ = v12
                case "":
                    if (!(!(492 == 1))) {
                        _ = len(v3)
                        var v13 int = int(v1)
                        switch (v13 & fi10) {
                        case 9223372036854775807:
                            v14 := GenS2{F0: (v7 % ((-880) | 1))}
                            _ = v14
                        default:
                            v7 -= (v7 * (^(v7)))
                        }
                        _ = v13
                    }
                case "a":
                    var v15 float32 = 0.0
                    _ = v15
                default:
                    var v16 float64 = 1e10
                    _ = v16
                }
            } else {
                switch fi10 {
                case 127:
                    if (false || true) {
                        switch -1 {
                        case 32768:
                            v7 += (-(v7))
                        case 142:
                            if false {
                                switch "hello" {
                                case "foo bar":
                                    if (true || (v7 > (v7 << ((v1) & 63)))) {
                                        v1 += 255
                                        fmt.Println(v8, v1)
                                        switch fi10 {
                                        case -9223372036854775808:
                                            v2[298] = "foo bar"
                                        case 9223372036854775807:
                                            if (!false) {
                                                v3 = append(v3, "hello")
                                                v7 += v7
                                            } else {
                                                if false {
                                                    fmt.Println(v8, v8, fi10)
                                                    for fi17 := 0; fi17 < 2; fi17++ {
                                                        if true {
                                                            v18 := GenS4{F0: v8, F1: 0, F2: 1000}
                                                            _ = v18
                                                        }
                                                        fmt.Println(v1, v8, fi17)
                                                        var v19 float32 = v9
                                                        _ = fi17
                                                        _ = v19
                                                    }
                                                }
                                                v20 := GenS5{F0: 4294967295}
                                                _ = v20
                                            }
                                        case 128:
                                            if false {
                                                var v21 int = fi10
                                                for fi22 := 0; fi22 < 2; fi22++ {
                                                    _ = len(v3)
                                                    var v23 int = (-(fi10))
                                                    switch (fi10 * fi10) {
                                                    case -9223372036854775808:
                                                        var v24 uint32 = 1
                                                        _ = v24
                                                    }
                                                    _ = fi22
                                                    _ = v23
                                                }
                                                _ = v21
                                            } else {
                                                if (("a" + "\n\t") >= "a") {
                                                    var v25 int8 = int8(fi10)
                                                    if len(v3) > 0 {
                                                        v3[0] = ("\n\t" + (("hello" + "foo bar") + "a"))
                                                    }
                                                    switch fi10 {
                                                    case 32768:
                                                        v2[2147483648] = ""
                                                    case 1:
                                                        v26 := GenS3{F0: v7, F1: v1}
                                                        _ = v26
                                                    case 762:
                                                        var v27 string = "foo bar"
                                                        _ = v27
                                                    }
                                                    _ = v25
                                                } else {
                                                    v1 += v1
                                                    for fi28 := 0; fi28 < 3; fi28++ {
                                                        if (false || true) {
                                                            if true {
                                                                var v29 int16 = v7
                                                                _ = v29
                                                            } else {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                var v30 uint64 = 18
                                                                v1 += 974
                                                                switch "foo bar" {
                                                                case "":
                                                                    var v31 uint16 = 65535
                                                                    _ = v31
                                                                }
                                                                _ = v30
                                                            }
                                                        } else {
                                                            var v32 int32 = -1
                                                            delete(v2, 2147483647)
                                                            _ = v32
                                                        }
                                                        _ = len(v2)
                                                        v7 = int16(v8)
                                                        _ = fi28
                                                    }
                                                }
                                                v7 = v7
                                            }
                                        default:
                                            var v33 int8 = 127
                                            _ = v33
                                        }
                                    } else {
                                        v8 += -9223372036854775808
                                        var v34 int64 = (v8 ^ v8)
                                        _ = v34
                                    }
                                case "\n\t":
                                    switch "a" {
                                    case "a":
                                        delete(v2, 719)
                                    default:
                                        v7 = v7
                                    }
                                }
                            }
                        default:
                            v7 = (v7 | v7)
                        }
                        v7 -= (^(v7))
                        v35 := GenS5{F0: 4294967295}
                        _ = v35
                    }
                case -517:
                    var v36 string = "hello"
                    _ = v36
                case -340:
                    for fi37 := 0; fi37 < 6; fi37++ {
                        v38 := v2[65535]
                        _ = fi37
                        _ = v38
                    }
                }
            }
        }
        var v39 float32 = (v9 / (v9 - v9))
        _ = fi10
        _ = v39
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        v41 := GenS1{F0: -97, F1: 497}
        v42 := GenS3{F0: -32768, F1: v1}
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 int64 = (-(v8))
    v1 = 0
    var v44 int = -9223372036854775808
    v45 := GenS1{F0: 42, F1: 1}
    v43 = (v43 & 1)
    var v46 int = (v44 * (v44 & 112))
    v7 -= v7
    v7 -= (^(v7))
    if (!(false && true)) {
        fmt.Println(v9)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v7
    _ = v8
    _ = v9
}
