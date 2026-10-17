package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 int32 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]string)
    v3 := []uint32{2147483647}
    var v4 int64 = 634
    switch ("a" + "") {
    case "foo bar":
        var v5 string = "a"
        _ = v5
    case "a":
        v3 = append(v3, 65535)
    default:
        v4 -= (-(v4))
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        v1 = (0.1 * ((float32((float64(((v1 / v1) - 3.14159)) + float64(0.1))) / (-2.25 - 0.0)) - ((((0.1 * (float32(float64(v1)) / v1)) + v1) - v1) * 0.1)))
        _ = fi6
    }
    v7 := GenS1{F0: 306, F1: 126, F2: "hello"}
    v8 := GenS1{F0: 42, F1: -1, F2: "foo bar"}
    for fi9 := 0; fi9 < 11; fi9++ {
        if true {
            v1 = 1.5
            v1 = (float32(1.5) - v1)
        }
        _ = fi9
    }
    var v10 float32 = v1
    if (!true) {
        var v11 int64 = (v4 * (v4 ^ (^(v4))))
        _ = v11
    }
    if (1e10 == -2.25) {
        var v12 int64 = v4
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 uint8 = 0
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 float32 = v1
            _ = fi16
            _ = v17
        }
        if false {
            v18 := v2[""]
            v10 += v10
            _ = v18
        } else {
            var v19 uint16 = 255
            v20 := GenS2{F0: -1}
            _ = v19
            _ = v20
        }
        _ = fi15
    }
    v1 = (((v10 / v10) * (1.5 / ((0.0 + ((((1.5 - v1) / 3.14159) - (1.5 - v1)) * v1)) + -2.25))) - v1)
    if len(v3) > 0 {
        v3[0] = 842
    }
    v4 -= v4
    var v21 string = ("hello" + "foo bar")
    switch (v21 + "a") {
    case "foo bar":
        var v22 float64 = 3.14159
        _ = v22
    case "hello":
        v1 = v1
    case "":
        v4 -= (v4 << ((42) & 63))
    default:
        if false {
            v21 = v21
        }
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        fmt.Println(v1)
        var v24 int16 = -1
        var v25 uint32 = 289
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
