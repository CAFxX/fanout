package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int8; F2 bool }
type GenS3 struct { F0 int64; F1 uint16; F2 int8 }
type GenS4 struct { F0 int32; F1 uint64; F2 int }
type GenS5 struct { F0 string; F1 uint; F2 uint64 }

func main() {
    var v1 int16 = 0
    v2 := []uint32{4294967295, 2147483647}
    v3 := []string{"foo bar", (("\n\t" + "a") + (("" + "hello") + "\n\t")), ""}
    v4 := []string{(("\n\t" + "foo bar") + "\n\t"), ("hello" + ""), "a", ("a" + "")}
    v5 := GenS1{F0: 9223372036854775807}
    fmt.Println(v1)
    v6 := GenS2{F0: 4294967295, F1: 127, F2: (127 == 242)}
    v7 := GenS3{F0: 9223372036854775807, F1: 264, F2: -28}
    var v8 int32 = -1
    if (((int64(v8) >= -1) && (523 > 374)) || (1e10 > (((1e10 + 1e10) * float64((float32(1.5) - (((1e10 * (1e10 - 1.5)) + (float32((0.0 * -2.25)) / 3.14159)) - 0.1)))) + (-2.25 / 1e10)))) {
        if len(v4) > 0 {
            v4[0] = "\n\t"
        }
        var v9 uint = 42
        var v10 int32 = v8
        _ = v10
        _ = v9
    } else {
        if ("\n\t" == "a") {
            v11 := GenS1{F0: -48}
            if (false || true) {
                v4 = append(v4, ("foo bar" + ("foo bar" + "a")))
            }
            v1 = (v1 ^ (v1 / ((v1) | 1)))
            _ = v11
        } else {
            v1 -= (v1 >> ((0) & 63))
            if len(v4) > 0 {
                v4[0] = ("a" + "a")
            }
        }
        switch ("hello" + (("foo bar" + "hello") + "a")) {
        case "foo bar":
            var v12 uint8 = 42
            _ = v12
        default:
            if true {
                v8 -= 170
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        v13 := GenS4{F0: v8, F1: 2, F2: 42}
        _ = v13
    }
    v14 := GenS5{F0: "", F1: 0, F2: 42}
    if (int(v1) < 42) {
        _ = len(v3)
    }
    if ((0.1 / float32(float64(0.1))) <= 0.1) {
        fmt.Println(v1)
        v15 := GenS4{F0: (v8 % ((v8) | 1)), F1: 519, F2: 648}
        _ = v15
    } else {
        if true {
            for fi16 := 0; fi16 < 12; fi16++ {
                v1 = v1
                _ = fi16
            }
            fmt.Println(v8)
        } else {
            var v17 uint16 = 948
            v17 += v17
            _ = v17
        }
    }
    v8 = v8
    var v18 uint8 = 115
    fmt.Println(v1, v8, v1)
    var v19 int8 = -28
    v19 -= -1
    fmt.Println(v1, v8, v8)
    switch -627 {
    case 4294967295:
        v20 := GenS3{F0: -445, F1: 65535, F2: v19}
        _ = v20
    default:
        v1 += (-(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
