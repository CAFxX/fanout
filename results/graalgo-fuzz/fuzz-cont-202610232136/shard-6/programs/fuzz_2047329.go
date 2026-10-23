package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 string; F1 int32; F2 uint64 }
func genFn0(p0 int8, p1 bool, p2 uint32) uint {
    var v5 bool = p1
    if (!v5) {
        var v6 float32 = 0.0
        var v7 uint32 = (p2 % (((p2 | (p2 >> ((898) & 63)))) | 1))
        _ = v6
        _ = v7
    }
    var v8 int8 = p0
    v5 = v5
    v9 := GenS1{F0: (^(p0)), F1: -1, F2: -170}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint, p1 int32, p2 uint8) int64 {
    v10 := []int{9223372036854775807, 501}
    switch 1 {
    case -895:
        v11 := GenS1{F0: 0, F1: 32767, F2: (p1 | (p1 & 42))}
        _ = v11
    case 42:
        p2 = 110
    case -1:
        p1 -= (p1 ^ p1)
    }
    var v12 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v12
    return 9223372036854775807
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]uint32)
    var v3 int32 = -716
    v4 := make(map[uint32]string)
    v4[65535] = "foo bar"
    for fi13 := 0; fi13 < 6; fi13++ {
        v3 -= (v3 + 2147483647)
        _ = fi13
    }
    if true {
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int16 = 32767
            _ = fi14
            _ = v15
        }
        v16 := v2["hello"]
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int16 = 1
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: -1, F1: 1, F2: v3}
    fmt.Println(v1)
    v20 := GenS2{F0: 18446744073709551615}
    v21 := GenS1{F0: 0, F1: 1, F2: v3}
    v22 := GenS3{F0: 0, F1: 32767}
    var v23 int8 = -1
    v24 := GenS4{F0: "", F1: v3, F2: 42}
    v23 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
