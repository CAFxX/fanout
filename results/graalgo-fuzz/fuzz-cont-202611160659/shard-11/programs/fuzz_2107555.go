package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint32 = 1
    var v2 int = 128
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    v5 := make(map[int]string)
    var v6 float64 = 0.1
    v7 := v5[(v2 + (v2 & (v2 & v2)))]
    v2 = v2
    var v8 uint = 0
    v6 = v6
    var v9 string = "hello"
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 float64 = 3.14159
        _ = fi10
        _ = v11
    }
    switch "hello" {
    case "foo bar":
        v12 := GenS1{F0: (v2 ^ v2)}
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int16 = -1
            v7 = (v9 + v9)
            _ = fi13
            _ = v14
        }
    case "hello":
        v15 := []int{(v2 << ((v8) & 63)), (v2 + (v2 + (v2 ^ v2))), v2, v2}
        _ = v15
    default:
        switch "hello" {
        case "a":
            if ((v2 << ((v8) & 63)) != v2) {
                v16 := GenS2{F0: 157, F1: v2}
                switch (-(v2)) {
                case 0:
                    if ("hello" > (v7 + "hello")) {
                        fmt.Println(v8, v1, v9)
                    } else {
                        v17 := GenS2{F0: 1, F1: v2}
                        v2 += (-(v2))
                        if (42 != 9223372036854775807) {
                            v1 = 0
                            v18 := GenS3{F0: 255}
                            _ = v18
                        } else {
                            v8 -= (v8 + (^(v8)))
                            if true {
                                v3[(v7 + v9)] = (v9 + "")
                            }
                        }
                        _ = v17
                    }
                case 2147483647:
                    for fi19 := 0; fi19 < 16; fi19++ {
                        for fi20 := 0; fi20 < 1; fi20++ {
                            var v21 int8 = 1
                            v1 = (v1 | 4294967295)
                            v2 -= fi20
                            _ = fi20
                            _ = v21
                        }
                        _ = fi19
                    }
                }
                _ = v16
            } else {
                fmt.Println(v6, v6)
                if (float64(1e10) <= ((v6 + float64(1.5)) + v6)) {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 bool = true
                        _ = fi22
                        _ = v23
                    }
                } else {
                    v1 = (v1 * v1)
                    v9 = "hello"
                }
                v24 := GenS1{F0: v2}
                _ = v24
            }
        case "":
            if (v9 <= (v7 + v7)) {
                v25 := GenS1{F0: -9223372036854775808}
                _ = v25
            } else {
                v2 = v2
                fmt.Println(v1, v8, v7)
            }
        case "foo bar":
            v26 := GenS3{F0: 64}
            _ = v26
        }
    }
    var v27 bool = (!(-614 <= 2147483647))
    fmt.Println(v27, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
