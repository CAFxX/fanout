package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 bool = (113 > 42)
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    v4 := make(map[uint32]bool)
    v1 = (255 == 999)
    for fi5 := 0; fi5 < 14; fi5++ {
        v1 = (-1 == 42)
        v6 := GenS1{F0: 0}
        fmt.Println(fi5, fi5)
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 402
    var v8 uint = 348
    fmt.Println(v8, v8)
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int32 = -2147483648
        for fi11 := 0; fi11 < 10; fi11++ {
            fmt.Println(v1)
            var v12 string = "hello"
            v13 := []int{(fi9 % ((int(v10)) | 1))}
            switch fi9 {
            case -550:
                v14 := GenS1{F0: 100}
                _ = v14
            }
            var v15 int16 = 860
            _ = fi11
            _ = v12
            _ = v13
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v1, v1, v7)
    v16 := GenS1{F0: -128}
    v7 += v7
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
