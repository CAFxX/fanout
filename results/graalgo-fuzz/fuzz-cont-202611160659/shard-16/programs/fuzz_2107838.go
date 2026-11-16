package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 bool }
func genFn0(p0 int16) int64 {
    p0 += p0
    p0 += -797
    _ = p0
    return 0
}

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 4294967295
    v3 := []string{}
    var v4 uint = 255
    v5 := make(map[uint32]uint32)
    v6 := []string{}
    v7 := GenS1{F0: 1}
    v8 := GenS2{F0: (v2 - 42)}
    fmt.Println(v2, v2, v2)
    var v9 bool = ((-1 != genFn0(32767)) || false)
    if v9 {
        v9 = (4294967295 <= 42)
        fmt.Println(v1, v9)
    } else {
        v4 -= 4294967295
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        _ = len(v6)
        delete(v5, 42)
        _ = fi10
    }
    switch "foo bar" {
    case "a":
        delete(v5, 42)
    default:
        var v11 uint16 = 2
        _ = v11
    }
    var v12 uint = v4
    v13 := GenS2{F0: (v2 + 766)}
    v1 = v1
    switch v1 {
    case "a":
        for fi14 := 0; fi14 < 18; fi14++ {
            for fi15 := 0; fi15 < 11; fi15++ {
                v9 = v9
                _ = fi15
            }
            v1 = (v1 + v1)
            v16 := GenS3{F0: uint16(v2), F1: 2147483647, F2: (!(((-2.25 - float32(0.0)) / 3.14159) < 0.0))}
            _ = fi14
            _ = v16
        }
    case "":
        v9 = v9
    case "\n\t":
        for fi17 := 0; fi17 < 19; fi17++ {
            v1 = v1
            _ = fi17
        }
    }
    v5[349] = 0
    var v18 uint = (v12 * v4)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v1 = v1
        var v20 int = 1
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
