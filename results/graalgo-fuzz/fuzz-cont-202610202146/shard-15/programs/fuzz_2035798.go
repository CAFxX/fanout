package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 bool; F1 int32; F2 int }
func genFn0() int32 {
    var v6 uint = 0
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 += (v6 + 1)
        _ = fi7
    }
    var v8 int32 = -2147483648
    v6 = (v6 / (((v6 ^ v6)) | 1))
    _ = v6
    _ = v8
    return -1
}
func genFn1(p0 uint8, p1 uint8) uint8 {
    var v9 int = 191
    v10 := GenS1{F0: "foo bar", F1: (p0 | (p1 + (p0 + (p1 % (((^(p0))) | 1)))))}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return p0
}

func main() {
    var v1 int = 128
    var v2 int32 = 42
    v3 := make(map[int]string)
    var v4 int32 = (v2 / ((v2) | 1))
    var v5 string = "a"
    v1 = v1
    var v11 int = (v1 + (v1 * v1))
    v1 -= v1
    v12 := []int{}
    if (127 >= 42) {
        switch ((v5 + (v5 + v5)) + v5) {
        case "\n\t":
            v2 -= (v4 - (v4 / (((-(v2))) | 1)))
        case "a":
            v13 := GenS2{F0: ((v2 << ((18446744073709551615) & 63)) <= genFn0()), F1: (v4 / ((v2) | 1)), F2: v1}
            _ = v13
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := GenS2{F0: (!true), F1: (v2 ^ v4), F2: fi14}
            _ = fi14
            _ = v15
        }
        fmt.Println(v11, v11)
    }
    v5 = (((v5 + v5) + ((v5 + "a") + ("foo bar" + "a"))) + "a")
    v5 = "foo bar"
    switch (v5 + (v5 + (v5 + v5))) {
    case "foo bar":
        v4 = (v2 | v4)
    default:
        v12 = append(v12, (v11 - (^(v1))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
