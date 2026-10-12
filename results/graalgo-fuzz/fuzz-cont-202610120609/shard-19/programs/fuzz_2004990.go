package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 string; F1 int; F2 int64 }
type GenS4 struct { F0 uint; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0(p0 string) int16 {
    var v5 int16 = 1
    for fi6 := 0; fi6 < 10; fi6++ {
        if ("foo bar" == p0) {
            p0 = ("a" + (("" + p0) + ("a" + "")))
        } else {
            for fi7 := 0; fi7 < 11; fi7++ {
                switch p0 {
                case "\n\t":
                    var v8 int16 = (v5 % ((v5) | 1))
                    _ = v8
                case "":
                    fmt.Println(fi7, v5)
                }
                _ = fi7
            }
        }
        v9 := []int{}
        for fi10 := 0; fi10 < 14; fi10++ {
            if true {
                if true {
                    p0 = (p0 + "hello")
                    var v11 int16 = (v5 >> ((255) & 63))
                    if true {
                        var v12 int64 = 1
                        var v13 uint64 = 2
                        _ = v12
                        _ = v13
                    } else {
                        var v14 int64 = 0
                        if (!(false || false)) {
                            var v15 float32 = 1.5
                            _ = v15
                        }
                        _ = v14
                    }
                    _ = v11
                } else {
                    p0 = p0
                    switch fi6 {
                    case 127:
                        var v16 uint32 = 140
                        _ = v16
                    case -698:
                        _ = len(v9)
                    }
                    if (((255 >= 255) || true) || ((float64((float32(1.5) * -2.25)) == -2.25) || true)) {
                        if ((true && (797 < 959)) && true) {
                            v9 = append(v9, (fi10 | (fi10 % ((fi10) | 1))))
                        }
                        v17 := GenS1{F0: -65}
                        if false {
                            v5 = v5
                            var v18 uint16 = 255
                            var v19 uint16 = v18
                            if false {
                                v20 := GenS2{F0: 900, F1: 42}
                                _ = v20
                            } else {
                                v19 = 0
                            }
                            fmt.Println(fi10)
                            v21 := GenS2{F0: 42, F1: 0}
                            v22 := GenS2{F0: 18446744073709551615, F1: 255}
                            v23 := GenS3{F0: ((((p0 + "foo bar") + p0) + p0) + ""), F1: (fi10 | fi10), F2: 1}
                            _ = v18
                            _ = v19
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = v17
                    }
                }
            } else {
                var v24 string = "hello"
                if (42 != 1) {
                    p0 = p0
                    var v25 uint16 = 656
                    _ = v25
                } else {
                    if len(v9) > 0 {
                        v9[0] = 127
                    }
                    v24 = v24
                    v5 -= -218
                }
                var v26 uint32 = 900
                _ = v24
                _ = v26
            }
            var v27 uint = 2
            v28 := GenS3{F0: p0, F1: fi6, F2: 1}
            _ = fi10
            _ = v27
            _ = v28
        }
        _ = fi6
        _ = v9
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 string = "foo bar"
        v31 := []int{}
        _ = fi29
        _ = v30
        _ = v31
    }
    _ = p0
    _ = v5
    return v5
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint8 = 26
    var v3 uint8 = (^(v2))
    v4 := []string{"\n\t"}
    v2 += v2
    switch "a" {
    case "foo bar":
        v2 -= 128
    case "a":
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS2{F0: (v1 + (v1 & 1)), F1: (v3 + (v2 | (v2 - uint8(v1))))}
            v3 = (^(v2))
            _ = fi32
            _ = v33
        }
    case "\n\t":
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if true {
                if (true || ((v2 % ((v3) | 1)) != (v2 / (((v2 & (v2 / ((128) | 1)))) | 1)))) {
                    var v34 int32 = -2147483648
                    fmt.Println(v2, v1)
                    _ = v34
                }
            }
            var v35 uint16 = 1
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        v3 = (v3 / ((v2) | 1))
        if (((float32(float64(((0.1 / 0.1) + 0.0))) - (3.14159 * (0.0 * (1e10 * 1e10)))) == (-2.25 - 0.0)) && (293 < 869)) {
            v37 := GenS4{F0: 2, F1: v1, F2: 1}
            fmt.Println(v1, v3)
            var v38 uint8 = 42
            _ = v37
            _ = v38
        } else {
            fmt.Println(fi36, v2)
        }
        _ = fi36
    }
    if (1 != 476) {
        v39 := GenS2{F0: (^(v1)), F1: 166}
        for fi40 := 0; fi40 < 8; fi40++ {
            v41 := GenS5{F0: 255, F1: 0}
            v4 = append(v4, "\n\t")
            fmt.Println(v2, v1)
            _ = fi40
            _ = v41
        }
        _ = v39
    } else {
        v3 += v3
    }
    v2 -= (^(v2))
    for fi42 := 0; fi42 < 8; fi42++ {
        fmt.Println(v1, v3)
        _ = fi42
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        fmt.Println(v1)
        v4 = append(v4, "a")
        if ((1.5 + 1.5) >= -2.25) {
            if false {
                fmt.Println(v3, v2)
            } else {
                v44 := GenS3{F0: "hello", F1: 982, F2: 206}
                _ = v44
            }
            var v45 uint64 = (^(v1))
            _ = v45
        } else {
            v46 := GenS5{F0: 42, F1: 255}
            var v47 int8 = -30
            _ = v46
            _ = v47
        }
        _ = fi43
    }
    if false {
        v48 := GenS2{F0: v1, F1: v2}
        _ = v48
    } else {
        var v49 float64 = 1e10
        var v50 uint = 255
        v2 = (v2 + v2)
        _ = v49
        _ = v50
    }
    var v51 uint8 = 42
    fmt.Println(v51, v2)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
}
