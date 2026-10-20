package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS4 struct { F0 bool }
func genFn0() int32 {
    v7 := GenS1{F0: -128}
    if (!true) {
        var v8 int8 = 0
        for fi9 := 0; fi9 < 17; fi9++ {
            if true {
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 int32 = 42
                    _ = fi10
                    _ = v11
                }
                v8 -= (v8 | v8)
            } else {
                v8 = (v8 - -90)
                v12 := GenS1{F0: 42}
                _ = v12
            }
            _ = fi9
        }
        _ = v8
    }
    _ = v7
    return 2147483647
}
func genFn1() int16 {
    if (496 != 4294967295) {
        var v13 uint64 = 4294967295
        fmt.Println(v13)
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 bool = false
        if (2 != 4294967295) {
            if ((-948 == 1) && v15) {
                v16 := GenS1{F0: 100}
                v17 := []int{72, -9223372036854775808, fi14}
                _ = v16
                _ = v17
            }
        } else {
            v18 := GenS1{F0: 1}
            v15 = (!(false || v15))
            _ = v18
        }
        if (v15 && (!(!(v15 || (152 >= 74))))) {
            if v15 {
                if false {
                    var v19 int16 = 1
                    if v15 {
                        v20 := []int{}
                        _ = v20
                    }
                    _ = v19
                }
            } else {
                if false {
                    v15 = (("foo bar" > ("\n\t" + "a")) || (!(1 > 2)))
                    for fi21 := 0; fi21 < 18; fi21++ {
                        v22 := GenS1{F0: 1}
                        _ = fi21
                        _ = v22
                    }
                }
                v23 := GenS2{F0: 6}
                _ = v23
            }
        }
        _ = fi14
        _ = v15
    }
    v24 := GenS3{F0: 2, F1: -980, F2: 1}
    _ = v24
    return 193
}

func main() {
    v1 := []int{}
    v2 := make(map[string]bool)
    var v3 int16 = -1
    var v4 float64 = (3.14159 * 0.1)
    v5 := make(map[int]int)
    var v6 int32 = 94
    var v25 uint8 = 0
    var v26 uint32 = 0
    v25 += uint8(v26)
    fmt.Println(v6, v4)
    v6 = v6
    for fi27 := 0; fi27 < 13; fi27++ {
        fmt.Println(v6, fi27)
        if (!false) {
            v3 += (v3 << ((567) & 63))
        }
        v28 := GenS3{F0: (v26 - v26), F1: (v6 * genFn0()), F2: (v26 << ((42) & 63))}
        _ = fi27
        _ = v28
    }
    if true {
        var v29 uint16 = 42
        v30 := v5[127]
        v31 := GenS3{F0: v26, F1: (v6 * (v6 / (((v6 | (v6 & v6))) | 1))), F2: 0}
        _ = v29
        _ = v30
        _ = v31
    } else {
        v32 := GenS4{F0: ("foo bar" == ("" + ((("a" + "") + "foo bar") + "foo bar")))}
        v25 += (v25 ^ v25)
        v33 := GenS4{F0: true}
        _ = v32
        _ = v33
    }
    var v34 uint16 = 521
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
