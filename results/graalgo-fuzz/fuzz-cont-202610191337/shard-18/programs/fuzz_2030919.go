package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int16; F2 int64 }

func main() {
    var v1 string = "foo bar"
    var v2 float32 = ((float32((1.5 * 1.5)) - (0.1 / (1.5 - ((0.1 * (3.14159 + -2.25)) + 0.0)))) - (0.0 - float32(0.0)))
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    v2 = 0.1
    switch -394 {
    case 127:
        if true {
            var v5 int16 = 818
            fmt.Println(v2, v2)
            var v6 int32 = 0
            _ = v5
            _ = v6
        }
    case 32768:
        if false {
            for fi7 := 0; fi7 < 2; fi7++ {
                v8 := GenS1{F0: 148}
                if false {
                    switch v1 {
                    case "a":
                        if true {
                            if ((!true) && ((-696 == int16(fi7)) || (fi7 >= fi7))) {
                                if false {
                                    if (1 <= 114) {
                                        v2 -= (-2.25 * float32(float64(1e10)))
                                        if false {
                                            v1 = (v1 + ((((v1 + "") + v1) + v1) + v1))
                                        } else {
                                            v2 = ((v2 / v2) - v2)
                                            if false {
                                                var v9 int = (fi7 | fi7)
                                                _ = v9
                                            }
                                            v1 = v1
                                        }
                                    } else {
                                        switch (fi7 - fi7) {
                                        case -363:
                                            var v10 float32 = v2
                                            _ = v10
                                        default:
                                            switch (fi7 % ((fi7) | 1)) {
                                            case -9223372036854775808:
                                                v4[4294967295] = 0
                                            case -1:
                                                v11 := []int{(fi7 | fi7), (^(fi7)), fi7, (fi7 << ((255) & 63))}
                                                _ = v11
                                            default:
                                                var v12 uint16 = 2
                                                _ = v12
                                            }
                                        }
                                        v1 = v1
                                    }
                                }
                                fmt.Println(fi7, v1, fi7)
                            }
                        } else {
                            var v13 int = (fi7 << ((18446744073709551615) & 63))
                            _ = v13
                        }
                    default:
                        fmt.Println(fi7)
                    }
                    v14 := GenS2{F0: "hello", F1: 768, F2: -117}
                    _ = v14
                }
                _ = fi7
                _ = v8
            }
            var v15 float64 = (-2.25 * 1.5)
            var v16 string = "hello"
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := GenS3{F0: 2, F1: 0, F2: 2}
        v2 += (v2 * float32((-2.25 / float64(0.1))))
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 float32 = v2
            v21 := GenS4{F0: 459, F1: 205}
            fmt.Println(fi17, v20, fi19)
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi17
        _ = v18
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            v2 = 0.0
            v24 := []int{}
            _ = fi23
            _ = v24
        }
        v1 = "a"
        _ = fi22
    }
    var v25 int = -2147483648
    v26 := GenS5{F0: 42}
    v25 += 701
    var v27 uint16 = 1
    switch "a" {
    case "a":
        v4[2147483648] = 65535
    case "\n\t":
        v28 := GenS6{F0: 405, F1: -659, F2: 323}
        _ = v28
    default:
        v29 := v4[566]
        _ = v29
    }
    var v30 bool = ((42 > 127) || true)
    for fi31 := 0; fi31 < 7; fi31++ {
        if v30 {
            var v32 uint = 42
            v25 = (fi31 << (((v32 >> ((v32) & 63))) & 63))
            _ = v32
        }
        var v33 int32 = 1
        _ = fi31
        _ = v33
    }
    var v34 int64 = 983
    v25 -= v25
    if (501 > 42) {
        var v35 float64 = float64((v2 + v2))
        v3[-2147483648] = (v25 * (v25 - (v25 * v25)))
        if false {
            for fi36 := 0; fi36 < 9; fi36++ {
                var v37 uint64 = 0
                if (false || ((^(v27)) > v27)) {
                    v1 = (v1 + (("hello" + (v1 + "\n\t")) + (v1 + "\n\t")))
                    v38 := v3[(v25 + v25)]
                    _ = v38
                } else {
                    v37 = (v37 * 18446744073709551615)
                    var v39 string = v1
                    var v40 uint32 = 265
                    _ = v39
                    _ = v40
                }
                v1 = (v1 + "foo bar")
                _ = fi36
                _ = v37
            }
            fmt.Println(v1, v35, v35)
            v2 -= (float32(v35) + 0.1)
        }
        _ = v35
    }
    v41 := v3[(^(v25))]
    var v42 int16 = 1000
    v27 -= 928
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v41
    _ = v42
}
