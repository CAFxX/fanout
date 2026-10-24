package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint64; F1 int32 }
func genFn0() int8 {
    var v6 int16 = 55
    v6 = v6
    v6 -= (v6 * v6)
    _ = v6
    return 127
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{("a" + ("foo bar" + "a"))}
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 int8 = 1
    var v7 float32 = (-2.25 + 0.0)
    var v8 uint32 = 2
    switch (("a" + "") + "hello") {
    case "foo bar":
        var v9 int = -705
        _ = v9
    }
    if (42 <= 42) {
        v10 := GenS1{F0: v5, F1: (v8 >> ((18446744073709551615) & 63)), F2: v8}
        _ = len(v2)
        _ = v10
    } else {
        var v11 string = "\n\t"
        if true {
            v12 := GenS2{F0: "hello", F1: (v5 | (-(v5)))}
            var v13 string = ("\n\t" + ("a" + (v11 + (v11 + v11))))
            _ = v12
            _ = v13
        }
        switch (v11 + (v11 + v11)) {
        case "hello":
            switch (v11 + ((v11 + (v11 + ("\n\t" + v11))) + "a")) {
            case "\n\t":
                var v14 int64 = -9223372036854775808
                _ = v14
            case "hello":
                v15 := GenS1{F0: v5, F1: (v8 << ((4294967295) & 63)), F2: v8}
                _ = v15
            }
        case "a":
            if true {
                for fi16 := 0; fi16 < 2; fi16++ {
                    v8 += v8
                    if true {
                        switch fi16 {
                        case 128:
                            fmt.Println(v11)
                        case -744:
                            delete(v4, 2147483647)
                        default:
                            fmt.Println(v8, v8, fi16)
                        }
                        if true {
                            if (v11 != ((v11 + v11) + "")) {
                                v17 := GenS3{F0: -2147483648, F1: ((1e10 * (-2.25 / 1.0)) == 3.14159)}
                                _ = v17
                            }
                            v8 -= (v8 / (((v8 >> ((651) & 63))) | 1))
                            v18 := GenS4{F0: fi16, F1: 2147483647}
                            _ = v18
                        } else {
                            if false {
                                var v19 uint64 = 745
                                if (2 <= 65535) {
                                    fmt.Println(v5, v11, fi16)
                                } else {
                                    v7 = 3.14159
                                    var v20 bool = (false && (v5 == 36))
                                    v8 = 416
                                    _ = v20
                                }
                                v21 := GenS2{F0: ("a" + (v11 + v11)), F1: v5}
                                _ = v19
                                _ = v21
                            }
                            _ = len(v3)
                        }
                    } else {
                        v2 = append(v2, "")
                    }
                    _ = fi16
                }
            }
        case "":
            v22 := GenS5{F0: 1, F1: -488}
            _ = v22
        default:
            v11 = "foo bar"
        }
        _ = v11
    }
    v5 = v5
    v23 := GenS2{F0: "hello", F1: (v5 ^ (v5 & 100))}
    var v24 uint32 = (v8 / ((0) | 1))
    v25 := GenS1{F0: (-(v5)), F1: (v8 + (^(v8))), F2: (^(v8))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        v8 += (^(v24))
        _ = len(v2)
        var v27 uint64 = 4294967295
        _ = fi26
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
