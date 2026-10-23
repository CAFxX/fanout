package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32; F1 int64 }
func genFn0(p0 uint32, p1 string, p2 uint32) bool {
    p2 = p2
    p1 = (p1 + p1)
    fmt.Println(p1, p0)
    v6 := []int{-425}
    if true {
        var v7 int64 = 9223372036854775807
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (false || false)
}
func genFn1() int {
    if false {
        var v8 uint16 = 0
        v8 = 2
        v9 := GenS1{F0: 551, F1: 2147483647}
        _ = v8
        _ = v9
    }
    var v10 int8 = 109
    v11 := GenS1{F0: 2147483647, F1: 1}
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 int64 = 42
    var v2 float32 = 0.0
    var v3 int8 = -18
    var v4 uint = 18446744073709551615
    var v5 string = ""
    switch 671 {
    case 1:
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 uint64 = 128
            v5 = "a"
            _ = fi12
            _ = v13
        }
    case -707:
        var v14 uint8 = 1
        _ = v14
    default:
        var v15 int = -1
        _ = v15
    }
    v2 = (-2.25 / ((v2 / v2) / 1.5))
    for fi16 := 0; fi16 < 9; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := []int{}
            _ = fi17
            _ = v18
        }
        var v19 int64 = (v1 / (((^(v1))) | 1))
        _ = fi16
        _ = v19
    }
    v20 := []int{190, genFn1()}
    v3 = -23
    _ = len(v20)
    switch 2 {
    case 32768:
        fmt.Println(v1)
    case -449:
        var v21 int = 32768
        _ = v21
    case 0:
        if (!(2147483647 != 4294967295)) {
            v2 = (float32(3.14159) / v2)
            switch 128 {
            case -280:
                var v22 uint16 = 32767
                _ = v22
            case -1:
                var v23 float32 = -2.25
                _ = v23
            case -465:
                v2 = (v2 * float32(float64((1.5 + float32((0.0 + (-2.25 * 0.0)))))))
            default:
                switch "hello" {
                case "hello":
                    v3 = (v3 + 127)
                }
            }
        }
    default:
        for fi24 := 0; fi24 < 16; fi24++ {
            v20 = append(v20, fi24)
            var v25 int16 = 0
            _ = fi24
            _ = v25
        }
    }
    var v26 uint = (^(v4))
    var v27 uint16 = 42
    _ = len(v20)
    switch 32768 {
    case 2147483647:
        v1 = (v1 & -9223372036854775808)
    case -226:
        v28 := GenS2{F0: -2147483648, F1: v1}
        _ = v28
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 int8 = (v3 >> ((65535) & 63))
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = fi29
        _ = v30
    }
    fmt.Println(v1)
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 uint32 = 65535
        _ = fi31
        _ = v32
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
