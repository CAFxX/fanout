package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int8 }
type GenS2 struct { F0 string; F1 uint; F2 bool }
type GenS3 struct { F0 int; F1 uint; F2 int }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 uint; F1 uint64; F2 int64 }
type GenS6 struct { F0 uint16; F1 uint; F2 string }
func genFn0(p0 uint64, p1 uint64, p2 uint32) uint64 {
    fmt.Println(p2, p2, p1)
    switch 128 {
    case 1:
        var v6 uint = 42
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}
func genFn1(p0 int8, p1 uint8, p2 uint8) int64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 2, F1: 32767, F2: (p0 - p0)}
        _ = fi7
        _ = v8
    }
    switch 128 {
    case 128:
        fmt.Println(p2, p1, p0)
    default:
        if (false && (2 == 2)) {
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
    }
    v10 := []int{128, 609, -16, -497}
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            p2 = 1
            fmt.Println(p2, fi11, p0)
            _ = fi12
        }
        var v13 uint = 4294967295
        _ = fi11
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    var v1 uint = 65535
    var v2 int64 = 1
    var v3 int16 = -1
    var v4 int32 = 42
    var v5 int16 = (v3 << ((42) & 63))
    var v14 int8 = 1
    v15 := []int{1}
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 bool = true
        _ = fi16
        _ = v17
    }
    if (0 != (^(v5))) {
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 int8 = 22
            fmt.Println(v4)
            for fi20 := 0; fi20 < 1; fi20++ {
                v19 = (v19 - 100)
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    }
    if len(v15) > 0 {
        v15[0] = 9223372036854775807
    }
    v21 := GenS2{F0: "\n\t", F1: 65535, F2: (32767 > 127)}
    v5 = v3
    v3 += (^(v5))
    fmt.Println(v3)
    if false {
        fmt.Println(v14, v14, v4)
        var v22 uint16 = 32767
        fmt.Println(v1)
        _ = v22
    } else {
        switch 127 {
        case 32767:
            if false {
                var v23 float32 = 3.14159
                _ = v23
            } else {
                if (1 == v1) {
                    v24 := GenS3{F0: -1, F1: (v1 % ((v1) | 1)), F2: 42}
                    switch -219 {
                    case 2147483647:
                        var v25 uint8 = 127
                        _ = v25
                    case -580:
                        var v26 int64 = (v2 & v2)
                        _ = v26
                    case 32767:
                        if len(v15) > 0 {
                            v15[0] = 9223372036854775807
                        }
                    }
                    _ = v24
                } else {
                    v27 := GenS4{F0: 42, F1: (v2 & v2)}
                    if false {
                        for fi28 := 0; fi28 < 5; fi28++ {
                            v29 := GenS5{F0: v1, F1: 618, F2: 1}
                            var v30 int16 = 42
                            for fi31 := 0; fi31 < 20; fi31++ {
                                _ = len(v15)
                                if len(v15) > 0 {
                                    v15[0] = 32767
                                }
                                v1 -= (v1 / (((v1 >> ((v1) & 63))) | 1))
                                switch "hello" {
                                case "hello":
                                    if (("\n\t" >= ("foo bar" + "a")) || true) {
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                    } else {
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                    }
                                case "":
                                    var v32 int32 = int32(v2)
                                    _ = v32
                                case "\n\t":
                                    if len(v15) > 0 {
                                        v15[0] = (fi31 / ((fi31) | 1))
                                    }
                                }
                                v2 -= 9223372036854775807
                                _ = fi31
                            }
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    }
                    v5 = v5
                    _ = v27
                }
                if (!false) {
                    fmt.Println(v14, v4)
                    v3 += (v3 / (((^(v5))) | 1))
                } else {
                    _ = len(v15)
                }
                v33 := GenS2{F0: "foo bar", F1: 928, F2: (v5 < v5)}
                _ = v33
            }
        case 946:
            var v34 int64 = (v2 + (^(v2)))
            _ = v34
        default:
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := GenS5{F0: v1, F1: 1, F2: (v2 | (v2 % ((genFn1((v14 / ((v14) | 1)), 128, 187)) | 1)))}
                var v37 string = "hello"
                _ = fi35
                _ = v36
                _ = v37
            }
        }
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        v39 := GenS6{F0: 255, F1: (v1 * (v1 >> (((^(v1))) & 63))), F2: "foo bar"}
        _ = fi38
        _ = v39
    }
    var v40 int16 = (v5 % ((-32768) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
