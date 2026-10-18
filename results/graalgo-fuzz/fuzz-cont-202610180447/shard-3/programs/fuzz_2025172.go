package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int; F2 uint64 }
func genFn0(p0 int32) int {
    var v5 int32 = p0
    v6 := GenS1{F0: 4294967295, F1: 128}
    v7 := GenS2{F0: 4294967295, F1: -9223372036854775808, F2: 0}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 bool = (true && ("" <= "\n\t"))
    v2 := make(map[string]bool)
    v3 := []uint32{}
    var v4 int16 = -754
    v8 := GenS2{F0: 2, F1: 32768, F2: 0}
    var v9 int32 = 1
    v1 = ((!(0 > 853)) && true)
    var v10 int8 = -1
    v11 := GenS2{F0: 2, F1: 180, F2: 18446744073709551615}
    switch 1 {
    case 4294967295:
        v12 := GenS2{F0: 42, F1: 127, F2: 1}
        _ = v12
    case 2147483647:
        delete(v2, "hello")
    }
    switch ("\n\t" + "hello") {
    case "":
        v13 := GenS1{F0: 0, F1: 128}
        _ = v13
    case "\n\t":
        fmt.Println(v9)
    case "a":
        var v14 int8 = (v10 | (v10 & (v10 + v10)))
        _ = v14
    default:
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS2{F0: 326, F1: fi15, F2: 4294967295}
            _ = fi15
            _ = v16
        }
    }
    var v17 uint = 684
    v18 := GenS2{F0: 18446744073709551615, F1: 4294967295, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
