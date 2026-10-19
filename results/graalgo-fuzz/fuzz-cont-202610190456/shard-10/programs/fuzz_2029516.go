package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 int8) uint32 {
    p0 += p0
    if (4294967295 != 572) {
        fmt.Println(p0)
    } else {
        p0 -= p0
        if (!true) {
            var v5 uint64 = 563
            v6 := []int{}
            _ = v5
            _ = v6
        }
    }
    _ = p0
    return 42
}

func main() {
    var v1 int8 = 0
    var v2 float32 = (1e10 * -2.25)
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    v7 := []int{9223372036854775807}
    if ((true || (true || true)) || false) {
        var v8 uint = 2
        _ = v8
    }
    switch 42 {
    case 1:
        v2 -= -2.25
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 int16 = 1000
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    }
    fmt.Println(v1, v2)
    switch "\n\t" {
    case "\n\t":
        v1 += (v1 ^ v1)
    case "foo bar":
        switch 1 {
        case 4294967295:
            var v12 uint8 = 128
            _ = v12
        default:
            fmt.Println(v2, v2)
        }
    default:
        if false {
            v1 += v1
            if true {
                v1 = v1
                if true {
                    fmt.Println(v1)
                } else {
                    v1 += v1
                }
                v13 := GenS1{F0: 1}
                _ = v13
            } else {
                if (1 < 0) {
                    v2 += float32(((((1e10 + 3.14159) * 1.5) / -2.25) + (0.1 + -2.25)))
                } else {
                    for fi14 := 0; fi14 < 16; fi14++ {
                        var v15 string = "a"
                        v15 = "foo bar"
                        var v16 int16 = 1
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    v17 := GenS2{F0: (v1 >> ((116) & 63)), F1: "a"}
                    v1 += (v1 << ((172) & 63))
                    _ = v17
                }
                switch ("\n\t" + ("foo bar" + "a")) {
                case "a":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                default:
                    var v18 float64 = 0.1
                    _ = v18
                }
            }
        }
    }
    var v19 uint64 = 625
    v20 := GenS3{F0: 0, F1: 65535}
    v1 = (v1 + (v1 ^ 65))
    var v21 uint = 541
    v22 := GenS3{F0: (v19 * (v19 | (v19 & (^(v19))))), F1: v21}
    var v23 float32 = float32(1e10)
    fmt.Println(v21)
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 float64 = 0.0
        v7 = append(v7, (fi24 * -2147483648))
        for fi26 := 0; fi26 < 2; fi26++ {
            if len(v7) > 0 {
                v7[0] = fi26
            }
            var v27 bool = false
            var v28 float64 = v25
            var v29 float32 = 3.14159
            fmt.Println(v27, v25)
            if true {
                v23 = 1.5
            } else {
                v25 += v28
            }
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi24
        _ = v25
    }
    if (!true) {
        v23 += (v23 / 1.5)
        v23 += v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
