package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint16, p1 uint16) int {
    if ((0.1 / ((1.5 + (-2.25 - -2.25)) - float32(float64(0.0)))) != 1.5) {
        if false {
            if (3.14159 != 3.14159) {
                fmt.Println(p0, p0)
            } else {
                var v6 uint16 = p0
                for fi7 := 0; fi7 < 4; fi7++ {
                    switch (fi7 ^ fi7) {
                    case 32767:
                        var v8 int32 = 42
                        _ = v8
                    case 2147483647:
                        v9 := GenS1{F0: (fi7 ^ (fi7 + (fi7 & 9223372036854775807))), F1: 726, F2: 125}
                        _ = v9
                    }
                    if false {
                        if ((1e10 / 1e10) <= 0.1) {
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                var v10 int8 = -125
                                _ = v10
                            default:
                                var v11 uint32 = 557
                                _ = v11
                            }
                            v6 = 316
                            for fi12 := 0; fi12 < 14; fi12++ {
                                var v13 uint = 1
                                p1 -= (p1 | p0)
                                _ = fi12
                                _ = v13
                            }
                        }
                    }
                    var v14 uint8 = 1
                    _ = fi7
                    _ = v14
                }
                _ = v6
            }
            fmt.Println(p1)
        } else {
            if true {
                p1 -= (p1 % ((p1) | 1))
            }
            v15 := []int{-78, 224}
            switch -1 {
            case -879:
                var v16 uint64 = 42
                _ = v16
            case 4294967295:
                p0 -= (^(p1))
            case 32767:
                for fi17 := 0; fi17 < 3; fi17++ {
                    p1 = 943
                    p1 = 366
                    v18 := GenS2{F0: 839, F1: 42, F2: 0}
                    _ = fi17
                    _ = v18
                }
            }
            _ = v15
        }
        if (873 > 1) {
            v19 := []int{int(p1), -1}
            _ = v19
        }
        p0 -= p1
    } else {
        v20 := []int{-978, 2147483647, -357, -1}
        if true {
            var v21 uint64 = 18446744073709551615
            switch 0 {
            case 4294967295:
                v21 = v21
            case 0:
                p0 += p1
            }
            _ = v21
        }
        _ = v20
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := []int{(fi22 * fi22)}
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 int32, p1 int8, p2 uint) uint32 {
    var v24 uint8 = 42
    var v25 uint32 = 65535
    v26 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v26
    return uint32(v24)
}

func main() {
    var v1 uint32 = 917
    var v2 float64 = (0.0 + ((0.0 * 0.1) / 1.0))
    var v3 int16 = -69
    var v4 bool = (true && (522 > 0))
    v5 := make(map[uint32]bool)
    v1 = v1
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 string = "\n\t"
        _ = fi27
        _ = v28
    }
    var v29 bool = (128 < 0)
    delete(v5, (v1 & v1))
    if (0.1 != 1e10) {
        v30 := GenS3{F0: 4294967295}
        _ = v30
    }
    v1 -= 65535
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 uint8 = 42
        switch "a" {
        case "hello":
            var v33 uint64 = 0
            _ = v33
        case "foo bar":
            v2 = 0.0
        case "\n\t":
            v29 = ("" < "hello")
        default:
            v5[(v1 % ((v1) | 1))] = v29
        }
        _ = fi31
        _ = v32
    }
    var v34 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
