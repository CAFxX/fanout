package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int8; F1 int }

func main() {
    v1 := []uint32{}
    var v2 uint32 = 2147483648
    var v3 bool = true
    v4 := make(map[string]bool)
    var v5 uint64 = 4294967295
    var v6 uint64 = v5
    var v7 int16 = 32767
    var v8 int = 89
    var v9 uint16 = 1
    v10 := GenS1{F0: (v8 << ((4294967295) & 63)), F1: v9}
    v11 := v4[("foo bar" + "foo bar")]
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS2{F0: 2, F1: 1, F2: v9}
        _ = len(v4)
        if (!(!true)) {
            v2 -= (v2 * v2)
        } else {
            v14 := GenS3{F0: 100, F1: (v8 ^ v8)}
            var v15 uint32 = v2
            _ = v14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    delete(v4, "hello")
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint16 = v9
        v2 -= 2147483647
        _ = fi16
        _ = v17
    }
    switch "" {
    case "":
        for fi18 := 0; fi18 < 14; fi18++ {
            v6 -= v5
            v2 -= (v2 * (^(v2)))
            v11 = true
            _ = fi18
        }
    case "a":
        v1 = append(v1, (v2 % (((v2 * v2)) | 1)))
    case "\n\t":
        var v19 int64 = 9223372036854775807
        _ = v19
    }
    var v20 int = 32767
    _ = len(v4)
    fmt.Println(v3, v8, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
