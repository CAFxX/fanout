package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8; F1 uint; F2 int16 }

func main() {
    var v1 int16 = -300
    var v2 uint64 = 42
    var v3 int = 884
    v4 := []uint32{2}
    v5 := make(map[uint32]bool)
    var v6 uint64 = v2
    if len(v4) > 0 {
        v4[0] = 575
    }
    v7 := GenS1{F0: "foo bar", F1: (0 < -1)}
    if false {
        var v8 string = "foo bar"
        _ = v8
    } else {
        var v9 uint16 = 2
        v10 := GenS1{F0: "\n\t", F1: false}
        var v11 bool = (uint32(v9) != 4294967295)
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(v1, v1, v1)
    switch v3 {
    case 2147483647:
        if (!false) {
            v12 := GenS1{F0: "foo bar", F1: false}
            var v13 string = (((("hello" + "hello") + "foo bar") + (("a" + ("a" + "a")) + "hello")) + "a")
            _ = v12
            _ = v13
        } else {
            if false {
                fmt.Println(v2, v2)
            } else {
                var v14 bool = (-2.25 >= float32(0.1))
                _ = v14
            }
        }
    }
    if (-94 == 127) {
        switch "hello" {
        case "hello":
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 uint8 = 255
                var v17 uint32 = 4294967295
                v18 := GenS1{F0: "\n\t", F1: false}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        default:
            if (284 == 4294967295) {
                v5[42] = (((v3 + v3) >= (-(v3))) && ((!((0.1 / 1.0) > 0.0)) || true))
                v6 -= v2
            }
        }
    }
    v19 := GenS1{F0: "a", F1: true}
    v4 = append(v4, 68)
    var v20 int16 = 1
    if (v6 != (v6 / ((v6) | 1))) {
        v21 := GenS2{F0: 18, F1: 2, F2: 0}
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            v3 = (^(v3))
            _ = fi23
        }
        if len(v4) > 0 {
            v4[0] = 0
        }
        _ = fi22
    }
    fmt.Println(v20, v20, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
