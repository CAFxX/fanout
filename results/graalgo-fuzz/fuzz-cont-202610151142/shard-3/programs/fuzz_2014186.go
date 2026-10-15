package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 float64 = 3.14159
    v2 := []int{9223372036854775807, 32767, 194}
    v3 := make(map[uint32]string)
    var v4 uint16 = 389
    var v5 string = ("" + (("foo bar" + "") + "foo bar"))
    v6 := make(map[int]bool)
    var v7 float64 = float64(float32((1e10 + 0.0)))
    var v8 int = 0
    var v9 uint = 2
    v9 += v9
    v9 += (v9 & (v9 ^ v9))
    var v10 float32 = ((1e10 + (3.14159 / 1.5)) * 0.1)
    if true {
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS1{F0: 42}
            v1 = (float64(0.0) * ((v1 * (float64(3.14159) - (float64(v10) - v1))) + (v1 / (-2.25 + (v1 + (v7 * ((1.5 - float64((float32(0.1) / ((0.0 / float32(v7)) + -2.25)))) + float64(v10))))))))
            _ = fi11
            _ = v12
        }
        switch "hello" {
        case "":
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 uint16 = 107
                _ = fi13
                _ = v14
            }
        case "hello":
            var v15 uint = v9
            _ = v15
        }
        v16 := GenS1{F0: -9223372036854775808}
        _ = v16
    } else {
        v17 := GenS1{F0: 42}
        var v18 float64 = 3.14159
        _ = v17
        _ = v18
    }
    v6[(v8 - -1)] = ((!true) && (false || true))
    var v19 int = 9223372036854775807
    var v20 int = (v8 ^ v8)
    v8 += v20
    for fi21 := 0; fi21 < 13; fi21++ {
        v4 += v4
        v22 := GenS1{F0: 0}
        v23 := GenS1{F0: -280}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: 60, F1: (128 == uint8(v4))}
    for fi25 := 0; fi25 < 16; fi25++ {
        switch v19 {
        case -9223372036854775808:
            var v26 uint8 = 255
            _ = v26
        case 500:
            if (!true) {
                v5 = v5
            } else {
                var v27 int64 = 9223372036854775807
                _ = v27
            }
        default:
            switch v5 {
            case "hello":
                var v28 uint32 = 2147483647
                _ = v28
            default:
                v29 := GenS3{F0: v4, F1: (v4 << ((v9) & 63))}
                _ = v29
            }
        }
        v9 = v9
        _ = fi25
    }
    fmt.Println(v4, v10)
    v30 := GenS2{F0: 2, F1: ((float32((v1 + ((float64(float32((1e10 - (v7 / v1)))) * -2.25) * (0.1 + (float64((1e10 + (((v10 / 3.14159) + 1.5) / v10))) / (-2.25 / 1.0)))))) + ((v10 - v10) - (v10 / v10))) >= v10)}
    v10 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
