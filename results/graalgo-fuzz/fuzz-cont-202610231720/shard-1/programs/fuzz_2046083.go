package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    v1 := []uint32{4294967295, 1, 65535, 4294967295}
    v2 := make(map[uint32]bool)
    var v3 float32 = (((((-2.25 / 3.14159) - 0.1) / (float32((3.14159 / (1.5 + float64(3.14159)))) * float32(((1.5 + 3.14159) / float64(1.5))))) / float32((1e10 * ((1.5 - 3.14159) * (((float64(1.5) / float64(1.5)) + (-2.25 / float64(1e10))) * (0.0 - ((3.14159 * 0.1) / 1.0))))))) * float32(((-2.25 - 0.0) - 0.0)))
    v4 := make(map[string]uint32)
    var v5 uint32 = 0
    var v6 float32 = v3
    if (42 == 757) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v4["a"] = (v5 - v5)
            _ = fi7
        }
        var v8 string = ("hello" + ("hello" + ("\n\t" + "")))
        var v9 int32 = -1
        _ = v8
        _ = v9
    } else {
        var v10 uint64 = 4294967295
        _ = v10
    }
    var v11 bool = (42 > 42)
    delete(v4, ("a" + "\n\t"))
    v11 = v11
    switch "hello" {
    case "hello":
        v6 = 0.0
    case "\n\t":
        _ = len(v2)
    case "foo bar":
        if (v11 && false) {
            v12 := GenS1{F0: -1, F1: v5}
            _ = v12
        } else {
            v3 = v6
            v1 = append(v1, 0)
        }
    default:
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 uint = 18446744073709551615
            v6 -= float32(1e10)
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 float32 = v3
            v6 -= v6
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: 38, F1: (^(v5))}
        _ = fi15
        _ = v18
    }
    v6 = 1.5
    fmt.Println(v6, v3, v3)
    var v19 uint32 = (v5 >> ((2) & 63))
    for fi20 := 0; fi20 < 14; fi20++ {
        fmt.Println(v5, fi20)
        var v21 float32 = (-2.25 / (v6 / v6))
        if (-355 > -2147483648) {
            v21 = (v21 / float32((((float64(v6) / 1.5) - 0.0) / float64(v6))))
            switch "hello" {
            case "\n\t":
                var v22 uint = 42
                _ = v22
            case "foo bar":
                v3 = 3.14159
            }
        } else {
            for fi23 := 0; fi23 < 17; fi23++ {
                v3 = 1e10
                v2[v19] = (true || (!((!(42 >= -32768)) && (1 < 1))))
                _ = fi23
            }
            if v11 {
                _ = len(v1)
            } else {
                if v11 {
                    fmt.Println(v21, fi20, v5)
                    var v24 bool = v11
                    _ = v24
                }
            }
        }
        _ = fi20
        _ = v21
    }
    v25 := GenS1{F0: int8(v5), F1: v19}
    v1 = append(v1, v5)
    for fi26 := 0; fi26 < 1; fi26++ {
        v19 = v5
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 int8 = 100
            switch "foo bar" {
            case "foo bar":
                v5 += (v19 ^ v5)
            }
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
