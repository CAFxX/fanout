package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0() int16 {
    var v4 uint32 = 442
    v4 += (v4 ^ (v4 + (v4 % (((v4 - 42)) | 1))))
    v4 -= v4
    var v5 uint16 = 0
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint64 = 612
        v8 := GenS1{F0: -599, F1: v4, F2: (^(v5))}
        v9 := GenS2{F0: 65535}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = v4
    _ = v5
    return -538
}

func main() {
    var v1 int16 = -108
    var v2 string = ("a" + "hello")
    var v3 uint = 65535
    v3 -= (v3 ^ v3)
    var v10 int16 = -180
    var v11 int16 = (v10 * genFn0())
    var v12 int16 = 32767
    var v13 int64 = 1
    var v14 string = "hello"
    v12 -= v1
    var v15 int8 = -1
    v16 := []int{9223372036854775807}
    v1 = v12
    v14 = ((v2 + ((v14 + ((("\n\t" + (("" + v2) + v2)) + "hello") + v14)) + v2)) + "")
    switch 9223372036854775807 {
    case -2147483648:
        switch v2 {
        case "hello":
            if len(v16) > 0 {
                v16[0] = -286
            }
        case "":
            var v17 uint64 = 18446744073709551615
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 13; fi18++ {
            fmt.Println(v14, v3)
            for fi19 := 0; fi19 < 12; fi19++ {
                switch v14 {
                case "a":
                    if (v14 == v2) {
                        v13 -= v13
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        v13 = int64(v3)
                    }
                case "foo bar":
                    var v20 float32 = float32(-2.25)
                    _ = v20
                case "":
                    switch (v14 + v2) {
                    case "\n\t":
                        switch v14 {
                        case "foo bar":
                            if true {
                                if (42 < (v13 >> ((v3) & 63))) {
                                    v21 := GenS3{F0: 4294967295}
                                    _ = v21
                                } else {
                                    v10 = (v12 * (v10 + (v11 - (v1 * (-(v1))))))
                                    v14 = v14
                                    var v22 uint = (^(v3))
                                    _ = len(v16)
                                    _ = v22
                                }
                                if len(v16) > 0 {
                                    v16[0] = 128
                                }
                            } else {
                                var v23 uint16 = 2
                                _ = v23
                            }
                        case "\n\t":
                            if false {
                                if len(v16) > 0 {
                                    v16[0] = fi19
                                }
                            } else {
                                fmt.Println(v14)
                            }
                        case "":
                            var v24 uint32 = 394
                            _ = v24
                        default:
                            var v25 int16 = v10
                            var v26 int32 = 1
                            v27 := GenS1{F0: 42, F1: 2147483648, F2: 334}
                            v12 = (v11 * (v1 & 32767))
                            v2 = v14
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    case "hello":
                        fmt.Println(v10, v3)
                    case "":
                        if len(v16) > 0 {
                            v16[0] = (fi18 - 32768)
                        }
                    default:
                        if (77 != 1) {
                            var v28 int = fi18
                            v1 -= v11
                            v11 -= -32768
                            _ = v28
                        } else {
                            var v29 uint8 = 138
                            var v30 int32 = 421
                            v2 = v14
                            _ = v29
                            _ = v30
                        }
                    }
                }
                _ = fi19
            }
            _ = fi18
        }
    }
    fmt.Println(v3, v10, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
