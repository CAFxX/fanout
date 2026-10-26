package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int32; F1 int }
func genFn0(p0 string, p1 uint) string {
    var v5 int8 = -72
    v6 := GenS1{F0: p0, F1: false}
    p1 = p1
    var v7 uint64 = 2
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1() bool {
    var v8 uint16 = 198
    v8 += 740
    v9 := []int{-434, -303, 2147483647, 294}
    _ = v8
    _ = v9
    return (!false)
}

func main() {
    var v1 uint32 = 956
    v2 := make(map[uint32]bool)
    var v3 string = ""
    var v4 uint8 = 128
    fmt.Println(v3)
    fmt.Println(v3, v4, v1)
    for fi10 := 0; fi10 < 16; fi10++ {
        v3 = (("a" + (v3 + (("\n\t" + (v3 + (v3 + ("a" + v3)))) + "foo bar"))) + ((("hello" + "a") + v3) + v3))
        var v11 uint = 653
        _ = fi10
        _ = v11
    }
    if true {
        _ = len(v2)
    } else {
        var v12 int64 = 9223372036854775807
        fmt.Println(v4)
        _ = v12
    }
    var v13 int = 331
    if (608 >= 42) {
        v14 := GenS1{F0: ((v3 + v3) + v3), F1: (v13 <= (v13 ^ (v13 & (-(v13)))))}
        if false {
            v15 := GenS1{F0: v3, F1: (!true)}
            _ = v15
        } else {
            var v16 float32 = (1.5 - float32(0.0))
            _ = v16
        }
        _ = v14
    } else {
        v3 = v3
    }
    if false {
        v3 = ("foo bar" + v3)
    }
    v17 := GenS2{F0: -128, F1: v4}
    fmt.Println(v3)
    var v18 int32 = 1
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(fi19, v3)
        _ = fi19
    }
    if false {
        switch v13 {
        case -9223372036854775808:
            v13 += 301
        }
        switch v3 {
        case "a":
            v20 := GenS3{F0: 127, F1: 749}
            _ = v20
        case "foo bar":
            for fi21 := 0; fi21 < 16; fi21++ {
                if true {
                    fmt.Println(v13)
                    v13 = fi21
                }
                v22 := GenS4{F0: 42, F1: v13}
                _ = len(v2)
                _ = fi21
                _ = v22
            }
        case "\n\t":
            for fi23 := 0; fi23 < 20; fi23++ {
                switch int(v4) {
                case -825:
                    if false {
                        v24 := GenS4{F0: (v18 | v18), F1: (v13 + int(v1))}
                        _ = v24
                    }
                case 22:
                    if (42 <= 242) {
                        var v25 int32 = (v18 | v18)
                        _ = v25
                    }
                }
                v3 = v3
                v13 = -31
                _ = fi23
            }
        default:
            v2[0] = (!(!(false && true)))
        }
    }
    v2[(v1 << ((2) & 63))] = false
    v26 := GenS2{F0: 127, F1: 255}
    v2[v1] = (v1 < 65535)
    _ = len(v2)
    var v27 uint = 303
    v1 = (v1 << (((v27 ^ v27)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
