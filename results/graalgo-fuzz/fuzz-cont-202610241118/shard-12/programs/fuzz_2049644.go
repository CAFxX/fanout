package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0(p0 int16, p1 int8) uint64 {
    if false {
        p1 += (p1 | (p1 & 0))
    }
    fmt.Println(p0, p1)
    switch 9223372036854775807 {
    case 899:
        var v4 int = 540
        _ = v4
    case 0:
        p1 -= -19
    }
    _ = p0
    _ = p1
    return 437
}

func main() {
    v1 := []string{("hello" + ("a" + "foo bar")), "", ""}
    var v2 int = 1
    v3 := make(map[int]uint32)
    v2 = v2
    v2 = v2
    fmt.Println(v2)
    v2 = v2
    _ = len(v3)
    v2 = v2
    var v5 float64 = 0.0
    var v6 int = 9223372036854775807
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        fmt.Println(v2, v5)
    }
    if ("hello" < "") {
        var v7 float32 = float32(((3.14159 / v5) + v5))
        v8 := GenS1{F0: 127}
        switch 32768 {
        case 9223372036854775807:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 4294967295:
            var v9 int8 = -98
            _ = v9
        case 128:
            v2 = 923
        default:
            var v10 float64 = v5
            _ = v10
        }
        _ = v7
        _ = v8
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 uint32 = 2147483647
            v14 := GenS1{F0: 1}
            var v15 int16 = 1
            var v16 uint8 = 0
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
            _ = v16
        }
        fmt.Println(v5, v2)
        var v17 string = ("\n\t" + ("" + "hello"))
        _ = fi11
        _ = v17
    }
    v18 := GenS1{F0: 2}
    fmt.Println(v2, v5, v2)
    v19 := GenS1{F0: 128}
    if (((!true) && true) || (4294967295 >= 18446744073709551615)) {
        v6 -= 4294967295
        fmt.Println(v2, v2, v5)
    } else {
        for fi20 := 0; fi20 < 13; fi20++ {
            v5 += ((v5 + 1.5) - (((1e10 + v5) + (v5 - 1.5)) * (-2.25 * (v5 - float64(-2.25)))))
            v6 -= (v2 + 2147483647)
            var v21 float64 = 1.5
            _ = fi20
            _ = v21
        }
        var v22 int16 = -1
        v23 := GenS2{F0: "", F1: 916, F2: 1}
        _ = v22
        _ = v23
    }
    v2 += (v6 / ((-2147483648) | 1))
    switch v2 {
    case -397:
        fmt.Println(v6, v6)
    case 0:
        v6 -= v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
