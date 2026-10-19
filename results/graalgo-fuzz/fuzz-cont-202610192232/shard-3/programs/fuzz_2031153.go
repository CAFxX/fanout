package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 uint16 }

func main() {
    v1 := []int{1}
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    var v4 int64 = -9223372036854775808
    v5 := []int{2147483647, 66, 2147483647, -36}
    var v6 uint16 = 0
    var v7 float64 = 0.0
    v3["a"] = 957
    _ = len(v3)
    var v8 uint16 = v6
    v9 := GenS1{F0: 42, F1: 1}
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint16 = 0
        fmt.Println(v8, v8, v11)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v4 = (v4 % (((v4 & (-(v4)))) | 1))
        if true {
            fmt.Println(fi12)
        } else {
            var v13 string = "foo bar"
            _ = v13
        }
        v14 := v2[fi12]
        _ = fi12
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        if (!true) {
            var v16 string = ("a" + "")
            if (false && (!true)) {
                v4 -= (v4 >> ((255) & 63))
            } else {
                v4 = v4
                var v17 float64 = 0.0
                switch v16 {
                case "hello":
                    switch (^(fi15)) {
                    case -1:
                        if (!false) {
                            v17 = float64(1.5)
                            fmt.Println(v16, v6)
                            v3[("" + "a")] = (-(fi15))
                        }
                    case 32767:
                        if (!true) {
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v4 += (v4 ^ v4)
                            delete(v3, (("hello" + "foo bar") + ("hello" + v16)))
                        }
                    case 199:
                        v6 -= v8
                    default:
                        v18 := GenS2{F0: ((true && true) || (!(true && true))), F1: v6}
                        _ = v18
                    }
                default:
                    var v19 uint8 = 127
                    _ = v19
                }
                _ = v17
            }
            _ = v16
        } else {
            var v20 uint64 = 570
            var v21 int16 = -1
            if (false && false) {
                v6 = v6
                if true {
                    v6 += v6
                    var v22 int32 = -244
                    _ = v22
                }
            }
            _ = v20
            _ = v21
        }
        _ = fi15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
