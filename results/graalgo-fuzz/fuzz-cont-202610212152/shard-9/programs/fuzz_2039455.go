package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0(p0 int8, p1 bool) string {
    var v5 int16 = 0
    var v6 int64 = 1
    var v7 uint32 = 2147483647
    var v8 int16 = (v5 * (v5 | (v5 | v5)))
    if p1 {
        v9 := []int{4294967295, -252, 2147483647}
        if len(v9) > 0 {
            v9[0] = -300
        }
        var v10 bool = ((((0.0 - 1e10) / 1e10) * float64(float32(0.0))) >= 0.1)
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: (v6 / (((-(v6))) | 1))}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 uint = 66
    var v2 int8 = -1
    var v3 uint16 = 42
    var v4 uint = (^(v1))
    v12 := []int{}
    v13 := GenS1{F0: 362}
    v14 := GenS2{F0: v2, F1: (v2 + v2), F2: 127}
    var v15 uint8 = 1
    if false {
        v4 += 14
        switch 127 {
        case -9223372036854775808:
            for fi16 := 0; fi16 < 10; fi16++ {
                if (4294967295 == 433) {
                    var v17 uint = 699
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 string = "\n\t"
                        if false {
                            fmt.Println(fi16)
                            if true {
                                var v20 uint16 = uint16(v15)
                                _ = v20
                            } else {
                                fmt.Println(fi16)
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                                switch fi18 {
                                case 32767:
                                    var v21 int8 = (v2 + v2)
                                    _ = v21
                                case -2147483648:
                                    var v22 string = ((("\n\t" + (v19 + v19)) + (("hello" + "a") + ((v19 + v19) + ("a" + "foo bar")))) + v19)
                                    _ = v22
                                case 4294967295:
                                    fmt.Println(v2)
                                default:
                                    v23 := GenS3{F0: 9223372036854775807}
                                    _ = v23
                                }
                            }
                            v2 = 127
                        }
                        _ = fi18
                        _ = v19
                    }
                }
                _ = fi16
            }
        default:
            var v24 float32 = (-2.25 + 3.14159)
            _ = v24
        }
    } else {
        switch 128 {
        case -2147483648:
            if (!(42 > 65535)) {
                var v25 bool = (971 >= -32768)
                var v26 int16 = 221
                _ = v25
                _ = v26
            } else {
                v27 := GenS4{F0: 729, F1: 274, F2: (v2 & (v2 ^ (v2 >> (((v1 / ((v4) | 1))) & 63))))}
                fmt.Println(v2, v3, v3)
                _ = v27
            }
        }
    }
    var v28 int64 = 0
    for fi29 := 0; fi29 < 8; fi29++ {
        v15 -= (v15 | (v15 + v15))
        _ = fi29
    }
    fmt.Println(v28, v2, v15)
    switch "" {
    case "hello":
        v2 -= (v2 + (^(v2)))
    case "":
        var v30 int64 = (v28 & (^(v28)))
        _ = v30
    case "\n\t":
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
