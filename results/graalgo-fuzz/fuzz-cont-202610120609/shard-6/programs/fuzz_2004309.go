package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 uint32) int32 {
    switch -9223372036854775808 {
    case 42:
        v7 := GenS1{F0: 128, F1: 18446744073709551615, F2: -884}
        _ = v7
    }
    v8 := GenS1{F0: 212, F1: 1, F2: 0}
    p0 -= 127
    _ = p0
    _ = p1
    _ = v8
    return 2147483647
}

func main() {
    var v1 int32 = -2147483648
    var v2 int = 128
    v3 := make(map[uint32]string)
    var v4 uint16 = 0
    v5 := []int{(v2 + v2)}
    var v6 int32 = v1
    v6 += v6
    v9 := GenS2{F0: "foo bar"}
    v6 = (v1 % ((-1) | 1))
    if len(v5) > 0 {
        v5[0] = v2
    }
    var v10 uint32 = 42
    var v11 int32 = (v1 % (((v1 | (v6 + v6))) | 1))
    v11 += (v11 / ((v6) | 1))
    if true {
        var v12 int16 = int16(v6)
        var v13 int16 = 1000
        fmt.Println(v11, v12, v4)
        _ = v12
        _ = v13
    } else {
        switch v2 {
        case 2147483647:
            delete(v3, v10)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
