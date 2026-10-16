package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint16, p1 uint, p2 int16) int {
    p2 += (p2 * (p2 << ((p1) & 63)))
    var v7 uint32 = 399
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 26
}

func main() {
    v1 := []string{"\n\t", (("hello" + "a") + "a"), "hello", "\n\t"}
    v2 := make(map[string]int)
    var v3 uint64 = 2
    v4 := make(map[uint32]int)
    v5 := make(map[string]string)
    var v6 uint8 = 0
    v9 := GenS1{F0: 0}
    v10 := v5[""]
    var v11 uint32 = 4294967295
    v6 -= v6
    var v12 uint32 = (v11 - uint32(v6))
    var v13 bool = false
    for fi14 := 0; fi14 < 16; fi14++ {
        v1 = append(v1, ((v10 + (((v10 + v10) + v10) + "a")) + v10))
        for fi15 := 0; fi15 < 12; fi15++ {
            v12 = v11
            var v16 float32 = 0.0
            var v17 float32 = 3.14159
            _ = fi15
            _ = v16
            _ = v17
        }
        v12 = (v11 - v11)
        _ = fi14
    }
    v10 = v10
    for fi18 := 0; fi18 < 18; fi18++ {
        if v13 {
            _ = len(v4)
        } else {
            for fi19 := 0; fi19 < 9; fi19++ {
                delete(v5, "foo bar")
                var v20 int8 = 28
                _ = fi19
                _ = v20
            }
            var v21 int = (-(fi18))
            _ = v21
        }
        var v22 int8 = -1
        var v23 uint8 = (v6 - v6)
        _ = fi18
        _ = v22
        _ = v23
    }
    fmt.Println(v10)
    fmt.Println(v12)
    var v24 int64 = -9223372036854775808
    var v25 uint64 = (v3 / ((v3) | 1))
    v11 += 449
    v26 := GenS2{F0: 42, F1: v11}
    var v27 int32 = -374
    if v13 {
        v13 = (v13 || (1 > -592))
        fmt.Println(v3, v24)
    }
    var v28 bool = v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
