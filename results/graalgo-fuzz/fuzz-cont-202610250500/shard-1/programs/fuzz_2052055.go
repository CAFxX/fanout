package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 bool; F2 uint32 }
func genFn0(p0 int64) int32 {
    fmt.Println(p0)
    var v5 int16 = -32768
    var v6 uint16 = 0
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int8 = 127
    v2 := []int{}
    v3 := make(map[uint32]string)
    v4 := []uint32{396, 0}
    var v7 bool = (((((1e10 == 1e10) && (539 <= 242)) || false) && false) || (!(false || false)))
    v1 -= -46
    switch "foo bar" {
    case "a":
        if false {
            for fi8 := 0; fi8 < 3; fi8++ {
                switch fi8 {
                case 2147483647:
                    v1 = 100
                case 128:
                    switch 9223372036854775807 {
                    case 42:
                        var v9 uint16 = 255
                        _ = v9
                    case -985:
                        v7 = v7
                    case 646:
                        var v10 uint16 = 42
                        _ = v10
                    default:
                        delete(v3, 4294967295)
                    }
                case -971:
                    var v11 uint64 = 2
                    _ = v11
                default:
                    delete(v3, 42)
                }
                v1 = -27
                var v12 uint32 = 509
                _ = fi8
                _ = v12
            }
            var v13 uint8 = 28
            v14 := v3[47]
            _ = v13
            _ = v14
        } else {
            v7 = v7
            var v15 uint16 = 605
            v7 = (true || false)
            _ = v15
        }
    }
    if false {
        if len(v2) > 0 {
            v2[0] = 32768
        }
        switch ("a" + (("\n\t" + "") + "")) {
        case "":
            v1 = (v1 / (((-(v1))) | 1))
        }
    }
    v16 := v3[2147483647]
    switch "foo bar" {
    case "a":
        var v17 uint = 255
        _ = v17
    default:
        var v18 int16 = -1
        _ = v18
    }
    var v19 float32 = 0.1
    v20 := GenS1{F0: 100, F1: 32767}
    _ = len(v2)
    v1 = 93
    var v21 uint32 = 65535
    var v22 uint = 0
    v23 := GenS1{F0: v1, F1: 32767}
    for fi24 := 0; fi24 < 13; fi24++ {
        v22 = v22
        v25 := GenS2{F0: v22, F1: 255}
        v26 := GenS3{F0: 0, F1: v7, F2: 2147483648}
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println(v7, v19, v21)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
