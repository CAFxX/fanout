package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint) int16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 int8 = 1
        var v7 int64 = 523
        v7 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{-1}
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 string = ""
        if (!false) {
            v11 := GenS1{F0: 32767, F1: (p0 >> (((p0 ^ (p0 - 255))) & 63)), F2: 4294967295}
            var v12 int64 = 371
            _ = v11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 int64 = -9223372036854775808
    fmt.Println(v13, v13)
    _ = p0
    _ = v13
    _ = v8
    return 0
}
func genFn1(p0 string) int {
    p0 = p0
    v14 := GenS2{F0: 9223372036854775807, F1: 1}
    _ = p0
    _ = v14
    return 25
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    v3 := []uint32{4294967295, 4294967295}
    var v4 uint = 18446744073709551615
    for fi15 := 0; fi15 < 14; fi15++ {
        if len(v1) > 0 {
            v1[0] = 42
        }
        v16 := GenS3{F0: 733}
        _ = fi15
        _ = v16
    }
    var v17 int32 = -946
    if true {
        if (!(1 >= 4294967295)) {
            v18 := v2[931]
            for fi19 := 0; fi19 < 11; fi19++ {
                v20 := GenS3{F0: -505}
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            v17 -= (v17 | -2147483648)
            if true {
                for fi21 := 0; fi21 < 11; fi21++ {
                    var v22 float32 = (1e10 - -2.25)
                    _ = fi21
                    _ = v22
                }
                var v23 int16 = genFn0(v4)
                var v24 int64 = 1
                _ = v23
                _ = v24
            } else {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v4, v4)
                }
            }
        }
    }
    if false {
        v17 = v17
        fmt.Println(v4)
    } else {
        _ = len(v2)
        v25 := GenS4{F0: 32767}
        for fi26 := 0; fi26 < 11; fi26++ {
            fmt.Println(v17)
            v27 := GenS2{F0: -9223372036854775808, F1: -581}
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    v17 -= v17
    v28 := GenS3{F0: 653}
    fmt.Println(v4)
    var v29 int16 = 42
    for fi30 := 0; fi30 < 14; fi30++ {
        for fi31 := 0; fi31 < 17; fi31++ {
            var v32 bool = false
            fmt.Println(v32, v29)
            var v33 uint64 = 0
            _ = fi31
            _ = v32
            _ = v33
        }
        var v34 int32 = 42
        if true {
            var v35 uint8 = 2
            _ = v35
        } else {
            if false {
                for fi36 := 0; fi36 < 3; fi36++ {
                    var v37 int8 = 1
                    var v38 uint8 = 128
                    if (244 < 2147483647) {
                        v37 += (v37 * (^(v37)))
                    } else {
                        var v39 int8 = (v37 >> ((v4) & 63))
                        if false {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            fmt.Println(v17)
                        }
                        _ = v39
                    }
                    switch (-(fi36)) {
                    case 32767:
                        var v40 uint32 = 331
                        _ = v40
                    case -1:
                        fmt.Println(fi36, fi36)
                    case 9223372036854775807:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        if (-2.25 < ((0.1 - (1e10 * (-2.25 * 3.14159))) * (-2.25 - (1.5 * 3.14159)))) {
                            var v41 uint32 = 4294967295
                            fmt.Println(v38, v34, fi30)
                            _ = v41
                        } else {
                            fmt.Println(v37, v38)
                        }
                    }
                    var v42 int = genFn1("hello")
                    v17 = (v34 + v34)
                    var v43 string = ""
                    if (v34 >= (v34 % ((18) | 1))) {
                        fmt.Println(v4, v29)
                        fmt.Println(v38)
                    }
                    _ = fi36
                    _ = v37
                    _ = v38
                    _ = v42
                    _ = v43
                }
                v2[47] = 2
            } else {
                fmt.Println(v34)
            }
            v29 += (v29 ^ v29)
        }
        _ = fi30
        _ = v34
    }
    fmt.Println(v29, v4, v29)
    v44 := GenS2{F0: -420, F1: (v29 / ((v29) | 1))}
    if (!false) {
        var v45 float64 = float64((1e10 / 1.5))
        v4 -= v4
        _ = v45
    }
    fmt.Println(v4, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v44
}
