package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint8 }
func genFn0(p0 uint64) string {
    var v7 int64 = -1
    p0 += p0
    v8 := GenS1{F0: 42}
    switch "hello" {
    case "foo bar":
        v9 := []int{}
        _ = v9
    case "":
        v10 := []int{-9223372036854775808}
        _ = v10
    default:
        fmt.Println(v7, p0)
    }
    _ = p0
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]bool)
    var v3 uint = v1
    var v4 uint = 64
    v5 := make(map[int]uint32)
    var v6 float64 = (1.5 + 1e10)
    if false {
        v3 += v4
    }
    v1 = 4294967295
    v11 := GenS2{F0: 24, F1: 255, F2: -128}
    v1 = 255
    switch ("\n\t" + ("foo bar" + ("\n\t" + "\n\t"))) {
    case "\n\t":
        v4 += 1
    }
    switch 4294967295 {
    case -269:
        v12 := v2[""]
        _ = v12
    case 32767:
        v6 = 0.1
    case 42:
        var v13 float32 = (float32(float64(0.0)) + 1.5)
        _ = v13
    }
    var v14 float32 = 0.0
    v3 -= v3
    if true {
        fmt.Println(v1, v6)
        v14 = v14
    } else {
        v15 := v5[-955]
        v2["\n\t"] = (2 == 2)
        _ = v15
    }
    v16 := GenS3{F0: 1, F1: 367, F2: 255}
    var v17 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
