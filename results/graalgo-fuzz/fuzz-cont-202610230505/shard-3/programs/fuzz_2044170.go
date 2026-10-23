package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
func genFn0() int8 {
    var v4 int16 = -32768
    var v5 int = 4294967295
    var v6 string = "a"
    var v7 int16 = (v4 + v4)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 int64 = -251
    fmt.Println(v3)
    for fi8 := 0; fi8 < 1; fi8++ {
        v3 = 381
        _ = fi8
    }
    v3 += (^(v3))
    v9 := []int{955, -9223372036854775808}
    if ((1.5 / (float64(0.1) - (0.1 * -2.25))) < float64(((float32((0.1 - (3.14159 * ((((0.1 - 3.14159) + 0.0) * 1.5) * float64((0.0 + 1e10)))))) - 3.14159) - (0.0 / (1e10 / 3.14159))))) {
        if true {
            var v10 int32 = 1
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi11
    }
    if (38 < 255) {
        v3 = v3
        v9 = append(v9, 818)
    }
    v12 := GenS1{F0: "\n\t", F1: ""}
    delete(v1, ("\n\t" + "\n\t"))
    var v13 uint16 = 0
    var v14 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
