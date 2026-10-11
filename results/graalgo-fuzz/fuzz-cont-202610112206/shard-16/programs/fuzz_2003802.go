package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint16 = 1
    var v2 string = "foo bar"
    v3 := make(map[uint32]string)
    v4 := []string{(("\n\t" + "\n\t") + ("\n\t" + v2)), ("" + ("\n\t" + ("" + ""))), (v2 + "\n\t"), (v2 + v2)}
    var v5 int16 = -1
    v2 = v2
    var v6 uint32 = 530
    if true {
        var v7 int8 = 2
        v8 := GenS1{F0: -9223372036854775808, F1: 592, F2: 774}
        v6 = (v6 * v6)
        _ = v7
        _ = v8
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v9 := GenS1{F0: int64(v5), F1: -9223372036854775808, F2: -316}
        _ = v9
    }
    v5 -= (v5 | 0)
    var v10 int64 = 9223372036854775807
    v11 := GenS2{F0: (^(v1)), F1: 100, F2: 2147483647}
    fmt.Println(v2, v1)
    var v12 int32 = 42
    fmt.Println(v10)
    var v13 int16 = 381
    var v14 uint64 = uint64(v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v15 uint16 = v1
    v16 := GenS3{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
