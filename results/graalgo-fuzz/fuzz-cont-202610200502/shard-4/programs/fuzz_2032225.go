package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32, p1 int16, p2 uint) string {
    p0 = 4294967295
    if true {
        v4 := GenS1{F0: 35, F1: -1}
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 uint32 = (p0 % (((p0 ^ p0)) | 1))
            var v7 uint16 = 65535
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = v4
    } else {
        var v8 int64 = -1
        v9 := []int{}
        var v10 uint8 = 127
        _ = v10
        _ = v8
        _ = v9
    }
    p0 += 2
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}
func genFn1(p0 int32) bool {
    fmt.Println(p0)
    var v11 float64 = 0.1
    _ = p0
    _ = v11
    return false
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 2147483647
    v3 := []int{442, 772}
    v2 = v2
    if true {
        var v12 int32 = v2
        switch -354 {
        case 731:
            if false {
                v13 := GenS1{F0: (v2 % (((v12 * 188)) | 1)), F1: 100}
                _ = v13
            } else {
                var v14 float32 = (3.14159 + (3.14159 / (0.1 - 1.5)))
                switch 1 {
                case 874:
                    switch "\n\t" {
                    case "\n\t":
                        switch -606 {
                        case -1:
                            v12 -= v2
                        default:
                            var v15 uint = 255
                            _ = v15
                        }
                    case "a":
                        v2 += v12
                    case "":
                        fmt.Println(v12, v2, v14)
                    default:
                        if true {
                            if ((v2 >> ((18446744073709551615) & 63)) != 2147483647) {
                                _ = len(v1)
                            } else {
                                var v16 float64 = (-2.25 - 0.0)
                                if (("foo bar" + "\n\t") <= ("foo bar" + "a")) {
                                    var v17 uint16 = 65535
                                    v18 := GenS1{F0: v2, F1: 0}
                                    var v19 uint32 = 4294967295
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                } else {
                                    v12 -= v2
                                }
                                _ = v16
                            }
                        }
                    }
                case 1:
                    if (!(!false)) {
                        var v20 float64 = 3.14159
                        for fi21 := 0; fi21 < 16; fi21++ {
                            var v22 bool = true
                            for fi23 := 0; fi23 < 18; fi23++ {
                                var v24 uint16 = 1
                                _ = fi23
                                _ = v24
                            }
                            v1[2147483648] = ("foo bar" + ("\n\t" + ""))
                            _ = fi21
                            _ = v22
                        }
                        v25 := GenS1{F0: v12, F1: 102}
                        _ = v20
                        _ = v25
                    }
                case 4294967295:
                    _ = len(v3)
                }
                _ = v14
            }
        case -1:
            v26 := GenS1{F0: 60, F1: 100}
            _ = v26
        }
        _ = v12
    }
    if true {
        if true {
            var v27 int8 = -38
            var v28 float64 = (3.14159 / float64((0.0 - 3.14159)))
            _ = v27
            _ = v28
        } else {
            v29 := GenS2{F0: 9223372036854775807, F1: 1}
            _ = v29
        }
    } else {
        for fi30 := 0; fi30 < 12; fi30++ {
            v2 += -2147483648
            _ = fi30
        }
    }
    fmt.Println(v2)
    v2 = (^(v2))
    v2 += 856
    var v31 uint64 = 18446744073709551615
    var v32 uint32 = 42
    for fi33 := 0; fi33 < 11; fi33++ {
        v3 = append(v3, (fi33 + fi33))
        v34 := GenS3{F0: (v32 <= v32)}
        switch fi33 {
        case 2147483647:
            v2 -= v2
        case -9223372036854775808:
            var v35 uint8 = 2
            _ = v35
        case 4294967295:
            var v36 int64 = -9223372036854775808
            _ = v36
        default:
            v31 += v31
        }
        _ = fi33
        _ = v34
    }
    var v37 uint16 = 0
    for fi38 := 0; fi38 < 8; fi38++ {
        var v39 uint8 = 128
        _ = fi38
        _ = v39
    }
    v40 := GenS1{F0: (v2 - v2), F1: -91}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v40
}
