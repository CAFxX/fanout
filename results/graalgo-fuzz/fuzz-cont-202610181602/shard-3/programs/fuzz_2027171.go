package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{-707, 32767}
    var v3 uint32 = 2147483647
    v4 := []string{"hello", (("hello" + ("foo bar" + ("" + (("\n\t" + (("hello" + "a") + "\n\t")) + "foo bar")))) + ""), (("\n\t" + "a") + ""), ("\n\t" + "")}
    var v5 uint64 = 42
    _ = len(v1)
    v5 -= v5
    var v6 int = 2147483647
    var v7 float32 = 1.5
    if (-128 >= -1) {
        switch (("a" + "a") + ("a" + "")) {
        case "":
            var v8 int8 = 1
            _ = v8
        case "hello":
            v7 = 0.1
        case "a":
            fmt.Println(v5)
        default:
            v9 := GenS1{F0: 42}
            _ = v9
        }
        v3 -= v3
    } else {
        var v10 float32 = v7
        fmt.Println(v6)
        for fi11 := 0; fi11 < 1; fi11++ {
            v10 = v10
            v12 := GenS2{F0: (v6 ^ v6)}
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v13 uint = 18446744073709551615
    v6 = -9223372036854775808
    v14 := GenS3{F0: ("foo bar" + ""), F1: -517, F2: 0}
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint = v13
        _ = fi15
        _ = v16
    }
    var v17 bool = false
    fmt.Println(v6, v7)
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 float64 = (0.1 * (((1e10 - 3.14159) - (((float64(v7) + ((((((-2.25 * (-2.25 + 1.5)) / 0.1) + 3.14159) * (-2.25 + float64(0.1))) / 1.0) + ((float64(3.14159) + -2.25) * 0.1))) * 3.14159) / float64((v7 / v7)))) / 1e10))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v5 += (v5 >> ((42) & 63))
        _ = fi18
        _ = v19
    }
    if false {
        v17 = ((2 != 127) && v17)
        for fi20 := 0; fi20 < 6; fi20++ {
            if (65535 >= v3) {
                _ = len(v1)
            } else {
                var v21 int = v6
                _ = v21
            }
            var v22 int8 = -118
            var v23 float32 = (1e10 * (v7 / 3.14159))
            _ = fi20
            _ = v22
            _ = v23
        }
    }
    var v24 int16 = 542
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
