package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint64 }

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = 42
    var v3 uint32 = 0
    v4 := []uint32{1, v3, (v3 | v3)}
    if len(v4) > 0 {
        v4[0] = v3
    }
    if len(v4) > 0 {
        v4[0] = v3
    }
    switch ("foo bar" + "hello") {
    case "hello":
        v5 := GenS1{F0: 127}
        _ = v5
    case "\n\t":
        v6 := GenS1{F0: 127}
        _ = v6
    }
    _ = len(v4)
    var v7 int8 = -128
    fmt.Println(v1)
    var v8 int64 = -1
    switch -99 {
    case 875:
        fmt.Println(v1, v8, v7)
    default:
        v9 := GenS2{F0: 42}
        _ = v9
    }
    if ((("\n\t" + ("foo bar" + (("" + "hello") + ("hello" + "hello")))) + "") >= "foo bar") {
        var v10 float64 = (((-2.25 - 1.5) / -2.25) / 1.0)
        v3 += v3
        for fi11 := 0; fi11 < 4; fi11++ {
            v8 = (v2 * (v8 * -1))
            var v12 uint16 = 65535
            fmt.Println(v8, v12, v7)
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        v13 := GenS3{F0: (-884 <= 0), F1: 931, F2: 4294967295}
        _ = v13
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v2)
    case "\n\t":
        var v14 int = -905
        _ = v14
    case "hello":
        v1 = v1
    default:
        v15 := GenS3{F0: (4294967295 < 128), F1: 255, F2: 0}
        _ = v15
    }
    if true {
        v1 += float32(1e10)
    }
    v7 = (v7 % (((v7 & -58)) | 1))
    fmt.Println(v1)
    var v16 bool = (105 > 0)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
