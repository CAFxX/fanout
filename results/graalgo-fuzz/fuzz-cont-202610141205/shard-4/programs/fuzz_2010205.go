package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64) int {
    var v6 int32 = 0
    if true {
        fmt.Println(v6, v6)
        var v7 uint = 4294967295
        _ = v7
    } else {
        p0 = (p0 & (p0 & (p0 & p0)))
    }
    p0 += p0
    v8 := []int{-98, 9223372036854775807, -1}
    _ = p0
    _ = v6
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 uint = 2
    var v2 bool = ((32767 >= 1) || (255 <= 107))
    var v3 int8 = 69
    var v4 float32 = ((float32((float64(0.1) * float64(float32(3.14159)))) / 1e10) + -2.25)
    v5 := make(map[uint32]string)
    if false {
        var v9 int8 = v3
        v10 := v5[42]
        _ = v10
        _ = v9
    } else {
        v4 += v4
        switch ("" + "\n\t") {
        case "a":
            delete(v5, 134)
        }
    }
    var v11 uint64 = 0
    v12 := GenS1{F0: "a", F1: 378}
    v1 = (v1 % ((4294967295) | 1))
    var v13 int64 = -1
    v4 = ((0.1 - v4) + v4)
    switch "hello" {
    case "hello":
        v13 += v13
    default:
        v14 := GenS1{F0: "hello", F1: -745}
        _ = v14
    }
    var v15 float32 = (v4 + float32(-2.25))
    if v2 {
        switch int(v3) {
        case 32767:
            switch 128 {
            case 0:
                for fi16 := 0; fi16 < 14; fi16++ {
                    for fi17 := 0; fi17 < 20; fi17++ {
                        if (42 != -24) {
                            v18 := GenS1{F0: ("\n\t" + "foo bar"), F1: 1}
                            if v2 {
                                delete(v5, 399)
                                if ((fi16 | fi16) <= (^(fi16))) {
                                    if true {
                                        v19 := GenS1{F0: ("\n\t" + "foo bar"), F1: 2147483647}
                                        _ = v19
                                    }
                                }
                                if v2 {
                                    fmt.Println(v4)
                                    if v2 {
                                        v20 := GenS1{F0: (("\n\t" + ("hello" + "\n\t")) + "hello"), F1: -44}
                                        _ = v20
                                    }
                                    var v21 uint64 = (^(v11))
                                    _ = v21
                                }
                            }
                            _ = v18
                        }
                        fmt.Println(v2, v3)
                        fmt.Println(v4, v4)
                        _ = fi17
                    }
                    if false {
                        var v22 uint = (^(v1))
                        v23 := GenS2{F0: 42, F1: 993}
                        _ = v22
                        _ = v23
                    }
                    _ = fi16
                }
            }
        }
    } else {
        v24 := v5[4294967295]
        _ = v24
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
