package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint32; F2 int16 }
type GenS4 struct { F0 int8; F1 uint; F2 int32 }
func genFn0(p0 int16, p1 uint16, p2 uint16) int8 {
    if false {
        var v7 uint64 = 0
        p0 = 32767
        _ = v7
    } else {
        p2 = p2
        v8 := GenS1{F0: 1, F1: 2147483647}
        _ = v8
    }
    fmt.Println(p0)
    if false {
        var v9 uint8 = 221
        p0 = 1000
        _ = v9
    } else {
        p1 = (^(p1))
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}
func genFn1(p0 uint64, p1 int32, p2 uint32) uint64 {
    if false {
        var v10 bool = true
        if (225 < 42) {
            p1 += p1
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := []int{}
                fmt.Println(v10)
                var v13 string = ""
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := GenS1{F0: 1, F1: p1}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 3; fi15++ {
                if v10 {
                    switch 4294967295 {
                    case 32767:
                        if v10 {
                            v16 := GenS1{F0: 1, F1: (-(p1))}
                            _ = v16
                        }
                    case 121:
                        for fi17 := 0; fi17 < 7; fi17++ {
                            var v18 uint16 = 316
                            _ = fi17
                            _ = v18
                        }
                    default:
                        p2 -= (p2 ^ p2)
                    }
                    fmt.Println(p1, v10, v10)
                }
                var v19 uint32 = p2
                v20 := []int{int(p0)}
                _ = fi15
                _ = v19
                _ = v20
            }
            v21 := []int{-513, 42, 0, -533}
            _ = v21
        }
        fmt.Println(p0)
        _ = v10
    }
    p1 = (p1 % ((p1) | 1))
    if false {
        switch ("\n\t" + "") {
        case "":
            v22 := []int{-9223372036854775808, -436, -1}
            _ = v22
        default:
            v23 := []int{}
            _ = v23
        }
    } else {
        p1 = p1
        v24 := []int{128, 127, 325, 9223372036854775807}
        p2 -= 4294967295
        _ = v24
    }
    v25 := GenS2{F0: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return uint64(p2)
}

func main() {
    var v1 uint32 = 190
    var v2 int32 = -1
    v3 := []int{}
    v4 := []int{32768}
    v5 := make(map[int]string)
    var v6 bool = false
    v5[2147483647] = ""
    for fi26 := 0; fi26 < 12; fi26++ {
        v1 = (v1 | (^(v1)))
        v2 -= v2
        fmt.Println(fi26, v1)
        _ = fi26
    }
    fmt.Println(v6, v2, v2)
    v1 = (v1 >> ((255) & 63))
    v27 := GenS2{F0: 9223372036854775807}
    var v28 uint32 = (v1 - v1)
    if true {
        switch "hello" {
        case "":
            v5[-1] = ("" + ("foo bar" + (("" + "hello") + "a")))
        case "hello":
            delete(v5, -9223372036854775808)
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            v30 := GenS2{F0: (fi29 * (fi29 * fi29))}
            v6 = true
            _ = fi29
            _ = v30
        }
        v1 = (^(v28))
    }
    fmt.Println(v6)
    var v31 uint8 = 120
    var v32 int8 = 42
    v33 := GenS3{F0: 18446744073709551615, F1: v28, F2: 32767}
    var v34 string = "foo bar"
    for fi35 := 0; fi35 < 18; fi35++ {
        v28 += (v1 + v28)
        v34 = v34
        _ = fi35
    }
    v1 -= 572
    var v36 int64 = 0
    v37 := GenS4{F0: v32, F1: 2, F2: (v2 << ((896) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
