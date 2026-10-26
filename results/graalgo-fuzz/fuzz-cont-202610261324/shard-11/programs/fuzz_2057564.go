package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
func genFn0() int32 {
    var v7 uint16 = 65535
    var v8 int8 = -49
    var v9 uint8 = 0
    v10 := []int{-2147483648}
    v7 += (v7 * v7)
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint32 = 2147483648
    var v2 int32 = -53
    v3 := make(map[uint32]int)
    var v4 uint32 = v1
    var v5 uint16 = 42
    v6 := []uint32{}
    v4 = v4
    fmt.Println(v4, v5, v1)
    if (!(uint32(v5) <= v1)) {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int8 = -124
            v12 = v12
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: int64(v5), F1: 42}
        _ = v13
    }
    switch (((("" + "a") + ("a" + "\n\t")) + (("\n\t" + "foo bar") + "hello")) + "") {
    case "hello":
        if (65535 <= v1) {
            v5 += v5
        }
    }
    var v14 float32 = 0.0
    var v15 uint8 = 255
    var v16 float32 = float32((1e10 + 0.0))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    delete(v3, (v1 + (v1 >> ((0) & 63))))
    var v17 bool = false
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 float64 = 1.5
            v15 += uint8(v4)
            v21 := GenS2{F0: 541, F1: (^(v5)), F2: -1}
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 uint64 = uint64(v1)
        _ = fi18
        _ = v22
    }
    v15 += (v15 ^ v15)
    for fi23 := 0; fi23 < 13; fi23++ {
        fmt.Println(v15)
        var v24 float32 = v16
        _ = len(v6)
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
