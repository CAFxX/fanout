package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
func genFn0(p0 int) int8 {
    var v6 int16 = 32767
    if true {
        var v7 uint64 = 2
        v8 := []int{p0, (p0 & (p0 >> ((42) & 63))), 9223372036854775807, (p0 - 32767)}
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                switch 914 {
                case -213:
                    v11 := GenS1{F0: -783, F1: 784}
                    _ = v11
                }
                _ = fi10
            }
            v7 = 1
            v12 := GenS2{F0: uint32(v7), F1: 18446744073709551615, F2: "\n\t"}
            _ = fi9
            _ = v12
        }
        _ = v7
        _ = v8
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := []int{fi13}
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    }
    v6 = (v6 | (v6 << ((18446744073709551615) & 63)))
    fmt.Println(p0, v6)
    v16 := []int{(p0 - p0), (p0 ^ -358)}
    _ = p0
    _ = v16
    _ = v6
    return -80
}
func genFn1(p0 int64, p1 int64, p2 int32) int8 {
    fmt.Println(p1, p1)
    p0 = (^(p1))
    var v17 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    v3 := make(map[string]string)
    var v4 uint32 = 2
    var v5 float32 = 0.0
    var v18 int64 = 9223372036854775807
    switch "\n\t" {
    case "":
        v5 += v5
    case "\n\t":
        var v19 uint = 4294967295
        _ = v19
    case "foo bar":
        _ = len(v3)
    }
    var v20 float32 = v5
    var v21 string = (("" + "foo bar") + (("a" + ("foo bar" + ("hello" + "foo bar"))) + ("\n\t" + "a")))
    v22 := GenS1{F0: (v18 + (v18 | v18)), F1: (v4 * (v4 / (((v4 * (^(v4)))) | 1)))}
    v23 := GenS2{F0: (v4 % (((v4 / (((^(v4))) | 1))) | 1)), F1: 0, F2: (v21 + (v21 + (v21 + v21)))}
    fmt.Println(v20, v5, v21)
    if false {
        v2 = v2
        v21 = v21
    }
    var v24 float32 = (v5 * v20)
    switch ("a" + ("foo bar" + v21)) {
    case "a":
        var v25 uint64 = 4294967295
        _ = v25
    case "hello":
        v26 := []int{32767}
        _ = v26
    default:
        _ = len(v1)
    }
    fmt.Println(v2, v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
