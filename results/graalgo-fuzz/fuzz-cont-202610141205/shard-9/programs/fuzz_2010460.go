package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }

func main() {
    var v1 uint8 = 39
    var v2 string = "foo bar"
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]bool)
    var v5 int64 = -1
    if true {
        fmt.Println(v1, v2, v1)
        v6 := GenS1{F0: 2147483648, F1: 18446744073709551615, F2: "\n\t"}
        switch (((v2 + "a") + "a") + "hello") {
        case "":
            var v7 float32 = 0.0
            _ = v7
        default:
            _ = len(v3)
        }
        _ = v6
    } else {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := []int{fi8}
            v5 += (v5 - (-(v5)))
            _ = fi8
            _ = v9
        }
        if (-969 > -2147483648) {
            v2 = v2
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 uint = 281
                v1 = (v1 ^ 27)
                _ = fi10
                _ = v11
            }
        }
    }
    var v12 int8 = 0
    if (true && true) {
        v1 += (v1 - 35)
    } else {
        if (!true) {
            v13 := []int{}
            _ = v13
        } else {
            fmt.Println(v1, v2, v1)
            var v14 float32 = float32((0.1 + (-2.25 + 1.5)))
            _ = v14
        }
        switch 107 {
        case -9223372036854775808:
            v15 := GenS1{F0: 2147483647, F1: 328, F2: "\n\t"}
            _ = v15
        case 2147483647:
            v2 = ((((v2 + ((((v2 + v2) + "hello") + "foo bar") + v2)) + ("\n\t" + "hello")) + v2) + (v2 + (v2 + ("" + (v2 + ("foo bar" + (v2 + ("a" + v2))))))))
        case 9223372036854775807:
            v16 := []int{-126}
            _ = v16
        default:
            v17 := GenS1{F0: 302, F1: 1, F2: "a"}
            _ = v17
        }
    }
    v2 = (("" + v2) + v2)
    v12 += -128
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(v2, v1, v2)
        _ = fi18
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
