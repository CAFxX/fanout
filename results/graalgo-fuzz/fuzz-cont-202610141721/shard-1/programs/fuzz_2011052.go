package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 uint16 }
func genFn0() uint16 {
    switch 4294967295 {
    case 127:
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int16 = -896
            v7 += v7
            _ = fi6
            _ = v7
        }
    case 2147483647:
        var v8 uint64 = 18446744073709551615
        _ = v8
    case 4294967295:
        v9 := []int{}
        _ = v9
    default:
        var v10 bool = (false && (false && (42 >= 255)))
        _ = v10
    }
    switch "" {
    case "foo bar":
        var v11 uint = 1
        _ = v11
    case "":
        if true {
            var v12 uint64 = 108
            _ = v12
        }
    }
    v13 := GenS1{F0: 2, F1: 0, F2: 255}
    _ = v13
    return 65535
}

func main() {
    var v1 uint8 = 22
    var v2 uint64 = 553
    var v3 uint32 = 4294967295
    var v4 int32 = 246
    var v5 uint32 = v3
    var v14 uint64 = v2
    switch "foo bar" {
    case "hello":
        if false {
            if ((!((54 > 32768) && false)) && false) {
                if true {
                    var v15 int = -1
                    for fi16 := 0; fi16 < 18; fi16++ {
                        switch "hello" {
                        case "hello":
                            v17 := []int{(fi16 ^ (fi16 >> ((4294967295) & 63))), v15, v15}
                            _ = v17
                        case "foo bar":
                            for fi18 := 0; fi18 < 7; fi18++ {
                                var v19 float64 = ((0.0 / float64(1.5)) + 0.1)
                                var v20 int8 = 100
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                        default:
                            v1 -= (v1 * v1)
                        }
                        _ = fi16
                    }
                    for fi21 := 0; fi21 < 4; fi21++ {
                        v22 := GenS1{F0: 255, F1: -81, F2: (v1 * v1)}
                        _ = fi21
                        _ = v22
                    }
                    _ = v15
                } else {
                    v14 += (v14 ^ v2)
                    var v23 uint = 770
                    v2 = v14
                    _ = v23
                }
                var v24 int32 = (v4 & (v4 & -2147483648))
                v14 += v2
                _ = v24
            }
            v14 = v14
        }
    case "\n\t":
        var v25 int = -2147483648
        _ = v25
    }
    v1 = (v1 ^ (v1 << ((130) & 63)))
    if false {
        v26 := []int{894, 426}
        v1 = (v1 - 2)
        var v27 int16 = 32767
        _ = v26
        _ = v27
    }
    if (!true) {
        var v28 uint16 = 1
        _ = v28
    }
    switch "" {
    case "":
        var v29 int8 = -51
        _ = v29
    case "hello":
        fmt.Println(v4, v3, v5)
    case "a":
        v2 -= v2
    default:
        v30 := GenS2{F0: 1, F1: -1, F2: 1}
        _ = v30
    }
    fmt.Println(v2, v1, v1)
    v2 -= 214
    v4 = -1
    switch 32768 {
    case 32767:
        v31 := GenS2{F0: 902, F1: 84, F2: 642}
        _ = v31
    case 2147483647:
        v3 -= (v5 ^ uint32(v1))
    case 32768:
        v4 += v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
