package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int64 = -728
    v2 := []uint32{1, 0, 104}
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v5 int64 = (v1 + -278)
    for fi6 := 0; fi6 < 2; fi6++ {
        v5 += v5
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        var v7 int = -2147483648
        _ = fi6
        _ = v7
    }
    switch ("foo bar" + "") {
    case "":
        v8 := v4["a"]
        _ = v8
    case "a":
        var v9 uint32 = 2
        _ = v9
    case "\n\t":
        if true {
            v10 := GenS1{F0: 141, F1: 1}
            var v11 uint = 65535
            v1 += 1
            _ = v10
            _ = v11
        } else {
            v5 = v1
            v5 = 110
            var v12 uint8 = 2
            _ = v12
        }
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v13 uint = 794
    v14 := GenS1{F0: 4294967295, F1: 72}
    switch ((("hello" + "foo bar") + "hello") + "") {
    case "foo bar":
        var v15 string = ""
        _ = v15
    case "a":
        v3[2] = 2147483647
    }
    v16 := v4[((("" + "") + "a") + "")]
    for fi17 := 0; fi17 < 14; fi17++ {
        v5 = -684
        var v18 uint = 611
        _ = fi17
        _ = v18
    }
    if true {
        for fi19 := 0; fi19 < 1; fi19++ {
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 float64 = -2.25
                if true {
                    var v22 uint32 = (^(v16))
                    var v23 int64 = (v5 ^ (v1 ^ (v1 << ((v13) & 63))))
                    _ = v22
                    _ = v23
                }
                var v24 int8 = 0
                _ = fi20
                _ = v21
                _ = v24
            }
            v13 -= 943
            _ = fi19
        }
        v2 = append(v2, 860)
    }
    v5 = (v5 / ((-9223372036854775808) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
