package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0(p0 bool, p1 uint64) uint16 {
    switch 102 {
    case 4294967295:
        v6 := []int{315, -254, 128}
        _ = v6
    }
    p0 = (903 < 971)
    p0 = (!(-9223372036854775808 <= -107))
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 uint16 = 567
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    v4 := []int{}
    v5 := []int{0}
    switch -9223372036854775808 {
    case 9223372036854775807:
        if true {
            switch 111 {
            case 1:
                if true {
                    var v7 float64 = (1.5 / 1.0)
                    fmt.Println(v7, v1)
                    _ = v7
                } else {
                    v1 = 255
                    var v8 uint8 = 2
                    var v9 string = "\n\t"
                    _ = v8
                    _ = v9
                }
            }
            var v10 float32 = 1e10
            switch 42 {
            case -2147483648:
                fmt.Println(v1)
            default:
                v1 = genFn0(("" == (("a" + "foo bar") + "foo bar")), 497)
            }
            _ = v10
        } else {
            var v11 string = "hello"
            fmt.Println(v1, v11)
            var v12 int8 = 42
            _ = v11
            _ = v12
        }
    }
    switch 32768 {
    case 128:
        switch 0 {
        case -46:
            v4 = append(v4, 9223372036854775807)
        case 32767:
            switch 1 {
            case -4:
                v13 := v3[128]
                _ = v13
            case 2147483647:
                var v14 string = "\n\t"
                _ = v14
            }
        default:
            if len(v4) > 0 {
                v4[0] = -395
            }
        }
    default:
        v1 += v1
    }
    v1 = 455
    fmt.Println(v1)
    v3[0] = ""
    v1 = v1
    v1 = v1
    v15 := GenS1{F0: int8(v1), F1: 58, F2: 920}
    v1 += genFn0(true, 0)
    v16 := GenS2{F0: 1, F1: -32768}
    if false {
        var v17 float32 = (3.14159 + 1e10)
        var v18 uint16 = v1
        _ = v17
        _ = v18
    } else {
        v1 += (v1 ^ v1)
        if false {
            fmt.Println(v1)
        }
    }
    v1 += 2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v19 := GenS1{F0: -1, F1: 255, F2: 280}
    v20 := GenS1{F0: 100, F1: 255, F2: 42}
    switch -774 {
    case 1:
        v21 := GenS2{F0: 42, F1: 0}
        _ = v21
    case -149:
        fmt.Println(v1)
    default:
        var v22 float64 = 0.0
        _ = v22
    }
    if (true || false) {
        fmt.Println(v1)
        var v23 bool = true
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
