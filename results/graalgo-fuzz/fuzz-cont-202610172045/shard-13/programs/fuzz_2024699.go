package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint16, p1 int8) string {
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: int32(p1), F1: 18446744073709551615, F2: 311}
            v8 := GenS1{F0: -2147483648, F1: 255, F2: 40}
            if (-9223372036854775808 != 0) {
                if (true && false) {
                    var v9 int32 = 0
                    v9 = (^(v9))
                    v10 := GenS1{F0: v9, F1: 4294967295, F2: 81}
                    var v11 int = (-(fi5))
                    _ = v10
                    _ = v11
                    _ = v9
                }
                if false {
                    if true {
                        var v12 int32 = -494
                        v13 := []int{}
                        _ = len(v13)
                        _ = v12
                        _ = v13
                    }
                    var v14 int16 = 1
                    _ = v14
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint64 = 18446744073709551615
        if (582 >= fi15) {
            v17 := []int{(fi15 >> ((255) & 63)), -855}
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    switch (((("a" + ("hello" + "")) + "foo bar") + ("a" + "hello")) + (((("\n\t" + ("a" + (("" + ("foo bar" + "foo bar")) + ""))) + "foo bar") + "\n\t") + "")) {
    case "foo bar":
        var v18 int = -859
        _ = v18
    case "":
        switch 32767 {
        case 42:
            v19 := []int{-9223372036854775808, 574}
            _ = v19
        }
    case "hello":
        fmt.Println(p1)
    default:
        var v20 int8 = -94
        _ = v20
    }
    _ = p0
    _ = p1
    return "hello"
}

func main() {
    v1 := []uint32{2}
    v2 := make(map[uint32]string)
    var v3 uint32 = 2
    var v4 int64 = 1
    v4 += (-(v4))
    var v21 bool = (42 <= -1)
    if ((false && (0.0 == 1.5)) && v21) {
        switch 920 {
        case -639:
            v22 := v2[216]
            _ = v22
        default:
            var v23 int = 200
            _ = v23
        }
        v24 := GenS1{F0: -463, F1: 65535, F2: 0}
        _ = v24
    } else {
        if v21 {
            for fi25 := 0; fi25 < 3; fi25++ {
                v26 := GenS1{F0: -2147483648, F1: 575, F2: 4294967295}
                var v27 int32 = 0
                _ = len(v1)
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        v4 = (v4 / ((v4) | 1))
        v3 += (v3 ^ 0)
    }
    if v21 {
        v3 -= v3
        if len(v1) > 0 {
            v1[0] = (v3 % ((0) | 1))
        }
        v28 := GenS2{F0: 641, F1: -1, F2: 42}
        _ = v28
    }
    if v21 {
        v29 := GenS3{F0: false}
        v4 = (v4 * v4)
        _ = v29
    } else {
        v3 -= (v3 << ((42) & 63))
        fmt.Println(v3, v21, v4)
        var v30 int8 = -1
        _ = v30
    }
    v21 = (!((v21 || (v21 && false)) || (0 == -32768)))
    var v31 uint64 = 571
    var v32 float32 = (0.0 * ((-2.25 + 3.14159) + 3.14159))
    v32 -= (-2.25 + v32)
    v33 := GenS4{F0: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
