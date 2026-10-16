package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{834, 2147483647, 269}
    v3 := []uint32{2, 42, 4294967295}
    var v4 uint = 539
    var v5 uint8 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v6 uint32 = 1
    v6 += 1
    v2 = append(v2, (v6 * (v6 >> ((v4) & 63))))
    switch 202 {
    case 4294967295:
        v7 := GenS1{F0: 0}
        _ = v7
    case 2147483647:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 32768:
        fmt.Println(v6)
    }
    fmt.Println(v5, v4, v5)
    var v8 uint16 = 844
    for fi9 := 0; fi9 < 14; fi9++ {
        v3 = append(v3, 0)
        var v10 uint32 = (v6 & (^(v6)))
        _ = fi9
        _ = v10
    }
    fmt.Println(v6, v5)
    fmt.Println(v6, v4)
    var v11 uint32 = 65535
    switch "foo bar" {
    case "a":
        fmt.Println(v1)
    }
    fmt.Println(v1, v5, v1)
    var v12 float64 = 3.14159
    if len(v2) > 0 {
        v2[0] = (^(v6))
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        switch "" {
        case "":
            v14 := GenS1{F0: (v8 * (v8 ^ v8))}
            _ = v14
        default:
            if len(v3) > 0 {
                v3[0] = (v11 >> ((42) & 63))
            }
        }
        v11 = v6
        _ = fi13
    }
    _ = len(v2)
    if true {
        fmt.Println(v11, v6)
        v15 := GenS1{F0: v8}
        var v16 uint8 = (v5 / (((v5 & 127)) | 1))
        _ = v15
        _ = v16
    } else {
        if false {
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 int = (fi17 * (fi17 + fi17))
                v19 := GenS2{F0: 42, F1: v5}
                var v20 int8 = 100
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            fmt.Println(v6)
            if (true && true) {
                v4 = (v4 >> (((^(v4))) & 63))
            } else {
                fmt.Println(v8, v5, v1)
            }
        } else {
            var v21 float32 = 1.5
            _ = v21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
