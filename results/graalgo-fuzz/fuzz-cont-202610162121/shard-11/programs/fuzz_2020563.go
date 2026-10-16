package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 bool; F2 uint }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint16; F1 int; F2 int64 }

func main() {
    var v1 int64 = -97
    var v2 float64 = ((1.5 * 1e10) * ((-2.25 + 1e10) * (0.1 + 3.14159)))
    var v3 string = "hello"
    v4 := make(map[int]int)
    var v5 float32 = 3.14159
    var v6 int = 32767
    delete(v4, -735)
    switch v6 {
    case 902:
        v4[v6] = 2147483647
    }
    var v7 int = v6
    v5 = ((((float32((v2 + v2)) / 1e10) + 1.5) + (0.0 + 1e10)) * (((1e10 * 1.5) - ((v5 / 3.14159) * (-2.25 / float32(v2)))) / ((1e10 - 0.0) + ((float32(v2) * v5) + ((v5 * 0.1) - v5)))))
    fmt.Println(v3)
    v4[v7] = v6
    var v8 bool = false
    if (!v8) {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: 1}
            var v11 int16 = 1000
            delete(v4, (v7 + (v6 << ((255) & 63))))
            _ = fi9
            _ = v10
            _ = v11
        }
        if v8 {
            v12 := GenS1{F0: -795}
            _ = v12
        }
    }
    v13 := GenS1{F0: 714}
    if (v5 != (3.14159 - v5)) {
        v14 := GenS2{F0: (v7 + v7), F1: v8, F2: 684}
        _ = v14
    }
    v15 := GenS3{F0: 0, F1: v3}
    v16 := GenS4{F0: 65535, F1: (v6 | (v6 * (v6 ^ v7))), F2: v1}
    switch (("a" + v3) + ((v3 + "") + v3)) {
    case "":
        var v17 uint8 = 255
        _ = v17
    case "a":
        v18 := []int{(v7 - v7)}
        _ = v18
    default:
        if v8 {
            var v19 uint64 = 2
            switch v3 {
            case "":
                var v20 uint32 = 503
                _ = v20
            case "foo bar":
                fmt.Println(v6)
            case "hello":
                fmt.Println(v19, v8)
            default:
                v6 -= 142
            }
            var v21 int8 = 1
            _ = v19
            _ = v21
        }
    }
    var v22 float32 = (1e10 + (v5 * v5))
    v23 := GenS1{F0: -428}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
