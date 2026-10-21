package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }
func genFn0(p0 int16) uint {
    var v4 int = 42
    v5 := GenS1{F0: (0 == 12), F1: 42, F2: 2}
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 string, p1 bool, p2 int16) string {
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 18; fi7++ {
        if (!p1) {
            v8 := []int{738, 2147483647, (fi7 ^ fi7)}
            var v9 uint64 = 2
            var v10 uint8 = 2
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 float32 = -2.25
            v11 += ((v11 / 1.0) / (0.1 * 0.1))
            if (p2 <= p2) {
                v12 := []int{32768, (fi7 >> ((736) & 63)), -612, fi7}
                _ = v12
            }
            _ = v11
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ("\n\t" + p0)
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 202
    v3 := make(map[uint32]bool)
    v2 -= v2
    fmt.Println(v2)
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := []int{fi13, fi13, 1}
        _ = fi13
        _ = v14
    }
    fmt.Println(v2)
    var v15 int64 = 9223372036854775807
    var v16 uint64 = 0
    for fi17 := 0; fi17 < 16; fi17++ {
        v16 += 1
        _ = fi17
    }
    var v18 uint64 = (v16 / (((v16 ^ v16)) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
