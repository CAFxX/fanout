package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 int32, p1 int32) uint64 {
    var v5 int64 = -463
    var v6 uint16 = 42
    switch "" {
    case "\n\t":
        var v7 float32 = 0.1
        _ = v7
    case "":
        p1 += 126
    case "hello":
        v8 := GenS1{F0: 2}
        _ = v8
    default:
        if true {
            fmt.Println(v6, p1)
            if false {
                var v9 bool = false
                _ = v9
            } else {
                v5 = 0
            }
        }
    }
    var v10 bool = false
    var v11 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v5
    _ = v6
    return 854
}
func genFn1(p0 uint16, p1 bool, p2 int16) int64 {
    var v12 int = 9223372036854775807
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint8 = 126
        switch v12 {
        case 350:
            var v15 int8 = 127
            _ = v15
        case 4294967295:
            for fi16 := 0; fi16 < 14; fi16++ {
                if (-32768 > (p2 * (p2 & p2))) {
                    v17 := []int{fi13, fi16}
                    _ = v17
                }
                _ = fi16
            }
        default:
            fmt.Println(p1, p0)
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 64
}

func main() {
    var v1 uint8 = 0
    var v2 float64 = ((((((1e10 / (0.1 - (((0.1 / -2.25) / 1e10) + (float64(-2.25) - 3.14159)))) + (1.5 / 1.0)) - -2.25) / (0.1 - 1.5)) - float64((1e10 + ((float32(1.5) - -2.25) / 3.14159)))) + 1e10)
    var v3 string = ""
    var v4 int16 = 1
    v18 := []int{}
    var v19 int32 = 1
    v18 = append(v18, -9223372036854775808)
    v18 = append(v18, 753)
    var v20 uint = 42
    _ = len(v18)
    var v21 uint8 = v1
    if true {
        if true {
            switch "" {
            case "a":
                for fi22 := 0; fi22 < 15; fi22++ {
                    if ((^(v20)) == v20) {
                        if len(v18) > 0 {
                            v18[0] = (fi22 << (((v20 * (v20 ^ 567))) & 63))
                        }
                        _ = len(v18)
                        fmt.Println(v1)
                    } else {
                        var v23 int = -127
                        if len(v18) > 0 {
                            v18[0] = -9223372036854775808
                        }
                        _ = v23
                    }
                    if len(v18) > 0 {
                        v18[0] = fi22
                    }
                    _ = fi22
                }
            case "":
                fmt.Println(v19)
            }
            if false {
                var v24 uint8 = 153
                v18 = append(v18, 941)
                _ = v24
            }
        } else {
            var v25 uint32 = 1
            v3 = ((((v3 + "a") + ((("hello" + (v3 + v3)) + ("foo bar" + v3)) + v3)) + v3) + v3)
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        v19 += (^(v19))
        var v27 int16 = 1000
        v28 := GenS2{F0: "foo bar", F1: v20}
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
