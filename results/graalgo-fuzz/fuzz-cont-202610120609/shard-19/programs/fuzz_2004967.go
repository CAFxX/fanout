package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []uint32{0, 65535}
    var v2 uint16 = 315
    var v3 bool = (v2 <= (v2 - 352))
    v4 := make(map[uint32]uint32)
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint16 = 320
        var v7 float32 = 3.14159
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int = -591
    var v9 int32 = 179
    if len(v1) > 0 {
        v1[0] = 786
    }
    v9 = (v9 % ((42) | 1))
    v10 := GenS1{F0: 0}
    var v11 int8 = 1
    var v12 string = ("\n\t" + "")
    switch v12 {
    case "":
        if false {
            var v13 uint64 = 2
            var v14 uint8 = 45
            _ = v13
            _ = v14
        } else {
            v8 = (v8 << ((847) & 63))
            v2 -= v2
            _ = len(v1)
        }
    case "\n\t":
        v3 = (v3 || true)
    case "a":
        if ((float64(3.14159) - 1e10) != (3.14159 + (-2.25 + 1.5))) {
            if (1 == 1) {
                delete(v4, 4294967295)
                var v15 int32 = -320
                _ = v15
            } else {
                switch (v8 ^ (v8 & (^(v8)))) {
                case -781:
                    v16 := GenS1{F0: 704}
                    _ = v16
                case 0:
                    v11 += (v11 | 1)
                }
                fmt.Println(v2)
            }
        }
    }
    switch -222 {
    case 1:
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS2{F0: v9}
            _ = fi17
            _ = v18
        }
    case 9223372036854775807:
        var v19 uint64 = 4294967295
        _ = v19
    case 528:
        var v20 uint = 18446744073709551615
        _ = v20
    default:
        var v21 bool = false
        _ = v21
    }
    switch (v8 ^ v8) {
    case 42:
        v11 -= v11
    case -2147483648:
        var v22 uint8 = 77
        _ = v22
    case 2147483647:
        v9 = v9
    default:
        var v23 int8 = (v11 + (v11 - v11))
        _ = v23
    }
    var v24 uint8 = 127
    var v25 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
