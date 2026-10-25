package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    v3 := []string{"", "hello"}
    var v4 float32 = float32((0.0 - 0.1))
    v5 := []uint32{376, 1, 1, 2147483647}
    if (29 >= -101) {
        v6 := GenS1{F0: ("" + "a"), F1: 1, F2: -1}
        _ = v6
    }
    if (true && (0 < 18446744073709551615)) {
        if false {
            if true {
                for fi7 := 0; fi7 < 8; fi7++ {
                    var v8 int = 986
                    _ = fi7
                    _ = v8
                }
            }
        } else {
            var v9 uint64 = 18446744073709551615
            var v10 float32 = (v4 / (float32(0.0) - ((v4 * -2.25) + v4)))
            if len(v5) > 0 {
                v5[0] = 358
            }
            _ = v10
            _ = v9
        }
        var v11 int = -354
        _ = v11
    }
    fmt.Println(v4)
    fmt.Println(v4)
    v1["hello"] = 1
    v4 += v4
    if true {
        v12 := GenS1{F0: "\n\t", F1: 871, F2: 57}
        v4 -= -2.25
        var v13 int32 = 277
        _ = v12
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = (("a" + "") + ("hello" + "\n\t"))
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint32 = 65535
        _ = fi14
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
