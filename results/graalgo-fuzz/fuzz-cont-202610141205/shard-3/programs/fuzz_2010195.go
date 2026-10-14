package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0(p0 uint, p1 uint8, p2 int64) int32 {
    var v7 uint16 = 32767
    fmt.Println(p1, v7, p2)
    var v8 int64 = 383
    if (1 < 2147483647) {
        if (true && true) {
            fmt.Println(v7, p2)
        } else {
            p1 = p1
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint32) uint8 {
    switch ("" + "hello") {
    case "hello":
        fmt.Println(p0)
    case "a":
        p0 = p0
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            p0 += (p0 >> ((255) & 63))
            _ = fi9
        }
    }
    switch "\n\t" {
    case "hello":
        var v10 string = "\n\t"
        _ = v10
    case "":
        p0 = (p0 + 0)
    }
    _ = p0
    return 255
}

func main() {
    var v1 float64 = (1e10 / 0.1)
    var v2 uint64 = 840
    v3 := make(map[string]string)
    var v4 int32 = -1
    var v5 float64 = (v1 / 1e10)
    var v6 uint64 = (v2 / (((^(v2))) | 1))
    var v11 int64 = -9223372036854775808
    var v12 int = 42
    var v13 int = v12
    delete(v3, ("foo bar" + ""))
    if (true || true) {
        if (((((-2.25 + 1.5) - 0.1) == -2.25) || true) || (!false)) {
            if false {
                var v14 int = v12
                _ = v14
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                v13 = 0
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 float32 = 3.14159
                    if true {
                        var v18 int16 = 0
                        var v19 bool = false
                        v20 := GenS1{F0: v13, F1: 0}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    _ = fi16
                    _ = v17
                }
                for fi21 := 0; fi21 < 7; fi21++ {
                    if (3.14159 <= (float32(v1) / float32(((float64(-2.25) - 1.5) * v5)))) {
                        var v22 uint8 = 120
                        _ = v22
                    }
                    var v23 int64 = (v11 | (-(v11)))
                    _ = fi21
                    _ = v23
                }
                _ = fi15
            }
        }
        v11 += -1
        v6 += v6
    } else {
        v5 += ((float64((((-2.25 / float32((v5 / float64(0.1)))) - 0.1) + 0.1)) + v1) - v1)
    }
    var v24 int = 32768
    var v25 int = v24
    v26 := GenS1{F0: (v25 * 32768), F1: 2147483648}
    var v27 uint16 = 42
    switch ("hello" + "hello") {
    case "":
        v28 := []int{(v13 | int(v4)), v24, -1}
        _ = v28
    default:
        var v29 uint32 = 219
        _ = v29
    }
    v30 := []int{v13, 9223372036854775807, v24}
    v3["foo bar"] = "foo bar"
    fmt.Println(v24, v13, v1)
    v31 := GenS1{F0: (v13 + (-(v12))), F1: 233}
    v13 -= (v12 & (v25 % ((v25) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
