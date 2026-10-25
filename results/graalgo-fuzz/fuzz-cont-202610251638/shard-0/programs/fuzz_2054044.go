package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 uint8, p1 uint16) uint32 {
    p0 += (p0 >> ((255) & 63))
    if false {
        v5 := GenS1{F0: 1, F1: 1, F2: 2147483648}
        _ = v5
    } else {
        var v6 uint = 4294967295
        _ = v6
    }
    _ = p0
    _ = p1
    return 502
}

func main() {
    var v1 int = 128
    v2 := []int{v1}
    v3 := []string{"", "hello", "hello", "a"}
    v4 := make(map[int]string)
    if false {
        var v7 uint8 = 1
        fmt.Println(v7)
        var v8 float64 = ((0.0 * -2.25) * float64((1.5 - 3.14159)))
        _ = v7
        _ = v8
    } else {
        var v9 uint32 = 1
        for fi10 := 0; fi10 < 4; fi10++ {
            if (29 == 0) {
                fmt.Println(v9, v1, fi10)
            }
            var v11 float64 = (((float64(0.1) / (1e10 / -2.25)) + (3.14159 + 1.5)) / (1.5 / (3.14159 + 1.5)))
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    var v12 int = 4294967295
    if true {
        v13 := GenS2{F0: 4294967295, F1: "a"}
        switch v1 {
        case 32767:
            v14 := GenS3{F0: 0, F1: (v12 | (-(v1)))}
            _ = v14
        case -510:
            var v15 int64 = 1
            _ = v15
        default:
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 bool = (4294967295 < 150)
                v17 = (v17 || v17)
                _ = fi16
                _ = v17
            }
        }
        fmt.Println(v1, v12)
        _ = v13
    } else {
        _ = len(v2)
    }
    var v18 uint = 255
    if false {
        v1 = v12
    } else {
        var v19 uint = v18
        _ = v19
    }
    var v20 int = (v1 + 82)
    v20 = (v12 / (((v1 * (-(v1)))) | 1))
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(v18, v12)
        if (!false) {
            var v22 uint64 = 18446744073709551615
            var v23 float32 = (1e10 * -2.25)
            v24 := GenS3{F0: -1, F1: v1}
            _ = v22
            _ = v23
            _ = v24
        } else {
            switch (("" + ("\n\t" + "foo bar")) + ("" + "foo bar")) {
            case "":
                v18 = v18
            case "a":
                fmt.Println(v1, v12)
            case "foo bar":
                v1 += (v1 - 9223372036854775807)
            default:
                v2 = append(v2, -454)
            }
            v12 = 4294967295
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 float64 = (1e10 / 3.14159)
                _ = fi25
                _ = v26
            }
        }
        var v27 uint32 = 852
        _ = fi21
        _ = v27
    }
    var v28 int8 = 100
    delete(v4, (v20 / ((v12) | 1)))
    var v29 uint16 = 65535
    var v30 uint16 = (v29 & v29)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
