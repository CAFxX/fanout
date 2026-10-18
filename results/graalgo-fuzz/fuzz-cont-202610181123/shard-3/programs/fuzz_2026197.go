package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 int8 }
func genFn0() int32 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint64 = 116
        if (!(!false)) {
            var v8 uint16 = 461
            v9 := GenS1{F0: false, F1: 1, F2: 0}
            var v10 bool = true
            _ = v10
            _ = v8
            _ = v9
        } else {
            if true {
                var v11 int16 = -32768
                _ = v11
            }
        }
        _ = fi6
        _ = v7
    }
    var v12 uint32 = 42
    v12 += (v12 ^ v12)
    _ = v12
    return 1
}
func genFn1(p0 uint64, p1 uint32, p2 uint32) uint64 {
    v13 := GenS1{F0: true, F1: 334, F2: -2147483648}
    fmt.Println(p2, p1, p2)
    var v14 int16 = -437
    var v15 int64 = 42
    var v16 float32 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 int64 = -625
    v5 := make(map[string]bool)
    switch 127 {
    case 32768:
        var v17 float32 = (((float32(0.0) + 0.1) / (3.14159 + ((1.5 * ((((1.5 / (1e10 - float32((3.14159 - -2.25)))) * float32((-2.25 - 1e10))) * float32(1.5)) / 3.14159)) - 0.1))) - 3.14159)
        _ = v17
    default:
        var v18 uint8 = 127
        _ = v18
    }
    var v19 float32 = 3.14159
    fmt.Println(v4, v19)
    var v20 uint16 = 750
    if (18446744073709551615 == genFn1(756, 963, 2)) {
        switch -2147483648 {
        case 32768:
            v4 = 5
        case 2147483647:
            v4 += (-(v4))
        case -9223372036854775808:
            if false {
                var v21 float32 = v19
                _ = v21
            } else {
                var v22 bool = false
                var v23 uint16 = v20
                if v22 {
                    v22 = true
                }
                _ = v22
                _ = v23
            }
        default:
            switch ("foo bar" + "hello") {
            case "a":
                if ((708 != 1) && ((float32(1e10) - (float32(-2.25) + ((((v19 + v19) * (0.0 + 0.1)) / 0.1) / v19))) >= (((0.1 / (v19 * v19)) * v19) / v19))) {
                    v24 := GenS2{F0: -124}
                    for fi25 := 0; fi25 < 2; fi25++ {
                        switch "a" {
                        case "foo bar":
                            v26 := []int{}
                            _ = v26
                        case "\n\t":
                            var v27 int = (fi25 << ((387) & 63))
                            _ = v27
                        case "a":
                            var v28 int = (fi25 | 42)
                            _ = v28
                        }
                        _ = fi25
                    }
                    _ = v24
                }
            case "foo bar":
                for fi29 := 0; fi29 < 16; fi29++ {
                    v30 := []int{(fi29 % (((fi29 & fi29)) | 1)), (fi29 ^ fi29)}
                    var v31 uint16 = (v20 / ((v20) | 1))
                    v32 := GenS1{F0: (!false), F1: 4294967295, F2: 0}
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
            default:
                v4 = 258
            }
        }
        var v33 int32 = -2147483648
        _ = v33
    }
    v34 := []int{-409, 9223372036854775807}
    var v35 bool = true
    v36 := GenS1{F0: (32767 < 0), F1: 671, F2: 42}
    var v37 int8 = -1
    v38 := GenS2{F0: -84}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
