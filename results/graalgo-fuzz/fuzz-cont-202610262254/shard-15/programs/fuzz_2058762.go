package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int64; F2 int64 }
type GenS3 struct { F0 int16; F1 bool; F2 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 bool, p1 uint16) int8 {
    v5 := GenS1{F0: (840 <= 4294967295)}
    var v6 uint64 = 708
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -128
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]string)
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]uint32)
    v2["\n\t"] = ("hello" + "foo bar")
    var v7 uint8 = 127
    if true {
        v8 := GenS2{F0: 802, F1: 9223372036854775807, F2: 42}
        var v9 float32 = (1e10 + (0.1 * ((float32(0.0) / 0.1) / 1.5)))
        _ = v8
        _ = v9
    }
    var v10 int8 = -128
    v7 += 199
    v11 := v2[""]
    v11 = v11
    v12 := GenS3{F0: 1000, F1: (2 <= 2), F2: 1}
    if false {
        v13 := []int{}
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v15 := GenS4{F0: -2147483648}
            if true {
                v11 = ((v11 + ("foo bar" + v11)) + "")
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        v10 = int8(v3)
        _ = len(v4)
    }
    v16 := []int{32768, 42, 0, -1}
    v10 = (v10 - (v10 & 1))
    v10 = int8(v7)
    v17 := v4[0]
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v2[((v11 + v11) + "hello")] = v11
        var v19 float64 = (v1 + v1)
        var v20 float32 = (1e10 - ((float32((v1 - (v19 - float64(1.5)))) + -2.25) / 3.14159))
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
