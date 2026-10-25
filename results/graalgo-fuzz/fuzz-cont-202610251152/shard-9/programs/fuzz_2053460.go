package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int16; F1 uint; F2 uint64 }
func genFn0(p0 uint16) int {
    var v5 int64 = 9223372036854775807
    v5 -= 42
    p0 = p0
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := []int{1, 127, -332, 9223372036854775807}
    var v2 uint16 = 707
    var v3 int = 0
    v4 := []string{"\n\t"}
    if (true || (-9223372036854775808 < -9223372036854775808)) {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := GenS1{F0: 42, F1: 255}
            _ = fi6
            _ = v7
        }
    } else {
        v2 -= v2
        v2 += (v2 & (v2 >> ((2) & 63)))
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    v3 += v3
    v8 := GenS2{F0: 0, F1: -1}
    var v9 int64 = 0
    var v10 int64 = 663
    var v11 int = genFn0((v2 >> ((493) & 63)))
    var v12 uint8 = 127
    var v13 float32 = (-2.25 - -2.25)
    v1 = append(v1, v11)
    v14 := GenS3{F0: 539, F1: uint(v2), F2: 393}
    fmt.Println(v3, v11, v9)
    v9 += -9223372036854775808
    var v15 float64 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
