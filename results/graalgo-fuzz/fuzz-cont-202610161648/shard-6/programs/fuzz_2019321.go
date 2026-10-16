package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int64; F2 uint32 }
func genFn0(p0 bool, p1 int16) int32 {
    v5 := []int{0}
    var v6 string = "\n\t"
    switch 826 {
    case -1:
        var v7 uint8 = 2
        _ = v7
    case 0:
        if len(v5) > 0 {
            v5[0] = 2147483647
        }
    case 42:
        fmt.Println(v6, p1, p0)
    }
    v6 = (((("\n\t" + (v6 + (v6 + v6))) + v6) + (v6 + ("foo bar" + v6))) + ("\n\t" + ((v6 + "") + "a")))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -514
}
func genFn1(p0 uint64, p1 int64, p2 int64) int64 {
    var v8 int32 = 2147483647
    v9 := GenS1{F0: 1000, F1: 1}
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (p1 | p2)
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = ((3.14159 * (0.1 * 3.14159)) * 0.1)
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    for fi10 := 0; fi10 < 17; fi10++ {
        switch fi10 {
        case 413:
            fmt.Println(fi10, fi10, fi10)
        case 0:
            var v11 uint = 607
            _ = v11
        default:
            switch "hello" {
            case "a":
                v12 := GenS1{F0: -1, F1: 9223372036854775807}
                _ = v12
            default:
                v2 = 1e10
            }
        }
        v2 = v2
        v13 := GenS2{F0: 50, F1: 109}
        _ = fi10
        _ = v13
    }
    v14 := GenS3{F0: genFn0((2 >= 4294967295), 32767)}
    fmt.Println(v2, v2)
    if (4294967295 <= 550) {
        var v15 uint64 = 2
        _ = v15
    }
    v16 := []int{0}
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            if (true && true) {
                var v19 int = (fi18 >> ((412) & 63))
                var v20 string = ("\n\t" + "a")
                _ = v19
                _ = v20
            } else {
                if false {
                    var v21 uint64 = 288
                    _ = v21
                } else {
                    v1 = (v1 + 255)
                }
                var v22 uint64 = 276
                var v23 uint64 = 18446744073709551615
                _ = v22
                _ = v23
            }
            v2 = v2
            var v24 int32 = -55
            var v25 uint16 = (v1 * (v1 - v1))
            if (v24 > -130) {
                switch (fi17 * fi17) {
                case 42:
                    var v26 uint64 = 42
                    _ = v26
                case 0:
                    v1 = 32767
                }
                var v27 uint32 = 4294967295
                _ = v27
            }
            _ = fi18
            _ = v24
            _ = v25
        }
        _ = fi17
    }
    v2 = (1e10 / 0.1)
    v2 = float64(1.5)
    v28 := GenS4{F0: 998, F1: 9223372036854775807, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
