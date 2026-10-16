package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint32 }
func genFn0(p0 uint, p1 uint64) string {
    for fi4 := 0; fi4 < 13; fi4++ {
        switch "foo bar" {
        case "\n\t":
            v5 := []int{fi4, fi4, (fi4 - fi4)}
            _ = v5
        case "foo bar":
            v6 := GenS1{F0: p1, F1: 143, F2: 2147483648}
            _ = v6
        case "hello":
            if (1.5 != (0.1 * 1.5)) {
                var v7 uint64 = p1
                fmt.Println(p0, v7)
                var v8 uint64 = (p1 * uint64(p0))
                _ = v7
                _ = v8
            }
        default:
            v9 := GenS1{F0: p1, F1: 255, F2: 0}
            _ = v9
        }
        _ = fi4
    }
    var v10 float32 = (3.14159 + float32(-2.25))
    _ = p0
    _ = p1
    _ = v10
    return ("a" + "\n\t")
}
func genFn1(p0 int8) int8 {
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint = 672
        v13 := []int{}
        if false {
            var v14 uint64 = 42
            _ = v14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v15 uint = 977
    _ = p0
    _ = v15
    return 1
}

func main() {
    var v1 string = ("" + (("\n\t" + ("" + "\n\t")) + "\n\t"))
    var v2 bool = true
    var v3 int = 2147483647
    fmt.Println(v2)
    v3 += -2147483648
    if false {
        v16 := []int{v3, v3, v3}
        var v17 float64 = -2.25
        switch ("" + (v1 + "hello")) {
        case "":
            fmt.Println(v3, v3, v1)
        default:
            fmt.Println(v2)
        }
        _ = v16
        _ = v17
    }
    v18 := []int{}
    v2 = ((!v2) && true)
    v1 = "a"
    var v19 uint8 = 231
    if v2 {
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    } else {
        var v20 bool = v2
        v21 := GenS2{F0: 127, F1: 0, F2: 42}
        var v22 int = v3
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: -33, F1: 0, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
}
