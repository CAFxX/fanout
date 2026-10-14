package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
func genFn0(p0 int8, p1 int32, p2 int) string {
    if (2 == 659) {
        p1 = (p1 ^ (-(p1)))
        p2 -= (p2 << ((0) & 63))
    } else {
        if false {
            v7 := []int{(p2 * p2), 2147483647}
            var v8 uint64 = 1
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 bool = false
                _ = fi9
                _ = v10
            }
            _ = v7
            _ = v8
        }
    }
    v11 := []int{(p2 / ((p2) | 1)), 0, 4294967295, 2147483647}
    switch ("hello" + "hello") {
    case "foo bar":
        var v12 float32 = 0.0
        _ = v12
    case "":
        var v13 int16 = -131
        _ = v13
    case "hello":
        v14 := GenS1{F0: ("" + ("" + ("\n\t" + ("a" + ("a" + "a"))))), F1: 32767, F2: int8(p1)}
        _ = v14
    default:
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 int32 = (p1 - -580)
                _ = len(v11)
                p1 = (v17 & (-(v17)))
                switch (^(fi16)) {
                case 4294967295:
                    _ = len(v11)
                case 9223372036854775807:
                    var v18 float64 = 1e10
                    _ = v18
                case 1:
                    var v19 bool = ("" >= "")
                    _ = v19
                }
                p1 += 1
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    }
    p1 -= (p1 - p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return ""
}
func genFn1(p0 int16, p1 uint) int8 {
    var v20 int = 0
    var v21 int32 = 1
    v22 := GenS1{F0: ("\n\t" + "a"), F1: p0, F2: 0}
    p0 += (p0 * (p0 * int16(v21)))
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    _ = v22
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 127
    v3 := []int{32768, 127, -2147483648}
    var v4 uint = 977
    v5 := make(map[string]uint32)
    var v6 bool = false
    var v23 float64 = (-2.25 - 0.0)
    var v24 int32 = -1
    v5["a"] = 2147483647
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v25 float32 = ((0.1 + float32(v23)) / 3.14159)
    if len(v3) > 0 {
        v3[0] = 128
    }
    if v6 {
        fmt.Println(v4, v25)
        for fi26 := 0; fi26 < 19; fi26++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi26
        }
    }
    if false {
        v4 = (v4 * 4294967295)
    } else {
        fmt.Println(v2, v2)
        for fi27 := 0; fi27 < 6; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 uint64 = 910
                _ = fi28
                _ = v29
            }
            var v30 uint64 = 714
            v31 := GenS1{F0: "foo bar", F1: -82, F2: 102}
            _ = fi27
            _ = v30
            _ = v31
        }
    }
    v4 = v4
    var v32 uint = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
