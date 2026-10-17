package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int16; F1 int64; F2 int16 }

func main() {
    var v1 int16 = 1
    v2 := make(map[int]int)
    var v3 uint16 = 786
    if ("" > "foo bar") {
        switch ("hello" + "hello") {
        case "a":
            v3 = v3
        case "":
            v1 -= v1
        default:
            _ = len(v2)
        }
    }
    if false {
        v3 += (^(v3))
        v4 := GenS1{F0: 2147483648, F1: -72}
        for fi5 := 0; fi5 < 7; fi5++ {
            for fi6 := 0; fi6 < 18; fi6++ {
                if false {
                    if (707 == 755) {
                        switch ("foo bar" + "a") {
                        case "a":
                            v7 := GenS1{F0: 2147483648, F1: 1}
                            _ = v7
                        case "\n\t":
                            v1 -= (v1 + 1000)
                        case "foo bar":
                            if true {
                                v8 := []int{fi6}
                                v8 = append(v8, (fi6 + (fi5 + (fi6 ^ fi5))))
                                var v9 float64 = (float64(1e10) / 3.14159)
                                var v10 uint8 = 0
                                _ = v10
                                _ = v8
                                _ = v9
                            } else {
                                v1 += (^(v1))
                                v3 -= v3
                                if (!(1 <= 1)) {
                                    var v11 bool = (2 > 4294967295)
                                    _ = v11
                                }
                            }
                        }
                        var v12 int16 = v1
                        _ = v12
                    }
                    v13 := GenS1{F0: uint32(v3), F1: -1}
                    fmt.Println(fi5, v3)
                    _ = v13
                }
                _ = fi6
            }
            _ = fi5
        }
        _ = v4
    }
    v14 := []int{-1, 128, 32768, 0}
    switch "\n\t" {
    case "":
        var v15 int = 4294967295
        _ = v15
    case "hello":
        var v16 int16 = v1
        _ = v16
    }
    v3 += (v3 * (v3 ^ (v3 | v3)))
    v1 = v1
    v1 = v1
    fmt.Println(v1, v3)
    v17 := v2[569]
    v18 := GenS2{F0: 1000, F1: 107, F2: -32768}
    for fi19 := 0; fi19 < 18; fi19++ {
        switch ("\n\t" + "") {
        case "\n\t":
            v17 = 9223372036854775807
        default:
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(v3)
                _ = fi20
            }
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
