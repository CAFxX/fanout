package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 int64, p1 uint, p2 int32) string {
    p0 += p0
    p1 += (p1 & p1)
    p2 = p2
    var v7 uint64 = 2
    if false {
        if (0.1 > 0.0) {
            var v8 int = -9223372036854775808
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: 1}
                var v11 uint8 = uint8(v7)
                var v12 int32 = p2
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v8
        } else {
            p1 -= p1
        }
        switch 1 {
        case -198:
            var v13 float32 = ((1e10 - 0.0) / 3.14159)
            _ = v13
        case -1:
            v7 = 362
        default:
            var v14 bool = true
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int = (^(fi15))
            v17 := GenS1{F0: 127}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "hello"
}

func main() {
    v1 := []int{1, 128}
    v2 := []uint32{2147483648, 2147483648, 1, 1}
    v3 := make(map[string]bool)
    var v4 uint32 = 2
    v5 := make(map[string]bool)
    var v6 uint32 = v4
    v4 -= (v6 % (((v6 & v6)) | 1))
    for fi18 := 0; fi18 < 18; fi18++ {
        v6 -= v6
        var v19 uint16 = 1
        _ = fi18
        _ = v19
    }
    _ = len(v1)
    v4 += v4
    if (0 < -326) {
        fmt.Println(v4, v4)
    }
    if true {
        var v20 float32 = 3.14159
        var v21 int = -1
        v4 = (v6 >> ((168) & 63))
        _ = v20
        _ = v21
    } else {
        var v22 int64 = 0
        var v23 string = "a"
        _ = v22
        _ = v23
    }
    switch 32768 {
    case 2147483647:
        for fi24 := 0; fi24 < 8; fi24++ {
            for fi25 := 0; fi25 < 18; fi25++ {
                v6 = v6
                v4 -= 474
                var v26 int8 = -1
                var v27 uint64 = 1
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi24
        }
    case -58:
        delete(v3, "")
    }
    if (-2147483648 == 2147483647) {
        var v28 uint16 = 32767
        v29 := GenS1{F0: 255}
        _ = v28
        _ = v29
    }
    var v30 string = "a"
    if true {
        v31 := GenS1{F0: 170}
        _ = v31
    } else {
        v32 := GenS2{F0: -14, F1: -326}
        _ = v32
    }
    v4 = (v6 | (v6 & 1))
    var v33 float32 = (1.5 * ((1e10 / float32((1.5 - 1e10))) + 1e10))
    v6 = v4
    if ((0.1 + 0.1) >= ((1e10 * -2.25) / 3.14159)) {
        v34 := GenS2{F0: -765, F1: 1}
        _ = v34
    }
    var v35 uint64 = 779
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
