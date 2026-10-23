package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 string }
func genFn0() uint64 {
    var v5 uint = 4294967295
    var v6 float32 = (0.1 + 1.5)
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int8, p1 int16) uint32 {
    if (false && false) {
        p0 = p0
        switch 2147483647 {
        case 9223372036854775807:
            v7 := GenS1{F0: 2147483648}
            _ = v7
        }
        if true {
            var v8 uint = 65535
            for fi9 := 0; fi9 < 8; fi9++ {
                switch "foo bar" {
                case "\n\t":
                    v8 -= (v8 % ((4294967295) | 1))
                case "foo bar":
                    var v10 uint16 = 338
                    _ = v10
                }
                _ = fi9
            }
            var v11 string = "\n\t"
            _ = v11
            _ = v8
        }
    }
    p0 = p0
    v12 := GenS1{F0: 979}
    v13 := []int{4294967295}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 426
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := []string{(("\n\t" + "") + "foo bar")}
    v3 := make(map[string]bool)
    v4 := []string{("foo bar" + ""), ("a" + "hello"), ((((("hello" + "hello") + "foo bar") + "a") + "hello") + ("a" + ""))}
    if false {
        var v14 uint32 = 4294967295
        _ = v14
    } else {
        if false {
            _ = len(v3)
            var v15 uint32 = 1
            _ = v15
        } else {
            if false {
                var v16 int16 = 42
                v16 -= v16
                _ = v16
            } else {
                if false {
                    var v17 uint8 = 186
                    _ = v17
                }
                v18 := GenS2{F0: -594, F1: "foo bar"}
                _ = v18
            }
            var v19 int8 = 13
            _ = len(v1)
            _ = v19
        }
        v4 = append(v4, "\n\t")
        var v20 uint8 = 1
        _ = v20
    }
    _ = len(v3)
    var v21 int16 = 813
    v22 := v3[(("\n\t" + "hello") + "hello")]
    fmt.Println(v21, v21)
    var v23 uint64 = 0
    var v24 int32 = -979
    fmt.Println(v23, v24)
    for fi25 := 0; fi25 < 19; fi25++ {
        _ = len(v2)
        for fi26 := 0; fi26 < 15; fi26++ {
            if v22 {
                if v22 {
                    var v27 uint16 = 65535
                    v28 := GenS1{F0: genFn1(int8(fi25), (v21 << ((366) & 63)))}
                    var v29 uint = 2
                    v1 = append(v1, (fi26 >> (((v29 * 82)) & 63)))
                    _ = v27
                    _ = v28
                    _ = v29
                }
            } else {
                v23 -= (v23 + genFn0())
                v21 = (v21 / ((-32768) | 1))
            }
            _ = fi26
        }
        fmt.Println(v21, v24)
        _ = fi25
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v23 = (v23 >> ((1) & 63))
        var v31 int16 = 780
        _ = fi30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
