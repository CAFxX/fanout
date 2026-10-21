package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint64) int {
    fmt.Println(p0)
    var v5 uint64 = (p0 >> ((1) & 63))
    var v6 uint8 = 238
    switch 314 {
    case 32767:
        v7 := []int{1, 32768}
        _ = v7
    case 0:
        var v8 int16 = 42
        _ = v8
    case 127:
        v5 += p0
    }
    _ = p0
    _ = v5
    _ = v6
    return 955
}
func genFn1(p0 uint32, p1 bool, p2 uint) uint {
    var v9 string = "\n\t"
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: 2}
        p2 = (p2 << ((p2) & 63))
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    var v3 uint8 = 0
    v4 := make(map[int]string)
    var v12 int = genFn0(42)
    fmt.Println(v3)
    v13 := GenS1{F0: (^(v3))}
    switch "hello" {
    case "foo bar":
        fmt.Println(v2, v3, v3)
    case "hello":
        fmt.Println(v2, v12)
    case "\n\t":
        v3 = 0
    default:
        v2 = (1e10 != 1.5)
    }
    fmt.Println(v12, v2)
    fmt.Println(v12, v12)
    var v14 uint16 = 0
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int32 = 938
        _ = len(v4)
        _ = fi15
        _ = v16
    }
    v12 += v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
