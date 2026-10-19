package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint64; F1 int; F2 int }
func genFn0(p0 int64, p1 string, p2 int16) uint {
    p1 = p1
    p0 -= p0
    v6 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}

func main() {
    v1 := []int{4294967295, 127, 812}
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    var v4 uint16 = 1
    v5 := []uint32{947, 42}
    v7 := GenS1{F0: -9223372036854775808, F1: 252, F2: 2}
    v4 = (v4 / ((v4) | 1))
    v8 := GenS1{F0: 42, F1: 0, F2: 65535}
    var v9 int32 = -2147483648
    if ("" < ("\n\t" + ("a" + (("\n\t" + "hello") + "a")))) {
        var v10 int32 = v9
        var v11 uint64 = 1
        for fi12 := 0; fi12 < 4; fi12++ {
            v13 := GenS1{F0: 127, F1: 255, F2: 290}
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    }
    switch 0 {
    case 546:
        v9 = v9
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            v3 = v3
            fmt.Println(v4, fi14, v3)
            _ = fi15
        }
        var v16 bool = true
        fmt.Println(v16, v4, fi14)
        _ = fi14
        _ = v16
    }
    v5 = append(v5, 2)
    if (float32(0.1) <= (3.14159 - -2.25)) {
        for fi17 := 0; fi17 < 20; fi17++ {
            switch (fi17 - 2147483647) {
            case 9223372036854775807:
                var v18 bool = false
                _ = v18
            case 537:
                switch "foo bar" {
                case "hello":
                    var v19 uint32 = 700
                    _ = v19
                case "":
                    v20 := GenS2{F0: -778, F1: (100 <= 127)}
                    _ = v20
                case "a":
                    v3 -= v3
                }
            }
            var v21 uint8 = 16
            var v22 int32 = 0
            _ = fi17
            _ = v21
            _ = v22
        }
        var v23 uint64 = 2
        v4 = (v4 - v4)
        _ = v23
    }
    switch ("foo bar" + "hello") {
    case "":
        v24 := GenS3{F0: 42, F1: 0, F2: -282}
        _ = v24
    case "a":
        switch -1 {
        case -1:
            var v25 int16 = v3
            _ = v25
        case 403:
            var v26 int64 = -1
            _ = v26
        default:
            fmt.Println(v9)
        }
    }
    var v27 int32 = v9
    fmt.Println(v27, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
