package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 bool }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint16 }
func genFn0() uint8 {
    if (743 == 780) {
        if (42 <= 127) {
            var v6 int8 = 42
            if (false || true) {
                for fi7 := 0; fi7 < 11; fi7++ {
                    v8 := GenS1{F0: 2147483647, F1: fi7, F2: true}
                    v9 := []int{-1, int(v6), fi7, (^(fi7))}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                var v10 uint64 = 76
                _ = v10
            }
            _ = v6
        } else {
            var v11 int8 = 42
            v11 = 65
            _ = v11
        }
    } else {
        v12 := GenS1{F0: 4294967295, F1: -120, F2: (1e10 <= (1e10 * 0.0))}
        _ = v12
    }
    var v13 int16 = 255
    _ = v13
    return 112
}

func main() {
    v1 := []string{}
    var v2 int64 = 0
    v3 := []int{522, -336, 32767}
    var v4 int8 = 42
    v5 := []int{-1, 42}
    fmt.Println(v2, v4)
    var v14 uint32 = 1
    if true {
        fmt.Println(v14, v14, v2)
        v15 := GenS2{F0: 32767, F1: 32767}
        _ = v15
    } else {
        fmt.Println(v2, v2, v2)
    }
    var v16 float64 = float64(1e10)
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(v16, fi17)
        var v18 int32 = 0
        fmt.Println(v16, v4, fi17)
        _ = fi17
        _ = v18
    }
    v19 := GenS3{F0: v2}
    if (((("a" + "hello") + ("hello" + (("a" + ("" + "foo bar")) + ""))) != "") || true) {
        v16 -= (v16 - v16)
        switch 128 {
        case 4294967295:
            v20 := GenS1{F0: 2, F1: -780, F2: true}
            _ = v20
        case 42:
            _ = len(v1)
        case 175:
            _ = len(v3)
        default:
            v21 := GenS4{F0: (^(v14)), F1: 4294967295, F2: 42}
            _ = v21
        }
    } else {
        v2 = -534
        fmt.Println(v2, v2)
    }
    if len(v3) > 0 {
        v3[0] = 128
    }
    switch (("a" + "a") + ("" + "")) {
    case "hello":
        v4 = 84
    case "":
        v4 = (v4 * (v4 + 100))
    }
    var v22 int64 = (v2 / ((v2) | 1))
    var v23 uint8 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
