package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
type GenS2 struct { F0 uint }
func genFn0(p0 int32) uint64 {
    var v6 bool = false
    var v7 float32 = 1.5
    v8 := []int{-1}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []uint32{437}
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    var v5 int = 128
    var v9 uint64 = 322
    for fi10 := 0; fi10 < 11; fi10++ {
        if ((-2147483648 < 279) && false) {
            v11 := GenS1{F0: 1, F1: "a", F2: (false && (-63 < 100))}
            _ = v11
        } else {
            if false {
                var v12 int16 = 42
                if len(v1) > 0 {
                    v1[0] = 786
                }
                v13 := GenS1{F0: 1, F1: "\n\t", F2: (1 <= 559)}
                _ = v12
                _ = v13
            }
            var v14 bool = true
            var v15 int = 42
            _ = v14
            _ = v15
        }
        var v16 int32 = 945
        _ = fi10
        _ = v16
    }
    v5 = 128
    for fi17 := 0; fi17 < 1; fi17++ {
        switch "\n\t" {
        case "":
            var v18 int32 = 122
            _ = v18
        default:
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 float32 = 3.14159
                _ = fi19
                _ = v20
            }
        }
        _ = fi17
    }
    v5 -= v5
    v21 := GenS2{F0: 18446744073709551615}
    v22 := GenS2{F0: 662}
    var v23 uint8 = 127
    var v24 uint32 = 65535
    v23 = v23
    fmt.Println(v24, v23, v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
