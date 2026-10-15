package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int64; F1 bool; F2 int64 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 bool; F1 int16; F2 uint64 }
func genFn0(p0 uint8, p1 int64) string {
    var v4 uint = 18446744073709551615
    var v5 float64 = float64(3.14159)
    if (p0 >= p0) {
        if false {
            switch "a" {
            case "\n\t":
                var v6 uint8 = p0
                _ = v6
            default:
                v4 -= 42
            }
            v5 -= 0.1
        } else {
            if (0 == (v4 ^ (v4 / (((v4 >> (((v4 | v4)) & 63))) | 1)))) {
                v7 := GenS1{F0: 32767, F1: p0, F2: 957}
                var v8 uint64 = 378
                switch "foo bar" {
                case "":
                    p0 -= 42
                case "hello":
                    v9 := GenS2{F0: 2, F1: 542}
                    _ = v9
                case "\n\t":
                    for fi10 := 0; fi10 < 16; fi10++ {
                        fmt.Println(v4)
                        if true {
                            v11 := GenS1{F0: 879, F1: (^(p0)), F2: v8}
                            _ = v11
                        }
                        v12 := []int{fi10, int(p0)}
                        _ = fi10
                        _ = v12
                    }
                default:
                    var v13 bool = false
                    _ = v13
                }
                _ = v7
                _ = v8
            } else {
                v14 := []int{}
                _ = v14
            }
            if false {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := []int{}
                    fmt.Println(v5)
                    var v17 float32 = ((-2.25 + 1.5) * 3.14159)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            v5 += (v5 * float64((1.5 * 0.0)))
        }
        v18 := []int{674}
        switch int(p0) {
        case 1:
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        }
        _ = v18
    } else {
        v19 := []int{-45}
        _ = v19
    }
    v20 := GenS3{F0: (p1 >> ((v4) & 63)), F1: false, F2: 42}
    v21 := GenS2{F0: 0, F1: -990}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    _ = v4
    _ = v5
    return "hello"
}
func genFn1(p0 uint8, p1 int64, p2 uint64) int64 {
    v22 := []int{0, -362, 2147483647, 32768}
    if ((false || ((("\n\t" + "a") + "hello") == "\n\t")) || true) {
        if true {
            p2 = (^(p2))
        }
        switch "" {
        case "hello":
            var v23 int = 931
            _ = v23
        }
        if false {
            var v24 int64 = p1
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        switch fi25 {
        case -9223372036854775808:
            var v26 uint64 = p2
            _ = v26
        default:
            p2 -= (^(p2))
        }
        _ = fi25
    }
    fmt.Println(p0)
    if true {
        for fi27 := 0; fi27 < 12; fi27++ {
            p2 = (p2 + 42)
            _ = fi27
        }
    } else {
        if (true || false) {
            fmt.Println(p0)
            fmt.Println(p2, p1)
            fmt.Println(p2, p0, p2)
        }
        var v28 int8 = -1
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 uint8 = (p0 | (p0 * (^(p0))))
            var v31 uint = 65535
            v31 = v31
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return p1
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 430
    var v3 uint64 = 1
    v32 := v1[-1]
    v33 := GenS4{F0: (v3 | (v3 ^ v3)), F1: false}
    v34 := []int{(v32 << ((255) & 63)), v32, (-(v32)), (v32 / ((v32) | 1))}
    v32 -= 42
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    var v35 int = v32
    v36 := GenS1{F0: -32768, F1: 139, F2: v3}
    for fi37 := 0; fi37 < 7; fi37++ {
        for fi38 := 0; fi38 < 3; fi38++ {
            switch "foo bar" {
            case "a":
                var v39 bool = ((!true) && (v2 <= 1))
                _ = v39
            case "\n\t":
                delete(v1, v35)
            default:
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
            }
            _ = fi38
        }
        var v40 uint8 = uint8(v3)
        fmt.Println(fi37)
        _ = fi37
        _ = v40
    }
    v32 += v35
    fmt.Println(v3)
    for fi41 := 0; fi41 < 3; fi41++ {
        fmt.Println(fi41, v3)
        v2 += v2
        var v42 float64 = 0.0
        _ = fi41
        _ = v42
    }
    v2 = 2
    if false {
        v43 := GenS5{F0: true, F1: 42, F2: v3}
        fmt.Println(v3, v3, v35)
        v2 = v2
        _ = v43
    } else {
        for fi44 := 0; fi44 < 2; fi44++ {
            v35 += 565
            _ = fi44
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
