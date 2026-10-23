package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    var v1 string = "a"
    v2 := []int{-249, 4294967295}
    v3 := []string{v1, ("\n\t" + "a"), (("foo bar" + (v1 + "hello")) + v1)}
    v4 := make(map[uint32]string)
    v5 := make(map[string]bool)
    v6 := make(map[int]int)
    var v7 int = 933
    if (true || true) {
        if false {
            v1 = v1
            var v8 uint64 = 42
            _ = v8
        }
    } else {
        fmt.Println(v1, v1)
        for fi9 := 0; fi9 < 17; fi9++ {
            if false {
                v3 = append(v3, v1)
                switch "hello" {
                case "hello":
                    fmt.Println(v1)
                }
            } else {
                switch (v1 + "a") {
                case "":
                    _ = len(v4)
                case "hello":
                    for fi10 := 0; fi10 < 9; fi10++ {
                        v1 = "hello"
                        _ = fi10
                    }
                case "foo bar":
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v1 = v1
                        _ = fi11
                    }
                default:
                    for fi12 := 0; fi12 < 17; fi12++ {
                        v7 -= fi9
                        v13 := GenS1{F0: (177 < 65535)}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi12
                        _ = v13
                    }
                }
                v7 -= (v7 | 1)
                fmt.Println(fi9, fi9, v7)
            }
            _ = fi9
        }
    }
    fmt.Println(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v14 uint32 = 42
    for fi15 := 0; fi15 < 8; fi15++ {
        v7 -= -879
        _ = fi15
    }
    v16 := GenS2{F0: 604, F1: 821}
    if true {
        switch ((("a" + v1) + v1) + (v1 + (v1 + ""))) {
        case "foo bar":
            fmt.Println(v14, v1)
        case "hello":
            var v17 uint = uint(v14)
            _ = v17
        case "\n\t":
            fmt.Println(v1, v7, v1)
        }
        if true {
            if ((true && (((949 != 255) || (!false)) || false)) && (128 <= 228)) {
                v1 = ((((v1 + v1) + v1) + ((v1 + v1) + "hello")) + v1)
                fmt.Println(v14)
                if (!(true && false)) {
                    var v18 uint32 = 42
                    var v19 int = (v7 | v7)
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS1{F0: (false && (0.1 <= 1e10))}
                    v21 := GenS2{F0: 110, F1: 0}
                    fmt.Println(v14, v14, v14)
                    _ = v20
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                v23 := GenS1{F0: false}
                v7 = (v7 & (-(v7)))
                switch v1 {
                case "hello":
                    v1 = v1
                case "\n\t":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    fmt.Println(v7, v14, v1)
                }
                _ = fi22
                _ = v23
            }
        }
    } else {
        v24 := GenS1{F0: (!true)}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v24
    }
    if ("a" < ("" + (v1 + v1))) {
        v4[2147483648] = (v1 + v1)
    } else {
        _ = len(v4)
    }
    if len(v2) > 0 {
        v2[0] = (v7 ^ (v7 * v7))
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        fmt.Println(v1, fi25)
        v7 = -1
        v2 = append(v2, (v7 | (v7 << ((4294967295) & 63))))
        _ = fi25
    }
    var v26 int32 = -1
    v27 := GenS1{F0: (3.14159 < 3.14159)}
    v28 := GenS2{F0: 42, F1: 886}
    if len(v3) > 0 {
        v3[0] = v1
    }
    v14 += (^(v14))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
