package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
func genFn0(p0 uint8) int8 {
    v5 := GenS1{F0: 2, F1: 526, F2: (p0 ^ (p0 & (p0 + (^(p0)))))}
    p0 = 128
    var v6 int64 = 0
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{}
        fmt.Println(p0)
        var v9 float32 = -2.25
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = true
    _ = p0
    _ = v10
    _ = v5
    _ = v6
    return int8(p0)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := []string{"foo bar", "", "hello", ""}
    var v4 int64 = 42
    v11 := GenS2{F0: -1}
    v1[27] = 0
    v12 := GenS1{F0: 284, F1: 65535, F2: 2}
    if false {
        for fi13 := 0; fi13 < 20; fi13++ {
            _ = len(v3)
            v4 += 347
            _ = fi13
        }
    }
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    if true {
        v1[1] = 2147483648
        v14 := GenS3{F0: (v4 ^ 43)}
        v4 = (v4 / ((-368) | 1))
        _ = v14
    } else {
        v15 := GenS2{F0: -19}
        delete(v1, 0)
        _ = len(v1)
        _ = v15
    }
    var v16 uint16 = 32767
    v16 = (^(v16))
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := GenS4{F0: (-(fi17))}
        _ = fi17
        _ = v18
    }
    v4 += (v4 >> ((575) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
