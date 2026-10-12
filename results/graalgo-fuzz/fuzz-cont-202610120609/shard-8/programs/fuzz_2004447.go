package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 bool; F1 uint; F2 int64 }
type GenS5 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint = 0
    var v2 bool = (-1 != -528)
    var v3 float64 = (0.0 * 1.5)
    var v4 uint64 = 2
    v1 += (v1 | 1)
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: 0}
        var v7 uint64 = (^(v4))
        v8 := GenS1{F0: -2147483648}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{32768}
    switch 2147483647 {
    case -1:
        var v10 uint64 = (^(v4))
        _ = v10
    case 4294967295:
        if (376 == 365) {
            fmt.Println(v1, v3, v1)
            fmt.Println(v2, v3)
        } else {
            fmt.Println(v4, v4)
        }
    }
    if len(v9) > 0 {
        v9[0] = 32767
    }
    v11 := GenS2{F0: v2, F1: 42, F2: 32767}
    v12 := GenS3{F0: 2147483647, F1: v1}
    var v13 int8 = -3
    fmt.Println(v3, v2, v2)
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS4{F0: true, F1: v1, F2: 711}
        if (((^(v13)) == v13) || true) {
            if (!true) {
                var v16 float32 = (1.5 / -2.25)
                _ = v16
            }
            v13 += (v13 >> ((v1) & 63))
            v17 := GenS5{F0: v1, F1: (v4 | 42)}
            _ = v17
        } else {
            var v18 int = 2147483647
            for fi19 := 0; fi19 < 7; fi19++ {
                if len(v9) > 0 {
                    v9[0] = fi14
                }
                _ = fi19
            }
            _ = v18
        }
        _ = len(v9)
        _ = fi14
        _ = v15
    }
    var v20 int32 = -743
    v21 := GenS1{F0: (v20 / (((v20 * 1)) | 1))}
    for fi22 := 0; fi22 < 15; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 uint32 = 2147483648
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v3 = v3
    var v25 int32 = v20
    switch "hello" {
    case "a":
        switch ("hello" + ("a" + "hello")) {
        case "":
            v2 = false
        }
    case "hello":
        if ((v3 + v3) != (float64(-2.25) + v3)) {
            v9 = append(v9, 42)
        } else {
            v26 := GenS4{F0: true, F1: v1, F2: -9223372036854775808}
            _ = v26
        }
    case "\n\t":
        fmt.Println(v1)
    default:
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 int16 = 97
            _ = fi27
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
