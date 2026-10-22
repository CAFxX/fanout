package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 bool; F1 int32; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 uint; F2 int8 }
func genFn0(p0 int16, p1 int16, p2 int) int8 {
    fmt.Println(p0, p1)
    p2 = (p2 | p2)
    if ("" == (("foo bar" + "foo bar") + "")) {
        var v6 int64 = 42
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int16) uint16 {
    v7 := []int{42}
    p0 -= (p0 << ((65535) & 63))
    _ = p0
    _ = v7
    return 255
}

func main() {
    var v1 float64 = 0.1
    var v2 uint8 = 1
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    var v5 uint64 = 42
    switch ("a" + "a") {
    case "hello":
        v5 -= v5
    case "":
        fmt.Println(v1, v2, v5)
    }
    if false {
        v8 := []int{128}
        var v9 int = 1
        v9 -= v9
        _ = v8
        _ = v9
    }
    v4[1] = uint32(v2)
    v10 := GenS1{F0: v5, F1: (v2 / ((v2) | 1))}
    v2 = v2
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int8 = 100
        v2 += 0
        for fi13 := 0; fi13 < 11; fi13++ {
            switch fi11 {
            case 32767:
                v12 += 100
            default:
                fmt.Println(fi11, v1, v2)
            }
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    var v14 int = 2147483647
    _ = len(v3)
    var v15 float64 = 3.14159
    v16 := GenS2{F0: (true && false), F1: -2147483648, F2: 424}
    var v17 uint16 = 1
    fmt.Println(v14)
    switch (v14 % (((-(v14))) | 1)) {
    case -469:
        v18 := []int{v14}
        _ = v18
    default:
        if false {
            v19 := GenS3{F0: 3}
            for fi20 := 0; fi20 < 3; fi20++ {
                v2 = 88
                var v21 uint = 570
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                v14 = (v14 - -945)
                var v23 uint16 = v17
                if false {
                    v15 += v1
                } else {
                    if true {
                        var v24 string = "hello"
                        var v25 uint16 = (v17 << ((1) & 63))
                        _ = v24
                        _ = v25
                    }
                    var v26 uint32 = 491
                    _ = v26
                }
                _ = fi22
                _ = v23
            }
            _ = v19
        } else {
            v27 := GenS2{F0: (false && (255 == v2)), F1: 305, F2: v14}
            if (v2 > v2) {
                fmt.Println(v14, v17)
                switch (v14 << ((513) & 63)) {
                case 9223372036854775807:
                    switch "\n\t" {
                    case "a":
                        var v28 float32 = (1e10 + 1.5)
                        _ = v28
                    case "hello":
                        for fi29 := 0; fi29 < 6; fi29++ {
                            var v30 int8 = 127
                            for fi31 := 0; fi31 < 10; fi31++ {
                                switch fi29 {
                                case -175:
                                    v30 += genFn0(-510, 32767, fi31)
                                case -467:
                                    var v32 int16 = 1
                                    v3[898] = 65535
                                    var v33 bool = (false || (2 != 0))
                                    _ = v32
                                    _ = v33
                                case 1:
                                    v34 := GenS4{F0: 809, F1: 2, F2: v30}
                                    _ = v34
                                }
                                var v35 int32 = 42
                                _ = fi31
                                _ = v35
                            }
                            _ = fi29
                            _ = v30
                        }
                    }
                case 32767:
                    v36 := []int{v14}
                    _ = v36
                case -539:
                    if true {
                        var v37 uint16 = 2
                        v38 := []int{127, (v14 * (-(v14))), v14, (v14 / ((v14) | 1))}
                        _ = v37
                        _ = v38
                    }
                }
            } else {
                v17 = v17
                var v39 int64 = 356
                _ = v39
            }
            _ = v27
        }
    }
    var v40 uint = 18446744073709551615
    v17 = (v17 % ((v17) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
