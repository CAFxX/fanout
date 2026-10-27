package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint }
func genFn0(p0 int64, p1 string) int64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{fi6, (fi6 * (^(fi6))), 2147483647, (fi6 * (fi6 >> ((697) & 63)))}
        var v8 string = ("foo bar" + "hello")
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p0)
    p0 -= -105
    var v9 int32 = -134
    _ = p0
    _ = p1
    _ = v9
    return p0
}

func main() {
    var v1 float64 = 1.5
    var v2 uint32 = 65535
    v3 := make(map[int]bool)
    var v4 int32 = 344
    var v5 uint8 = 2
    v10 := GenS1{F0: "foo bar", F1: -474, F2: 18446744073709551615}
    fmt.Println(v4)
    var v11 uint32 = (v2 & (v2 - v2))
    switch (("foo bar" + ("" + ((("" + "\n\t") + "\n\t") + "foo bar"))) + "\n\t") {
    case "hello":
        var v12 int32 = 512
        _ = v12
    case "foo bar":
        switch "\n\t" {
        case "":
            v11 += 2
        case "foo bar":
            v13 := GenS1{F0: "foo bar", F1: 747, F2: 0}
            _ = v13
        default:
            v1 += v1
        }
    case "":
        var v14 int16 = 1
        _ = v14
    }
    _ = len(v3)
    delete(v3, -947)
    _ = len(v3)
    for fi15 := 0; fi15 < 9; fi15++ {
        v1 += -2.25
        var v16 int32 = v4
        _ = fi15
        _ = v16
    }
    delete(v3, 128)
    v17 := []int{int(v11), 42}
    var v18 float32 = -2.25
    switch "" {
    case "hello":
        v1 = ((float64(v18) - v1) + ((1e10 - (v1 + v1)) / v1))
    case "\n\t":
        if len(v17) > 0 {
            v17[0] = -584
        }
    case "":
        if len(v17) > 0 {
            v17[0] = 32767
        }
    }
    if ((v5 ^ (v5 % ((v5) | 1))) < 2) {
        var v19 int16 = 1
        switch "a" {
        case "hello":
            var v20 float32 = float32(float64(v18))
            _ = v20
        default:
            var v21 string = "hello"
            _ = v21
        }
        if false {
            v22 := GenS1{F0: "foo bar", F1: v19, F2: 18446744073709551615}
            for fi23 := 0; fi23 < 5; fi23++ {
                if true {
                    var v24 int8 = int8(v4)
                    _ = v24
                }
                _ = fi23
            }
            v17 = append(v17, 4294967295)
            _ = v22
        } else {
            v18 += 1.5
        }
        _ = v19
    }
    _ = len(v3)
    v25 := GenS1{F0: "hello", F1: 32767, F2: 18446744073709551615}
    var v26 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
