package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint8; F1 uint64; F2 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 int16; F2 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 0
    v3 := make(map[int]string)
    var v4 uint = 791
    var v5 string = ""
    fmt.Println(v5)
    v6 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615, F2: 1}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = true
        if (!v8) {
            switch "\n\t" {
            case "a":
                v5 = (v5 + v5)
            case "":
                v5 = ((v5 + "") + v5)
            }
            var v9 uint16 = (^(v2))
            _ = v9
        } else {
            for fi10 := 0; fi10 < 6; fi10++ {
                switch fi10 {
                case 462:
                    v11 := []int{(fi7 << ((432) & 63))}
                    _ = v11
                case 128:
                    var v12 bool = (!true)
                    _ = v12
                case -1:
                    delete(v1, 65535)
                default:
                    var v13 uint32 = 2147483648
                    _ = v13
                }
                var v14 float64 = float64(0.1)
                v15 := []int{(-(fi10)), (-(fi10)), (fi10 << ((65535) & 63))}
                var v16 uint32 = 65535
                v17 := GenS2{F0: 42, F1: -1}
                _ = fi10
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            if (v2 >= v2) {
                v4 -= (v4 ^ v4)
            } else {
                v2 += 228
                switch fi18 {
                case 2147483647:
                    fmt.Println(v2, v8)
                case -2147483648:
                    switch "a" {
                    case "":
                        v5 = "\n\t"
                    }
                case 32768:
                    if v8 {
                        switch "" {
                        case "\n\t":
                            v19 := GenS3{F0: 144, F1: (v5 + "\n\t")}
                            var v20 float32 = ((1.5 * float32((float64((0.1 + (3.14159 - (float32((1e10 / float64(3.14159))) - 1.5)))) / 3.14159))) + 3.14159)
                            v21 := GenS3{F0: 42, F1: "\n\t"}
                            v22 := v1[0]
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                        case "foo bar":
                            v23 := []int{(fi7 >> (((v4 ^ 1)) & 63)), 1}
                            _ = v23
                        }
                        var v24 float32 = 1e10
                        _ = v24
                    }
                default:
                    var v25 int64 = 42
                    _ = v25
                }
                v8 = (-2.25 <= 3.14159)
            }
            _ = fi18
        }
        _ = fi7
        _ = v8
    }
    v2 += v2
    v2 = v2
    if (v5 < (v5 + v5)) {
        v2 -= 140
        fmt.Println(v5, v2, v5)
    }
    v26 := []int{2147483647, 32768, 42, 4294967295}
    for fi27 := 0; fi27 < 15; fi27++ {
        _ = len(v3)
        switch fi27 {
        case 32767:
            var v28 uint = v4
            _ = v28
        }
        switch v5 {
        case "":
            if (561 >= -614) {
                if true {
                    switch (fi27 * fi27) {
                    case 128:
                        v5 = (("" + v5) + "\n\t")
                    case 4294967295:
                        if true {
                            if false {
                                if (2 <= 2147483648) {
                                    v4 += 1
                                    for fi29 := 0; fi29 < 16; fi29++ {
                                        v4 = (v4 << (((v4 / (((v4 % (((v4 >> ((255) & 63))) | 1))) | 1))) & 63))
                                        _ = fi29
                                    }
                                }
                            }
                            var v30 uint8 = 42
                            _ = v30
                        } else {
                            v1[1] = (fi27 % ((249) | 1))
                            fmt.Println(v4, v4, v2)
                        }
                    case -854:
                        if (!(true && false)) {
                            v31 := GenS4{F0: 229, F1: 1, F2: true}
                            v32 := GenS5{F0: v5}
                            _ = v31
                            _ = v32
                        } else {
                            var v33 uint16 = 65535
                            v34 := GenS6{F0: (!(1 < 15))}
                            _ = v33
                            _ = v34
                        }
                    default:
                        fmt.Println(v2, v4, fi27)
                    }
                    if (!(!true)) {
                        var v35 uint32 = 956
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 7; fi36++ {
                        v2 += 42
                        var v37 uint32 = 2
                        var v38 bool = (!true)
                        v39 := GenS7{F0: fi27}
                        v4 += (v4 & 320)
                        _ = fi36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                } else {
                    switch "foo bar" {
                    case "a":
                        var v40 string = "\n\t"
                        _ = v40
                    default:
                        for fi41 := 0; fi41 < 10; fi41++ {
                            v42 := GenS8{F0: false, F1: 1, F2: fi27}
                            fmt.Println(v5)
                            _ = fi41
                            _ = v42
                        }
                    }
                }
                v43 := GenS6{F0: (0.1 == 0.1)}
                v5 = "\n\t"
                _ = v43
            } else {
                fmt.Println(v2, v2)
                var v44 uint32 = uint32(v2)
                v5 = ("hello" + v5)
                _ = v44
            }
        }
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
