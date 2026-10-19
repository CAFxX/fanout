package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = 9223372036854775807
    var v3 bool = (true && (390 > (v2 - -9223372036854775808)))
    var v4 uint = 65535
    v5 := make(map[uint32]bool)
    var v6 string = (("\n\t" + "") + "hello")
    fmt.Println(v1, v1, v3)
    var v7 uint64 = 4294967295
    var v8 uint32 = 1
    var v9 bool = (208 != uint8(v8))
    if v3 {
        if (-459 >= -981) {
            v6 = v6
            fmt.Println(v9, v7)
            var v10 uint16 = 2
            _ = v10
        }
        _ = len(v5)
    }
    v11 := GenS1{F0: (v4 % (((v4 | (v4 / (((v4 & v4)) | 1)))) | 1)), F1: 1}
    fmt.Println(v1, v4, v3)
    var v12 int16 = 32767
    v7 -= (v7 << ((1) & 63))
    v12 = (v12 / (((-(v12))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
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
