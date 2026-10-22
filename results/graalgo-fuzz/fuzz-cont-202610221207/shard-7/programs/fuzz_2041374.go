package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 string, p1 bool, p2 uint16) int8 {
    fmt.Println(p1, p2, p0)
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 56
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 54
    var v3 float64 = (1.5 - 0.0)
    var v4 int16 = 1000
    v2 -= 9
    v7 := []int{-9223372036854775808, 32767}
    for fi8 := 0; fi8 < 4; fi8++ {
        v2 = genFn0("foo bar", true, 649)
        _ = fi8
    }
    var v9 int32 = -1
    switch 128 {
    case 0:
        switch "a" {
        case "":
            var v10 int64 = 1
            _ = v10
        }
    case 2147483647:
        switch -9223372036854775808 {
        case 562:
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 float64 = (0.0 / (3.14159 / float64(3.14159)))
                v13 := GenS1{F0: 2, F1: 255, F2: 784}
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            var v14 string = "foo bar"
            _ = v14
        }
    }
    v1[-1] = "hello"
    var v15 uint32 = 65535
    if true {
        var v16 int64 = 867
        v2 = (v2 << ((65535) & 63))
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            for fi18 := 0; fi18 < 8; fi18++ {
                var v19 int64 = -1
                v19 = -1
                v20 := GenS2{F0: -1, F1: 0}
                v4 = (v4 % ((42) | 1))
                _ = fi18
                _ = v19
                _ = v20
            }
            v7 = append(v7, 808)
            _ = fi17
        }
        var v21 int = int(v2)
        _ = v21
    }
    v4 += 1000
    v4 += (v4 ^ 0)
    for fi22 := 0; fi22 < 2; fi22++ {
        _ = len(v7)
        if len(v7) > 0 {
            v7[0] = (fi22 | (fi22 << ((55) & 63)))
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
