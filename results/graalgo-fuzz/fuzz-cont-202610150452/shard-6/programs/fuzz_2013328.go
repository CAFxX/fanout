package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]string)
    var v5 uint16 = 0
    var v6 int8 = -14
    v6 = (v6 ^ v6)
    fmt.Println(v5, v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int32 = -567
        _ = fi7
        _ = v8
    }
    v6 = 44
    var v9 int = -984
    _ = len(v1)
    var v10 int8 = (v6 >> ((993) & 63))
    var v11 uint64 = 618
    if len(v1) > 0 {
        v1[0] = 1
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS1{F0: 1, F1: 0, F2: 1}
        _ = fi12
        _ = v13
    }
    var v14 int64 = 392
    switch v9 {
    case 2147483647:
        var v15 int8 = v10
        _ = v15
    case -168:
        fmt.Println(v14, v11)
    case -9223372036854775808:
        v10 -= -20
    }
    fmt.Println(v10, v10)
    if true {
        if ((v14 ^ v14) > (v14 | v14)) {
            var v16 string = ("hello" + "foo bar")
            _ = v16
        } else {
            switch v9 {
            case 127:
                fmt.Println(v6)
            case 2147483647:
                v17 := GenS1{F0: 9223372036854775807, F1: 42, F2: int32(v9)}
                _ = v17
            }
        }
        var v18 float64 = 1.5
        _ = v18
    } else {
        var v19 string = ("" + "a")
        _ = v19
    }
    v20 := GenS2{F0: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
