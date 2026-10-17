package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int16 }
type GenS2 struct { F0 bool; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 string }
func genFn0(p0 uint32) uint16 {
    v6 := []int{382, 124, 42, 128}
    switch "a" {
    case "hello":
        p0 += 4294967295
    default:
        p0 = (p0 % ((9) | 1))
    }
    _ = p0
    _ = v6
    return 2
}

func main() {
    var v1 float32 = ((0.1 / float32(1e10)) + 1.5)
    var v2 bool = (!(!(((((3.14159 / v1) + 3.14159) * 1e10) <= v1) || (1 < 0))))
    v3 := make(map[uint32]string)
    var v4 int16 = 978
    v5 := make(map[int]int)
    if (2 == 2147483648) {
        v4 -= (-(v4))
        if v2 {
            if true {
                v7 := GenS1{F0: 42, F1: 32767, F2: 222}
                switch ((("foo bar" + "foo bar") + "hello") + ("hello" + (("" + "a") + "\n\t"))) {
                case "a":
                    switch -167 {
                    case -295:
                        v2 = (!false)
                    }
                case "foo bar":
                    fmt.Println(v1, v2, v1)
                default:
                    v8 := v5[32768]
                    _ = v8
                }
                _ = v7
            }
        }
        if v2 {
            fmt.Println(v2)
            var v9 int16 = -425
            v10 := GenS2{F0: (!v2), F1: ((v2 || ("foo bar" <= "a")) && (1 <= 141)), F2: 42}
            _ = v10
            _ = v9
        }
    } else {
        v11 := []int{}
        _ = v11
    }
    v12 := GenS1{F0: -9223372036854775808, F1: 42, F2: (v4 / (((v4 >> ((65535) & 63))) | 1))}
    v2 = v2
    switch 358 {
    case 2147483647:
        for fi13 := 0; fi13 < 16; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(fi14)
                switch "foo bar" {
                case "hello":
                    var v15 int32 = 213
                    _ = v15
                case "foo bar":
                    fmt.Println(v1, fi14)
                    var v16 uint8 = 1
                    _ = v16
                case "":
                    v17 := GenS1{F0: 42, F1: 590, F2: 32767}
                    _ = v17
                }
                _ = fi14
            }
            _ = fi13
        }
    }
    v1 -= v1
    v18 := []int{2, 489, 32767, 660}
    if v2 {
        var v19 int = -186
        var v20 int64 = -9223372036854775808
        switch v19 {
        case 1:
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 uint16 = 0
                _ = fi21
                _ = v22
            }
        case 4294967295:
            var v23 uint64 = 2
            _ = v23
        case 32767:
            v18 = append(v18, -2147483648)
        }
        _ = v19
        _ = v20
    }
    v1 = -2.25
    v24 := GenS1{F0: -145, F1: 1, F2: v4}
    v25 := GenS3{F0: 255}
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 uint = 42
        var v28 int32 = 0
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 uint16 = 32767
    switch 570 {
    case 887:
        v2 = false
    case 42:
        for fi30 := 0; fi30 < 10; fi30++ {
            if (!(!true)) {
                for fi31 := 0; fi31 < 17; fi31++ {
                    v4 = 690
                    var v32 int = -9223372036854775808
                    v29 = (v29 >> ((65535) & 63))
                    _ = fi31
                    _ = v32
                }
                fmt.Println(fi30)
            } else {
                for fi33 := 0; fi33 < 20; fi33++ {
                    v2 = (1 <= 1)
                    if (1 <= 42) {
                        v1 = 0.0
                        var v34 float32 = ((1e10 * v1) + float32(3.14159))
                        _ = v34
                    } else {
                        v35 := GenS4{F0: (v29 * v29), F1: "\n\t"}
                        _ = v35
                    }
                    switch fi30 {
                    case -906:
                        var v36 int16 = v4
                        _ = v36
                    case 42:
                        var v37 float32 = v1
                        _ = v37
                    case -1:
                        fmt.Println(v1, v2)
                    default:
                        v38 := GenS1{F0: 613, F1: v29, F2: (v4 / ((v4) | 1))}
                        _ = v38
                    }
                    _ = fi33
                }
            }
            var v39 float32 = ((v1 - v1) / float32(0.1))
            _ = fi30
            _ = v39
        }
    case -384:
        v2 = ((v29 % (((v29 | (v29 & (v29 * (^(v29)))))) | 1)) == 329)
    }
    v3[594] = ("" + ((("hello" + ("hello" + (("a" + ("a" + ("" + "hello"))) + ((("hello" + (("a" + "foo bar") + ("a" + "a"))) + ("foo bar" + "\n\t")) + (("\n\t" + "") + ""))))) + ("a" + (("a" + "") + "\n\t"))) + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
