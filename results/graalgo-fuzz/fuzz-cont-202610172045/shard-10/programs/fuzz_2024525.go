package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{0, 188, -458}
    var v3 uint = 0
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    var v6 uint64 = 365
    var v7 uint64 = v6
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 32767, F1: 100, F2: 4294967295}
        _ = fi8
        _ = v9
    }
    delete(v5, "hello")
    v2 = append(v2, 127)
    var v10 int64 = -9223372036854775808
    var v11 int32 = 410
    fmt.Println(v7, v10)
    var v12 int32 = v11
    switch "" {
    case "":
        if (v3 >= v3) {
            for fi13 := 0; fi13 < 16; fi13++ {
                v2 = append(v2, 633)
                for fi14 := 0; fi14 < 13; fi14++ {
                    var v15 int8 = -124
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
            var v16 int32 = (-(v12))
            if true {
                v11 -= (v12 & v16)
                var v17 uint64 = (v7 - uint64(v3))
                _ = len(v1)
                _ = v17
            }
            _ = v16
        }
    }
    delete(v4, 1)
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 uint32 = 2147483647
        var v20 uint64 = 0
        v6 = v7
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 int = -1
    v11 += (v12 << (((v3 - 337)) & 63))
    fmt.Println(v10, v11)
    switch ("\n\t" + "hello") {
    case "hello":
        v22 := GenS2{F0: v12}
        _ = v22
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        fmt.Println(v10, v3)
    default:
        fmt.Println(v10, v11)
    }
    var v23 int32 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
