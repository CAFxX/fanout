package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint }

func main() {
    v1 := []int{0, 128, 42}
    var v2 string = ""
    v3 := make(map[uint32]bool)
    var v4 uint64 = 18446744073709551615
    var v5 int64 = 672
    var v6 bool = false
    var v7 bool = true
    v4 = v4
    if (!((v2 + ("\n\t" + "hello")) < v2)) {
        v5 = (v5 + 9223372036854775807)
        var v8 uint8 = 0
        var v9 uint32 = 2147483647
        _ = v8
        _ = v9
    }
    var v10 uint64 = v4
    v11 := GenS1{F0: -960, F1: -577}
    var v12 int = 32767
    if (42 <= 393) {
        for fi13 := 0; fi13 < 16; fi13++ {
            v4 = 856
            fmt.Println(v12)
            _ = fi13
        }
        _ = len(v1)
    } else {
        _ = len(v1)
        switch (v12 - v12) {
        case 1:
            v6 = ((127 < 111) || v6)
        case 0:
            delete(v3, 2147483647)
        case 128:
            var v14 bool = v6
            _ = v14
        default:
            v15 := GenS2{F0: 0, F1: v7}
            _ = v15
        }
    }
    if false {
        if false {
            var v16 float32 = 0.1
            fmt.Println(v16, v2)
            switch (v2 + (((((v2 + v2) + v2) + "foo bar") + "") + v2)) {
            case "":
                if (!(65535 <= 1)) {
                    v12 -= v12
                }
            default:
                if true {
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v4 -= v10
                        _ = fi17
                    }
                    fmt.Println(v4, v5, v7)
                    if (v12 >= (v12 + v12)) {
                        v3[462] = v6
                        delete(v3, 2147483648)
                        var v18 uint8 = 127
                        _ = v18
                    } else {
                        var v19 int32 = -1
                        v6 = v7
                        _ = v19
                    }
                }
            }
            _ = v16
        }
    } else {
        var v20 int64 = 42
        var v21 string = "\n\t"
        _ = v20
        _ = v21
    }
    v4 -= v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v22 := GenS2{F0: -702, F1: (!(-1 == 0))}
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := GenS3{F0: 65535}
        switch "a" {
        case "hello":
            v25 := GenS1{F0: 240, F1: -802}
            _ = v25
        case "foo bar":
            v1 = append(v1, -1)
        case "a":
            if (-69 <= 0) {
                v10 = v10
                delete(v3, 1)
                var v26 float32 = 0.0
                _ = v26
            }
        default:
            var v27 int64 = v5
            _ = v27
        }
        v28 := GenS1{F0: -333, F1: int32(v4)}
        _ = fi23
        _ = v24
        _ = v28
    }
    if (!(!(!((3.14159 / -2.25) <= ((float64(1.5) / 1.0) + (1.5 - (((float64(0.0) / -2.25) - 3.14159) / 1.0))))))) {
        v29 := v3[2147483647]
        fmt.Println(v29, v6)
        _ = v29
    } else {
        var v30 float32 = 1e10
        v6 = v6
        v31 := GenS3{F0: 0}
        _ = v30
        _ = v31
    }
    fmt.Println(v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
