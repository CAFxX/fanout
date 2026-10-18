package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int64, p1 int64) int64 {
    p0 -= (p0 ^ 9223372036854775807)
    v7 := GenS1{F0: ("hello" + ("foo bar" + "foo bar")), F1: 618, F2: 512}
    v8 := []int{4294967295, 2147483647, -424}
    if ((3.14159 - 3.14159) == float32((3.14159 + (-2.25 * 3.14159)))) {
        p1 -= p1
    } else {
        p0 = (p1 | 42)
        v9 := GenS1{F0: "\n\t", F1: 32767, F2: -605}
        _ = v9
    }
    switch -46 {
    case -274:
        p1 = (p0 / ((276) | 1))
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 bool, p1 uint8, p2 int8) uint64 {
    var v10 uint64 = 157
    v11 := GenS2{F0: 42, F1: 649, F2: 42}
    p0 = p0
    var v12 uint64 = (v10 - v10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 238
}

func main() {
    var v1 int32 = 690
    var v2 bool = ((3.14159 - 0.0) == 1.5)
    var v3 uint8 = 127
    v4 := make(map[string]bool)
    var v5 int32 = v1
    var v6 float32 = (float32(1e10) * -2.25)
    v2 = ((1 >= 0) || ((v5 * v1) == 0))
    v13 := []int{}
    for fi14 := 0; fi14 < 18; fi14++ {
        _ = len(v4)
        var v15 int8 = 79
        fmt.Println(v5, v3)
        _ = fi14
        _ = v15
    }
    v3 += (^(v3))
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := GenS2{F0: 32767, F1: 4294967295, F2: fi17}
            var v19 float32 = (v6 * (v6 - (float32(-2.25) * ((1e10 - float32(0.0)) - v6))))
            _ = len(v13)
            _ = fi17
            _ = v18
            _ = v19
        }
        fmt.Println(v5, v3)
        v20 := GenS1{F0: ("" + ("a" + "foo bar")), F1: 80, F2: 1}
        _ = fi16
        _ = v20
    }
    var v21 int = 841
    var v22 uint64 = 1
    var v23 uint16 = 65535
    v5 = -894
    v24 := GenS3{F0: 42, F1: v22}
    for fi25 := 0; fi25 < 5; fi25++ {
        delete(v4, "foo bar")
        var v26 bool = false
        v27 := GenS2{F0: v23, F1: 2, F2: (v21 ^ -289)}
        _ = fi25
        _ = v26
        _ = v27
    }
    fmt.Println(v22, v3)
    v28 := v4[""]
    switch v21 {
    case 997:
        var v29 int = v21
        _ = v29
    case -1:
        var v30 uint8 = v3
        _ = v30
    }
    v3 = (v3 << ((663) & 63))
    v31 := GenS3{F0: 42, F1: v22}
    var v32 bool = v28
    v4[("\n\t" + "foo bar")] = v28
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
