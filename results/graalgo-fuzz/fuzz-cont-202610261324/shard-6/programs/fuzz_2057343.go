package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 int8, p2 int8) uint16 {
    v6 := []int{122, 9223372036854775807, -423}
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 uint, p1 string, p2 uint32) uint16 {
    p1 = "hello"
    var v7 float32 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    var v1 int64 = 1
    v2 := []string{}
    v3 := []uint32{382}
    var v4 uint16 = 27
    v5 := []uint32{19}
    v1 -= v1
    if false {
        switch ("foo bar" + "hello") {
        case "":
            v5 = append(v5, 65535)
        case "\n\t":
            fmt.Println(v4)
        case "a":
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v8 uint = 845
                switch "\n\t" {
                case "hello":
                    fmt.Println(v4, v8, v1)
                case "\n\t":
                    for fi9 := 0; fi9 < 10; fi9++ {
                        if len(v5) > 0 {
                            v5[0] = 375
                        }
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        _ = fi9
                    }
                case "foo bar":
                    var v10 uint = v8
                    _ = v10
                default:
                    for fi11 := 0; fi11 < 4; fi11++ {
                        if ((1e10 < (1.5 * 1e10)) || false) {
                            switch 128 {
                            case 4294967295:
                                var v12 int8 = 42
                                _ = v12
                            }
                            v13 := GenS1{F0: 42, F1: 2147483647}
                            _ = v13
                        } else {
                            v14 := GenS2{F0: 401}
                            var v15 bool = (-738 >= -1)
                            fmt.Println(fi11, v8, fi11)
                            _ = v14
                            _ = v15
                        }
                        var v16 float32 = 0.0
                        v1 -= -1
                        _ = fi11
                        _ = v16
                    }
                }
                v8 = (^(v8))
                _ = v8
            }
        }
    } else {
        v4 -= v4
    }
    var v17 uint8 = 174
    if false {
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := GenS1{F0: 2, F1: 1}
                _ = fi19
                _ = v20
            }
            var v21 bool = true
            v17 = (v17 - (v17 % (((^(v17))) | 1)))
            _ = fi18
            _ = v21
        }
        v22 := GenS1{F0: 42, F1: 4294967295}
        _ = v22
    } else {
        if (18446744073709551615 <= 0) {
            v23 := GenS1{F0: uint64(v17), F1: 299}
            var v24 string = "a"
            _ = v23
            _ = v24
        } else {
            if (1 == 0) {
                fmt.Println(v1, v17)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                v25 := GenS2{F0: 32767}
                v1 -= (v1 << ((1) & 63))
                v1 += -489
                _ = v25
            }
            var v26 uint = 710
            _ = v26
        }
        v1 -= (v1 / (((v1 + -993)) | 1))
        v27 := GenS2{F0: -721}
        _ = v27
    }
    var v28 int32 = -347
    v4 = 42
    switch 9223372036854775807 {
    case 1:
        v4 = genFn1(255, "hello", 220)
    case -2147483648:
        v29 := GenS1{F0: 0, F1: 1}
        _ = v29
    case 4294967295:
        var v30 string = "\n\t"
        _ = v30
    default:
        v31 := GenS1{F0: 42, F1: 65535}
        _ = v31
    }
    if len(v2) > 0 {
        v2[0] = (("a" + "") + (("" + "foo bar") + "foo bar"))
    }
    v4 = genFn1(4294967295, "foo bar", 2147483647)
    var v32 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
