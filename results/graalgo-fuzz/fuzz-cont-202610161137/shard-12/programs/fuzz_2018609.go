package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 float32 = (1.5 / 1e10)
    var v3 int = 963
    var v4 uint8 = 18
    var v5 uint16 = 65535
    v6 := GenS1{F0: 124, F1: 321, F2: (v3 * -9223372036854775808)}
    v1[(v3 ^ v3)] = v3
    var v7 int32 = 1
    v8 := []int{-1, v3, (v3 - v3)}
    var v9 float64 = 0.1
    if true {
        v7 = -1
        v10 := GenS2{F0: 32767}
        _ = v10
    }
    if ((v3 % (((v3 * (v3 | v3))) | 1)) != int(v5)) {
        if false {
            v11 := GenS3{F0: 42}
            _ = v11
        } else {
            fmt.Println(v9, v9)
            v7 = (v7 ^ -1)
        }
        v12 := GenS4{F0: (v5 * v5), F1: -225}
        v7 -= -119
        _ = v12
    } else {
        var v13 float32 = (((v2 + float32((float64(v2) * v9))) + v2) + v2)
        switch "hello" {
        case "\n\t":
            v7 = v7
        case "a":
            for fi14 := 0; fi14 < 7; fi14++ {
                v1[(v3 - v3)] = fi14
                _ = fi14
            }
        case "foo bar":
            var v15 string = "\n\t"
            _ = v15
        default:
            var v16 int = 216
            _ = v16
        }
        delete(v1, (v3 >> ((2) & 63)))
        _ = v13
    }
    var v17 int64 = -1
    if (true && true) {
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 9; fi19++ {
                delete(v1, -9223372036854775808)
                _ = fi19
            }
            _ = fi18
        }
    }
    var v20 int = 168
    fmt.Println(v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
