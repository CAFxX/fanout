package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint64 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int32 = 42
        fmt.Println(fi5)
        _ = fi5
        _ = v6
    }
    if false {
        var v7 int16 = -1
        _ = v7
    } else {
        var v8 uint32 = 223
        fmt.Println(v8)
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{fi9, (fi9 % ((fi9) | 1)), (fi9 ^ 127), fi9}
        if len(v10) > 0 {
            v10[0] = 2147483647
        }
        v11 := GenS1{F0: 0, F1: 2, F2: 4294967295}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint = 2
    _ = v12
    return -446
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 uint8 = 255
    v13 := []int{-306, -1, 128, -977}
    var v14 uint = 255
    var v15 int32 = 1
    v15 = (v15 & v15)
    v16 := GenS1{F0: 32767, F1: 1, F2: v14}
    fmt.Println(v4, v15, v4)
    if true {
        v2[0] = 0
        v1 = int8(v15)
        _ = len(v13)
    } else {
        switch "hello" {
        case "\n\t":
            v1 -= (v1 % ((-128) | 1))
        case "":
            switch ((("a" + (((("" + "a") + "foo bar") + (("\n\t" + "") + "hello")) + "hello")) + "\n\t") + "foo bar") {
            case "hello":
                if ((genFn0() <= (^(v15))) && true) {
                    var v17 int8 = (^(v1))
                    v18 := GenS1{F0: 2, F1: (v4 | v4), F2: (v14 + v14)}
                    if (("a" > "a") || ((1e10 - ((((1e10 + float32(float64(3.14159))) - ((3.14159 - 0.1) + (float32(((1.5 * 1e10) / 1.0)) - (1e10 - 1.5)))) - (((3.14159 / (1.5 + ((0.1 + 1e10) + (-2.25 + 0.0)))) - -2.25) - 3.14159)) + (0.1 - float32(((((1.5 * 1.5) / -2.25) / 0.1) * -2.25))))) == 1e10)) {
                        v1 = v17
                        v19 := GenS1{F0: 255, F1: (v4 >> ((255) & 63)), F2: v14}
                        switch 127 {
                        case -529:
                            var v20 float32 = float32(1.5)
                            _ = v20
                        default:
                            v15 = v15
                        }
                        _ = v19
                    } else {
                        var v21 string = "hello"
                        v4 = v4
                        _ = v21
                    }
                    _ = v17
                    _ = v18
                } else {
                    v22 := GenS2{F0: v4}
                    var v23 string = "\n\t"
                    if (!true) {
                        var v24 uint16 = 42
                        _ = v24
                    }
                    _ = v22
                    _ = v23
                }
            }
        }
        var v25 int32 = v15
        v26 := GenS3{F0: -1, F1: 42}
        _ = v25
        _ = v26
    }
    var v27 uint64 = 42
    _ = len(v13)
    var v28 uint16 = 600
    if false {
        v27 += v27
        fmt.Println(v27, v27, v1)
        var v29 uint64 = v27
        _ = v29
    }
    switch 127 {
    case 2147483647:
        v27 = uint64(v28)
    case -2147483648:
        var v30 int16 = int16(v15)
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
