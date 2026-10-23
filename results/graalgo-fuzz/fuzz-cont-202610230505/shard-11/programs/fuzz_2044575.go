package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint64) uint {
    var v7 uint8 = 42
    if (true || true) {
        v7 = (v7 | v7)
        fmt.Println(p0, p0)
    }
    fmt.Println(p0, v7)
    v8 := []int{32767, 2147483647}
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = 1
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    var v6 int32 = 42
    v6 = (v6 + (v6 ^ (v6 | (v6 >> ((genFn0(359)) & 63)))))
    var v9 int16 = -858
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v1, fi10, fi10)
        if false {
            v11 := GenS1{F0: 2, F1: (!false)}
            if false {
                v12 := GenS2{F0: 42}
                v13 := []int{}
                _ = v12
                _ = v13
            } else {
                fmt.Println(v9, v9, v6)
                switch fi10 {
                case 0:
                    var v14 float32 = float32((3.14159 / (float64(-2.25) / (1e10 * 0.1))))
                    _ = v14
                case 127:
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 uint32 = 1
                        _ = fi15
                        _ = v16
                    }
                }
                var v17 int64 = 562
                _ = v17
            }
            _ = v11
        }
        fmt.Println(fi10)
        _ = fi10
    }
    if (42 != 42) {
        v9 = 0
        var v18 uint = 2
        v9 = v9
        _ = v18
    } else {
        switch "" {
        case "foo bar":
            if (0.1 > ((0.0 + ((-2.25 * 3.14159) / (1.5 / 1e10))) * 3.14159)) {
                v6 += (v6 << ((184) & 63))
            }
        case "\n\t":
            var v19 uint16 = 574
            _ = v19
        case "":
            for fi20 := 0; fi20 < 11; fi20++ {
                for fi21 := 0; fi21 < 17; fi21++ {
                    if true {
                        v9 = (v9 + (v9 << ((304) & 63)))
                    }
                    _ = fi21
                }
                _ = fi20
            }
        }
    }
    fmt.Println(v9, v6, v6)
    v6 += (-(v6))
    fmt.Println(v6, v6, v1)
    v22 := GenS3{F0: (v6 | v6)}
    fmt.Println(v6, v9, v9)
    switch ((("foo bar" + "a") + ("a" + ("foo bar" + "a"))) + "") {
    case "hello":
        var v23 uint16 = 255
        _ = v23
    case "a":
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 uint16 = 0
            for fi26 := 0; fi26 < 8; fi26++ {
                if (127 >= -61) {
                    _ = len(v5)
                    switch "\n\t" {
                    case "\n\t":
                        switch "a" {
                        case "":
                            fmt.Println(v6)
                        case "foo bar":
                            var v27 string = "a"
                            _ = v27
                        }
                    case "a":
                        var v28 int64 = 529
                        _ = v28
                    }
                }
                v25 = v25
                v9 = v9
                switch (("foo bar" + "foo bar") + ("a" + "")) {
                case "\n\t":
                    var v29 int16 = 42
                    _ = v29
                case "a":
                    v25 -= v25
                default:
                    fmt.Println(fi26, v25)
                }
                var v30 int = fi26
                _ = fi26
                _ = v30
            }
            _ = fi24
            _ = v25
        }
    case "":
        var v31 float32 = (((0.1 * (1.5 - (3.14159 / (1e10 / 1.0)))) + ((((1e10 + (3.14159 / -2.25)) * 1.5) * 1.5) * -2.25)) * 0.1)
        _ = v31
    default:
        for fi32 := 0; fi32 < 2; fi32++ {
            v6 -= 465
            fmt.Println(v9)
            _ = fi32
        }
    }
    v9 -= (v9 * v9)
    var v33 uint = 18446744073709551615
    if (true && false) {
        v1 += (^(v1))
        for fi34 := 0; fi34 < 19; fi34++ {
            fmt.Println(v1, v33, fi34)
            v1 = v1
            _ = fi34
        }
        var v35 uint16 = 521
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
