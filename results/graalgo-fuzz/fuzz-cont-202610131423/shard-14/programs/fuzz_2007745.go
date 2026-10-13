package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16; F1 int64; F2 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 int16, p1 int) uint {
    var v7 int16 = (p0 << ((0) & 63))
    var v8 uint16 = 0
    for fi9 := 0; fi9 < 20; fi9++ {
        if (true || (543 >= 65535)) {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 string = ("a" + ("" + "hello"))
                if (757 != 18446744073709551615) {
                    var v12 int16 = v7
                    p1 = int(v8)
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            fmt.Println(fi9, v8)
            switch ((("\n\t" + "hello") + "hello") + "foo bar") {
            case "hello":
                p1 = 933
            }
        }
        _ = fi9
    }
    fmt.Println(v7, v7)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 432
}
func genFn1(p0 uint16, p1 int8) int16 {
    switch 292 {
    case 32768:
        v13 := GenS1{F0: 1, F1: -636}
        _ = v13
    case 1:
        var v14 int = 1
        _ = v14
    default:
        var v15 uint16 = p0
        _ = v15
    }
    if (!(float64(((1.5 / 1.0) * (3.14159 + -2.25))) < float64(-2.25))) {
        var v16 bool = true
        _ = v16
    } else {
        var v17 int64 = 711
        var v18 string = "\n\t"
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: 0, F1: 42, F2: 42}
    v20 := []int{2147483647, 127}
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 42
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 uint8 = v1
    var v5 float32 = float32(((0.1 + (1e10 / -2.25)) - (0.1 * 1e10)))
    var v6 int64 = 951
    var v21 int8 = 127
    v22 := GenS3{F0: -80}
    switch "\n\t" {
    case "\n\t":
        v5 = float32(float64(-2.25))
    case "a":
        var v23 string = "hello"
        _ = v23
    case "hello":
        var v24 uint8 = 1
        _ = v24
    default:
        var v25 uint8 = v4
        _ = v25
    }
    delete(v2, -132)
    var v26 int = 0
    v5 = v5
    v2[v26] = 0
    delete(v3, (v26 & v26))
    switch 32767 {
    case -9223372036854775808:
        v27 := GenS2{F0: -992, F1: (v6 / (((v6 % (((v6 / ((181) | 1))) | 1))) | 1)), F2: 42}
        _ = v27
    default:
        v28 := v3[(v26 + v26)]
        _ = v28
    }
    v1 = v1
    v5 += v5
    v29 := GenS2{F0: 42, F1: (v6 * v6), F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
