package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint }
type GenS3 struct { F0 uint64; F1 int; F2 bool }

func main() {
    var v1 float32 = (float32(float64((0.0 - 0.1))) / (-2.25 + float32(1.5)))
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v5 int8 = 42
    var v6 int = 4294967295
    v7 := v4[4294967295]
    delete(v3, v7)
    v7 = v7
    switch v7 {
    case "foo bar":
        v8 := GenS1{F0: 255}
        _ = v8
    }
    var v9 string = v7
    fmt.Println(v1)
    var v10 uint64 = 234
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(v7, fi12)
            v7 = ("foo bar" + v7)
            _ = fi12
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 float32 = 1e10
            v15 := GenS2{F0: 0, F1: 1, F2: 145}
            switch v9 {
            case "hello":
                switch (v9 + (v9 + ("a" + v7))) {
                case "a":
                    if ((!false) || false) {
                        var v16 uint = 64
                        v17 := GenS2{F0: 489, F1: 2, F2: v16}
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(v9, v14, v14)
                    }
                }
            case "":
                if (2147483647 == 1) {
                    v10 -= 18446744073709551615
                } else {
                    v9 = ""
                    var v18 int16 = -1
                    _ = v18
                }
            }
            v10 += (v10 * (v10 & (^(v10))))
            fmt.Println(v5, v7, v1)
            _ = fi13
            _ = v14
            _ = v15
        }
        v19 := GenS1{F0: 128}
        _ = fi11
        _ = v19
    }
    var v20 uint = 2
    v21 := []int{-9223372036854775808}
    v22 := GenS1{F0: uint8(v10)}
    v23 := GenS3{F0: uint64(v20), F1: 128, F2: true}
    var v24 uint32 = 65535
    var v25 uint8 = 1
    if (127 <= (v5 >> ((v20) & 63))) {
        var v26 float64 = -2.25
        _ = len(v21)
        var v27 int16 = 1000
        _ = v26
        _ = v27
    } else {
        v25 = v25
        v28 := GenS3{F0: 0, F1: (v6 + 780), F2: (!false)}
        _ = v28
    }
    if false {
        v10 = 9
        v20 -= v20
        switch 0 {
        case 643:
            if (((float64(0.0) + 0.0) - 3.14159) == 0.1) {
                fmt.Println(v9)
                if false {
                    v5 = v5
                    v6 = 42
                }
            }
        default:
            switch (-(v6)) {
            case 1:
                switch "foo bar" {
                case "a":
                    v20 -= v20
                case "foo bar":
                    v29 := GenS2{F0: 1, F1: v20, F2: v20}
                    _ = v29
                case "":
                    v7 = (v9 + "hello")
                default:
                    fmt.Println(v9, v20)
                }
            case 0:
                v30 := GenS3{F0: v10, F1: v6, F2: false}
                _ = v30
            case -31:
                var v31 uint = (v20 + v20)
                _ = v31
            default:
                v10 = 821
            }
        }
    } else {
        if ((0.1 - 0.1) < (((3.14159 / (1e10 + float64(v1))) * 3.14159) * float64((v1 / v1)))) {
            var v32 uint = v20
            _ = v32
        } else {
            v6 = v6
            v33 := v2[v6]
            var v34 int32 = 1
            _ = v33
            _ = v34
        }
        v35 := GenS3{F0: 4294967295, F1: (v6 & 128), F2: true}
        var v36 uint8 = (v25 % ((v25) | 1))
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
