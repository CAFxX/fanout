package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }

func main() {
    v1 := []int{32767, 0, -2147483648, 32767}
    v2 := []string{("hello" + (("" + ((("" + (("\n\t" + "") + "hello")) + ("\n\t" + "hello")) + "\n\t")) + "hello"))}
    v3 := make(map[int]int)
    var v4 float32 = -2.25
    v4 -= v4
    v5 := GenS1{F0: 65535, F1: 1, F2: -1}
    v4 += v4
    if false {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 uint16 = 2
            var v8 string = ("hello" + ("foo bar" + "hello"))
            _ = fi6
            _ = v7
            _ = v8
        }
        v4 += v4
    } else {
        v9 := GenS1{F0: 1, F1: 0, F2: -1}
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: 1, F2: -872}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 uint64 = 1
    v11 = (v11 * 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
