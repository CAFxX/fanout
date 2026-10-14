package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
func genFn0() int8 {
    var v6 uint64 = 1
    v7 := []int{4294967295}
    switch 128 {
    case 127:
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int8 = 100
            _ = fi8
            _ = v9
        }
    case 128:
        v6 -= v6
    case -92:
        switch ((("foo bar" + ("hello" + ("hello" + "a"))) + "a") + "") {
        case "":
            switch (("\n\t" + "") + "hello") {
            case "\n\t":
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint = 0
                    _ = fi10
                    _ = v11
                }
            case "hello":
                v6 += (v6 | v6)
            }
        case "foo bar":
            var v12 uint32 = 2147483648
            _ = v12
        case "hello":
            fmt.Println(v6)
        }
    }
    var v13 uint32 = 0
    if len(v7) > 0 {
        v7[0] = -912
    }
    _ = v13
    _ = v6
    _ = v7
    return -82
}
func genFn1(p0 uint16) uint {
    p0 -= p0
    var v14 uint64 = 555
    _ = p0
    _ = v14
    return 2
}

func main() {
    var v1 float64 = ((((((-2.25 * (float64((0.0 + -2.25)) + 0.1)) - ((1.5 - 3.14159) * 1.5)) * 1e10) + (-2.25 / 1.0)) / 1.5) / 1.0)
    v2 := []int{-2147483648, 42, 9223372036854775807}
    var v3 int8 = 1
    var v4 uint64 = 1
    var v5 uint = 42
    v4 -= 18446744073709551615
    v15 := GenS1{F0: 128, F1: 42}
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 uint8 = 0
        _ = fi16
        _ = v17
    }
    fmt.Println(v1)
    v18 := GenS1{F0: 128, F1: -1}
    v1 -= (((((((0.1 - 1.5) - v1) + 0.0) / 1e10) - (1e10 / v1)) - (v1 - (1.5 + 0.1))) * v1)
    for fi19 := 0; fi19 < 11; fi19++ {
        v3 = genFn0()
        _ = fi19
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
