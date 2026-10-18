package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int; F1 int32; F2 int32 }
type GenS3 struct { F0 string; F1 uint8; F2 int }
func genFn0(p0 uint16, p1 string) uint16 {
    fmt.Println(p0, p0)
    if (1.5 > ((float64(0.1) + (((3.14159 / ((1e10 - 0.0) + 1.5)) * 3.14159) * ((((0.0 * 0.1) * 3.14159) * 1e10) / 1.0))) / 1e10)) {
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 1
    v3 := make(map[string]int)
    var v4 int64 = 795
    var v5 int8 = -1
    var v6 int = 42
    v7 := []int{}
    var v8 bool = true
    v5 -= v5
    var v9 int8 = (^(v5))
    _ = len(v3)
    switch ("\n\t" + (("foo bar" + ("foo bar" + "")) + "hello")) {
    case "\n\t":
        v6 -= (v6 % ((int(v5)) | 1))
    case "hello":
        v10 := GenS1{F0: 18446744073709551615, F1: 349}
        _ = v10
    case "foo bar":
        var v11 int32 = -198
        _ = v11
    default:
        fmt.Println(v8)
    }
    var v12 uint64 = 2
    v12 += v12
    v2 -= v2
    for fi13 := 0; fi13 < 13; fi13++ {
        if (1e10 == float32(3.14159)) {
            var v14 uint64 = 1
            _ = v14
        }
        var v15 int32 = -1
        switch int(v4) {
        case 9223372036854775807:
            var v16 uint64 = v12
            _ = v16
        case 42:
            v9 += (^(v9))
        }
        _ = fi13
        _ = v15
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if (v2 > (v2 >> ((65535) & 63))) {
            for fi18 := 0; fi18 < 3; fi18++ {
                v19 := GenS2{F0: (fi18 ^ 128), F1: 1, F2: -1}
                _ = fi18
                _ = v19
            }
        }
        fmt.Println(v2)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 int8 = (^(v9))
        if v8 {
            if (255 >= 0) {
                v22 := GenS3{F0: "foo bar", F1: uint8(v2), F2: 127}
                v8 = false
                v21 -= (v9 % ((v9) | 1))
                _ = v22
            }
        }
        _ = fi20
        _ = v21
    }
    v1["a"] = ((!(-32768 >= 1)) && (float32((1e10 * 0.1)) != 1.5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
