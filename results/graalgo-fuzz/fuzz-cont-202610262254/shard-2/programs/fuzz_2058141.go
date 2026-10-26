package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int32 }
func genFn0(p0 string, p1 int64, p2 int8) uint {
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{fi6, fi6, (fi6 ^ fi6)}
        _ = fi6
        _ = v7
    }
    switch int(p1) {
    case 127:
        p0 = ((("a" + p0) + "") + p0)
    default:
        if (32767 <= 4294967295) {
            p0 = ""
            fmt.Println(p2)
            v8 := GenS1{F0: 1, F1: p0, F2: 127}
            _ = v8
        }
    }
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return 616
}

func main() {
    v1 := []uint32{630, 2147483647}
    v2 := make(map[int]int)
    v3 := []uint32{}
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    var v9 int16 = 773
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS2{F0: 128}
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 uint64 = 2
            var v14 float32 = (float32(3.14159) + 1.5)
            v13 += 18446744073709551615
            v15 := GenS3{F0: 0, F1: fi10}
            v3 = append(v3, 4294967295)
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    v16 := GenS4{F0: 0}
    fmt.Println(v9)
    var v17 uint32 = 590
    v17 += 737
    var v18 uint8 = 255
    if (255 == 65535) {
        v19 := GenS3{F0: 255, F1: 1}
        _ = v19
    } else {
        v18 = v18
        v20 := GenS1{F0: 752, F1: "\n\t", F2: v18}
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
