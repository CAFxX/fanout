package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = -18
    v3 := make(map[int]uint32)
    var v4 uint8 = 42
    v4 = v4
    var v5 bool = ((0.0 / 0.1) != (0.1 - 0.1))
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            switch 9223372036854775807 {
            case 1:
                if true {
                    v4 += v4
                    v8 := []int{}
                    v9 := GenS1{F0: 192, F1: -110}
                    _ = v8
                    _ = v9
                } else {
                    if v5 {
                        v10 := []int{fi7, fi7, (fi6 % ((9223372036854775807) | 1))}
                        v11 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
                        switch fi7 {
                        case 2147483647:
                            v3[fi7] = 2
                        case 42:
                            if ((!(((v4 >> ((18446744073709551615) & 63)) < v4) && v5)) || ((!(!v5)) && v5)) {
                                v12 := v1[0]
                                _ = v12
                            }
                        case 396:
                            var v13 uint = 0
                            _ = v13
                        }
                        _ = v10
                        _ = v11
                    }
                }
            case 821:
                var v14 int8 = 100
                _ = v14
            case 2147483647:
                v5 = (-425 > 437)
            }
            fmt.Println(v5)
            _ = fi7
        }
        _ = fi6
    }
    var v15 int16 = -32768
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS1{F0: 1000, F1: 1}
            v15 += 42
            v19 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807}
            if v5 {
                fmt.Println(v15, fi16)
                v15 += v15
                switch ("" + "foo bar") {
                case "foo bar":
                    var v20 uint8 = v4
                    _ = v20
                case "":
                    v21 := GenS2{F0: v4}
                    switch (("\n\t" + ("\n\t" + ("hello" + ("" + "foo bar")))) + "foo bar") {
                    case "":
                        v22 := []int{128}
                        _ = v22
                    case "a":
                        switch "a" {
                        case "a":
                            v23 := GenS2{F0: (v4 & v4)}
                            _ = v23
                        default:
                            v4 = v4
                        }
                    }
                    _ = v21
                default:
                    v4 -= (v4 + 1)
                    var v24 uint16 = 255
                    _ = v24
                }
            } else {
                v25 := GenS1{F0: 68, F1: 9223372036854775807}
                var v26 uint16 = 32767
                v27 := v1[4294967295]
                if false {
                    var v28 uint32 = v27
                    _ = v28
                }
                v26 += (^(v26))
                v29 := GenS2{F0: uint8(v27)}
                _ = v25
                _ = v26
                _ = v27
                _ = v29
            }
            fmt.Println(v15, v15)
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    fmt.Println(v5, v4, v4)
    v30 := []int{-682, 32768, -324}
    v5 = (!(!(-708 <= 157)))
    if v5 {
        v15 -= v15
        v2 = (^(v2))
    }
    v4 -= 2
    if len(v30) > 0 {
        v30[0] = -502
    }
    v30 = append(v30, 32768)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
