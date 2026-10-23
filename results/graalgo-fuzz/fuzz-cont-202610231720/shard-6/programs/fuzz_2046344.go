package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int32; F1 int16; F2 int8 }

func main() {
    v1 := []string{"\n\t", ("foo bar" + "\n\t")}
    v2 := []int{42}
    var v3 bool = false
    var v4 int32 = 0
    var v5 int32 = (v4 ^ (v4 >> ((2) & 63)))
    v6 := GenS1{F0: 2147483647, F1: 153}
    var v7 int = -979
    var v8 uint8 = 19
    var v9 string = "foo bar"
    v5 = (v5 & (v4 ^ 1))
    switch v9 {
    case "\n\t":
        if false {
            if v3 {
                v9 = v9
            }
            v10 := GenS2{F0: 84, F1: 42}
            var v11 string = v9
            _ = v10
            _ = v11
        }
    case "foo bar":
        if (!false) {
            var v12 int = (v7 << ((4294967295) & 63))
            _ = v12
        }
    case "hello":
        v9 = v9
    default:
        v2 = append(v2, v7)
    }
    switch (v7 + v7) {
    case 128:
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint32 = 42
            switch fi13 {
            case 32768:
                var v15 int32 = v4
                _ = v15
            case -9223372036854775808:
                var v16 int = fi13
                _ = v16
            default:
                if v3 {
                    var v17 int = 2147483647
                    var v18 uint64 = 4294967295
                    if v3 {
                        fmt.Println(v18, v17, v18)
                        var v19 int32 = v4
                        v9 = v9
                        _ = v19
                    } else {
                        v1 = append(v1, (v9 + "hello"))
                        if (!v3) {
                            fmt.Println(v3, v7, v5)
                            for fi20 := 0; fi20 < 19; fi20++ {
                                if v3 {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v18)
                                } else {
                                    v3 = false
                                }
                                v5 += (v5 % ((-2147483648) | 1))
                                var v21 uint = 2
                                _ = fi20
                                _ = v21
                            }
                        }
                    }
                    _ = v17
                    _ = v18
                }
            }
            if len(v1) > 0 {
                v1[0] = (v9 + v9)
            }
            _ = fi13
            _ = v14
        }
    case 223:
        var v22 int8 = -123
        _ = v22
    }
    if v3 {
        var v23 bool = (!(-32768 <= 0))
        v9 = "foo bar"
        fmt.Println(v23, v9, v5)
        _ = v23
    }
    v2 = append(v2, (v7 << ((65535) & 63)))
    v4 += (v4 - v4)
    v24 := GenS3{F0: 631, F1: 771, F2: 100}
    if v3 {
        var v25 float32 = float32(3.14159)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v25
    } else {
        fmt.Println(v4, v3)
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        for fi27 := 0; fi27 < 19; fi27++ {
            if (255 >= 4294967295) {
                switch -233 {
                case 42:
                    fmt.Println(fi27)
                    if (float32(1.5) == 1e10) {
                        if ((true && true) || (float64(1.5) == (float64((1.5 * 1e10)) - 0.1))) {
                            v3 = (!v3)
                            var v28 int64 = -185
                            _ = v28
                        }
                        v5 += v5
                    } else {
                        var v29 int16 = -32768
                        _ = v29
                    }
                }
                var v30 int16 = 629
                var v31 uint16 = 442
                _ = v30
                _ = v31
            } else {
                fmt.Println(fi26, v7, v3)
                v4 -= (-(v5))
            }
            v8 += (v8 & (v8 ^ (v8 & 128)))
            v3 = (v4 != 282)
            _ = fi27
        }
        if v3 {
            v32 := GenS2{F0: 42, F1: 255}
            _ = v32
        }
        var v33 string = "foo bar"
        _ = fi26
        _ = v33
    }
    var v34 float64 = (((float64(float32((1.5 - 1.5))) / 1.0) - float64((1e10 / 0.1))) + ((0.0 - 3.14159) + float64(float32((((0.0 * 0.1) - (0.1 * 3.14159)) - -2.25)))))
    v8 = v8
    v5 = v4
    v2 = append(v2, (-(v7)))
    fmt.Println(v9, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
