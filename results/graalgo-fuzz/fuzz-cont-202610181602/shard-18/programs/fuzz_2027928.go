package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 uint, p1 int64, p2 uint) int64 {
    v7 := []int{2147483647}
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (-(p1))
}

func main() {
    v1 := []int{-331, -721, 42}
    var v2 uint16 = 0
    v3 := []int{32768}
    var v4 uint16 = 350
    var v5 int64 = 42
    var v6 uint16 = (v2 & v2)
    if true {
        v6 -= 1
        var v8 float32 = float32(1.5)
        _ = v8
    }
    fmt.Println(v6, v4, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v9 := GenS1{F0: 800}
        v2 -= v6
        v4 = v2
        _ = v9
    }
    v5 = (v5 ^ v5)
    if true {
        v3 = append(v3, 32768)
        switch "\n\t" {
        case "hello":
            switch "hello" {
            case "a":
                if true {
                    switch 9223372036854775807 {
                    case 128:
                        v3 = append(v3, 0)
                    case -1:
                        v5 = v5
                    default:
                        var v10 uint64 = 2
                        _ = v10
                    }
                    var v11 bool = false
                    _ = v11
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi12 := 0; fi12 < 8; fi12++ {
                        fmt.Println(v2)
                        v6 -= v2
                        v6 += (v4 ^ v2)
                        _ = fi12
                    }
                    v13 := GenS1{F0: -641}
                    _ = v13
                }
            default:
                var v14 uint = 909
                _ = v14
            }
        case "\n\t":
            var v15 int32 = 248
            _ = v15
        default:
            switch -1 {
            case -887:
                switch -2147483648 {
                case -1:
                    for fi16 := 0; fi16 < 10; fi16++ {
                        if (0.0 != (1e10 / 1.0)) {
                            v5 = (v5 + v5)
                            v17 := GenS2{F0: -128, F1: 1}
                            fmt.Println(v6, fi16, v2)
                            _ = v17
                        } else {
                            var v18 int64 = (v5 - (^(v5)))
                            v18 -= v18
                            fmt.Println(v5, v2, v2)
                            _ = v18
                        }
                        fmt.Println(v4, fi16)
                        _ = fi16
                    }
                }
            }
        }
    }
    v4 += 65535
    switch "foo bar" {
    case "\n\t":
        switch ("foo bar" + "") {
        case "a":
            for fi19 := 0; fi19 < 7; fi19++ {
                v4 = v2
                _ = fi19
            }
        case "\n\t":
            _ = len(v3)
        default:
            fmt.Println(v2)
        }
    case "a":
        v3 = append(v3, -461)
    case "hello":
        for fi20 := 0; fi20 < 8; fi20++ {
            v5 = (v5 >> ((0) & 63))
            var v21 uint = 18446744073709551615
            _ = fi20
            _ = v21
        }
    default:
        var v22 int8 = -1
        _ = v22
    }
    var v23 int64 = (v5 << ((65535) & 63))
    var v24 int64 = (v23 - v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
