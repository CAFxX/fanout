package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0(p0 int, p1 uint32) int8 {
    v6 := GenS1{F0: 42, F1: 128, F2: 886}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint8 = 128
        if false {
            fmt.Println(p1, p0, fi7)
        }
        _ = fi7
        _ = v8
    }
    v9 := []int{(p0 % ((p0) | 1)), p0, 867, (-(p0))}
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return 116
}

func main() {
    v1 := []int{}
    var v2 int64 = -1
    var v3 int16 = 42
    var v4 int16 = v3
    var v5 string = "\n\t"
    v10 := GenS1{F0: -1, F1: 255, F2: 0}
    v1 = append(v1, 658)
    var v11 int32 = -907
    var v12 int64 = (v2 | (-(v2)))
    v3 -= (-(v4))
    v2 += v2
    switch "hello" {
    case "\n\t":
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 int8 = 42
            switch (fi13 << ((0) & 63)) {
            case 2147483647:
                var v15 float64 = ((-2.25 / 1e10) + float64(-2.25))
                _ = v15
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi13
            _ = v14
        }
    }
    v2 += (v12 | v12)
    v5 = v5
    var v16 int = 0
    var v17 int8 = 1
    var v18 string = ""
    for fi19 := 0; fi19 < 16; fi19++ {
        v1 = append(v1, fi19)
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(v3, v12)
            var v21 uint32 = 4294967295
            if true {
                if len(v1) > 0 {
                    v1[0] = int(v17)
                }
                v2 = v12
                switch v5 {
                case "foo bar":
                    v1 = append(v1, fi20)
                    if false {
                        if ((false || true) && (genFn0(1, (^(v21))) <= (v17 % ((int8(v21)) | 1)))) {
                            v16 += (v16 << ((655) & 63))
                            if len(v1) > 0 {
                                v1[0] = (v16 >> ((1) & 63))
                            }
                        }
                        if (4294967295 <= 255) {
                            var v22 int32 = (v11 & (-(v11)))
                            _ = v22
                        }
                        fmt.Println(fi20, v2)
                    } else {
                        v18 = v18
                        v5 = ""
                    }
                    v23 := GenS2{F0: (-(fi19)), F1: v3}
                    _ = v23
                case "a":
                    v24 := GenS2{F0: int(v4), F1: 32767}
                    _ = v24
                case "hello":
                    v21 = v21
                }
            } else {
                var v25 float64 = float64((float32(((1e10 - float64(0.1)) + (((1e10 * 1e10) * (0.0 + 0.1)) - (0.0 - 3.14159)))) - (0.1 * ((1.5 - -2.25) / ((0.1 + ((0.1 + 1.5) * 3.14159)) / 1.0)))))
                v1 = append(v1, (^(fi19)))
                _ = v25
            }
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    if (-259 >= (v16 ^ -1)) {
        var v26 uint32 = 250
        var v27 uint = 255
        _ = v26
        _ = v27
    } else {
        v5 = ((("foo bar" + v5) + (v5 + v18)) + v18)
    }
    if true {
        fmt.Println(v3)
        v12 = -9223372036854775808
    }
    v18 = (v5 + v18)
    v28 := GenS2{F0: (-(v16)), F1: -32768}
    fmt.Println(v17, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
