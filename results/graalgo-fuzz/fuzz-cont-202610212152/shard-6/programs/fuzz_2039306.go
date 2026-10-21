package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint64) string {
    var v5 uint64 = (p0 | p0)
    v6 := GenS1{F0: 4294967295, F1: -1}
    var v7 int32 = 42
    v5 -= 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    v1 := []int{9223372036854775807, -9223372036854775808, 9223372036854775807}
    var v2 float32 = 0.0
    v3 := make(map[string]bool)
    v4 := []string{"\n\t"}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 float32 = v2
        fmt.Println(v2, v2)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v2 += (((v2 / v2) + (v2 / (float32(float64((((v2 - v2) - v2) / ((0.0 * 0.0) - ((((v2 * v2) + 0.0) + 0.1) - 1e10))))) / 0.1))) / v2)
        _ = fi10
    }
    _ = len(v3)
    if true {
        delete(v3, "")
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (true || true) {
        v11 := GenS2{F0: (1 > 768)}
        var v12 uint16 = 32767
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v11
        _ = v12
    } else {
        v2 = (v2 + (0.1 / 3.14159))
        if (true || ((!false) && (true && true))) {
            var v13 uint = 2
            v14 := GenS2{F0: false}
            v13 += 255
            _ = v13
            _ = v14
        }
        if false {
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 float32 = float32(((-2.25 - (3.14159 / -2.25)) - 3.14159))
                _ = fi15
                _ = v16
            }
        } else {
            fmt.Println(v2)
            v2 -= (v2 / 3.14159)
            if true {
                var v17 uint16 = 0
                v18 := v3["hello"]
                _ = v17
                _ = v18
            }
        }
    }
    v19 := v3[""]
    var v20 float32 = (1.5 / (v2 / -2.25))
    v21 := GenS1{F0: 1, F1: 0}
    var v22 uint64 = 0
    v3[("\n\t" + "a")] = true
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 float32 = (0.1 + (v20 * v20))
        for fi25 := 0; fi25 < 18; fi25++ {
            if (751 > 787) {
                var v26 int16 = -1
                _ = v26
            }
            var v27 int16 = -510
            _ = fi25
            _ = v27
        }
        v2 -= v20
        _ = fi23
        _ = v24
    }
    v2 = float32(float64(0.1))
    var v28 int32 = -667
    v20 = (v20 / v20)
    var v29 int = 32768
    var v30 int = (v29 ^ v29)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
