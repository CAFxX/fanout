package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint32 }
func genFn0(p0 int16, p1 int) int32 {
    v7 := []int{(p1 & -694), p1, p1, (p1 >> ((4294967295) & 63))}
    var v8 uint8 = 255
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int8 = 100
    v2 := []string{("\n\t" + (("foo bar" + ("hello" + "hello")) + "hello"))}
    var v3 int16 = 42
    var v4 bool = false
    v5 := make(map[int]bool)
    v6 := []string{}
    v9 := GenS1{F0: int16(v1), F1: -1, F2: 127}
    v5[-9223372036854775808] = (9223372036854775807 != 0)
    var v10 uint16 = 527
    v11 := GenS1{F0: v3, F1: -912, F2: 128}
    v4 = v4
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := GenS2{F0: 0, F1: 127, F2: 42}
        _ = fi12
        _ = v13
    }
    switch "foo bar" {
    case "":
        v10 += 838
    default:
        var v14 uint8 = 0
        _ = v14
    }
    var v15 uint32 = 306
    fmt.Println(v3)
    switch 127 {
    case -259:
        fmt.Println(v15, v1)
    }
    delete(v5, int(v3))
    _ = len(v2)
    v5[637] = (4294967295 >= 4294967295)
    var v16 int8 = (v1 << ((42) & 63))
    fmt.Println(v4, v15, v1)
    fmt.Println(v1)
    v4 = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
