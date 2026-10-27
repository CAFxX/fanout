package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 int32; F2 uint64 }
func genFn0() int8 {
    var v4 uint64 = 1
    v4 -= v4
    var v5 uint16 = 255
    var v6 int64 = 9223372036854775807
    _ = v4
    _ = v5
    _ = v6
    return int8(v6)
}

func main() {
    v1 := []string{"foo bar"}
    var v2 bool = (((127 < 42) && true) || false)
    var v3 uint = 2
    _ = len(v1)
    for fi7 := 0; fi7 < 12; fi7++ {
        v1 = append(v1, ("\n\t" + "foo bar"))
        switch "" {
        case "foo bar":
            var v8 string = ("hello" + (("foo bar" + "a") + "hello"))
            _ = v8
        }
        var v9 bool = (float64(0.1) >= 0.0)
        _ = fi7
        _ = v9
    }
    fmt.Println(v3, v3)
    v2 = false
    _ = len(v1)
    v1 = append(v1, "hello")
    v10 := GenS1{F0: 350, F1: -107}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ((("a" + ("a" + "")) + ("\n\t" + "\n\t")) + "") {
    case "foo bar":
        fmt.Println(v2, v3)
    case "a":
        var v11 uint64 = 42
        _ = v11
    case "hello":
        v1 = append(v1, "\n\t")
    }
    v1 = append(v1, "\n\t")
    var v12 uint32 = 2147483648
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: 42, F1: fi13}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        if (1 != 2147483647) {
            if (!false) {
                v12 += v12
                var v16 string = ("" + "\n\t")
                _ = v16
            } else {
                var v17 bool = v2
                _ = v17
            }
            switch "\n\t" {
            case "\n\t":
                if true {
                    fmt.Println(v2)
                    if (v2 || v2) {
                        fmt.Println(v12)
                    }
                    v18 := GenS1{F0: 930, F1: (fi15 / (((^(fi15))) | 1))}
                    _ = v18
                }
            }
            fmt.Println(v3, v3, v2)
        } else {
            v19 := GenS1{F0: 42, F1: 128}
            _ = len(v1)
            switch fi15 {
            case 2147483647:
                v3 += v3
            default:
                if ((0.1 * float32((3.14159 + 0.1))) >= 0.0) {
                    v12 += v12
                } else {
                    var v20 int32 = -392
                    for fi21 := 0; fi21 < 17; fi21++ {
                        switch ("hello" + "a") {
                        case "hello":
                            var v22 uint16 = 678
                            _ = v22
                        case "":
                            switch "hello" {
                            case "hello":
                                v3 = 255
                            case "":
                                v3 -= v3
                                var v23 int16 = -32768
                                _ = v23
                            default:
                                v24 := GenS1{F0: 97, F1: fi21}
                                _ = v24
                            }
                        case "a":
                            var v25 int8 = 1
                            _ = v25
                        default:
                            v26 := GenS1{F0: 2, F1: (fi15 - (fi15 & fi21))}
                            _ = v26
                        }
                        v27 := GenS1{F0: 32767, F1: (fi15 ^ fi15)}
                        v12 -= v12
                        _ = fi21
                        _ = v27
                    }
                    _ = v20
                }
            }
            _ = v19
        }
        _ = fi15
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v28 := GenS2{F0: -414, F1: 1, F2: 146}
    for fi29 := 0; fi29 < 3; fi29++ {
        v30 := GenS2{F0: 0, F1: -1, F2: 2}
        _ = fi29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v28
    _ = v3
}
