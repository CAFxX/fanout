package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 9223372036854775807
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    v5 := []int{1}
    v2 = v2
    if (13 >= 65535) {
        v2 = (v2 >> ((65535) & 63))
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(v2, v2)
            var v7 uint64 = 318
            _ = fi6
            _ = v7
        }
        if true {
            var v8 uint64 = 18446744073709551615
            _ = v8
        }
    } else {
        var v9 int8 = -128
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: 793}
            v12 := GenS1{F0: 255}
            _ = fi10
            _ = v11
            _ = v12
        }
        delete(v3, "")
        _ = v9
    }
    var v13 uint64 = 2
    v14 := GenS2{F0: 42}
    _ = len(v1)
    v15 := GenS1{F0: 719}
    if false {
        var v16 int = 2147483647
        var v17 int16 = 635
        fmt.Println(v2, v16, v2)
        _ = v16
        _ = v17
    } else {
        v18 := GenS1{F0: 2}
        _ = v18
    }
    v2 -= 4294967295
    v19 := GenS2{F0: 42}
    if true {
        var v20 uint = 2
        var v21 uint16 = 65535
        v21 -= uint16(v13)
        _ = v20
        _ = v21
    } else {
        v4[1] = "a"
        var v22 int32 = 42
        v22 = v22
        _ = v22
    }
    if true {
        for fi23 := 0; fi23 < 8; fi23++ {
            var v24 string = "\n\t"
            v25 := v4[553]
            _ = fi23
            _ = v24
            _ = v25
        }
        for fi26 := 0; fi26 < 1; fi26++ {
            fmt.Println(v13)
            for fi27 := 0; fi27 < 4; fi27++ {
                v2 = 4294967295
                var v28 float32 = ((0.0 * float32(1.5)) * 1.5)
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 15; fi29++ {
                var v30 uint16 = 388
                v30 += v30
                v13 = (^(v13))
                _ = fi29
                _ = v30
            }
            _ = fi26
        }
    }
    v31 := GenS1{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
