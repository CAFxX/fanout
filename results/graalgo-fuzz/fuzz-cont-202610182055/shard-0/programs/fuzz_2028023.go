package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 float32 = 1e10
    var v4 string = ("foo bar" + "")
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    var v7 string = v4
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            fmt.Println(v4)
            var v10 int8 = -88
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 0, F1: 2147483647}
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(fi8)
            var v13 int16 = -32768
            v7 = "a"
            _ = fi12
            _ = v13
        }
        _ = fi8
        _ = v11
    }
    if ((!false) && (!(false || true))) {
        var v14 int16 = 42
        var v15 uint64 = 0
        delete(v2, 0)
        _ = v14
        _ = v15
    } else {
        fmt.Println(v3)
    }
    _ = len(v6)
    var v16 uint8 = 127
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 int = -555
        fmt.Println(fi17, v16, v7)
        v19 := GenS2{F0: -2147483648, F1: 32767, F2: true}
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int8 = -16
    for fi21 := 0; fi21 < 19; fi21++ {
        if false {
            v22 := GenS2{F0: 2147483647, F1: 282, F2: true}
            _ = v22
        }
        var v23 int32 = 1
        _ = fi21
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
