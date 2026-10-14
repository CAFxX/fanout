package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
func genFn0(p0 bool) uint {
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 int64 = 1
            fmt.Println(p0, fi6)
            v7 -= (v7 / ((v7) | 1))
            _ = fi6
            _ = v7
        }
        var v8 int32 = 2147483647
        _ = fi5
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return 56
}

func main() {
    var v1 int16 = 42
    v2 := make(map[string]uint32)
    var v3 int16 = 1000
    v4 := make(map[uint32]string)
    var v9 int = 664
    var v10 string = "hello"
    if true {
        v11 := GenS1{F0: ((v10 + (v10 + v10)) + "a")}
        v12 := GenS2{F0: 1}
        _ = v11
        _ = v12
    } else {
        v10 = v10
        v13 := v4[42]
        if true {
            v10 = v10
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS2{F0: -1}
                if ((v10 + ("a" + "")) >= ((("\n\t" + "") + (v13 + v10)) + v10)) {
                    _ = len(v2)
                    var v16 uint64 = 2
                    v17 := GenS1{F0: v10}
                    _ = v16
                    _ = v17
                }
                v1 = v1
                _ = fi14
                _ = v15
            }
        } else {
            v18 := v2[v13]
            _ = v18
        }
        _ = v13
    }
    v3 -= (v3 | (v1 ^ v3))
    v1 += 409
    if (true && false) {
        var v19 uint16 = 32767
        var v20 float64 = ((1e10 + 1.5) - 0.1)
        v21 := []int{127, -333, 1, (v9 | 9223372036854775807)}
        _ = v19
        _ = v20
        _ = v21
    } else {
        fmt.Println(v9, v9)
        fmt.Println(v10, v10, v3)
        _ = len(v2)
    }
    v10 = v10
    v10 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
