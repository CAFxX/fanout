package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 string }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0() string {
    v4 := []int{2147483647}
    v5 := GenS1{F0: 185, F1: 1, F2: ("" + ("\n\t" + (("a" + "hello") + "")))}
    v6 := GenS2{F0: 128, F1: 127}
    var v7 uint64 = 1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return "a"
}
func genFn1(p0 int16) int8 {
    v8 := []int{0}
    var v9 string = ("foo bar" + ("\n\t" + "foo bar"))
    v10 := GenS1{F0: 1, F1: 0, F2: v9}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return -55
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    var v3 string = "\n\t"
    var v11 uint = 18446744073709551615
    fmt.Println(v3, v3)
    if (!(1 > 127)) {
        fmt.Println(v11, v3)
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        switch v3 {
        case "a":
            fmt.Println(v3, v11, fi12)
        default:
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 int64 = -903
                v11 += 4294967295
                _ = fi13
                _ = v14
            }
        }
        var v15 uint8 = uint8(v11)
        _ = fi12
        _ = v15
    }
    fmt.Println(v3)
    v3 = ("a" + (("foo bar" + ("foo bar" + ("\n\t" + v3))) + (v3 + "a")))
    if true {
        v16 := GenS1{F0: 241, F1: 1, F2: "foo bar"}
        v3 = "\n\t"
        v11 -= (^(v11))
        _ = v16
    }
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
