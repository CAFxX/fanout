package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int8 = 0
    var v2 float32 = ((0.0 / ((1e10 * 1.5) * -2.25)) - 3.14159)
    v3 := make(map[string]bool)
    v4 := []uint32{4294967295, 1, 4294967295}
    var v5 float64 = 3.14159
    v6 := make(map[uint32]bool)
    switch 32767 {
    case -614:
        var v7 string = "foo bar"
        _ = v7
    case 128:
        for fi8 := 0; fi8 < 13; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                v1 += (-(v1))
                fmt.Println(fi9)
                _ = fi9
            }
            _ = fi8
        }
    case -396:
        fmt.Println(v2, v2, v5)
    default:
        for fi10 := 0; fi10 < 6; fi10++ {
            v4 = append(v4, 822)
            var v11 uint8 = 127
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v5 += (-2.25 * (3.14159 * (float64((v2 * 0.0)) - 0.0)))
        _ = fi12
    }
    if (!("\n\t" < "")) {
        v2 = float32(v5)
        if (42 < 2) {
            switch -2147483648 {
            case 42:
                var v13 uint8 = 255
                _ = v13
            default:
                var v14 uint32 = 721
                _ = v14
            }
            var v15 uint = 897
            if (!((v2 / v2) == v2)) {
                v16 := GenS1{F0: -1, F1: 1, F2: 547}
                var v17 uint16 = 592
                fmt.Println(v15, v17, v1)
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 uint64 = 18446744073709551615
                    if false {
                        for fi20 := 0; fi20 < 2; fi20++ {
                            v21 := GenS1{F0: 42, F1: -437, F2: 836}
                            v22 := GenS1{F0: -1, F1: -610, F2: 65535}
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            }
            _ = v15
        } else {
            var v23 int8 = v1
            switch 128 {
            case 9223372036854775807:
                var v24 uint32 = 65535
                _ = v24
            case -61:
                v3["a"] = (539 == -9223372036854775808)
            }
            _ = v23
        }
        switch -1 {
        case 612:
            v2 -= v2
        }
    }
    if false {
        for fi25 := 0; fi25 < 17; fi25++ {
            for fi26 := 0; fi26 < 13; fi26++ {
                switch 473 {
                case 390:
                    v2 += float32(v5)
                case 2147483647:
                    if len(v4) > 0 {
                        v4[0] = 610
                    }
                case 603:
                    v1 = (v1 % ((v1) | 1))
                }
                delete(v3, "\n\t")
                var v27 int32 = 831
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    } else {
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 uint16 = 42
            _ = fi28
            _ = v29
        }
        v1 -= -1
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 int8 = v1
        _ = fi30
        _ = v31
    }
    _ = len(v4)
    for fi32 := 0; fi32 < 6; fi32++ {
        v33 := GenS1{F0: -152, F1: 42, F2: 32767}
        if true {
            v34 := GenS1{F0: 811, F1: -434, F2: 255}
            v1 = (v1 << ((662) & 63))
            _ = v34
        }
        _ = fi32
        _ = v33
    }
    v2 = 0.1
    v35 := GenS1{F0: 783, F1: -1, F2: 255}
    if false {
        if (42 == -965) {
            var v36 uint8 = 42
            _ = len(v6)
            _ = v36
        } else {
            var v37 uint8 = 224
            _ = v37
        }
        var v38 float32 = float32(v5)
        _ = v38
    } else {
        if (987 >= 255) {
            v2 -= ((0.0 / (v2 * 0.1)) / (1e10 - (0.1 * v2)))
        }
        v39 := GenS2{F0: 42}
        var v40 string = ("foo bar" + ("a" + "\n\t"))
        _ = v39
        _ = v40
    }
    v1 -= -1
    v1 = v1
    switch "\n\t" {
    case "a":
        for fi41 := 0; fi41 < 14; fi41++ {
            v42 := GenS2{F0: 255}
            _ = fi41
            _ = v42
        }
    case "\n\t":
        switch -9223372036854775808 {
        case 735:
            v3[((("a" + ("\n\t" + "")) + "") + "hello")] = (2 != 42)
        }
    }
    switch 351 {
    case 128:
        v2 = v2
    case -9223372036854775808:
        v5 = 0.1
    default:
        var v43 int32 = -2147483648
        _ = v43
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        var v45 uint64 = 18446744073709551615
        var v46 float32 = (1e10 - v2)
        _ = fi44
        _ = v45
        _ = v46
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
