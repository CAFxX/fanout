package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 int64, p1 uint, p2 uint16) uint8 {
    if false {
        fmt.Println(p0, p2, p1)
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{445}
        var v7 float64 = float64(3.14159)
        fmt.Println(p1, v7)
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 bool, p1 int32) uint16 {
    p1 = -2147483648
    p1 = (p1 ^ (p1 % ((p1) | 1)))
    if p0 {
        fmt.Println(p0, p0)
    } else {
        var v8 int32 = p1
        fmt.Println(p0, v8)
        _ = v8
    }
    switch ("foo bar" + "a") {
    case "":
        v9 := []int{32768, 42, -1}
        _ = v9
    case "a":
        for fi10 := 0; fi10 < 1; fi10++ {
            p1 += 730
            _ = fi10
        }
    case "foo bar":
        v11 := GenS1{F0: 212}
        _ = v11
    default:
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{65535}
    var v3 float64 = 3.14159
    var v4 int64 = -218
    var v12 float32 = -2.25
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 uint32 = 42
        v14 = (v14 - v14)
        v12 = 0.0
        _ = fi13
        _ = v14
    }
    switch "hello" {
    case "foo bar":
        v4 += (v4 ^ -9223372036854775808)
    }
    var v15 float32 = (3.14159 / 0.1)
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS1{F0: 255}
        _ = fi16
        _ = v17
    }
    v3 = v3
    v2 = append(v2, 65535)
    var v18 uint16 = 0
    fmt.Println(v18)
    v2 = append(v2, 836)
    for fi19 := 0; fi19 < 6; fi19++ {
        v2 = append(v2, 65535)
        _ = fi19
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        switch "\n\t" {
        case "":
            if false {
                v21 := GenS1{F0: 1}
                if false {
                    var v22 int32 = 42
                    _ = v22
                } else {
                    delete(v1, (("\n\t" + "\n\t") + ((("" + "hello") + ("hello" + "")) + "\n\t")))
                }
                fmt.Println(fi20, v12)
                _ = v21
            } else {
                if false {
                    fmt.Println(v3, v3, v18)
                    var v23 int16 = 1000
                    _ = v23
                }
            }
        case "\n\t":
            if (!(65535 == 870)) {
                switch (^(fi20)) {
                case 32767:
                    var v24 bool = (!false)
                    _ = v24
                default:
                    fmt.Println(v12, v18, fi20)
                }
                var v25 string = ""
                if true {
                    v26 := GenS2{F0: fi20, F1: 42}
                    if true {
                        v27 := GenS2{F0: (-(fi20)), F1: 661}
                        _ = v27
                    } else {
                        var v28 int8 = -128
                        _ = v28
                    }
                    _ = len(v1)
                    _ = v26
                }
                _ = v25
            } else {
                var v29 string = ""
                _ = v29
            }
        default:
            v12 += (v15 / float32((-2.25 / v3)))
        }
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
