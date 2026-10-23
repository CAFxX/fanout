package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int16 }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    var v1 uint = 65535
    var v2 uint8 = 144
    var v3 int16 = 1
    v4 := make(map[int]uint32)
    var v5 int16 = -32768
    for fi6 := 0; fi6 < 13; fi6++ {
        v4[-615] = 786
        v2 += (^(v2))
        _ = fi6
    }
    if (v1 <= v1) {
        v7 := GenS1{F0: 957, F1: (v3 * 32767), F2: (v3 - 1)}
        var v8 uint8 = 1
        var v9 int16 = (v5 - -741)
        _ = v7
        _ = v8
        _ = v9
    }
    if true {
        if (false || false) {
            var v10 uint16 = 0
            var v11 int32 = 1
            for fi12 := 0; fi12 < 20; fi12++ {
                fmt.Println(v2, v10, v10)
                if ((v11 & (^(v11))) < v11) {
                    _ = len(v4)
                    if (-9223372036854775808 <= 32767) {
                        var v13 float64 = float64((3.14159 + (-2.25 / 1.5)))
                        var v14 uint8 = 139
                        switch "foo bar" {
                        case "foo bar":
                            v4[(fi12 << ((1) & 63))] = 2147483648
                        case "\n\t":
                            var v15 int8 = 31
                            _ = v15
                        }
                        _ = v13
                        _ = v14
                    }
                    var v16 uint8 = (v2 % ((v2) | 1))
                    _ = v16
                }
                switch "\n\t" {
                case "foo bar":
                    var v17 uint16 = 255
                    _ = v17
                default:
                    var v18 uint16 = v10
                    _ = v18
                }
                _ = fi12
            }
            _ = v10
            _ = v11
        }
        v1 = 964
        v19 := GenS1{F0: 1, F1: (v3 + 0), F2: (^(v3))}
        _ = v19
    } else {
        if false {
            for fi20 := 0; fi20 < 11; fi20++ {
                v2 += 174
                _ = fi20
            }
            v3 -= -202
            for fi21 := 0; fi21 < 6; fi21++ {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 int16 = (v3 / ((v5) | 1))
                    _ = fi22
                    _ = v23
                }
                var v24 bool = true
                for fi25 := 0; fi25 < 7; fi25++ {
                    v26 := GenS1{F0: 0, F1: v5, F2: -1}
                    _ = fi25
                    _ = v26
                }
                _ = fi21
                _ = v24
            }
        }
        v1 -= (v1 + (v1 % (((v1 / (((^(v1))) | 1))) | 1)))
    }
    v5 = (v5 % (((v5 - v3)) | 1))
    var v27 string = ""
    var v28 int = 875
    v29 := []int{v28, (v28 >> (((v1 + v1)) & 63)), (v28 | v28), v28}
    v28 = 954
    var v30 float64 = 1.5
    v31 := GenS2{F0: "hello", F1: 127}
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 uint8 = 231
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
