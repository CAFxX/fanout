package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 string; F1 int64 }

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]bool)
    v3 := []uint32{209, 65535, 1, 37}
    v4 := make(map[uint32]uint32)
    var v5 int64 = 42
    var v6 int8 = 43
    var v7 int16 = 1
    delete(v2, 806)
    var v8 uint8 = 2
    var v9 bool = ((453 == 2) && false)
    if v9 {
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int16 = (v7 >> ((574) & 63))
            v5 += (v5 ^ v5)
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v1, v6, v6)
        _ = len(v3)
        v13 := GenS1{F0: 2}
        _ = fi12
        _ = v13
    }
    if ((v8 < v8) && (0 <= 4294967295)) {
        v7 -= 616
        if false {
            var v14 uint32 = 42
            _ = v14
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 int16 = (-(v7))
            _ = fi15
            _ = v16
        }
    }
    v1 = 255
    if (!(v9 || ((v7 >> ((v1) & 63)) >= v7))) {
        var v17 uint64 = 290
        if v9 {
            var v18 int32 = 1
            for fi19 := 0; fi19 < 18; fi19++ {
                if true {
                    var v20 int8 = 0
                    if v9 {
                        if (18446744073709551615 <= (v17 - v17)) {
                            v8 = v8
                        } else {
                            for fi21 := 0; fi21 < 5; fi21++ {
                                v20 = (-(v20))
                                if (("hello" + "\n\t") >= "hello") {
                                    switch "a" {
                                    case "a":
                                        if ((v18 - (v18 << ((712) & 63))) <= (-(v18))) {
                                            var v22 int64 = 527
                                            fmt.Println(v8, v18, fi19)
                                            _ = v22
                                        }
                                    default:
                                        v23 := GenS1{F0: 2}
                                        _ = v23
                                    }
                                    var v24 int16 = v7
                                    _ = v24
                                }
                                _ = fi21
                            }
                            fmt.Println(v20)
                        }
                    }
                    v9 = false
                    _ = v20
                }
                v5 += v5
                fmt.Println(v8, v7, v5)
                _ = fi19
            }
            _ = v18
        } else {
            v5 += (v5 / ((v5) | 1))
            if (2147483647 >= 187) {
                switch 787 {
                case 693:
                    if (2 > 32767) {
                        var v25 int64 = (v5 % (((v5 * (^(v5)))) | 1))
                        switch "foo bar" {
                        case "foo bar":
                            fmt.Println(v7, v8, v1)
                        }
                        _ = v25
                    } else {
                        switch -2147483648 {
                        case -2147483648:
                            v6 = (v6 << ((uint(v17)) & 63))
                        }
                        var v26 uint = (v1 << (((v1 << (((v1 / ((v1) | 1))) & 63))) & 63))
                        _ = v26
                    }
                case 32767:
                    for fi27 := 0; fi27 < 14; fi27++ {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            var v29 int = (fi28 * (fi27 ^ (fi28 << ((v1) & 63))))
                            var v30 float64 = 0.1
                            var v31 string = "\n\t"
                            switch 42 {
                            case 128:
                                if v9 {
                                    v32 := GenS2{F0: uint32(v8), F1: 163, F2: 2147483647}
                                    _ = v32
                                }
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        _ = fi27
                    }
                case 696:
                    v6 = v6
                }
                var v33 int = 2147483647
                if (v17 > (v17 * (v17 / (((v17 >> ((v1) & 63))) | 1)))) {
                    v8 -= uint8(v17)
                    _ = len(v4)
                } else {
                    fmt.Println(v33, v1)
                    switch "" {
                    case "hello":
                        v34 := GenS3{F0: true, F1: ("foo bar" + "hello")}
                        _ = v34
                    case "\n\t":
                        v17 = v17
                    case "":
                        if (v5 == v5) {
                            var v35 int = 104
                            v7 -= (v7 & -1)
                            _ = v35
                        } else {
                            for fi36 := 0; fi36 < 2; fi36++ {
                                var v37 int8 = -98
                                fmt.Println(v6)
                                if v9 {
                                    delete(v2, 1)
                                } else {
                                    v1 += (v1 - v1)
                                }
                                _ = fi36
                                _ = v37
                            }
                        }
                    }
                }
                _ = v33
            }
            if v9 {
                v9 = v9
            }
        }
        _ = v17
    } else {
        fmt.Println(v5, v9)
    }
    _ = len(v4)
    v7 -= -1
    v38 := GenS1{F0: v1}
    switch -9223372036854775808 {
    case 32768:
        var v39 int32 = 957
        _ = v39
    default:
        switch "a" {
        case "a":
            v40 := GenS4{F0: (3.14159 > (0.0 * float64(1e10))), F1: 4294967295}
            _ = v40
        case "":
            v8 += v8
        default:
            for fi41 := 0; fi41 < 14; fi41++ {
                v42 := GenS5{F0: "a", F1: v5}
                _ = fi41
                _ = v42
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
