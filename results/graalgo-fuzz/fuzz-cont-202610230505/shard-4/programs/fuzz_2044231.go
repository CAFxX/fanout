package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int64; F2 int }

func main() {
    v1 := make(map[int]int)
    var v2 string = ""
    var v3 int64 = 42
    v4 := make(map[uint32]int)
    var v5 uint32 = 356
    v6 := v1[9223372036854775807]
    var v7 bool = false
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int8 = 127
        _ = fi8
        _ = v9
    }
    if v7 {
        switch v2 {
        case "a":
            if true {
                _ = len(v4)
                if v7 {
                    v7 = ((v5 & v5) < 42)
                }
            }
        default:
            if (1 > 334) {
                switch v2 {
                case "\n\t":
                    v10 := GenS1{F0: v7}
                    _ = v10
                }
            } else {
                switch int(v3) {
                case -431:
                    switch "\n\t" {
                    case "":
                        v5 += (v5 >> ((2) & 63))
                    case "\n\t":
                        for fi11 := 0; fi11 < 6; fi11++ {
                            v12 := []int{9223372036854775807}
                            for fi13 := 0; fi13 < 2; fi13++ {
                                v14 := v1[fi13]
                                var v15 uint8 = 127
                                _ = len(v12)
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                            _ = fi11
                            _ = v12
                        }
                    }
                default:
                    switch (v6 & (v6 << ((18446744073709551615) & 63))) {
                    case 42:
                        v7 = true
                    }
                }
                switch 0 {
                case 4294967295:
                    switch int(v3) {
                    case 127:
                        fmt.Println(v7)
                    }
                case 42:
                    v16 := GenS2{F0: -128, F1: 42, F2: (v6 | (-(v6)))}
                    _ = v16
                case 0:
                    _ = len(v1)
                default:
                    delete(v4, (^(v5)))
                }
                var v17 bool = v7
                _ = v17
            }
        }
    } else {
        var v18 uint = 18446744073709551615
        if v7 {
            v19 := v4[v5]
            delete(v1, int(v3))
            _ = v19
        }
        v18 += (v18 | v18)
        _ = v18
    }
    var v20 int8 = int8(v6)
    v1[(v6 & v6)] = (^(v6))
    v2 = (((("foo bar" + v2) + v2) + v2) + v2)
    var v21 uint32 = 4294967295
    v22 := GenS1{F0: v7}
    if v7 {
        delete(v4, 2147483647)
        v20 = v20
    } else {
        v5 += v21
        if ((0.1 != 0.0) && true) {
            if v7 {
                fmt.Println(v7, v5, v7)
                v21 = (^(v21))
            } else {
                v20 = -1
                fmt.Println(v21, v2, v21)
            }
        } else {
            switch -70 {
            case -932:
                switch v2 {
                case "hello":
                    v2 = "hello"
                default:
                    v7 = ((v3 % ((1) | 1)) > (v3 ^ (v3 + v3)))
                }
            case -697:
                if false {
                    fmt.Println(v3, v5)
                }
            default:
                v20 = v20
            }
        }
        v6 += 9223372036854775807
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
