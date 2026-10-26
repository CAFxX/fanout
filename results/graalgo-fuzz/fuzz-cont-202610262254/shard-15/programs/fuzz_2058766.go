package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    v1 := []uint32{}
    var v2 uint64 = 0
    v3 := make(map[int]uint32)
    var v4 uint16 = 0
    fmt.Println(v2)
    if true {
        var v5 int = 0
        v5 += 654
        _ = v5
    }
    v6 := GenS1{F0: 1, F1: -49}
    if false {
        switch "\n\t" {
        case "foo bar":
            var v7 int32 = -1
            _ = v7
        }
        fmt.Println(v2, v4)
    } else {
        _ = len(v3)
        fmt.Println(v4, v4)
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        switch (("a" + ("hello" + ("a" + "\n\t"))) + ("a" + "")) {
        case "\n\t":
            if false {
                v9 := GenS1{F0: -639, F1: -128}
                _ = v9
            }
        case "a":
            var v10 float32 = ((1e10 * (3.14159 - 0.0)) + 1e10)
            _ = v10
        case "hello":
            fmt.Println(fi8)
        }
        var v11 uint16 = 940
        _ = fi8
        _ = v11
    }
    fmt.Println(v2, v4)
    if (!(v2 > v2)) {
        for fi12 := 0; fi12 < 8; fi12++ {
            fmt.Println(v4)
            var v13 bool = false
            if (true && (true && v13)) {
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 uint8 = 1
                    var v16 bool = v13
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                v13 = (v13 && v13)
                switch fi12 {
                case 9223372036854775807:
                    switch (fi12 >> ((612) & 63)) {
                    case -1:
                        if (false || (false && v13)) {
                            if true {
                                var v17 uint8 = 1
                                v18 := GenS1{F0: 1, F1: 1}
                                var v19 int64 = 819
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                            v2 = (v2 | v2)
                            switch 0 {
                            case -795:
                                fmt.Println(v2)
                            default:
                                v20 := v3[fi12]
                                _ = v20
                            }
                        } else {
                            if v13 {
                                v21 := GenS2{F0: ("a" + "\n\t"), F1: v2}
                                v22 := GenS2{F0: "\n\t", F1: (v2 & v2)}
                                var v23 uint8 = 2
                                _ = v21
                                _ = v22
                                _ = v23
                            } else {
                                if ((127 != 0) && false) {
                                    v2 -= (^(v2))
                                } else {
                                    v2 -= 202
                                    for fi24 := 0; fi24 < 3; fi24++ {
                                        v13 = (v13 || true)
                                        v13 = v13
                                        _ = fi24
                                    }
                                }
                            }
                        }
                    case -9223372036854775808:
                        if (0 == 4294967295) {
                            switch (("hello" + (("foo bar" + "") + ("foo bar" + ("" + "hello")))) + ((("" + "") + "hello") + ("a" + ""))) {
                            case "foo bar":
                                var v25 uint64 = 18446744073709551615
                                _ = v25
                            case "hello":
                                fmt.Println(v2)
                            case "a":
                                if (!false) {
                                    v4 = (v4 / (((v4 * v4)) | 1))
                                    fmt.Println(v4, fi12)
                                }
                            default:
                                switch (("a" + ("" + "a")) + "foo bar") {
                                case "":
                                    fmt.Println(v2, v2)
                                case "hello":
                                    if false {
                                        for fi26 := 0; fi26 < 18; fi26++ {
                                            v4 += v4
                                            _ = len(v3)
                                            v4 += 0
                                            _ = fi26
                                        }
                                    } else {
                                        v3[fi12] = 2147483648
                                    }
                                default:
                                    var v27 float64 = -2.25
                                    _ = v27
                                }
                            }
                            v13 = (1 >= 100)
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                        }
                    }
                case -325:
                    for fi28 := 0; fi28 < 7; fi28++ {
                        v29 := GenS2{F0: "foo bar", F1: (v2 << ((143) & 63))}
                        _ = fi28
                        _ = v29
                    }
                default:
                    for fi30 := 0; fi30 < 5; fi30++ {
                        _ = len(v3)
                        _ = fi30
                    }
                }
            }
            _ = fi12
            _ = v13
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (1 != 0) {
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 uint64 = (^(v2))
            _ = fi31
            _ = v32
        }
        v4 = v4
    }
    v2 = v2
    fmt.Println(v4, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
