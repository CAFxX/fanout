package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int64; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 592
    var v3 uint64 = 42
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: (v2 << ((v2) & 63))}
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 196}
    var v7 int64 = 1
    v7 -= v7
    v3 = v3
    _ = len(v1)
    if (0 == int8(v2)) {
        fmt.Println(v7)
        if true {
            fmt.Println(v7, v3)
        }
        switch 127 {
        case -9223372036854775808:
            if false {
                for fi8 := 0; fi8 < 6; fi8++ {
                    v2 += v2
                    if false {
                        if ((^(v2)) < (^(v2))) {
                            v9 := GenS1{F0: 1}
                            _ = v9
                        } else {
                            v10 := []int{(fi8 % ((fi8) | 1)), (fi8 % (((fi8 << (((v2 / ((v2) | 1))) & 63))) | 1))}
                            _ = v10
                        }
                        v3 = v3
                        if (true && (!((v7 + (v7 & 49)) > 0))) {
                            var v11 int64 = (-(v7))
                            _ = v11
                        }
                    } else {
                        v12 := GenS1{F0: v2}
                        _ = v12
                    }
                    _ = fi8
                }
            } else {
                var v13 int16 = 32767
                v14 := []int{42, -2147483648, 2147483647}
                _ = v13
                _ = v14
            }
        default:
            v15 := GenS1{F0: v2}
            _ = v15
        }
    } else {
        var v16 int32 = 0
        var v17 int8 = 1
        for fi18 := 0; fi18 < 12; fi18++ {
            _ = len(v1)
            for fi19 := 0; fi19 < 7; fi19++ {
                if ((v17 | 100) != v17) {
                    v16 = (v16 - (v16 >> ((v2) & 63)))
                    if (((-(v16)) == v16) && true) {
                        switch "foo bar" {
                        case "\n\t":
                            v20 := GenS2{F0: uint32(v2), F1: v7, F2: 32767}
                            _ = v20
                        }
                    }
                    var v21 string = "hello"
                    _ = v21
                } else {
                    if true {
                        v17 -= 0
                        v16 = (v16 ^ v16)
                        fmt.Println(v3, v7, v16)
                    } else {
                        _ = len(v1)
                    }
                    v16 -= (v16 / (((v16 * v16)) | 1))
                    fmt.Println(v2, v16)
                    fmt.Println(v17)
                }
                _ = fi19
            }
            _ = fi18
        }
        _ = v16
        _ = v17
    }
    var v22 int8 = -86
    for fi23 := 0; fi23 < 12; fi23++ {
        var v24 uint8 = 156
        delete(v1, -194)
        _ = fi23
        _ = v24
    }
    var v25 float64 = 0.1
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS2{F0: 2, F1: 9223372036854775807, F2: 243}
        var v28 uint16 = uint16(v2)
        v29 := GenS3{F0: (!false)}
        _ = fi26
        _ = v27
        _ = v28
        _ = v29
    }
    var v30 int32 = 0
    if ((!(!true)) && (!true)) {
        delete(v1, -9223372036854775808)
        v3 -= (v3 << ((698) & 63))
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        for fi32 := 0; fi32 < 15; fi32++ {
            if false {
                v22 = (-(v22))
            } else {
                var v33 int32 = (v30 & (v30 / ((2147483647) | 1)))
                var v34 int8 = -44
                _ = v33
                _ = v34
            }
            _ = fi32
        }
        v3 = v3
        v35 := GenS4{F0: v30, F1: 42}
        _ = fi31
        _ = v35
    }
    v25 -= 3.14159
    var v36 int32 = 490
    var v37 string = ("hello" + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v6
    _ = v7
}
