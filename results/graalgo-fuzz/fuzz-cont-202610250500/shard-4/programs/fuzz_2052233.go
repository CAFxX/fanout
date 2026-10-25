package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint; F2 bool }
func genFn0(p0 int16, p1 bool, p2 uint8) int64 {
    v5 := []int{0}
    v6 := GenS1{F0: 65535, F1: p0, F2: 2147483647}
    p0 += 0
    if len(v5) > 0 {
        v5[0] = 32767
    }
    switch ((("" + "a") + "a") + ("" + "\n\t")) {
    case "\n\t":
        p0 -= (p0 >> ((722) & 63))
    case "a":
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                fmt.Println(p1, p2, p2)
                _ = fi8
            }
            v9 := GenS2{F0: (^(fi7))}
            if (!p1) {
                switch fi7 {
                case 9223372036854775807:
                    var v10 uint8 = 128
                    _ = v10
                }
                v11 := GenS1{F0: 4294967295, F1: p0, F2: 775}
                _ = v11
            } else {
                fmt.Println(fi7, p2, fi7)
                p2 = 255
            }
            _ = fi7
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 uint64 = 876
    var v2 int32 = -941
    v3 := make(map[string]uint32)
    var v4 int32 = 464
    var v12 int = 32768
    var v13 float64 = 3.14159
    v14 := []int{}
    v12 = (v12 >> ((0) & 63))
    v15 := GenS2{F0: v12}
    var v16 float64 = 1.5
    for fi17 := 0; fi17 < 18; fi17++ {
        fmt.Println(v1)
        fmt.Println(v16, fi17)
        v18 := GenS3{F0: true, F1: 65535, F2: false}
        _ = fi17
        _ = v18
    }
    var v19 int = (v12 / (((v12 - (v12 / ((v12) | 1)))) | 1))
    fmt.Println(v16)
    var v20 string = "\n\t"
    var v21 float64 = float64((-2.25 * 0.0))
    fmt.Println(v16, v13, v2)
    fmt.Println(v1, v16)
    v21 -= v21
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
