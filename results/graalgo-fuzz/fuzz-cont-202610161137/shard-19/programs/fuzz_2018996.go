package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 int16) uint8 {
    v5 := GenS1{F0: 2, F1: -812, F2: true}
    if (344 > 32767) {
        for fi6 := 0; fi6 < 14; fi6++ {
            if ((!false) || true) {
                if true {
                    var v7 int16 = (p0 / (((p0 & (^(p0)))) | 1))
                    _ = v7
                }
                v8 := GenS2{F0: 128, F1: 469, F2: 2147483647}
                v9 := GenS2{F0: 128, F1: 18446744073709551615, F2: 0}
                _ = v8
                _ = v9
            }
            if (567 == 0) {
                switch "foo bar" {
                case "\n\t":
                    for fi10 := 0; fi10 < 9; fi10++ {
                        var v11 float64 = 1.5
                        _ = fi10
                        _ = v11
                    }
                case "foo bar":
                    switch "" {
                    case "hello":
                        var v12 uint = 513
                        _ = v12
                    }
                }
                p0 -= p0
            }
            switch "a" {
            case "\n\t":
                switch "\n\t" {
                case "\n\t":
                    var v13 uint16 = 65535
                    _ = v13
                default:
                    v14 := GenS3{F0: -1, F1: (^(fi6))}
                    _ = v14
                }
            }
            _ = fi6
        }
    }
    switch "foo bar" {
    case "":
        p0 += 32767
    case "\n\t":
        var v15 uint64 = 18446744073709551615
        _ = v15
    case "foo bar":
        if (0 == 0) {
            fmt.Println(p0)
            var v16 uint16 = 291
            _ = v16
        }
    default:
        v17 := []int{4294967295, 2147483647}
        _ = v17
    }
    switch "\n\t" {
    case "foo bar":
        switch 42 {
        case 4294967295:
            var v18 uint16 = 65535
            _ = v18
        case 32768:
            fmt.Println(p0)
        }
    default:
        v19 := GenS2{F0: 12, F1: 255, F2: 0}
        _ = v19
    }
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 1
    var v3 int8 = 127
    v4 := []string{((("a" + "") + "a") + ""), "foo bar", "\n\t", (("" + "") + (("hello" + "foo bar") + ("" + ("\n\t" + "a"))))}
    if false {
        fmt.Println(v3, v2)
        delete(v1, 42)
        delete(v1, 2147483647)
    } else {
        if (genFn0(-604) != 2) {
            v3 = (v3 >> ((1) & 63))
            v2 = v2
        } else {
            for fi20 := 0; fi20 < 18; fi20++ {
                v21 := GenS3{F0: 0, F1: fi20}
                var v22 int8 = (v3 / ((v3) | 1))
                v2 -= 255
                _ = fi20
                _ = v21
                _ = v22
            }
            var v23 bool = (!false)
            var v24 int64 = 0
            _ = v23
            _ = v24
        }
    }
    var v25 string = (("" + "\n\t") + "\n\t")
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 int16 = 32767
        _ = fi26
        _ = v27
    }
    fmt.Println(v25, v25)
    var v28 float32 = 0.0
    v25 = v25
    var v29 uint8 = 2
    v30 := GenS3{F0: 42, F1: 1}
    if (-448 > 0) {
        v28 = v28
    }
    var v31 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
