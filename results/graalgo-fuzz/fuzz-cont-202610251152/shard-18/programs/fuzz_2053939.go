package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int, p1 uint16) uint32 {
    v4 := GenS1{F0: 92}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v4
    return uint32(p1)
}

func main() {
    var v1 bool = false
    var v2 string = "hello"
    v3 := []uint32{290}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1, v2)
    v1 = v1
    var v5 uint = 42
    if (2 <= 0) {
        if (51 != 0) {
            var v6 uint8 = 102
            fmt.Println(v1)
            _ = v6
        } else {
            var v7 float32 = (1.5 - (float32((float64(float32(-2.25)) / (1.5 / 1.5))) + float32(0.1)))
            _ = v7
        }
        var v8 string = "hello"
        _ = v8
    } else {
        _ = len(v3)
        v2 = (v2 + "foo bar")
        v3 = append(v3, 0)
    }
    var v9 uint64 = 322
    v1 = v1
    if v1 {
        v10 := GenS2{F0: -128}
        _ = v10
    }
    switch ("\n\t" + ("foo bar" + v2)) {
    case "hello":
        var v11 string = "a"
        _ = v11
    case "":
        fmt.Println(v9, v2)
    default:
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 int32 = 42
                var v15 bool = (v1 || (!(2147483648 <= 191)))
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
    }
    v3 = append(v3, 2147483647)
    if len(v3) > 0 {
        v3[0] = 120
    }
    if (!(93 < 1)) {
        var v16 uint = (v5 / (((v5 % ((2) | 1))) | 1))
        _ = v16
    } else {
        var v17 uint64 = 4294967295
        var v18 int8 = 52
        v9 -= v17
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        if (!(!((0.0 / ((3.14159 * 0.1) / (0.1 + 1e10))) <= 0.0))) {
            if len(v3) > 0 {
                v3[0] = 0
            }
        } else {
            fmt.Println(fi19, v2)
        }
        _ = fi19
    }
    var v20 string = "hello"
    var v21 uint = (v5 | 2)
    if len(v3) > 0 {
        v3[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v5
    _ = v9
}
