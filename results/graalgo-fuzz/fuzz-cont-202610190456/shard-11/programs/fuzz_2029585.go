package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 string }
func genFn0(p0 int64, p1 uint32, p2 int16) uint64 {
    p1 -= (p1 ^ 65535)
    if false {
        var v5 uint64 = 229
        p2 = p2
        v6 := GenS1{F0: true, F1: true}
        _ = v5
        _ = v6
    }
    p0 = 9223372036854775807
    p0 -= -9223372036854775808
    v7 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    v1 := []string{("" + ((("\n\t" + "\n\t") + "\n\t") + "hello")), "foo bar", "foo bar"}
    var v2 uint32 = 2147483648
    v3 := make(map[string]string)
    var v4 int16 = 637
    v8 := v3["foo bar"]
    switch -2147483648 {
    case -994:
        v4 = (^(v4))
    case 0:
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 int32 = 2147483647
            _ = fi9
            _ = v10
        }
    case 127:
        if true {
            v2 -= (v2 & 239)
            var v11 int16 = -455
            _ = v11
        } else {
            switch 127 {
            case 127:
                fmt.Println(v8)
            default:
                for fi12 := 0; fi12 < 16; fi12++ {
                    v2 += v2
                    for fi13 := 0; fi13 < 16; fi13++ {
                        switch (fi12 * 1) {
                        case 9223372036854775807:
                            v8 = (v8 + ((v8 + "foo bar") + (v8 + "\n\t")))
                        }
                        _ = fi13
                    }
                    _ = fi12
                }
            }
            v4 = v4
            v8 = (v8 + ("foo bar" + v8))
        }
    }
    var v14 bool = false
    v15 := GenS2{F0: v8}
    v8 = (v8 + (("" + (v8 + "")) + (v8 + "\n\t")))
    v14 = true
    fmt.Println(v14, v4)
    var v16 int32 = 289
    v8 = ""
    v2 = v2
    var v17 int16 = int16(v16)
    fmt.Println(v4, v14)
    v8 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
