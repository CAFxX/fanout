package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint64 = v1
    var v3 uint16 = 255
    var v4 uint = 733
    var v5 uint16 = 0
    v6 := GenS1{F0: 771, F1: ("\n\t" + "\n\t"), F2: 127}
    var v7 uint32 = 0
    v1 = (v2 << (((v4 + (^(v4)))) & 63))
    v7 += v7
    v8 := []int{32767}
    var v9 int8 = 0
    if true {
        v8 = append(v8, 137)
        fmt.Println(v2, v4)
        v10 := GenS1{F0: 616, F1: ((("foo bar" + "hello") + "a") + ""), F2: -4}
        _ = v10
    }
    switch 32768 {
    case 901:
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v1, v1, v5)
            if len(v8) > 0 {
                v8[0] = fi11
            }
            _ = fi11
        }
    case -956:
        v4 -= 2
    case -830:
        v1 += (v2 - v1)
    }
    var v12 bool = true
    var v13 uint64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
