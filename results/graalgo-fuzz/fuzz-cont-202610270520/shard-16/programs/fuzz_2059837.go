package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 128
    var v3 int16 = 570
    v4 := make(map[uint32]uint32)
    var v5 int = 127
    v6 := []int{32767, int(v3)}
    v3 = (v3 - v3)
    var v7 float32 = (0.1 - 1.5)
    _ = len(v4)
    v3 -= (v3 / ((686) | 1))
    if (false && false) {
        v8 := GenS1{F0: 228, F1: ("hello" + "a"), F2: 255}
        v6 = append(v6, (v5 % ((-505) | 1)))
        _ = v8
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := v1[0]
        _ = fi9
        _ = v10
    }
    if ((0.0 * (0.0 * 1e10)) >= (1e10 + float64(v7))) {
        if true {
            if (true || false) {
                var v11 string = "hello"
                _ = v11
            } else {
                v1[2147483647] = "a"
                var v12 int = (^(v5))
                _ = v12
            }
            for fi13 := 0; fi13 < 13; fi13++ {
                v3 -= v3
                _ = fi13
            }
            if true {
                v14 := GenS1{F0: 1, F1: ("a" + "hello"), F2: 42}
                _ = v14
            } else {
                v5 = (v5 - v5)
            }
        } else {
            var v15 int32 = 2147483647
            v2 -= (v2 - (v2 >> ((992) & 63)))
            _ = v15
        }
    } else {
        v16 := GenS1{F0: 779, F1: "foo bar", F2: 65535}
        v7 += v7
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
