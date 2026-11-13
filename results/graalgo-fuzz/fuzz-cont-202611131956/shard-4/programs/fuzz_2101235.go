package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int16 }
type GenS2 struct { F0 int16; F1 int64; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
func genFn0(p0 bool) bool {
    p0 = (!((p0 || (-2.25 > 1.5)) || true))
    for fi4 := 0; fi4 < 20; fi4++ {
        p0 = ((fi4 & fi4) >= -1)
        p0 = (p0 && (0 >= 1))
        _ = fi4
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        switch "" {
        case "\n\t":
            var v6 uint = 18446744073709551615
            _ = v6
        case "hello":
            v7 := []int{972, fi5, fi5}
            _ = v7
        }
        _ = fi5
    }
    p0 = (4294967295 == 18446744073709551615)
    _ = p0
    return p0
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v8 := GenS1{F0: 42, F1: -9223372036854775808, F2: 32767}
    for fi9 := 0; fi9 < 15; fi9++ {
        v1 = v1
        _ = fi9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(v1, v1)
        for fi11 := 0; fi11 < 17; fi11++ {
            if v1 {
                v1 = true
                switch (fi11 & (fi11 - fi10)) {
                case -605:
                    v1 = (!v1)
                default:
                    v3[("hello" + "")] = 32767
                }
                v12 := []int{32767, (fi11 | (-(fi11))), 201, fi11}
                _ = v12
            }
            v1 = (!(42 < 604))
            _ = fi11
        }
        _ = fi10
    }
    if v1 {
        v13 := []int{127, 2147483647, 240}
        if v1 {
            fmt.Println(v1)
            var v14 float64 = -2.25
            var v15 float32 = 1.5
            _ = v14
            _ = v15
        } else {
            var v16 uint8 = 127
            _ = v16
        }
        v17 := GenS2{F0: 42, F1: 1, F2: 9223372036854775807}
        _ = v13
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v19 := GenS2{F0: -32768, F1: -9223372036854775808, F2: 32767}
        v20 := []int{2147483647, fi18}
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{-1}
    for fi22 := 0; fi22 < 10; fi22++ {
        v1 = v1
        fmt.Println(fi22, fi22)
        _ = fi22
    }
    fmt.Println(v1)
    switch -1 {
    case -516:
        if true {
            var v23 int16 = 1000
            _ = v23
        } else {
            delete(v2, 2)
            if v1 {
                var v24 bool = false
                _ = v24
            }
            v21 = append(v21, 2147483647)
        }
    case 32768:
        switch -2 {
        case 0:
            var v25 int16 = -32768
            _ = v25
        case 9223372036854775807:
            v1 = v1
        case 42:
            v26 := GenS3{F0: ((("a" + ("" + "a")) + "a") + "")}
            _ = v26
        default:
            fmt.Println(v1)
        }
    default:
        v2[96] = "hello"
    }
    var v27 uint32 = 2
    var v28 float32 = (float32((3.14159 / 1.5)) + float32(float64(1.5)))
    v27 = (^(v27))
    v21 = append(v21, 127)
    v29 := GenS4{F0: 31}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v8
}
