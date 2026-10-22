package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = 467
    var v2 uint8 = 2
    var v3 string = ""
    v4 := GenS1{F0: 42}
    fmt.Println(v1, v3)
    v5 := []int{-419, 127, 127}
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    var v6 float32 = (((0.1 - 0.1) - ((0.1 / 1e10) + (1.5 * -2.25))) / (float32(1e10) - ((1.5 + 0.1) - 0.0)))
    if (-460 > 1) {
        if true {
            var v7 uint8 = v2
            var v8 bool = (false && false)
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 2; fi9++ {
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 int8 = 100
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
            var v12 int = -2147483648
            _ = v12
        }
        if true {
            _ = len(v5)
            v3 = ("a" + v3)
        }
        v3 = v3
    }
    fmt.Println(v6)
    v1 = v1
    var v13 float32 = v6
    var v14 uint32 = 2147483647
    var v15 int64 = v1
    switch 2147483647 {
    case 4294967295:
        _ = len(v5)
    }
    var v16 uint64 = 6
    if (true || true) {
        var v17 int32 = 42
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
