package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int16, p1 int) string {
    var v5 uint = 494
    var v6 float64 = 3.14159
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 bool = ((1.5 + 1e10) > float32((float64(float32(0.0)) * v6)))
        _ = fi7
        _ = v8
    }
    if false {
        v9 := GenS1{F0: 340}
        _ = v9
    } else {
        var v10 int = (p1 << ((v5) & 63))
        v5 = v5
        switch -2147483648 {
        case 127:
            v11 := []int{(p1 >> ((v5) & 63)), v10, v10}
            _ = v11
        case 32768:
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 int32 = -854
                v14 := GenS1{F0: 1}
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = v10
    }
    p1 -= 32768
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 int8 = -128
    delete(v1, "a")
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int64 = 9223372036854775807
        v4 -= (v4 & v4)
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 0}
    v4 += -29
    var v18 uint8 = 127
    v19 := GenS1{F0: 300}
    var v20 uint32 = 336
    if false {
        v21 := v1["\n\t"]
        _ = v21
    } else {
        v22 := []int{455, -281}
        var v23 uint16 = 702
        v24 := GenS1{F0: 182}
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 float32 = 1.5
        _ = fi25
        _ = v26
    }
    if true {
        v4 = (v4 ^ (v4 - v4))
        v4 = (v4 << ((2) & 63))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
