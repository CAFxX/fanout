package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint) int16 {
    if (!false) {
        var v6 int8 = 1
        _ = v6
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(p0)
    case "":
        p0 = (p0 + 18446744073709551615)
    case "\n\t":
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: p0, F1: -856}
            if true {
                var v9 string = "\n\t"
                _ = v9
            }
            p0 = p0
            _ = fi7
            _ = v8
        }
    }
    p0 = 819
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: 399, F1: 0}
        _ = fi10
        _ = v11
    }
    p0 = p0
    _ = p0
    return -32768
}
func genFn1(p0 int64, p1 uint32) bool {
    p1 -= 42
    v12 := []int{-9223372036854775808, 128, 4294967295, 4294967295}
    _ = p0
    _ = p1
    _ = v12
    return false
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 uint64 = 18446744073709551615
    var v5 int16 = 1000
    v5 = v5
    if true {
        v13 := GenS1{F0: 771, F1: 248}
        switch "a" {
        case "hello":
            fmt.Println(v5)
        case "\n\t":
            if true {
                var v14 float32 = (0.0 / float32(0.1))
                v15 := []int{-959, 32768}
                var v16 uint32 = 2147483647
                _ = v14
                _ = v15
                _ = v16
            } else {
                var v17 int64 = -1
                _ = v17
            }
        }
        v18 := GenS1{F0: 42, F1: 1}
        _ = v13
        _ = v18
    } else {
        v4 += (v4 | 1)
        switch 42 {
        case 4294967295:
            v5 = genFn0(705)
        default:
            v5 = (v5 + (v5 * (^(v5))))
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            _ = len(v2)
            var v20 uint8 = 2
            v21 := GenS1{F0: 588, F1: 0}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v22 := GenS1{F0: 192, F1: -338}
    v23 := GenS2{F0: 128, F1: ("foo bar" + ("foo bar" + "foo bar")), F2: 100}
    switch 1 {
    case 638:
        if false {
            if true {
                v5 = genFn0(42)
                if (313 != 1) {
                    switch 128 {
                    case -2147483648:
                        fmt.Println(v4)
                    case 1:
                        v4 -= v4
                    case 42:
                        v4 += v4
                    }
                    var v24 uint16 = 255
                    _ = v24
                }
                var v25 float64 = 1e10
                _ = v25
            }
        }
    case 2147483647:
        v5 = (v5 & (v5 | (v5 >> ((2) & 63))))
    case 42:
        var v26 uint32 = 2147483647
        _ = v26
    default:
        fmt.Println(v4, v4)
    }
    var v27 uint64 = v4
    var v28 int64 = 604
    var v29 string = (("\n\t" + "a") + "\n\t")
    fmt.Println(v27, v5, v27)
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS3{F0: fi30}
        _ = fi30
        _ = v31
    }
    var v32 uint64 = v27
    v33 := GenS1{F0: uint(v27), F1: -191}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
