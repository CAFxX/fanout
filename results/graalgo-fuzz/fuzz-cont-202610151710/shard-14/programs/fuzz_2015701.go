package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int64 = 42
    var v2 int32 = 66
    v3 := make(map[uint32]int)
    var v4 string = ("hello" + ("hello" + "hello"))
    fmt.Println(v4)
    if false {
        var v5 uint = 4294967295
        var v6 bool = true
        _ = v5
        _ = v6
    }
    fmt.Println(v4)
    v7 := []int{-1}
    v3[0] = int(v2)
    v2 += (v2 >> ((255) & 63))
    for fi8 := 0; fi8 < 6; fi8++ {
        v1 -= (v1 - v1)
        var v9 bool = false
        switch 926 {
        case 32768:
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: 0}
                _ = fi10
                _ = v11
            }
        case -677:
            _ = len(v3)
        case 127:
            var v12 int = (fi8 ^ -160)
            _ = v12
        default:
            for fi13 := 0; fi13 < 4; fi13++ {
                fmt.Println(v1)
                var v14 uint8 = 1
                fmt.Println(fi13, v2)
                _ = fi13
                _ = v14
            }
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v2)
    v15 := GenS1{F0: 1}
    var v16 int64 = v1
    for fi17 := 0; fi17 < 1; fi17++ {
        v1 += v16
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
