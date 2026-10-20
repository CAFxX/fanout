package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int16 }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
type GenS3 struct { F0 int64 }
func genFn0() uint64 {
    if false {
        v6 := GenS1{F0: 0, F1: 107, F2: 531}
        _ = v6
    } else {
        var v7 int32 = 384
        fmt.Println(v7)
        _ = v7
    }
    if false {
        if false {
            for fi8 := 0; fi8 < 7; fi8++ {
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 string = (("hello" + "hello") + "")
                    var v11 uint16 = 1
                    if (v11 <= v11) {
                        var v12 int32 = 346
                        v12 = (-(v12))
                        v13 := []int{fi9}
                        _ = v12
                        _ = v13
                    } else {
                        switch v10 {
                        case "foo bar":
                            v14 := GenS2{F0: -121, F1: v10, F2: 9223372036854775807}
                            v15 := []int{(fi8 * -611)}
                            _ = v14
                            _ = v15
                        }
                        v10 = "hello"
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
        }
    } else {
        switch "\n\t" {
        case "hello":
            v16 := GenS2{F0: -1, F1: "\n\t", F2: 1}
            _ = v16
        default:
            var v17 float64 = float64(0.0)
            _ = v17
        }
        var v18 int16 = -1
        v19 := []int{-1, int(v18), 29, 9223372036854775807}
        _ = v18
        _ = v19
    }
    return 1
}

func main() {
    v1 := []int{-1, 42, 4294967295, 9223372036854775807}
    v2 := make(map[int]string)
    var v3 float64 = 3.14159
    var v4 bool = false
    var v5 uint32 = 486
    v20 := GenS3{F0: 9223372036854775807}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 uint = 4294967295
        var v23 int8 = -48
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v5, v5, v5)
    v24 := GenS3{F0: 1}
    var v25 string = ("\n\t" + "a")
    var v26 float32 = 0.0
    var v27 int16 = -387
    v27 -= v27
    switch "\n\t" {
    case "hello":
        v25 = v25
    }
    var v28 int = -9223372036854775808
    var v29 uint64 = 18446744073709551615
    v30 := GenS2{F0: 42, F1: "hello", F2: 523}
    var v31 bool = (!(!(605 == 18446744073709551615)))
    switch v25 {
    case "hello":
        v27 += v27
    case "\n\t":
        if false {
            v1 = append(v1, (v28 | 4294967295))
        } else {
            var v32 string = (v25 + (v25 + (v25 + "")))
            v33 := GenS2{F0: v27, F1: (v32 + "a"), F2: 1}
            v2[v28] = (v32 + v25)
            _ = v32
            _ = v33
        }
    case "foo bar":
        v4 = ((v26 * ((float32((v3 / 1.5)) * (v26 - v26)) + -2.25)) > -2.25)
    }
    var v34 int = v28
    for fi35 := 0; fi35 < 19; fi35++ {
        v31 = (0 >= 42)
        var v36 uint64 = (v29 * (v29 & v29))
        _ = fi35
        _ = v36
    }
    var v37 float64 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v37
    _ = v4
    _ = v5
}
