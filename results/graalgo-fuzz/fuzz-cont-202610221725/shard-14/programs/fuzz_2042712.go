package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int32; F2 int64 }
type GenS6 struct { F0 int16; F1 uint16; F2 int8 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 bool = false
    var v3 uint32 = 2147483648
    v4 := make(map[int]uint32)
    var v5 int16 = 1000
    var v6 int = -354
    if v2 {
        switch "\n\t" {
        case "\n\t":
            v7 := GenS1{F0: (^(v6)), F1: 124}
            _ = v7
        case "foo bar":
            if ("hello" == "a") {
                var v8 uint64 = 29
                var v9 float32 = (1.5 / 1.0)
                _ = v8
                _ = v9
            } else {
                if false {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        v11 := GenS1{F0: fi10, F1: 0}
                        if v2 {
                            v1 -= v1
                            v1 = (v1 + 0)
                        } else {
                            for fi12 := 0; fi12 < 15; fi12++ {
                                v13 := GenS1{F0: (fi12 | fi10), F1: 209}
                                fmt.Println(fi10, v3)
                                v4[(^(fi10))] = v3
                                _ = fi12
                                _ = v13
                            }
                            switch 0 {
                            case -2147483648:
                                var v14 int8 = -128
                                _ = v14
                            }
                        }
                        _ = fi10
                        _ = v11
                    }
                    for fi15 := 0; fi15 < 8; fi15++ {
                        v16 := GenS1{F0: 0, F1: 128}
                        var v17 uint8 = 226
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    var v18 float32 = 3.14159
                    _ = v18
                }
            }
        default:
            switch -2147483648 {
            case -687:
                var v19 uint8 = 42
                _ = v19
            case -249:
                v1 -= (^(v1))
            default:
                v1 = v1
            }
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int64 = 1
            var v22 int8 = 10
            var v23 float64 = (1e10 / 0.1)
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        switch ("foo bar" + (("foo bar" + ("foo bar" + "foo bar")) + "hello")) {
        case "\n\t":
            v24 := []int{v6, (v6 * (v6 | 813)), v6}
            _ = v24
        case "foo bar":
            for fi25 := 0; fi25 < 2; fi25++ {
                if (int32(v3) == -1) {
                    v5 -= -96
                    switch "a" {
                    case "a":
                        delete(v4, v6)
                    default:
                        if ((0.1 + 1e10) < float64(-2.25)) {
                            switch v6 {
                            case 32768:
                                v6 -= fi25
                            default:
                                fmt.Println(v3)
                            }
                            for fi26 := 0; fi26 < 18; fi26++ {
                                if (((^(v5)) == v5) && v2) {
                                    v4[32768] = v3
                                } else {
                                    var v27 bool = false
                                    v3 -= 789
                                    _ = v27
                                }
                                _ = fi26
                            }
                        } else {
                            for fi28 := 0; fi28 < 11; fi28++ {
                                switch "hello" {
                                case "hello":
                                    v6 = fi25
                                case "\n\t":
                                    var v29 int16 = v5
                                    _ = v29
                                case "":
                                    v30 := GenS2{F0: 42, F1: ("foo bar" + ("foo bar" + ("a" + "")))}
                                    _ = v30
                                }
                                var v31 float32 = 1e10
                                v32 := GenS3{F0: v5, F1: -476}
                                var v33 uint16 = 242
                                v31 -= ((-2.25 * 0.0) + v31)
                                v34 := GenS4{F0: int32(v5), F1: 39}
                                v1 -= v1
                                _ = fi28
                                _ = v31
                                _ = v32
                                _ = v33
                                _ = v34
                            }
                            var v35 float64 = (float64(0.0) / 1.5)
                            for fi36 := 0; fi36 < 1; fi36++ {
                                v3 = 4294967295
                                if (true && true) {
                                    var v37 bool = v2
                                    switch (^(fi25)) {
                                    case 32767:
                                        v6 += 32768
                                    case -539:
                                        v6 = fi36
                                    case 880:
                                        if ((-48 == 42) && v37) {
                                            var v38 int32 = int32(v1)
                                            v39 := []int{}
                                            var v40 float32 = (1.5 * 3.14159)
                                            if (v37 && true) {
                                                var v41 int64 = -761
                                                v1 = (v1 * (v1 | v1))
                                                _ = v41
                                            } else {
                                                v4[(fi25 - v6)] = v3
                                                if len(v39) >= 2 {
                                                    v39 = v39[1:len(v39)]
                                                }
                                            }
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                    }
                                    _ = v37
                                }
                                var v42 string = "foo bar"
                                _ = fi36
                                _ = v42
                            }
                            _ = v35
                        }
                    }
                } else {
                    if v2 {
                        var v43 uint32 = 476
                        if v2 {
                            fmt.Println(v43)
                        }
                        _ = v43
                    } else {
                        if (!(v5 != v5)) {
                            v2 = v2
                        }
                        var v44 float64 = (-2.25 - ((1e10 / 1.5) - 0.0))
                        fmt.Println(fi25, v2, v44)
                        _ = v44
                    }
                    v45 := []int{42, v6, (v6 % ((-9223372036854775808) | 1))}
                    _ = v45
                }
                fmt.Println(v3, v1)
                if (v5 < (v5 << (((v1 % ((v1) | 1))) & 63))) {
                    v46 := GenS5{F0: 127, F1: -1, F2: 496}
                    var v47 uint16 = 2
                    _ = v46
                    _ = v47
                } else {
                    if v2 {
                        v48 := GenS2{F0: 42, F1: "foo bar"}
                        var v49 int64 = 1
                        _ = v48
                        _ = v49
                    }
                }
                _ = fi25
            }
        }
    } else {
        var v50 int8 = -1
        v51 := []int{v6, (v6 * -565)}
        v6 -= 32768
        _ = v50
        _ = v51
    }
    v3 = (v3 | (v3 % ((65535) | 1)))
    switch "\n\t" {
    case "\n\t":
        var v52 int8 = 80
        _ = v52
    default:
        v53 := GenS2{F0: 42, F1: (("\n\t" + ("\n\t" + "foo bar")) + ("\n\t" + "hello"))}
        _ = v53
    }
    fmt.Println(v3, v2)
    v54 := GenS6{F0: 1000, F1: 255, F2: 1}
    var v55 bool = v2
    v5 -= (v5 - 310)
    fmt.Println(v5, v3)
    if (!(v2 || (!(!("hello" != ("hello" + "a")))))) {
        fmt.Println(v5, v5)
        fmt.Println(v55, v3, v5)
        v56 := GenS2{F0: 4294967295, F1: "hello"}
        _ = v56
    }
    v57 := []int{}
    v58 := GenS1{F0: (v6 / ((v6) | 1)), F1: 114}
    _ = len(v4)
    v57 = append(v57, (^(v6)))
    _ = len(v4)
    if v2 {
        for fi59 := 0; fi59 < 11; fi59++ {
            fmt.Println(v3, fi59)
            switch (v6 % ((128) | 1)) {
            case 127:
                delete(v4, v6)
            default:
                v55 = true
            }
            _ = fi59
        }
        var v60 int16 = (v5 % (((v5 >> (((v1 << ((255) & 63))) & 63))) | 1))
        for fi61 := 0; fi61 < 20; fi61++ {
            for fi62 := 0; fi62 < 20; fi62++ {
                var v63 float32 = float32(3.14159)
                _ = fi62
                _ = v63
            }
            var v64 int32 = 42
            _ = fi61
            _ = v64
        }
        _ = v60
    } else {
        var v65 float64 = (-2.25 * 1.5)
        _ = v65
    }
    for fi66 := 0; fi66 < 7; fi66++ {
        v2 = (606 >= 0)
        var v67 uint64 = 408
        _ = len(v57)
        _ = fi66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v57
    _ = v58
    _ = v6
}
