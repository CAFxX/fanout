package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
func genFn0(p0 uint64, p1 uint8) int16 {
    p0 += (p0 / (((p0 % ((p0) | 1))) | 1))
    v7 := GenS1{F0: 947}
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = v7
    return -32768
}
func genFn1(p0 uint8) uint64 {
    var v8 uint8 = (p0 >> ((2) & 63))
    v9 := GenS1{F0: 406}
    _ = p0
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int32 = 42
    var v2 int32 = v1
    v3 := make(map[uint32]bool)
    var v4 int = 128
    var v5 string = ""
    var v6 string = "\n\t"
    v4 = (v4 * v4)
    if true {
        v10 := GenS2{F0: 9223372036854775807}
        v1 = (v2 - v1)
        v4 += 127
        _ = v10
    } else {
        v5 = (("foo bar" + "foo bar") + v6)
        var v11 int16 = 32767
        _ = v11
    }
    v2 -= v2
    v5 = "foo bar"
    var v12 uint64 = 4294967295
    v12 += 809
    var v13 float32 = (((1e10 - 0.1) * float32(-2.25)) - (0.0 - (0.1 * 3.14159)))
    v13 += float32(float64(float32((0.1 + 1.5))))
    var v14 uint16 = uint16(v12)
    var v15 int32 = v2
    v6 = v5
    v16 := GenS3{F0: 1000, F1: (v12 >> ((1) & 63)), F2: (v12 >> ((42) & 63))}
    for fi17 := 0; fi17 < 15; fi17++ {
        _ = len(v3)
        _ = fi17
    }
    var v18 uint16 = v14
    v19 := []int{(-(v4)), v4, (v4 % ((0) | 1))}
    fmt.Println(v1, v2, v14)
    if (255 <= 42) {
        if false {
            var v20 uint16 = uint16(v12)
            if false {
                if true {
                    var v21 string = ("" + ("\n\t" + "\n\t"))
                    fmt.Println(v6, v14)
                    _ = v21
                } else {
                    v20 = 65535
                }
            }
            _ = v20
        }
    } else {
        v13 = float32((0.0 / (1.5 + 0.0)))
    }
    v18 = 697
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
