package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string }
func genFn0(p0 uint32, p1 int8, p2 int8) int {
    v4 := GenS1{F0: 1}
    var v5 uint16 = 534
    v6 := GenS1{F0: (v5 | v5)}
    var v7 uint16 = 589
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint) uint8 {
    switch 2147483647 {
    case -9223372036854775808:
        switch 128 {
        case 1:
            v8 := []int{genFn0(0, 100, 3), -961, -1}
            _ = v8
        case 814:
            if true {
                p0 -= p0
                v9 := []int{280}
                _ = v9
            }
        case -9223372036854775808:
            fmt.Println(p0)
        default:
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 int = fi10
                _ = fi10
                _ = v11
            }
        }
    case 2147483647:
        var v12 uint = (p0 + (p0 << (((p0 | (p0 % (((p0 >> (((p0 + (^(p0)))) & 63))) | 1)))) & 63)))
        _ = v12
    default:
        if false {
            var v13 int64 = 0
            v14 := []int{32767, 42}
            _ = v13
            _ = v14
        } else {
            p0 = (p0 ^ 2)
            v15 := []int{627, genFn0(65535, 127, 100), 0, 0}
            for fi16 := 0; fi16 < 3; fi16++ {
                p0 = 81
                var v17 int32 = 2147483647
                v18 := GenS1{F0: 65535}
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v15
        }
    }
    p0 = 1
    v19 := []int{127, -9223372036854775808, 633}
    _ = p0
    _ = v19
    return 42
}

func main() {
    var v1 int = 2147483647
    var v2 uint = 449
    var v3 int16 = 32767
    if (((155 > 128) || (true && false)) || ("foo bar" > ("foo bar" + "foo bar"))) {
        switch "hello" {
        case "hello":
            for fi20 := 0; fi20 < 19; fi20++ {
                for fi21 := 0; fi21 < 18; fi21++ {
                    v22 := GenS2{F0: -884, F1: 0}
                    v23 := []int{}
                    if (9223372036854775807 < -9223372036854775808) {
                        fmt.Println(fi20, v1, v2)
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    v25 := GenS3{F0: 0, F1: 744, F2: (v2 & (v2 / ((255) | 1)))}
                    var v26 uint = v2
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                v27 := GenS4{F0: 2147483647}
                _ = fi20
                _ = v27
            }
        case "a":
            v2 -= 2
        }
    } else {
        v28 := []int{(^(v1)), v1, (v1 & v1), (v1 / (((-(v1))) | 1))}
        _ = v28
    }
    v1 = v1
    var v29 uint16 = 42
    v30 := GenS5{F0: (v1 | 1)}
    if false {
        var v31 uint8 = 255
        var v32 string = ("hello" + "a")
        _ = v31
        _ = v32
    } else {
        var v33 int = v1
        v34 := GenS1{F0: (v29 % ((32767) | 1))}
        _ = v33
        _ = v34
    }
    v2 = v2
    v3 -= (v3 ^ int16(v29))
    v35 := GenS6{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
}
