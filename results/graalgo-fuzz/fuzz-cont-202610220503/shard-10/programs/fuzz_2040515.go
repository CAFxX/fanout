package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS4 struct { F0 string; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint; F1 int16; F2 uint32 }
func genFn0(p0 uint8, p1 int, p2 int16) uint8 {
    v7 := []int{-715, 4294967295, p1, 747}
    if true {
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: p2, F1: 911, F2: 1}
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 int32 = -2147483648
    v4 := []int{1, -9223372036854775808, 203, -481}
    v5 := []uint32{65535, 42, 0}
    var v6 bool = false
    v3 -= v3
    v3 += v3
    if true {
        switch "foo bar" {
        case "hello":
            var v10 int = 520
            _ = v10
        case "":
            if v6 {
                if (155 != 128) {
                    if ("a" > "foo bar") {
                        fmt.Println(v3)
                        v11 := GenS2{F0: 968}
                        _ = v11
                    }
                } else {
                    v6 = ((!v6) || false)
                }
                fmt.Println(v3, v3)
            }
        case "foo bar":
            _ = len(v4)
        }
        v12 := GenS3{F0: 1, F1: 0, F2: 814}
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v6 = (!v6)
        _ = fi13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v3 -= (v3 << ((255) & 63))
        fmt.Println(fi14)
        _ = fi14
    }
    v15 := GenS4{F0: (("hello" + "foo bar") + ("" + "foo bar")), F1: -99, F2: 622}
    for fi16 := 0; fi16 < 20; fi16++ {
        v3 += 42
        _ = fi16
    }
    fmt.Println(v3)
    for fi17 := 0; fi17 < 4; fi17++ {
        v6 = v6
        _ = fi17
    }
    if (!v6) {
        var v18 int32 = v3
        if true {
            v19 := v2[0]
            if len(v5) > 0 {
                v5[0] = 964
            }
            switch "" {
            case "":
                v20 := GenS2{F0: -249}
                _ = v20
            default:
                var v21 int8 = -89
                _ = v21
            }
            _ = v19
        }
        v3 = (v18 / (((v3 % ((42) | 1))) | 1))
        _ = v18
    } else {
        v3 -= 0
    }
    switch 358 {
    case 957:
        var v22 string = (("hello" + ("a" + "hello")) + ("a" + "foo bar"))
        _ = v22
    default:
        if (v6 && (v6 || false)) {
            v23 := v1[417]
            v6 = v6
            v24 := GenS3{F0: 560, F1: v3, F2: 210}
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 16; fi25++ {
                if v6 {
                    fmt.Println(fi25, v3, fi25)
                    v1[470] = "foo bar"
                }
                _ = fi25
            }
        }
    }
    var v26 int16 = -1
    fmt.Println(v6)
    var v27 int = -1
    var v28 float32 = 0.1
    v28 += v28
    switch "" {
    case "a":
        var v29 uint32 = 635
        _ = v29
    case "":
        var v30 int32 = (v3 | int32(v27))
        _ = v30
    default:
        if (float32(1e10) > 0.0) {
            switch ("a" + "foo bar") {
            case "":
                fmt.Println(v6)
            case "foo bar":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v28 = v28
        }
    }
    v31 := GenS5{F0: 65535, F1: (v26 | (^(v26))), F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
