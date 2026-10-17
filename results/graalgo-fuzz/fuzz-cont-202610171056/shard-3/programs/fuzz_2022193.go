package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint, p1 uint, p2 int64) uint8 {
    v7 := []int{-2147483648, -91}
    if len(v7) > 0 {
        v7[0] = 0
    }
    v8 := GenS1{F0: "\n\t", F1: "\n\t"}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 65535
    v3 := []uint32{0, 2, 2147483647, 984}
    var v4 float32 = -2.25
    v5 := make(map[string]uint32)
    var v6 int32 = 480
    for fi9 := 0; fi9 < 18; fi9++ {
        if (429 != 2) {
            v10 := GenS2{F0: 32768, F1: 9223372036854775807}
            var v11 uint32 = 4294967295
            fmt.Println(v4)
            _ = v10
            _ = v11
        } else {
            if (32767 < -32768) {
                switch (("a" + ("\n\t" + "hello")) + "a") {
                case "a":
                    var v12 int16 = -1
                    _ = v12
                case "hello":
                    var v13 uint16 = 312
                    _ = v13
                case "foo bar":
                    if true {
                        v5["foo bar"] = 2147483647
                        var v14 uint8 = 2
                        _ = v14
                    }
                default:
                    var v15 bool = ((false || false) && ((^(v6)) != v6))
                    _ = v15
                }
            } else {
                v5[((("\n\t" + "hello") + ("\n\t" + (("\n\t" + "a") + "a"))) + ("foo bar" + "hello"))] = 2
                var v16 uint64 = 981
                if false {
                    fmt.Println(v4, v16, v6)
                }
                _ = v16
            }
            var v17 float32 = (0.1 / (v4 + v4))
            v18 := GenS3{F0: 2147483647, F1: (true && true)}
            _ = v17
            _ = v18
        }
        v19 := v1["a"]
        v4 = v4
        _ = fi9
        _ = v19
    }
    switch 42 {
    case 2147483647:
        if (!(42 >= 0)) {
            var v20 int64 = -1
            _ = v20
        }
    }
    var v21 uint64 = 4294967295
    v22 := GenS3{F0: uint32(v21), F1: (((-1 > 42) && ((true || true) && (false && false))) || (-1 <= (v6 * (v6 + (v6 % ((-1) | 1))))))}
    v23 := GenS4{F0: 1}
    v4 = 0.0
    var v24 uint32 = 65535
    v2 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v24 -= (v24 - 2)
    var v25 uint = 110
    var v26 int32 = (-(v6))
    v5["foo bar"] = 274
    delete(v1, "\n\t")
    if (!false) {
        fmt.Println(v2, v2, v24)
        var v27 int = 292
        v1["\n\t"] = ""
        _ = v27
    } else {
        v28 := GenS3{F0: 2147483647, F1: (true || false)}
        v25 -= (v2 * (v2 - v25))
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        fmt.Println(v26, v25)
        v21 = v21
        var v30 float32 = float32((float64(0.1) / 1.0))
        _ = fi29
        _ = v30
    }
    var v31 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
