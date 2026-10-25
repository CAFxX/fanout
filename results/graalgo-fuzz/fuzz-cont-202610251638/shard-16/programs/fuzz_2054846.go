package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }
func genFn0(p0 uint8, p1 uint8) string {
    var v6 uint16 = 42
    v6 = (v6 >> ((65535) & 63))
    _ = p0
    _ = p1
    _ = v6
    return "a"
}

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{}
    var v3 uint = 255
    var v4 uint8 = 42
    var v5 uint32 = 2147483648
    switch 9223372036854775807 {
    case 42:
        v1 = -2147483648
    }
    var v7 bool = (0 > (v4 % (((v4 % (((v4 + 44)) | 1))) | 1)))
    var v8 int = 128
    v8 = (v8 * (v8 + int(v4)))
    var v9 int8 = 1
    v7 = true
    v10 := GenS1{F0: 456, F1: (v1 >> (((^(v3))) & 63)), F2: 9223372036854775807}
    var v11 int16 = -32768
    var v12 uint16 = 255
    if (0.1 < float64(-2.25)) {
        if len(v2) > 0 {
            v2[0] = 821
        }
    } else {
        var v13 int64 = 1
        var v14 uint16 = v12
        _ = v13
        _ = v14
    }
    if (127 < v9) {
        switch (v8 & (v8 << ((v3) & 63))) {
        case 0:
            var v15 uint16 = (^(v12))
            _ = v15
        case 2147483647:
            v7 = (uint32(v4) <= (^(v5)))
        }
    }
    var v16 uint64 = 42
    v1 += -2147483648
    if (v11 >= v11) {
        var v17 float32 = 3.14159
        var v18 uint8 = uint8(v12)
        if len(v2) > 0 {
            v2[0] = (v5 | 1)
        }
        _ = v17
        _ = v18
    }
    v7 = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
