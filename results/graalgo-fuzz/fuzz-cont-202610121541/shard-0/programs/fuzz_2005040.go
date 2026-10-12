package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    v1 := []int{}
    v2 := []string{}
    var v3 uint64 = 18446744073709551615
    var v4 uint16 = 2
    var v5 uint16 = v4
    v6 := GenS1{F0: -1}
    fmt.Println(v3)
    fmt.Println(v4)
    v7 := GenS2{F0: 42, F1: -1}
    switch ("hello" + "") {
    case "a":
        v5 += (v4 >> ((57) & 63))
    }
    v2 = append(v2, ("foo bar" + ((("\n\t" + ("foo bar" + "\n\t")) + "hello") + "")))
    v4 = v4
    if false {
        v4 = (v5 * (v4 / (((v5 & (v5 & (v4 | 65535)))) | 1)))
        for fi8 := 0; fi8 < 14; fi8++ {
            if false {
                var v9 float64 = 1.5
                v10 := GenS2{F0: fi8, F1: -50}
                _ = v10
                _ = v9
            }
            _ = fi8
        }
    } else {
        v4 = 255
        switch -983 {
        case 632:
            var v11 uint8 = 42
            _ = v11
        case 9223372036854775807:
            if len(v2) > 0 {
                v2[0] = ("hello" + "\n\t")
            }
        case 1:
            if (-1 == -1) {
                if false {
                    v12 := GenS1{F0: 173}
                    v4 -= (v4 + (v4 % (((v4 ^ (v4 / (((v5 << ((164) & 63))) | 1)))) | 1)))
                    if true {
                        for fi13 := 0; fi13 < 18; fi13++ {
                            v4 -= v5
                            _ = fi13
                        }
                        for fi14 := 0; fi14 < 20; fi14++ {
                            v5 += 42
                            _ = fi14
                        }
                    }
                    _ = v12
                } else {
                    fmt.Println(v3, v5)
                    var v15 uint32 = 0
                    _ = v15
                }
                v3 = v3
                v16 := GenS1{F0: 267}
                _ = v16
            }
        default:
            v4 -= v5
        }
    }
    _ = len(v1)
    v5 = 0
    v3 += v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
