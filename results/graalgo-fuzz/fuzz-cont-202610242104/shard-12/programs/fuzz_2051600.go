package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
func genFn0(p0 int32) int8 {
    p0 -= 0
    if true {
        var v5 float64 = float64(1e10)
        v6 := []int{137}
        if len(v6) > 0 {
            v6[0] = 1
        }
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: ((("" + ("\n\t" + "foo bar")) + "a") + "\n\t"), F1: -32768}
    _ = p0
    _ = v7
    return -122
}
func genFn1(p0 int) bool {
    p0 += p0
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint32 = 65535
        v10 := []int{(p0 ^ fi8)}
        p0 -= fi8
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{p0, p0}
    _ = p0
    _ = v11
    return true
}

func main() {
    v1 := []uint32{731, 528, 0, 2}
    v2 := make(map[int]bool)
    var v3 uint = 440
    var v4 uint32 = 2
    var v12 float64 = (-2.25 - float64(float32((1e10 + (1e10 - 0.1)))))
    v13 := GenS1{F0: (("hello" + ("a" + "\n\t")) + ""), F1: 1000}
    var v14 uint32 = (v4 % (((v4 * (v4 % ((81) | 1)))) | 1))
    var v15 int64 = 9223372036854775807
    v15 = -582
    fmt.Println(v3)
    if (180 > 1) {
        if true {
            for fi16 := 0; fi16 < 17; fi16++ {
                fmt.Println(v4)
                var v17 uint32 = (v14 - (v14 % (((v4 * (v4 | (^(v14))))) | 1)))
                _ = fi16
                _ = v17
            }
        }
        var v18 string = "\n\t"
        _ = v18
    }
    v4 -= 943
    var v19 string = "hello"
    switch v19 {
    case "":
        var v20 int32 = -861
        _ = v20
    case "\n\t":
        fmt.Println(v4)
    default:
        var v21 string = ((((((v19 + ((v19 + ((v19 + "") + (v19 + (v19 + "hello")))) + ((v19 + (("a" + ("\n\t" + (((v19 + (v19 + v19)) + v19) + v19))) + (v19 + v19))) + (v19 + "a")))) + v19) + (v19 + ("hello" + "\n\t"))) + v19) + v19) + v19)
        _ = v21
    }
    if true {
        for fi22 := 0; fi22 < 17; fi22++ {
            v12 = -2.25
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 uint8 = 2
                var v25 int = fi22
                _ = fi23
                _ = v24
                _ = v25
            }
            v19 = ((v19 + v19) + "\n\t")
            _ = fi22
        }
    }
    v1 = append(v1, v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
