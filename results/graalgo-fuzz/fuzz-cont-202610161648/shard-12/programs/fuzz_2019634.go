package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 int16, p1 uint16) int32 {
    var v7 int16 = -268
    v8 := GenS1{F0: p1, F1: -75}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int) uint64 {
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := []int{}
        fmt.Println(p0, fi9)
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 1
    v12 := GenS1{F0: 941, F1: 100}
    for fi13 := 0; fi13 < 19; fi13++ {
        v11 += v11
        if true {
            var v14 string = "\n\t"
            _ = v14
        }
        _ = fi13
    }
    var v15 string = "foo bar"
    _ = p0
    _ = v11
    _ = v12
    _ = v15
    return 577
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{133, 42, 208, 2147483647}
    var v3 int16 = 42
    var v4 float64 = (0.1 + 0.1)
    v5 := []uint32{0, 4294967295}
    var v6 uint64 = 18446744073709551615
    v4 += v4
    v6 = v6
    var v16 int16 = (v3 & -1)
    _ = len(v2)
    var v17 int8 = -128
    if false {
        v4 += ((((float64(1e10) - v4) + v4) + (v4 * (float64(float32(float64(0.1))) / 1.5))) - 0.0)
    }
    v18 := GenS1{F0: 42, F1: v17}
    var v19 uint8 = 127
    v20 := GenS1{F0: 1, F1: (^(v17))}
    var v21 string = "foo bar"
    v22 := v1[127]
    var v23 int32 = -2147483648
    switch v21 {
    case "hello":
        var v24 bool = (1 >= 741)
        _ = v24
    case "foo bar":
        var v25 uint32 = (v22 | 147)
        _ = v25
    default:
        for fi26 := 0; fi26 < 1; fi26++ {
            fmt.Println(v22)
            if (true && (!(!(!true)))) {
                v27 := GenS1{F0: 379, F1: (v17 * int8(v3))}
                v4 = v4
                _ = v27
            }
            for fi28 := 0; fi28 < 9; fi28++ {
                v23 = genFn0((v3 * -998), 2)
                if ((v19 / (((v19 * v19)) | 1)) == v19) {
                    v29 := GenS2{F0: 42, F1: v23}
                    v22 = (v22 - v22)
                    _ = v29
                } else {
                    v30 := GenS1{F0: 42, F1: (v17 + int8(v6))}
                    _ = v30
                }
                _ = fi28
            }
            _ = fi26
        }
    }
    if len(v2) > 0 {
        v2[0] = (v22 >> ((81) & 63))
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS1{F0: 32767, F1: v17}
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        if false {
            v3 -= 1
            var v34 bool = true
            _ = v34
        }
        _ = fi33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
