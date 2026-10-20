package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 int }
func genFn0(p0 int32) uint16 {
    var v4 uint64 = 2
    var v5 uint8 = 231
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 42
    v3 := []uint32{658, 0, 349, 42}
    v2 = v2
    v6 := GenS1{F0: 0}
    v2 = -944
    v7 := v1[1]
    var v8 bool = (1e10 > float64(3.14159))
    v9 := GenS1{F0: 42}
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 string = "a"
        fmt.Println(v7, fi10)
        v11 = (v11 + (v11 + v11))
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v2 = (v2 >> ((4294967295) & 63))
        _ = fi12
    }
    v13 := GenS2{F0: (v7 + 74), F1: -54, F2: -943}
    for fi14 := 0; fi14 < 3; fi14++ {
        switch "hello" {
        case "\n\t":
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 int64 = (v2 ^ 58)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi15
                _ = v16
            }
        case "":
            for fi17 := 0; fi17 < 19; fi17++ {
                v8 = (!(uint(v7) <= 42))
                _ = fi17
            }
        default:
            switch ("\n\t" + ("hello" + "foo bar")) {
            case "hello":
                var v18 uint64 = 672
                _ = v18
            case "":
                v8 = (2 <= 7)
            case "a":
                var v19 float32 = 0.0
                _ = v19
            }
        }
        _ = fi14
    }
    v20 := GenS1{F0: 4294967295}
    fmt.Println(v8)
    switch 930 {
    case 9:
        fmt.Println(v2)
    case -1:
        v21 := GenS1{F0: 4294967295}
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
