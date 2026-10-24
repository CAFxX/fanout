package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 string; F2 uint8 }
func genFn0(p0 int16, p1 string) int32 {
    v6 := []int{}
    p1 = ("a" + "")
    v7 := GenS1{F0: (p0 / ((p0) | 1))}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint = 1
    var v2 uint16 = 1
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    v5 := []int{}
    for fi8 := 0; fi8 < 20; fi8++ {
        v2 += 255
        _ = fi8
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v1 = 174
    var v9 float32 = (3.14159 + (1e10 - 0.1))
    delete(v3, 0)
    fmt.Println(v1, v1)
    v10 := GenS1{F0: 1}
    v1 = v1
    if (571 > 695) {
        v9 = v9
        var v11 uint16 = v2
        fmt.Println(v2, v1)
        _ = v11
    } else {
        v1 = (v1 & 65535)
    }
    fmt.Println(v9)
    switch 127 {
    case 2147483647:
        if false {
            v12 := GenS2{F0: -1}
            v2 = v2
            _ = v12
        } else {
            if false {
                for fi13 := 0; fi13 < 14; fi13++ {
                    var v14 float64 = 0.0
                    v9 += 1.5
                    _ = fi13
                    _ = v14
                }
            }
            if len(v5) > 0 {
                v5[0] = -571
            }
            v15 := GenS2{F0: -1}
            _ = v15
        }
    case 42:
        v9 = (v9 / (v9 * v9))
    case 960:
        v16 := GenS2{F0: 1}
        _ = v16
    default:
        for fi17 := 0; fi17 < 13; fi17++ {
            switch "a" {
            case "\n\t":
                if true {
                    v9 = (float32(3.14159) * 0.1)
                    var v18 uint = v1
                    _ = v18
                }
            case "hello":
                v9 -= float32((float64(0.0) * 0.0))
            }
            fmt.Println(v2)
            v19 := GenS1{F0: 243}
            _ = fi17
            _ = v19
        }
    }
    var v20 int = 128
    v20 += (v20 >> ((65535) & 63))
    fmt.Println(v20, v1, v1)
    v21 := GenS3{F0: 9223372036854775807, F1: ("\n\t" + ("foo bar" + "hello")), F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
