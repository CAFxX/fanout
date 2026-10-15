package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint32 = 4294967295
    var v2 int16 = 42
    v3 := make(map[string]int)
    var v4 float32 = float32(float64(3.14159))
    var v5 int16 = (v2 + 1)
    var v6 uint32 = 0
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v4)
        var v8 uint = 18446744073709551615
        _ = fi7
        _ = v8
    }
    v9 := []int{-203}
    if len(v9) > 0 {
        v9[0] = -1
    }
    fmt.Println(v5)
    v10 := GenS1{F0: 255, F1: 821}
    for fi11 := 0; fi11 < 15; fi11++ {
        switch fi11 {
        case 4294967295:
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 uint16 = 1
                fmt.Println(fi11)
                v6 += v1
                fmt.Println(v1, v2, v4)
                var v14 uint32 = (^(v1))
                _ = fi12
                _ = v13
                _ = v14
            }
        case -9223372036854775808:
            v4 = v4
        case 25:
            v2 = (v2 >> ((4294967295) & 63))
        default:
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 uint64 = 790
                v5 -= (v2 - -1)
                fmt.Println(v2, v16)
                var v17 uint64 = 38
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                var v18 int16 = v5
                switch (fi11 ^ 2147483647) {
                case 4294967295:
                    fmt.Println(v17)
                }
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        _ = fi11
    }
    var v19 string = "hello"
    switch "\n\t" {
    case "a":
        var v20 uint16 = 2
        _ = v20
    case "\n\t":
        v21 := GenS2{F0: (false && ((!true) || true))}
        _ = v21
    case "":
        v19 = (v19 + (("hello" + (("\n\t" + "hello") + v19)) + ((v19 + "foo bar") + (v19 + ("\n\t" + "")))))
    default:
        v6 = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
