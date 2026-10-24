package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int32 }
type GenS2 struct { F0 bool; F1 string; F2 int8 }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0() uint32 {
    var v4 bool = true
    v4 = (!v4)
    v5 := []int{128, -9223372036854775808, 42}
    v5 = append(v5, -2147483648)
    _ = v4
    _ = v5
    return 2147483648
}

func main() {
    var v1 int32 = 1
    var v2 int32 = (v1 + v1)
    var v3 int32 = v2
    switch "foo bar" {
    case "foo bar":
        var v6 int16 = 1
        _ = v6
    case "hello":
        v7 := []int{-558, -533, 859, 2147483647}
        _ = v7
    case "a":
        var v8 float32 = (1.5 - 0.0)
        _ = v8
    }
    v9 := []int{32767, -173, 32767}
    var v10 uint64 = 0
    v11 := GenS1{F0: 2, F1: 128, F2: (v2 + 2147483647)}
    fmt.Println(v10)
    v9 = append(v9, 127)
    v1 = v1
    if (!(-47 < -43)) {
        if (!(("\n\t" + "foo bar") >= ("" + ("foo bar" + ("foo bar" + ("" + "foo bar")))))) {
            if ((v1 - v1) == (v2 * v3)) {
                v2 = v3
                var v12 uint64 = (v10 * v10)
                _ = v12
            } else {
                var v13 string = "foo bar"
                _ = v13
            }
            v14 := GenS2{F0: ((1.5 == (float64(((1.5 / 3.14159) / 1.0)) - 1.5)) || false), F1: ("foo bar" + ""), F2: 42}
            for fi15 := 0; fi15 < 7; fi15++ {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                for fi16 := 0; fi16 < 3; fi16++ {
                    v17 := GenS1{F0: 640, F1: fi15, F2: v3}
                    fmt.Println(fi15, fi15)
                    v18 := GenS2{F0: true, F1: "foo bar", F2: 127}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                switch "\n\t" {
                case "hello":
                    v19 := GenS1{F0: 0, F1: fi15, F2: v1}
                    _ = v19
                case "":
                    v2 = (v3 >> ((2) & 63))
                case "\n\t":
                    for fi20 := 0; fi20 < 14; fi20++ {
                        v3 = (v3 - (v2 ^ (v2 ^ (v1 - (v3 * (v3 - v1))))))
                        switch fi15 {
                        case 32767:
                            v21 := GenS2{F0: ((true || (v10 > v10)) && ((fi20 >> ((0) & 63)) == (fi15 % ((319) | 1)))), F1: "hello", F2: -49}
                            _ = v21
                        case -731:
                            switch (fi15 << ((534) & 63)) {
                            case 42:
                                v3 -= v3
                                if (0 > -98) {
                                    fmt.Println(v3, fi20, v10)
                                    v10 = v10
                                } else {
                                    v10 = (^(v10))
                                }
                            case 128:
                                var v22 uint16 = 0
                                _ = v22
                            }
                        default:
                            var v23 uint16 = 1
                            _ = v23
                        }
                        _ = fi20
                    }
                default:
                    for fi24 := 0; fi24 < 4; fi24++ {
                        fmt.Println(v10, fi15, fi15)
                        _ = fi24
                    }
                }
                _ = fi15
            }
            _ = v14
        }
        var v25 int64 = 869
        _ = v25
    }
    var v26 float32 = 3.14159
    for fi27 := 0; fi27 < 7; fi27++ {
        if true {
            v28 := GenS3{F0: 0, F1: -9223372036854775808}
            _ = v28
        } else {
            fmt.Println(v10)
            var v29 string = "\n\t"
            v1 -= -791
            _ = v29
        }
        if (559 > -751) {
            v30 := GenS3{F0: -582, F1: 9223372036854775807}
            for fi31 := 0; fi31 < 20; fi31++ {
                if true {
                    v3 -= -2147483648
                    switch "foo bar" {
                    case "a":
                        fmt.Println(fi27)
                    case "":
                        fmt.Println(v3, fi27, v26)
                    default:
                        if ((v26 >= v26) && false) {
                            v26 += 3.14159
                            fmt.Println(v10, v2)
                        }
                    }
                    v32 := GenS2{F0: true, F1: "", F2: 100}
                    _ = len(v9)
                    _ = v32
                } else {
                    var v33 int = (fi31 & -750)
                    _ = v33
                }
                _ = fi31
            }
            _ = v30
        } else {
            var v34 int16 = 1000
            _ = v34
        }
        _ = fi27
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v35 := GenS3{F0: 1000, F1: -1}
    if len(v9) > 0 {
        v9[0] = -498
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v9
}
