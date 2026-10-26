package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    v5 := []int{612, 32767, 121, 9223372036854775807}
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 string = "a"
    v5 = append(v5, 520)
    var v8 int32 = 1
    v8 += 1
    if ((!(true && false)) || false) {
        v8 = (v8 - 2147483647)
        if (18446744073709551615 < 18446744073709551615) {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 string = v7
                _ = fi9
                _ = v10
            }
        }
        var v11 uint8 = 1
        _ = v11
    } else {
        var v12 int64 = -1
        _ = v12
    }
    switch -243 {
    case 9223372036854775807:
        if true {
            if (-1 != -9223372036854775808) {
                for fi13 := 0; fi13 < 15; fi13++ {
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v7 = ((("hello" + "\n\t") + "a") + v7)
                        _ = fi14
                    }
                    _ = fi13
                }
                v7 = (v7 + (("a" + "foo bar") + v7))
            }
        } else {
            v8 += v8
        }
    case 0:
        if ((42 >= 18446744073709551615) || false) {
            switch ((v7 + "hello") + "foo bar") {
            case "":
                if ((true && (-214 <= v8)) && ((-186 != -1) || false)) {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        fmt.Println(v8, v8)
                        fmt.Println(v7)
                        _ = fi15
                    }
                }
            case "foo bar":
                v7 = v7
            case "hello":
                switch 4294967295 {
                case 42:
                    v16 := GenS1{F0: 1000}
                    _ = v16
                case 568:
                    v7 = v7
                case 127:
                    fmt.Println(v7)
                default:
                    if true {
                        v17 := GenS1{F0: 159}
                        _ = v17
                    } else {
                        v18 := GenS2{F0: 422, F1: -4}
                        var v19 float32 = (float32(((1.5 * 1.5) - 0.0)) + float32(float64(float32((0.1 + 0.1)))))
                        _ = len(v5)
                        _ = v18
                        _ = v19
                    }
                }
            }
        }
    case 794:
        var v20 uint64 = 42
        _ = v20
    default:
        var v21 float32 = 3.14159
        _ = v21
    }
    v22 := GenS3{F0: 42}
    v23 := GenS3{F0: -828}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
