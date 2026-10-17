package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int16; F1 int64 }

func main() {
    var v1 uint32 = 2
    var v2 float32 = (1e10 - (0.0 - float32((1e10 - (float64(1.5) / -2.25)))))
    var v3 uint32 = v1
    var v4 uint64 = 2
    var v5 bool = true
    var v6 int64 = -249
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int16 = 0
        v5 = (true || (v2 >= v2))
        _ = fi7
        _ = v8
    }
    var v9 uint32 = v1
    v4 -= 273
    if (318 >= 255) {
        switch ("hello" + ("\n\t" + "hello")) {
        case "a":
            v6 += (v6 ^ 42)
        case "foo bar":
            for fi10 := 0; fi10 < 10; fi10++ {
                v11 := []int{fi10, (fi10 + fi10)}
                _ = fi10
                _ = v11
            }
        default:
            v12 := GenS1{F0: 4294967295, F1: 1, F2: ""}
            _ = v12
        }
        if false {
            fmt.Println(v9, v3, v1)
            var v13 int64 = -1
            _ = v13
        } else {
            v14 := GenS1{F0: 128, F1: 640, F2: (("hello" + "") + "\n\t")}
            _ = v14
        }
        switch "a" {
        case "":
            v6 -= (v6 << ((925) & 63))
        case "hello":
            v2 -= v2
        default:
            var v15 uint32 = v1
            _ = v15
        }
    } else {
        var v16 uint8 = 73
        var v17 float64 = 3.14159
        switch "\n\t" {
        case "\n\t":
            v6 = (^(v6))
        case "foo bar":
            v18 := []int{-9223372036854775808}
            _ = v18
        case "a":
            v16 += (v16 & 111)
        default:
            v3 -= (^(v3))
        }
        _ = v16
        _ = v17
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        if ((0 > 0) && v5) {
            v20 := []int{(fi19 / ((1) | 1)), fi19, fi19, 2147483647}
            _ = v20
        } else {
            v1 = (^(v9))
            var v21 float32 = 3.14159
            v22 := GenS1{F0: fi19, F1: uint(v4), F2: "hello"}
            _ = v21
            _ = v22
        }
        v23 := GenS2{F0: 1000, F1: 849}
        _ = fi19
        _ = v23
    }
    v5 = v5
    switch 32767 {
    case -631:
        switch "\n\t" {
        case "a":
            var v24 bool = (!v5)
            _ = v24
        }
    case -9223372036854775808:
        var v25 float64 = (-2.25 * 1e10)
        _ = v25
    case 127:
        var v26 uint32 = (v9 & v9)
        _ = v26
    default:
        v6 = (v6 >> ((0) & 63))
    }
    fmt.Println(v9)
    for fi27 := 0; fi27 < 4; fi27++ {
        fmt.Println(v1, v4)
        _ = fi27
    }
    v28 := []int{9223372036854775807, 604}
    v29 := GenS3{F0: 0, F1: v6}
    v28 = append(v28, 128)
    _ = len(v28)
    for fi30 := 0; fi30 < 6; fi30++ {
        _ = len(v28)
        var v31 int8 = -56
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
