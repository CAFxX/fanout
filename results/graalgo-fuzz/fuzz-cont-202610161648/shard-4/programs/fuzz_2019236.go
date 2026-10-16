package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 string; F1 uint16; F2 bool }
type GenS3 struct { F0 int; F1 int16; F2 int }
type GenS4 struct { F0 string; F1 int32 }
func genFn0(p0 bool, p1 uint, p2 int) int {
    if p0 {
        v5 := GenS1{F0: (("hello" + "") + (("\n\t" + "hello") + (((("foo bar" + "a") + (("" + "hello") + "hello")) + "a") + "a"))), F1: -1, F2: 680}
        v6 := []int{(^(p2))}
        _ = len(v6)
        _ = v5
        _ = v6
    }
    var v7 uint8 = 240
    v8 := GenS2{F0: (((("\n\t" + "hello") + "a") + "a") + "\n\t"), F1: 828, F2: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p2 / (((p2 << ((326) & 63))) | 1))
}
func genFn1(p0 uint) uint16 {
    var v9 uint8 = uint8(p0)
    v10 := GenS3{F0: -9223372036854775808, F1: 1000, F2: 0}
    p0 -= p0
    _ = p0
    _ = v10
    _ = v9
    return 465
}

func main() {
    var v1 uint16 = 376
    var v2 float64 = 0.1
    var v3 uint32 = 65535
    v4 := make(map[uint32]uint32)
    v11 := []int{}
    v3 -= (v3 << ((894) & 63))
    if false {
        switch -2147483648 {
        case -9223372036854775808:
            var v12 int8 = 0
            _ = v12
        case -268:
            v13 := v4[(v3 ^ 42)]
            _ = v13
        case -868:
            delete(v4, v3)
        }
    }
    fmt.Println(v1)
    var v14 int32 = 685
    for fi15 := 0; fi15 < 4; fi15++ {
        if (514 > 0) {
            var v16 float32 = -2.25
            var v17 string = ("a" + "\n\t")
            _ = v16
            _ = v17
        } else {
            v1 -= v1
        }
        _ = fi15
    }
    if ((!false) && false) {
        for fi18 := 0; fi18 < 13; fi18++ {
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 uint64 = 949
                switch fi18 {
                case -1:
                    var v21 float32 = 0.0
                    _ = v21
                case 9223372036854775807:
                    v2 = v2
                }
                if false {
                    var v22 int8 = 0
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
            v1 -= (v1 >> ((188) & 63))
            switch ((("\n\t" + "foo bar") + "") + ("\n\t" + "")) {
            case "":
                switch 0 {
                case 803:
                    _ = len(v4)
                case -1:
                    _ = len(v4)
                default:
                    v23 := v4[42]
                    _ = v23
                }
            case "foo bar":
                v24 := GenS4{F0: "hello", F1: v14}
                _ = v24
            case "a":
                v14 += v14
            default:
                v25 := GenS4{F0: "a", F1: v14}
                _ = v25
            }
            _ = fi18
        }
    } else {
        var v26 int32 = v14
        fmt.Println(v26, v3, v1)
        _ = v26
    }
    v27 := GenS2{F0: "hello", F1: v1, F2: ((128 == 127) || true)}
    v2 += float64(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
