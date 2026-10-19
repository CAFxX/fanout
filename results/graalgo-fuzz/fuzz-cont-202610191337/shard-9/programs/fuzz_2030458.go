package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int16; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 string }
func genFn0(p0 string, p1 uint8) int16 {
    if false {
        for fi4 := 0; fi4 < 9; fi4++ {
            p1 = (p1 & (^(p1)))
            _ = fi4
        }
        var v5 uint8 = 1
        p1 = (^(v5))
        _ = v5
    } else {
        v6 := []int{}
        if false {
            if len(v6) > 0 {
                v6[0] = 32767
            }
            _ = len(v6)
        } else {
            fmt.Println(p0, p1)
            p0 = "hello"
            for fi7 := 0; fi7 < 15; fi7++ {
                p0 = p0
                _ = fi7
            }
        }
        v8 := GenS1{F0: 123, F1: 18446744073709551615}
        _ = v6
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        v10 := []int{(fi9 / ((1) | 1))}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi9
        _ = v10
    }
    if true {
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS1{F0: 21, F1: 2}
            var v13 float64 = 0.1
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int64 = -1
        _ = v14
    }
    p1 += 42
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 int8, p1 uint16, p2 uint) string {
    if true {
        p2 = p2
        v15 := GenS2{F0: 127, F1: 1000, F2: 128}
        _ = v15
    } else {
        p0 = p0
        v16 := []int{}
        fmt.Println(p1, p1, p0)
        _ = v16
    }
    v17 := GenS3{F0: 2}
    if false {
        fmt.Println(p2, p0)
    } else {
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return ""
}

func main() {
    var v1 int8 = 0
    var v2 string = ""
    var v3 uint64 = 18446744073709551615
    if (255 <= 308) {
        switch (v2 + "hello") {
        case "a":
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := []int{9223372036854775807, (^(fi18)), 983, (-(fi18))}
                _ = fi18
                _ = v19
            }
        case "hello":
            v20 := []int{128, 997, 9223372036854775807}
            _ = v20
        }
        fmt.Println(v3, v2)
    } else {
        var v21 float64 = (float64(1e10) - ((0.1 - (0.1 - 0.1)) / 0.1))
        var v22 float32 = 0.0
        fmt.Println(v21, v22, v21)
        _ = v21
        _ = v22
    }
    fmt.Println(v2, v1, v1)
    switch -971 {
    case 9223372036854775807:
        v23 := []int{32768, 9223372036854775807, -493}
        _ = v23
    case 343:
        var v24 float32 = (((1e10 / -2.25) + ((0.1 - -2.25) * 0.0)) / (1.5 / -2.25))
        _ = v24
    case 571:
        for fi25 := 0; fi25 < 15; fi25++ {
            for fi26 := 0; fi26 < 10; fi26++ {
                fmt.Println(fi25, v3)
                v3 -= v3
                if (927 <= 4294967295) {
                    v2 = "\n\t"
                    v2 = v2
                } else {
                    var v27 int64 = -1
                    _ = v27
                }
                _ = fi26
            }
            var v28 uint64 = 2
            _ = fi25
            _ = v28
        }
    }
    v29 := GenS2{F0: 128, F1: -797, F2: 128}
    fmt.Println(v1, v1)
    fmt.Println(v3, v1)
    v1 = v1
    fmt.Println(v1)
    v30 := GenS4{F0: v3, F1: 2, F2: v2}
    v31 := []int{335}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
