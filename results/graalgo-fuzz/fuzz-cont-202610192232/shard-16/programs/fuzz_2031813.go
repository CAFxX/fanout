package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
func genFn0(p0 int8, p1 string, p2 string) int64 {
    if (149 == 1) {
        v6 := GenS1{F0: 0}
        p2 = (p2 + p1)
        _ = v6
    } else {
        if false {
            p0 = -128
            p0 = p0
            var v7 float64 = (0.1 * 3.14159)
            _ = v7
        }
    }
    var v8 float64 = (((1e10 * (((-2.25 - -2.25) * 0.0) + -2.25)) / (1.5 + float64((1.5 - 0.0)))) / 1.0)
    var v9 uint8 = 201
    var v10 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 0
    v2 := []int{-589, 127}
    var v3 uint8 = 172
    var v4 uint8 = v3
    var v5 float32 = -2.25
    if true {
        v3 += v3
        switch 127 {
        case 798:
            v11 := GenS1{F0: 0}
            _ = v11
        }
    }
    fmt.Println(v4, v1)
    v4 -= 127
    _ = len(v2)
    v5 += (v5 * ((v5 / ((v5 - (v5 * v5)) - float32((((3.14159 + (float64(0.1) / 1.5)) + (0.1 / -2.25)) * 3.14159)))) * float32(0.0)))
    if true {
        fmt.Println(v4)
    }
    var v12 int16 = 1
    v5 -= float32(1e10)
    v5 = (3.14159 * v5)
    switch (((("foo bar" + ("a" + ("\n\t" + "foo bar"))) + "\n\t") + "") + "hello") {
    case "a":
        v5 -= (v5 / v5)
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v13 int8 = int8(v4)
        _ = v13
    }
    v14 := GenS2{F0: v12, F1: (("" + "hello") + "a"), F2: 395}
    v1 += v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 794 {
    case 32767:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 128:
        v5 = v5
    case 243:
        v2 = append(v2, 0)
    default:
        if false {
            fmt.Println(v12, v5)
        } else {
            for fi15 := 0; fi15 < 19; fi15++ {
                if ((-111 <= 32) && (true && true)) {
                    var v16 uint64 = 1
                    fmt.Println(v5, v3, v16)
                    _ = v16
                }
                v12 += v12
                _ = fi15
            }
            v1 -= 2
            v17 := GenS2{F0: (v12 >> ((308) & 63)), F1: "a", F2: 2}
            _ = v17
        }
    }
    var v18 uint16 = 32767
    v19 := GenS2{F0: -32768, F1: "\n\t", F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
