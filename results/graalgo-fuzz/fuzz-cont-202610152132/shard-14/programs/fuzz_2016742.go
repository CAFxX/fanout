package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int16 }

func main() {
    v1 := []uint32{42, 572, 2147483647, 2}
    v2 := make(map[int]uint32)
    v3 := []int{32767, 4294967295}
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: 0}
        var v6 bool = false
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int = 1
    if true {
        var v8 float32 = ((1e10 - 0.1) + 3.14159)
        _ = v8
    }
    var v9 uint32 = 2
    var v10 uint32 = 429
    var v11 uint8 = 42
    if false {
        if false {
            switch ((("" + "foo bar") + "\n\t") + "hello") {
            case "hello":
                var v12 int8 = int8(v9)
                _ = v12
            default:
                for fi13 := 0; fi13 < 18; fi13++ {
                    v11 = 127
                    if len(v3) > 0 {
                        v3[0] = -664
                    }
                    _ = fi13
                }
            }
        } else {
            v14 := GenS2{F0: (^(v10)), F1: -32768}
            _ = v14
        }
        v15 := GenS1{F0: -240}
        _ = v15
    } else {
        delete(v2, (v7 & v7))
        switch ("\n\t" + "hello") {
        case "foo bar":
            v9 += 464
        case "a":
            v16 := GenS1{F0: -906}
            _ = v16
        default:
            var v17 float64 = 1.5
            _ = v17
        }
    }
    v7 = 977
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
