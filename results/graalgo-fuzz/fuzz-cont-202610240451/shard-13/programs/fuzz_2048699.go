package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int64 }
type GenS2 struct { F0 uint64; F1 bool; F2 int }
type GenS3 struct { F0 string; F1 int64; F2 uint64 }
type GenS4 struct { F0 string }
func genFn0(p0 uint32) int {
    if (-2.25 < 0.1) {
        p0 = p0
        if false {
            if (false && true) {
                v7 := GenS1{F0: 255, F1: ("foo bar" + "a"), F2: 9223372036854775807}
                p0 += p0
                _ = v7
            }
        }
    }
    if false {
        if false {
            p0 = 0
            for fi8 := 0; fi8 < 20; fi8++ {
                if true {
                    v9 := GenS1{F0: 32767, F1: "hello", F2: -1}
                    fmt.Println(p0, p0)
                    _ = v9
                }
                _ = fi8
            }
        } else {
            var v10 uint = 0
            _ = v10
        }
    } else {
        if false {
            p0 = p0
        }
        v11 := GenS1{F0: 1, F1: "", F2: 9223372036854775807}
        var v12 float64 = (3.14159 - -2.25)
        _ = v11
        _ = v12
    }
    if false {
        fmt.Println(p0)
        v13 := []int{42, -2147483648, 1}
        if false {
            _ = len(v13)
        }
        _ = v13
    } else {
        v14 := GenS2{F0: 2, F1: true, F2: 140}
        fmt.Println(p0)
        _ = v14
    }
    v15 := GenS2{F0: 18446744073709551615, F1: (!false), F2: 127}
    _ = p0
    _ = v15
    return 9223372036854775807
}
func genFn1(p0 uint, p1 int32) uint64 {
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint8 = 255
        _ = fi16
        _ = v17
    }
    v18 := GenS3{F0: "", F1: 22, F2: 619}
    switch genFn0(274) {
    case -1:
        var v19 uint16 = 32767
        _ = v19
    case 32767:
        p1 -= p1
    case 0:
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 uint64 = 4294967295
            v22 := []int{fi20, fi20, (-(fi20))}
            _ = fi20
            _ = v21
            _ = v22
        }
    default:
        v23 := GenS3{F0: (("" + ("foo bar" + "hello")) + ""), F1: -9223372036854775808, F2: 408}
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v18
    return 197
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{2147483648, 2147483647}
    var v3 int8 = -128
    v4 := []int{633, 32767, -291}
    var v5 uint = 65535
    var v6 uint16 = 589
    v6 -= (v6 * (v6 / ((42) | 1)))
    _ = len(v2)
    v5 = v5
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 int8 = (v3 - v3)
        _ = fi24
        _ = v25
    }
    v26 := GenS4{F0: v1}
    v6 -= (v6 + (v6 | v6))
    for fi27 := 0; fi27 < 15; fi27++ {
        v6 = 2
        _ = fi27
    }
    var v28 int = 42
    v29 := GenS1{F0: (v6 >> (((^(v5))) & 63)), F1: "foo bar", F2: 791}
    var v30 int = v28
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
