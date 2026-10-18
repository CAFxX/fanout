package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint }
func genFn0(p0 string, p1 bool, p2 uint32) uint {
    fmt.Println(p0)
    var v5 float32 = (((1e10 - 1e10) * 3.14159) - 1e10)
    fmt.Println(v5, p1, v5)
    var v6 float32 = v5
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 string = ("hello" + "foo bar")
    var v2 uint = 1
    v3 := make(map[uint32]uint32)
    var v4 int8 = 100
    v4 = (v4 / ((0) | 1))
    for fi7 := 0; fi7 < 8; fi7++ {
        v3[2147483648] = 0
        _ = fi7
    }
    switch 32768 {
    case -653:
        switch 1 {
        case 1:
            v4 += (v4 ^ (v4 % (((v4 + 47)) | 1)))
        case 128:
            v4 = v4
        default:
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 bool = (!(!false))
                var v10 int8 = v4
                var v11 uint16 = 1
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: 32767, F1: -1}
            switch v1 {
            case "hello":
                v14 := []int{(fi12 << ((v2) & 63)), 649, 32768}
                _ = v14
            case "a":
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 uint = (v2 % ((genFn0(v1, ((1.5 - ((0.1 / 0.1) / 0.1)) <= ((float32(0.1) * (3.14159 + 1.5)) + 0.1)), 42)) | 1))
                    v17 := GenS2{F0: (^(v16))}
                    v18 := GenS1{F0: 0, F1: -1}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            case "":
                if true {
                    v1 = ("hello" + "hello")
                    v4 += v4
                    var v19 uint8 = 42
                    _ = v19
                } else {
                    var v20 int8 = v4
                    _ = v20
                }
            }
            _ = fi12
            _ = v13
        }
    }
    if false {
        var v21 float64 = 1e10
        _ = v21
    }
    switch v1 {
    case "\n\t":
        v2 -= 65535
    case "foo bar":
        fmt.Println(v1, v2, v4)
    default:
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS2{F0: (^(v2))}
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(v4, v2)
    var v24 uint32 = 423
    v2 = genFn0("foo bar", (0.0 < 1e10), (v24 * v24))
    v24 = v24
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
