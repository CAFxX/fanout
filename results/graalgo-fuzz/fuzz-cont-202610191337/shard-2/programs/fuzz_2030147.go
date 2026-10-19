package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint16) uint8 {
    p1 = (p1 >> ((400) & 63))
    p0 = p1
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    var v5 float32 = ((0.1 / float32((float64(0.1) * (-2.25 + 3.14159)))) / 3.14159)
    v1 -= v1
    v5 += (v5 / -2.25)
    v1 -= v1
    v6 := []int{-507, 1}
    _ = len(v6)
    v1 = 1.5
    v5 -= ((v5 - v5) - float32((v1 + float64(0.1))))
    v7 := GenS1{F0: ((-9223372036854775808 != 473) && (true && (("hello" + "") != ("a" + (("" + ((("a" + ("a" + "foo bar")) + "\n\t") + ("a" + "a"))) + "foo bar")))))}
    v1 = v1
    if false {
        v5 += v5
        v5 -= (((v5 / v5) - float32((0.0 + (v1 + ((v1 + v1) * 3.14159))))) / 0.1)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 float32 = (-2.25 / 1e10)
        _ = fi8
        _ = v9
    }
    var v10 int8 = -72
    var v11 uint64 = 0
    if ((2147483647 < 267) || (v11 != 557)) {
        _ = len(v2)
        fmt.Println(v11, v5)
    } else {
        v10 = v10
    }
    var v12 uint16 = 42
    switch 1 {
    case 128:
        for fi13 := 0; fi13 < 4; fi13++ {
            fmt.Println(v5, v10, v5)
            _ = fi13
        }
    case 32767:
        if ((v12 ^ (v12 - 0)) == 65535) {
            _ = len(v2)
            switch 42 {
            case 130:
                if len(v6) > 0 {
                    v6[0] = -361
                }
            case 530:
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 int8 = -80
                    var v16 int16 = 42
                    fmt.Println(v10, v11)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            case 132:
                v17 := GenS2{F0: -472, F1: 0}
                _ = v17
            default:
                v5 = v5
            }
            v11 = v11
        } else {
            v18 := GenS3{F0: 0}
            _ = v18
        }
    case -677:
        var v19 uint8 = 127
        _ = v19
    }
    v6 = append(v6, -678)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
