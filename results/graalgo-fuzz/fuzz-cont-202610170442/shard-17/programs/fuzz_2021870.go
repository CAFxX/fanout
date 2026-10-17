package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 int64) string {
    v5 := []int{534, 9223372036854775807, -625}
    var v6 string = "foo bar"
    var v7 string = (v6 + "foo bar")
    v6 = ("" + v6)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return v6
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 32
    v3 := make(map[int]bool)
    var v4 uint32 = 1000
    v8 := []int{(v2 * 128), v2, (v2 - (v2 | 42)), (v2 + v2)}
    if (("foo bar" + "hello") > "foo bar") {
        v9 := v1[("hello" + "hello")]
        fmt.Println(v9)
        _ = v9
    }
    fmt.Println(v2, v2)
    var v10 int32 = 673
    var v11 string = "hello"
    var v12 int = -1
    if (1 >= -1) {
        switch "" {
        case "a":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        case "hello":
            for fi13 := 0; fi13 < 7; fi13++ {
                v12 += (v12 << ((4294967295) & 63))
                for fi14 := 0; fi14 < 10; fi14++ {
                    if false {
                        v4 -= v4
                    }
                    _ = fi14
                }
                _ = fi13
            }
        case "foo bar":
            v4 = (v4 << ((969) & 63))
        default:
            v4 += (^(v4))
        }
    }
    v1[v11] = 769
    v4 = v4
    v4 = 4294967295
    v15 := GenS1{F0: -32768, F1: v11}
    if ((2 >= 42) && (v11 > (v11 + v11))) {
        switch "a" {
        case "hello":
            var v16 uint64 = 18446744073709551615
            _ = v16
        case "foo bar":
            switch v11 {
            case "":
                delete(v1, "foo bar")
            case "\n\t":
                switch v11 {
                case "hello":
                    var v17 uint16 = 0
                    _ = v17
                case "":
                    v1[v11] = v4
                case "foo bar":
                    var v18 uint32 = v4
                    _ = v18
                }
            case "hello":
                v11 = "\n\t"
            default:
                v4 += v4
            }
        }
        var v19 int64 = -610
        var v20 int8 = 0
        _ = v19
        _ = v20
    } else {
        var v21 uint = 65535
        v2 += 32767
        if (-1 < 9223372036854775807) {
            var v22 int32 = v10
            if ((false && (v11 != (("hello" + v11) + (v11 + v11)))) || true) {
                if (!false) {
                    for fi23 := 0; fi23 < 18; fi23++ {
                        switch (v2 ^ fi23) {
                        case -908:
                            if false {
                                v12 -= 9223372036854775807
                                v12 += fi23
                            } else {
                                v24 := GenS2{F0: v11, F1: 9223372036854775807}
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                                if (!((v10 < v10) && (true && (0 == -9223372036854775808)))) {
                                    if (!(0.0 >= (0.1 * 0.0))) {
                                        var v25 uint64 = 154
                                        _ = v25
                                    }
                                    var v26 uint = v21
                                    _ = v26
                                }
                                _ = v24
                            }
                        }
                        _ = fi23
                    }
                } else {
                    v11 = v11
                }
                for fi27 := 0; fi27 < 20; fi27++ {
                    fmt.Println(fi27)
                    _ = fi27
                }
            }
            _ = v22
        }
        _ = v21
    }
    v12 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
