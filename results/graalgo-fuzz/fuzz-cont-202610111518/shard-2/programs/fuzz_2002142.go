package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int64, p1 uint32) int {
    fmt.Println(p0)
    var v5 int8 = -56
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    var v1 uint = 1
    v2 := make(map[string]uint32)
    v3 := []string{"", "hello", "", "\n\t"}
    v4 := make(map[uint32]string)
    var v6 uint64 = 252
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := v4[4294967295]
        _ = fi7
        _ = v8
    }
    fmt.Println(v6, v1)
    var v9 int64 = 901
    var v10 float32 = 1.5
    v11 := GenS1{F0: 32767}
    v3 = append(v3, "a")
    var v12 int32 = 738
    var v13 int64 = (-(v9))
    v13 -= v9
    switch int(v1) {
    case 127:
        _ = len(v2)
    default:
        for fi14 := 0; fi14 < 16; fi14++ {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            _ = len(v4)
            _ = fi14
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
