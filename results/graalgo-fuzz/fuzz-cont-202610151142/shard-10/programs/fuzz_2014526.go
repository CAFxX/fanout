package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint; F1 bool; F2 int }
type GenS4 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint32 = 4294967295
    var v2 int = int(v1)
    var v3 int16 = 1000
    v4 := make(map[int]bool)
    v5 := GenS1{F0: 1, F1: v2}
    v6 := []int{1, (v2 / (((v2 | v2)) | 1)), (v2 | (v2 - v2)), (v2 >> ((255) & 63))}
    _ = len(v6)
    v7 := v4[v2]
    var v8 string = (("" + (("hello" + "a") + "")) + "\n\t")
    if ((v2 << ((4294967295) & 63)) <= v2) {
        v9 := GenS2{F0: 0, F1: 128}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v1)
            fmt.Println(v2)
            if (v7 || v7) {
                var v11 int16 = 42
                v12 := GenS3{F0: 2, F1: (v7 && v7), F2: (v2 * (^(fi10)))}
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint64 = 1
            v15 := GenS1{F0: 255, F1: (fi13 / (((v2 | fi13)) | 1))}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    fmt.Println(v2, v8)
    if false {
        v16 := GenS4{F0: 2, F1: 111}
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
