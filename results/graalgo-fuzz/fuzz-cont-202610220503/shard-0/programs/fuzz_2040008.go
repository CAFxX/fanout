package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 string; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0(p0 int32, p1 bool) uint {
    var v5 int = -892
    p0 = (-(p0))
    var v6 uint16 = 65535
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint32, p1 uint64) int16 {
    var v7 uint16 = 721
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: fi8, F1: 184}
        v10 := GenS2{F0: ("foo bar" + "hello"), F1: int64(fi8), F2: (p1 | p1)}
        v11 := GenS2{F0: "a", F1: 9223372036854775807, F2: 23}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 32767
}

func main() {
    var v1 uint16 = 32767
    var v2 float32 = (1.5 * 3.14159)
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    var v12 uint = 4294967295
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 bool = (18446744073709551615 < 2)
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 int16 = 1000
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v1 = v1
    var v17 int = 2147483647
    fmt.Println(v2)
    v2 += -2.25
    fmt.Println(v2, v1)
    v12 -= (^(v12))
    v3[("foo bar" + "foo bar")] = ("a" + "")
    delete(v4, "hello")
    switch "" {
    case "foo bar":
        v18 := GenS3{F0: (v1 >> ((165) & 63)), F1: 0}
        _ = v18
    default:
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 uint32 = 42
            _ = fi19
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
