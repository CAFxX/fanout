package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int8, p1 int32) uint64 {
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            p1 = p1
            if (88 > 1) {
                v9 := GenS1{F0: 42, F1: "a"}
                v10 := GenS1{F0: 42, F1: "a"}
                p1 += p1
                p1 -= (p1 >> ((348) & 63))
                var v11 int64 = 450
                if true {
                    p0 += (p0 >> ((65535) & 63))
                    v12 := GenS2{F0: 32767, F1: ((fi8 / (((fi8 / (((fi7 / ((2147483647) | 1))) | 1))) | 1)) <= fi8), F2: p0}
                    _ = v12
                } else {
                    var v13 int16 = 1000
                    if (2 == 251) {
                        if false {
                            fmt.Println(fi7)
                            var v14 int16 = (v13 % (((v13 << ((42) & 63))) | 1))
                            _ = v14
                        } else {
                            var v15 float32 = (1.5 - 1.5)
                            _ = v15
                        }
                    }
                    _ = v13
                }
                p0 = (p0 & 0)
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi8
        }
        p0 = (-(p0))
        _ = fi7
    }
    var v16 uint64 = 42
    var v17 uint8 = 183
    p1 = p1
    var v18 float64 = 1e10
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return (v16 - 18446744073709551615)
}

func main() {
    var v1 int8 = 1
    var v2 bool = (!((594 != -9223372036854775808) || false))
    v3 := make(map[uint32]int)
    v4 := []string{"foo bar"}
    var v5 int32 = 0
    v6 := make(map[string]string)
    v19 := GenS1{F0: 2, F1: "\n\t"}
    for fi20 := 0; fi20 < 3; fi20++ {
        v4 = append(v4, "foo bar")
        _ = fi20
    }
    fmt.Println(v5, v1, v2)
    delete(v6, "\n\t")
    v21 := GenS3{F0: 32767}
    fmt.Println(v2, v5)
    if (0.1 < (1e10 * float64((0.1 - (-2.25 * (0.1 / 1e10)))))) {
        fmt.Println(v1, v1, v1)
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                v2 = v2
            } else {
                if ((v1 != (v1 - (v1 & 127))) && (42 != 128)) {
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                } else {
                    if (!((-2.25 + (1e10 + 1.5)) <= 1e10)) {
                        var v22 int = 42
                        _ = v22
                    }
                }
            }
            v23 := GenS2{F0: -1, F1: (v2 && false), F2: v1}
            _ = v23
        }
        var v24 uint8 = 255
        _ = v24
    } else {
        v1 -= v1
        var v25 int8 = (v1 * v1)
        fmt.Println(v1, v25, v25)
        _ = v25
    }
    var v26 int16 = 1
    var v27 int8 = -94
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
