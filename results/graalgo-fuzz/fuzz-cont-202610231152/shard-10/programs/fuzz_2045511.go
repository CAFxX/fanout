package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int32; F1 int }
func genFn0(p0 uint64, p1 uint8) uint32 {
    if (0 > 0) {
        if ((false || (-1 <= 824)) && false) {
            var v4 int64 = -929
            _ = v4
        }
        var v5 uint32 = 2147483647
        var v6 int32 = 1
        _ = v5
        _ = v6
    }
    var v7 uint = 857
    v7 += v7
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 int64 = 42
    v2 := []string{"a"}
    var v3 bool = false
    v8 := GenS1{F0: 987, F1: -1}
    var v9 int16 = 32767
    v9 += 1000
    v1 = -1
    v2 = append(v2, "foo bar")
    var v10 int = -676
    if len(v2) > 0 {
        v2[0] = ("a" + ((("foo bar" + ("hello" + "foo bar")) + "\n\t") + ""))
    }
    fmt.Println(v10, v9)
    _ = len(v2)
    v11 := GenS1{F0: 455, F1: v9}
    switch "hello" {
    case "a":
        for fi12 := 0; fi12 < 4; fi12++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v10, v9, v1)
            v10 = 9223372036854775807
            _ = fi12
        }
    }
    v13 := GenS1{F0: 2147483648, F1: int16(v10)}
    switch ((("a" + "foo bar") + "") + ("foo bar" + ("" + "\n\t"))) {
    case "\n\t":
        var v14 float64 = 0.0
        _ = v14
    case "a":
        v15 := GenS1{F0: 65535, F1: v9}
        _ = v15
    case "hello":
        v1 += v1
    default:
        if v3 {
            fmt.Println(v3, v3, v9)
            var v16 int64 = v1
            v3 = (1 <= 65535)
            _ = v16
        }
    }
    if (4294967295 > 1) {
        v10 = (v10 ^ (^(v10)))
        switch "\n\t" {
        case "foo bar":
            if (!v3) {
                var v17 int64 = 1
                v17 = (v1 & -40)
                _ = v17
            }
        case "a":
            _ = len(v2)
        case "":
            v3 = true
        default:
            for fi18 := 0; fi18 < 9; fi18++ {
                for fi19 := 0; fi19 < 13; fi19++ {
                    v1 -= v1
                    v3 = (!(128 > 127))
                    var v20 uint16 = 226
                    _ = fi19
                    _ = v20
                }
                var v21 float32 = float32((0.1 / 1.5))
                var v22 uint32 = genFn0(18446744073709551615, 127)
                _ = fi18
                _ = v21
                _ = v22
            }
        }
        v1 -= 0
    }
    v10 = v10
    v3 = v3
    fmt.Println(v10)
    v23 := GenS2{F0: 2147483647, F1: v10}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v8
    _ = v9
}
