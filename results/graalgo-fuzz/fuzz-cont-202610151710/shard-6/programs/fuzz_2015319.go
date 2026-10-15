package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
func genFn0() bool {
    if false {
        switch "hello" {
        case "foo bar":
            var v7 string = "a"
            _ = v7
        default:
            var v8 float64 = 0.1
            _ = v8
        }
        v9 := GenS1{F0: 597, F1: -1, F2: 547}
        var v10 uint32 = 1
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        fmt.Println(fi11)
        _ = fi11
    }
    var v12 uint = 420
    v12 = v12
    _ = v12
    return (4294967295 == 42)
}

func main() {
    v1 := []string{("" + "\n\t"), "foo bar", "foo bar", "foo bar"}
    var v2 uint32 = 345
    var v3 uint8 = 0
    var v4 uint = 676
    var v5 bool = false
    var v6 float64 = 0.1
    v5 = (0 < 127)
    v13 := GenS1{F0: 42, F1: -846, F2: 0}
    v1 = append(v1, "hello")
    var v14 uint32 = (v2 | 1)
    var v15 uint32 = (v14 >> (((v4 & 1)) & 63))
    v16 := GenS1{F0: 1, F1: -663, F2: 539}
    _ = len(v1)
    if (!(-1 < 0)) {
        fmt.Println(v14, v14, v5)
        v6 = 1.5
    } else {
        v15 += (^(v2))
        _ = len(v1)
    }
    var v17 uint8 = v3
    var v18 uint8 = v3
    if ((true && v5) && true) {
        var v19 float32 = ((0.1 * 1e10) + (3.14159 * 1.5))
        var v20 float32 = v19
        v17 -= v18
        _ = v19
        _ = v20
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
