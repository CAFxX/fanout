package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint16 }
type GenS3 struct { F0 int32; F1 int; F2 bool }
func genFn0() bool {
    v7 := GenS1{F0: 255, F1: 1}
    v8 := []int{}
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := []int{}
    var v2 float32 = (1e10 * 0.0)
    var v3 int = 9223372036854775807
    var v4 int32 = 0
    var v5 uint32 = 49
    v6 := []uint32{v5, v5}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 := GenS2{F0: 453, F1: v3, F2: 2}
    switch (v3 ^ v3) {
    case 127:
        v10 := GenS3{F0: v4, F1: (v3 & v3), F2: ((!((true && (3.14159 == 0.0)) || (!(4294967295 != 42)))) || false)}
        _ = v10
    case -594:
        var v11 int64 = 0
        _ = v11
    case 128:
        var v12 bool = false
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        fmt.Println(v2)
        if (!(311 <= 171)) {
            var v14 uint16 = 1
            _ = len(v6)
            var v15 uint32 = v5
            _ = v14
            _ = v15
        } else {
            if len(v6) > 0 {
                v6[0] = (v5 % ((v5) | 1))
            }
            v2 += float32((0.1 + 0.1))
        }
        _ = fi13
    }
    v4 += v4
    switch ("hello" + "hello") {
    case "":
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 uint8 = 42
            _ = fi16
            _ = v17
        }
    case "hello":
        var v18 uint64 = 763
        _ = v18
    case "a":
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v4, v5, fi19)
            _ = fi19
        }
    default:
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
    }
    switch "foo bar" {
    case "a":
        var v20 uint = 18446744073709551615
        _ = v20
    default:
        v3 = 32767
    }
    v1 = append(v1, v3)
    if len(v6) > 0 {
        v6[0] = (v5 ^ v5)
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 int16 = 42
        var v23 uint16 = 2
        v4 += v4
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v5, v3)
    fmt.Println(v2, v5, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
