package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int; F2 int16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint16, p2 string) uint32 {
    var v4 uint32 = uint32(p1)
    var v5 int16 = -32768
    fmt.Println(p2, p1)
    if (p2 >= p2) {
        v6 := []int{-254, 243, 9223372036854775807, 847}
        v5 = (v5 + v5)
        _ = v6
    }
    var v7 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v7
    return v4
}
func genFn1(p0 uint, p1 uint8, p2 int16) int16 {
    var v8 uint8 = p1
    if false {
        var v9 int32 = -86
        if ((p2 & p2) <= 32767) {
            var v10 int16 = p2
            _ = v10
        } else {
            p2 = p2
        }
        _ = v9
    } else {
        v11 := []int{}
        p0 += (p0 << ((p0) & 63))
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p2 + 1)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 3.14159
    var v3 uint16 = 0
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(fi12)
        v13 := []int{fi12}
        _ = len(v1)
        _ = fi12
        _ = v13
    }
    switch "hello" {
    case "":
        v14 := GenS1{F0: -9223372036854775808}
        _ = v14
    }
    v2 = 3.14159
    v15 := v1[uint32(v3)]
    var v16 float32 = (3.14159 - (float32(((1e10 * (0.0 + float64(1.5))) + ((((0.0 + 3.14159) - float64((v2 + (((v2 * v2) / (float32((3.14159 * (1.5 + 3.14159))) + float32(3.14159))) - ((-2.25 + ((v2 + 0.1) - -2.25)) * v2))))) * 1.5) + 1.5))) * v2))
    v17 := []int{127, 1}
    fmt.Println(v16, v15, v2)
    _ = len(v1)
    switch ("" + "\n\t") {
    case "foo bar":
        v18 := GenS2{F0: 42, F1: 42, F2: v3}
        _ = v18
    case "":
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS3{F0: 255, F1: 946, F2: 0}
            v16 = ((v2 - (v2 + float32(((3.14159 / 1e10) / 1.0)))) - 1.5)
            _ = fi19
            _ = v20
        }
    case "\n\t":
        fmt.Println(v16, v16, v15)
    }
    v21 := GenS1{F0: 42}
    switch 807 {
    case -320:
        v15 = (1e10 != 1.5)
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 18; fi23++ {
            fmt.Println(v15)
            _ = fi23
        }
        v24 := GenS4{F0: -9223372036854775808}
        v2 = (v2 - 3.14159)
        _ = fi22
        _ = v24
    }
    var v25 int16 = 769
    delete(v1, 154)
    v26 := GenS3{F0: 27, F1: 4294967295, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
}
