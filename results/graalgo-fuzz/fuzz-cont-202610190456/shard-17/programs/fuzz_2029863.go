package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int16; F1 int16; F2 int64 }
func genFn0(p0 int, p1 int) bool {
    fmt.Println(p0)
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v6
    return ((p0 ^ p1) < p0)
}

func main() {
    var v1 uint32 = 65535
    var v2 uint = uint(v1)
    v3 := []int{160, 2147483647, 32768}
    v4 := make(map[uint32]uint32)
    v5 := []uint32{v1, (^(v1)), v1, (v1 % (((^(v1))) | 1))}
    v7 := GenS1{F0: 255, F1: (!true)}
    if (true && true) {
        v1 = (v1 | (^(v1)))
    } else {
        if (255 >= 2) {
            if (!false) {
                _ = len(v4)
                for fi8 := 0; fi8 < 6; fi8++ {
                    var v9 float64 = 1e10
                    switch "a" {
                    case "foo bar":
                        v9 = float64(1.5)
                    }
                    _ = fi8
                    _ = v9
                }
            }
            var v10 uint8 = 127
            _ = v10
        } else {
            var v11 uint64 = 4294967295
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint8 = 42
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        _ = len(v3)
        for fi15 := 0; fi15 < 3; fi15++ {
            if len(v3) > 0 {
                v3[0] = 32767
            }
            v1 = (v1 * (v1 / ((v1) | 1)))
            _ = fi15
        }
        v1 += (v1 / (((v1 >> ((v2) & 63))) | 1))
        _ = fi14
    }
    var v16 int32 = -1
    delete(v4, 916)
    fmt.Println(v1, v16, v2)
    _ = len(v3)
    switch -1 {
    case 0:
        switch -1 {
        case 127:
            v17 := GenS1{F0: 161, F1: true}
            _ = v17
        case -721:
            for fi18 := 0; fi18 < 6; fi18++ {
                fmt.Println(v16, v2, fi18)
                if ((0 == 745) && false) {
                    switch (^(fi18)) {
                    case 780:
                        if true {
                            var v19 float64 = (((float64(0.0) - ((0.0 - 0.1) - 0.1)) - 0.1) * float64(-2.25))
                            _ = v19
                        }
                    }
                    delete(v4, v1)
                }
                _ = fi18
            }
        case 32767:
            var v20 bool = (!false)
            _ = v20
        default:
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 int32 = v16
                var v23 int = fi21
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    case 127:
        switch 32767 {
        case 0:
            v5 = append(v5, (v1 * (v1 << (((v2 ^ v2)) & 63))))
        }
    default:
        var v24 int64 = -426
        _ = v24
    }
    var v25 bool = (!true)
    for fi26 := 0; fi26 < 3; fi26++ {
        _ = len(v3)
        switch "foo bar" {
        case "\n\t":
            if (0 > 42) {
                fmt.Println(v16, v25)
                v27 := GenS1{F0: 125, F1: (((-2.25 + 0.1) / 0.1) < (float64(1e10) / 0.1))}
                _ = v27
            } else {
                for fi28 := 0; fi28 < 1; fi28++ {
                    v29 := GenS1{F0: 42, F1: v25}
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        v30 := GenS1{F0: 0, F1: v25}
                        _ = v30
                    case "\n\t":
                        v2 = v2
                    case "":
                        var v31 bool = v25
                        _ = v31
                    }
                    _ = fi28
                    _ = v29
                }
            }
        case "":
            if v25 {
                fmt.Println(v16, v25, fi26)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                for fi32 := 0; fi32 < 19; fi32++ {
                    if (!(v25 || false)) {
                        switch "foo bar" {
                        case "":
                            v33 := GenS2{F0: 1, F1: -502, F2: 0}
                            _ = v33
                        case "a":
                            v16 += -535
                        default:
                            switch "hello" {
                            case "":
                                v34 := v4[2147483647]
                                var v35 uint8 = 230
                                if len(v3) > 0 {
                                    v3[0] = (fi26 % ((fi26) | 1))
                                }
                                fmt.Println(v1, fi32)
                                v3 = append(v3, fi32)
                                _ = v34
                                _ = v35
                            case "a":
                                if v25 {
                                    v36 := GenS1{F0: 25, F1: false}
                                    v1 = 42
                                    var v37 int8 = 100
                                    _ = v36
                                    _ = v37
                                }
                            default:
                                var v38 uint32 = v1
                                _ = v38
                            }
                        }
                        v1 = (^(v1))
                    } else {
                        switch (("" + "") + "foo bar") {
                        case "hello":
                            v2 = (v2 | v2)
                        case "\n\t":
                            v16 = v16
                        }
                    }
                    var v39 uint32 = (v1 & (v1 >> (((v2 / ((18446744073709551615) | 1))) & 63)))
                    v1 = (^(v1))
                    _ = fi32
                    _ = v39
                }
            }
        }
        var v40 uint16 = 537
        _ = fi26
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
