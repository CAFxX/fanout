package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint8 }
func genFn0(p0 int64) int64 {
    if (!(-36 < -1)) {
        if (false && true) {
            v7 := GenS1{F0: -202, F1: "hello", F2: 127}
            if (0 > 2) {
                for fi8 := 0; fi8 < 11; fi8++ {
                    fmt.Println(fi8)
                    _ = fi8
                }
            } else {
                if false {
                    var v9 uint8 = 0
                    var v10 float64 = (3.14159 / 0.1)
                    _ = v10
                    _ = v9
                } else {
                    v11 := GenS1{F0: -1, F1: "\n\t", F2: 2}
                    switch "hello" {
                    case "\n\t":
                        var v12 uint = 0
                        _ = v12
                    case "foo bar":
                        v13 := GenS1{F0: -32768, F1: "", F2: 127}
                        _ = v13
                    default:
                        for fi14 := 0; fi14 < 7; fi14++ {
                            for fi15 := 0; fi15 < 3; fi15++ {
                                if (!(32767 != 32767)) {
                                    v16 := []int{fi15, fi14}
                                    _ = v16
                                }
                                p0 += p0
                                v17 := []int{32767, (fi15 / ((fi15) | 1)), (fi15 % (((fi14 - (fi14 / ((fi15) | 1)))) | 1))}
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                _ = fi15
                                _ = v17
                            }
                            fmt.Println(p0, p0)
                            _ = fi14
                        }
                    }
                    _ = v11
                }
                var v18 uint16 = 421
                if (-2.25 <= 0.1) {
                    v19 := []int{838, 0}
                    _ = v19
                }
                _ = v18
            }
            _ = v7
        }
        var v20 int64 = (p0 / ((p0) | 1))
        _ = v20
    } else {
        var v21 uint64 = 894
        p0 -= p0
        _ = v21
    }
    p0 = -242
    var v22 uint8 = 0
    _ = p0
    _ = v22
    return (^(p0))
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]uint32)
    var v3 int = 127
    v4 := []uint32{101}
    var v5 uint = 255
    var v6 int64 = 1
    fmt.Println(v1, v5)
    if false {
        v23 := GenS1{F0: 1, F1: "foo bar", F2: 42}
        _ = v23
    } else {
        fmt.Println(v5)
    }
    fmt.Println(v6)
    var v24 uint16 = 0
    v1 = (v1 << ((v5) & 63))
    v3 += int(v6)
    var v25 int = int(v1)
    fmt.Println(v24, v5, v25)
    v26 := GenS1{F0: int16(v1), F1: "", F2: 236}
    v6 = (v6 << ((v5) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
