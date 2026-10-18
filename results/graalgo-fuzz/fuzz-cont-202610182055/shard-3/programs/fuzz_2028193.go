package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int32; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint8; F1 bool; F2 int32 }
func genFn0() uint32 {
    if true {
        v7 := []int{1, 533, 76}
        var v8 int16 = -830
        v8 = v8
        _ = v7
        _ = v8
    }
    var v9 string = ("hello" + ("\n\t" + "a"))
    _ = v9
    return 84
}
func genFn1(p0 bool, p1 uint32, p2 int64) int64 {
    var v10 uint64 = 703
    v10 = (v10 >> ((65535) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (p2 - (p2 - (p2 * p2)))
}

func main() {
    v1 := []string{"foo bar", "a", "hello", "hello"}
    var v2 float64 = 1.5
    var v3 uint16 = 176
    v4 := []uint32{806, 1, 2147483648, 2147483648}
    var v5 int16 = 0
    var v6 bool = true
    v5 -= v5
    v5 = -974
    switch "\n\t" {
    case "\n\t":
        var v11 uint64 = 754
        _ = v11
    case "a":
        if v6 {
            v3 -= (v3 >> ((18446744073709551615) & 63))
        } else {
            var v12 int16 = 32767
            if len(v4) > 0 {
                v4[0] = genFn0()
            }
            _ = v12
        }
    case "":
        if true {
            if (4294967295 != 1) {
                v13 := GenS1{F0: 2}
                _ = v13
            } else {
                v14 := GenS1{F0: 128}
                v15 := GenS1{F0: 42}
                switch ("hello" + "hello") {
                case "":
                    if false {
                        var v16 uint8 = 255
                        var v17 bool = v6
                        _ = v16
                        _ = v17
                    }
                case "hello":
                    var v18 string = ""
                    _ = v18
                }
                _ = v14
                _ = v15
            }
        } else {
            v19 := GenS2{F0: v6, F1: v5}
            _ = v19
        }
    }
    v20 := GenS1{F0: 50}
    var v21 int16 = (v5 - (v5 * (v5 / (((v5 % ((v5) | 1))) | 1))))
    v1 = append(v1, ("foo bar" + "a"))
    fmt.Println(v6)
    v22 := GenS3{F0: 2, F1: v21}
    if ((v6 && (103 <= 0)) || v6) {
        if false {
            v23 := GenS4{F0: -975, F1: 901, F2: v21}
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 int32 = 2147483647
                v26 := GenS3{F0: v3, F1: (v21 << ((42) & 63))}
                _ = fi24
                _ = v25
                _ = v26
            }
            switch 128 {
            case -326:
                switch 32767 {
                case -679:
                    if v6 {
                        v6 = v6
                        var v27 bool = true
                        if ("hello" < ("\n\t" + "foo bar")) {
                            var v28 uint16 = v3
                            var v29 uint16 = 2
                            _ = v28
                            _ = v29
                        } else {
                            if false {
                                v5 += v21
                                for fi30 := 0; fi30 < 7; fi30++ {
                                    if (uint(v3) <= 991) {
                                        switch "a" {
                                        case "foo bar":
                                            v27 = v27
                                        case "\n\t":
                                            v2 -= float64(0.1)
                                        case "hello":
                                            for fi31 := 0; fi31 < 11; fi31++ {
                                                v27 = false
                                                switch "foo bar" {
                                                case "a":
                                                    v32 := GenS5{F0: 127, F1: (-9223372036854775808 >= -9223372036854775808), F2: 2147483647}
                                                    _ = v32
                                                case "hello":
                                                    v6 = v6
                                                case "foo bar":
                                                    var v33 int64 = -1
                                                    _ = v33
                                                default:
                                                    var v34 float64 = float64(1.5)
                                                    _ = v34
                                                }
                                                _ = fi31
                                            }
                                        }
                                        v21 += v21
                                    }
                                    fmt.Println(v3, v21)
                                    _ = fi30
                                }
                                v5 += int16(v3)
                            } else {
                                switch "" {
                                case "":
                                    v3 -= v3
                                default:
                                    fmt.Println(v6)
                                }
                            }
                        }
                        _ = v27
                    } else {
                        if v6 {
                            if (-1 != 100) {
                                var v35 bool = v6
                                _ = v35
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            v5 -= (^(v5))
                            v5 += v21
                        }
                    }
                }
            default:
                var v36 uint16 = (v3 / (((v3 | v3)) | 1))
                _ = v36
            }
            _ = v23
        }
    } else {
        switch 9223372036854775807 {
        case 213:
            var v37 uint8 = 1
            _ = v37
        case 32767:
            if false {
                for fi38 := 0; fi38 < 8; fi38++ {
                    switch (^(fi38)) {
                    case 127:
                        v5 += v21
                    case -9223372036854775808:
                        v39 := GenS1{F0: 0}
                        _ = v39
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = len(v4)
                    _ = fi38
                }
                v5 += (-(v5))
            }
        case 1:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        var v40 float32 = 0.1
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
