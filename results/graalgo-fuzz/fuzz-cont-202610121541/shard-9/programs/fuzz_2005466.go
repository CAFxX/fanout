package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 string }

func main() {
    var v1 int32 = 692
    var v2 uint16 = 0
    var v3 bool = ((-9223372036854775808 < -318) || true)
    var v4 float32 = float32(3.14159)
    var v5 int8 = -1
    v6 := make(map[uint32]bool)
    if v3 {
        v7 := []int{183, int(v5)}
        v8 := GenS1{F0: 255}
        _ = v7
        _ = v8
    } else {
        v2 -= 1
        _ = len(v6)
        if ((v2 % ((2) | 1)) <= (v2 ^ (v2 ^ (v2 + v2)))) {
            if (((-2.25 / 0.1) + ((((0.1 + 0.0) - (-2.25 * (1.5 / -2.25))) / -2.25) / ((-2.25 + 1e10) - float64(v4)))) == 3.14159) {
                fmt.Println(v4, v2, v1)
                v2 += (^(v2))
            } else {
                if false {
                    if ((v1 > (v1 - -1)) || (float32(0.0) == float32((-2.25 + (float64(v4) - -2.25))))) {
                        fmt.Println(v3)
                        var v9 float32 = float32(0.1)
                        _ = v9
                    }
                    var v10 int8 = (v5 - (v5 >> ((42) & 63)))
                    fmt.Println(v10)
                    _ = v10
                }
                for fi11 := 0; fi11 < 7; fi11++ {
                    v12 := GenS1{F0: 42}
                    v13 := GenS1{F0: 2}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            v14 := GenS2{F0: 255, F1: "hello"}
            _ = v14
        }
    }
    v15 := []int{-9223372036854775808, 9223372036854775807, 42, 42}
    var v16 int16 = 25
    var v17 float32 = v4
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v16 = v16
            v1 += (v1 + (^(v1)))
            var v20 uint32 = 527
            if true {
                if (uint64(v20) > 21) {
                    if (!v3) {
                        if true {
                            var v21 uint64 = 1
                            fmt.Println(v4, fi18, v5)
                            _ = v21
                        }
                        switch "a" {
                        case "":
                            fmt.Println(fi19, v5, v1)
                        case "\n\t":
                            var v22 int32 = v1
                            _ = v22
                        default:
                            v23 := v6[v20]
                            _ = v23
                        }
                        if (!(v2 > (v2 | (v2 - v2)))) {
                            if v3 {
                                v24 := GenS1{F0: 160}
                                _ = v24
                            } else {
                                v15 = append(v15, fi19)
                                v16 = v16
                            }
                        }
                        var v25 uint32 = v20
                        var v26 uint16 = v2
                        v27 := GenS2{F0: 172, F1: ""}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    delete(v6, v20)
                }
            }
            _ = fi19
            _ = v20
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            fmt.Println(v4, v16, v3)
            v1 = -661
            _ = fi28
        }
        _ = fi18
    }
    v1 += v1
    var v29 int16 = v16
    fmt.Println(v5, v17, v4)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
