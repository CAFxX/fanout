package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint8; F1 string; F2 int16 }
func genFn0() bool {
    var v4 uint32 = 4294967295
    v4 += v4
    _ = v4
    return true
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "foo bar"
    var v3 uint64 = 42
    v3 = (v3 ^ (v3 - (v3 | (v3 - v3))))
    var v5 int32 = -50
    v6 := []int{-9223372036854775808}
    var v7 uint16 = 50
    v2 = (v2 + ((v2 + "hello") + v2))
    _ = len(v1)
    if (!(-835 != 461)) {
        v6 = append(v6, 1)
        if ((128 == 42) && (v5 >= v5)) {
            switch 2147483647 {
            case 4294967295:
                for fi8 := 0; fi8 < 16; fi8++ {
                    if (2 <= 368) {
                        fmt.Println(v5, fi8)
                        v5 = v5
                        var v9 uint64 = v3
                        _ = v9
                    } else {
                        if true {
                            if true {
                                for fi10 := 0; fi10 < 10; fi10++ {
                                    v2 = v2
                                    _ = fi10
                                }
                                delete(v1, fi8)
                            } else {
                                var v11 uint16 = v7
                                v3 += v3
                                _ = v11
                            }
                            v7 = 2
                        } else {
                            var v12 bool = false
                            v3 = v3
                            v12 = ((v7 != v7) && v12)
                            _ = v12
                        }
                    }
                    _ = fi8
                }
            case -250:
                if false {
                    if false {
                        v13 := GenS1{F0: (v7 / ((480) | 1)), F1: 128, F2: v2}
                        _ = v13
                    } else {
                        v1[-420] = 42
                        if true {
                            fmt.Println(v3, v2, v3)
                        } else {
                            for fi14 := 0; fi14 < 7; fi14++ {
                                fmt.Println(v5)
                                var v15 float32 = -2.25
                                for fi16 := 0; fi16 < 10; fi16++ {
                                    var v17 uint32 = uint32(v3)
                                    var v18 int16 = 289
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi14
                                _ = v15
                            }
                            fmt.Println(v2)
                        }
                        v19 := GenS2{F0: 128, F1: (^(v7)), F2: 4294967295}
                        _ = v19
                    }
                    var v20 uint = 258
                    _ = v20
                } else {
                    v21 := GenS3{F0: 9223372036854775807, F1: (579 >= 2147483647)}
                    _ = v21
                }
            case 0:
                var v22 int8 = -128
                _ = v22
            default:
                var v23 int32 = 46
                _ = v23
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                v25 := GenS1{F0: v7, F1: (fi24 ^ (fi24 * (fi24 | (fi24 * (fi24 ^ (fi24 * (fi24 << ((65535) & 63)))))))), F2: v2}
                if false {
                    v7 = (v7 % ((v7) | 1))
                } else {
                    var v26 uint = 18446744073709551615
                    v7 = v7
                    if true {
                        if len(v6) > 0 {
                            v6[0] = int(v3)
                        }
                        v27 := GenS3{F0: 312, F1: true}
                        v5 -= v5
                        _ = v27
                    }
                    _ = v26
                }
                _ = fi24
                _ = v25
            }
        } else {
            v6 = append(v6, 32768)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    }
    fmt.Println(v7, v7)
    for fi28 := 0; fi28 < 14; fi28++ {
        v7 = v7
        _ = fi28
    }
    switch 1 {
    case 127:
        _ = len(v1)
    case 4294967295:
        v7 += v7
    default:
        var v29 float64 = 0.1
        _ = v29
    }
    switch 32767 {
    case 128:
        v3 = 332
    case 32767:
        v1[32767] = 4294967295
    case 127:
        switch 4294967295 {
        case 2147483647:
            v2 = "\n\t"
        case -2147483648:
            v30 := GenS2{F0: 42, F1: (v7 / ((65535) | 1)), F2: 1}
            _ = v30
        case 32767:
            v31 := GenS1{F0: v7, F1: 9223372036854775807, F2: "hello"}
            _ = v31
        default:
            if ((1.5 - 1e10) > 0.1) {
                if (false && true) {
                    if (!false) {
                        if true {
                            for fi32 := 0; fi32 < 7; fi32++ {
                                var v33 float32 = 3.14159
                                v2 = "hello"
                                _ = fi32
                                _ = v33
                            }
                            v34 := GenS3{F0: 9223372036854775807, F1: (((3.14159 >= 0.1) || ((((v2 + v2) + ("foo bar" + ((("a" + "") + v2) + ("a" + v2)))) != (v2 + v2)) && (4294967295 != 32768))) && ((!false) || true))}
                            _ = v34
                        } else {
                            v35 := GenS4{F0: 255, F1: v2, F2: 494}
                            var v36 int8 = 1
                            _ = v35
                            _ = v36
                        }
                        var v37 int64 = int64(v7)
                        _ = v37
                    }
                }
            }
        }
    }
    var v38 string = "\n\t"
    for fi39 := 0; fi39 < 10; fi39++ {
        for fi40 := 0; fi40 < 18; fi40++ {
            _ = len(v6)
            v2 = (v2 + "hello")
            if (1e10 >= (float32(0.1) / 1.0)) {
                fmt.Println(v3, v5)
                if ((3.14159 <= float64(0.1)) || (255 > 255)) {
                    v5 -= (^(v5))
                    var v41 uint16 = 65535
                    fmt.Println(v3, v5, v38)
                    _ = v41
                }
            } else {
                v42 := GenS3{F0: 1, F1: (2 <= 240)}
                v5 -= (v5 << ((42) & 63))
                _ = v42
            }
            _ = fi40
        }
        _ = fi39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v5
    _ = v6
    _ = v7
}
