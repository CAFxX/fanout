package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint16 }
func genFn0() int32 {
    v6 := []int{-989, 384, 32768, 128}
    var v7 int8 = -1
    var v8 uint8 = 42
    _ = v6
    _ = v7
    _ = v8
    return 884
}
func genFn1(p0 bool, p1 string) int32 {
    v9 := []int{-9223372036854775808, 2147483647, 128, 654}
    var v10 uint = 255
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    v3 := []string{("hello" + ("hello" + "")), "hello", "", ("hello" + "\n\t")}
    v4 := []string{("foo bar" + ""), "\n\t", "hello", "a"}
    var v5 int16 = -32768
    v11 := GenS1{F0: 255}
    var v12 int64 = 42
    switch ("foo bar" + "a") {
    case "\n\t":
        fmt.Println(v12, v12)
    default:
        v13 := v2[369]
        _ = v13
    }
    var v14 int32 = 2147483647
    switch 42 {
    case -185:
        var v15 int = -9223372036854775808
        _ = v15
    default:
        if (3.14159 <= -2.25) {
            delete(v2, 748)
            fmt.Println(v12, v5)
            if ((0.1 + float32(1e10)) >= float32((((3.14159 - (1e10 + float64(1e10))) - -2.25) + -2.25))) {
                fmt.Println(v5, v14)
                delete(v2, 42)
            }
        } else {
            switch 9223372036854775807 {
            case -1:
                var v16 int8 = 0
                _ = v16
            case 1:
                for fi17 := 0; fi17 < 20; fi17++ {
                    if (!false) {
                        v18 := v2[42]
                        _ = v18
                    }
                    _ = fi17
                }
            default:
                fmt.Println(v5, v5)
            }
            v1 = append(v1, 4294967295)
            v19 := GenS2{F0: 551, F1: 42, F2: 227}
            _ = v19
        }
    }
    v12 = (v12 | v12)
    var v20 bool = true
    fmt.Println(v5, v12, v14)
    if false {
        v21 := GenS2{F0: 18446744073709551615, F1: 4294967295, F2: -12}
        _ = v21
    } else {
        switch 128 {
        case 107:
            var v22 uint32 = 2147483647
            _ = v22
        case 9223372036854775807:
            for fi23 := 0; fi23 < 4; fi23++ {
                v14 += genFn1((!(42 > 0)), ("foo bar" + ("foo bar" + "\n\t")))
                var v24 uint32 = 42
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                _ = fi23
                _ = v24
            }
        }
        v12 = int64(v14)
    }
    var v25 int64 = (v12 | v12)
    if v20 {
        fmt.Println(v14)
    } else {
        var v26 float32 = (float32((((float64(((0.1 * 0.1) + (1.5 + -2.25))) / 3.14159) / -2.25) - float64(float32(1.5)))) + (0.0 * 3.14159))
        v20 = v20
        v27 := GenS3{F0: 1, F1: 0, F2: 32767}
        _ = v26
        _ = v27
    }
    if v20 {
        v12 += (v25 + (v12 & 42))
    }
    var v28 string = "\n\t"
    switch -428 {
    case 272:
        v4 = append(v4, (v28 + ("" + (v28 + ("" + v28)))))
    }
    var v29 uint16 = 1
    v5 += (v5 / ((v5) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v20
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
