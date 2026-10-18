package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int; F1 int16; F2 int16 }
type GenS4 struct { F0 int8; F1 bool }
func genFn0(p0 uint16) uint {
    v6 := []int{9223372036854775807, 127, -41, 572}
    v7 := GenS1{F0: 374}
    p0 += (p0 >> ((0) & 63))
    var v8 uint32 = 1
    var v9 string = ("a" + ("\n\t" + "hello"))
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint, p1 uint8, p2 uint) int64 {
    for fi10 := 0; fi10 < 11; fi10++ {
        p0 += p2
        p1 = p1
        if (491 > 620) {
            v11 := []int{fi10, 4294967295, fi10, fi10}
            var v12 int = (fi10 & (fi10 | 2147483647))
            fmt.Println(p0, p2)
            _ = v11
            _ = v12
        } else {
            p2 += p0
            v13 := []int{}
            _ = v13
        }
        _ = fi10
    }
    fmt.Println(p2, p1, p0)
    v14 := []int{-971, 4294967295, 128, 1}
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = true
    var v3 bool = false
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    v2 = v2
    v15 := v4[345]
    v15 = v15
    v2 = false
    v2 = v3
    v15 = v15
    var v16 bool = (!(v3 && (2147483648 >= 65535)))
    v17 := GenS2{F0: -1, F1: v15}
    if v3 {
        for fi18 := 0; fi18 < 17; fi18++ {
            fmt.Println(v15, v3, v16)
            v3 = (2147483647 != 111)
            if false {
                fmt.Println(v16)
                if true {
                    switch v15 {
                    case "\n\t":
                        v3 = (1 >= 4294967295)
                    default:
                        v15 = v15
                    }
                } else {
                    switch "a" {
                    case "hello":
                        v15 = v15
                    default:
                        v15 = v15
                    }
                    if (!v3) {
                        var v19 uint = 626
                        for fi20 := 0; fi20 < 17; fi20++ {
                            var v21 uint16 = 0
                            if v16 {
                                v15 = "a"
                                var v22 int = fi18
                                v23 := GenS1{F0: 2}
                                _ = v22
                                _ = v23
                            }
                            _ = fi20
                            _ = v21
                        }
                        var v24 float32 = (0.0 * 0.1)
                        _ = v19
                        _ = v24
                    } else {
                        v1[("" + ((v15 + v15) + v15))] = 65535
                    }
                }
            }
            _ = fi18
        }
        v25 := GenS1{F0: 54}
        v3 = true
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 6; fi27++ {
            delete(v4, 2)
            _ = fi27
        }
        _ = fi26
    }
    if (v2 && true) {
        fmt.Println(v15, v3)
    }
    var v28 int64 = -1
    v16 = (!("" != v15))
    v29 := v5[693]
    switch (v15 + (v15 + v15)) {
    case "\n\t":
        v29 -= (v29 * v29)
    case "":
        switch -9223372036854775808 {
        case 32767:
            v3 = v3
        case 273:
            for fi30 := 0; fi30 < 9; fi30++ {
                v31 := GenS3{F0: fi30, F1: 1, F2: 32767}
                fmt.Println(v16, v29, v2)
                _ = len(v1)
                _ = fi30
                _ = v31
            }
        }
    case "a":
        v32 := GenS4{F0: 90, F1: (genFn0(1) == genFn0(32767))}
        _ = v32
    default:
        _ = len(v1)
    }
    v29 += v29
    v29 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
