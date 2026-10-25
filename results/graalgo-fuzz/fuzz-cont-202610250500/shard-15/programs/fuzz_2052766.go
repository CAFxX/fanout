package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int; F2 int64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint, p1 bool) bool {
    switch ("foo bar" + "foo bar") {
    case "hello":
        v6 := []int{2147483647}
        _ = v6
    default:
        fmt.Println(p1)
    }
    p0 = (p0 / (((p0 % ((p0) | 1))) | 1))
    v7 := GenS1{F0: 32767}
    v8 := GenS1{F0: 591}
    if (183 < 255) {
        p0 -= 42
        for fi9 := 0; fi9 < 7; fi9++ {
            if ((33 != 98) && ((p1 || (128 <= 1)) || (((0.1 - 0.1) <= 1e10) && p1))) {
                p1 = true
                p0 -= p0
                var v10 bool = (!(!p1))
                _ = v10
            }
            _ = fi9
        }
    } else {
        switch "a" {
        case "hello":
            var v11 uint16 = 42
            _ = v11
        default:
            switch "" {
            case "foo bar":
                var v12 int64 = 340
                _ = v12
            case "\n\t":
                p1 = p1
            case "a":
                if ((p1 || p1) || p1) {
                    fmt.Println(p0, p1)
                    v13 := []int{127, 32767, 32767, 0}
                    _ = v13
                } else {
                    var v14 int16 = 1
                    p0 = (p0 * (p0 - 0))
                    v15 := []int{-2147483648, 0, 536}
                    _ = v14
                    _ = v15
                }
            default:
                switch -2147483648 {
                case -2147483648:
                    p0 -= 426
                case -296:
                    switch -747 {
                    case -388:
                        var v16 string = ("\n\t" + "")
                        _ = v16
                    case 9223372036854775807:
                        p1 = ((!(true && p1)) || p1)
                    case 1000:
                        var v17 float32 = 1e10
                        _ = v17
                    default:
                        if true {
                            p1 = p1
                            var v18 float32 = ((3.14159 * float32(0.0)) - (1.5 - float32(1e10)))
                            _ = v18
                        }
                    }
                case 0:
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v20 := GenS1{F0: 42}
                        var v21 int16 = 0
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
            }
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 bool) int16 {
    p0 = (910 <= -1)
    var v22 uint16 = 42
    _ = p0
    _ = v22
    return -273
}

func main() {
    v1 := []uint32{4294967295, 4294967295}
    var v2 int = -37
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    var v5 bool = false
    var v23 uint32 = 2147483647
    v5 = (v5 || v5)
    var v24 string = (((("\n\t" + "foo bar") + ("foo bar" + "")) + "hello") + "hello")
    var v25 int64 = 497
    v26 := GenS1{F0: 32767}
    fmt.Println(v2, v24, v2)
    if false {
        switch ("" + (v24 + v24)) {
        case "":
            for fi27 := 0; fi27 < 1; fi27++ {
                v2 = 32768
                _ = fi27
            }
        default:
            v24 = v24
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            fmt.Println(v2, v24, v25)
            v29 := GenS1{F0: -32768}
            fmt.Println(v5)
            _ = fi28
            _ = v29
        }
        v5 = false
    }
    if (-1 > -128) {
        fmt.Println(v25, v5, v25)
    }
    if (v24 < v24) {
        v30 := GenS2{F0: 4294967295, F1: -585, F2: (v25 - (^(v25)))}
        _ = len(v1)
        _ = v30
    } else {
        v24 = ""
        fmt.Println(v5, v23)
    }
    v25 -= v25
    fmt.Println(v2)
    v31 := GenS3{F0: (v23 - 1)}
    v32 := GenS2{F0: 1, F1: v2, F2: (v25 * v25)}
    v24 = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
