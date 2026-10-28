package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 uint64; F2 uint16 }
func genFn0(p0 int32, p1 int32, p2 int) uint8 {
    fmt.Println(p2, p1)
    v4 := GenS1{F0: 145}
    p1 += -707
    v5 := GenS1{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 84
}
func genFn1(p0 uint, p1 int64, p2 string) uint {
    var v6 uint64 = 0
    switch p2 {
    case "hello":
        if true {
            var v7 int64 = 503
            var v8 float32 = 1e10
            switch 0 {
            case 4294967295:
                var v9 float64 = 0.0
                _ = v9
            case -624:
                var v10 uint32 = 335
                _ = v10
            default:
                if true {
                    fmt.Println(v6, v6)
                    v11 := GenS1{F0: 2}
                    _ = v11
                } else {
                    v12 := GenS1{F0: genFn0(-2147483648, 42, 32768)}
                    v13 := GenS1{F0: genFn0(-1, -139, 2147483647)}
                    p0 = (p0 % ((2) | 1))
                    _ = v12
                    _ = v13
                }
            }
            _ = v7
            _ = v8
        } else {
            var v14 string = (p2 + "a")
            p2 = v14
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint = p0
        switch 32767 {
        case 32767:
            p2 = p2
        }
        v17 := []int{fi15, 32767}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        if false {
            v6 += (^(v6))
        } else {
            v19 := []int{fi18, fi18}
            var v20 bool = true
            _ = v19
            _ = v20
        }
        v21 := []int{fi18, (fi18 << ((191) & 63)), (fi18 << ((p0) & 63))}
        v6 += v6
        _ = fi18
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 + (^(p0)))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := []uint32{853, 784, 4294967295, 2}
    var v22 int8 = 100
    var v23 int8 = v22
    fmt.Println(v22, v22)
    var v24 int8 = v22
    switch 9223372036854775807 {
    case 0:
        fmt.Println(v24, v22)
    case 2147483647:
        if false {
            var v25 uint8 = 2
            v23 += (v24 * (v23 / ((-98) | 1)))
            _ = v25
        }
    default:
        v22 -= (v23 & (^(v22)))
    }
    v24 = v23
    v26 := GenS1{F0: 54}
    if (((false && false) || (65535 != 65535)) && true) {
        if (true && true) {
            for fi27 := 0; fi27 < 2; fi27++ {
                v28 := GenS1{F0: 252}
                if len(v3) > 0 {
                    v3[0] = 42
                }
                _ = fi27
                _ = v28
            }
            switch 9223372036854775807 {
            case 1:
                switch 42 {
                case 9223372036854775807:
                    v29 := GenS2{F0: "hello", F1: 31, F2: 255}
                    _ = v29
                case 736:
                    if len(v3) > 0 {
                        v3[0] = 2147483648
                    }
                default:
                    _ = len(v2)
                }
            case 128:
                var v30 int16 = 42
                _ = v30
            case -1:
                fmt.Println(v23)
            }
            v31 := GenS1{F0: genFn0(0, 2147483647, -9223372036854775808)}
            _ = v31
        }
    }
    var v32 bool = (!true)
    v24 -= v23
    v32 = (42 <= 752)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v32
}
