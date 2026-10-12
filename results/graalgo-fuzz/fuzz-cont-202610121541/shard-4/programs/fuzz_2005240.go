package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int }
func genFn0(p0 int8, p1 uint64, p2 uint16) int32 {
    switch 127 {
    case 2147483647:
        p0 = -53
    case 32768:
        switch ((((("foo bar" + "a") + ("" + "a")) + ((("hello" + ("a" + "hello")) + "\n\t") + "hello")) + (("foo bar" + "") + "a")) + ((("hello" + "") + "\n\t") + "")) {
        case "a":
            p1 -= p1
        case "hello":
            var v5 uint8 = 0
            _ = v5
        default:
            for fi6 := 0; fi6 < 18; fi6++ {
                for fi7 := 0; fi7 < 17; fi7++ {
                    var v8 int8 = (p0 << ((376) & 63))
                    _ = fi7
                    _ = v8
                }
                var v9 uint64 = 4294967295
                var v10 float32 = 0.0
                _ = fi6
                _ = v10
                _ = v9
            }
        }
    default:
        v11 := []int{32768, -934}
        _ = v11
    }
    var v12 int16 = -1
    v13 := []int{}
    p2 -= 891
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return -445
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = (0.1 - 0.0)
    var v3 uint32 = 65535
    v4 := []string{"", ""}
    if true {
        switch (((((("\n\t" + "a") + "") + "foo bar") + "foo bar") + "a") + "hello") {
        case "hello":
            var v14 float32 = 3.14159
            _ = v14
        case "\n\t":
            for fi15 := 0; fi15 < 17; fi15++ {
                v3 = v3
                fmt.Println(v3, fi15, fi15)
                _ = fi15
            }
        }
        fmt.Println(v3)
    } else {
        var v16 float32 = (-2.25 * v2)
        v17 := GenS1{F0: 65535, F1: ("foo bar" + ("a" + ""))}
        v18 := GenS1{F0: uint16(v3), F1: "foo bar"}
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS2{F0: -9223372036854775808}
    var v20 bool = (0.1 < -2.25)
    v3 += (v3 % (((v3 % ((v3) | 1))) | 1))
    v2 = v2
    var v21 int32 = -505
    var v22 uint64 = 668
    var v23 int64 = 42
    var v24 int16 = 42
    v2 = v2
    delete(v1, "foo bar")
    v24 += v24
    fmt.Println(v22, v24, v2)
    var v25 int32 = genFn0(127, (v22 - (^(v22))), 471)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
