package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 string }
func genFn0() uint16 {
    if (0 != -959) {
        var v6 uint32 = 1
        _ = v6
    } else {
        var v7 uint = 312
        v7 = (v7 | 892)
        _ = v7
    }
    var v8 int = -559
    v9 := []int{v8, (^(v8)), v8}
    var v10 bool = ((v8 + 4294967295) < (v8 % ((v8) | 1)))
    fmt.Println(v10, v8)
    _ = v10
    _ = v8
    _ = v9
    return 263
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]string)
    var v3 uint64 = 1
    v4 := make(map[string]int)
    var v5 int16 = 290
    if true {
        var v11 uint8 = 186
        for fi12 := 0; fi12 < 19; fi12++ {
            if false {
                v13 := GenS1{F0: int32(v3), F1: v5}
                v11 += (^(v11))
                for fi14 := 0; fi14 < 4; fi14++ {
                    v15 := []int{(fi12 >> ((2) & 63)), (fi14 | (-(fi14))), (fi12 & (fi14 << ((18446744073709551615) & 63))), fi12}
                    fmt.Println(fi14)
                    if len(v15) > 0 {
                        v15[0] = 32768
                    }
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            _ = fi12
        }
        var v16 int16 = (v5 / (((v5 + v5)) | 1))
        _ = v11
        _ = v16
    } else {
        var v17 string = ("\n\t" + ("foo bar" + "a"))
        v18 := v4[(("foo bar" + v17) + ("\n\t" + (("a" + v17) + ("\n\t" + (v17 + v17)))))]
        _ = v17
        _ = v18
    }
    if true {
        if true {
            var v19 int = 281
            var v20 int8 = 0
            _ = v19
            _ = v20
        }
        delete(v2, 127)
    } else {
        v5 += 1
        if false {
            var v21 uint16 = 255
            _ = v21
        }
    }
    v1 -= float32((3.14159 + -2.25))
    v22 := GenS1{F0: 0, F1: (v5 & 0)}
    v1 += (float32(1.5) + (1.5 * v1))
    v1 = ((v1 - v1) + 0.1)
    fmt.Println(v5, v1)
    v5 += v5
    for fi23 := 0; fi23 < 3; fi23++ {
        v24 := v4[("a" + "\n\t")]
        fmt.Println(v24, v3)
        fmt.Println(v1, v1)
        _ = fi23
        _ = v24
    }
    var v25 uint32 = 615
    var v26 int64 = -304
    v27 := []int{128, -137, 42, 739}
    var v28 uint = 449
    v29 := GenS2{F0: v28, F1: (v5 / ((v5) | 1))}
    var v30 int = 711
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    if false {
        v5 += 5
    } else {
        fmt.Println(v25, v3, v28)
        v31 := GenS3{F0: -21, F1: v28, F2: "foo bar"}
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
