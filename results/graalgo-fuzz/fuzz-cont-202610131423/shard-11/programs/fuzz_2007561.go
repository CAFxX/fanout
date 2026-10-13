package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -1
    v3 := make(map[int]uint32)
    v4 := []uint32{2}
    var v5 int16 = -32768
    var v6 int16 = (^(v5))
    var v7 int64 = v2
    fmt.Println(v2, v5, v2)
    v5 += v5
    v5 += 589
    var v8 int64 = v7
    var v9 uint = 4294967295
    fmt.Println(v6, v7)
    fmt.Println(v8)
    var v10 int32 = -846
    var v11 uint32 = 2
    var v12 bool = true
    if v12 {
        var v13 int = 0
        _ = v13
    }
    switch 4294967295 {
    case -350:
        if v12 {
            v14 := GenS1{F0: -58}
            _ = v14
        } else {
            v12 = v12
        }
    case 4294967295:
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint8 = 128
            _ = fi15
            _ = v16
        }
    case 128:
        v7 -= v2
    default:
        if v12 {
            for fi17 := 0; fi17 < 2; fi17++ {
                v6 = v6
                var v18 bool = (!(1 < 0))
                fmt.Println(v2, v11, v10)
                _ = fi17
                _ = v18
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v19 := GenS1{F0: 127}
            _ = v19
        } else {
            v4 = append(v4, v11)
            var v20 int16 = 42
            v2 = v7
            _ = v20
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
