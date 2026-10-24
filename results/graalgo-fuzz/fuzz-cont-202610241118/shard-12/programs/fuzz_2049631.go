package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 int64) uint8 {
    v7 := GenS1{F0: (("hello" + "") + ("hello" + "hello"))}
    if false {
        switch 32767 {
        case 4294967295:
            p0 += -248
        }
        var v8 int64 = p1
        _ = v8
    } else {
        fmt.Println(p1, p0)
    }
    var v9 int = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 1
}

func main() {
    v1 := []string{""}
    v2 := []uint32{65535}
    var v3 uint64 = 4294967295
    var v4 int8 = 100
    var v5 uint16 = 32767
    var v6 uint16 = (^(v5))
    var v10 int8 = (v4 >> ((2) & 63))
    var v11 uint64 = 899
    v2 = append(v2, 794)
    var v12 int64 = 1
    v5 = 1
    if (0 != 2147483647) {
        fmt.Println(v6, v11)
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 int8 = -128
            if false {
                if ((false && false) || (!false)) {
                    fmt.Println(v14, v10, v12)
                    v12 -= (-(v12))
                    v14 -= (^(v4))
                }
                for fi15 := 0; fi15 < 5; fi15++ {
                    switch -34 {
                    case -9223372036854775808:
                        fmt.Println(v10, v10, v14)
                        v16 := GenS1{F0: (("\n\t" + (("\n\t" + "a") + ("" + ("a" + "\n\t")))) + ("" + "foo bar"))}
                        _ = v16
                    default:
                        v17 := GenS1{F0: ((("hello" + "foo bar") + "hello") + "hello")}
                        fmt.Println(v11)
                        _ = v17
                    }
                    v18 := GenS2{F0: (^(fi15)), F1: 52}
                    v19 := GenS3{F0: -695}
                    _ = fi15
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 16; fi20++ {
                    fmt.Println(v14, v3, fi13)
                    if (true || false) {
                        fmt.Println(v12)
                        v2 = append(v2, 2)
                        _ = len(v2)
                    }
                    _ = fi20
                }
            }
            _ = fi13
            _ = v14
        }
        if false {
            var v21 int = int(v12)
            _ = v21
        } else {
            var v22 float64 = (3.14159 / (((3.14159 + ((0.1 / float64((-2.25 - 0.0))) - 1e10)) - float64(-2.25)) / 1.0))
            _ = v22
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v2 = append(v2, 232)
    v11 = 846
    for fi23 := 0; fi23 < 1; fi23++ {
        v4 = 38
        var v24 uint16 = v6
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
