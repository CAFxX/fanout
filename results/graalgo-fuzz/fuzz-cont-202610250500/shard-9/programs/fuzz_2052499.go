package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 bool; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int }
type GenS4 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint; F1 uint }
func genFn0() uint64 {
    v7 := GenS1{F0: -2147483648}
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS2{F0: 36, F1: false, F2: 0}
        v10 := []int{}
        fmt.Println(fi8)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v7
    return 2
}
func genFn1(p0 bool, p1 int16) int {
    if false {
        var v11 uint8 = 2
        var v12 uint16 = 21
        _ = v11
        _ = v12
    } else {
        var v13 int8 = 127
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        p1 = p1
        _ = fi14
    }
    var v15 int32 = 1
    switch -1 {
    case 569:
        var v16 uint64 = 18446744073709551615
        _ = v16
    case -1:
        if false {
            var v17 int16 = (p1 << ((1) & 63))
            var v18 float64 = (-2.25 * -2.25)
            _ = v17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = v15
    return 353
}

func main() {
    var v1 uint64 = 42
    var v2 float32 = (0.1 + (1e10 + float32(float64((1e10 + (((0.0 + (1.5 / float32((float64(3.14159) * 0.1)))) - (3.14159 * (((0.1 * 1e10) * (1e10 * -2.25)) - 3.14159))) * -2.25))))))
    var v3 uint = 825
    var v4 int = -751
    v5 := make(map[uint32]bool)
    v6 := make(map[string]string)
    v19 := GenS3{F0: 255, F1: 808, F2: 42}
    v2 = (((float32(float64(0.1)) * 1e10) + v2) + ((v2 * (v2 * 3.14159)) * v2))
    v20 := GenS2{F0: (v1 % ((uint64(v3)) | 1)), F1: false, F2: -9223372036854775808}
    switch v4 {
    case -148:
        v21 := []int{genFn1(false, -281)}
        _ = v21
    case 2147483647:
        for fi22 := 0; fi22 < 15; fi22++ {
            var v23 string = ("foo bar" + ("" + "hello"))
            v24 := GenS4{F0: 0, F1: 100, F2: 255}
            _ = fi22
            _ = v23
            _ = v24
        }
    case 9223372036854775807:
        switch ((("a" + "") + "a") + "hello") {
        case "hello":
            var v25 float64 = 0.1
            _ = v25
        default:
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 int64 = 1
                v2 = ((-2.25 + 1.5) * 1.5)
                fmt.Println(v4)
                _ = fi26
                _ = v27
            }
        }
    }
    _ = len(v6)
    var v28 uint8 = 0
    if (!(int16(v1) <= -100)) {
        fmt.Println(v1, v2)
        var v29 uint32 = 2147483648
        v30 := GenS2{F0: v1, F1: true, F2: 807}
        _ = v29
        _ = v30
    }
    v31 := []int{(v4 / ((genFn1((!(42 != 567)), 287)) | 1)), v4, -2147483648}
    fmt.Println(v2, v3, v28)
    if false {
        v1 = (v1 | (v1 + v1))
        if false {
            fmt.Println(v28, v3)
        }
        v1 = 950
    }
    v32 := GenS5{F0: (v3 % ((779) | 1)), F1: v3}
    var v33 uint64 = v1
    var v34 int = (v4 - int(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
