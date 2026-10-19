package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 60
    v3 := make(map[uint32]int)
    switch 32768 {
    case 32768:
        for fi4 := 0; fi4 < 16; fi4++ {
            v5 := []int{(fi4 & fi4), (fi4 / ((fi4) | 1)), fi4, (fi4 % (((fi4 + (fi4 - (fi4 | (fi4 << ((461) & 63)))))) | 1))}
            v2 += 255
            v2 += v2
            _ = fi4
            _ = v5
        }
    case 732:
        v6 := GenS1{F0: -2147483648}
        _ = v6
    }
    var v7 uint64 = v1
    v8 := GenS1{F0: 42}
    var v9 string = (("a" + "hello") + "hello")
    v10 := GenS2{F0: -1}
    for fi11 := 0; fi11 < 10; fi11++ {
        v7 = 918
        _ = fi11
    }
    var v12 int32 = 42
    var v13 bool = true
    v14 := GenS2{F0: 105}
    var v15 int32 = (v12 + v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
