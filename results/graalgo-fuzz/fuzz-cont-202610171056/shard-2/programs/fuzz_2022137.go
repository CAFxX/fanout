package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint; F1 int; F2 int8 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint16 }
func genFn0(p0 string, p1 int32) uint8 {
    p1 = (p1 | (p1 >> ((42) & 63)))
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(fi5, p0)
        _ = fi5
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        if (!(1 > 925)) {
            var v7 int16 = 32767
            var v8 string = p0
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 uint16 = 65535
                _ = fi9
                _ = v10
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]uint32)
    var v3 uint64 = 0
    v4 := make(map[string]int)
    v11 := GenS1{F0: 0, F1: -128}
    var v12 bool = ((-2.25 * 1.5) != ((3.14159 - 0.0) / 0.1))
    for fi13 := 0; fi13 < 15; fi13++ {
        switch fi13 {
        case 127:
            if v12 {
                fmt.Println(fi13, fi13)
            } else {
                v14 := []int{(fi13 + (fi13 + fi13))}
                var v15 uint = 18446744073709551615
                _ = v14
                _ = v15
            }
        case 1:
            if (!true) {
                switch (fi13 + (-(fi13))) {
                case -398:
                    delete(v2, 42)
                case 125:
                    v16 := []int{42, int(v3)}
                    _ = v16
                default:
                    var v17 int = (fi13 * fi13)
                    _ = v17
                }
            }
        }
        _ = fi13
    }
    fmt.Println(v1, v12)
    if v12 {
        fmt.Println(v3, v3, v3)
        fmt.Println(v1, v3)
    } else {
        v18 := GenS1{F0: -128, F1: 42}
        _ = v18
    }
    fmt.Println(v3, v1, v1)
    v19 := GenS2{F0: v1, F1: -861, F2: 1}
    if v12 {
        switch "" {
        case "a":
            switch "" {
            case "foo bar":
                v12 = false
            case "hello":
                fmt.Println(v12, v12)
            case "a":
                v1 = 42
            }
        default:
            v20 := GenS2{F0: (v1 + v1), F1: -9223372036854775808, F2: -119}
            _ = v20
        }
        var v21 string = (("" + ("" + ("a" + ""))) + ("hello" + ((("a" + ("a" + (("a" + "") + "\n\t"))) + "a") + "")))
        _ = v21
    } else {
        v22 := GenS3{F0: (v3 / ((v3) | 1)), F1: -9223372036854775808, F2: uint16(v1)}
        _ = v22
    }
    v23 := []int{1}
    var v24 float32 = 1e10
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 uint64 = (v3 ^ v3)
        _ = fi25
        _ = v26
    }
    fmt.Println(v1, v1)
    var v27 uint64 = (^(v3))
    switch 4294967295 {
    case -1:
        var v28 int64 = -352
        _ = v28
    case 32767:
        if len(v23) > 0 {
            v23[0] = 32768
        }
    case -153:
        var v29 int16 = 0
        _ = v29
    default:
        var v30 uint = (v1 << (((v1 | (v1 - 0))) & 63))
        _ = v30
    }
    if len(v23) > 0 {
        v23[0] = 346
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
