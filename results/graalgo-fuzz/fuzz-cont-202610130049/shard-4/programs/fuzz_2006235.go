package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
func genFn0(p0 uint64) int {
    switch (("\n\t" + ("hello" + "a")) + ((("\n\t" + "\n\t") + ("a" + "\n\t")) + "foo bar")) {
    case "foo bar":
        v7 := []int{750, 74, 128, 218}
        _ = v7
    case "a":
        switch -1 {
        case 56:
            v8 := []int{-391, 128, 1}
            _ = v8
        default:
            if false {
                fmt.Println(p0)
                if (662 != 18446744073709551615) {
                    v9 := []int{0, 32767, 32768, 127}
                    var v10 int64 = 397
                    fmt.Println(v10, p0)
                    _ = v10
                    _ = v9
                }
            }
        }
    default:
        p0 = p0
    }
    v11 := []int{2147483647}
    if ("hello" <= "a") {
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: 0, F1: 42, F2: 51}
            var v14 float32 = -2.25
            v14 += (v14 / float32(0.1))
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 string = "\n\t"
        _ = v15
    }
    var v16 uint8 = 1
    var v17 int8 = -128
    _ = p0
    _ = v11
    _ = v16
    _ = v17
    return -258
}

func main() {
    v1 := []string{"", "\n\t", ("a" + "hello")}
    var v2 float32 = float32((0.1 / 1e10))
    v3 := []uint32{}
    var v4 int64 = -321
    v5 := make(map[int]uint32)
    v6 := make(map[string]string)
    switch "foo bar" {
    case "foo bar":
        if false {
            var v18 float64 = (1e10 / (0.1 * -2.25))
            if len(v1) > 0 {
                v1[0] = (("foo bar" + ("hello" + "a")) + "a")
            }
            _ = v18
        } else {
            v19 := GenS1{F0: 0, F1: 440, F2: 127}
            var v20 uint16 = 1
            if len(v3) > 0 {
                v3[0] = 42
            }
            _ = v19
            _ = v20
        }
    case "hello":
        v21 := GenS1{F0: 467, F1: 1, F2: 0}
        _ = v21
    }
    fmt.Println(v4, v2)
    if true {
        var v22 float32 = (v2 - v2)
        v4 = v4
        _ = v22
    } else {
        v4 = 9223372036854775807
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        v24 := GenS1{F0: 531, F1: 255, F2: 100}
        v4 += (v4 | (v4 * v4))
        v1 = append(v1, ("\n\t" + "hello"))
        _ = fi23
        _ = v24
    }
    fmt.Println(v2)
    switch 831 {
    case 226:
        v2 = ((v2 / (1.5 / (v2 / -2.25))) - (((float32(1e10) + 3.14159) - float32(1e10)) / float32(0.1)))
    default:
        fmt.Println(v4, v4)
    }
    var v25 int32 = 2147483647
    if len(v1) > 0 {
        v1[0] = (("\n\t" + (("\n\t" + ("" + "a")) + (("a" + (("a" + "") + ("a" + "hello"))) + ("hello" + "\n\t")))) + "a")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
