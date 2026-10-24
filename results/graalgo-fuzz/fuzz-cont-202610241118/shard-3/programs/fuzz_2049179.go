package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int; F1 int16 }

func main() {
    var v1 string = "a"
    var v2 uint16 = 42
    var v3 int64 = 9223372036854775807
    v4 := make(map[uint32]bool)
    var v5 float32 = 1e10
    var v6 uint = 0
    var v7 int32 = 250
    v8 := GenS1{F0: v2, F1: v7}
    v9 := []int{2147483647, 0, 2147483647}
    var v10 int8 = 100
    fmt.Println(v5)
    var v11 bool = false
    if (4294967295 < 256) {
        v2 += uint16(v6)
        for fi12 := 0; fi12 < 15; fi12++ {
            switch ("\n\t" + v1) {
            case "a":
                v13 := GenS2{F0: fi12, F1: int16(v3)}
                _ = v13
            case "foo bar":
                var v14 string = (v1 + (v1 + "a"))
                _ = v14
            case "hello":
                fmt.Println(fi12)
            default:
                var v15 int32 = (v7 | -2147483648)
                _ = v15
            }
            _ = fi12
        }
    }
    var v16 int8 = -128
    var v17 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
