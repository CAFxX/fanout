package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0(p0 int) string {
    fmt.Println(p0)
    v4 := []int{}
    _ = p0
    _ = v4
    return "\n\t"
}
func genFn1(p0 string, p1 int32, p2 int32) int16 {
    var v5 uint16 = 0
    if (1000 < 0) {
        var v6 float32 = 0.0
        var v7 uint16 = v5
        var v8 int = -655
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -1}
    if false {
        if true {
            v10 := []int{103, 1, -9223372036854775808, 618}
            if true {
                var v11 float32 = (1e10 * (3.14159 / (0.1 + 1.5)))
                p2 -= (p2 << ((1) & 63))
                var v12 string = p0
                _ = v11
                _ = v12
            }
            p2 += (p2 << ((573) & 63))
            _ = v10
        }
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            p0 = p0
            var v14 uint8 = 255
            _ = fi13
            _ = v14
        }
        switch -9223372036854775808 {
        case -2147483648:
            v15 := GenS1{F0: 367}
            _ = v15
        default:
            v16 := []int{-145, 42, 127}
            _ = v16
        }
        p1 = (p1 % ((0) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    v3 := make(map[uint32]string)
    if (523 != 42) {
        if (!true) {
            var v17 uint = 945
            _ = v17
        }
        v2 += 0
    } else {
        v2 += (v2 | 42)
        v18 := v1["foo bar"]
        _ = v18
    }
    if false {
        v19 := GenS2{F0: 23}
        _ = v19
    } else {
        fmt.Println(v2)
        v20 := []int{}
        v21 := GenS3{F0: 48}
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        switch fi22 {
        case 9223372036854775807:
            fmt.Println(v2)
        case 127:
            v2 += -1
        }
        v2 += (^(v2))
        v23 := GenS4{F0: v2, F1: v2}
        _ = fi22
        _ = v23
    }
    var v24 int = 32768
    v25 := []int{v24, -2147483648, v24}
    v25 = append(v25, v24)
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := GenS5{F0: genFn1(("hello" + "a"), v2, (v2 ^ (-(v2)))), F1: -9223372036854775808}
        _ = fi26
        _ = v27
    }
    v2 = v2
    switch v24 {
    case 93:
        v28 := GenS3{F0: 32767}
        _ = v28
    default:
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 float64 = -2.25
            _ = fi29
            _ = v30
        }
    }
    switch "" {
    case "hello":
        for fi31 := 0; fi31 < 3; fi31++ {
            fmt.Println(v24, v2, v2)
            v32 := GenS1{F0: (fi31 >> ((286) & 63))}
            _ = fi31
            _ = v32
        }
    case "foo bar":
        v2 -= (v2 >> ((4294967295) & 63))
    default:
        var v33 uint32 = 2147483647
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
