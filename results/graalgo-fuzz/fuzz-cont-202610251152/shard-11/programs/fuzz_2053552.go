package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 string; F1 int64 }
func genFn0() int16 {
    var v5 uint16 = 2
    var v6 int = 2147483647
    v7 := GenS1{F0: "a", F1: 4294967295, F2: -718}
    _ = v5
    _ = v6
    _ = v7
    return -32768
}

func main() {
    var v1 uint32 = 0
    var v2 int8 = -26
    var v3 uint32 = 476
    v4 := make(map[uint32]int)
    var v8 int8 = v2
    var v9 int16 = 0
    var v10 float64 = 0.0
    _ = len(v4)
    v1 = 4294967295
    for fi11 := 0; fi11 < 12; fi11++ {
        fmt.Println(v1, v10)
        v3 = (v1 << ((4294967295) & 63))
        if (!false) {
            v10 += (v10 - 1.5)
            v4[v3] = fi11
            if ((0.1 - -2.25) >= (1e10 / ((-2.25 * v10) + v10))) {
                v12 := v4[2]
                var v13 int32 = -2147483648
                _ = v12
                _ = v13
            }
        }
        _ = fi11
    }
    v14 := GenS2{F0: "hello", F1: -1}
    fmt.Println(v3, v10)
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := []int{1}
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 uint32 = v3
            v3 = 765
            switch "\n\t" {
            case "hello":
                var v19 int32 = 42
                _ = v19
            case "a":
                var v20 uint = 83
                _ = v20
            default:
                var v21 int8 = (v8 << ((255) & 63))
                _ = v21
            }
            _ = fi17
            _ = v18
        }
        var v22 int = fi15
        _ = fi15
        _ = v16
        _ = v22
    }
    v23 := GenS1{F0: "hello", F1: (v1 * (^(v3))), F2: genFn0()}
    if ((!(255 == 103)) || false) {
        v8 -= (v2 | v2)
    } else {
        _ = len(v4)
        if true {
            v24 := v4[v3]
            fmt.Println(v24, v1)
            v9 = v9
            _ = v24
        }
    }
    var v25 int = 999
    v2 = v2
    switch v25 {
    case 32768:
        fmt.Println(v2, v25, v8)
    default:
        var v26 float32 = -2.25
        _ = v26
    }
    v27 := []int{1, (v25 >> ((1) & 63)), v25}
    for fi28 := 0; fi28 < 4; fi28++ {
        if (-393 >= 1) {
            v3 = v1
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            fmt.Println(v9)
        }
        var v29 int = (v25 ^ fi28)
        var v30 uint64 = 1
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v23
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
