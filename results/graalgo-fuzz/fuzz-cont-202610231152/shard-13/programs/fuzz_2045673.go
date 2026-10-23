package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint64 }

func main() {
    var v1 float64 = -2.25
    var v2 int64 = 42
    var v3 float32 = (0.0 + ((0.1 / 0.1) * 1.5))
    var v4 uint = 927
    var v5 uint8 = 2
    v6 := make(map[uint32]uint32)
    v7 := v6[uint32(v4)]
    if (true || (!true)) {
        if (false || (-1 != -128)) {
            var v8 uint8 = v5
            var v9 int8 = 42
            fmt.Println(v5)
            _ = v8
            _ = v9
        } else {
            var v10 int8 = -39
            v11 := GenS1{F0: (("\n\t" + ("foo bar" + (("\n\t" + ("a" + ((("hello" + "\n\t") + "") + ("a" + "hello")))) + "\n\t"))) + "hello"), F1: (v5 << ((v4) & 63)), F2: 308}
            _ = v10
            _ = v11
        }
        switch "" {
        case "foo bar":
            v1 += (v1 - (float64(v3) + ((0.1 * (-2.25 / (0.1 * v1))) / ((v1 - v1) * (v1 / (((float64(((v3 * ((0.1 * v3) + -2.25)) + v3)) + 0.0) / float64(1.5)) / (0.1 + float64(v3))))))))
        case "hello":
            v12 := []int{0, -828}
            _ = v12
        case "a":
            v13 := []int{}
            _ = v13
        default:
            fmt.Println(v4, v4)
        }
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int64 = v2
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: (("\n\t" + "a") + "hello"), F1: (^(v5)), F2: 0}
    v17 := GenS1{F0: "hello", F1: (v5 >> (((^(v4))) & 63)), F2: 4294967295}
    var v18 int64 = v2
    var v19 int8 = -128
    fmt.Println(v1, v2)
    v20 := GenS1{F0: ("foo bar" + ("a" + "hello")), F1: v5, F2: 2}
    v4 += (v4 ^ 301)
    for fi21 := 0; fi21 < 14; fi21++ {
        v2 += (v18 - -9223372036854775808)
        v7 -= (v7 << ((v4) & 63))
        v4 += v4
        _ = fi21
    }
    var v22 uint = v4
    v5 = v5
    if true {
        v23 := v6[65535]
        var v24 int32 = 329
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        fmt.Println(v18, v1)
        var v26 bool = (!(false && (!(-1 <= (v19 << (((v22 << (((v22 >> ((368) & 63))) & 63))) & 63))))))
        v26 = (!(!(-32768 == -935)))
        _ = fi25
        _ = v26
    }
    v7 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
