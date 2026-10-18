package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0() uint32 {
    var v6 int16 = 42
    v6 -= v6
    var v7 int32 = 817
    if (true || true) {
        var v8 uint32 = 1
        v6 = v6
        v6 += 934
        _ = v8
    } else {
        var v9 int16 = v6
        var v10 uint8 = 1
        v11 := GenS1{F0: 200, F1: 4294967295}
        _ = v10
        _ = v11
        _ = v9
    }
    _ = v6
    _ = v7
    return 559
}

func main() {
    var v1 uint8 = 0
    var v2 bool = true
    var v3 uint64 = 42
    var v4 uint32 = 0
    var v5 int8 = 0
    v2 = (127 < 32768)
    v3 = v3
    var v12 int = -1
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := []int{v12}
        fmt.Println(v2, v2, v3)
        _ = fi13
        _ = v14
    }
    v15 := []int{int(v3)}
    var v16 float64 = (float64((0.0 / (1e10 + 3.14159))) / 3.14159)
    fmt.Println(v12, v3)
    var v17 float64 = float64((1.5 * 1e10))
    if v2 {
        v18 := GenS1{F0: (v12 >> ((uint(v4)) & 63)), F1: 2}
        switch v12 {
        case -1:
            var v19 uint = 988
            _ = v19
        }
        _ = v18
    } else {
        if v2 {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 int16 = 0
                var v22 string = "foo bar"
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := GenS1{F0: 2147483647, F1: 255}
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        v3 -= (v3 * (v3 * 293))
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
