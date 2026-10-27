package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 uint16 }
func genFn0() bool {
    v5 := GenS1{F0: 9223372036854775807, F1: 42}
    var v6 uint8 = 131
    _ = v5
    _ = v6
    return (639 >= -151)
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]uint32)
    var v3 string = ""
    var v4 uint8 = 2
    fmt.Println(v3, v1)
    var v7 uint32 = 65535
    fmt.Println(v4, v3)
    var v8 float32 = 0.1
    _ = len(v2)
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := []int{(fi9 + (fi9 / (((fi9 - 32768)) | 1)))}
        v11 := v2[v3]
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS2{F0: 2, F1: uint(v4), F2: (v1 * (v1 & (v1 - v4)))}
        v3 = v3
        _ = fi12
        _ = v13
    }
    var v14 uint64 = 1
    v1 = 2
    var v15 int = 2147483647
    v16 := GenS3{F0: uint16(v14)}
    switch 2147483647 {
    case 128:
        v7 += (v7 | v7)
    case 1:
        v17 := GenS4{F0: -2147483648, F1: 434}
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v7 = (v7 + 2147483648)
        v3 = ("" + "hello")
        _ = fi18
    }
    _ = len(v2)
    fmt.Println(v14)
    var v19 uint8 = (v4 - (v1 - (v1 - v4)))
    v15 = (v15 - v15)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
