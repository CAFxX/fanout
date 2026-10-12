package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS2 struct { F0 bool }
func genFn0() int64 {
    var v6 uint64 = 18446744073709551615
    v7 := []int{-2147483648}
    _ = v6
    _ = v7
    return 731
}
func genFn1(p0 uint16) string {
    fmt.Println(p0)
    if (("a" + ((("" + "") + ("foo bar" + ("foo bar" + "a"))) + ("a" + "hello"))) >= ("" + "a")) {
        p0 = 1
        switch 0 {
        case 32768:
            for fi8 := 0; fi8 < 20; fi8++ {
                if true {
                    for fi9 := 0; fi9 < 15; fi9++ {
                        var v10 float32 = -2.25
                        switch ("" + "hello") {
                        case "":
                            if (((p0 - (p0 % (((^(p0))) | 1))) != (p0 & p0)) && true) {
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    v11 := GenS1{F0: p0, F1: 1, F2: 790}
                                    _ = v11
                                case "a":
                                    switch "a" {
                                    case "":
                                        v10 = v10
                                    case "foo bar":
                                        v10 = (v10 / (((v10 + 0.0) / (1.5 + 1e10)) * v10))
                                    }
                                default:
                                    var v12 uint16 = 65535
                                    _ = v12
                                }
                                var v13 int16 = -1
                                _ = v13
                            }
                        case "foo bar":
                            if false {
                                v14 := []int{fi8, fi9}
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                _ = v14
                            }
                        }
                        var v15 int8 = 1
                        _ = fi9
                        _ = v10
                        _ = v15
                    }
                } else {
                    fmt.Println(fi8)
                    p0 = p0
                }
                _ = fi8
            }
        case 1000:
            var v16 int32 = -1
            _ = v16
        }
    } else {
        var v17 float64 = 0.1
        var v18 uint = 42
        fmt.Println(p0, p0)
        _ = v17
        _ = v18
    }
    _ = p0
    return "hello"
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    var v3 uint32 = 4294967295
    var v4 int8 = 0
    var v5 float64 = float64(v1)
    v19 := GenS1{F0: 1000, F1: v3, F2: (^(v3))}
    if (v4 >= -128) {
        if false {
            v1 -= 0.1
            v20 := GenS1{F0: 65535, F1: (v3 >> ((681) & 63)), F2: 295}
            var v21 float32 = (1e10 + v1)
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 17; fi22++ {
            v4 = 1
            _ = fi22
        }
    }
    v3 = (^(v3))
    var v23 uint32 = v3
    for fi24 := 0; fi24 < 3; fi24++ {
        v3 -= v3
        if ((fi24 >> ((0) & 63)) > fi24) {
            v5 -= v5
            var v25 uint64 = 4294967295
            v23 = v23
            _ = v25
        }
        v23 += 1
        _ = fi24
    }
    fmt.Println(v5, v5, v1)
    var v26 int32 = -285
    if false {
        if (219 == 168) {
            switch 42 {
            case -9223372036854775808:
                v27 := GenS1{F0: 411, F1: 65535, F2: v23}
                _ = v27
            default:
                switch "foo bar" {
                case "":
                    v28 := GenS2{F0: (46 != 42)}
                    _ = v28
                case "\n\t":
                    fmt.Println(v23, v5, v4)
                case "hello":
                    for fi29 := 0; fi29 < 9; fi29++ {
                        v30 := []int{(fi29 ^ -2147483648)}
                        v31 := GenS2{F0: true}
                        fmt.Println(fi29, v1, v23)
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                default:
                    var v32 float64 = 1.5
                    _ = v32
                }
            }
            var v33 int16 = 619
            _ = v33
        }
        if true {
            v26 -= v26
            v23 -= v3
        }
    }
    v34 := []int{-9223372036854775808, 127, 828}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
