package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint64; F2 int }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int; F1 int; F2 uint }

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 uint64 = 2
    v5 := make(map[uint32]string)
    v4 = 18446744073709551615
    var v6 uint64 = v4
    var v7 uint32 = 1
    v8 := []int{-946, -1, -872, 32768}
    fmt.Println(v7)
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(fi10)
            if ((fi10 < fi9) || true) {
                v4 = (v6 | (v4 ^ v6))
                if false {
                    v11 := GenS1{F0: 180}
                    var v12 int = -902
                    switch (fi10 << ((170) & 63)) {
                    case 0:
                        v13 := GenS1{F0: 1}
                        _ = v13
                    case -984:
                        v14 := v2[778]
                        _ = v14
                    case 127:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    default:
                        v6 -= v4
                    }
                    _ = v11
                    _ = v12
                } else {
                    delete(v5, v7)
                    v15 := GenS2{F0: -55, F1: v4, F2: (fi9 * -258)}
                    var v16 float32 = (1e10 - 0.1)
                    _ = v15
                    _ = v16
                }
                var v17 uint16 = 2
                _ = v17
            } else {
                switch ("a" + "a") {
                case "":
                    switch (^(fi10)) {
                    case 32767:
                        v18 := GenS1{F0: 405}
                        _ = v18
                    case 1:
                        v1 = -60
                    case -9223372036854775808:
                        switch fi9 {
                        case 32768:
                            v7 = uint32(v4)
                        case 32767:
                            var v19 uint16 = 2
                            _ = v19
                        default:
                            var v20 uint = 1
                            _ = v20
                        }
                    }
                default:
                    switch fi10 {
                    case -14:
                        _ = len(v5)
                    case 32768:
                        v6 = v6
                    case 0:
                        var v21 int64 = -436
                        var v22 uint32 = v7
                        v7 += v7
                        v7 = 65535
                        _ = v21
                        _ = v22
                    }
                    v7 += 42
                }
            }
            if true {
                v1 = v1
                v7 -= uint32(v6)
                v23 := GenS2{F0: -416, F1: v6, F2: fi9}
                v7 = (v7 & (v7 | v7))
                _ = v23
            } else {
                v24 := GenS2{F0: -820, F1: 1, F2: (fi9 & (^(fi10)))}
                _ = v24
            }
            v7 -= 936
            _ = fi10
        }
        v4 = v6
        v6 = v4
        _ = fi9
    }
    if (false && (int(v7) > 128)) {
        switch "foo bar" {
        case "":
            if false {
                var v25 uint8 = 28
                v6 = (v6 >> ((42) & 63))
                fmt.Println(v7)
                _ = v25
            }
        case "\n\t":
            var v26 string = "foo bar"
            _ = v26
        case "hello":
            v1 -= (v1 << ((0) & 63))
        }
        if (0 > 2) {
            fmt.Println(v4, v6, v6)
            for fi27 := 0; fi27 < 7; fi27++ {
                v1 = v1
                _ = fi27
            }
            if true {
                if (127 > 0) {
                    v4 -= v6
                    if false {
                        var v28 uint8 = 1
                        v3[("hello" + "")] = true
                        _ = v28
                    } else {
                        v29 := GenS3{F0: 195, F1: int(v4)}
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        v30 := GenS3{F0: -9223372036854775808, F1: 273}
                        _ = v29
                        _ = v30
                    }
                }
            } else {
                v31 := GenS4{F0: -9223372036854775808, F1: 0, F2: 18446744073709551615}
                _ = v31
            }
        }
    }
    var v32 float64 = 1e10
    var v33 float64 = (float64((-2.25 / 0.1)) * v32)
    v34 := GenS3{F0: -9223372036854775808, F1: 9223372036854775807}
    fmt.Println(v4)
    var v35 int8 = 111
    var v36 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
