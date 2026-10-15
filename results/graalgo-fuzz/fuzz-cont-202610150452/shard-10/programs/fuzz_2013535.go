package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    var v4 int32 = 0
    var v5 int = 1
    v4 -= (v4 % ((v4) | 1))
    var v6 int = v5
    var v7 int32 = v4
    var v8 uint8 = 190
    var v9 uint64 = 18446744073709551615
    v7 = v4
    v10 := GenS1{F0: 32767}
    v5 = -450
    v5 -= v5
    var v11 uint64 = (^(v9))
    if (0 >= 228) {
        switch v5 {
        case 32767:
            v6 = v6
        case -562:
            if v1 {
                var v12 bool = v1
                _ = v12
            }
        case 307:
            var v13 int32 = (^(v7))
            _ = v13
        }
        var v14 uint8 = (v8 ^ (^(v8)))
        v2["a"] = (v6 << ((0) & 63))
        _ = v14
    }
    v2["hello"] = (v6 >> ((879) & 63))
    var v15 uint16 = 255
    var v16 float32 = (0.0 + 1e10)
    v17 := v2["foo bar"]
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := GenS1{F0: 0}
            fmt.Println(fi19)
            var v21 int8 = 29
            _ = fi19
            _ = v20
            _ = v21
        }
        if v1 {
            v22 := GenS1{F0: 32767}
            v9 -= (v9 + (v9 * (v9 / ((1) | 1))))
            _ = len(v3)
            _ = v22
        } else {
            delete(v2, "")
            v23 := v2["a"]
            _ = v23
        }
        v5 -= (v6 / (((^(v5))) | 1))
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
