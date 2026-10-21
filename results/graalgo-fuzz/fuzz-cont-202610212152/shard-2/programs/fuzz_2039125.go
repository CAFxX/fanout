package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int32 = 1
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    fmt.Println(v1)
    v4 := []int{32767}
    switch "a" {
    case "":
        var v5 int32 = 42
        _ = v5
    default:
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: 18446744073709551615}
            v1 = v1
            var v8 uint32 = 4294967295
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    fmt.Println(v1)
    if (!(!false)) {
        v9 := GenS1{F0: 18446744073709551615}
        var v10 float64 = (0.1 - ((0.1 / 1.5) - -2.25))
        _ = v10
        _ = v9
    }
    var v11 bool = (true || (!((1 < 2) || (false || true))))
    var v12 float32 = float32(3.14159)
    if (v11 && (false && false)) {
        var v13 int32 = (v1 & v1)
        v14 := GenS1{F0: 0}
        fmt.Println(v12, v1)
        _ = v13
        _ = v14
    } else {
        v11 = (!(!v11))
        v15 := GenS1{F0: 649}
        switch ("\n\t" + ("foo bar" + "foo bar")) {
        case "hello":
            v1 = v1
        default:
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 uint8 = 128
                _ = fi16
                _ = v17
            }
        }
        _ = v15
    }
    fmt.Println(v12)
    var v18 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
