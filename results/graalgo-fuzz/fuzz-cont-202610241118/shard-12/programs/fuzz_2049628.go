package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }
func genFn0() uint {
    var v5 int64 = 0
    if true {
        v5 += v5
    } else {
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(fi6, v5)
            var v7 uint16 = 2
            _ = fi6
            _ = v7
        }
        v5 -= v5
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := []int{}
            switch ("foo bar" + "a") {
            case "foo bar":
                v5 -= v5
            default:
                var v10 uint16 = 65535
                _ = v10
            }
            v5 -= 42
            _ = fi8
            _ = v9
        }
    }
    _ = v5
    return 42
}

func main() {
    var v1 int8 = 42
    var v2 float32 = (0.1 * 0.0)
    var v3 string = ""
    v4 := []uint32{2}
    if (int64(v1) < 9223372036854775807) {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int32 = -28
            var v13 int8 = -82
            if len(v4) > 0 {
                v4[0] = 42
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int8 = v1
        _ = v14
    } else {
        v1 = 25
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS1{F0: 293, F1: 127, F2: -280}
        _ = fi15
        _ = v16
    }
    if len(v4) > 0 {
        v4[0] = 2
    }
    if false {
        v1 -= (v1 << ((255) & 63))
        if (-694 < 42) {
            v17 := GenS1{F0: 65535, F1: 0, F2: 0}
            var v18 int = 128
            _ = v17
            _ = v18
        }
        v2 = (v2 - -2.25)
    } else {
        v19 := GenS2{F0: v3, F1: 238, F2: 32767}
        if (float64((1.5 - v2)) > float64((v2 / 3.14159))) {
            var v20 int = 42
            v2 = v2
            v1 = -128
            _ = v20
        }
        v21 := GenS1{F0: 42, F1: 1, F2: -1}
        _ = v19
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        switch v3 {
        case "a":
            for fi23 := 0; fi23 < 13; fi23++ {
                fmt.Println(v2, v2)
                v2 -= (float32(1e10) * v2)
                _ = fi23
            }
        case "\n\t":
            var v24 string = ("hello" + "\n\t")
            _ = v24
        default:
            var v25 uint8 = 173
            _ = v25
        }
        v2 = v2
        for fi26 := 0; fi26 < 20; fi26++ {
            v1 += (v1 - (v1 + v1))
            _ = fi26
        }
        _ = fi22
    }
    v27 := GenS1{F0: 1, F1: 127, F2: 1}
    for fi28 := 0; fi28 < 18; fi28++ {
        v2 = 0.1
        var v29 uint8 = 42
        _ = fi28
        _ = v29
    }
    fmt.Println(v1, v3)
    for fi30 := 0; fi30 < 20; fi30++ {
        v2 -= 0.1
        var v31 uint8 = 216
        var v32 float64 = float64(v2)
        _ = fi30
        _ = v31
        _ = v32
    }
    switch v3 {
    case "foo bar":
        v33 := GenS2{F0: (v3 + (v3 + "\n\t")), F1: 0, F2: 255}
        _ = v33
    case "hello":
        if (1 >= (v1 >> ((286) & 63))) {
            var v34 uint = 42
            for fi35 := 0; fi35 < 5; fi35++ {
                if len(v4) > 0 {
                    v4[0] = 42
                }
                _ = fi35
            }
            _ = v34
        } else {
            var v36 int64 = -9223372036854775808
            v37 := GenS2{F0: "", F1: 698, F2: 65535}
            _ = v36
            _ = v37
        }
    case "\n\t":
        if (-1 > 1) {
            var v38 bool = (true && (500 != 465))
            switch v3 {
            case "":
                var v39 int = 128
                _ = v39
            case "hello":
                var v40 bool = (208 <= 220)
                _ = v40
            case "a":
                v41 := GenS2{F0: "\n\t", F1: -911, F2: 32767}
                _ = v41
            default:
                _ = len(v4)
            }
            _ = v38
        } else {
            v2 = 1.5
        }
    }
    var v42 uint = 4294967295
    v42 -= (v42 / (((v42 * (v42 & v42))) | 1))
    for fi43 := 0; fi43 < 7; fi43++ {
        var v44 uint = v42
        v1 += v1
        var v45 float32 = ((float32(float64(v2)) + (1.5 * (0.1 - ((0.1 / 1.5) + 0.0)))) - v2)
        _ = fi43
        _ = v44
        _ = v45
    }
    switch 223 {
    case 9223372036854775807:
        if false {
            if ((0.0 / float64(((v2 / 3.14159) * 0.0))) != 1e10) {
                var v46 int64 = -9223372036854775808
                var v47 uint16 = 255
                _ = v46
                _ = v47
            } else {
                for fi48 := 0; fi48 < 15; fi48++ {
                    v2 += v2
                    v4 = append(v4, 2147483647)
                    _ = fi48
                }
                var v49 uint = (v42 - (v42 + (v42 ^ (v42 ^ (v42 ^ genFn0())))))
                _ = v49
            }
            v2 -= -2.25
        }
    case 2147483647:
        v42 += (v42 ^ (v42 >> ((2) & 63)))
    default:
        v50 := GenS2{F0: (v3 + (("a" + v3) + v3)), F1: 876, F2: 630}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v42
}
