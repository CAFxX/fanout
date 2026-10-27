package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 uint8; F1 uint8; F2 int }
func genFn0() uint32 {
    v4 := []int{-716}
    var v5 uint8 = 11
    var v6 int32 = -2147483648
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []uint32{713, 0, 449}
    var v2 bool = (42 < 536)
    var v3 string = ("\n\t" + "hello")
    switch v3 {
    case "a":
        if v2 {
            v7 := GenS1{F0: 648, F1: (v3 + v3), F2: 83}
            v8 := GenS1{F0: 2, F1: ((v3 + "foo bar") + "a"), F2: 255}
            _ = v7
            _ = v8
        } else {
            var v9 int32 = 1
            _ = v9
        }
    default:
        switch -9223372036854775808 {
        case 130:
            v3 = (((v3 + v3) + v3) + (("a" + (("\n\t" + (v3 + ("" + ("foo bar" + (v3 + (v3 + v3)))))) + "foo bar")) + "a"))
        case 2147483647:
            switch v3 {
            case "":
                for fi10 := 0; fi10 < 13; fi10++ {
                    _ = len(v1)
                    _ = fi10
                }
            }
        }
    }
    var v11 uint64 = 4294967295
    var v12 int64 = 1
    v13 := GenS2{F0: 255, F1: 128, F2: -1}
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS1{F0: 949, F1: "foo bar", F2: 18446744073709551615}
        if (v2 || v2) {
            var v16 float32 = -2.25
            fmt.Println(v11)
            if len(v1) > 0 {
                v1[0] = genFn0()
            }
            _ = v16
        }
        if (!(v2 && false)) {
            _ = len(v1)
            switch fi14 {
            case 32768:
                var v17 uint8 = 2
                _ = v17
            case 4294967295:
                var v18 int64 = (v12 ^ int64(v11))
                _ = v18
            }
        }
        _ = fi14
        _ = v15
    }
    _ = len(v1)
    if (true || (!(((0.1 * 1.5) + -2.25) >= 3.14159))) {
        fmt.Println(v3, v12, v11)
        _ = len(v1)
    } else {
        v19 := GenS2{F0: 255, F1: 255, F2: -1}
        v20 := GenS1{F0: 306, F1: v3, F2: 921}
        _ = v19
        _ = v20
    }
    switch 2147483647 {
    case 9223372036854775807:
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := GenS2{F0: 0, F1: 127, F2: fi21}
            v2 = (2 != 720)
            _ = fi21
            _ = v22
        }
    case -82:
        for fi23 := 0; fi23 < 19; fi23++ {
            v12 = 599
            var v24 int64 = (v12 * 1)
            var v25 uint16 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        v3 = v3
        _ = fi26
    }
    if false {
        v27 := GenS1{F0: 32767, F1: v3, F2: 65535}
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        for fi29 := 0; fi29 < 8; fi29++ {
            v1 = append(v1, 65535)
            _ = fi29
        }
        fmt.Println(fi28, v2, v12)
        v3 = v3
        _ = fi28
    }
    v12 = (v12 ^ (v12 ^ (v12 % (((v12 / ((v12) | 1))) | 1))))
    v11 += 356
    if (false && v2) {
        v1 = append(v1, genFn0())
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
