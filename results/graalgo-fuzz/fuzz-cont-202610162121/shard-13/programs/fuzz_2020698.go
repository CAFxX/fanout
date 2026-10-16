package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint32 }
func genFn0() uint {
    switch -132 {
    case -521:
        if true {
            switch ("" + (("a" + "hello") + "a")) {
            case "":
                if (0 <= 100) {
                    v6 := []int{-2147483648}
                    var v7 uint32 = 2147483647
                    _ = v6
                    _ = v7
                }
            case "a":
                v8 := []int{42, -1}
                _ = v8
            case "hello":
                for fi9 := 0; fi9 < 20; fi9++ {
                    switch "hello" {
                    case "a":
                        switch (fi9 + fi9) {
                        case -1:
                            v10 := []int{(-(fi9)), fi9, (fi9 - fi9), (fi9 | 32767)}
                            _ = v10
                        }
                    case "hello":
                        fmt.Println(fi9)
                    }
                    var v11 int64 = 42
                    for fi12 := 0; fi12 < 5; fi12++ {
                        v13 := GenS1{F0: 18446744073709551615, F1: 1}
                        v14 := GenS2{F0: 2}
                        v11 = (v11 >> ((255) & 63))
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = fi9
                    _ = v11
                }
            }
            v15 := []int{-1, 369, 127, 0}
            _ = v15
        }
    }
    var v16 uint32 = 0
    v17 := GenS1{F0: 65535, F1: -127}
    _ = v16
    _ = v17
    return 0
}
func genFn1(p0 string) int16 {
    var v18 uint64 = 115
    v18 = (v18 * v18)
    v19 := GenS1{F0: 584, F1: -1}
    _ = p0
    _ = v18
    _ = v19
    return 0
}

func main() {
    var v1 bool = true
    v2 := []uint32{206}
    v3 := make(map[string]string)
    var v4 uint32 = 2147483648
    var v5 float64 = ((-2.25 * -2.25) - 0.0)
    if v1 {
        for fi20 := 0; fi20 < 15; fi20++ {
            v1 = ("hello" > "")
            var v21 uint16 = 65535
            _ = fi20
            _ = v21
        }
    }
    v22 := v3["\n\t"]
    v4 = v4
    v23 := GenS1{F0: 1, F1: 100}
    v3[v22] = v22
    if len(v2) > 0 {
        v2[0] = v4
    }
    var v24 uint = 516
    if v1 {
        v5 = 1e10
    }
    var v25 uint16 = 507
    v22 = v22
    for fi26 := 0; fi26 < 11; fi26++ {
        fmt.Println(v24, v22, v1)
        var v27 uint32 = uint32(v25)
        _ = fi26
        _ = v27
    }
    var v28 uint16 = (v25 - v25)
    fmt.Println(v5, v4)
    for fi29 := 0; fi29 < 17; fi29++ {
        fmt.Println(v4)
        var v30 uint32 = (v4 << (((v24 >> (((v24 << ((genFn0()) & 63))) & 63))) & 63))
        switch v22 {
        case "hello":
            fmt.Println(v4, v24, v25)
        case "\n\t":
            var v31 uint64 = 4294967295
            _ = v31
        case "foo bar":
            fmt.Println(v5, v30, v30)
        }
        _ = fi29
        _ = v30
    }
    switch 2147483647 {
    case -174:
        var v32 uint = v24
        _ = v32
    }
    var v33 uint64 = 4294967295
    v25 += (v28 + (^(v25)))
    var v34 uint32 = 943
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
