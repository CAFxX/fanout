package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{32768}
        _ = fi6
        _ = v7
    }
    var v8 uint = 851
    fmt.Println(v8)
    _ = v8
    return 227
}
func genFn1() int8 {
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(fi9)
        v10 := []int{fi9, (fi9 * fi9), fi9}
        _ = fi9
        _ = v10
    }
    var v11 string = "foo bar"
    _ = v11
    return -128
}

func main() {
    var v1 uint32 = 808
    var v2 int16 = -146
    var v3 uint16 = 155
    var v4 string = "foo bar"
    v5 := []uint32{}
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(v3, v1, v4)
        _ = fi12
    }
    v13 := GenS1{F0: (false && false)}
    var v14 uint32 = 0
    fmt.Println(v1)
    switch "foo bar" {
    case "\n\t":
        v3 = 108
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(fi15)
        v16 := GenS1{F0: true}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint8 = 42
        _ = fi17
        _ = v18
    }
    if false {
        var v19 int64 = -1
        fmt.Println(v3, v19)
        _ = v19
    } else {
        var v20 uint32 = 0
        _ = v20
    }
    switch 209 {
    case 1:
        switch 787 {
        case 4294967295:
            var v21 float64 = -2.25
            _ = v21
        case 32768:
            for fi22 := 0; fi22 < 4; fi22++ {
                switch (v4 + (v4 + (v4 + v4))) {
                case "":
                    if false {
                        var v23 int16 = (v2 - v2)
                        _ = len(v5)
                        v2 = v2
                        _ = v23
                    } else {
                        v5 = append(v5, 579)
                        var v24 bool = true
                        switch v4 {
                        case "a":
                            v25 := GenS1{F0: false}
                            _ = v25
                        case "foo bar":
                            v26 := GenS2{F0: (v2 >= 356), F1: -52, F2: -978}
                            _ = v26
                        default:
                            v1 += (^(v14))
                        }
                        _ = v24
                    }
                default:
                    v1 += v14
                }
                _ = fi22
            }
        case -2147483648:
            if (false || false) {
                if (false || false) {
                    var v27 string = ""
                    for fi28 := 0; fi28 < 15; fi28++ {
                        var v29 uint64 = 4294967295
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                } else {
                    var v30 string = v4
                    _ = v30
                }
                v5 = append(v5, v1)
            }
        }
    default:
        v1 -= v1
    }
    v4 = v4
    v31 := GenS1{F0: (!true)}
    fmt.Println(v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
