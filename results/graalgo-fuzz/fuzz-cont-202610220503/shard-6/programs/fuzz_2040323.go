package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint32; F1 int }

func main() {
    var v1 float64 = 1.5
    var v2 float64 = 1e10
    var v3 int = -967
    switch (^(v3)) {
    case -9223372036854775808:
        for fi4 := 0; fi4 < 3; fi4++ {
            v5 := []int{(v3 % ((650) | 1)), (-(fi4))}
            if len(v5) > 0 {
                v5[0] = (fi4 << ((1) & 63))
            }
            _ = fi4
            _ = v5
        }
    default:
        v6 := []int{(v3 ^ (v3 & v3)), (v3 | v3)}
        _ = v6
    }
    v7 := []int{-526, 2147483647, 42}
    fmt.Println(v1, v2, v3)
    var v8 uint32 = 969
    fmt.Println(v1, v2, v8)
    v9 := GenS1{F0: 0, F1: v3}
    if len(v7) > 0 {
        v7[0] = (v3 % ((-9223372036854775808) | 1))
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 uint32 = (^(v8))
        if (-32768 > -21) {
            if ("foo bar" == (("" + "") + "foo bar")) {
                v1 = float64(1.5)
            } else {
                _ = len(v7)
                v3 = fi10
                if false {
                    var v12 int8 = 0
                    _ = v12
                } else {
                    var v13 int16 = 733
                    v14 := GenS1{F0: 565, F1: (v3 + (fi10 >> ((42) & 63)))}
                    var v15 string = "hello"
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
        }
        _ = fi10
        _ = v11
    }
    if ((^(v8)) <= 65535) {
        var v16 float32 = float32(1e10)
        switch "hello" {
        case "":
            v17 := GenS1{F0: 148, F1: v3}
            _ = v17
        case "\n\t":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case "a":
            v7 = append(v7, -1)
        default:
            var v18 float64 = 0.1
            _ = v18
        }
        _ = v16
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        v1 = 0.1
        _ = fi19
    }
    var v20 int8 = 42
    v21 := GenS2{F0: (v8 >> ((803) & 63)), F1: v3}
    v2 -= float64(-2.25)
    v22 := GenS2{F0: (v8 / ((v8) | 1)), F1: (v3 << ((65535) & 63))}
    v1 += v2
    if ((147 < 2) || true) {
        v7 = append(v7, v3)
        v1 -= v2
        if false {
            if false {
                if false {
                    if (("a" + "foo bar") > "hello") {
                        if (!true) {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        }
                    }
                    var v23 int8 = -75
                    _ = v23
                }
                fmt.Println(v20)
                v3 = (v3 >> ((0) & 63))
            }
            fmt.Println(v20, v20, v1)
        } else {
            var v24 uint64 = 1
            _ = v24
        }
    }
    if true {
        var v25 uint64 = 4294967295
        var v26 int8 = (v20 + v20)
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: (v8 + 2147483647), F1: (v3 >> ((0) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
