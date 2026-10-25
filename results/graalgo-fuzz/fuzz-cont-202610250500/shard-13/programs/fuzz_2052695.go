package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]uint32)
    var v3 int8 = -128
    var v4 string = ("hello" + "a")
    v1 = (v1 * v1)
    var v5 int32 = 0
    var v6 int8 = -128
    var v7 int = 0
    v2[808] = 2147483647
    if (-9223372036854775808 == 42) {
        if false {
            if (false || false) {
                if (65535 < 1) {
                    v1 -= 1e10
                }
            }
            v8 := []int{(-(v7))}
            var v9 uint8 = 2
            _ = v8
            _ = v9
        }
    } else {
        v10 := GenS1{F0: (!(255 == 644)), F1: -21, F2: v7}
        _ = v10
    }
    var v11 uint32 = 987
    v7 = v7
    var v12 float64 = float64(v1)
    var v13 float32 = 0.1
    switch ("\n\t" + ("\n\t" + v4)) {
    case "a":
        var v14 int32 = v5
        _ = v14
    }
    var v15 uint8 = 71
    var v16 int64 = -9223372036854775808
    v17 := GenS1{F0: true, F1: v16, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
