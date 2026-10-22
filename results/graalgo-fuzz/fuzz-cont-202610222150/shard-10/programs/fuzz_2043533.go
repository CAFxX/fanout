package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -128
    var v3 uint16 = 255
    v4 := make(map[int]bool)
    v3 = v3
    v5 := []int{}
    var v6 uint = 1
    var v7 uint = v6
    var v8 uint16 = (v3 + (v3 / ((v3) | 1)))
    v9 := GenS1{F0: false, F1: 4294967295}
    v1[9223372036854775807] = 4294967295
    fmt.Println(v7, v2, v7)
    fmt.Println(v3, v3)
    if false {
        if (42 >= -1) {
            fmt.Println(v2, v6)
            var v10 float64 = 1.5
            _ = v10
        }
    } else {
        if (2 >= 4294967295) {
            var v11 string = (("foo bar" + ("" + "\n\t")) + "a")
            v12 := GenS1{F0: (738 >= 18446744073709551615), F1: 759}
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            if false {
                delete(v4, 1000)
                _ = len(v5)
                v2 += v2
            } else {
                switch (("" + ("" + "\n\t")) + ("hello" + "")) {
                case "foo bar":
                    v6 -= v7
                default:
                    if (-416 == -547) {
                        v3 -= (v8 - (v3 / (((v8 | 73)) | 1)))
                        switch fi14 {
                        case -26:
                            if false {
                                _ = len(v4)
                                fmt.Println(v3)
                                if len(v5) > 0 {
                                    v5[0] = fi13
                                }
                                switch -110 {
                                case 401:
                                    if (v2 != v2) {
                                        if (!(!((((((3.14159 - 0.0) * (0.0 - (3.14159 + (1e10 / 0.1)))) / -2.25) < (3.14159 / ((-2.25 / (0.0 - ((0.0 / (-2.25 - ((((((0.0 - (((1.5 * 1.5) - 0.0) - ((1.5 / float64(-2.25)) + 3.14159))) - 0.1) + 0.0) - 0.0) - 3.14159) + 1e10))) - -2.25))) * 3.14159))) && (((!true) && true) && true)) || true))) {
                                            _ = len(v5)
                                            v6 += v6
                                            var v15 string = "a"
                                            _ = v15
                                        }
                                        fmt.Println(v6, v3)
                                        v16 := GenS2{F0: 2147483647, F1: 2147483648}
                                        switch "" {
                                        case "foo bar":
                                            v5 = append(v5, fi13)
                                        case "\n\t":
                                            v6 += v7
                                        case "hello":
                                            v8 += v3
                                        }
                                        _ = v16
                                    }
                                case 182:
                                    if (!true) {
                                        v8 -= 849
                                        v6 -= 190
                                        v17 := GenS1{F0: true, F1: 4294967295}
                                        _ = v17
                                    } else {
                                        if (!true) {
                                            if (true || false) {
                                                _ = len(v5)
                                            }
                                            v18 := GenS1{F0: false, F1: 247}
                                            _ = v18
                                        }
                                        fmt.Println(fi13, v7, v2)
                                    }
                                }
                                switch fi13 {
                                case 0:
                                    switch fi14 {
                                    case 941:
                                        var v19 float64 = 3.14159
                                        _ = v19
                                    case 127:
                                        var v20 int32 = 0
                                        _ = v20
                                    }
                                case 491:
                                    if (false && ((!((-(fi13)) >= fi13)) && (v7 < (v7 + (v7 / ((v6) | 1)))))) {
                                        v8 = (^(v8))
                                    } else {
                                        v21 := GenS2{F0: -980, F1: 481}
                                        v22 := GenS1{F0: false, F1: 42}
                                        _ = v21
                                        _ = v22
                                    }
                                default:
                                    if false {
                                        fmt.Println(fi13)
                                        delete(v1, fi14)
                                    }
                                }
                            }
                        }
                        switch ("" + ("foo bar" + ("hello" + "\n\t"))) {
                        case "foo bar":
                            var v23 int32 = -968
                            _ = v23
                        case "\n\t":
                            v6 = (v6 % ((v7) | 1))
                        }
                    }
                }
            }
            var v24 uint64 = 45
            _ = fi14
            _ = v24
        }
        v25 := GenS2{F0: int32(v8), F1: 1}
        v6 = (v6 & 18446744073709551615)
        _ = fi13
        _ = v25
    }
    var v26 int8 = v2
    fmt.Println(v7, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
