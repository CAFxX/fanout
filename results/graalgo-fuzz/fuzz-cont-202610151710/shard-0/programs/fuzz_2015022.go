package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int64) uint64 {
    v4 := []int{1, 32767, 32768}
    switch 975 {
    case 42:
        p0 += 1
    }
    _ = p0
    _ = p1
    _ = v4
    return 2
}
func genFn1(p0 uint8) uint {
    switch 9223372036854775807 {
    case -9223372036854775808:
        v5 := GenS1{F0: 9223372036854775807}
        _ = v5
    default:
        if true {
            fmt.Println(p0)
            p0 = p0
            if (!false) {
                if (!(!true)) {
                    v6 := GenS1{F0: 32767}
                    _ = v6
                } else {
                    v7 := []int{2147483647}
                    var v8 int16 = 0
                    _ = v7
                    _ = v8
                }
                var v9 uint8 = (p0 & p0)
                _ = v9
            }
        }
    }
    fmt.Println(p0)
    v10 := GenS1{F0: 0}
    _ = p0
    _ = v10
    return 2
}

func main() {
    var v1 float32 = (0.0 / 1.0)
    var v2 int = 127
    var v3 int8 = 1
    v11 := GenS1{F0: (v2 & (v2 & v2))}
    fmt.Println(v3, v1, v3)
    v12 := GenS2{F0: v3, F1: 95}
    v1 += v1
    switch (v2 / ((v2) | 1)) {
    case -181:
        var v13 float64 = (float64(float32((((float64(1e10) + 3.14159) - (1.5 + float64(v1))) + float64(v1)))) - 0.1)
        _ = v13
    }
    if true {
        if true {
            var v14 uint64 = 2
            _ = v14
        }
    }
    v15 := GenS2{F0: 100, F1: 0}
    v2 = v2
    fmt.Println(v1)
    v1 += 0.0
    for fi16 := 0; fi16 < 4; fi16++ {
        if false {
            switch "" {
            case "a":
                v1 = ((v1 + (v1 / v1)) + 0.0)
            case "hello":
                v2 = -326
            case "foo bar":
                for fi17 := 0; fi17 < 12; fi17++ {
                    v1 = v1
                    switch "" {
                    case "\n\t":
                        v1 += v1
                    default:
                        var v18 int = v2
                        _ = v18
                    }
                    _ = fi17
                }
            }
            if false {
                for fi19 := 0; fi19 < 20; fi19++ {
                    v20 := []int{9223372036854775807, v2}
                    var v21 uint64 = 0
                    var v22 uint = 0
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
        }
        v3 += (v3 | 77)
        _ = fi16
    }
    fmt.Println(v2, v2, v3)
    if false {
        v1 = 0.0
        fmt.Println(v1, v3)
        var v23 uint32 = 0
        _ = v23
    } else {
        v24 := GenS3{F0: 2}
        _ = v24
    }
    if false {
        v1 = 0.1
        v25 := GenS1{F0: v2}
        var v26 bool = (!(v2 > (v2 * -9223372036854775808)))
        _ = v25
        _ = v26
    } else {
        for fi27 := 0; fi27 < 1; fi27++ {
            v28 := []int{(-(fi27))}
            var v29 int32 = 2147483647
            v1 += v1
            _ = fi27
            _ = v28
            _ = v29
        }
        v30 := []int{v2, (v2 ^ 32767), (v2 >> ((651) & 63))}
        if (float64((v1 / v1)) <= 3.14159) {
            v2 = (v2 >> ((1) & 63))
        }
        _ = v30
    }
    v31 := []int{9223372036854775807, (v2 / (((v2 & (v2 * (v2 ^ 42)))) | 1))}
    v1 = ((v1 - ((v1 / 1.0) - 3.14159)) * 1e10)
    var v32 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
}
