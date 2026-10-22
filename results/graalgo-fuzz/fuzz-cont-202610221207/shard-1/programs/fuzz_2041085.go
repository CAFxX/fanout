package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 int32, p1 uint16) int64 {
    var v5 uint = 65535
    v6 := []int{128}
    var v7 float32 = -2.25
    if (3.14159 > v7) {
        if (((42 > -1) && true) && true) {
            var v8 uint = (v5 - (v5 ^ v5))
            p1 = (p1 * 255)
            v6 = append(v6, 127)
            _ = v8
        } else {
            var v9 uint16 = p1
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint32) uint32 {
    var v10 int32 = -2147483648
    switch "" {
    case "":
        var v11 int16 = -1
        _ = v11
    default:
        var v12 int16 = 42
        _ = v12
    }
    var v13 int32 = (v10 % (((v10 ^ v10)) | 1))
    for fi14 := 0; fi14 < 4; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            if true {
                v13 = (^(v13))
            }
            var v16 uint64 = 4294967295
            v13 += v10
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    var v17 uint64 = 917
    _ = p0
    _ = v10
    _ = v13
    _ = v17
    return (^(p0))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (127 <= 116)
    var v3 uint = 4294967295
    v4 := []uint32{}
    switch 2147483647 {
    case -798:
        var v18 string = ("a" + ("" + ""))
        _ = v18
    }
    if (0.1 != 3.14159) {
        fmt.Println(v2, v3)
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        var v20 uint32 = 2147483648
        if v2 {
            var v21 string = "hello"
            v22 := GenS1{F0: (v20 ^ v20), F1: 1}
            _ = v21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    switch 9223372036854775807 {
    case 128:
        var v23 bool = v2
        _ = v23
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        if true {
            var v25 int64 = 238
            var v26 int16 = 42
            v26 -= (^(v26))
            _ = v25
            _ = v26
        }
        var v27 bool = ((0.0 / 0.1) != float32((1e10 + 1e10)))
        v27 = ((0 == 143) && true)
        _ = fi24
        _ = v27
    }
    var v28 int8 = -91
    if v2 {
        _ = len(v1)
        fmt.Println(v2, v28, v3)
        for fi29 := 0; fi29 < 8; fi29++ {
            v2 = true
            _ = fi29
        }
    } else {
        var v30 bool = v2
        _ = v30
    }
    var v31 uint8 = 0
    delete(v1, (("hello" + "") + "foo bar"))
    var v32 int16 = -32768
    v31 = 210
    _ = len(v1)
    var v33 bool = false
    v34 := v1[("\n\t" + ("hello" + ""))]
    if false {
        for fi35 := 0; fi35 < 5; fi35++ {
            fmt.Println(v28, v33)
            v32 -= (v32 + -147)
            var v36 uint16 = 255
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            v34 = (v34 | v34)
            for fi38 := 0; fi38 < 4; fi38++ {
                v1[("hello" + ("hello" + "hello"))] = genFn1(v34)
                _ = fi38
            }
            switch "foo bar" {
            case "foo bar":
                for fi39 := 0; fi39 < 3; fi39++ {
                    v40 := GenS1{F0: 65535, F1: -128}
                    v34 = genFn1(v34)
                    var v41 int32 = 2147483647
                    _ = fi39
                    _ = v40
                    _ = v41
                }
            case "\n\t":
                if v33 {
                    for fi42 := 0; fi42 < 6; fi42++ {
                        switch "foo bar" {
                        case "\n\t":
                            v31 += (v31 % ((v31) | 1))
                        case "foo bar":
                            v34 -= v34
                        case "hello":
                            fmt.Println(v3, fi37, v28)
                        }
                        v1[""] = (^(v34))
                        fmt.Println(fi37, v32, v2)
                        _ = fi42
                    }
                }
            default:
                if (v33 || v2) {
                    v2 = (-2147483648 < -926)
                    _ = len(v4)
                } else {
                    v43 := GenS1{F0: (v34 % ((2147483647) | 1)), F1: (v28 / (((^(v28))) | 1))}
                    _ = v43
                }
            }
            _ = fi37
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 -= (v3 & v3)
    for fi44 := 0; fi44 < 9; fi44++ {
        var v45 uint32 = genFn1((v34 / ((v34) | 1)))
        _ = fi44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
