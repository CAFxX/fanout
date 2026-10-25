package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint16; F1 int8; F2 int8 }
func genFn0(p0 int32, p1 uint8) string {
    var v5 uint64 = 18446744073709551615
    v5 = v5
    v6 := GenS1{F0: (p0 % ((p0) | 1)), F1: 148, F2: 0}
    for fi7 := 0; fi7 < 6; fi7++ {
        p1 -= p1
        var v8 uint8 = p1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (p0 - p0), F1: 2, F2: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int16 = 858
    v2 := make(map[string]string)
    v3 := []string{"hello", "hello", "\n\t"}
    v4 := make(map[int]bool)
    if false {
        v10 := GenS1{F0: -1, F1: 65535, F2: 1}
        _ = v10
    }
    var v11 uint32 = 4294967295
    _ = len(v4)
    v12 := GenS2{F0: 42, F1: true}
    if ((^(v1)) <= v1) {
        v13 := GenS1{F0: -2147483648, F1: 65535, F2: 768}
        _ = v13
    }
    v1 = int16(v11)
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint = 329
        _ = fi14
        _ = v15
    }
    var v16 uint32 = (v11 & 2147483647)
    v11 += v11
    var v17 int32 = 0
    for fi18 := 0; fi18 < 2; fi18++ {
        if (255 == 18446744073709551615) {
            if true {
                v19 := GenS3{F0: int16(fi18)}
                fmt.Println(v11)
                v16 -= 1
                _ = v19
            } else {
                var v20 string = ("foo bar" + "")
                var v21 bool = (!(1 < 390))
                fmt.Println(v20)
                _ = v20
                _ = v21
            }
            v22 := GenS3{F0: v1}
            _ = v22
        } else {
            var v23 bool = (0 != 65535)
            _ = v23
        }
        if (4294967295 > 18446744073709551615) {
            if true {
                v17 += (v17 >> ((18446744073709551615) & 63))
            } else {
                var v24 string = "hello"
                _ = v24
            }
            var v25 bool = (!(405 <= 368))
            v1 = (v1 ^ (v1 & (v1 ^ (v1 * (v1 / ((0) | 1))))))
            _ = v25
        } else {
            fmt.Println(v17)
            var v26 uint64 = 42
            if (1 <= (^(v1))) {
                switch -2147483648 {
                case -769:
                    v27 := GenS2{F0: v26, F1: ((true && false) && (-1 <= 418))}
                    _ = v27
                }
            }
            _ = v26
        }
        var v28 float64 = ((1e10 - 0.1) - (float64(3.14159) - 0.0))
        _ = fi18
        _ = v28
    }
    v29 := GenS4{F0: 18446744073709551615}
    v30 := GenS5{F0: (^(v11)), F1: 255}
    v31 := GenS6{F0: 610, F1: 0, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
