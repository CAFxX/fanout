package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 string }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2
    v3 := make(map[string]uint32)
    if ("\n\t" > ("\n\t" + "hello")) {
        if true {
            v4 := v1[680]
            v4 = (false || v4)
            v5 := GenS1{F0: 0, F1: 282, F2: "hello"}
            _ = v4
            _ = v5
        } else {
            delete(v1, -612)
            var v6 int8 = 15
            var v7 int64 = 0
            _ = v6
            _ = v7
        }
        v8 := []int{968, -785, 32768}
        v2 -= 1
        _ = v8
    }
    v9 := v3[""]
    v9 = v2
    var v10 int8 = -1
    v10 += (v10 | (v10 | (v10 ^ -1)))
    v11 := []int{214, 67}
    var v12 int16 = 1000
    var v13 int8 = 51
    _ = len(v1)
    var v14 int = -344
    delete(v1, (v14 >> ((0) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
