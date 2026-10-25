package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 int8 }
type GenS3 struct { F0 int16; F1 string; F2 int64 }
type GenS4 struct { F0 int; F1 int16 }
func genFn0(p0 uint8, p1 uint) uint8 {
    switch 32767 {
    case 112:
        p0 = uint8(p1)
    }
    var v4 int64 = -1
    v4 += (v4 >> (((p1 % ((p1) | 1))) & 63))
    fmt.Println(p1, p0, p0)
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p0 / ((2) | 1))
}
func genFn1(p0 uint32) string {
    if (2 < 776) {
        if (1 > 65535) {
            p0 = 428
        } else {
            var v6 float32 = float32((0.0 + 0.1))
            _ = v6
        }
    } else {
        v7 := GenS1{F0: 386}
        v8 := GenS1{F0: 956}
        _ = v7
        _ = v8
    }
    v9 := []int{2147483647, 4294967295, -899}
    for fi10 := 0; fi10 < 11; fi10++ {
        p0 = 4294967295
        v11 := GenS2{F0: 2, F1: 42, F2: 100}
        p0 -= p0
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v9
    return ""
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = -2147483648
    v3 := make(map[string]bool)
    var v12 int16 = 32767
    var v13 uint64 = 4294967295
    var v14 uint8 = 39
    if (((-2.25 + (1.5 + 0.0)) - 0.1) == float32((-2.25 + float64((0.1 / ((0.1 / 0.1) / 3.14159)))))) {
        switch "hello" {
        case "hello":
            fmt.Println(v2, v12, v2)
        }
        var v15 float32 = (float32(1e10) - float32((3.14159 / -2.25)))
        _ = v15
    }
    v13 -= (v13 - 58)
    v2 -= v2
    var v16 int16 = v12
    fmt.Println(v13, v16, v2)
    for fi17 := 0; fi17 < 10; fi17++ {
        fmt.Println(v2, v13)
        var v18 uint16 = 36
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        v20 := GenS3{F0: 42, F1: "hello", F2: 42}
        fmt.Println(v16)
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(v2)
        _ = fi21
    }
    v22 := GenS3{F0: 37, F1: "", F2: 42}
    v2 += v2
    v23 := GenS4{F0: v2, F1: (^(v12))}
    var v24 string = ((((("\n\t" + "a") + "\n\t") + "\n\t") + "") + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
