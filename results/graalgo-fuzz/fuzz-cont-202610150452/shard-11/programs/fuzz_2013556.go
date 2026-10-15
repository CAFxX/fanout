package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 0
    v3 := []string{}
    var v4 float32 = (0.1 * 3.14159)
    var v5 int8 = 93
    var v6 int = 32767
    fmt.Println(v2, v6, v2)
    if (!(((v6 << ((433) & 63)) <= -1) || true)) {
        var v7 float64 = (-2.25 - ((0.0 * float64(((v4 + ((v4 / ((((v4 * 0.0) / ((v4 / float32((((0.0 + 0.0) / ((1e10 + 1e10) + float64(v4))) / (float64(v4) / (float64(v4) / 3.14159))))) + ((1.5 - ((float32((3.14159 + 0.0)) - 0.1) / v4)) / v4))) / (v4 * v4)) * ((0.1 / v4) * ((((float32(float64((((v4 / v4) / 1e10) + v4))) / float32(float64(float32(float64((v4 / (v4 * ((v4 * v4) + v4)))))))) - v4) - ((1.5 * float32((((float64(v4) - 1.5) * 0.1) * 3.14159))) - 1e10)) + (v4 * v4))))) + v4)) * (1e10 - v4)))) - -2.25))
        _ = v7
    } else {
        var v8 uint8 = (v2 % ((v2) | 1))
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        v5 = v5
        _ = fi9
    }
    v2 = v2
    fmt.Println(v4)
    v10 := GenS1{F0: -929, F1: 0, F2: 2147483647}
    v6 = v6
    v11 := GenS1{F0: 695, F1: 255, F2: 65535}
    if true {
        var v12 uint32 = 4294967295
        var v13 int64 = -126
        v13 += int64(v12)
        _ = v12
        _ = v13
    } else {
        v2 += v2
    }
    v4 = v4
    var v14 int64 = 9223372036854775807
    if ((0.0 / -2.25) != (1.5 * 1e10)) {
        v6 = (v6 % ((v6) | 1))
    }
    if (255 == 42) {
        var v15 bool = true
        _ = v15
    } else {
        var v16 int16 = 1
        _ = v16
    }
    if (((1e10 / v4) - v4) <= v4) {
        if (((!(!(!((^(v14)) != -361)))) || false) && true) {
            if true {
                v14 -= (v14 ^ -9223372036854775808)
            }
        } else {
            if (766 <= 672) {
                fmt.Println(v4, v2)
                fmt.Println(v2)
            }
        }
        if (2 <= (v2 - 0)) {
            v5 += v5
            if true {
                var v17 uint8 = (v2 >> ((219) & 63))
                fmt.Println(v6)
                v17 = v17
                _ = v17
            } else {
                fmt.Println(v14, v4)
            }
            switch "hello" {
            case "":
                v3 = append(v3, (("hello" + ("a" + "foo bar")) + "a"))
            case "hello":
                var v18 uint16 = 82
                _ = v18
            case "foo bar":
                v1[(v6 + v6)] = 1
            }
        } else {
            delete(v1, (v6 | 32768))
            var v19 uint64 = 2
            _ = v19
        }
        var v20 uint64 = 4294967295
        _ = v20
    }
    var v21 int64 = -5
    if false {
        v22 := v1[v6]
        _ = v22
    } else {
        v2 += 204
        v6 = (v6 / ((v6) | 1))
        v6 -= v6
    }
    var v23 uint16 = 201
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
