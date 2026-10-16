package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
func genFn0() int64 {
    var v5 uint = 4294967295
    fmt.Println(v5)
    v5 = 788
    v6 := GenS1{F0: (v5 & 40), F1: 42}
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int, p1 uint8, p2 bool) int8 {
    p1 = (p1 - p1)
    var v7 int = -49
    switch 1 {
    case 42:
        v8 := []int{}
        _ = v8
    case -1:
        fmt.Println(p0, v7)
    }
    p2 = (p2 || (!p2))
    p0 += p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]string)
    var v3 int = -2147483648
    v4 := make(map[int]bool)
    var v9 uint64 = 5
    fmt.Println(v9)
    v3 = (-(v3))
    var v10 int32 = -990
    var v11 int32 = (v10 & v10)
    v12 := []int{v3, int(v10), (v3 / ((9223372036854775807) | 1))}
    if (!false) {
        v12 = append(v12, int(v11))
    } else {
        if true {
            if false {
                for fi13 := 0; fi13 < 2; fi13++ {
                    var v14 uint16 = 766
                    _ = len(v4)
                    if false {
                        var v15 uint64 = (v9 | v9)
                        for fi16 := 0; fi16 < 15; fi16++ {
                            v17 := v4[9223372036854775807]
                            v10 += v10
                            _ = fi16
                            _ = v17
                        }
                        _ = v15
                    }
                    _ = fi13
                    _ = v14
                }
            } else {
                switch "foo bar" {
                case "\n\t":
                    for fi18 := 0; fi18 < 5; fi18++ {
                        var v19 int64 = 360
                        if len(v12) > 0 {
                            v12[0] = v3
                        }
                        _ = fi18
                        _ = v19
                    }
                default:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v10 = v10
                        var v21 int64 = -353
                        var v22 int = v3
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            }
            var v23 uint64 = v9
            _ = v23
        } else {
            switch "" {
            case "hello":
                for fi24 := 0; fi24 < 6; fi24++ {
                    if len(v12) > 0 {
                        v12[0] = 552
                    }
                    for fi25 := 0; fi25 < 13; fi25++ {
                        var v26 int = (fi24 ^ -868)
                        delete(v4, -2147483648)
                        _ = fi25
                        _ = v26
                    }
                    if true {
                        v10 = (v10 ^ 0)
                    } else {
                        var v27 float64 = 1e10
                        _ = v27
                    }
                    _ = fi24
                }
            default:
                fmt.Println(v10, v11, v10)
            }
            delete(v2, 623)
        }
        v9 = (v9 * (v9 - (v9 + v9)))
    }
    var v28 uint32 = 0
    v29 := GenS2{F0: uint16(v28), F1: 469, F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
