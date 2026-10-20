package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
func genFn0(p0 string) int {
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 int = (fi5 & fi5)
        v7 := GenS1{F0: 128, F1: 42, F2: (^(fi5))}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{-2147483648, -697, -9223372036854775808, 4294967295}
    p0 = (p0 + p0)
    if len(v8) > 0 {
        v8[0] = 1
    }
    var v9 bool = (3.14159 > float32(0.1))
    _ = p0
    _ = v8
    _ = v9
    return 127
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 65535
    v3 := make(map[uint32]string)
    var v4 uint64 = 2
    for fi10 := 0; fi10 < 16; fi10++ {
        v4 = (^(v4))
        var v11 int8 = 127
        _ = fi10
        _ = v11
    }
    if (false && (false || false)) {
        v4 = (^(v4))
        fmt.Println(v4)
        v12 := v1[9223372036854775807]
        _ = v12
    }
    var v13 uint32 = 2147483647
    var v14 uint64 = v4
    if true {
        for fi15 := 0; fi15 < 13; fi15++ {
            for fi16 := 0; fi16 < 17; fi16++ {
                v2 += v2
                v13 -= (v13 / ((v13) | 1))
                _ = fi16
            }
            var v17 uint8 = 2
            _ = fi15
            _ = v17
        }
    }
    var v18 int8 = 1
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 string = ("\n\t" + ("foo bar" + "foo bar"))
        _ = fi19
        _ = v20
    }
    v21 := []int{400, 441, 2147483647}
    switch "" {
    case "":
        v22 := GenS1{F0: 1, F1: 0, F2: -569}
        _ = v22
    case "foo bar":
        if false {
            delete(v3, (^(v13)))
            var v23 uint32 = 2
            _ = v23
        } else {
            if true {
                if len(v21) > 0 {
                    v21[0] = 0
                }
            } else {
                for fi24 := 0; fi24 < 1; fi24++ {
                    fmt.Println(fi24, v13)
                    v4 = 2
                    _ = fi24
                }
            }
            v2 = (v2 ^ v2)
            var v25 int16 = 32767
            _ = v25
        }
    case "hello":
        fmt.Println(v2)
    default:
        var v26 uint = 0
        _ = v26
    }
    var v27 uint = 4294967295
    fmt.Println(v2)
    var v28 int16 = -379
    v4 = (v4 * 0)
    var v29 float64 = (float64((3.14159 / -2.25)) + -2.25)
    var v30 float64 = float64(3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
