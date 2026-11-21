package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64, p1 int64, p2 int32) int {
    if (("\n\t" == "a") || (65535 == 255)) {
        v6 := GenS1{F0: (p2 / ((p2) | 1)), F1: 883, F2: 914}
        _ = v6
    } else {
        var v7 uint8 = 42
        for fi8 := 0; fi8 < 7; fi8++ {
            if ((!(793 == 960)) || false) {
                p2 = p2
            }
            var v9 int32 = int32(fi8)
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v10 := []int{}
    var v11 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 0
}

func main() {
    v1 := []uint32{2147483647}
    v2 := make(map[string]string)
    v3 := make(map[int]bool)
    var v4 string = ("" + "foo bar")
    var v5 float32 = 0.0
    v4 = ((v4 + "a") + v4)
    switch -186 {
    case 4294967295:
        if (false || (32767 > 42)) {
            v5 = (1.5 - (v5 + 1.5))
        }
    case -971:
        v4 = v4
    case 9223372036854775807:
        var v12 float32 = v5
        _ = v12
    default:
        var v13 bool = (!true)
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 string = (v4 + ("a" + ("\n\t" + v4)))
        v16 := GenS2{F0: 886}
        for fi17 := 0; fi17 < 8; fi17++ {
            if len(v1) > 0 {
                v1[0] = 1
            }
            if ((-2.25 > 3.14159) && true) {
                v18 := GenS2{F0: -255}
                v4 = "foo bar"
                _ = v18
            }
            v15 = (v4 + v15)
            _ = fi17
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    v19 := GenS3{F0: 1}
    _ = len(v2)
    v5 = 0.1
    _ = len(v1)
    if true {
        if true {
            if len(v1) > 0 {
                v1[0] = 802
            }
        }
    }
    var v20 uint64 = 2
    v20 = (v20 >> ((4294967295) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
