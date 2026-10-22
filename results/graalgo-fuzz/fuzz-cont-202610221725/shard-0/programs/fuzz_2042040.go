package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint }
func genFn0(p0 int16, p1 uint) uint16 {
    switch ("\n\t" + ("a" + "\n\t")) {
    case "hello":
        v4 := GenS1{F0: 2}
        _ = v4
    case "foo bar":
        v5 := GenS1{F0: 127}
        _ = v5
    default:
        p1 = p1
    }
    switch "foo bar" {
    case "hello":
        var v6 uint32 = 766
        _ = v6
    case "":
        fmt.Println(p1)
    case "a":
        if (!(p0 >= p0)) {
            switch 32768 {
            case 32:
                p0 = p0
            case -9223372036854775808:
                switch "\n\t" {
                case "hello":
                    if ((p1 >> ((p1) & 63)) == (p1 * 388)) {
                        p0 += p0
                    }
                default:
                    v7 := []int{32767, 128}
                    _ = v7
                }
            case 42:
                for fi8 := 0; fi8 < 5; fi8++ {
                    switch (fi8 - fi8) {
                    case -1:
                        var v9 int8 = 42
                        _ = v9
                    case -274:
                        v10 := GenS1{F0: 127}
                        _ = v10
                    default:
                        switch fi8 {
                        case -597:
                            v11 := []int{fi8, -753}
                            _ = v11
                        case 858:
                            p0 += 42
                        default:
                            p1 = (p1 - 4294967295)
                        }
                    }
                    _ = fi8
                }
            default:
                p1 = 2
            }
        }
    }
    v12 := []int{2147483647, -484, 0}
    var v13 float32 = 1.5
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{127, 32767, -759}
    var v3 uint8 = 97
    var v14 uint64 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3)
    var v15 uint = 622
    v15 -= v15
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        v3 += v3
        _ = fi16
    }
    var v17 int32 = 42
    v17 += -958
    if (false && true) {
        fmt.Println(v17, v14, v14)
        v15 = (v15 + 0)
    }
    _ = len(v1)
    v17 = (v17 ^ (v17 - v17))
    v18 := GenS1{F0: (v3 % ((v3) | 1))}
    if false {
        v19 := GenS2{F0: 18446744073709551615, F1: v17}
        _ = len(v1)
        var v20 uint8 = (v3 | (v3 ^ 194))
        _ = v19
        _ = v20
    } else {
        switch 589 {
        case 4294967295:
            var v21 uint32 = 2147483648
            _ = v21
        case 128:
            var v22 int32 = v17
            _ = v22
        case -508:
            if ((v17 + v17) <= (v17 & (v17 & (v17 / (((-(v17))) | 1))))) {
                var v23 bool = true
                var v24 uint32 = 491
                v2 = append(v2, 128)
                _ = v23
                _ = v24
            } else {
                v1[9223372036854775807] = 0
                var v25 uint64 = (v14 * 37)
                var v26 int32 = (v17 % (((v17 << ((1) & 63))) | 1))
                _ = v25
                _ = v26
            }
        default:
            var v27 uint8 = (v3 + 42)
            _ = v27
        }
        if false {
            v3 = (v3 * v3)
        } else {
            delete(v1, 1)
            v28 := GenS3{F0: 4294967295}
            _ = v28
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            fmt.Println(v14)
            _ = fi29
        }
    }
    v3 = (v3 + v3)
    v3 += 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
