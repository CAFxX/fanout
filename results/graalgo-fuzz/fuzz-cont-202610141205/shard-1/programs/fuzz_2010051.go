package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 uint8; F1 int8 }
func genFn0(p0 int32) bool {
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            p0 -= p0
            switch fi6 {
            case 4294967295:
                var v8 uint32 = 617
                _ = v8
            default:
                fmt.Println(fi6)
            }
            v9 := GenS1{F0: 0, F1: (p0 - 2147483647), F2: 2}
            _ = fi7
            _ = v9
        }
        var v10 float32 = 0.0
        _ = fi6
        _ = v10
    }
    fmt.Println(p0)
    p0 += (^(p0))
    _ = p0
    return true
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = ((1.5 + 0.1) / 3.14159)
    v3 := []string{"hello", "a", "hello"}
    var v4 uint = 496
    var v5 float64 = 0.0
    if (((v4 * v4) == 313) && false) {
        var v11 int64 = 9223372036854775807
        switch -1 {
        case 42:
            v2 = -2.25
        }
        _ = v11
    }
    var v12 float64 = 3.14159
    switch "a" {
    case "foo bar":
        var v13 int64 = -979
        _ = v13
    }
    var v14 int16 = 807
    for fi15 := 0; fi15 < 5; fi15++ {
        if true {
            switch (fi15 & int(v4)) {
            case -72:
                v12 = 0.0
            case -457:
                fmt.Println(v14, v4)
            }
        }
        if true {
            if (v2 >= 0.0) {
                if (!false) {
                    v3 = append(v3, "a")
                } else {
                    v16 := GenS1{F0: 0, F1: 1, F2: 0}
                    switch fi15 {
                    case -9223372036854775808:
                        if true {
                            switch 278 {
                            case 128:
                                var v17 uint16 = 42
                                _ = v17
                            case 32767:
                                switch (((("a" + "hello") + (("" + "\n\t") + ("a" + "\n\t"))) + "hello") + "foo bar") {
                                case "hello":
                                    delete(v1, -9223372036854775808)
                                case "":
                                    if false {
                                        if (!false) {
                                            v18 := GenS2{F0: v14, F1: ""}
                                            if ((1 == (v14 << (((v4 ^ (v4 / (((v4 | (v4 >> (((^(v4))) & 63)))) | 1)))) & 63))) || (-9223372036854775808 <= 0)) {
                                                var v19 uint32 = 95
                                                _ = v19
                                            }
                                            _ = v18
                                        }
                                    } else {
                                        var v20 int8 = 91
                                        delete(v1, -1)
                                        _ = v20
                                    }
                                default:
                                    if false {
                                        var v21 int32 = -1
                                        v14 += -1
                                        v22 := GenS2{F0: (^(v14)), F1: "a"}
                                        _ = v21
                                        _ = v22
                                    }
                                }
                            case -819:
                                switch 32768 {
                                case 4294967295:
                                    v23 := GenS1{F0: (v14 & v14), F1: -2147483648, F2: 76}
                                    _ = v23
                                default:
                                    for fi24 := 0; fi24 < 4; fi24++ {
                                        var v25 int64 = 0
                                        _ = fi24
                                        _ = v25
                                    }
                                }
                            }
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    default:
                        fmt.Println(fi15)
                    }
                    _ = v16
                }
                v4 = 756
            } else {
                v2 -= -2.25
                var v26 uint = (^(v4))
                v27 := GenS3{F0: (v4 << ((v26) & 63)), F1: 2147483648}
                _ = v26
                _ = v27
            }
            if true {
                _ = len(v1)
                switch ("a" + "") {
                case "foo bar":
                    switch (fi15 / ((127) | 1)) {
                    case 519:
                        var v28 int = -394
                        _ = v28
                    case -1:
                        var v29 int32 = -258
                        _ = v29
                    default:
                        v5 -= ((3.14159 / 1.0) + v12)
                    }
                case "\n\t":
                    for fi30 := 0; fi30 < 8; fi30++ {
                        var v31 int64 = 493
                        v32 := GenS4{F0: 0}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                case "a":
                    switch -634 {
                    case 0:
                        var v33 uint16 = 255
                        _ = v33
                    case -1:
                        switch "" {
                        case "hello":
                            switch ("foo bar" + "") {
                            case "\n\t":
                                _ = len(v1)
                            }
                        }
                    default:
                        if false {
                            v5 = (((v5 + (v12 - ((float64(float32(v5)) - float64(3.14159)) * 3.14159))) / ((v5 - float64(float32(0.1))) * 3.14159)) * 3.14159)
                        } else {
                            v14 -= 42
                            _ = len(v3)
                            var v34 int64 = 42
                            _ = v34
                        }
                    }
                default:
                    var v35 bool = false
                    _ = v35
                }
            } else {
                v36 := GenS5{F0: "a", F1: v4}
                v2 += ((v2 - v2) * (3.14159 - 1.5))
                var v37 uint = 18446744073709551615
                _ = v36
                _ = v37
            }
        } else {
            var v38 int32 = -479
            _ = v38
        }
        v39 := GenS6{F0: v14, F1: (fi15 ^ (fi15 + -933))}
        _ = fi15
        _ = v39
    }
    v14 -= 42
    if false {
        v40 := GenS6{F0: 0, F1: -1}
        if len(v3) > 0 {
            v3[0] = ""
        }
        _ = v40
    }
    _ = len(v1)
    if true {
        for fi41 := 0; fi41 < 9; fi41++ {
            v14 = v14
            fmt.Println(fi41)
            _ = fi41
        }
        fmt.Println(v2, v5)
        v42 := GenS3{F0: (v4 * (^(v4))), F1: 143}
        _ = v42
    } else {
        fmt.Println(v2, v4)
        v43 := GenS4{F0: 65535}
        var v44 int32 = -943
        _ = v43
        _ = v44
    }
    v45 := GenS7{F0: 221, F1: -1}
    var v46 string = "hello"
    var v47 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
