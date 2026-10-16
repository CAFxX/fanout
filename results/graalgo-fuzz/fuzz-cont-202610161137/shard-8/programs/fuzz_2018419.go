package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int }
func genFn0(p0 uint8, p1 uint16, p2 uint) int32 {
    var v7 uint32 = 0
    if (-1 > 42) {
        fmt.Println(v7, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}

func main() {
    var v1 int = 1
    var v2 bool = true
    var v3 int16 = -393
    var v4 float64 = 1.5
    var v5 int8 = 100
    v6 := make(map[string]string)
    v1 = -9223372036854775808
    v8 := []int{}
    fmt.Println(v5)
    if v2 {
        for fi9 := 0; fi9 < 19; fi9++ {
            v3 -= 1
            var v10 int32 = 2147483647
            _ = fi9
            _ = v10
        }
        fmt.Println(v3, v2)
        var v11 uint = 255
        _ = v11
    }
    var v12 uint64 = 4294967295
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS1{F0: "\n\t"}
        fmt.Println(v3, v5)
        _ = fi13
        _ = v14
    }
    switch v1 {
    case -97:
        if v2 {
            fmt.Println(v1, v4, v12)
            delete(v6, "hello")
            v2 = (!v2)
        } else {
            var v15 uint32 = 2147483647
            if (127 != -732) {
                v4 = float64(float32((((v4 + (float64(0.1) * v4)) + v4) - 3.14159)))
            } else {
                if true {
                    var v16 int16 = (v3 ^ 42)
                    var v17 uint64 = v12
                    var v18 uint8 = 207
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    fmt.Println(v5, v15, v4)
                    for fi19 := 0; fi19 < 3; fi19++ {
                        switch v1 {
                        case 0:
                            v1 += int(v5)
                        case 9223372036854775807:
                            fmt.Println(v3, v5)
                        default:
                            v4 = float64(3.14159)
                        }
                        fmt.Println(v4, v3)
                        _ = fi19
                    }
                }
            }
            _ = v15
        }
    case 9223372036854775807:
        if true {
            for fi20 := 0; fi20 < 1; fi20++ {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v5 = -128
                    v22 := GenS1{F0: ""}
                    _ = fi21
                    _ = v22
                }
                var v23 float32 = 0.0
                _ = fi20
                _ = v23
            }
        } else {
            var v24 float32 = 0.0
            _ = v24
        }
    case 42:
        v8 = append(v8, (v1 ^ 32767))
    }
    switch (v1 - int(v5)) {
    case -610:
        fmt.Println(v5, v1, v3)
    case 42:
        fmt.Println(v3)
    default:
        v1 = v1
    }
    fmt.Println(v3)
    var v25 float64 = (v4 + (0.0 + (v4 * float64(0.1))))
    v4 = v4
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS2{F0: (v1 * (^(v1)))}
        _ = fi26
        _ = v27
    }
    if v2 {
        var v28 bool = false
        if (!v2) {
            fmt.Println(v25, v5)
            v29 := GenS3{F0: 42, F1: 1}
            _ = v29
        }
        v5 += (^(v5))
        _ = v28
    }
    fmt.Println(v25)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
