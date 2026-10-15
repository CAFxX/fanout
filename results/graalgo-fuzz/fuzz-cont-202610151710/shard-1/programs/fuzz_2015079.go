package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 string }

func main() {
    var v1 uint16 = 42
    v2 := []int{2147483647, 32767, 0, 9223372036854775807}
    var v3 uint64 = 2
    v4 := []uint32{42}
    v5 := make(map[uint32]bool)
    var v6 int = 127
    fmt.Println(v3, v1)
    var v7 int16 = 0
    v1 -= (v1 << ((65535) & 63))
    v3 = 233
    if (true || true) {
        v7 = v7
        if false {
            delete(v5, 0)
        } else {
            v1 = v1
        }
        v6 = (v6 & 0)
    }
    var v8 uint64 = v3
    fmt.Println(v8, v7, v6)
    if true {
        for fi9 := 0; fi9 < 9; fi9++ {
            v3 = v8
            v3 -= (v3 * v3)
            _ = fi9
        }
        var v10 int32 = 929
        v6 = v6
        _ = v10
    }
    v11 := GenS1{F0: 42, F1: 1, F2: ("foo bar" + "foo bar")}
    if true {
        var v12 int16 = (v7 | (v7 - 42))
        if (-615 > -2147483648) {
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 bool = true
                v3 -= uint64(v1)
                v7 -= (v12 / (((v7 + (v12 << ((18446744073709551615) & 63)))) | 1))
                _ = fi13
                _ = v14
            }
        } else {
            v15 := v5[42]
            _ = v15
        }
        var v16 int16 = (v7 + v12)
        _ = v12
        _ = v16
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v1)
    }
    fmt.Println(v8, v6, v8)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
