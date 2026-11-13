package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0(p0 int16) int8 {
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    return 119
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]int)
    var v3 uint32 = 829
    var v4 uint = v1
    fmt.Println(v4)
    v1 = 18446744073709551615
    v5 := GenS1{F0: 255, F1: v3, F2: -32768}
    for fi6 := 0; fi6 < 9; fi6++ {
        if true {
            for fi7 := 0; fi7 < 9; fi7++ {
                if true {
                    var v8 bool = false
                    v4 -= v1
                    _ = v8
                }
                v4 = v4
                _ = fi7
            }
            v1 -= v4
            var v9 uint = v1
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS2{F0: 2}
    _ = len(v2)
    v11 := []int{629}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    switch -962 {
    case 42:
        var v12 bool = (!((v4 >> (((v4 % ((uint(v3)) | 1))) & 63)) != 1))
        _ = v12
    case -898:
        for fi13 := 0; fi13 < 8; fi13++ {
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            fmt.Println(v4, v4)
            v14 := GenS3{F0: 0, F1: 42}
            _ = fi13
            _ = v14
        }
    case 910:
        var v15 int16 = -32768
        _ = v15
    default:
        for fi16 := 0; fi16 < 14; fi16++ {
            v4 += v4
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := GenS2{F0: 4294967295}
                var v19 string = "a"
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
