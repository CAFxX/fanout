package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 1
    var v3 float32 = float32(((1.5 / 1.0) / 1.0))
    v4 := GenS1{F0: 42, F1: 18446744073709551615}
    v5 := GenS1{F0: 1, F1: 18446744073709551615}
    var v6 int32 = -844
    fmt.Println(v3, v3, v3)
    var v7 uint = 459
    var v8 int32 = (v6 << (((v7 * v7)) & 63))
    v9 := []int{2147483647, -775, 0}
    v6 = v8
    if false {
        var v10 bool = (("foo bar" + (("foo bar" + ("\n\t" + (("\n\t" + "\n\t") + ""))) + "hello")) <= "hello")
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(v8, v8)
            v1[(fi11 / ((fi11) | 1))] = (fi11 | fi11)
            v3 = 0.1
            _ = fi11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(v3, v10, v10)
            _ = fi12
        }
        _ = v10
    }
    v13 := GenS1{F0: (v6 & (v8 & v6)), F1: 1}
    switch -902 {
    case 32768:
        var v14 int16 = 0
        _ = v14
    case 665:
        v3 = ((-2.25 / (v3 - 3.14159)) / (v3 * 0.1))
    case 348:
        if (0 > 2147483647) {
            var v15 float64 = -2.25
            if true {
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 int64 = -9223372036854775808
                    v18 := v1[fi16]
                    fmt.Println(v15)
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v2 = (v2 << ((2) & 63))
                v6 = (v6 << (((v7 % ((v7) | 1))) & 63))
            }
            v6 = (v6 << ((18446744073709551615) & 63))
            _ = v15
        } else {
            var v19 uint = uint(v2)
            _ = v19
        }
    default:
        switch 714 {
        case -428:
            v9 = append(v9, 897)
        default:
            var v20 uint32 = 2147483648
            _ = v20
        }
    }
    v6 = v6
    if ((0 > (v8 | v8)) && false) {
        if (true && ((uint8(v2) != 127) || ((true || ((v6 / ((v6) | 1)) >= (v8 | -1))) && ((!true) && (!false))))) {
            fmt.Println(v2, v7, v7)
        } else {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        var v21 int = 9223372036854775807
        v22 := GenS1{F0: v8, F1: 649}
        _ = v21
        _ = v22
    } else {
        if false {
            v23 := GenS2{F0: 4294967295, F1: 32767}
            v6 = (^(v6))
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v23
        } else {
            var v24 float64 = float64(float32(1.5))
            for fi25 := 0; fi25 < 1; fi25++ {
                v2 -= 840
                if (!(true || (!(569 != 1)))) {
                    var v26 int32 = v8
                    _ = v26
                }
                _ = fi25
            }
            var v27 uint16 = v2
            _ = v24
            _ = v27
        }
    }
    v2 += (v2 & (v2 / ((v2) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
