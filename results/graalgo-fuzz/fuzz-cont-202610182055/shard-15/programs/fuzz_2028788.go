package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint32 }
func genFn0(p0 int16, p1 int64, p2 uint64) uint {
    v6 := []int{896, 232}
    var v7 float32 = 1e10
    v8 := GenS1{F0: 1, F1: 4294967295, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1() uint64 {
    v9 := GenS1{F0: 70, F1: 448, F2: 666}
    var v10 uint64 = 42
    if (3.14159 < 3.14159) {
        fmt.Println(v10)
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v10)
            v12 := []int{127, (fi11 - 32768), (fi11 - (fi11 + fi11))}
            var v13 int16 = -237
            _ = fi11
            _ = v12
            _ = v13
        }
        fmt.Println(v10)
    } else {
        v14 := []int{-2147483648, -2147483648}
        _ = v14
    }
    _ = v10
    _ = v9
    return (v10 + 0)
}

func main() {
    v1 := []uint32{}
    var v2 float32 = (((-2.25 * 0.0) / 1.0) / -2.25)
    var v3 float32 = 1.5
    var v4 uint32 = 4294967295
    v5 := make(map[string]int)
    var v15 int64 = -97
    fmt.Println(v2, v2, v2)
    v4 -= v4
    if (!(0 > 0)) {
        if false {
            var v16 uint32 = v4
            v17 := GenS1{F0: 2, F1: v4, F2: (v16 - (v16 % (((v4 >> ((4294967295) & 63))) | 1)))}
            var v18 int8 = -101
            _ = v16
            _ = v17
            _ = v18
        }
        fmt.Println(v2, v3, v2)
    }
    v3 = float32(0.1)
    if false {
        v4 = 53
    }
    var v19 float32 = ((1e10 * float32((1.5 + 3.14159))) / 1.0)
    switch "\n\t" {
    case "":
        var v20 int32 = 1
        _ = v20
    case "hello":
        for fi21 := 0; fi21 < 6; fi21++ {
            fmt.Println(fi21, v15, v15)
            _ = fi21
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 42 {
    case 32768:
        var v22 uint = 2
        _ = v22
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 string = ((("" + "\n\t") + ("" + ("foo bar" + "\n\t"))) + ("a" + "a"))
        for fi25 := 0; fi25 < 3; fi25++ {
            if (42 > 42) {
                fmt.Println(v24)
                v24 = "hello"
                if (float32(3.14159) < v2) {
                    v26 := GenS1{F0: 0, F1: v4, F2: 0}
                    v27 := GenS1{F0: uint16(v4), F1: (v4 & (v4 * (v4 % (((v4 >> ((1) & 63))) | 1)))), F2: (v4 >> ((4294967295) & 63))}
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS1{F0: 743, F1: v4, F2: v4}
                    var v29 int64 = 42
                    _ = v28
                    _ = v29
                }
                v19 -= 0.0
            }
            v19 += float32((0.1 + (1.5 - 3.14159)))
            _ = fi25
        }
        fmt.Println(v19, v24, fi23)
        _ = fi23
        _ = v24
    }
    if true {
        if len(v1) > 0 {
            v1[0] = v4
        }
    }
    var v30 uint32 = (^(v4))
    v31 := GenS1{F0: 255, F1: 715, F2: (v4 << ((843) & 63))}
    var v32 int16 = 1000
    var v33 uint64 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
