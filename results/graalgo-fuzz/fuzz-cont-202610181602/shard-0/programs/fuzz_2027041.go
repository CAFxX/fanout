package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int) uint8 {
    var v6 int8 = 127
    v7 := GenS1{F0: 4294967295, F1: -2147483648}
    for fi8 := 0; fi8 < 12; fi8++ {
        p0 += 127
        v6 -= (v6 >> ((18446744073709551615) & 63))
        _ = fi8
    }
    v9 := GenS2{F0: 0}
    v6 -= v6
    _ = p0
    _ = v6
    _ = v7
    _ = v9
    return 236
}

func main() {
    var v1 bool = (true && true)
    var v2 int = -956
    var v3 int = 9223372036854775807
    var v4 int8 = 100
    var v5 float64 = 1e10
    v3 -= (v3 / (((v2 ^ (^(v3)))) | 1))
    var v10 int64 = -742
    var v11 int = v2
    switch ((("a" + "foo bar") + ("" + (("" + "a") + (("hello" + ("\n\t" + "a")) + "foo bar")))) + ((("\n\t" + "foo bar") + "foo bar") + "\n\t")) {
    case "":
        switch (-(v2)) {
        case 127:
            v12 := []int{(v3 >> ((2) & 63)), 32768}
            _ = v12
        }
    }
    v3 -= (v11 & 165)
    fmt.Println(v4)
    v3 -= (v2 * v2)
    var v13 uint8 = 255
    if v1 {
        var v14 float64 = v5
        v15 := []int{}
        _ = v14
        _ = v15
    } else {
        v16 := []int{v2, v11, v11, v2}
        for fi17 := 0; fi17 < 17; fi17++ {
            v3 = v11
            _ = fi17
        }
        switch "a" {
        case "hello":
            for fi18 := 0; fi18 < 12; fi18++ {
                switch "\n\t" {
                case "":
                    v19 := GenS3{F0: 42}
                    _ = v19
                case "a":
                    v4 += 21
                }
                var v20 uint16 = 0
                var v21 string = ("a" + "a")
                _ = fi18
                _ = v20
                _ = v21
            }
        case "\n\t":
            fmt.Println(v13, v2)
        case "a":
            v13 = (v13 - (v13 * 126))
        }
        _ = v16
    }
    fmt.Println(v5)
    fmt.Println(v4, v5, v13)
    fmt.Println(v1)
    v22 := []int{}
    v11 = (v2 >> ((700) & 63))
    v23 := GenS1{F0: 697, F1: 1}
    v22 = append(v22, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
