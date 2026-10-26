package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint64; F2 bool }
func genFn0() int32 {
    switch "hello" {
    case "a":
        var v4 int32 = -2147483648
        _ = v4
    case "\n\t":
        if (!false) {
            if (!(false && true)) {
                var v5 uint8 = 128
                var v6 uint8 = v5
                if (1 != 348) {
                    fmt.Println(v5, v5)
                    if (-454 == 1000) {
                        v7 := []int{}
                        _ = v7
                    }
                    fmt.Println(v6)
                }
                _ = v5
                _ = v6
            } else {
                switch 32768 {
                case -1:
                    for fi8 := 0; fi8 < 15; fi8++ {
                        var v9 uint32 = 1
                        v10 := GenS1{F0: -1, F1: (((("\n\t" + "\n\t") + "") + "") + "foo bar")}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                case 32767:
                    v11 := GenS2{F0: 1}
                    _ = v11
                }
            }
            switch ("" + ("\n\t" + (("\n\t" + ("\n\t" + "hello")) + "a"))) {
            case "foo bar":
                switch 32767 {
                case 128:
                    switch -2147483648 {
                    case -845:
                        var v12 uint8 = 255
                        _ = v12
                    }
                case 127:
                    var v13 uint8 = 128
                    _ = v13
                case 207:
                    var v14 int64 = -638
                    _ = v14
                }
            case "hello":
                var v15 uint64 = 0
                _ = v15
            default:
                if (false && (-2147483648 != -309)) {
                    v16 := GenS3{F0: -1, F1: 938, F2: -32768}
                    var v17 int16 = 293
                    _ = v16
                    _ = v17
                }
            }
        } else {
            var v18 bool = (!(-2.25 != 3.14159))
            _ = v18
        }
    case "foo bar":
        v19 := []int{4294967295, -1, 1}
        _ = v19
    }
    if false {
        v20 := GenS4{F0: 876, F1: 2, F2: (false || (!false))}
        _ = v20
    } else {
        var v21 int16 = 1000
        _ = v21
    }
    var v22 bool = (-679 != -1)
    switch 581 {
    case 32767:
        v22 = ((255 >= 508) || v22)
    }
    _ = v22
    return 42
}

func main() {
    var v1 int32 = 754
    v2 := []string{"a"}
    v3 := make(map[uint32]bool)
    v1 = v1
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 -= (^(v1))
    v1 -= (v1 - genFn0())
    for fi23 := 0; fi23 < 7; fi23++ {
        v3[42] = (!(false || ("" != "foo bar")))
        if false {
            v3[42] = (false && false)
        }
        _ = fi23
    }
    v1 -= (v1 * (v1 >> ((0) & 63)))
    v1 += genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
