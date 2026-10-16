package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int64 }
func genFn0() int8 {
    var v7 uint16 = 65535
    v8 := GenS1{F0: 42, F1: 42, F2: -9223372036854775808}
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint) string {
    fmt.Println(p0)
    var v9 uint8 = 22
    if false {
        var v10 int = -526
        if (!true) {
            v9 = 255
            var v11 bool = false
            _ = v11
        }
        _ = v10
    } else {
        v12 := GenS1{F0: (v9 % ((128) | 1)), F1: 100, F2: 42}
        _ = v12
    }
    _ = p0
    _ = v9
    return "foo bar"
}

func main() {
    v1 := []int{-9223372036854775808, -2147483648, 128}
    var v2 uint64 = 42
    v3 := []uint32{65535, 858, 2147483647, 805}
    v4 := []string{(("a" + "foo bar") + (("a" + ("\n\t" + "")) + "\n\t"))}
    v5 := make(map[uint32]int)
    v6 := []uint32{2, uint32(v2), 0}
    fmt.Println(v2)
    v2 = 0
    var v13 uint32 = 2
    var v14 uint8 = uint8(v2)
    var v15 int = 257
    v15 = (-(v15))
    var v16 int32 = -412
    v16 = 0
    var v17 uint8 = v14
    v17 -= (v14 - (^(v14)))
    for fi18 := 0; fi18 < 9; fi18++ {
        fmt.Println(fi18, v14)
        _ = fi18
    }
    if len(v6) > 0 {
        v6[0] = v13
    }
    delete(v5, (v13 - uint32(v2)))
    fmt.Println(v14, v14, v14)
    v19 := GenS1{F0: v14, F1: 42, F2: 117}
    var v20 int8 = genFn0()
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
