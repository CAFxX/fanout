package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int32) uint {
    p0 -= (p0 & (p0 << ((4294967295) & 63)))
    v6 := GenS1{F0: "a", F1: 119}
    _ = p0
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    v5 := []string{"a", (("\n\t" + "a") + "hello")}
    switch ("" + "\n\t") {
    case "foo bar":
        var v7 float32 = 0.0
        _ = v7
    default:
        switch "hello" {
        case "hello":
            fmt.Println(v1)
        }
    }
    v1 += v1
    var v8 int32 = (v1 + (v1 ^ (^(v1))))
    switch ("\n\t" + "hello") {
    case "":
        v5 = append(v5, "hello")
    }
    v1 -= v8
    fmt.Println(v1, v1)
    v9 := GenS2{F0: 9223372036854775807}
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := v4[636]
        v2[fi10] = (352 != 305)
        _ = fi10
        _ = v11
    }
    v8 = (v8 * v8)
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS3{F0: "foo bar", F1: 120, F2: 32767}
        _ = fi12
        _ = v13
    }
    v14 := GenS4{F0: 4}
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint16 = 255
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
