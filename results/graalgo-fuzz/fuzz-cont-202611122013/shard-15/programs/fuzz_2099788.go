package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS3 struct { F0 string; F1 string; F2 int }

func main() {
    var v1 uint = 0
    v2 := make(map[int]string)
    var v3 int8 = 42
    v4 := []int{-9223372036854775808, -9223372036854775808}
    var v5 int16 = -783
    var v6 float32 = (3.14159 * float32(0.0))
    var v7 int = 0
    for fi8 := 0; fi8 < 6; fi8++ {
        v6 += 0.0
        var v9 int16 = (-(v5))
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: (v7 | int(v1)), F1: "\n\t"}
    if len(v4) > 0 {
        v4[0] = (v7 - (v7 - v7))
    }
    var v11 string = "a"
    switch (((v11 + v11) + ("foo bar" + v11)) + v11) {
    case "hello":
        _ = len(v4)
    case "":
        _ = len(v4)
    default:
        for fi12 := 0; fi12 < 4; fi12++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi12
        }
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS2{F0: ((!((v7 / ((128) | 1)) != 232)) && ((v5 <= v5) && (!true))), F1: 255, F2: 0}
        if true {
            var v15 uint = 255
            v16 := v2[128]
            if (!(255 < 127)) {
                _ = len(v4)
            }
            _ = v15
            _ = v16
        }
        switch fi13 {
        case -666:
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v17 float64 = -2.25
                var v18 int32 = 2147483647
                _ = v17
                _ = v18
            } else {
                v19 := GenS3{F0: "hello", F1: "foo bar", F2: (fi13 - 128)}
                fmt.Println(v3, v3)
                v11 = "foo bar"
                _ = v19
            }
        case 9223372036854775807:
            _ = len(v4)
        case 2147483647:
            v5 = v5
        }
        _ = fi13
        _ = v14
    }
    v20 := GenS2{F0: ((true && ((v5 << ((v1) & 63)) != (v5 * -11))) || (65535 == 2147483647)), F1: 2, F2: 32767}
    if true {
        var v21 int16 = 0
        v7 = v7
        var v22 uint16 = 1
        _ = v21
        _ = v22
    } else {
        var v23 int64 = -1
        _ = v23
    }
    var v24 float32 = 0.1
    v24 = 1e10
    switch ("hello" + v11) {
    case "foo bar":
        var v25 int = (v7 * (v7 / ((2147483647) | 1)))
        _ = v25
    default:
        var v26 int32 = 0
        _ = v26
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
