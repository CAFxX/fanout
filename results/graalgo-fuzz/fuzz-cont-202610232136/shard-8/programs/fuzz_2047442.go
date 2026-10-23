package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint16; F2 string }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int16, p1 bool, p2 bool) uint8 {
    fmt.Println(p1, p0)
    v5 := []int{32767}
    if p2 {
        p2 = p1
        p0 -= 1
    } else {
        switch 127 {
        case 299:
            var v6 int64 = 0
            _ = v6
        case 1:
            var v7 int16 = p0
            _ = v7
        case 128:
            if len(v5) > 0 {
                v5[0] = -1
            }
        default:
            v8 := GenS1{F0: 455, F1: 18446744073709551615, F2: 2147483647}
            _ = v8
        }
        if len(v5) > 0 {
            v5[0] = 1
        }
    }
    if p1 {
        if (p2 || ((!true) || p2)) {
            for fi9 := 0; fi9 < 1; fi9++ {
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 uint = 1
                    v12 := GenS1{F0: 2, F1: (^(v11)), F2: 2147483647}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 string = "\n\t"
                fmt.Println(p2)
                _ = fi9
                _ = v13
            }
            v5 = append(v5, -1)
        }
    }
    p1 = (-799 <= 329)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 int16, p1 string, p2 string) string {
    v14 := []int{656, 128, 42}
    v15 := GenS2{F0: "\n\t"}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return p1
}

func main() {
    var v1 int16 = 1
    var v2 int = 983
    v3 := make(map[string]bool)
    var v4 int16 = (-(v1))
    switch "" {
    case "\n\t":
        v16 := GenS1{F0: 402, F1: 1, F2: 399}
        _ = v16
    case "a":
        fmt.Println(v1, v1)
    case "":
        v17 := v3[""]
        _ = v17
    default:
        var v18 int64 = -9223372036854775808
        _ = v18
    }
    var v19 int16 = (v4 + v1)
    fmt.Println(v1, v1)
    if (255 > 144) {
        var v20 uint64 = 19
        v21 := []int{v2}
        _ = v20
        _ = v21
    } else {
        switch v2 {
        case 4294967295:
            if false {
                switch "hello" {
                case "hello":
                    v1 = 1
                default:
                    v22 := GenS3{F0: 45}
                    _ = v22
                }
            }
        case -1:
            v1 = v1
        default:
            v4 = v1
        }
        v23 := GenS4{F0: 1, F1: 712, F2: "a"}
        _ = v23
    }
    switch (v2 & (v2 >> ((506) & 63))) {
    case -1:
        var v24 uint32 = 2147483647
        _ = v24
    case 830:
        v3["a"] = (!(1 > 4294967295))
    default:
        v4 -= -1
    }
    v25 := GenS3{F0: -703}
    v26 := GenS5{F0: 2}
    fmt.Println(v19)
    var v27 int64 = -482
    var v28 uint = 18446744073709551615
    var v29 uint32 = 2
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
