package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 int8; F2 string }
func genFn0(p0 uint32) int64 {
    v6 := []int{2147483647}
    if len(v6) > 0 {
        v6[0] = -479
    }
    var v7 uint64 = 112
    v8 := GenS1{F0: "foo bar"}
    v9 := GenS2{F0: 100}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -531
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 0
    v3 := []int{32768}
    v4 := make(map[int]string)
    var v5 float64 = float64(0.1)
    v10 := GenS3{F0: 32767, F1: (!false)}
    var v11 uint32 = v2
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v11 -= (v2 ^ 0)
        v2 = (v11 * 2147483648)
        for fi13 := 0; fi13 < 18; fi13++ {
            if false {
                var v14 int = -1
                v5 -= v5
                _ = v14
            }
            if false {
                v15 := v4[-9223372036854775808]
                v11 = (v2 | (v2 >> ((4294967295) & 63)))
                switch (fi12 & fi12) {
                case -9223372036854775808:
                    v15 = "a"
                case 378:
                    if ((^(v1)) <= v1) {
                        var v16 string = "hello"
                        v2 = 0
                        if false {
                            v15 = v16
                            v17 := GenS4{F0: 42}
                            if true {
                                v16 = "hello"
                                var v18 uint8 = 0
                                if (100 < (v2 * (v2 >> ((1) & 63)))) {
                                    v19 := v4[460]
                                    fmt.Println(v15, fi12, v1)
                                    _ = v19
                                }
                                _ = v18
                            }
                            _ = v17
                        } else {
                            v5 = v5
                            v16 = ("foo bar" + v16)
                        }
                        _ = v16
                    }
                case 4294967295:
                    var v20 int64 = 766
                    _ = v20
                }
                _ = v15
            }
            _ = fi13
        }
        _ = fi12
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        v5 -= (v5 + 0.0)
        if false {
            v5 = v5
            switch (-(fi21)) {
            case -1:
                v3 = append(v3, (fi21 | fi21))
            case 0:
                fmt.Println(v2, v5)
            case -61:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            var v22 uint32 = v2
            var v23 int64 = 912
            switch "" {
            case "\n\t":
                for fi24 := 0; fi24 < 11; fi24++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v25 uint = 18446744073709551615
                    _ = len(v4)
                    _ = fi24
                    _ = v25
                }
            case "foo bar":
                if true {
                    fmt.Println(fi21, v1, v1)
                } else {
                    v26 := GenS3{F0: 32767, F1: (18446744073709551615 <= 18446744073709551615)}
                    v11 += (v22 - 42)
                    if len(v3) > 0 {
                        v3[0] = fi21
                    }
                    _ = v26
                }
            default:
                _ = len(v3)
            }
            _ = v22
            _ = v23
        }
        if true {
            for fi27 := 0; fi27 < 6; fi27++ {
                if (!((2 != 255) && false)) {
                    fmt.Println(fi21)
                    v2 = (v11 + (v11 % ((v2) | 1)))
                    var v28 int32 = 2147483647
                    _ = v28
                }
                v1 -= v1
                _ = fi27
            }
            var v29 int8 = (v1 * v1)
            var v30 int = 32768
            _ = v29
            _ = v30
        } else {
            _ = len(v4)
            v31 := GenS1{F0: "a"}
            v5 = v5
            _ = v31
        }
        _ = fi21
    }
    v32 := GenS5{F0: 127, F1: v1, F2: (("" + ("\n\t" + (("\n\t" + ("a" + "a")) + ("foo bar" + "")))) + "foo bar")}
    v5 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
