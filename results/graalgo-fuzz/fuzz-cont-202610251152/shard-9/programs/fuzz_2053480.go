package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 int32; F2 uint }
type GenS5 struct { F0 uint }
func genFn0() uint {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 998, F1: 65535}
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := []int{349, fi9, fi9, fi7}
            switch -587 {
            case -919:
                v11 := GenS1{F0: 4294967295, F1: 0}
                _ = v11
            case 127:
                var v12 int16 = int16(fi9)
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        var v13 int8 = 121
        _ = fi7
        _ = v13
        _ = v8
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := []int{(fi14 ^ fi14), (fi14 ^ (fi14 % ((4294967295) | 1))), (fi14 | fi14)}
        v16 := GenS2{F0: 9223372036854775807, F1: 291}
        _ = fi14
        _ = v15
        _ = v16
    }
    return 745
}
func genFn1(p0 int64) int8 {
    v17 := []int{32768, 42, -587, 0}
    var v18 uint64 = 1
    _ = p0
    _ = v17
    _ = v18
    return 127
}

func main() {
    var v1 uint32 = 383
    var v2 float32 = 1e10
    var v3 int8 = -1
    var v4 int64 = 1
    var v5 string = "\n\t"
    var v6 uint = 65535
    for fi19 := 0; fi19 < 10; fi19++ {
        v3 += (v3 << (((^(v6))) & 63))
        _ = fi19
    }
    v20 := []int{-750, -625, 224}
    v20 = append(v20, 32768)
    var v21 int32 = -1
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 string = "foo bar"
        v24 := GenS3{F0: (v6 % (((^(v6))) | 1))}
        var v25 int8 = (v3 - (v3 * 0))
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 int8 = 82
    if false {
        var v27 uint8 = 72
        var v28 int = 9223372036854775807
        _ = v27
        _ = v28
    }
    if true {
        v2 = v2
        var v29 uint16 = 935
        var v30 float64 = float64(v2)
        _ = v29
        _ = v30
    } else {
        v1 = (v1 >> ((v6) & 63))
    }
    v20 = append(v20, 32767)
    v1 += (v1 / ((0) | 1))
    fmt.Println(v6)
    if false {
        v21 += (v21 * (v21 >> ((v6) & 63)))
        v31 := GenS4{F0: v21, F1: v21, F2: (v6 >> (((^(v6))) & 63))}
        v21 = v21
        _ = v31
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    if (127 == v26) {
        v5 = v5
    } else {
        switch 127 {
        case 4294967295:
            var v32 float32 = (0.1 / v2)
            _ = v32
        }
        for fi33 := 0; fi33 < 14; fi33++ {
            v26 -= (-(v26))
            v34 := GenS1{F0: 42, F1: v1}
            _ = fi33
            _ = v34
        }
    }
    var v35 int = 4294967295
    v36 := GenS5{F0: 0}
    var v37 int = v35
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
