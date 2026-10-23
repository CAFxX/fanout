package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 42
    var v3 int32 = -51
    var v4 bool = true
    v5 := []string{"\n\t", "hello"}
    v6 := make(map[string]uint32)
    fmt.Println(v3)
    v6["hello"] = 2147483647
    if (95 <= -258) {
        switch 32767 {
        case 2147483647:
            var v7 uint32 = 4294967295
            _ = v7
        case 127:
            var v8 int8 = 5
            _ = v8
        case -534:
            v5 = append(v5, "a")
        default:
            var v9 float32 = 0.1
            _ = v9
        }
        fmt.Println(v3, v3)
    } else {
        var v10 int64 = v2
        _ = v10
    }
    switch 207 {
    case -370:
        if (((1.5 - -2.25) * ((1e10 - 1e10) + (float32(1.5) + 1e10))) <= ((-2.25 / 1.0) / (-2.25 / 0.1))) {
            var v11 int = 9223372036854775807
            _ = v11
        } else {
            switch -1 {
            case 32768:
                var v12 uint32 = 2147483647
                _ = v12
            case -868:
                v2 += v2
            }
            v2 += v2
            v2 += (v2 + (v2 * v2))
        }
    case -388:
        var v13 int8 = 100
        _ = v13
    case 548:
        var v14 float32 = 0.1
        _ = v14
    default:
        v5 = append(v5, "")
    }
    v15 := v1[""]
    v15 = (-(v15))
    var v16 uint32 = 4294967295
    fmt.Println(v2, v3, v4)
    v4 = (true && ((float32(0.1) + (3.14159 - (1e10 - 0.1))) == 0.1))
    if v4 {
        fmt.Println(v16, v4)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v17 := GenS1{F0: -1, F1: v15}
        _ = v17
    } else {
        _ = len(v5)
    }
    v18 := GenS1{F0: 32767, F1: (v15 * -609)}
    for fi19 := 0; fi19 < 13; fi19++ {
        v15 += v15
        var v20 int = fi19
        _ = fi19
        _ = v20
    }
    v2 -= (v2 & 1)
    v4 = v4
    switch "hello" {
    case "a":
        if (128 <= 42) {
            switch -362 {
            case 9223372036854775807:
                for fi21 := 0; fi21 < 5; fi21++ {
                    v5 = append(v5, "hello")
                    _ = fi21
                }
            case 198:
                var v22 string = "foo bar"
                _ = v22
            }
            var v23 uint64 = 683
            var v24 uint16 = 1
            _ = v23
            _ = v24
        }
    case "hello":
        fmt.Println(v15, v16)
    case "foo bar":
        fmt.Println(v16, v4)
    default:
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := GenS1{F0: -1, F1: 1}
            v4 = v4
            var v27 int64 = -301
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    v4 = v4
    for fi28 := 0; fi28 < 5; fi28++ {
        v2 = (v2 * 136)
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
