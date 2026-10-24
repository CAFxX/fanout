package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int64, p1 int8) int16 {
    v4 := []int{}
    var v5 float64 = (((-2.25 / (1e10 * (1e10 / 1.5))) + 3.14159) / 1.5)
    p0 -= 690
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 2
    v3 := []uint32{0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v1)
    if true {
        if (!((255 <= 65535) || ((!true) && (!(false && (2147483648 == 42)))))) {
            if false {
                v2 += (v2 >> ((1) & 63))
                if len(v3) > 0 {
                    v3[0] = 2
                }
            }
        }
        fmt.Println(v2)
    }
    switch 1 {
    case 128:
        var v6 int64 = -967
        _ = v6
    case -9223372036854775808:
        var v7 float32 = 1e10
        _ = v7
    case -1:
        v2 = 0
    default:
        for fi8 := 0; fi8 < 7; fi8++ {
            if true {
                var v9 uint32 = 2147483647
                _ = v9
            } else {
                if true {
                    var v10 int16 = -221
                    v10 += v10
                    var v11 int64 = 0
                    _ = v10
                    _ = v11
                }
            }
            if len(v3) > 0 {
                v3[0] = 2
            }
            _ = fi8
        }
    }
    fmt.Println(v2)
    v2 = (v2 & (v2 | v2))
    switch 127 {
    case 4294967295:
        fmt.Println(v2)
    }
    var v12 bool = (true || (-812 != 1))
    switch "foo bar" {
    case "foo bar":
        v13 := GenS1{F0: (-102 <= 30)}
        _ = v13
    case "":
        v3 = append(v3, 1)
    default:
        v14 := v1[669]
        _ = v14
    }
    v12 = true
    v12 = true
    if v12 {
        var v15 bool = v12
        var v16 bool = (218 != 42)
        _ = v15
        _ = v16
    } else {
        if (3.14159 <= 1.5) {
            for fi17 := 0; fi17 < 5; fi17++ {
                _ = len(v1)
                _ = fi17
            }
            v18 := GenS2{F0: 65535}
            v2 -= v2
            _ = v18
        }
        _ = len(v3)
    }
    v19 := GenS3{F0: (v2 >> ((2) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
}
