package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 int; F2 int16 }

func main() {
    var v1 uint = 2
    v2 := []string{"a", ("hello" + "\n\t"), "a"}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]bool)
    v7 := GenS1{F0: -2147483648, F1: 42}
    v1 -= 4294967295
    v1 -= v1
    fmt.Println(v1)
    var v8 uint16 = 65535
    v9 := GenS2{F0: -32768, F1: 531, F2: 0}
    var v10 int = 32768
    v11 := GenS3{F0: 18446744073709551615}
    switch v10 {
    case 0:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS4{F0: 42}
            var v14 uint = (v1 & v1)
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    v4[2147483647] = "hello"
    v15 := GenS5{F0: 0, F1: (v10 | (v10 ^ v10)), F2: 0}
    fmt.Println(v10, v1)
    switch v10 {
    case 356:
        var v16 int64 = -1
        _ = v16
    case 42:
        v1 = 18446744073709551615
    default:
        v8 = v8
    }
    switch "\n\t" {
    case "hello":
        var v17 int8 = 1
        _ = v17
    default:
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 bool = (3.14159 <= (((0.0 + -2.25) / 0.1) / 1.0))
            for fi20 := 0; fi20 < 10; fi20++ {
                fmt.Println(fi18, fi20, fi18)
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
