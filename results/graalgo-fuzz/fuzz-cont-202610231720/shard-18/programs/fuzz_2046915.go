package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint32 }

func main() {
    var v1 int32 = 503
    var v2 float64 = (1.5 + 0.1)
    var v3 int = -940
    v4 := make(map[string]string)
    v5 := GenS1{F0: 18446744073709551615}
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 822}
            v9 := GenS1{F0: 388}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v10 := []int{}
    var v11 int = v3
    var v12 uint16 = 0
    switch "foo bar" {
    case "":
        if len(v10) > 0 {
            v10[0] = 198
        }
    case "foo bar":
        switch "foo bar" {
        case "foo bar":
            for fi13 := 0; fi13 < 4; fi13++ {
                if len(v10) > 0 {
                    v10[0] = 32767
                }
                _ = fi13
            }
        case "hello":
            v14 := GenS1{F0: 293}
            _ = v14
        }
    case "a":
        var v15 bool = (42 >= 9223372036854775807)
        _ = v15
    default:
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint32 = 65535
            v12 -= v12
            _ = fi16
            _ = v17
        }
    }
    var v18 uint = 1
    var v19 float32 = ((0.1 - 0.1) * 3.14159)
    switch v3 {
    case 2147483647:
        if (100 >= -113) {
            var v20 uint8 = 127
            v21 := GenS2{F0: 219, F1: -807}
            _ = v20
            _ = v21
        } else {
            v3 -= (v11 / ((v3) | 1))
            v22 := GenS3{F0: false}
            _ = v22
        }
    default:
        if (("\n\t" + "\n\t") >= "") {
            for fi23 := 0; fi23 < 2; fi23++ {
                if (!false) {
                    var v24 float32 = (1e10 - (0.1 / v19))
                    switch "hello" {
                    case "a":
                        v25 := GenS1{F0: 18446744073709551615}
                        _ = v25
                    default:
                        switch fi23 {
                        case 881:
                            if false {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                            }
                        default:
                            v1 -= (v1 - (v1 + 42))
                        }
                    }
                    _ = v24
                }
                if (!true) {
                    if false {
                        v26 := GenS2{F0: 1, F1: -9223372036854775808}
                        _ = v26
                    }
                    switch 145 {
                    case 357:
                        var v27 uint64 = 1
                        _ = v27
                    case -19:
                        v2 -= (0.1 - 3.14159)
                    }
                    _ = len(v4)
                }
                _ = fi23
            }
            switch ("" + "a") {
            case "":
                v3 += 32768
            case "hello":
                v2 = v2
            case "\n\t":
                _ = len(v4)
            default:
                switch 9223372036854775807 {
                case 1:
                    var v28 float32 = (v19 - (v19 * (v19 / 0.1)))
                    _ = v28
                case 0:
                    var v29 int16 = 1000
                    _ = v29
                case 9223372036854775807:
                    v30 := GenS2{F0: 922, F1: 0}
                    _ = v30
                }
            }
        }
    }
    v31 := GenS4{F0: (0.0 >= v2), F1: 2147483648}
    var v32 float64 = float64((float32(float64(v19)) / v19))
    var v33 uint8 = 112
    v34 := v4["hello"]
    fmt.Println(v1, v12)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
