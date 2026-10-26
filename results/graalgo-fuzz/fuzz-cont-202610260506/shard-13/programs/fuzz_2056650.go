package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 bool }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 int32, p1 uint16, p2 uint8) bool {
    var v5 uint16 = 482
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ((false || ((false && ((p2 & p2) != (p2 % (((p2 >> ((820) & 63))) | 1)))) || false)) && true)
}

func main() {
    var v1 int = 950
    var v2 bool = false
    var v3 int8 = 100
    v4 := make(map[int]bool)
    fmt.Println(v3, v1, v3)
    switch ("foo bar" + ("foo bar" + ("foo bar" + ""))) {
    case "hello":
        switch -2147483648 {
        case 2147483647:
            fmt.Println(v3, v1, v3)
        }
    case "foo bar":
        var v6 uint16 = 574
        _ = v6
    default:
        v2 = v2
    }
    fmt.Println(v1, v3, v1)
    v7 := GenS1{F0: 341}
    switch (v1 / ((v1) | 1)) {
    case -492:
        fmt.Println(v2, v2)
    }
    if false {
        var v8 uint64 = 2
        _ = len(v4)
        v1 += v1
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    v10 := GenS2{F0: 0, F1: 32767, F2: false}
    v11 := GenS3{F0: (^(v1)), F1: "foo bar", F2: 127}
    switch v1 {
    case 465:
        v12 := []int{v1, (-(v1)), v1}
        _ = v12
    default:
        switch 1 {
        case -732:
            _ = len(v4)
        case 127:
            if true {
                v2 = (0.1 >= 0.0)
                v13 := GenS1{F0: 2}
                if v2 {
                    switch v1 {
                    case 127:
                        var v14 int32 = 2147483647
                        _ = v14
                    case 1:
                        v15 := []int{v1, (v1 & (v1 / (((v1 * (v1 / (((v1 >> ((18446744073709551615) & 63))) | 1)))) | 1))), v1}
                        _ = v15
                    case -934:
                        v16 := []int{v1, 128}
                        _ = v16
                    default:
                        var v17 string = "hello"
                        _ = v17
                    }
                    if (!((!((1e10 * float64(((1e10 * (3.14159 / 0.1)) / 1e10))) == 0.0)) || false)) {
                        switch ("hello" + "hello") {
                        case "foo bar":
                            var v18 string = "\n\t"
                            _ = v18
                        default:
                            v19 := []int{v1, 42, (v1 << ((18446744073709551615) & 63))}
                            _ = v19
                        }
                        var v20 string = "a"
                        _ = v20
                    }
                } else {
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v22 := GenS4{F0: (("a" + ((("foo bar" + "\n\t") + ("hello" + "hello")) + "hello")) + "")}
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v13
            }
        default:
            v23 := GenS4{F0: "foo bar"}
            _ = v23
        }
    }
    var v24 int = (v1 & v1)
    if v2 {
        switch (v1 ^ v1) {
        case 32767:
            var v25 int32 = int32(v24)
            _ = v25
        case 0:
            v1 += -1
        case -2147483648:
            fmt.Println(v1, v9)
        default:
            v26 := []int{9223372036854775807, v1, v1, v24}
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
