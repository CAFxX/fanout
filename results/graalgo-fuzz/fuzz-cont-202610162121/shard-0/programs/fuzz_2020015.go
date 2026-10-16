package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint16; F1 string }
func genFn0() int8 {
    var v7 uint = 65535
    if false {
        if false {
            var v8 uint8 = 42
            if (0 >= 666) {
                v8 -= (v8 + (v8 % (((v8 / ((v8) | 1))) | 1)))
                fmt.Println(v7, v8)
            } else {
                v9 := GenS1{F0: 1}
                v10 := GenS1{F0: 982}
                _ = v10
                _ = v9
            }
            var v11 int64 = 9223372036854775807
            _ = v11
            _ = v8
        }
        var v12 int8 = 1
        _ = v12
    } else {
        var v13 uint32 = 0
        v7 = uint(v13)
        _ = v13
    }
    _ = v7
    return -1
}
func genFn1(p0 bool, p1 int) int64 {
    v14 := []int{p1}
    if (!p0) {
        v15 := GenS2{F0: 0, F1: 2}
        fmt.Println(p1)
        _ = v15
    } else {
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = v14
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]string)
    var v3 uint64 = 0
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 2
    var v6 float32 = float32(-2.25)
    v16 := GenS3{F0: 0, F1: (^(v3))}
    v17 := []int{-9223372036854775808}
    v18 := v2[-304]
    var v19 uint32 = 65535
    switch "a" {
    case "\n\t":
        v20 := GenS4{F0: 100, F1: v1, F2: (v5 + 2)}
        _ = v20
    case "hello":
        if true {
            v21 := v4[(v19 + (v19 + v19))]
            if false {
                v22 := GenS1{F0: -1}
                _ = v22
            } else {
                fmt.Println(v6)
                for fi23 := 0; fi23 < 2; fi23++ {
                    if false {
                        switch fi23 {
                        case 1:
                            var v24 int64 = 9223372036854775807
                            _ = v24
                        case -946:
                            fmt.Println(fi23)
                        default:
                            v19 -= 2
                        }
                        if (uint64(v1) >= v3) {
                            v18 = v18
                            if len(v17) > 0 {
                                v17[0] = -1
                            }
                            var v25 int8 = genFn0()
                            _ = v25
                        }
                        fmt.Println(v6, v6)
                    }
                    v26 := GenS5{F0: 42, F1: 42}
                    v6 -= float32((((1.5 + float64(((float32(1.5) + v6) + 3.14159))) * (0.0 - (-2.25 + -2.25))) / (0.0 + 3.14159)))
                    _ = fi23
                    _ = v26
                }
                for fi27 := 0; fi27 < 15; fi27++ {
                    _ = len(v2)
                    _ = fi27
                }
            }
            _ = v21
        } else {
            if (!(0 == -1)) {
                v19 -= v19
            }
        }
    case "foo bar":
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 float64 = 1.5
            _ = fi28
            _ = v29
        }
    default:
        if len(v17) > 0 {
            v17[0] = 32767
        }
    }
    v1 += (v1 >> ((4294967295) & 63))
    v3 = 604
    v19 = (v19 >> ((4294967295) & 63))
    v30 := v4[(v19 ^ v19)]
    v6 = -2.25
    var v31 uint = 0
    fmt.Println(v5, v31, v18)
    fmt.Println(v3, v30)
    if true {
        var v32 int8 = int8(v1)
        _ = v32
    } else {
        var v33 int16 = -755
        if (867 != 1) {
            v30 = 266
            var v34 uint = 2
            for fi35 := 0; fi35 < 15; fi35++ {
                for fi36 := 0; fi36 < 6; fi36++ {
                    fmt.Println(v1, v30)
                    _ = len(v17)
                    _ = fi36
                }
                _ = fi35
            }
            _ = v34
        }
        _ = v33
    }
    fmt.Println(v6, v19)
    var v37 uint32 = (v30 | v30)
    var v38 uint32 = v30
    v39 := GenS6{F0: (v1 >> (((v31 % ((uint(v5)) | 1))) & 63)), F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
