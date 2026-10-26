package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int) bool {
    p0 += 4294967295
    for fi5 := 0; fi5 < 20; fi5++ {
        p0 -= 127
        var v6 float32 = 3.14159
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := []int{(^(fi7))}
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    p0 += p0
    v9 := []int{p0, (p0 / ((0) | 1)), p0}
    v10 := GenS1{F0: p0, F1: false, F2: 127}
    _ = p0
    _ = v10
    _ = v9
    return (!true)
}
func genFn1(p0 int32) string {
    p0 = (p0 << ((0) & 63))
    p0 += p0
    v11 := []int{}
    if len(v11) > 0 {
        v11[0] = 128
    }
    p0 += 2147483647
    _ = p0
    _ = v11
    return (("hello" + "\n\t") + (("a" + "foo bar") + (("" + ("foo bar" + "foo bar")) + ((("hello" + "foo bar") + (((("foo bar" + ("hello" + "\n\t")) + ("a" + "hello")) + ("\n\t" + "foo bar")) + ("\n\t" + "a"))) + ("" + "a")))))
}

func main() {
    var v1 int16 = 144
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    var v4 int64 = -9223372036854775808
    fmt.Println(v1, v4)
    v12 := v2[-875]
    switch "hello" {
    case "":
        v12 = true
    case "\n\t":
        v13 := []int{32768, -1}
        _ = v13
    case "hello":
        if (v12 && true) {
            if true {
                v1 = 1
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 uint8 = 255
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v4 = (-(v4))
                        if false {
                            v12 = v12
                        } else {
                            var v17 uint = 255
                            _ = v17
                        }
                        if v12 {
                            var v18 uint32 = 708
                            switch "\n\t" {
                            case "foo bar":
                                var v19 float32 = 1e10
                                _ = v19
                            default:
                                v18 += v18
                            }
                            _ = v18
                        } else {
                            if (!true) {
                                var v20 int8 = 0
                                _ = v20
                            } else {
                                _ = len(v3)
                                var v21 uint16 = 717
                                _ = v21
                            }
                            v22 := v2[int(v4)]
                            var v23 string = (("foo bar" + ("\n\t" + "\n\t")) + "")
                            _ = v22
                            _ = v23
                        }
                        _ = fi16
                    }
                    _ = fi14
                    _ = v15
                }
            }
            v24 := GenS1{F0: 629, F1: true, F2: 22}
            _ = v24
        }
    default:
        v1 = 0
    }
    switch "foo bar" {
    case "a":
        var v25 int8 = 0
        _ = v25
    case "foo bar":
        fmt.Println(v1)
    }
    v26 := GenS2{F0: -128}
    var v27 bool = true
    v2[32768] = v12
    v1 = v1
    v4 -= 1
    for fi28 := 0; fi28 < 17; fi28++ {
        v29 := GenS2{F0: -27}
        var v30 string = "\n\t"
        _ = fi28
        _ = v29
        _ = v30
    }
    if (100 < 42) {
        var v31 float64 = float64(float32(0.1))
        fmt.Println(v4)
        var v32 bool = (v27 || (v4 == (-(v4))))
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        var v34 uint8 = 42
        v27 = true
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
