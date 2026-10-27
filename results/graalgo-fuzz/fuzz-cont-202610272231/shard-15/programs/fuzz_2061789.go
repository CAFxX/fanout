package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 int8; F1 int8; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint; F1 int8; F2 int }
func genFn0(p0 int16, p1 int32) bool {
    fmt.Println(p0, p0)
    v5 := GenS1{F0: 42, F1: false}
    if false {
        for fi6 := 0; fi6 < 3; fi6++ {
            for fi7 := 0; fi7 < 20; fi7++ {
                if (false || ((^(p0)) > p0)) {
                    if false {
                        p1 = (p1 & (p1 ^ (p1 ^ p1)))
                        var v8 int32 = -2147483648
                        if true {
                            p0 = p0
                            fmt.Println(p1, v8)
                        }
                        _ = v8
                    } else {
                        v9 := GenS1{F0: 60, F1: false}
                        p0 += p0
                        var v10 int8 = 100
                        _ = v10
                        _ = v9
                    }
                }
                v11 := []int{fi6, 2147483647, (fi7 * (fi7 * (fi7 + -9223372036854775808)))}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                v12 := GenS1{F0: int64(p0), F1: false}
                var v13 string = ""
                v14 := GenS2{F0: 1, F1: 952, F2: -32768}
                _ = fi7
                _ = v11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi6
        }
        fmt.Println(p1)
    }
    v15 := GenS2{F0: 46, F1: 970, F2: 1}
    _ = p0
    _ = p1
    _ = v15
    _ = v5
    return (!(344 <= 425))
}
func genFn1(p0 int16, p1 int16) int32 {
    var v16 uint32 = 4294967295
    var v17 int = 2147483647
    v18 := GenS2{F0: 42, F1: 24, F2: p1}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    var v3 int32 = 42
    var v4 float32 = 1e10
    switch "a" {
    case "":
        delete(v2, 621)
    }
    v19 := GenS2{F0: 127, F1: 32767, F2: 351}
    var v20 bool = false
    v21 := []int{32767, 4294967295}
    v22 := GenS3{F0: 255, F1: 42}
    var v23 float64 = 0.0
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 bool = true
        if len(v21) > 0 {
            v21[0] = (fi24 << ((1) & 63))
        }
        _ = fi24
        _ = v25
    }
    v23 = v23
    for fi26 := 0; fi26 < 8; fi26++ {
        if true {
            v27 := GenS4{F0: -91, F1: 127, F2: 4294967295}
            _ = v27
        }
        v28 := GenS5{F0: 100}
        v4 += 3.14159
        _ = fi26
        _ = v28
    }
    var v29 bool = (false || v20)
    v30 := GenS6{F0: false}
    fmt.Println(v3)
    v29 = (153 >= 4294967295)
    fmt.Println(v29, v20)
    for fi31 := 0; fi31 < 5; fi31++ {
        if ((!(18446744073709551615 <= 42)) || (2147483648 > 2147483648)) {
            v32 := GenS7{F0: 1, F1: 100, F2: (fi31 % ((fi31) | 1))}
            _ = len(v2)
            _ = v32
        }
        var v33 bool = (v29 || (("" + "a") > ("" + "a")))
        _ = fi31
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
