package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 int8, p1 int32) string {
    fmt.Println(p1)
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 707, F1: "\n\t", F2: 127}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    return "a"
}
func genFn1(p0 uint32, p1 uint8) uint8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 string = ""
        _ = fi7
        _ = v8
    }
    v9 := []int{-264, 42, 2147483647}
    var v10 uint32 = 42
    switch -989 {
    case 32768:
        var v11 int32 = -1
        _ = v11
    case 2147483647:
        switch 32768 {
        case 128:
            v12 := GenS1{F0: p0, F1: "a", F2: (p1 >> ((138) & 63))}
            _ = v12
        case -1:
            v10 = (^(p0))
        }
    default:
        var v13 uint = 3
        _ = v13
    }
    if true {
        var v14 int = 1
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 int32 = 42
    var v2 int32 = v1
    var v3 uint8 = 195
    var v4 float32 = float32(1e10)
    switch "" {
    case "hello":
        v15 := GenS1{F0: 2147483647, F1: "\n\t", F2: v3}
        _ = v15
    case "":
        var v16 uint64 = 1
        _ = v16
    }
    var v17 int64 = 204
    v18 := GenS1{F0: 0, F1: "a", F2: 108}
    var v19 uint8 = v3
    v17 = v17
    var v20 float32 = -2.25
    v21 := GenS1{F0: 65535, F1: "a", F2: genFn1(1, 127)}
    var v22 int16 = -32768
    var v23 int8 = 100
    v24 := []int{-1, -770, 128, 2147483647}
    var v25 string = ""
    v26 := GenS2{F0: v23, F1: int16(v17)}
    if len(v24) > 0 {
        v24[0] = 32767
    }
    _ = len(v24)
    var v27 string = (("foo bar" + (v25 + v25)) + v25)
    var v28 float32 = (v4 + ((v4 - 1.5) + v4))
    v23 += v23
    v1 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
