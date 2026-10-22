package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 uint64, p1 bool, p2 int8) int16 {
    fmt.Println(p2, p1, p2)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := []int{fi7, 1, (fi7 + (fi7 << ((659) & 63))), (fi7 >> ((745) & 63))}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint16 = 32767
    var v2 int = 9223372036854775807
    v3 := make(map[string]uint32)
    var v4 uint64 = 0
    var v5 bool = false
    v6 := []uint32{}
    v9 := v3["a"]
    var v10 uint8 = 42
    v5 = v5
    _ = len(v6)
    _ = len(v6)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v5, v5, v2)
    var v11 int64 = 0
    _ = len(v3)
    var v12 int32 = -1
    for fi13 := 0; fi13 < 12; fi13++ {
        if (!(600 <= 1)) {
            if ((v12 % ((v12) | 1)) >= 42) {
                switch v2 {
                case 127:
                    if v5 {
                        v11 -= 1
                    } else {
                        for fi14 := 0; fi14 < 17; fi14++ {
                            fmt.Println(v11, v5)
                            fmt.Println(v2)
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = fi14
                        }
                        delete(v3, "a")
                        delete(v3, "foo bar")
                    }
                case 32768:
                    v15 := GenS1{F0: uint(v4), F1: int32(v1)}
                    _ = v15
                }
                if (9223372036854775807 == int(v11)) {
                    switch "" {
                    case "":
                        if (65535 != 1) {
                            var v16 float32 = (-2.25 + 0.0)
                            v2 = (-(v2))
                            v6 = append(v6, v9)
                            _ = v16
                        }
                    }
                } else {
                    var v17 float32 = (1.5 * 0.0)
                    _ = len(v6)
                    if (uint(v4) == 846) {
                        fmt.Println(v5)
                        var v18 uint = 2
                        var v19 uint64 = (^(v4))
                        _ = v18
                        _ = v19
                    }
                    _ = v17
                }
                v5 = ((v1 - (v1 | 65535)) == v1)
            }
            switch -1 {
            case 2147483647:
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 float32 = (0.0 / (((3.14159 + float32((3.14159 / (-2.25 * 1e10)))) / 0.1) / float32(0.1)))
                    var v22 uint16 = 65535
                    var v23 int8 = 42
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
            default:
                for fi24 := 0; fi24 < 12; fi24++ {
                    var v25 uint16 = (v1 % (((v1 + (v1 & (v1 % ((v1) | 1))))) | 1))
                    v9 += (^(v9))
                    _ = fi24
                    _ = v25
                }
            }
        }
        if v5 {
            v6 = append(v6, (v9 | v9))
        }
        _ = fi13
    }
    v26 := GenS1{F0: 15, F1: -1}
    if v5 {
        fmt.Println(v5, v10, v4)
        v5 = v5
    }
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 int = v2
        var v29 uint64 = (v4 | (v4 / ((v4) | 1)))
        v6 = append(v6, v9)
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := GenS1{F0: uint(v10), F1: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
