package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    v2 := []int{385, 32768}
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 uint32 = 2147483648
        _ = fi4
        _ = v5
    }
    var v6 bool = (127 < -1)
    v6 = v6
    v6 = (true || v6)
    var v7 int16 = 528
    v8 := v1[((("hello" + "foo bar") + ("" + "a")) + "")]
    if v6 {
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                v8 -= (fi9 << ((0) & 63))
                _ = fi10
            }
            var v11 uint32 = 4294967295
            _ = fi9
            _ = v11
        }
        switch -9223372036854775808 {
        case -658:
            var v12 uint = 42
            _ = v12
        case -9223372036854775808:
            delete(v3, 504)
        }
        var v13 int16 = (-(v7))
        _ = v13
    }
    fmt.Println(v7)
    if (127 <= 216) {
        v14 := GenS1{F0: false}
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
