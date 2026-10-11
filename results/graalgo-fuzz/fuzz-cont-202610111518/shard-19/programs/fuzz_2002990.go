package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
type GenS3 struct { F0 int64; F1 int64; F2 int32 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint16, p1 string, p2 int8) bool {
    v4 := GenS1{F0: 4294967295, F1: (((p1 + p1) + p1) + p1)}
    var v5 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}
func genFn1() int64 {
    v6 := []int{42, 514, 32767, -2147483648}
    var v7 string = ("hello" + "foo bar")
    _ = v6
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 127
    v2 := []string{}
    var v3 int16 = -32768
    if true {
        v3 = 0
        v8 := GenS1{F0: 208, F1: "\n\t"}
        _ = v8
    }
    var v9 float32 = (1e10 * 0.1)
    var v10 float32 = (0.0 - v9)
    var v11 int16 = -737
    var v12 float32 = v9
    v13 := GenS2{F0: false, F1: 255, F2: (!false)}
    for fi14 := 0; fi14 < 2; fi14++ {
        v9 = float32((1e10 + float64((0.1 / v10))))
        _ = fi14
    }
    v2 = append(v2, "\n\t")
    v15 := GenS1{F0: 1, F1: "a"}
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v16 := GenS3{F0: 0, F1: 42, F2: 700}
    _ = len(v2)
    v17 := GenS3{F0: 9223372036854775807, F1: 42, F2: 42}
    fmt.Println(v12)
    if (float32(0.1) == v10) {
        for fi18 := 0; fi18 < 15; fi18++ {
            if true {
                switch fi18 {
                case 32767:
                    v1 += (^(v1))
                default:
                    var v19 uint = 732
                    _ = v19
                }
                var v20 uint32 = 0
                v12 += -2.25
                _ = v20
            } else {
                switch fi18 {
                case -1:
                    v21 := GenS4{F0: (v11 << ((65535) & 63)), F1: 100}
                    _ = v21
                }
                if (false || false) {
                    v2 = append(v2, "a")
                    v12 += 0.1
                    v9 = 1e10
                } else {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v23 := GenS5{F0: false}
                        if len(v2) > 0 {
                            v2[0] = ("\n\t" + "\n\t")
                        }
                        var v24 uint = 806
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS2{F0: false, F1: 632, F2: (false && false)}
                    if true {
                        v10 -= v10
                    } else {
                        var v26 int = fi18
                        v27 := GenS3{F0: 42, F1: genFn1(), F2: 578}
                        _ = v26
                        _ = v27
                    }
                    _ = v25
                }
            }
            _ = len(v2)
            _ = fi18
        }
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v9
}
