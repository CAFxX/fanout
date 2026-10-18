package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int, p2 uint8) uint32 {
    if true {
        var v5 int8 = 114
        v6 := GenS1{F0: 32767, F1: 1}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 509, F1: 42}
    var v8 int8 = 1
    v9 := []int{(p1 + (-(p1))), (p1 - 401)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 string = ""
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS1{F0: -536, F1: 2147483648}
        fmt.Println(v4, fi10)
        v4 = "hello"
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 100}
    v13 := GenS1{F0: -105, F1: 1}
    fmt.Println(v4, v1)
    v14 := GenS3{F0: 32767, F1: 1}
    v15 := v3[525]
    var v16 bool = false
    v17 := GenS3{F0: -32768, F1: 9223372036854775807}
    v18 := GenS4{F0: 9223372036854775807}
    switch 32768 {
    case 128:
        var v19 bool = true
        _ = v19
    }
    var v20 float32 = (float32(1.5) * (((-2.25 + (0.0 + -2.25)) * 1.5) * 0.1))
    v21 := GenS1{F0: 42, F1: 785}
    for fi22 := 0; fi22 < 12; fi22++ {
        switch fi22 {
        case 0:
            v16 = v16
        case 9223372036854775807:
            if v16 {
                v23 := GenS5{F0: 1}
                _ = v23
            }
        default:
            var v24 int = (fi22 & fi22)
            _ = v24
        }
        v25 := GenS4{F0: 1}
        switch v4 {
        case "\n\t":
            v26 := []int{fi22, (-(fi22)), fi22}
            _ = v26
        case "foo bar":
            v27 := GenS4{F0: 42}
            _ = v27
        default:
            if v15 {
                if v16 {
                    var v28 bool = v16
                    var v29 int32 = 2147483647
                    _ = v28
                    _ = v29
                } else {
                    var v30 int = -2147483648
                    delete(v3, fi22)
                    _ = v30
                }
                v1 += v1
            }
        }
        _ = fi22
        _ = v25
    }
    switch v4 {
    case "\n\t":
        if v15 {
            v1 -= v1
            var v31 int64 = -1
            if (((1e10 - 3.14159) < ((-2.25 - float64(v20)) - (1.5 + 1e10))) || v16) {
                v16 = v15
            }
            _ = v31
        }
    }
    v3[-9223372036854775808] = v16
    var v32 uint32 = 861
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v4
}
