package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint16 = 42
    var v3 float64 = -2.25
    var v4 int8 = 42
    v5 := []uint32{(v1 * v1), 42, 357}
    v1 += v1
    fmt.Println(v3, v3)
    var v6 uint8 = 255
    v4 = -128
    v5 = append(v5, 2147483647)
    if ((127 > 9223372036854775807) || ((v6 == 0) || ((((712 > 95) && (!(-355 < 738))) || false) && (42 == 0)))) {
        v7 := GenS1{F0: uint32(v2)}
        _ = v7
    }
    if false {
        var v8 uint8 = (v6 << ((255) & 63))
        if false {
            v6 = v6
            switch int(v2) {
            case 32767:
                var v9 uint16 = (^(v2))
                _ = v9
            case 9223372036854775807:
                var v10 int8 = (v4 & 117)
                _ = v10
            }
        }
        _ = len(v5)
        _ = v8
    }
    fmt.Println(v2, v4, v1)
    var v11 int64 = 0
    var v12 int = -695
    if (3.14159 >= 0.0) {
        for fi13 := 0; fi13 < 2; fi13++ {
            if len(v5) > 0 {
                v5[0] = (v1 + 2147483648)
            }
            _ = fi13
        }
        var v14 string = ("hello" + "\n\t")
        _ = v14
    }
    var v15 uint32 = 2147483648
    v16 := GenS1{F0: (v1 % (((v15 / (((v15 & 2147483647)) | 1))) | 1))}
    v12 = v12
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
