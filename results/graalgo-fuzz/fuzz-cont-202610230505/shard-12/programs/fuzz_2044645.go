package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 2
    var v3 uint32 = 42
    var v4 uint = uint(v2)
    var v5 string = (("foo bar" + ("foo bar" + "a")) + ("hello" + ("foo bar" + "hello")))
    var v6 uint = (v4 + (v4 / ((v4) | 1)))
    switch v5 {
    case "\n\t":
        var v7 float32 = 0.1
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 19; fi8++ {
            if (255 != 2) {
                for fi9 := 0; fi9 < 14; fi9++ {
                    switch v5 {
                    case "hello":
                        var v10 int64 = 319
                        _ = v10
                    case "\n\t":
                        if (42 >= 1) {
                            fmt.Println(fi8, v2)
                            var v11 uint = (v6 & v6)
                            v12 := GenS1{F0: v5, F1: (v3 - (v3 & v3)), F2: 176}
                            _ = v11
                            _ = v12
                        }
                    default:
                        switch v5 {
                        case "a":
                            v2 += (v2 / ((v2) | 1))
                        }
                    }
                    _ = fi9
                }
            }
            _ = fi8
        }
    case "":
        v13 := GenS1{F0: (v5 + (v5 + v5)), F1: (v3 % ((v3) | 1)), F2: 42}
        _ = v13
    }
    v14 := GenS2{F0: v5, F1: (v3 | (v3 | v3))}
    for fi15 := 0; fi15 < 11; fi15++ {
        v16 := GenS2{F0: v5, F1: 2147483647}
        v17 := GenS3{F0: (42 != 413)}
        fmt.Println(v2, v4)
        _ = fi15
        _ = v16
        _ = v17
    }
    v2 = v2
    v18 := GenS4{F0: 127}
    var v19 string = v5
    for fi20 := 0; fi20 < 12; fi20++ {
        _ = len(v1)
        _ = fi20
    }
    var v21 int16 = 42
    var v22 int8 = -14
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 uint16 = v2
        if true {
            var v25 uint8 = uint8(v24)
            switch 32767 {
            case -9223372036854775808:
                for fi26 := 0; fi26 < 11; fi26++ {
                    if (!(!(-1 >= -1))) {
                        fmt.Println(v2)
                        var v27 float32 = (-2.25 / 1.0)
                        v28 := GenS3{F0: true}
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS5{F0: (v4 ^ v6)}
                        fmt.Println(fi26, fi23)
                        _ = v29
                    }
                    _ = fi26
                }
            default:
                v1[fi23] = (v3 ^ v3)
            }
            _ = v25
        }
        var v30 uint16 = (v2 - v24)
        _ = fi23
        _ = v24
        _ = v30
    }
    switch -1 {
    case 9223372036854775807:
        for fi31 := 0; fi31 < 4; fi31++ {
            switch (fi31 + (fi31 << ((1) & 63))) {
            case 0:
                v1[(fi31 | 1)] = (^(v3))
            default:
                var v32 uint = 0
                _ = v32
            }
            var v33 bool = true
            v34 := []int{fi31, (fi31 << ((v4) & 63)), fi31}
            _ = fi31
            _ = v33
            _ = v34
        }
    case -1:
        delete(v1, 32767)
    case 745:
        fmt.Println(v22)
    default:
        if false {
            v35 := GenS2{F0: v5, F1: 995}
            v3 = (v3 << ((v4) & 63))
            switch (v5 + "\n\t") {
            case "a":
                var v36 uint32 = (v3 % ((v3) | 1))
                _ = v36
            default:
                v22 += (v22 | v22)
            }
            _ = v35
        } else {
            var v37 int = 32768
            _ = v37
        }
    }
    if false {
        v1[1] = uint32(v2)
        var v38 int16 = (v21 & v21)
        v21 = -882
        _ = v38
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        var v40 uint8 = 0
        _ = fi39
        _ = v40
    }
    v5 = (v19 + v19)
    v41 := GenS5{F0: 119}
    v21 -= (v21 + v21)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
