package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint32; F2 int8 }
func genFn0() uint8 {
    v6 := []int{}
    switch "hello" {
    case "a":
        var v7 bool = true
        _ = v7
    default:
        v8 := GenS1{F0: -2147483648}
        _ = v8
    }
    _ = v6
    return 42
}
func genFn1() string {
    var v9 uint32 = 42
    var v10 int8 = -1
    switch "hello" {
    case "":
        v11 := []int{-79, int(v10), 9223372036854775807}
        _ = v11
    case "a":
        switch "foo bar" {
        case "a":
            for fi12 := 0; fi12 < 5; fi12++ {
                v13 := []int{(fi12 / ((fi12) | 1)), 9223372036854775807, (fi12 ^ fi12), (fi12 - fi12)}
                var v14 uint64 = 4294967295
                _ = fi12
                _ = v13
                _ = v14
            }
        case "":
            if (1 > -9223372036854775808) {
                var v15 float32 = 1e10
                _ = v15
            } else {
                v16 := GenS1{F0: 906}
                _ = v16
            }
        default:
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 uint64 = 18446744073709551615
                var v19 int32 = -2147483648
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    default:
        v9 -= (v9 + 0)
    }
    v20 := []int{-686, 9223372036854775807, -2147483648, 1}
    var v21 uint32 = 843
    _ = v10
    _ = v20
    _ = v21
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int64 = -948
    var v2 uint8 = 26
    var v3 int = -9223372036854775808
    var v4 uint32 = 2147483647
    var v5 float64 = 1.5
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int32 = 0
        if true {
            var v24 string = ("a" + "foo bar")
            v25 := []int{(v3 / ((fi22) | 1)), (fi22 << ((346) & 63))}
            _ = v24
            _ = v25
        } else {
            fmt.Println(v2, v4)
            if false {
                for fi26 := 0; fi26 < 13; fi26++ {
                    fmt.Println(v1, v1, v1)
                    v27 := GenS1{F0: v23}
                    _ = fi26
                    _ = v27
                }
                if ("a" <= "a") {
                    v1 = (v1 << ((0) & 63))
                    v28 := []int{(-(fi22))}
                    if len(v28) > 0 {
                        v28[0] = int(v2)
                    }
                    _ = v28
                }
            }
        }
        if true {
            var v29 float64 = v5
            _ = v29
        }
        _ = fi22
        _ = v23
    }
    v1 = 496
    for fi30 := 0; fi30 < 11; fi30++ {
        v31 := GenS2{F0: 4294967295}
        _ = fi30
        _ = v31
    }
    v5 -= 0.0
    v2 = v2
    v32 := GenS3{F0: v3, F1: (v4 % ((v4) | 1)), F2: -128}
    var v33 bool = false
    v5 = (v5 + (v5 + v5))
    switch "hello" {
    case "foo bar":
        v34 := []int{793, (v3 * (v3 & 9223372036854775807)), (-(v3)), (-(v3))}
        _ = v34
    case "\n\t":
        switch "hello" {
        case "foo bar":
            var v35 string = "\n\t"
            _ = v35
        }
    case "hello":
        v36 := []int{v3, (v3 + v3), (v3 | v3)}
        _ = v36
    }
    v37 := []int{613}
    for fi38 := 0; fi38 < 9; fi38++ {
        fmt.Println(v3)
        _ = fi38
    }
    if len(v37) > 0 {
        v37[0] = (v3 | (v3 | v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
