package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 int }
func genFn0(p0 int32, p1 int) int {
    fmt.Println(p1)
    if false {
        fmt.Println(p0, p0)
        for fi6 := 0; fi6 < 10; fi6++ {
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := GenS1{F0: 18446744073709551615, F1: 0}
                fmt.Println(p0, fi7, fi7)
                if (1 < 1000) {
                    p0 += (p0 / (((p0 + (p0 << ((4294967295) & 63)))) | 1))
                    var v9 int32 = (p0 / ((p0) | 1))
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
            for fi10 := 0; fi10 < 5; fi10++ {
                var v11 float32 = 3.14159
                v12 := []int{fi10, fi10}
                var v13 string = "foo bar"
                switch (p1 & fi10) {
                case -227:
                    var v14 uint = 4294967295
                    _ = v14
                case -9223372036854775808:
                    var v15 uint8 = 0
                    _ = v15
                default:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            for fi16 := 0; fi16 < 10; fi16++ {
                p1 = (p1 & (^(fi16)))
                v17 := []int{(fi16 & 52), 858, (p1 ^ (fi16 + (p1 >> ((4294967295) & 63))))}
                _ = fi16
                _ = v17
            }
            _ = fi6
        }
        p0 = p0
    }
    if (!(4294967295 > 1)) {
        fmt.Println(p0, p0)
        p0 = p0
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 uint, p1 int8) bool {
    v18 := GenS1{F0: 18446744073709551615, F1: p0}
    p1 += 42
    var v19 uint64 = 271
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return false
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 255
    var v3 uint = v2
    var v4 uint16 = (v1 >> (((v3 - (^(v2)))) & 63))
    var v5 string = ("hello" + "foo bar")
    fmt.Println(v1)
    var v20 int16 = 238
    v21 := []int{-1, genFn0(2147483647, -9223372036854775808)}
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 bool = (!false)
        _ = fi22
        _ = v23
    }
    var v24 uint64 = 4294967295
    v5 = (("foo bar" + "") + "a")
    var v25 uint32 = 2
    _ = len(v21)
    if len(v21) > 0 {
        v21[0] = -333
    }
    switch 42 {
    case 32768:
        var v26 uint64 = v24
        _ = v26
    case 42:
        v5 = v5
    }
    var v27 int32 = 42
    v4 = v4
    var v28 uint32 = 2147483648
    for fi29 := 0; fi29 < 16; fi29++ {
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
        if (!((v28 / (((v25 - (v28 % ((1) | 1)))) | 1)) < (v28 | v25))) {
            fmt.Println(v4)
            var v30 int = fi29
            _ = v30
        } else {
            v5 = v5
        }
        v24 = (^(v24))
        _ = fi29
    }
    var v31 uint8 = 42
    var v32 float64 = (1.5 / 3.14159)
    switch 42 {
    case 1:
        v33 := GenS2{F0: v27, F1: 4294967295, F2: 835}
        _ = v33
    case 32768:
        v34 := GenS1{F0: v24, F1: (v2 ^ 255)}
        _ = v34
    case 948:
        v35 := GenS1{F0: v24, F1: v3}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
