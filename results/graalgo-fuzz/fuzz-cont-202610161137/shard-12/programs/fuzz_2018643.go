package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 bool }
type GenS2 struct { F0 uint; F1 uint64; F2 uint }

func main() {
    var v1 uint64 = 2
    var v2 uint = 443
    var v3 uint64 = (v1 >> (((v2 << ((v2) & 63))) & 63))
    var v4 float32 = 1e10
    v2 = v2
    for fi5 := 0; fi5 < 3; fi5++ {
        v1 = 0
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: 991, F1: (("\n\t" + "\n\t") + ("hello" + "a")), F2: true}
            _ = fi6
            _ = v7
        }
        v8 := []int{}
        _ = fi5
        _ = v8
    }
    if false {
        v1 += 18446744073709551615
        var v9 int64 = 0
        _ = v9
    } else {
        if false {
            v10 := []int{}
            for fi11 := 0; fi11 < 20; fi11++ {
                if true {
                    var v12 int16 = 884
                    var v13 float64 = (0.1 - (1e10 + -2.25))
                    _ = v12
                    _ = v13
                }
                switch "hello" {
                case "foo bar":
                    var v14 int64 = 1
                    _ = v14
                default:
                    fmt.Println(v4)
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi11
            }
            var v15 float64 = 1.5
            _ = v10
            _ = v15
        } else {
            for fi16 := 0; fi16 < 5; fi16++ {
                switch "hello" {
                case "a":
                    v2 -= v2
                case "":
                    v17 := []int{32768, (fi16 - -387), fi16, fi16}
                    _ = v17
                case "\n\t":
                    v3 = v1
                default:
                    v18 := GenS1{F0: 0, F1: "\n\t", F2: true}
                    _ = v18
                }
                var v19 float64 = (-2.25 / 1.5)
                v20 := []int{415, fi16, 235, (fi16 * fi16)}
                _ = fi16
                _ = v19
                _ = v20
            }
        }
    }
    v21 := GenS1{F0: -672, F1: "a", F2: ((false || (1 == 42)) || (v4 > (-2.25 / float32(0.1))))}
    if true {
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := []int{fi22, -2147483648}
            switch ("a" + (("foo bar" + ("foo bar" + "hello")) + "hello")) {
            case "foo bar":
                var v24 int = fi22
                _ = v24
            case "\n\t":
                if (!(0.0 == v4)) {
                    fmt.Println(fi22)
                    switch 9223372036854775807 {
                    case 42:
                        var v25 float64 = 0.1
                        _ = v25
                    default:
                        v4 -= ((((v4 / 0.1) * (float32(3.14159) - 0.0)) - v4) + (((float32(0.0) * -2.25) / (v4 + v4)) / (v4 - v4)))
                    }
                } else {
                    if true {
                        var v26 float32 = v4
                        _ = v26
                    }
                }
            default:
                var v27 bool = (!false)
                _ = v27
            }
            _ = fi22
            _ = v23
        }
    } else {
        var v28 uint8 = 1
        var v29 int64 = 42
        _ = v28
        _ = v29
    }
    v2 -= v2
    v30 := GenS2{F0: v2, F1: 4294967295, F2: (v2 << ((v2) & 63))}
    fmt.Println(v3)
    fmt.Println(v3, v2, v3)
    v31 := []int{-854}
    v3 -= (v3 ^ (v1 - v3))
    var v32 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
