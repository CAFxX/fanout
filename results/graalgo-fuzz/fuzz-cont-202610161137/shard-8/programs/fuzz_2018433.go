package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }

func main() {
    v1 := []string{"hello", ("foo bar" + "a"), ("\n\t" + "")}
    var v2 uint = 0
    var v3 uint8 = 42
    v4 := make(map[string]int)
    v5 := []int{32767, 804, 42}
    fmt.Println(v3, v2)
    var v6 int16 = -1
    v7 := GenS1{F0: 876, F1: -753, F2: ""}
    v6 = 42
    if true {
        var v8 int64 = 0
        _ = v8
    }
    v6 += (v6 + (v6 / (((v6 ^ v6)) | 1)))
    if (255 > 32767) {
        var v9 int16 = 234
        v3 += (v3 + 128)
        _ = v9
    }
    var v10 uint16 = 2
    v11 := v4["hello"]
    v12 := v4["foo bar"]
    var v13 int = v11
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
