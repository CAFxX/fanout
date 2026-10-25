package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    var v3 bool = (42 <= 767)
    var v4 float32 = -2.25
    v5 := make(map[uint32]string)
    v6 := make(map[string]bool)
    var v7 int8 = 76
    for fi8 := 0; fi8 < 7; fi8++ {
        v1 = append(v1, (fi8 + fi8))
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(v4, v4)
            _ = fi9
        }
        _ = fi8
    }
    if true {
        v10 := GenS1{F0: 42}
        fmt.Println(v3, v3)
        fmt.Println(v4, v7, v7)
        _ = v10
    } else {
        switch 32768 {
        case 1:
            v7 -= (v7 >> ((739) & 63))
        default:
            for fi11 := 0; fi11 < 5; fi11++ {
                if (v3 && v3) {
                    switch "hello" {
                    case "":
                        v7 += (v7 + (v7 & (v7 - 17)))
                    case "a":
                        var v12 string = "hello"
                        _ = v12
                    default:
                        switch (fi11 % (((fi11 % (((fi11 >> ((0) & 63))) | 1))) | 1)) {
                        case 127:
                            if v3 {
                                var v13 uint8 = 2
                                _ = v13
                            }
                        case 703:
                            v4 = v4
                        }
                    }
                    var v14 string = ""
                    var v15 uint = 18446744073709551615
                    _ = v14
                    _ = v15
                } else {
                    switch "" {
                    case "a":
                        for fi16 := 0; fi16 < 8; fi16++ {
                            _ = len(v1)
                            v17 := GenS2{F0: true, F1: (-(fi16))}
                            _ = fi16
                            _ = v17
                        }
                    case "foo bar":
                        v18 := GenS1{F0: 9223372036854775807}
                        _ = v18
                    default:
                        var v19 uint32 = 1
                        _ = v19
                    }
                    v20 := v2[2147483648]
                    _ = v20
                }
                _ = fi11
            }
        }
        var v21 int64 = 42
        _ = v21
    }
    v3 = v3
    var v22 uint64 = 960
    var v23 uint = 42
    v24 := GenS2{F0: (-1 >= 4294967295), F1: 127}
    v25 := GenS3{F0: 32767}
    v26 := GenS2{F0: (672 <= 42), F1: -1}
    switch ("hello" + "") {
    case "hello":
        var v27 float64 = (0.0 + -2.25)
        _ = v27
    case "\n\t":
        v7 -= (v7 + -1)
    }
    v7 -= (-(v7))
    var v28 uint32 = 674
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
