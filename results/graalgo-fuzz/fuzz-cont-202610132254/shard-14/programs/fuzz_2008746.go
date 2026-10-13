package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int16 }

func main() {
    var v1 bool = (-32768 < 1000)
    v2 := make(map[int]int)
    v3 := []string{}
    v4 := []uint32{0, 4294967295, 0, 0}
    v5 := []string{"\n\t"}
    var v6 int8 = 100
    var v7 uint = 2
    v1 = v1
    var v8 int16 = 1
    var v9 float64 = 1.5
    if (true && (v1 || v1)) {
        var v10 int32 = -262
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int = 2147483647
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            if (true || false) {
                v14 := GenS1{F0: 0}
                _ = v14
            }
            _ = fi13
        }
    }
    var v15 int = 861
    v16 := GenS2{F0: 0, F1: v8, F2: (-(v8))}
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int8 = 1
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 int = 128
        var v21 int = 127
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 float32 = 0.0
    v2[32768] = (v15 % ((v15) | 1))
    fmt.Println(v6)
    var v23 uint32 = 2147483648
    var v24 uint64 = 18446744073709551615
    var v25 int8 = (v6 << ((v7) & 63))
    var v26 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
