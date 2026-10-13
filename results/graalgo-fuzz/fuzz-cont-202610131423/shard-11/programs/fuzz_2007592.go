package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64, p1 string) uint32 {
    p1 = ("" + ("" + p1))
    v6 := GenS1{F0: 2, F1: 318, F2: 65535}
    switch "\n\t" {
    case "hello":
        v7 := []int{32768, 42}
        _ = v7
    case "":
        var v8 int = -2147483648
        _ = v8
    case "\n\t":
        v9 := GenS2{F0: (p0 - (p0 % (((^(p0))) | 1)))}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 uint8 = 201
    v2 := make(map[uint32]bool)
    var v3 int64 = -1
    var v4 uint8 = v1
    v5 := make(map[uint32]string)
    var v10 uint64 = 42
    if (2 == 756) {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint32 = genFn0((^(v10)), ("foo bar" + "a"))
            switch "a" {
            case "foo bar":
                v3 -= (v3 / (((-(v3))) | 1))
            case "hello":
                v12 = v12
            default:
                v10 = uint64(v12)
            }
            switch "hello" {
            case "foo bar":
                var v13 uint = uint(v10)
                _ = v13
            default:
                var v14 uint32 = 0
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    v4 = v1
    var v15 int16 = -32768
    for fi16 := 0; fi16 < 10; fi16++ {
        v5[uint32(v4)] = "\n\t"
        switch "" {
        case "\n\t":
            if (25 == 66) {
                fmt.Println(fi16, v15)
            } else {
                v17 := GenS3{F0: 613}
                v4 -= 242
                _ = v17
            }
        }
        _ = fi16
    }
    var v18 int16 = 1
    var v19 int32 = 1
    var v20 int8 = -128
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 bool = (!false)
        v10 = (v10 & (v10 % ((v10) | 1)))
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
