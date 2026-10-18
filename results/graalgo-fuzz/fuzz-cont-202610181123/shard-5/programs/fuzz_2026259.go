package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint8 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 2; fi7++ {
        if false {
            if (0 > 2147483647) {
                var v8 int32 = 550
                _ = v8
            }
            var v9 string = "\n\t"
            _ = v9
        }
        v10 := []int{}
        var v11 int32 = 252
        _ = fi7
        _ = v10
        _ = v11
    }
    if (-32768 > -524) {
        v12 := GenS1{F0: "foo bar", F1: ((-90 >= 42) && (913 < 554)), F2: 42}
        var v13 int8 = -55
        _ = v12
        _ = v13
    }
    switch -396 {
    case 1:
        var v14 string = ""
        _ = v14
    case 4294967295:
        switch "hello" {
        case "\n\t":
            v15 := GenS1{F0: "a", F1: false, F2: 0}
            _ = v15
        case "a":
            v16 := []int{127, -2147483648}
            _ = v16
        case "foo bar":
            v17 := GenS1{F0: "foo bar", F1: (!false), F2: -2147483648}
            _ = v17
        }
    case -9223372036854775808:
        v18 := []int{}
        _ = v18
    default:
        v19 := []int{}
        _ = v19
    }
    if (-832 >= -1) {
        v20 := []int{-70, 738}
        v21 := GenS2{F0: 401}
        _ = v20
        _ = v21
    } else {
        if true {
            var v22 int64 = 1
            v23 := GenS1{F0: "", F1: false, F2: 924}
            v22 = (v22 % ((9223372036854775807) | 1))
            _ = v22
            _ = v23
        }
    }
    return 42
}

func main() {
    var v1 uint64 = 208
    v2 := make(map[int]bool)
    var v3 int16 = 769
    var v4 uint16 = 255
    var v5 uint64 = v1
    v6 := []uint32{uint32(v4), 65535}
    var v24 float32 = (1.5 - -2.25)
    fmt.Println(v3)
    v2[9223372036854775807] = (864 <= 31)
    switch "" {
    case "\n\t":
        v24 += float32((float64((0.1 * -2.25)) + 0.0))
    }
    var v25 uint64 = 18446744073709551615
    v26 := GenS3{F0: 101, F1: (v3 % (((v3 << ((255) & 63))) | 1))}
    if (!false) {
        if false {
            v3 -= -996
        } else {
            v27 := GenS4{F0: (v4 - (v4 << ((18446744073709551615) & 63))), F1: v4, F2: 62}
            _ = len(v6)
            _ = v27
        }
        var v28 string = ("\n\t" + ((("" + "") + "\n\t") + "foo bar"))
        _ = v28
    }
    fmt.Println(v5)
    var v29 uint16 = (v4 << ((529) & 63))
    v5 = uint64(v4)
    v2[127] = (!((false && ((0.0 * 0.1) != (v24 * v24))) || ((float32((1e10 + -2.25)) + 0.1) > v24)))
    v5 -= (v25 * (v5 * uint64(v4)))
    if false {
        if false {
            var v30 bool = true
            _ = v30
        } else {
            v29 += v4
            switch 0 {
            case 32767:
                if true {
                    v31 := GenS1{F0: "\n\t", F1: true, F2: 42}
                    _ = v31
                }
            default:
                v32 := GenS2{F0: 707}
                _ = v32
            }
            var v33 float64 = 1.5
            _ = v33
        }
    } else {
        var v34 int64 = 9223372036854775807
        v24 -= float32(0.0)
        _ = v34
    }
    var v35 float64 = float64(1.5)
    v24 += 1.5
    fmt.Println(v29, v24)
    v3 = v3
    v29 -= (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
