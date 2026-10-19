package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 4294967295
    var v3 uint32 = 910
    var v4 int32 = 1
    var v5 float32 = 3.14159
    var v6 int = -9223372036854775808
    fmt.Println(v5)
    fmt.Println(v4, v3)
    switch "a" {
    case "hello":
        switch "hello" {
        case "":
            var v7 int = (v6 << ((4294967295) & 63))
            _ = v7
        }
    case "":
        v4 += -1
    }
    v8 := GenS1{F0: 924, F1: 42, F2: 42}
    fmt.Println(v3, v6, v4)
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: -9223372036854775808, F1: 0, F2: 100}
        var v12 uint64 = v2
        _ = fi10
        _ = v11
        _ = v12
    }
    delete(v1, v9)
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int8 = 0
        v1[(v9 + v9)] = ((((0.0 * 0.0) + ((3.14159 / (1.5 + -2.25)) - 1.5)) <= 1.5) || true)
        v3 += 2147483648
        _ = fi13
        _ = v14
    }
    delete(v1, v9)
    v4 -= (v4 & v4)
    var v15 float32 = float32(1.5)
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS2{F0: (v2 & (v2 * v2))}
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 uint16 = 65535
            fmt.Println(v4, v3, v6)
            v20 := v1[("a" + v9)]
            var v21 int8 = 1
            var v22 uint32 = v3
            switch ("\n\t" + v9) {
            case "foo bar":
                var v23 int16 = 42
                _ = v23
            default:
                v20 = v20
            }
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
            _ = v22
        }
        switch "\n\t" {
        case "a":
            var v24 bool = (!(18446744073709551615 != 4294967295))
            _ = v24
        }
        _ = fi16
        _ = v17
    }
    v25 := GenS3{F0: 2, F1: 1}
    var v26 float64 = 0.1
    v27 := GenS1{F0: -773, F1: 65535, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
