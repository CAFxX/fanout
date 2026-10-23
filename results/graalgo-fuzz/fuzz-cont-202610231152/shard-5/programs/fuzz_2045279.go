package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 226
    var v3 uint16 = 65535
    var v4 int8 = 127
    v5 := []uint32{uint32(v2)}
    v6 := make(map[int]string)
    if true {
        v2 = uint64(v3)
        v4 += (v4 >> ((4294967295) & 63))
        var v7 uint = 18446744073709551615
        _ = v7
    }
    v4 = (v4 & v4)
    v3 = v3
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            v5 = append(v5, uint32(v3))
            v2 += v2
            v10 := GenS1{F0: -1, F1: -559}
            fmt.Println(fi9)
            _ = fi9
            _ = v10
        }
        var v11 uint32 = 916
        _ = fi8
        _ = v11
    }
    fmt.Println(v2, v2, v2)
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(fi12, v2, v4)
        var v13 uint8 = 1
        fmt.Println(fi12, fi12)
        _ = fi12
        _ = v13
    }
    v3 = (v3 << ((18446744073709551615) & 63))
    switch "\n\t" {
    case "hello":
        v4 += 127
    case "\n\t":
        for fi14 := 0; fi14 < 4; fi14++ {
            v3 = uint16(v2)
            _ = fi14
        }
    case "":
        v15 := GenS1{F0: -792, F1: 0}
        _ = v15
    }
    v16 := GenS1{F0: -359, F1: -32768}
    for fi17 := 0; fi17 < 17; fi17++ {
        v5 = append(v5, uint32(v2))
        _ = fi17
    }
    v3 += 42
    for fi18 := 0; fi18 < 12; fi18++ {
        v3 += 1
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int64 = 0
            v6[388] = ""
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
