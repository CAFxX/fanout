package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 float32 = 3.14159
    var v2 int = 32768
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v5 := []uint32{209}
    var v6 uint64 = 18446744073709551615
    var v7 bool = (!false)
    switch v2 {
    case 128:
        switch 863 {
        case 32768:
            v5 = append(v5, 2)
        case 4294967295:
            var v8 uint8 = 42
            _ = v8
        default:
            var v9 int = (v2 - -63)
            _ = v9
        }
    case 4294967295:
        fmt.Println(v7, v2, v7)
    }
    var v10 uint64 = 1
    var v11 uint64 = v6
    v12 := GenS1{F0: 82, F1: 1}
    fmt.Println(v11)
    if ((42 < -951) && (!(32767 > 727))) {
        v2 += (v2 & v2)
    } else {
        var v13 int64 = 0
        var v14 string = "a"
        _ = len(v4)
        _ = v13
        _ = v14
    }
    var v15 uint32 = 915
    _ = len(v5)
    var v16 string = ("" + "foo bar")
    var v17 uint64 = 4294967295
    for fi18 := 0; fi18 < 9; fi18++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if (v1 >= ((v1 / (-2.25 + (3.14159 / (1e10 / v1)))) / 1.0)) {
            v17 += (^(v11))
            v10 = 528
            if true {
                v1 += (0.0 * (float32(1e10) + (float32((-2.25 + 1e10)) * (1.5 + 0.0))))
            } else {
                v19 := GenS2{F0: 836, F1: -9223372036854775808}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v19
            }
        } else {
            var v20 uint = 65535
            v10 += v10
            _ = v20
        }
        _ = fi18
    }
    _ = len(v5)
    _ = len(v4)
    fmt.Println(v10)
    switch "hello" {
    case "":
        v21 := GenS2{F0: -32768, F1: 9223372036854775807}
        _ = v21
    case "hello":
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 uint32 = (v15 % (((^(v15))) | 1))
                v25 := GenS2{F0: -32768, F1: -9223372036854775808}
                _ = fi23
                _ = v24
                _ = v25
            }
            fmt.Println(v17, v1, v7)
            if v7 {
                var v26 float64 = (float64(((((0.0 + float32(((0.1 / float64((0.0 - v1))) / (1e10 * float64(0.1))))) * (v1 - (0.0 + 0.0))) / v1) / v1)) + 0.0)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v26
            } else {
                v6 -= v6
                var v27 string = v16
                _ = v27
            }
            _ = fi22
        }
    case "a":
        var v28 uint8 = 1
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
