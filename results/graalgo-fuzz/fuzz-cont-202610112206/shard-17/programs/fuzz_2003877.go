package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int32; F1 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 int = 2147483647
    var v3 int8 = 127
    var v4 uint16 = 531
    v5 := make(map[uint32]int)
    var v6 int64 = -606
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 bool = ((!(!(false || false))) || true)
        v8 = (v4 <= (v4 % (((v4 | v4)) | 1)))
        _ = fi7
        _ = v8
    }
    fmt.Println(v6, v4, v4)
    var v9 float32 = (-2.25 * -2.25)
    v10 := GenS1{F0: 1, F1: 1}
    if (600 == 1) {
        if (((!(((1 > 1) || (true && false)) || (false || (42 == 0)))) && (0.0 > v9)) && ("a" >= "foo bar")) {
            v2 -= v2
        } else {
            var v11 float32 = -2.25
            v12 := GenS1{F0: v3, F1: 42}
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 uint = 89
                _ = fi13
                _ = v14
            }
            _ = v11
            _ = v12
        }
        fmt.Println(v6)
    }
    v15 := GenS1{F0: v3, F1: 733}
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 uint64 = 18446744073709551615
        var v18 int16 = -32768
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 int8 = v3
    v20 := GenS2{F0: 767, F1: 42}
    v21 := []int{(v2 - -2147483648), (v2 % (((v2 | 4294967295)) | 1)), v2}
    v3 = (v19 << ((4294967295) & 63))
    delete(v1, v2)
    fmt.Println(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
