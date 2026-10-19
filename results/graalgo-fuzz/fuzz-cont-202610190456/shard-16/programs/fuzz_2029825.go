package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 uint, p1 int16, p2 uint8) string {
    p2 += 128
    for fi5 := 0; fi5 < 2; fi5++ {
        p2 = 255
        var v6 float64 = 0.1
        v6 = v6
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("\n\t" + "a")
}
func genFn1(p0 int16, p1 bool) int {
    v8 := []int{697, 4294967295, -2147483648}
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int8 = -128
        _ = fi9
        _ = v10
    }
    var v11 float32 = -2.25
    switch "a" {
    case "hello":
        if (!true) {
            switch "hello" {
            case "a":
                p1 = (!p1)
            }
        }
    case "":
        for fi12 := 0; fi12 < 2; fi12++ {
            switch "foo bar" {
            case "":
                p0 = (p0 ^ p0)
            default:
                switch (fi12 - -668) {
                case 2147483647:
                    v11 += 0.0
                case 541:
                    fmt.Println(fi12, fi12, p1)
                }
            }
            _ = fi12
        }
    }
    switch 127 {
    case 42:
        if true {
            v11 = v11
        }
    case -307:
        fmt.Println(v11)
    case 4294967295:
        _ = len(v8)
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return 2147483647
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]bool)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]bool)
    var v13 uint32 = 558
    var v14 uint32 = 2147483648
    v15 := v2[65535]
    if v15 {
        var v16 uint8 = 42
        _ = v16
    }
    switch 1 {
    case 0:
        v15 = true
    default:
        fmt.Println(v1, v15)
    }
    v14 -= uint32(v3)
    for fi17 := 0; fi17 < 11; fi17++ {
        if (((1e10 + (0.1 + (1.5 + 1e10))) + 1.5) == (float32(float64((float32(1.5) / (0.1 + 3.14159)))) - 0.1)) {
            v18 := GenS1{F0: -32768, F1: -793}
            v14 = 0
            _ = v18
        } else {
            v19 := GenS1{F0: 0, F1: -2147483648}
            v13 -= 2
            v1 = "hello"
            _ = v19
        }
        fmt.Println(v13, v15)
        var v20 int8 = -128
        _ = fi17
        _ = v20
    }
    delete(v4, v13)
    var v21 int16 = 545
    var v22 int16 = (-(v21))
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 int64 = 987
        v13 = v14
        var v25 string = v1
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: v22, F1: 1}
    _ = len(v4)
    if (!(!v15)) {
        if true {
            var v27 uint32 = 286
            if true {
                var v28 int = 1
                _ = v28
            }
            var v29 uint16 = 42
            _ = v27
            _ = v29
        } else {
            v22 = (-(v21))
            v30 := GenS2{F0: v1}
            var v31 float32 = -2.25
            _ = v30
            _ = v31
        }
        v15 = true
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
}
