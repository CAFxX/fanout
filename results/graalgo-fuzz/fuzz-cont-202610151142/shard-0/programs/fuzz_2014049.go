package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint }
type GenS3 struct { F0 bool; F1 bool; F2 int }
type GenS4 struct { F0 int64; F1 uint64 }
func genFn0(p0 bool) bool {
    v6 := []int{}
    for fi7 := 0; fi7 < 6; fi7++ {
        if p0 {
            p0 = (0.0 < 1e10)
            p0 = false
        }
        v8 := GenS1{F0: 0, F1: true, F2: 2}
        _ = fi7
        _ = v8
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    p0 = (p0 && false)
    var v9 string = "hello"
    _ = p0
    _ = v6
    _ = v9
    return (p0 || p0)
}

func main() {
    var v1 int32 = -2147483648
    var v2 string = ""
    var v3 float32 = 0.1
    v4 := make(map[int]int)
    var v5 uint16 = 901
    var v10 int32 = (v1 & (v1 / (((-(v1))) | 1)))
    v10 -= -2147483648
    fmt.Println(v1, v1, v3)
    if true {
        _ = len(v4)
        v11 := GenS1{F0: 2, F1: (1000 >= -32768), F2: 42}
        _ = v11
    }
    if (!((!(0 <= 4294967295)) && (true || true))) {
        v1 += (v1 >> ((255) & 63))
        for fi12 := 0; fi12 < 6; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                v10 -= v1
                v14 := []int{-2147483648, fi12}
                var v15 string = "hello"
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
    } else {
        var v16 int32 = v1
        for fi17 := 0; fi17 < 7; fi17++ {
            fmt.Println(fi17, v10, v5)
            switch (^(fi17)) {
            case 32767:
                v2 = "hello"
            default:
                delete(v4, fi17)
            }
            var v18 bool = (false || (false && (2 > 127)))
            _ = fi17
            _ = v18
        }
        v19 := GenS1{F0: 18446744073709551615, F1: true, F2: 2}
        _ = v16
        _ = v19
    }
    v20 := []int{-330, 824, 555}
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v21 := GenS1{F0: 1, F1: (2 > 0), F2: 2147483648}
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 uint = 508
        _ = fi22
        _ = v23
    }
    v3 = float32((-2.25 - 1.5))
    if false {
        v24 := v4[32768]
        var v25 int64 = -9223372036854775808
        if false {
            if (((false && false) || (!false)) && false) {
                switch v2 {
                case "":
                    for fi26 := 0; fi26 < 7; fi26++ {
                        switch (v24 * 0) {
                        case -76:
                            var v27 float64 = 0.0
                            _ = v27
                        default:
                            if len(v20) > 0 {
                                v20[0] = 4294967295
                            }
                        }
                        _ = fi26
                    }
                case "hello":
                    v10 -= v1
                }
                v20 = append(v20, v24)
                v28 := GenS2{F0: 2147483647, F1: 18446744073709551615, F2: 271}
                _ = v28
            }
            v29 := GenS3{F0: (!(((3.14159 - v3) - (v3 + v3)) != v3)), F1: (!true), F2: v24}
            _ = v29
        } else {
            _ = len(v4)
        }
        _ = v24
        _ = v25
    } else {
        v30 := GenS4{F0: 0, F1: 42}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
