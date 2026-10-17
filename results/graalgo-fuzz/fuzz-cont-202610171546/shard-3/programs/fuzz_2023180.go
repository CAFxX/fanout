package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint, p1 string) uint8 {
    p0 = p0
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := []int{(fi5 ^ 9223372036854775807), 42, (fi5 << ((p0) & 63)), 127}
        p1 = ("hello" + p1)
        _ = fi5
        _ = v6
    }
    var v7 int32 = 0
    v8 := GenS1{F0: 87}
    if true {
        for fi9 := 0; fi9 < 6; fi9++ {
            if true {
                p1 = (p1 + p1)
                for fi10 := 0; fi10 < 16; fi10++ {
                    v7 = -2147483648
                    p1 = "a"
                    _ = fi10
                }
                v11 := []int{}
                _ = v11
            }
            p0 = p0
            _ = fi9
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 201
}
func genFn1(p0 uint64) int {
    for fi12 := 0; fi12 < 12; fi12++ {
        p0 -= p0
        _ = fi12
    }
    p0 += p0
    v13 := []int{-1, -2147483648, 32767, 4294967295}
    switch "\n\t" {
    case "\n\t":
        p0 = 976
    }
    var v14 uint8 = 46
    _ = p0
    _ = v13
    _ = v14
    return 32768
}

func main() {
    var v1 uint32 = 412
    v2 := []uint32{(v1 | (v1 + v1)), (v1 | v1), (v1 | v1)}
    var v3 uint32 = 196
    v4 := []uint32{}
    v15 := GenS1{F0: 931}
    switch 127 {
    case 4294967295:
        var v16 uint = 255
        _ = v16
    default:
        v2 = append(v2, (^(v1)))
    }
    v17 := GenS1{F0: 1}
    if len(v2) > 0 {
        v2[0] = (v3 | 0)
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := GenS1{F0: 42}
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: (v3 % (((v1 << ((18446744073709551615) & 63))) | 1)), F1: (("foo bar" + ("" + "foo bar")) + "")}
    var v21 uint8 = 1
    v22 := GenS3{F0: v3, F1: 18446744073709551615}
    if ((false && false) && (2147483647 == 43)) {
        for fi23 := 0; fi23 < 9; fi23++ {
            v24 := GenS2{F0: 1, F1: "\n\t"}
            _ = fi23
            _ = v24
        }
        v1 += v3
        if (2 >= (v3 * v1)) {
            v21 -= (v21 & 102)
            _ = len(v4)
            var v25 int32 = -671
            _ = v25
        }
    }
    fmt.Println(v21)
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if false {
        if len(v2) > 0 {
            v2[0] = (^(v3))
        }
        _ = len(v4)
        v21 = v21
    }
    var v26 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
}
