package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint32; F2 int64 }
func genFn0(p0 uint64, p1 int32) bool {
    p1 = p1
    v7 := GenS1{F0: "a", F1: 437, F2: 202}
    var v8 int8 = 0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}

func main() {
    var v1 uint8 = 128
    var v2 uint32 = 65535
    v3 := []uint32{(v2 << ((2) & 63)), v2, (v2 - v2), v2}
    v4 := make(map[uint32]string)
    v5 := []uint32{(v2 - v2), v2}
    v6 := make(map[int]int)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        if true {
            v9 := GenS1{F0: "\n\t", F1: 188, F2: v2}
            _ = v9
        } else {
            v1 += (v1 >> ((255) & 63))
            v1 -= 128
            var v10 int = 2147483647
            _ = v10
        }
        v11 := GenS1{F0: ("hello" + ""), F1: v2, F2: v2}
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 float32 = ((((1e10 * 0.1) / 1e10) + float32(((1e10 - 3.14159) / 0.1))) - 0.0)
        delete(v6, (fi12 + fi12))
        v3 = append(v3, v2)
        _ = fi12
        _ = v13
    }
    if (!("\n\t" > "a")) {
        v14 := GenS1{F0: "a", F1: (v2 | v2), F2: 2}
        v1 += v1
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v1 = v1
        v16 := GenS1{F0: "a", F1: (^(v2)), F2: v2}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        switch ("hello" + ("" + "a")) {
        case "":
            var v18 int = 830
            _ = v18
        case "foo bar":
            v1 -= v1
        case "\n\t":
            delete(v6, fi17)
        }
        _ = fi17
    }
    v2 -= 2147483647
    v19 := GenS2{F0: 158}
    for fi20 := 0; fi20 < 16; fi20++ {
        v21 := GenS3{F0: 65535, F1: v2, F2: 1}
        _ = len(v4)
        var v22 int8 = -41
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println(v1, v2)
    var v23 string = "\n\t"
    v24 := v4[v2]
    var v25 float32 = 0.0
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
