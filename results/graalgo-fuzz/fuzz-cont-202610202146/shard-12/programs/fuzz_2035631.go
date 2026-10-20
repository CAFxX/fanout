package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0() int64 {
    var v5 uint32 = 2
    v6 := []int{127, 705, 9223372036854775807, 32767}
    _ = v5
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := []string{"a", "\n\t", "a", (("hello" + "") + "\n\t")}
    var v2 bool = (0 >= 172)
    v3 := []string{("" + "hello"), "hello", "foo bar", "hello"}
    var v4 int64 = -1
    switch 433 {
    case 4294967295:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v4 += v4
    v7 := GenS1{F0: 199, F1: -964}
    if (!(!(127 > -128))) {
        fmt.Println(v2, v2)
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 float32 = 0.0
            v4 += v4
            _ = fi8
            _ = v9
        }
    }
    v2 = (((1.5 * 3.14159) + 3.14159) != 0.0)
    var v10 uint64 = 4294967295
    if false {
        var v11 string = "hello"
        var v12 int32 = -627
        _ = v11
        _ = v12
    }
    if v2 {
        switch "" {
        case "\n\t":
            for fi13 := 0; fi13 < 17; fi13++ {
                v4 = v4
                _ = fi13
            }
        case "a":
            for fi14 := 0; fi14 < 9; fi14++ {
                if (32767 <= 1) {
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v1 = append(v1, "")
                        if (!(-59 != -35)) {
                            fmt.Println(v2, v2)
                            var v16 int32 = -730
                            v10 += v10
                            v17 := GenS2{F0: (v4 & 1), F1: ((v10 / ((870) | 1)) <= 622)}
                            _ = v16
                            _ = v17
                        } else {
                            var v18 float64 = 1.5
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                            _ = v18
                        }
                        var v19 int32 = -355
                        _ = fi15
                        _ = v19
                    }
                } else {
                    v20 := GenS2{F0: v4, F1: false}
                    var v21 float64 = 1.5
                    _ = v20
                    _ = v21
                }
                if true {
                    switch "" {
                    case "a":
                        var v22 string = "\n\t"
                        _ = v22
                    case "hello":
                        v4 = v4
                    default:
                        if (!v2) {
                            v4 -= v4
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                        }
                    }
                    var v23 bool = true
                    var v24 uint16 = 530
                    _ = v23
                    _ = v24
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi14
            }
        case "hello":
            v2 = (127 <= 255)
        default:
            var v25 int16 = 32767
            _ = v25
        }
    }
    var v26 uint8 = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v7
}
