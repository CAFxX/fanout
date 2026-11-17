package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 bool; F1 int8; F2 int }
type GenS4 struct { F0 int; F1 bool; F2 int }
func genFn0(p0 uint64, p1 uint8) int32 {
    fmt.Println(p1, p0)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []string{"foo bar", (("\n\t" + "") + "hello"), "hello", "hello"}
    var v3 int64 = 681
    var v4 int = (v1 << ((18446744073709551615) & 63))
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint64 = 18446744073709551615
        v7 := GenS1{F0: 100, F1: (fi5 / ((-779) | 1)), F2: 2}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 bool = (!false)
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v2 = append(v2, "hello")
        var v10 float32 = ((0.0 - (float32(0.0) * 0.1)) * ((-2.25 + 1e10) - float32(float64(1.5))))
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 422, F1: ((-2.25 + (-2.25 / float64(-2.25))) == 1e10)}
    v12 := GenS3{F0: ((-2.25 / 1.0) < 0.0), F1: 42, F2: (v4 & (v1 ^ -235))}
    switch 32767 {
    case -1:
        if v8 {
            if false {
                v13 := GenS2{F0: 42, F1: true}
                var v14 float32 = 0.0
                v8 = false
                _ = v13
                _ = v14
            } else {
                v4 -= (v4 >> ((4294967295) & 63))
            }
            var v15 uint = 1
            _ = v15
        } else {
            var v16 int16 = 543
            _ = v16
        }
    case 833:
        fmt.Println(v3, v1, v4)
    }
    if ((v8 || v8) || true) {
        var v17 string = ""
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 uint8 = 77
                v21 := GenS1{F0: -112, F1: (-(v1)), F2: 255}
                fmt.Println(fi18, v4)
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        var v22 int64 = -9223372036854775808
        var v23 uint64 = 1
        _ = v22
        _ = v23
    }
    fmt.Println(v8)
    v2 = append(v2, ("a" + ""))
    var v24 string = "a"
    v25 := GenS4{F0: (v4 << ((255) & 63)), F1: (!v8), F2: v4}
    v2 = append(v2, v24)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v8
}
