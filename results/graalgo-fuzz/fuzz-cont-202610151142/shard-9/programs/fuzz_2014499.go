package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
func genFn0() int16 {
    var v4 uint8 = 1
    for fi5 := 0; fi5 < 18; fi5++ {
        switch "a" {
        case "":
            v4 = (^(v4))
        case "hello":
            var v6 uint32 = 2147483648
            _ = v6
        case "foo bar":
            fmt.Println(fi5, fi5)
        }
        v7 := GenS1{F0: 725}
        _ = fi5
        _ = v7
    }
    switch "\n\t" {
    case "foo bar":
        v8 := GenS1{F0: -42}
        _ = v8
    case "\n\t":
        v4 = (v4 >> ((333) & 63))
    }
    fmt.Println(v4)
    _ = v4
    return 1
}
func genFn1() int {
    v9 := []int{-589, 42, -29}
    if len(v9) > 0 {
        v9[0] = -281
    }
    var v10 uint32 = 42
    var v11 uint32 = (^(v10))
    var v12 float32 = (((float32(1.5) + float32((1e10 + 1e10))) + (-2.25 * 0.1)) - float32(1e10))
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return -94
}

func main() {
    v1 := []string{"", "a"}
    var v2 uint = 454
    v3 := []string{""}
    v1 = append(v1, ("foo bar" + "hello"))
    v2 = 76
    fmt.Println(v2)
    v2 = (v2 * (v2 | v2))
    v2 = (v2 & (v2 - (v2 << (((v2 >> (((v2 - (v2 + v2))) & 63))) & 63))))
    v2 += (v2 % (((v2 ^ (v2 % (((v2 / ((279) | 1))) | 1)))) | 1))
    var v13 uint = 1
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "")
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS1{F0: -161}
        v16 := GenS2{F0: (fi14 + genFn1())}
        _ = fi14
        _ = v15
        _ = v16
    }
    if len(v1) > 0 {
        v1[0] = ("a" + "")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
}
