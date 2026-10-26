package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 string; F1 uint64; F2 bool }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 int16, p1 int64) uint16 {
    if true {
        var v4 float32 = -2.25
        fmt.Println(p0)
        for fi5 := 0; fi5 < 2; fi5++ {
            for fi6 := 0; fi6 < 12; fi6++ {
                p1 -= (p1 % (((p1 << ((2) & 63))) | 1))
                switch "hello" {
                case "":
                    p1 += p1
                case "foo bar":
                    switch "\n\t" {
                    case "a":
                        var v7 int8 = 42
                        switch 0 {
                        case 32768:
                            v8 := []int{(^(fi5)), (^(fi5)), fi5, (fi5 | fi6)}
                            _ = v8
                        case 1:
                            v9 := GenS1{F0: 42, F1: v7, F2: 128}
                            _ = v9
                        default:
                            switch "hello" {
                            case "":
                                v10 := GenS1{F0: 353, F1: (-(v7)), F2: 1}
                                _ = v10
                            case "a":
                                v4 += v4
                            default:
                                v11 := GenS2{F0: 66, F1: p1, F2: 1}
                                _ = v11
                            }
                        }
                        var v12 string = ""
                        v13 := GenS1{F0: 255, F1: 1, F2: 28}
                        _ = v12
                        _ = v13
                        _ = v7
                    case "foo bar":
                        if false {
                            var v14 int = fi5
                            _ = v14
                        }
                        var v15 int32 = 404
                        _ = v15
                    case "\n\t":
                        fmt.Println(p0, v4)
                    default:
                        var v16 int64 = 9223372036854775807
                        _ = v16
                    }
                }
                _ = fi6
            }
            switch ("hello" + "\n\t") {
            case "a":
                var v17 int64 = (p1 - p1)
                _ = v17
            case "hello":
                v18 := []int{fi5, fi5, (fi5 * (-(fi5))), (fi5 | (fi5 | fi5))}
                _ = v18
            case "":
                v19 := GenS3{F0: 18446744073709551615, F1: (fi5 & fi5)}
                _ = v19
            default:
                fmt.Println(fi5, p1)
            }
            _ = fi5
        }
        _ = v4
    }
    var v20 uint16 = 1
    if ((2147483647 > -34) || (float32(3.14159) == (1.5 * 3.14159))) {
        v20 = (v20 * 2)
        v21 := []int{42}
        v22 := GenS2{F0: 43, F1: (p1 << ((1) & 63)), F2: 1}
        _ = v21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v20
    return (v20 % (((^(v20))) | 1))
}

func main() {
    var v1 int = -906
    var v2 int32 = 0
    var v3 uint64 = 18446744073709551615
    v3 = v3
    for fi23 := 0; fi23 < 5; fi23++ {
        if false {
            var v24 bool = ((false || (-32768 != -400)) || (!false))
            var v25 int16 = 42
            _ = v24
            _ = v25
        } else {
            fmt.Println(v3)
            v3 += v3
        }
        v26 := GenS3{F0: v3, F1: (^(fi23))}
        _ = fi23
        _ = v26
    }
    fmt.Println(v1, v3)
    var v27 uint32 = 42
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 11; fi29++ {
            v1 -= (fi28 * fi28)
            v30 := GenS4{F0: "\n\t", F1: 1, F2: false}
            v1 = (fi29 ^ (^(v1)))
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            v3 = v3
            _ = fi31
        }
        _ = fi28
    }
    switch ("foo bar" + "hello") {
    case "":
        fmt.Println(v1, v3, v3)
    }
    v32 := GenS3{F0: 4294967295, F1: v1}
    var v33 int32 = (v2 & v2)
    v34 := []int{}
    var v35 int8 = 1
    v34 = append(v34, v1)
    switch v1 {
    case 1:
        fmt.Println(v35, v2)
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        if len(v34) > 0 {
            v34[0] = (-(fi36))
        }
        v37 := GenS1{F0: uint(v27), F1: (v35 % (((^(v35))) | 1)), F2: 0}
        v27 = (v27 >> ((286) & 63))
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        _ = len(v34)
        _ = fi38
    }
    v39 := GenS1{F0: 2, F1: v35, F2: 128}
    if false {
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
        v33 = v33
    } else {
        var v40 uint8 = 42
        if ((v35 >> ((18446744073709551615) & 63)) <= 42) {
            var v41 uint8 = v40
            _ = v41
        }
        _ = v40
    }
    v42 := GenS5{F0: false, F1: (v35 + v35)}
    for fi43 := 0; fi43 < 2; fi43++ {
        switch "a" {
        case "a":
            var v44 bool = true
            _ = v44
        case "":
            if true {
                v45 := GenS3{F0: (v3 ^ v3), F1: fi43}
                v46 := GenS6{F0: 42, F1: -128}
                var v47 int8 = 127
                _ = v45
                _ = v46
                _ = v47
            } else {
                if true {
                    v48 := GenS7{F0: 0}
                    v33 -= v33
                    v49 := GenS8{F0: 2147483647}
                    _ = v48
                    _ = v49
                }
                v27 += 65535
                if len(v34) > 0 {
                    v34[0] = (v1 ^ int(v2))
                }
            }
        case "\n\t":
            for fi50 := 0; fi50 < 5; fi50++ {
                v27 -= 65535
                _ = fi50
            }
        }
        for fi51 := 0; fi51 < 11; fi51++ {
            if len(v34) >= 2 {
                v34 = v34[1:len(v34)]
            }
            _ = fi51
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v39
    _ = v42
}
