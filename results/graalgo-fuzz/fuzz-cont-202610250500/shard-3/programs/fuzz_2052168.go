package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }
type GenS2 struct { F0 bool; F1 bool }
func genFn0() int32 {
    var v5 int64 = -9223372036854775808
    fmt.Println(v5)
    var v6 float64 = (((0.1 + 1e10) - float64(0.1)) - 0.1)
    var v7 string = "hello"
    v6 = (v6 / float64((-2.25 / 0.1)))
    _ = v5
    _ = v6
    _ = v7
    return 604
}
func genFn1(p0 int32) int64 {
    p0 = p0
    p0 += (p0 >> ((179) & 63))
    _ = p0
    return -594
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 65535
    var v3 int32 = 42
    var v4 float64 = float64(1.5)
    switch 32768 {
    case 32767:
        v8 := GenS1{F0: ("\n\t" + "a"), F1: 0, F2: 0}
        _ = v8
    default:
        v2 += 2
    }
    if (("hello" + "foo bar") != "") {
        v3 = v3
        var v9 string = ("\n\t" + "\n\t")
        if false {
            if (!(!(246 <= 0))) {
                v9 = v9
                var v10 float32 = -2.25
                for fi11 := 0; fi11 < 9; fi11++ {
                    v9 = v9
                    v12 := GenS1{F0: ((v9 + "\n\t") + "foo bar"), F1: 1, F2: -964}
                    var v13 uint16 = (v2 % ((v2) | 1))
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = v10
            } else {
                switch -9223372036854775808 {
                case -2147483648:
                    for fi14 := 0; fi14 < 17; fi14++ {
                        switch fi14 {
                        case 1:
                            fmt.Println(v9, v4)
                        }
                        _ = fi14
                    }
                }
                var v15 uint8 = 255
                _ = v15
            }
        }
        _ = v9
    }
    var v16 uint = 422
    v1[42] = 4294967295
    fmt.Println(v4, v3)
    v4 += ((-2.25 / 1.0) - v4)
    if ((false && true) && true) {
        var v17 float32 = 1e10
        fmt.Println(v2, v2)
        var v18 string = ""
        _ = v17
        _ = v18
    }
    v2 += 42
    fmt.Println(v2)
    v2 = (v2 << ((v16) & 63))
    fmt.Println(v2, v4, v16)
    fmt.Println(v2, v16, v16)
    v19 := GenS2{F0: (165 != 753), F1: true}
    v1[42] = 855
    v4 = 1.5
    for fi20 := 0; fi20 < 3; fi20++ {
        if false {
            var v21 uint32 = 0
            if true {
                switch "" {
                case "hello":
                    fmt.Println(v21)
                case "foo bar":
                    v22 := []int{}
                    _ = v22
                default:
                    v3 -= (-(v3))
                }
            }
            _ = v21
        }
        _ = fi20
    }
    v3 += genFn0()
    v16 = (v16 >> ((42) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
