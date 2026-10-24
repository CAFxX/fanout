package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := []uint32{2, 4294967295}
    v2 := make(map[uint32]int)
    var v3 int64 = -9223372036854775808
    v4 := []string{"foo bar", "foo bar", (("a" + "a") + "foo bar"), "hello"}
    var v5 uint8 = 127
    v2[42] = 32768
    var v6 uint = 2
    v5 = 128
    var v7 uint32 = 2147483648
    v7 = 391
    v7 = 2
    if (-1 == -1) {
        v8 := GenS1{F0: uint32(v6), F1: 9223372036854775807}
        v3 = v3
        _ = v8
    }
    v9 := GenS1{F0: v7, F1: 2147483647}
    if (-1 <= -128) {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 string = ("" + "\n\t")
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 int32 = 0
                var v14 uint32 = 2147483648
                var v15 int16 = 1000
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            v6 = uint(v7)
            _ = fi10
            _ = v11
        }
    }
    var v16 uint8 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
