package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0() int32 {
    var v4 uint64 = 800
    v5 := []int{0, 42}
    fmt.Println(v4)
    var v6 int8 = 1
    fmt.Println(v4, v6)
    _ = v4
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1() int64 {
    var v7 bool = true
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := []int{-371, 1}
        switch (^(fi8)) {
        case -1:
            _ = len(v9)
        default:
            fmt.Println(fi8, fi8)
        }
        fmt.Println(fi8, v7)
        _ = fi8
        _ = v9
    }
    var v10 string = ("a" + "")
    _ = v10
    _ = v7
    return 1
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    var v11 string = ""
    var v12 string = (((v11 + ("\n\t" + v11)) + v11) + ((v11 + v11) + (v11 + "a")))
    if false {
        switch -460 {
        case 1:
            v13 := GenS1{F0: ("\n\t" + v11), F1: 42, F2: 0}
            _ = v13
        case 993:
            v14 := []int{127, 4294967295, 872}
            _ = v14
        case -2147483648:
            for fi15 := 0; fi15 < 2; fi15++ {
                switch (-(fi15)) {
                case -9223372036854775808:
                    switch (-(fi15)) {
                    case 9223372036854775807:
                        v16 := []int{}
                        _ = v16
                    }
                case 842:
                    var v17 uint = 385
                    _ = v17
                case 127:
                    v11 = "foo bar"
                default:
                    if true {
                        var v18 int = 4294967295
                        _ = v18
                    } else {
                        var v19 bool = (912 != 445)
                        v20 := []int{(-(fi15)), fi15, (fi15 - 32767)}
                        v1 = genFn0()
                        _ = v19
                        _ = v20
                    }
                }
                v21 := GenS1{F0: "a", F1: 379, F2: 32767}
                _ = fi15
                _ = v21
            }
        }
    }
    switch v11 {
    case "hello":
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 bool = false
            _ = fi22
            _ = v23
        }
    case "":
        v24 := []int{0, 615}
        _ = v24
    }
    v25 := GenS1{F0: ("" + "foo bar"), F1: 1, F2: 223}
    v26 := v3[-642]
    var v27 int = 9223372036854775807
    var v28 int16 = 1000
    var v29 bool = v26
    v1 -= v1
    switch (v27 ^ v27) {
    case 42:
        v30 := GenS1{F0: "a", F1: 18446744073709551615, F2: 6}
        _ = v30
    default:
        v31 := GenS2{F0: 42, F1: 296}
        _ = v31
    }
    var v32 int8 = 127
    v32 -= (v32 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
}
