package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8; F1 string; F2 uint64 }
type GenS3 struct { F0 int64; F1 bool; F2 uint32 }
func genFn0(p0 string) uint16 {
    p0 = p0
    var v5 uint64 = 42
    v6 := GenS1{F0: 127, F1: v5, F2: 1}
    p0 = ((p0 + p0) + ((p0 + p0) + p0))
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint8, p1 int64) string {
    var v7 float32 = 3.14159
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ""
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    v3 := []uint32{2147483647, 65535, 65535}
    var v4 uint32 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    v1[v4] = -389
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 string = "\n\t"
        _ = fi9
        _ = v10
    }
    delete(v1, v4)
    v4 = (v4 * 121)
    var v11 uint64 = 4294967295
    v12 := v2[-1]
    fmt.Println(v11, v11)
    fmt.Println(v11)
    v4 += (v12 ^ v12)
    v11 = (^(v11))
    v4 += (v12 >> ((65535) & 63))
    v13 := GenS2{F0: 0, F1: "hello", F2: (v11 - (v11 * uint64(v12)))}
    if false {
        var v14 float64 = 0.1
        v14 += v14
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        switch fi15 {
        case 127:
            if false {
                v16 := GenS3{F0: 1, F1: (!false), F2: (v12 >> ((855) & 63))}
                var v17 uint16 = 255
                _ = v16
                _ = v17
            }
        case -847:
            if (45 < 42) {
                var v18 uint = 2
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 float32 = 3.14159
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 int = (fi15 >> ((v18) & 63))
                    _ = fi21
                    _ = v22
                }
                _ = v18
            }
        }
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
