package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
func genFn0() string {
    var v4 uint32 = 410
    v5 := []int{42, 128}
    v5 = append(v5, -2147483648)
    v4 += 1
    _ = v4
    _ = v5
    return "foo bar"
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]string)
    var v3 uint16 = 936
    v6 := GenS1{F0: 9223372036854775807}
    v7 := GenS1{F0: 42}
    v1 += v1
    v1 = 235
    v8 := GenS1{F0: 2147483647}
    for fi9 := 0; fi9 < 5; fi9++ {
        fmt.Println(fi9, v1)
        fmt.Println(fi9)
        var v10 uint16 = 527
        _ = fi9
        _ = v10
    }
    if (!true) {
        switch "" {
        case "":
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 string = (((((("a" + "\n\t") + "foo bar") + "") + "a") + "hello") + "")
                fmt.Println(v12, v1)
                _ = fi11
                _ = v12
            }
        case "hello":
            var v13 string = "hello"
            _ = v13
        case "foo bar":
            var v14 uint8 = 156
            _ = v14
        }
    }
    v15 := GenS2{F0: true, F1: 0, F2: v3}
    if false {
        v3 = v3
    } else {
        fmt.Println(v3, v1)
        v1 = v1
        fmt.Println(v1, v1)
    }
    v16 := GenS2{F0: true, F1: -992, F2: (v3 | (v3 >> ((18446744073709551615) & 63)))}
    v1 -= (^(v1))
    var v17 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
