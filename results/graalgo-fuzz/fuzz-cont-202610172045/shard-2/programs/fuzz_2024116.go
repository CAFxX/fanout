package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int = 4294967295
    var v2 uint8 = 93
    v3 := []uint32{}
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int64 = 345
        _ = fi4
        _ = v5
    }
    var v6 int16 = -1
    v2 += (v2 - 1)
    var v7 float32 = ((1e10 * -2.25) * (((0.0 + float32(1e10)) + float32(float64((-2.25 / 3.14159)))) / 1.5))
    if (!(3.14159 == v7)) {
        if true {
            var v8 int32 = 42
            v9 := GenS1{F0: "", F1: "hello"}
            if true {
                v2 -= v2
            } else {
                v3 = append(v3, 2)
                v3 = append(v3, 44)
            }
            _ = v8
            _ = v9
        } else {
            fmt.Println(v7, v2, v6)
            _ = len(v3)
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v7 = (((3.14159 / (v7 / (v7 + (((v7 / (3.14159 + 1e10)) - ((v7 + v7) + 0.1)) - (float32(-2.25) - float32(-2.25)))))) - v7) * ((v7 / v7) / ((v7 / (v7 * 1e10)) / v7)))
            _ = fi10
        }
        var v11 bool = ((0 >= 1) && (!(false && false)))
        _ = v11
    }
    switch "hello" {
    case "hello":
        v6 += v6
    default:
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    }
    _ = len(v3)
    if (true || ((42 <= -599) && true)) {
        switch (v1 & 128) {
        case 128:
            v12 := GenS2{F0: 127}
            _ = v12
        default:
            v13 := GenS3{F0: (v1 + v1)}
            _ = v13
        }
        v14 := GenS3{F0: v1}
        var v15 uint8 = v2
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: "", F1: "foo bar"}
    v17 := GenS4{F0: 60}
    for fi18 := 0; fi18 < 14; fi18++ {
        v6 = (v6 | (v6 + -1))
        v19 := GenS1{F0: ("hello" + ("a" + "foo bar")), F1: "a"}
        _ = fi18
        _ = v19
    }
    if true {
        v20 := GenS5{F0: (v6 << ((89) & 63))}
        _ = v20
    }
    v1 -= (v1 << ((2) & 63))
    if len(v3) > 0 {
        v3[0] = 2
    }
    v21 := GenS1{F0: "", F1: "\n\t"}
    fmt.Println(v1)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v7
}
