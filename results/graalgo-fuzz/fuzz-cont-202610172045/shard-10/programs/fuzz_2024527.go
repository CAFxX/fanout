package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 string; F1 string; F2 int64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := []int{-641}
    v3 := []int{398}
    v4 := make(map[int]bool)
    var v5 uint32 = 2147483647
    v6 := GenS1{F0: 9223372036854775807, F1: v5}
    var v7 uint8 = 2
    v8 := GenS2{F0: "foo bar", F1: "a", F2: -1}
    for fi9 := 0; fi9 < 15; fi9++ {
        v1 = (v1 << ((2) & 63))
        switch ("foo bar" + "a") {
        case "":
            switch fi9 {
            case 127:
                v1 = (^(v1))
            }
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = len(v4)
        _ = fi9
    }
    v7 -= v7
    for fi10 := 0; fi10 < 11; fi10++ {
        v1 = v1
        if (127 <= 142) {
            v5 = 42
            v1 = v1
            v11 := GenS2{F0: "hello", F1: "a", F2: 9223372036854775807}
            _ = v11
        } else {
            _ = len(v2)
            if len(v2) > 0 {
                v2[0] = (fi10 % ((fi10) | 1))
            }
            if (!true) {
                if (38 <= -1) {
                    v5 = (v5 * v5)
                } else {
                    v5 = (v5 & v5)
                }
            } else {
                var v12 int8 = -128
                if (!true) {
                    var v13 bool = (!(0.1 == ((3.14159 / 1.5) - 1e10)))
                    _ = v13
                } else {
                    v5 = 2147483648
                    if true {
                        var v14 uint64 = (^(v1))
                        _ = v14
                    } else {
                        var v15 bool = ("foo bar" == "\n\t")
                        _ = v15
                    }
                }
                if (!(2 == 897)) {
                    switch "foo bar" {
                    case "\n\t":
                        if ("" > ("foo bar" + "hello")) {
                            if true {
                                for fi16 := 0; fi16 < 4; fi16++ {
                                    var v17 int = (fi16 / (((fi16 >> ((255) & 63))) | 1))
                                    var v18 uint = 81
                                    v1 -= 797
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                            } else {
                                v5 += (v5 + v5)
                                for fi19 := 0; fi19 < 4; fi19++ {
                                    v1 = v1
                                    _ = fi19
                                }
                            }
                            for fi20 := 0; fi20 < 1; fi20++ {
                                var v21 int = 128
                                if true {
                                    v21 = -276
                                }
                                _ = fi20
                                _ = v21
                            }
                        }
                    case "hello":
                        var v22 bool = ((((1 == v5) && (2 > 255)) && (740 < -1)) || ((1e10 / ((((-2.25 / float32((((3.14159 * 1e10) + (float64((3.14159 * -2.25)) + 0.0)) / -2.25))) * (3.14159 * (1e10 / 1.0))) - 1.5) / (((0.0 - 0.0) - ((1e10 / -2.25) / -2.25)) - 0.0))) > 3.14159))
                        _ = v22
                    }
                    if len(v3) > 0 {
                        v3[0] = (fi10 >> ((255) & 63))
                    }
                    for fi23 := 0; fi23 < 19; fi23++ {
                        var v24 uint64 = v1
                        _ = fi23
                        _ = v24
                    }
                }
                _ = v12
            }
        }
        _ = fi10
    }
    var v25 uint16 = 927
    v2 = append(v2, 128)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
