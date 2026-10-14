package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 int32 }

func main() {
    v1 := []int{2147483647}
    var v2 int64 = 9223372036854775807
    var v3 float32 = (1e10 + 0.1)
    var v4 uint16 = 65535
    v5 := make(map[uint32]uint32)
    var v6 int32 = 2147483647
    var v7 string = ""
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int8 = 1
        v10 := GenS1{F0: 0}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 string = ""
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v12 int64 = v2
    if (-128 <= 1) {
        for fi13 := 0; fi13 < 2; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 int64 = (v12 / (((v12 | (v2 - v2))) | 1))
                switch (fi13 << ((903) & 63)) {
                case -1:
                    v5[2147483647] = 1
                case 1:
                    _ = len(v1)
                case -2147483648:
                    var v16 int32 = (v6 | -58)
                    _ = v16
                default:
                    v2 += (v12 >> ((65535) & 63))
                }
                _ = len(v1)
                if true {
                    _ = len(v5)
                } else {
                    var v17 int8 = int8(fi13)
                    switch (fi14 * fi14) {
                    case -1:
                        v1 = append(v1, fi13)
                    }
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
            fmt.Println(v7)
            _ = fi13
        }
        v18 := GenS2{F0: 42, F1: 0, F2: 42}
        _ = v18
    } else {
        v19 := GenS1{F0: int16(v6)}
        var v20 int16 = 0
        _ = v19
        _ = v20
    }
    var v21 uint8 = 1
    var v22 string = ((v7 + v11) + (v11 + (v7 + ((v7 + "a") + v7))))
    v23 := GenS1{F0: -1}
    var v24 int16 = 0
    switch 32767 {
    case -426:
        fmt.Println(v22)
    }
    delete(v5, 4294967295)
    v25 := GenS1{F0: v24}
    v24 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
