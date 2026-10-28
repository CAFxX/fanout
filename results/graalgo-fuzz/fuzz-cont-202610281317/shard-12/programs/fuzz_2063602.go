package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0(p0 uint16, p1 string) uint16 {
    var v4 uint64 = 0
    switch (p1 + "a") {
    case "foo bar":
        v5 := []int{42, 42, 988}
        _ = v5
    default:
        fmt.Println(v4, p1)
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{fi6, fi6, (-(fi6))}
        var v8 string = p1
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    return (p0 * (p0 % ((p0) | 1)))
}
func genFn1(p0 uint32, p1 int8) int8 {
    fmt.Println(p1)
    fmt.Println(p1)
    switch -2147483648 {
    case 138:
        p0 -= p0
    case -451:
        switch 128 {
        case 32767:
            if true {
                for fi9 := 0; fi9 < 15; fi9++ {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        if false {
                            v11 := GenS1{F0: 778, F1: 797}
                            _ = v11
                        }
                        _ = fi10
                    }
                    p0 += p0
                    v12 := []int{(fi9 * (^(fi9)))}
                    _ = fi9
                    _ = v12
                }
            }
        }
    case 128:
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v13 float32 = 1.5
            _ = v13
        case 2147483647:
            p0 -= 42
        case 814:
            var v14 uint32 = p0
            _ = v14
        default:
            fmt.Println(p0, p0)
        }
    }
    _ = p0
    _ = p1
    return (p1 & p1)
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]int)
    var v3 int8 = 0
    v15 := []int{42, 32768, -2147483648}
    v16 := GenS1{F0: -1, F1: 0}
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(fi17)
        _ = fi17
    }
    var v18 uint = v1
    if len(v15) > 0 {
        v15[0] = 128
    }
    if (false && false) {
        for fi19 := 0; fi19 < 19; fi19++ {
            if true {
                if (0.0 < (float32(1e10) / 0.1)) {
                    var v20 int64 = -1
                    var v21 string = "hello"
                    if false {
                        var v22 float64 = (1.5 - 0.1)
                        for fi23 := 0; fi23 < 5; fi23++ {
                            v3 = v3
                            _ = fi23
                        }
                        v20 += v20
                        _ = v22
                    }
                    _ = v20
                    _ = v21
                } else {
                    if (v1 > 560) {
                        v24 := GenS2{F0: (^(fi19)), F1: 822, F2: 42}
                        _ = v24
                    } else {
                        v25 := GenS2{F0: (-(fi19)), F1: 4294967295, F2: 128}
                        v18 = v1
                        v2[2147483647] = (fi19 % ((fi19) | 1))
                        _ = v25
                    }
                }
            }
            if ((2 <= 32767) && true) {
                v1 += v1
            }
            _ = len(v15)
            _ = fi19
        }
    }
    if ((v18 * v18) != 1) {
        for fi26 := 0; fi26 < 14; fi26++ {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = fi26
        }
        for fi27 := 0; fi27 < 1; fi27++ {
            fmt.Println(v3, fi27, v3)
            _ = fi27
        }
    } else {
        fmt.Println(v3, v1, v18)
        var v28 uint8 = 134
        var v29 int64 = 42
        _ = v28
        _ = v29
    }
    var v30 uint64 = 42
    var v31 int32 = 42
    var v32 uint64 = (^(v30))
    v33 := GenS2{F0: 0, F1: 2147483648, F2: 1}
    v30 = v30
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
