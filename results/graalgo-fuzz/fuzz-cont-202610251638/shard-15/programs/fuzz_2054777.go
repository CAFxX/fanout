package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32) int8 {
    var v6 int32 = -157
    fmt.Println(p0)
    var v7 float32 = 3.14159
    _ = p0
    _ = v6
    _ = v7
    return -128
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 850
    var v3 uint8 = 127
    var v4 uint16 = v2
    var v5 uint = 47
    v8 := GenS1{F0: (^(v3))}
    var v9 int16 = 29
    switch 0 {
    case 9223372036854775807:
        fmt.Println(v9, v9, v2)
    case 42:
        v4 = (v4 ^ 65535)
    case 32768:
        v10 := GenS2{F0: "\n\t"}
        _ = v10
    }
    switch -9223372036854775808 {
    case 4294967295:
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(v5, v3)
            fmt.Println(v9, v9, v5)
            _ = fi11
        }
    case 495:
        v12 := GenS3{F0: (2 != 502)}
        _ = v12
    }
    v3 = v3
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := []int{(fi13 % ((-1) | 1))}
        _ = fi13
        _ = v14
    }
    fmt.Println(v9, v5)
    var v15 int8 = -1
    var v16 float32 = 3.14159
    v17 := GenS2{F0: "\n\t"}
    var v18 string = (((("a" + "\n\t") + "") + "hello") + "a")
    v19 := []int{-1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
