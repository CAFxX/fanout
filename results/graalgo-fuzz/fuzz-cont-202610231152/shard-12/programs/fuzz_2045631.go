package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16; F1 string; F2 string }
func genFn0(p0 string) uint16 {
    if (p0 == "\n\t") {
        p0 = "foo bar"
        p0 = p0
        v4 := GenS1{F0: (1 != 1)}
        _ = v4
    } else {
        v5 := GenS2{F0: 42, F1: 701, F2: 18446744073709551615}
        _ = v5
    }
    var v6 uint16 = 32767
    v6 = (v6 << ((197) & 63))
    _ = p0
    _ = v6
    return v6
}
func genFn1(p0 uint64, p1 string, p2 uint32) uint8 {
    switch 32767 {
    case -617:
        fmt.Println(p1, p1)
    case -503:
        v7 := []int{42, -1, 83}
        _ = v7
    case 32768:
        var v8 int64 = 42
        _ = v8
    }
    p1 = "\n\t"
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int8 = -119
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]string)
    v3 := []uint32{0, 164, 65535}
    switch "" {
    case "":
        var v11 uint64 = 370
        _ = v11
    default:
        v3 = append(v3, 459)
    }
    var v12 uint64 = 42
    v13 := GenS1{F0: ((65535 != 65535) && false)}
    v14 := v1[761]
    fmt.Println(v14)
    var v15 int = 2147483647
    if (-1 == -941) {
        v12 = (v12 >> ((1) & 63))
    }
    v12 += (v12 * v12)
    v16 := GenS1{F0: ((v12 >> ((255) & 63)) <= v12)}
    v14 = v14
    if (!(2147483647 >= -1)) {
        fmt.Println(v14, v14, v14)
        var v17 uint64 = (v12 & 134)
        v18 := GenS3{F0: -32768, F1: (("a" + ("a" + ("foo bar" + ""))) + "hello"), F2: "a"}
        _ = v17
        _ = v18
    } else {
        var v19 uint16 = 65535
        var v20 int64 = 1
        _ = v19
        _ = v20
    }
    v12 += v12
    _ = len(v1)
    if v14 {
        v21 := GenS3{F0: -967, F1: "foo bar", F2: ("foo bar" + "foo bar")}
        var v22 float32 = 1e10
        _ = v21
        _ = v22
    }
    var v23 string = "foo bar"
    v12 = (^(v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
}
