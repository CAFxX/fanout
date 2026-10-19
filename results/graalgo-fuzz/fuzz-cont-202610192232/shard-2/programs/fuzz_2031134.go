package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint; F2 int64 }
func genFn0(p0 int8, p1 int16, p2 uint64) int8 {
    var v6 int8 = p0
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (v6 * 0)
}

func main() {
    var v1 bool = (-9223372036854775808 > 42)
    var v2 int8 = 100
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    v5 := make(map[string]uint32)
    v7 := GenS1{F0: 4294967295, F1: 210}
    v1 = v1
    v1 = v1
    v1 = (-483 == -1)
    fmt.Println(v2, v1)
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS2{F0: -882}
        _ = fi8
        _ = v9
    }
    v10 := []int{1}
    v11 := GenS3{F0: (v2 % ((v2) | 1)), F1: 1, F2: -695}
    fmt.Println(v2)
    v12 := GenS3{F0: -1, F1: 718, F2: 882}
    v13 := v5[((((("" + "a") + "a") + "a") + (("a" + ("" + "hello")) + "a")) + "\n\t")]
    var v14 float32 = ((float32(float64(3.14159)) - (3.14159 * float32(float64((0.0 * 1e10))))) / 1e10)
    var v15 uint16 = 1
    switch 128 {
    case 2147483647:
        var v16 int64 = -1
        _ = v16
    default:
        fmt.Println(v2, v2)
    }
    v17 := v3["\n\t"]
    var v18 string = (((v17 + ("\n\t" + (((v17 + "") + "hello") + ((v17 + v17) + (v17 + (v17 + v17)))))) + ((((v17 + (v17 + v17)) + v17) + (v17 + v17)) + v17)) + v17)
    var v19 uint = 31
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
