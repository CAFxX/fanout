package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
func genFn0(p0 uint32, p1 int32, p2 uint32) uint16 {
    switch -249 {
    case 982:
        p2 = (^(p0))
    case 32767:
        p1 = p1
    default:
        fmt.Println(p2)
    }
    var v5 bool = false
    var v6 uint16 = 0
    v7 := GenS1{F0: ((("hello" + "\n\t") + "hello") + "foo bar"), F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 2
    v3 := []uint32{4294967295, 4294967295, 212, 1}
    var v4 float64 = 0.1
    fmt.Println(v1, v1)
    fmt.Println(v1)
    fmt.Println(v2, v2)
    v4 = float64(1e10)
    switch "" {
    case "hello":
        v1 = 1.5
    case "":
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
    case "a":
        var v10 string = "a"
        _ = v10
    }
    v2 += v2
    var v11 uint64 = 18446744073709551615
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            v3 = append(v3, 42)
            var v14 int = (fi12 << ((460) & 63))
            _ = fi13
            _ = v14
        }
        v2 = 206
        _ = fi12
    }
    if false {
        var v15 uint = 65535
        switch "\n\t" {
        case "":
            fmt.Println(v4, v4)
        case "\n\t":
            for fi16 := 0; fi16 < 17; fi16++ {
                v3 = append(v3, 2147483647)
                _ = fi16
            }
        case "a":
            v2 = (v2 >> ((255) & 63))
        }
        var v17 int16 = -1
        _ = v15
        _ = v17
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v4, v4)
    v4 = (float64((3.14159 * float32(v4))) - v4)
    var v18 uint = 224
    fmt.Println(v4, v2, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
