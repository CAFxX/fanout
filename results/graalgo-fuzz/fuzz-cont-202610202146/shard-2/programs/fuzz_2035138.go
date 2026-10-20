package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }

func main() {
    v1 := []int{}
    var v2 int64 = -639
    var v3 uint = 255
    v4 := make(map[uint32]bool)
    v5 := []uint32{42, 65535, 42, 2147483647}
    v6 := []string{"hello", ("hello" + ("hello" + ("\n\t" + "foo bar")))}
    if false {
        v3 = v3
    }
    v2 -= (v2 << ((v3) & 63))
    var v7 bool = (!(!true))
    if ((v3 >> ((v3) & 63)) > v3) {
        v5 = append(v5, 2147483648)
        var v8 string = "hello"
        v3 += 0
        _ = v8
    } else {
        var v9 int8 = 1
        if (146 <= 255) {
            var v10 int16 = -60
            v11 := GenS1{F0: 784, F1: v3}
            for fi12 := 0; fi12 < 6; fi12++ {
                v9 = v9
                v9 = (v9 | (v9 << ((v3) & 63)))
                _ = fi12
            }
            _ = v10
            _ = v11
        } else {
            var v13 uint32 = 2147483647
            v9 = (v9 >> ((v3) & 63))
            _ = v13
        }
        fmt.Println(v2, v7, v3)
        _ = v9
    }
    var v14 int = 128
    if v7 {
        var v15 uint32 = 2147483647
        v7 = (uint8(v3) <= 42)
        fmt.Println(v7, v14, v7)
        _ = v15
    }
    v2 += (v2 ^ 42)
    if v7 {
        switch ("foo bar" + ("a" + "\n\t")) {
        case "foo bar":
            _ = len(v6)
        case "":
            var v16 uint8 = 0
            _ = v16
        case "hello":
            var v17 int = (v14 * v14)
            _ = v17
        default:
            v18 := GenS2{F0: 622, F1: 97, F2: false}
            _ = v18
        }
        var v19 uint64 = 4294967295
        _ = v19
    } else {
        v2 = -376
        if (v7 && ((("hello" + "foo bar") + "") != "\n\t")) {
            var v20 int64 = (v2 + v2)
            for fi21 := 0; fi21 < 9; fi21++ {
                if len(v5) > 0 {
                    v5[0] = 1
                }
                v14 += (fi21 / (((v14 | 0)) | 1))
                _ = fi21
            }
            v22 := v4[42]
            _ = v20
            _ = v22
        } else {
            v6 = append(v6, "\n\t")
        }
    }
    var v23 int8 = -128
    for fi24 := 0; fi24 < 17; fi24++ {
        var v25 uint16 = 42
        _ = fi24
        _ = v25
    }
    switch "\n\t" {
    case "\n\t":
        v26 := GenS2{F0: (^(v3)), F1: 0, F2: (-1 < 111)}
        _ = v26
    case "hello":
        fmt.Println(v14, v2)
    }
    if ((2147483647 < 0) || (false || true)) {
        fmt.Println(v7)
    } else {
        v27 := GenS1{F0: 65535, F1: (^(v3))}
        fmt.Println(v2)
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
