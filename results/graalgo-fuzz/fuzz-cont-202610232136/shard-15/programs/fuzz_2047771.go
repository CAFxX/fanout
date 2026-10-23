package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int64; F1 int8; F2 bool }
type GenS4 struct { F0 uint }
func genFn0() uint {
    var v5 uint8 = 1
    v6 := []int{}
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 int16, p1 int64) uint8 {
    switch 1 {
    case 32767:
        var v7 int = 32767
        _ = v7
    case 1:
        fmt.Println(p0, p1)
    default:
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := []int{}
            _ = fi8
            _ = v9
        }
    }
    switch 2147483647 {
    case 4294967295:
        for fi10 := 0; fi10 < 2; fi10++ {
            fmt.Println(fi10)
            switch "hello" {
            case "foo bar":
                p0 = -732
            case "\n\t":
                switch 541 {
                case -671:
                    v11 := GenS1{F0: p1}
                    _ = v11
                default:
                    var v12 int32 = 126
                    _ = v12
                }
            case "a":
                if (127 == -128) {
                    var v13 int8 = 127
                    var v14 int64 = -889
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(p1, p0, p0)
                    switch (^(fi10)) {
                    case 32768:
                        switch "a" {
                        case "a":
                            switch ("\n\t" + "hello") {
                            case "":
                                fmt.Println(fi10, p1, fi10)
                            }
                        default:
                            if true {
                                var v15 uint = 4294967295
                                p1 = p1
                                _ = v15
                            } else {
                                switch ("" + "foo bar") {
                                case "":
                                    switch fi10 {
                                    case 9223372036854775807:
                                        for fi16 := 0; fi16 < 13; fi16++ {
                                            if (!(224 == 42)) {
                                                p1 -= (-(p1))
                                                var v17 bool = (true && true)
                                                var v18 bool = (2 <= 780)
                                                var v19 int8 = 0
                                                v19 = (v19 + (v19 * (v19 << ((475) & 63))))
                                                _ = v17
                                                _ = v18
                                                _ = v19
                                            } else {
                                                p1 = -1
                                                v20 := []int{(fi10 / ((fi10) | 1)), fi10, 0, (fi10 >> ((160) & 63))}
                                                _ = v20
                                            }
                                            var v21 int16 = p0
                                            var v22 uint16 = 2
                                            _ = fi16
                                            _ = v21
                                            _ = v22
                                        }
                                    case 1:
                                        switch (fi10 ^ fi10) {
                                        case 1:
                                            fmt.Println(p0, p1, fi10)
                                        case -9223372036854775808:
                                            p1 = 445
                                        case 32768:
                                            v23 := GenS2{F0: 18446744073709551615, F1: 797}
                                            _ = v23
                                        default:
                                            p1 += (p1 + -1)
                                        }
                                    case -546:
                                        fmt.Println(fi10)
                                    }
                                case "a":
                                    var v24 uint16 = 255
                                    _ = v24
                                case "\n\t":
                                    for fi25 := 0; fi25 < 11; fi25++ {
                                        v26 := GenS3{F0: 42, F1: 0, F2: (!(869 >= p0))}
                                        _ = fi25
                                        _ = v26
                                    }
                                }
                            }
                        }
                    case 128:
                        for fi27 := 0; fi27 < 6; fi27++ {
                            fmt.Println(p0)
                            v28 := []int{(fi27 << ((42) & 63)), fi10, fi27}
                            _ = fi27
                            _ = v28
                        }
                    default:
                        var v29 uint8 = 42
                        _ = v29
                    }
                    v30 := []int{(fi10 >> ((255) & 63))}
                    _ = v30
                }
            }
            _ = fi10
        }
    case -828:
        for fi31 := 0; fi31 < 12; fi31++ {
            p0 = (p0 / ((1000) | 1))
            var v32 int64 = 46
            _ = fi31
            _ = v32
        }
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v33 := GenS1{F0: 9223372036854775807}
        _ = v33
    case "":
        v34 := GenS4{F0: genFn0()}
        _ = v34
    case "hello":
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := []int{}
    v2 := make(map[int]string)
    var v3 uint64 = 4294967295
    v4 := make(map[int]int)
    fmt.Println(v3)
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v35 uint16 = 1
    for fi36 := 0; fi36 < 4; fi36++ {
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 float32 = (((-2.25 * float32((1.5 / (float64(0.0) + -2.25)))) * (float32(0.0) * ((3.14159 / (((((0.0 / 1.0) + 1.5) / (1e10 * 0.1)) / -2.25) + -2.25)) - float32(float64(0.0))))) + (float32(float64(0.1)) - 0.0))
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    v35 = v35
    switch -176 {
    case -1:
        switch 1 {
        case 9223372036854775807:
            fmt.Println(v3, v35)
        }
    }
    v39 := GenS2{F0: 909, F1: 743}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
}
