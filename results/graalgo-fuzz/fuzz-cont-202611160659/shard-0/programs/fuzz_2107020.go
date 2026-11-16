package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 uint64, p1 uint64, p2 int16) bool {
    v5 := []int{-1, 32768, -1}
    p0 = p1
    p2 = p2
    if (1 > -201) {
        p1 += (p0 & (p0 ^ p1))
        var v6 int32 = 459
        _ = v6
    } else {
        v7 := GenS1{F0: 255}
        if false {
            if ((("" + "a") + (("" + ("hello" + ("" + "foo bar"))) + ("foo bar" + "foo bar"))) == (("a" + ("a" + "")) + "a")) {
                var v8 float64 = (1.5 + -2.25)
                fmt.Println(p1, p2, v8)
                _ = v8
            }
        }
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(p0, fi9, fi9)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi9
        }
        _ = v7
    }
    p2 += 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return false
}
func genFn1(p0 uint16) bool {
    for fi10 := 0; fi10 < 3; fi10++ {
        p0 += (p0 / ((p0) | 1))
        for fi11 := 0; fi11 < 9; fi11++ {
            p0 = p0
            var v12 uint = 2
            var v13 float32 = 1.5
            p0 = 255
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    v14 := []int{-9223372036854775808, 0, 4294967295, 0}
    p0 = (p0 >> ((229) & 63))
    _ = p0
    _ = v14
    return true
}

func main() {
    var v1 uint = 255
    var v2 uint8 = 127
    var v3 int64 = 425
    var v4 int32 = -1
    fmt.Println(v2, v1)
    switch -802 {
    case 32768:
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 uint16 = 297
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    case 321:
        v2 += 190
    default:
        var v18 string = "\n\t"
        _ = v18
    }
    v4 -= (v4 ^ v4)
    var v19 uint8 = (v2 / ((255) | 1))
    v20 := []int{-9223372036854775808, 9223372036854775807}
    switch "hello" {
    case "hello":
        var v21 int8 = 0
        _ = v21
    case "\n\t":
        fmt.Println(v1, v3)
    case "foo bar":
        for fi22 := 0; fi22 < 8; fi22++ {
            v2 = v2
            v20 = append(v20, (fi22 ^ -624))
            if false {
                if (!(0 >= 39)) {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        fmt.Println(v3, v3)
                        v24 := GenS2{F0: v19}
                        _ = len(v20)
                        _ = fi23
                        _ = v24
                    }
                    switch fi22 {
                    case 0:
                        for fi25 := 0; fi25 < 10; fi25++ {
                            var v26 bool = true
                            if (v1 <= 343) {
                                v19 -= 42
                                v27 := GenS3{F0: v3, F1: 0}
                                _ = v27
                            } else {
                                var v28 uint = (^(v1))
                                _ = v28
                            }
                            _ = fi25
                            _ = v26
                        }
                    }
                    var v29 bool = false
                    _ = v29
                } else {
                    var v30 uint8 = (v19 % ((255) | 1))
                    _ = v30
                }
                for fi31 := 0; fi31 < 19; fi31++ {
                    v2 = (v2 ^ (v2 / (((v19 ^ v19)) | 1)))
                    _ = fi31
                }
            }
            _ = fi22
        }
    default:
        _ = len(v20)
    }
    if len(v20) > 0 {
        v20[0] = 127
    }
    var v32 int16 = 1000
    var v33 int = int(v1)
    if len(v20) > 0 {
        v20[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
