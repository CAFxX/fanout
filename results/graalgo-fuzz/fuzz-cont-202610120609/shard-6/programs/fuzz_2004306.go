package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
func genFn0(p0 uint64) int32 {
    v5 := GenS1{F0: 65535, F1: 224, F2: 4294967295}
    var v6 int16 = 32767
    _ = p0
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := []uint32{321, 1, 528}
    var v2 int64 = 42
    v3 := []uint32{}
    v4 := make(map[int]bool)
    if true {
        v2 += (v2 * (v2 << ((65535) & 63)))
        v7 := GenS1{F0: 0, F1: 255, F2: 2147483648}
        _ = v7
    }
    if (9223372036854775807 > -553) {
        v8 := v4[-478]
        _ = v8
    }
    v9 := v4[-2147483648]
    var v10 int16 = -1
    v2 += 42
    v11 := v4[2147483647]
    var v12 string = ""
    v12 = ("" + ((v12 + v12) + ("\n\t" + ((v12 + v12) + (v12 + (v12 + v12))))))
    var v13 int = 127
    var v14 uint = 2
    fmt.Println(v13, v11)
    v11 = true
    v14 = (v14 ^ (v14 ^ v14))
    v11 = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
