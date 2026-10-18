package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 bool; F2 int }
type GenS5 struct { F0 uint64; F1 int32; F2 int64 }
type GenS6 struct { F0 int }
func genFn0(p0 string) int8 {
    switch 32768 {
    case 4294967295:
        v6 := GenS1{F0: 520}
        _ = v6
    case 127:
        if false {
            if (false || (42 < 295)) {
                p0 = p0
                if (!true) {
                    fmt.Println(p0)
                    v7 := []int{450, 42, 471}
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 8; fi8++ {
                        v9 := []int{fi8, fi8, (fi8 - fi8), (-(fi8))}
                        _ = fi8
                        _ = v9
                    }
                    v10 := []int{4, 4294967295}
                    _ = v10
                }
            }
            if (42 < 2) {
                var v11 float32 = -2.25
                _ = v11
            }
        }
    case 42:
        fmt.Println(p0)
    }
    v12 := GenS1{F0: 1}
    p0 = "a"
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = v12
    return -1
}
func genFn1(p0 int16, p1 int32) int64 {
    var v13 bool = (42 >= 742)
    var v14 int64 = 1
    for fi15 := 0; fi15 < 17; fi15++ {
        switch "" {
        case "a":
            v16 := GenS2{F0: -1}
            _ = v16
        case "foo bar":
            v17 := GenS2{F0: 256}
            _ = v17
        case "\n\t":
            var v18 float64 = (-2.25 - ((0.1 + 3.14159) * 1e10))
            _ = v18
        }
        if false {
            p0 -= (p0 & (-(p0)))
        }
        _ = fi15
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS2{F0: (p1 | int32(p0))}
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return (v14 ^ (v14 / ((127) | 1)))
}

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    var v3 uint8 = 255
    var v4 int32 = 574
    var v5 int8 = 100
    fmt.Println(v3)
    _ = len(v1)
    var v21 float64 = 1.5
    var v22 string = ("foo bar" + "\n\t")
    switch -834 {
    case 128:
        if false {
            var v23 float32 = 1e10
            var v24 uint32 = 2
            _ = v23
            _ = v24
        }
    case 155:
        v25 := GenS1{F0: 255}
        _ = v25
    case 42:
        var v26 uint8 = v3
        _ = v26
    default:
        fmt.Println(v5, v5, v3)
    }
    var v27 int16 = -32768
    var v28 uint32 = 1
    switch v22 {
    case "foo bar":
        switch "foo bar" {
        case "hello":
            v1 = append(v1, 2147483647)
        default:
            if true {
                v27 += v27
                var v29 bool = (1 == v4)
                var v30 uint8 = 0
                _ = v29
                _ = v30
            } else {
                var v31 uint32 = v28
                if (true || (((((((((v21 * v21) * (3.14159 - v21)) * v21) / (v21 * float64(3.14159))) / float64((1e10 - float32(3.14159)))) * 0.1) / (float64(-2.25) + v21)) / 1.0) <= (v21 + 1.5))) {
                    if true {
                        if true {
                            v27 = 1000
                        }
                    } else {
                        var v32 bool = (!false)
                        v33 := GenS3{F0: 65535}
                        fmt.Println(v3, v28)
                        _ = v32
                        _ = v33
                    }
                    v22 = "\n\t"
                    var v34 int8 = v5
                    _ = v34
                } else {
                    v28 -= 1
                    switch -125 {
                    case 32768:
                        fmt.Println(v27)
                    default:
                        fmt.Println(v4, v3)
                    }
                }
                switch "a" {
                case "hello":
                    var v35 float32 = (-2.25 + (1.5 + 0.1))
                    _ = v35
                case "foo bar":
                    v21 -= 0.0
                case "a":
                    fmt.Println(v27)
                default:
                    if false {
                        for fi36 := 0; fi36 < 3; fi36++ {
                            if (!(!(!true))) {
                                for fi37 := 0; fi37 < 10; fi37++ {
                                    var v38 uint32 = (v28 / ((v28) | 1))
                                    var v39 float32 = (((float32(v21) + 1.5) / 3.14159) * (3.14159 * (float32((float64((-2.25 - (1e10 + 1.5))) * float64(-2.25))) * 0.0)))
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                                v40 := GenS4{F0: v28, F1: (0 != 1), F2: (fi36 - fi36)}
                                v41 := GenS2{F0: int32(v28)}
                                _ = v40
                                _ = v41
                            }
                            _ = fi36
                        }
                        fmt.Println(v21, v3, v4)
                        v22 = (v22 + "foo bar")
                    } else {
                        switch v22 {
                        case "foo bar":
                            fmt.Println(v28)
                        case "a":
                            _ = len(v2)
                        case "hello":
                            for fi42 := 0; fi42 < 20; fi42++ {
                                var v43 uint32 = (^(v28))
                                _ = fi42
                                _ = v43
                            }
                        default:
                            v27 = (v27 / ((v27) | 1))
                        }
                    }
                }
                _ = v31
            }
        }
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        if (false && ((v27 - v27) < v27)) {
            for fi44 := 0; fi44 < 6; fi44++ {
                var v45 bool = (1 != 18446744073709551615)
                switch (fi44 % (((fi44 * (fi44 * 783))) | 1)) {
                case -9223372036854775808:
                    var v46 string = (v22 + ((v22 + ((v22 + v22) + (v22 + v22))) + v22))
                    _ = v46
                case 128:
                    var v47 uint = 42
                    _ = v47
                case 0:
                    fmt.Println(v45)
                }
                _ = fi44
                _ = v45
            }
            if true {
                var v48 uint8 = v3
                _ = v48
            } else {
                switch 9223372036854775807 {
                case 4294967295:
                    var v49 int8 = v5
                    _ = v49
                case -1:
                    v2[(("" + (((v22 + v22) + ((v22 + v22) + v22)) + v22)) + v22)] = v22
                }
            }
        }
    }
    var v50 uint32 = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi51 := 0; fi51 < 8; fi51++ {
        switch (fi51 << ((21) & 63)) {
        case 2147483647:
            v52 := GenS5{F0: 2, F1: v4, F2: -110}
            _ = v52
        default:
            v50 = (v28 - 65535)
        }
        fmt.Println(v28)
        _ = fi51
    }
    if (-9223372036854775808 >= 9223372036854775807) {
        v27 += v27
        for fi53 := 0; fi53 < 14; fi53++ {
            v22 = v22
            _ = fi53
        }
        var v54 uint8 = 255
        _ = v54
    } else {
        v55 := GenS4{F0: (v28 ^ 42), F1: true, F2: 128}
        switch ("foo bar" + "foo bar") {
        case "":
            fmt.Println(v27)
        default:
            v21 = 0.1
        }
        _ = v55
    }
    if (488 != 496) {
        v56 := GenS6{F0: 64}
        _ = v56
    } else {
        var v57 bool = true
        _ = v57
    }
    if (false || (!true)) {
        _ = len(v2)
    }
    if true {
        v28 -= (^(v28))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
