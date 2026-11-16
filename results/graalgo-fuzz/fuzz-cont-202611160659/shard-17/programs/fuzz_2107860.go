package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int16; F2 bool }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int32 }
func genFn0() int {
    if true {
        switch -9223372036854775808 {
        case 127:
            if false {
                var v6 int16 = -1
                _ = v6
            }
        case -2147483648:
            var v7 uint64 = 477
            _ = v7
        }
        var v8 bool = (3.14159 > 3.14159)
        _ = v8
    }
    var v9 float32 = 3.14159
    v10 := []int{}
    _ = v10
    _ = v9
    return -1
}
func genFn1(p0 uint32, p1 uint64, p2 uint32) bool {
    var v11 string = "foo bar"
    p0 = 676
    var v12 bool = (0 > 0)
    v13 := []int{}
    fmt.Println(v12, v12, v12)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return true
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 2147483647
    v3 := []uint32{}
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]string)
    switch "foo bar" {
    case "\n\t":
        var v14 int32 = (v2 / (((v2 ^ (v2 >> ((18446744073709551615) & 63)))) | 1))
        _ = v14
    case "":
        fmt.Println(v2)
    case "a":
        var v15 uint8 = 208
        _ = v15
    }
    var v16 uint8 = 42
    v2 = -2147483648
    for fi17 := 0; fi17 < 12; fi17++ {
        if (127 != -1) {
            v18 := GenS1{F0: -559}
            _ = v18
        }
        v3 = append(v3, 736)
        if false {
            v16 -= (v16 % (((v16 % ((v16) | 1))) | 1))
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 float64 = ((-2.25 * 1e10) / 1e10)
        fmt.Println(fi19, v20)
        var v21 int = int(v16)
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: 1000}
    var v23 uint16 = 819
    v24 := GenS1{F0: -951}
    for fi25 := 0; fi25 < 9; fi25++ {
        fmt.Println(v23)
        fmt.Println(v16, fi25, v2)
        for fi26 := 0; fi26 < 17; fi26++ {
            v23 -= v23
            if len(v3) > 0 {
                v3[0] = 248
            }
            v2 = (^(v2))
            _ = fi26
        }
        _ = fi25
    }
    var v27 int64 = -1
    _ = len(v3)
    if false {
        _ = len(v4)
        for fi28 := 0; fi28 < 10; fi28++ {
            switch "foo bar" {
            case "a":
                for fi29 := 0; fi29 < 2; fi29++ {
                    v23 = v23
                    _ = fi29
                }
            case "\n\t":
                for fi30 := 0; fi30 < 5; fi30++ {
                    v31 := GenS2{F0: 42}
                    _ = fi30
                    _ = v31
                }
            case "hello":
                for fi32 := 0; fi32 < 17; fi32++ {
                    fmt.Println(fi32, fi28)
                    v27 += (v27 - v27)
                    _ = fi32
                }
            }
            var v33 int64 = v27
            var v34 uint64 = 42
            _ = fi28
            _ = v33
            _ = v34
        }
    }
    var v35 float32 = 1.5
    v23 = 255
    if (v16 <= 114) {
        fmt.Println(v2)
        if (!(!(!true))) {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v36 := GenS3{F0: ("\n\t" + "hello"), F1: 1, F2: (!(false || false))}
                for fi37 := 0; fi37 < 6; fi37++ {
                    for fi38 := 0; fi38 < 7; fi38++ {
                        var v39 float32 = ((v35 * (float32(-2.25) * v35)) * ((1.5 + v35) - 0.0))
                        v40 := GenS4{F0: 1, F1: true}
                        v2 -= (v2 | 1)
                        if (v2 != 0) {
                            v41 := GenS2{F0: 0}
                            fmt.Println(v27)
                            if (true && true) {
                                switch "" {
                                case "\n\t":
                                    if false {
                                        v42 := GenS5{F0: 4294967295, F1: 42}
                                        _ = v42
                                    } else {
                                        var v43 int16 = -772
                                        _ = v43
                                    }
                                case "":
                                    var v44 float32 = -2.25
                                    _ = v44
                                case "a":
                                    v45 := v4[331]
                                    _ = v45
                                default:
                                    v46 := GenS6{F0: (v16 + v16), F1: -2147483648}
                                    _ = v46
                                }
                                v39 += v35
                            } else {
                                fmt.Println(fi38, v39)
                            }
                            _ = v41
                        }
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    var v47 uint8 = v16
                    _ = fi37
                    _ = v47
                }
                _ = v36
            }
            if (false && (472 <= 18446744073709551615)) {
                for fi48 := 0; fi48 < 6; fi48++ {
                    var v49 float32 = (float32(1e10) + -2.25)
                    _ = fi48
                    _ = v49
                }
                var v50 float32 = 0.0
                _ = v50
            }
        } else {
            v16 += (v16 + (v16 & (v16 << ((47) & 63))))
            fmt.Println(v35)
            var v51 int32 = (-(v2))
            _ = v51
        }
        v2 = (v2 / ((v2) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
