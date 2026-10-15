package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int }
func genFn0(p0 int8, p1 uint16) int8 {
    var v4 uint64 = 42
    p1 = 0
    v5 := GenS1{F0: int32(p0), F1: p1, F2: 42}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 string, p1 uint8) int8 {
    if false {
        for fi6 := 0; fi6 < 10; fi6++ {
            p0 = p0
            fmt.Println(p0, p1)
            _ = fi6
        }
        var v7 float64 = float64(1e10)
        var v8 uint = 2
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 8; fi9++ {
            switch fi9 {
            case -793:
                fmt.Println(p1)
            case -9223372036854775808:
                p1 -= (^(p1))
            case 4294967295:
                var v10 string = "a"
                _ = v10
            default:
                p0 = p0
            }
            v11 := GenS1{F0: 1, F1: 685, F2: (fi9 & fi9)}
            switch ((p0 + p0) + p0) {
            case "foo bar":
                var v12 uint16 = 313
                _ = v12
            }
            _ = fi9
            _ = v11
        }
        p0 = p0
    }
    p1 += (p1 * (p1 % ((p1) | 1)))
    if true {
        fmt.Println(p1, p1)
    }
    v13 := []int{-499, -80, 42}
    _ = p0
    _ = p1
    _ = v13
    return 0
}

func main() {
    var v1 float64 = (3.14159 * (0.1 - 0.0))
    v2 := make(map[uint32]string)
    var v3 string = "foo bar"
    _ = len(v2)
    v3 = "foo bar"
    v1 -= v1
    var v14 int8 = 100
    for fi15 := 0; fi15 < 14; fi15++ {
        fmt.Println(v3)
        _ = fi15
    }
    switch 9223372036854775807 {
    case 446:
        v16 := GenS1{F0: 160, F1: 0, F2: 42}
        _ = v16
    default:
        if false {
            v3 = v3
        } else {
            fmt.Println(v1, v1)
        }
    }
    fmt.Println(v1)
    v17 := []int{}
    if (!(float32((1e10 * v1)) <= 3.14159)) {
        v14 -= 1
        if false {
            switch 0 {
            case -2147483648:
                var v18 float64 = v1
                _ = v18
            case 32767:
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
            case 9223372036854775807:
                for fi19 := 0; fi19 < 20; fi19++ {
                    v14 += v14
                    _ = fi19
                }
            }
        }
        var v20 int8 = (v14 & (v14 + v14))
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS1{F0: -2147483648, F1: 1, F2: 42}
        v3 = ""
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
}
