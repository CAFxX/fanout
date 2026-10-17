package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int32 }
func genFn0(p0 int8) string {
    var v6 bool = (((!true) && (false || false)) && (0 != 32767))
    var v7 float32 = 0.1
    _ = p0
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1(p0 bool, p1 uint) int16 {
    for fi8 := 0; fi8 < 10; fi8++ {
        p1 = 238
        fmt.Println(p0, fi8, p0)
        v9 := GenS1{F0: (fi8 / (((fi8 << (((p1 << (((p1 + p1)) & 63))) & 63))) | 1))}
        _ = fi8
        _ = v9
    }
    p1 = 65535
    v10 := []int{}
    var v11 uint = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -32768
}

func main() {
    var v1 float32 = float32(1.5)
    v2 := make(map[string]string)
    var v3 uint8 = 128
    v4 := []string{(("a" + ("a" + "a")) + "hello"), "hello", ("foo bar" + "a")}
    v5 := make(map[uint32]string)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        var v12 int8 = 0
        fmt.Println(v3, v1)
        for fi13 := 0; fi13 < 5; fi13++ {
            switch "foo bar" {
            case "":
                var v14 int = 0
                _ = v14
            case "a":
                fmt.Println(v12)
            default:
                if true {
                    if false {
                        v4 = append(v4, ("a" + ("a" + "a")))
                    }
                } else {
                    var v15 int8 = (^(v12))
                    var v16 uint = 1
                    _ = v15
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := GenS2{F0: 5, F1: -1}
                v1 -= (v1 + 1.5)
                _ = fi17
                _ = v18
            }
            v19 := GenS2{F0: (v3 / ((v3) | 1)), F1: 1}
            _ = fi13
            _ = v19
        }
        _ = v12
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        if false {
            var v21 uint64 = 268
            _ = v21
        } else {
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 int = (fi20 - fi22)
                fmt.Println(fi20, fi22)
                _ = fi22
                _ = v23
            }
            v1 = v1
        }
        var v24 uint16 = 1
        _ = fi20
        _ = v24
    }
    if (2 > 232) {
        v3 = (^(v3))
        var v25 int16 = genFn1(((^(v3)) <= 128), 627)
        var v26 uint32 = 1
        _ = v25
        _ = v26
    }
    var v27 uint8 = v3
    v1 -= ((1.5 - v1) + v1)
    fmt.Println(v27, v3)
    v2["foo bar"] = "a"
    var v28 uint16 = 0
    v1 = v1
    if false {
        for fi29 := 0; fi29 < 2; fi29++ {
            v1 -= v1
            _ = fi29
        }
    }
    v5[679] = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
