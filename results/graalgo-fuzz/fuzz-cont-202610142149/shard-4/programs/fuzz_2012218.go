package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 bool) uint {
    p0 = true
    if (!((("hello" + "a") + ("foo bar" + "hello")) <= "foo bar")) {
        var v5 float32 = float32(0.0)
        var v6 int32 = -2147483648
        _ = v5
        _ = v6
    }
    switch "a" {
    case "foo bar":
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: true, F1: 69}
                if (-1 > 127) {
                    p0 = p0
                } else {
                    v10 := []int{(fi7 << ((65535) & 63)), fi7}
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    fmt.Println(p0, fi7)
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case "\n\t":
        var v11 bool = ("\n\t" <= ("" + "hello"))
        _ = v11
    default:
        p0 = (p0 || true)
    }
    _ = p0
    return 65535
}

func main() {
    var v1 int16 = 1000
    var v2 bool = false
    var v3 int8 = 34
    var v4 string = "a"
    v12 := GenS1{F0: v2, F1: 2147483647}
    var v13 uint16 = 42
    v14 := GenS1{F0: v2, F1: 2147483647}
    var v15 uint = 2
    if v2 {
        if v2 {
            for fi16 := 0; fi16 < 15; fi16++ {
                fmt.Println(fi16, v1)
                var v17 uint16 = v13
                switch fi16 {
                case 372:
                    switch (^(fi16)) {
                    case 32767:
                        v3 -= v3
                    case -158:
                        var v18 int8 = (-(v3))
                        _ = v18
                    case 2147483647:
                        v19 := GenS2{F0: 0, F1: (v15 & v15)}
                        _ = v19
                    default:
                        v20 := []int{fi16, (fi16 & 1)}
                        _ = v20
                    }
                case 0:
                    var v21 int8 = v3
                    _ = v21
                case 1:
                    v17 += (v17 & v13)
                }
                _ = fi16
                _ = v17
            }
            v3 -= v3
        } else {
            if (882 < (v1 & (v1 >> ((v15) & 63)))) {
                var v22 int64 = 9223372036854775807
                v23 := []int{-899, 0, -726}
                var v24 int16 = v1
                _ = v22
                _ = v23
                _ = v24
            } else {
                fmt.Println(v3, v4, v4)
            }
            v25 := GenS2{F0: -2147483648, F1: 1}
            _ = v25
        }
        var v26 int64 = -1
        v1 = (^(v1))
        _ = v26
    } else {
        fmt.Println(v1, v15, v15)
        switch 127 {
        case 185:
            fmt.Println(v4, v3)
        case 887:
            var v27 int32 = 2147483647
            _ = v27
        default:
            var v28 uint32 = 2
            _ = v28
        }
        var v29 uint32 = 1
        _ = v29
    }
    v30 := GenS1{F0: (0.1 != float64((float32(1e10) - 1.5))), F1: 197}
    for fi31 := 0; fi31 < 4; fi31++ {
        v13 -= (v13 & (v13 * (v13 | 32767)))
        _ = fi31
    }
    if v2 {
        v3 -= v3
        if false {
            v32 := GenS1{F0: true, F1: 32768}
            for fi33 := 0; fi33 < 16; fi33++ {
                switch (-(fi33)) {
                case -456:
                    v1 -= v1
                }
                if true {
                    var v34 float32 = 1.5
                    v35 := GenS1{F0: (-9223372036854775808 != int64(v15)), F1: (fi33 % (((fi33 | fi33)) | 1))}
                    var v36 bool = v2
                    _ = v34
                    _ = v35
                    _ = v36
                }
                _ = fi33
            }
            _ = v32
        }
    } else {
        v13 = (v13 | (v13 | 1))
        v13 = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
