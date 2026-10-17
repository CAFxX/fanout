package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []string{"\n\t", ""}
    v2 := []string{(("\n\t" + "") + "foo bar")}
    var v3 uint16 = 32767
    v4 := make(map[int]bool)
    var v5 float64 = (0.0 * (1.5 - 3.14159))
    v6 := GenS1{F0: 1}
    var v7 uint32 = 1
    v7 = (v7 >> ((65535) & 63))
    if true {
        fmt.Println(v7, v7)
        if (787 >= 2147483647) {
            fmt.Println(v5)
            var v8 string = "\n\t"
            var v9 float32 = 1e10
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 20; fi10++ {
            if (!true) {
                v11 := GenS1{F0: -2147483648}
                _ = v11
            } else {
                fmt.Println(v5, v5, v7)
            }
            var v12 int32 = 42
            _ = fi10
            _ = v12
        }
    }
    var v13 uint32 = v7
    v5 += v5
    v14 := GenS1{F0: 1}
    var v15 uint32 = (v7 | 0)
    var v16 int64 = -673
    v17 := GenS2{F0: 42}
    v15 = (v13 ^ 2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
