package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    v4 := []int{969}
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint = 559
        _ = fi5
        _ = v6
    }
    fmt.Println(v1)
    v7 := GenS1{F0: 2147483647}
    fmt.Println(v1)
    switch "a" {
    case "\n\t":
        switch -9223372036854775808 {
        case 528:
            for fi8 := 0; fi8 < 1; fi8++ {
                switch -9223372036854775808 {
                case -167:
                    var v9 bool = true
                    _ = v9
                case 9223372036854775807:
                    if (!(2 >= 154)) {
                        switch ("a" + "a") {
                        case "\n\t":
                            if len(v4) > 0 {
                                v4[0] = (fi8 / (((fi8 & fi8)) | 1))
                            }
                        }
                        var v10 uint8 = (v1 % ((v1) | 1))
                        _ = v10
                    }
                default:
                    v11 := GenS1{F0: 309}
                    _ = v11
                }
                _ = fi8
            }
        case -9223372036854775808:
            for fi12 := 0; fi12 < 6; fi12++ {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 int16 = 1
                    v15 := GenS1{F0: 296}
                    var v16 float32 = 0.0
                    v16 += 3.14159
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                v1 = (v1 + (v1 | v1))
                fmt.Println(v1)
                _ = fi12
            }
        }
    case "foo bar":
        if false {
            v1 = v1
        }
    default:
        var v17 float32 = -2.25
        _ = v17
    }
    v18 := GenS2{F0: -128}
    switch 32767 {
    case -407:
        var v19 float64 = 3.14159
        _ = v19
    default:
        v3[0] = 788
    }
    var v20 float64 = 0.0
    if (!((!(true && true)) && (!((v1 * v1) != 1)))) {
        var v21 uint16 = 255
        _ = v21
    }
    v1 = (v1 + (v1 - 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
}
