package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int16; F1 int64 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float32 = 0.1
        _ = fi7
        _ = v8
    }
    v9 := []int{-2147483648, -9223372036854775808, -536, -1}
    var v10 bool = (1 <= -1)
    var v11 uint64 = 392
    if len(v9) > 0 {
        v9[0] = -1
    }
    _ = v10
    _ = v11
    _ = v9
    return 1000
}

func main() {
    v1 := []string{"hello", "hello"}
    v2 := []int{-382, 2147483647}
    var v3 string = "foo bar"
    var v4 uint = 1
    var v5 int8 = 42
    v6 := []uint32{42}
    for fi12 := 0; fi12 < 1; fi12++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v13 := GenS1{F0: (v5 - int8(v4)), F1: fi12}
        _ = fi12
        _ = v13
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 -= v4
    var v14 uint8 = 1
    if (true || (((2 != 24) || (v3 != "a")) || ((!(65535 > 2)) || (false || (true || (1 >= 163)))))) {
        switch 32768 {
        case 32768:
            fmt.Println(v14, v4, v3)
        case -563:
            v3 = (v3 + (v3 + (v3 + v3)))
        default:
            fmt.Println(v4, v4, v14)
        }
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v3 = v3
        v16 := GenS1{F0: v5, F1: fi15}
        switch -605 {
        case 248:
            for fi17 := 0; fi17 < 10; fi17++ {
                v6 = append(v6, 821)
                switch (v3 + v3) {
                case "foo bar":
                    var v18 uint = (v4 - (v4 - (^(v4))))
                    _ = v18
                case "a":
                    v4 = 1
                default:
                    var v19 bool = (((((0.0 - 3.14159) + ((-2.25 * (((1.5 * float64(((0.1 * -2.25) - (1e10 - (float32(0.0) - 0.0))))) * -2.25) - (0.0 * 1e10))) / (-2.25 - 0.0))) * 0.0) / 0.1) < 3.14159)
                    _ = v19
                }
                var v20 string = ""
                switch 127 {
                case 128:
                    v21 := GenS2{F0: 32767, F1: -9223372036854775808}
                    var v22 int = 127
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v22 = (v22 << ((5) & 63))
                    fmt.Println(v22, v20, v22)
                    var v23 float32 = (1.5 + (-2.25 * (1e10 * float32((float64(0.1) + 3.14159)))))
                    var v24 uint64 = 572
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi17
                _ = v20
            }
        case -105:
            if (true || (2147483647 > 2147483648)) {
                var v25 int64 = -9223372036854775808
                v4 = (v4 >> (((^(v4))) & 63))
                _ = v25
            }
        }
        _ = fi15
        _ = v16
    }
    if ((false || true) || (!(!(false || (((!false) || true) && (!false)))))) {
        fmt.Println(v4, v14, v14)
    }
    var v26 uint = 413
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
