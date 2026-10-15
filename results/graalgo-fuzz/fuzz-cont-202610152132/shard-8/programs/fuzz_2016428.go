package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
func genFn0(p0 uint8, p1 uint) int8 {
    var v5 uint32 = 0
    p0 -= (p0 | p0)
    var v6 uint32 = (v5 ^ (v5 * (v5 / (((v5 % (((v5 >> (((p1 * p1)) & 63))) | 1))) | 1))))
    v7 := []int{32767, -829, 32768}
    var v8 string = "foo bar"
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return int8(p0)
}
func genFn1(p0 int32) uint16 {
    p0 -= p0
    var v9 int32 = 0
    v10 := GenS1{F0: (p0 & p0)}
    _ = p0
    _ = v10
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]bool)
    v3 := []string{}
    var v4 float32 = (0.1 - (0.1 / 1e10))
    v11 := GenS2{F0: 42}
    if true {
        var v12 int64 = -1
        _ = v12
    }
    v13 := GenS3{F0: ("a" + ("a" + "foo bar"))}
    v4 -= v4
    v14 := GenS1{F0: -583}
    var v15 int64 = 2
    switch ("hello" + "foo bar") {
    case "\n\t":
        if ("" != "a") {
            _ = len(v3)
            v16 := GenS2{F0: 28}
            v4 += v4
            _ = v16
        }
    case "a":
        for fi17 := 0; fi17 < 6; fi17++ {
            switch fi17 {
            case 127:
                var v18 int16 = -682
                _ = v18
            case -2147483648:
                switch "a" {
                case "hello":
                    for fi19 := 0; fi19 < 4; fi19++ {
                        _ = len(v3)
                        _ = fi19
                    }
                case "a":
                    v20 := GenS1{F0: 0}
                    _ = v20
                }
            default:
                var v21 uint64 = 2
                _ = v21
            }
            v15 -= v15
            _ = fi17
        }
    default:
        var v22 float32 = (-2.25 + 3.14159)
        _ = v22
    }
    v15 = v15
    for fi23 := 0; fi23 < 16; fi23++ {
        if true {
            for fi24 := 0; fi24 < 6; fi24++ {
                v15 -= (v15 % ((v15) | 1))
                _ = fi24
            }
        } else {
            fmt.Println(v15, v15)
            v4 -= (1e10 / (v4 - (v4 * ((v4 - float32(0.1)) - float32(0.0)))))
            v15 = (-(v15))
        }
        delete(v1, (fi23 | fi23))
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
