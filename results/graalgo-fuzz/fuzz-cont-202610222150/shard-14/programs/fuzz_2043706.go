package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 int8 }
func genFn0() uint32 {
    var v4 uint16 = 824
    v5 := GenS1{F0: (^(v4))}
    v6 := []int{}
    if false {
        var v7 float64 = (0.1 + float64(3.14159))
        var v8 uint64 = 2
        v8 = v8
        _ = v7
        _ = v8
    } else {
        var v9 uint64 = 18446744073709551615
        if false {
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := GenS2{F0: fi10, F1: 2147483647, F2: 1}
                _ = fi10
                _ = v11
            }
        }
        _ = v9
    }
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int64) int32 {
    p0 = (^(p0))
    if true {
        var v12 float32 = 1.5
        var v13 string = "a"
        p0 -= (p0 << ((18446744073709551615) & 63))
        _ = v12
        _ = v13
    } else {
        var v14 string = ("\n\t" + ("" + (((("a" + "foo bar") + "hello") + "") + "")))
        _ = v14
    }
    var v15 int64 = (p0 ^ (p0 * -9223372036854775808))
    v16 := GenS1{F0: 65535}
    _ = p0
    _ = v15
    _ = v16
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -1
    v3 := []int{0}
    var v17 uint = 2
    v2 += v2
    v2 = (v2 % (((-(v2))) | 1))
    var v18 string = (("\n\t" + "foo bar") + "")
    v18 = v18
    var v19 string = ((((v18 + (v18 + (v18 + (((v18 + "") + ("" + v18)) + (v18 + v18))))) + (v18 + "a")) + v18) + (v18 + v18))
    v2 = (v2 + (v2 * (v2 + (v2 & v2))))
    if (!true) {
        switch 1 {
        case 9223372036854775807:
            for fi20 := 0; fi20 < 5; fi20++ {
                if false {
                    v1[678] = 1
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v22 := GenS1{F0: 0}
                        v23 := GenS2{F0: fi20, F1: 0, F2: -75}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    var v24 float32 = (-2.25 - 0.0)
                    _ = v24
                } else {
                    v25 := GenS2{F0: 4294967295, F1: 4294967295, F2: -82}
                    v18 = (v19 + "")
                    _ = v25
                }
                _ = fi20
            }
        case 2147483647:
            delete(v1, 496)
        case -96:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            var v26 int16 = 0
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
