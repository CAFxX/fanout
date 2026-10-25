package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0(p0 int) uint16 {
    var v5 int16 = 32767
    v6 := GenS1{F0: 127}
    _ = p0
    _ = v5
    _ = v6
    return 59
}

func main() {
    v1 := []string{"hello"}
    var v2 int16 = -599
    var v3 int64 = 1
    var v4 int16 = v2
    v7 := GenS2{F0: 1, F1: 784, F2: 42}
    if ((!false) && false) {
        v2 = v2
        if (0 == 250) {
            var v8 int64 = v3
            _ = v8
        } else {
            for fi9 := 0; fi9 < 10; fi9++ {
                v10 := GenS3{F0: 408, F1: (v4 >> ((463) & 63))}
                if ("\n\t" >= "") {
                    var v11 uint8 = 255
                    _ = v11
                } else {
                    if ((fi9 % ((fi9) | 1)) < -151) {
                        switch ("" + "\n\t") {
                        case "a":
                            var v12 int16 = v4
                            _ = v12
                        case "":
                            for fi13 := 0; fi13 < 7; fi13++ {
                                var v14 string = ("foo bar" + "foo bar")
                                _ = fi13
                                _ = v14
                            }
                        case "hello":
                            for fi15 := 0; fi15 < 12; fi15++ {
                                var v16 bool = (!true)
                                _ = fi15
                                _ = v16
                            }
                        }
                    } else {
                        if (false || true) {
                            v4 = -962
                            var v17 int32 = -2147483648
                            if (!false) {
                                if (false && true) {
                                    for fi18 := 0; fi18 < 10; fi18++ {
                                        var v19 uint = 1
                                        _ = fi18
                                        _ = v19
                                    }
                                    v4 = v2
                                    var v20 float64 = (-2.25 / 3.14159)
                                    _ = v20
                                }
                            }
                            _ = v17
                        }
                    }
                }
                _ = fi9
                _ = v10
            }
        }
    } else {
        v1 = append(v1, "hello")
        switch "a" {
        case "hello":
            if (777 >= 42) {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
            }
        case "a":
            v4 = -32768
        case "\n\t":
            for fi21 := 0; fi21 < 8; fi21++ {
                v4 = -32768
                var v22 uint64 = 0
                _ = fi21
                _ = v22
            }
        default:
            fmt.Println(v2)
        }
    }
    var v23 int64 = 9223372036854775807
    _ = len(v1)
    var v24 float64 = ((0.1 - 3.14159) / 3.14159)
    if true {
        switch "a" {
        case "foo bar":
            var v25 uint64 = 2
            _ = v25
        case "hello":
            v26 := GenS1{F0: 86}
            _ = v26
        }
    }
    v2 = -734
    switch -684 {
    case 9223372036854775807:
        var v27 int32 = -360
        _ = v27
    case 127:
        var v28 uint64 = 1
        _ = v28
    case 128:
        fmt.Println(v24)
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        if (("\n\t" == "") || true) {
            var v30 float32 = 3.14159
            v2 += (v4 / ((v2) | 1))
            _ = v30
        }
        v31 := GenS4{F0: (v4 % ((924) | 1)), F1: 270}
        _ = fi29
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        v3 = v23
        v33 := GenS5{F0: (v2 - (v2 / ((1) | 1))), F1: (v3 & (v3 ^ -1))}
        v34 := GenS3{F0: 42, F1: v4}
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
}
