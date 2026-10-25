package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int8; F2 int16 }
func genFn0(p0 int32, p1 int16) uint8 {
    switch "a" {
    case "\n\t":
        switch 0 {
        case -797:
            fmt.Println(p1, p1)
        case 32767:
            v6 := GenS1{F0: 9223372036854775807, F1: 377}
            _ = v6
        default:
            fmt.Println(p0, p1)
        }
    case "a":
        switch -407 {
        case 580:
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 int64 = -9223372036854775808
                v9 := GenS1{F0: v8, F1: 65535}
                v10 := []int{fi7, -200}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        case 0:
            v11 := GenS2{F0: 9223372036854775807, F1: p0}
            _ = v11
        default:
            var v12 int = 32767
            _ = v12
        }
    default:
        var v13 uint16 = 0
        _ = v13
    }
    p1 += 0
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint8 = 255
    v2 := []int{}
    var v3 float64 = 0.1
    var v4 uint64 = 0
    var v5 float32 = (3.14159 + float32(float64(0.1)))
    for fi14 := 0; fi14 < 4; fi14++ {
        v1 = v1
        var v15 uint8 = v1
        _ = fi14
        _ = v15
    }
    fmt.Println(v5, v1)
    switch "a" {
    case "":
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 string = "hello"
            var v18 int64 = -9223372036854775808
            v19 := GenS1{F0: v18, F1: 2}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
    case "\n\t":
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(v4, fi20, v1)
            _ = fi20
        }
    }
    v2 = append(v2, 1)
    fmt.Println(v4)
    v3 = v3
    var v21 string = "hello"
    var v22 int8 = 79
    var v23 uint32 = 774
    v23 -= (v23 - 391)
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v24 uint64 = v4
    for fi25 := 0; fi25 < 5; fi25++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "hello" {
        case "":
            for fi26 := 0; fi26 < 5; fi26++ {
                v24 = (v24 ^ 1)
                if true {
                    switch (^(fi25)) {
                    case 318:
                        v1 += v1
                    case -9223372036854775808:
                        var v27 uint16 = 435
                        _ = v27
                    case -928:
                        switch -438 {
                        case 1:
                            var v28 uint32 = v23
                            _ = v28
                        }
                    default:
                        v4 = v24
                    }
                    v29 := GenS3{F0: 42}
                    _ = v29
                }
                fmt.Println(fi26)
                _ = fi26
            }
        case "foo bar":
            if false {
                v24 -= (v4 & 0)
                var v30 float32 = v5
                for fi31 := 0; fi31 < 19; fi31++ {
                    fmt.Println(v1, v30, v22)
                    v23 = (^(v23))
                    _ = fi31
                }
                _ = v30
            } else {
                v4 += 42
                v1 += v1
            }
        default:
            fmt.Println(v5, v23)
        }
        _ = fi25
    }
    _ = len(v2)
    v1 = (v1 % ((v1) | 1))
    var v32 bool = (!(!true))
    v33 := GenS4{F0: 1, F1: (v22 - (v22 + (v22 * v22))), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
