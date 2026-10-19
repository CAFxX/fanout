package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int64 = 436
    var v2 string = ("foo bar" + "\n\t")
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    switch 4294967295 {
    case -416:
        switch 127 {
        case -55:
            var v5 int64 = (v1 << ((255) & 63))
            _ = v5
        default:
            fmt.Println(v1, v2)
        }
    case -724:
        var v6 uint32 = 844
        _ = v6
    case -9223372036854775808:
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 int8 = 11
            _ = fi7
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(v1, v1, v1)
            _ = fi9
        }
    }
    v10 := GenS1{F0: 127, F1: 1}
    var v11 int32 = -1
    v3[v2] = 2
    v12 := v3["hello"]
    switch v2 {
    case "":
        var v13 bool = ((!(572 < 32768)) || true)
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS1{F0: 1, F1: fi14}
            for fi16 := 0; fi16 < 7; fi16++ {
                v11 = -2147483648
                var v17 uint64 = 0
                var v18 float32 = -2.25
                _ = fi16
                _ = v17
                _ = v18
            }
            if (((-2.25 / (float64(1e10) * 3.14159)) <= 0.0) && false) {
                var v19 uint8 = 209
                if true {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v1 = v1
                        var v21 uint64 = 42
                        if ((false && true) && (v21 == v21)) {
                            v2 = v2
                            v12 += (v12 ^ 42)
                            if (32767 >= 831) {
                                var v22 int = fi14
                                _ = v22
                            } else {
                                v23 := GenS1{F0: v19, F1: fi20}
                                var v24 int = 127
                                switch 308 {
                                case -1:
                                    v2 = v2
                                case 424:
                                    if (!(((float64(3.14159) * (1e10 / (0.0 - -2.25))) / -2.25) != 0.1)) {
                                        var v25 int64 = v1
                                        _ = len(v4)
                                        var v26 string = (v2 + v2)
                                        v25 = v25
                                        _ = v25
                                        _ = v26
                                    }
                                }
                                _ = v23
                                _ = v24
                            }
                            v27 := GenS2{F0: v19}
                            _ = v27
                        } else {
                            v12 -= 0
                        }
                        _ = fi20
                        _ = v21
                    }
                    v2 = (("" + v2) + "")
                }
                v2 = ("a" + "hello")
                _ = v19
            } else {
                v12 = (v12 + (v12 % ((0) | 1)))
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 bool = (3.14159 <= 3.14159)
                    if v29 {
                        v1 -= 42
                    }
                    _ = fi28
                    _ = v29
                }
                v12 = v12
            }
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v12, v12)
    switch "\n\t" {
    case "":
        v1 = v1
    case "a":
        v30 := GenS1{F0: 120, F1: -2147483648}
        _ = v30
    default:
        var v31 uint32 = (v12 + v12)
        _ = v31
    }
    var v32 int32 = -2147483648
    var v33 string = v2
    v34 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
