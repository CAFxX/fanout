package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16; F1 uint; F2 int16 }
func genFn0(p0 string, p1 string, p2 uint16) int8 {
    v4 := GenS1{F0: 65535}
    v5 := GenS2{F0: p1, F1: 0, F2: 65535}
    for fi6 := 0; fi6 < 3; fi6++ {
        p1 = p0
        fmt.Println(p2, fi6)
        v7 := GenS2{F0: "foo bar", F1: 135, F2: 643}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -27
}

func main() {
    v1 := []uint32{65535}
    var v2 int16 = 803
    v3 := make(map[uint32]uint32)
    var v8 float32 = float32(1e10)
    v2 = (v2 + (^(v2)))
    switch "\n\t" {
    case "":
        if (-10 <= -1) {
            v8 = 1e10
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS2{F0: "hello", F1: 63, F2: 42}
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: 425}
            _ = v11
        } else {
            v1 = append(v1, 867)
            switch -2147483648 {
            case -23:
                v12 := GenS3{F0: 978, F1: 2, F2: 1}
                _ = v12
            case 9223372036854775807:
                var v13 int8 = genFn0("a", "", 1)
                _ = v13
            case 127:
                v14 := GenS2{F0: "", F1: 65535, F2: 0}
                _ = v14
            }
        }
    }
    v15 := GenS1{F0: 2}
    v8 = (((v8 + (v8 * v8)) / (v8 - float32(1.5))) / ((float32(1e10) + -2.25) / v8))
    var v16 int = 32767
    v2 -= (v2 & v2)
    var v17 bool = (42 > 252)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
