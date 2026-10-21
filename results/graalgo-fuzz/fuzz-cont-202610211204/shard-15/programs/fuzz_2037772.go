package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int64 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint16) uint16 {
    v6 := []int{127, 32768, 32768}
    v7 := GenS1{F0: 42, F1: 1, F2: 190}
    v6 = append(v6, 0)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (^(p1))
}
func genFn1(p0 int16) int8 {
    p0 += p0
    var v8 int32 = 2147483647
    p0 += -97
    p0 += (p0 >> ((2) & 63))
    var v9 int32 = -1
    _ = p0
    _ = v8
    _ = v9
    return 100
}

func main() {
    var v1 float32 = ((1e10 / -2.25) * ((3.14159 + 0.0) * 1e10))
    var v2 uint32 = 42
    v3 := make(map[string]uint32)
    v4 := []string{"hello"}
    var v5 int8 = 127
    v1 = 0.1
    var v10 uint64 = 42
    v10 = v10
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int16 = 42
        v5 += genFn1(v12)
        _ = fi11
        _ = v12
    }
    if (!false) {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(v1, v10, v1)
            _ = fi13
        }
        for fi14 := 0; fi14 < 1; fi14++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (false || (((v1 / 1.5) / v1) != (v1 / (v1 + 1e10)))) {
                v15 := v3["\n\t"]
                var v16 uint16 = 255
                var v17 int16 = 42
                _ = v15
                _ = v16
                _ = v17
            } else {
                if (255 > 1) {
                    var v18 int8 = (v5 / (((v5 % ((v5) | 1))) | 1))
                    _ = v18
                }
                v19 := GenS2{F0: 1}
                if len(v4) > 0 {
                    v4[0] = (("\n\t" + "") + "a")
                }
                _ = v19
            }
            _ = fi14
        }
    }
    var v20 int = 42
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            v2 = (v2 | 1)
            _ = fi22
        }
        v2 = (^(v2))
        _ = fi21
    }
    v10 = (v10 | 0)
    var v23 int64 = 9223372036854775807
    var v24 uint16 = 0
    fmt.Println(v10, v20, v20)
    var v25 uint64 = v10
    var v26 bool = true
    v27 := GenS2{F0: -9223372036854775808}
    v25 = 0
    v20 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
