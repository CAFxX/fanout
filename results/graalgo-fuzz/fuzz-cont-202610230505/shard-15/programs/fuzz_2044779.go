package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0(p0 int16, p1 int32, p2 uint8) int32 {
    var v4 uint = 1
    var v5 uint8 = p2
    p2 = p2
    var v6 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (p1 * (^(p1)))
}

func main() {
    var v1 uint16 = 0
    var v2 uint8 = 192
    var v3 uint = 629
    v7 := []int{4294967295, 0, -710, 4294967295}
    var v8 int8 = 127
    switch "foo bar" {
    case "hello":
        v3 = (v3 - (v3 - (^(v3))))
    case "foo bar":
        v8 = (v8 >> ((uint(v2)) & 63))
    case "\n\t":
        var v9 uint8 = (v2 + 42)
        _ = v9
    default:
        var v10 int8 = (^(v8))
        _ = v10
    }
    v11 := GenS1{F0: false, F1: (v1 + v1)}
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            switch (("" + ((("a" + ("" + "\n\t")) + "foo bar") + "\n\t")) + ("a" + ("a" + "foo bar"))) {
            case "hello":
                switch fi12 {
                case 1:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    switch (fi12 * (fi12 & -9223372036854775808)) {
                    case -16:
                        var v14 string = "\n\t"
                        _ = v14
                    case 42:
                        if len(v7) > 0 {
                            v7[0] = -65
                        }
                        v15 := GenS1{F0: (false && true), F1: 492}
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        _ = v15
                    default:
                        switch fi13 {
                        case -973:
                            var v16 uint64 = 18446744073709551615
                            _ = v16
                        }
                    }
                default:
                    fmt.Println(fi12)
                }
            }
            _ = fi13
        }
        _ = fi12
    }
    fmt.Println(v1, v8, v8)
    var v17 float32 = 0.0
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int32 = 5
        if ((0.0 == ((1e10 + (((1.5 - 3.14159) * (1e10 * 1.5)) / -2.25)) / 1.0)) || (!(!true))) {
            if true {
                v20 := GenS1{F0: false, F1: v1}
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v20
            } else {
                switch (fi18 >> (((v3 % ((1) | 1))) & 63)) {
                case -9223372036854775808:
                    v2 = 0
                case 128:
                    var v21 uint64 = 0
                    _ = v21
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            v1 = (v1 / ((v1) | 1))
        }
        _ = fi18
        _ = v19
    }
    v8 += v8
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 uint64 = uint64(v2)
        if (874 >= 0) {
            var v24 uint64 = v23
            _ = len(v7)
            if len(v7) > 0 {
                v7[0] = fi22
            }
            _ = v24
        } else {
            fmt.Println(fi22, v17, v3)
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint = uint(v2)
            v17 = ((v17 + (v17 / (float32((0.1 * 0.1)) * v17))) - v17)
            v27 := GenS1{F0: false, F1: (v1 - 32767)}
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi22
        _ = v23
    }
    if len(v7) > 0 {
        v7[0] = -97
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v17, v8)
    case "":
        if true {
            v7 = append(v7, 0)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        } else {
            var v28 bool = false
            v28 = v28
            _ = v28
        }
    case "\n\t":
        v2 -= (v2 + uint8(v3))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
