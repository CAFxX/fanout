package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 int16 }
func genFn0() uint {
    v5 := []int{54, -9223372036854775808, -391, -323}
    var v6 bool = false
    if false {
        var v7 uint16 = 42
        var v8 uint16 = v7
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 float64 = ((3.14159 + -2.25) * float64(3.14159))
            switch fi9 {
            case 9223372036854775807:
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
            case 646:
                switch (-(fi9)) {
                case -1:
                    fmt.Println(v7, v10, fi9)
                default:
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 int64 = 1
                        var v13 float64 = v10
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
            case 898:
                v10 -= (v10 * (float64((((0.1 - 1.5) + (1.5 / (3.14159 + 1.5))) / ((1.5 / 3.14159) / 1.5))) / (0.0 * (v10 + v10))))
            default:
                var v14 uint64 = 2
                _ = v14
            }
            v7 += 1
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        v15 := GenS1{F0: 9223372036854775807, F1: 1, F2: 128}
        _ = len(v5)
        _ = v15
    }
    switch (("hello" + "a") + ("\n\t" + "hello")) {
    case "":
        v16 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807, F2: 223}
        _ = v16
    }
    _ = v5
    _ = v6
    return 609
}

func main() {
    var v1 float32 = 3.14159
    var v2 bool = ("hello" != "a")
    v3 := make(map[string]uint32)
    var v4 uint = 1
    for fi17 := 0; fi17 < 19; fi17++ {
        v2 = (-1 <= -106)
        _ = fi17
    }
    delete(v3, "\n\t")
    v4 = (^(v4))
    switch "\n\t" {
    case "a":
        v1 -= 0.0
    case "foo bar":
        v2 = v2
    case "\n\t":
        v1 += v1
    default:
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 float64 = 3.14159
            for fi20 := 0; fi20 < 3; fi20++ {
                if false {
                    v21 := GenS2{F0: 42, F1: fi18, F2: -7}
                    if v2 {
                        v3["foo bar"] = 42
                        if v2 {
                            var v22 uint8 = 127
                            switch 2147483647 {
                            case -2147483648:
                                var v23 uint64 = 18446744073709551615
                                _ = v23
                            default:
                                switch "" {
                                case "hello":
                                    if (v2 || (fi20 < fi18)) {
                                        switch fi20 {
                                        case 127:
                                            fmt.Println(v4)
                                            v3["\n\t"] = 2147483647
                                            v4 += (v4 - v4)
                                        case 9223372036854775807:
                                            v24 := GenS2{F0: 42, F1: (fi20 / (((fi20 / ((646) | 1))) | 1)), F2: 717}
                                            _ = v24
                                        }
                                    }
                                    delete(v3, ("\n\t" + "\n\t"))
                                    v25 := []int{fi18, (^(fi20)), (fi18 - -2147483648), fi18}
                                    _ = v25
                                }
                                v2 = v2
                                var v26 int8 = -128
                                _ = v26
                            }
                            _ = v22
                        }
                    } else {
                        v4 -= 4294967295
                        v2 = ((2147483648 < 2) && v2)
                        var v27 uint16 = 0
                        _ = v27
                    }
                    v28 := []int{fi20, fi20}
                    if true {
                        var v29 uint8 = 42
                        var v30 float64 = (v19 / v19)
                        _ = v29
                        _ = v30
                    }
                    _ = v21
                    _ = v28
                }
                var v31 float64 = v19
                if true {
                    if (false || (v2 || (!(9223372036854775807 > -9223372036854775808)))) {
                        fmt.Println(v19, v31, v31)
                        fmt.Println(v31)
                        v1 = (v1 / (v1 - ((float32(v31) * v1) - v1)))
                    }
                }
                _ = fi20
                _ = v31
            }
            _ = fi18
            _ = v19
        }
    }
    v2 = (!(v2 && v2))
    fmt.Println(v4, v2)
    var v32 float32 = v1
    v33 := []int{757, 127, 9223372036854775807}
    var v34 uint8 = 1
    var v35 uint64 = 484
    switch "foo bar" {
    case "a":
        for fi36 := 0; fi36 < 17; fi36++ {
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            var v37 int16 = 1000
            var v38 float64 = 0.1
            _ = fi36
            _ = v37
            _ = v38
        }
    default:
        v2 = (94 >= -9223372036854775808)
    }
    var v39 int16 = 471
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v4
}
