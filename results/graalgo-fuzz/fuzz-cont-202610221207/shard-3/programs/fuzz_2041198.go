package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 string = "foo bar"
        v6 := []int{fi4}
        fmt.Println(v5, v5)
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int = 42
    v8 := []int{(v7 / ((v7) | 1)), v7, (v7 % ((v7) | 1))}
    v9 := GenS1{F0: -1}
    _ = v7
    _ = v8
    _ = v9
    return 128
}
func genFn1(p0 uint8) uint16 {
    var v10 uint64 = 607
    p0 += 221
    if false {
        v11 := GenS2{F0: (^(p0))}
        _ = v11
    } else {
        fmt.Println(p0, p0)
        v12 := []int{}
        _ = v12
    }
    v10 += 0
    v10 -= (v10 << ((1) & 63))
    _ = p0
    _ = v10
    return 2
}

func main() {
    v1 := []uint32{0, 4294967295, 253}
    v2 := make(map[int]string)
    var v3 float64 = ((((1e10 / (1e10 / 1.5)) * float64(((0.1 - (float32(1.5) - 0.0)) * float32((0.0 - float64((-2.25 - 0.1))))))) - 0.0) + -2.25)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = ((v3 + (-2.25 / float64((-2.25 + 0.1)))) + (3.14159 * v3))
    switch -191 {
    case 128:
        var v13 float32 = 1.5
        _ = v13
    case 651:
        for fi14 := 0; fi14 < 6; fi14++ {
            _ = len(v1)
            v3 -= float64(1e10)
            _ = fi14
        }
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(fi15, fi15)
        _ = fi15
    }
    if (32767 < 1000) {
        var v16 int8 = -128
        fmt.Println(v3)
        _ = v16
    }
    v17 := GenS3{F0: -1}
    for fi18 := 0; fi18 < 15; fi18++ {
        v3 += float64(1.5)
        _ = fi18
    }
    switch ("foo bar" + "hello") {
    case "foo bar":
        delete(v2, 0)
    case "a":
        var v19 uint8 = 42
        _ = v19
    case "":
        switch "hello" {
        case "a":
            switch 9223372036854775807 {
            case -2147483648:
                var v20 uint8 = 177
                _ = v20
            case -9223372036854775808:
                var v21 int64 = -886
                _ = v21
            }
        case "foo bar":
            for fi22 := 0; fi22 < 16; fi22++ {
                v23 := GenS1{F0: 0}
                var v24 int32 = -839
                _ = fi22
                _ = v23
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 11; fi25++ {
                v3 -= v3
                _ = fi25
            }
        }
    default:
        v3 = -2.25
    }
    v3 = 1.5
    var v26 uint = 42
    v27 := GenS3{F0: -1}
    fmt.Println(v26, v26)
    v3 = 1.5
    v28 := GenS3{F0: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
