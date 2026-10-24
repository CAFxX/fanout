package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 uint; F1 uint8; F2 uint16 }
func genFn0(p0 uint64) uint16 {
    if true {
        v7 := []int{-69, 0, -1, 127}
        v8 := GenS1{F0: p0, F1: 42}
        _ = v7
        _ = v8
    }
    p0 = (p0 | (^(p0)))
    _ = p0
    return 2
}
func genFn1(p0 bool, p1 uint8, p2 int8) uint32 {
    if p0 {
        p2 = -84
        p0 = p0
    }
    v9 := GenS2{F0: 2147483647}
    switch 1 {
    case 624:
        if (!(9223372036854775807 != 128)) {
            if true {
                switch "hello" {
                case "\n\t":
                    if (!p0) {
                        if ("foo bar" <= "hello") {
                            switch "foo bar" {
                            case "":
                                p0 = (1 < 0)
                            case "foo bar":
                                var v10 string = "hello"
                                _ = v10
                            case "hello":
                                v11 := GenS1{F0: 81, F1: 475}
                                _ = v11
                            }
                            if p0 {
                                v12 := []int{2147483647, 1, 128, 32768}
                                _ = v12
                            } else {
                                fmt.Println(p0)
                                var v13 float32 = 0.0
                                _ = v13
                            }
                        } else {
                            if (p0 || false) {
                                v14 := []int{-1, int(p2)}
                                v15 := GenS1{F0: 499, F1: 2}
                                _ = v14
                                _ = v15
                            } else {
                                if p0 {
                                    if true {
                                        for fi16 := 0; fi16 < 7; fi16++ {
                                            var v17 bool = (278 == 4294967295)
                                            _ = fi16
                                            _ = v17
                                        }
                                    }
                                    var v18 bool = p0
                                    var v19 uint32 = 826
                                    _ = v18
                                    _ = v19
                                }
                                var v20 int16 = 1
                                if (p0 || ((!true) && false)) {
                                    if true {
                                        var v21 uint64 = 619
                                        _ = v21
                                    } else {
                                        if true {
                                            v20 = -919
                                            v22 := []int{32767}
                                            _ = v22
                                        }
                                        var v23 uint = 1
                                        v24 := []int{-565, 1, -881, 128}
                                        _ = v23
                                        _ = v24
                                    }
                                    var v25 uint32 = 1
                                    p0 = true
                                    _ = v25
                                } else {
                                    v20 += v20
                                }
                                _ = v20
                            }
                            for fi26 := 0; fi26 < 6; fi26++ {
                                v27 := GenS3{F0: 2147483648, F1: 0}
                                switch "\n\t" {
                                case "\n\t":
                                    var v28 uint8 = 127
                                    _ = v28
                                case "a":
                                    fmt.Println(p2, p0)
                                default:
                                    for fi29 := 0; fi29 < 20; fi29++ {
                                        if (-1 != int64(fi26)) {
                                            v30 := []int{(fi29 - fi26)}
                                            v31 := GenS1{F0: 0, F1: 458}
                                            v32 := GenS4{F0: p1, F1: 32767, F2: 2147483647}
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = fi29
                                    }
                                }
                                _ = fi26
                                _ = v27
                            }
                            if p0 {
                                var v33 bool = false
                                _ = v33
                            }
                        }
                        fmt.Println(p1)
                        fmt.Println(p2, p0, p1)
                    }
                case "":
                    fmt.Println(p0)
                default:
                    var v34 int16 = 1
                    _ = v34
                }
                p2 = (^(p2))
                fmt.Println(p2, p1)
            } else {
                switch -1 {
                case -247:
                    p1 += p1
                case -9223372036854775808:
                    fmt.Println(p1, p2, p0)
                case -621:
                    var v35 uint64 = 2
                    _ = v35
                default:
                    var v36 int64 = 568
                    _ = v36
                }
            }
        }
    case 4294967295:
        var v37 int = 128
        _ = v37
    default:
        p1 -= 226
    }
    fmt.Println(p0, p1)
    v38 := GenS5{F0: 127, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    _ = v9
    return 2147483648
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float32 = float32(1.5)
    var v3 int16 = 42
    var v4 float64 = 0.0
    v5 := []string{("" + "\n\t"), "a", ""}
    var v6 bool = false
    var v39 float64 = 0.1
    v2 = v2
    fmt.Println(v6, v2, v4)
    var v40 float32 = -2.25
    fmt.Println(v3, v1)
    v41 := GenS4{F0: 2, F1: (v3 ^ (^(v3))), F2: 2}
    v42 := GenS6{F0: 18446744073709551615, F1: 55, F2: 1}
    v2 += -2.25
    for fi43 := 0; fi43 < 13; fi43++ {
        fmt.Println(v40)
        _ = fi43
    }
    if (2147483648 != 2147483648) {
        var v44 uint8 = 1
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
