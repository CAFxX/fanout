package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint; F2 uint }
type GenS4 struct { F0 uint; F1 int64; F2 int16 }
type GenS5 struct { F0 string; F1 int64; F2 int32 }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 int64 }
func genFn0(p0 int64, p1 string, p2 int32) int {
    p2 -= (p2 * 94)
    if false {
        var v7 bool = true
        _ = v7
    } else {
        switch -2147483648 {
        case 518:
            var v8 float32 = 1.5
            _ = v8
        }
        switch 9223372036854775807 {
        case 32767:
            v9 := []int{}
            _ = v9
        case 578:
            var v10 uint64 = 515
            _ = v10
        case -719:
            switch -9223372036854775808 {
            case -1:
                if true {
                    fmt.Println(p0)
                }
            case 279:
                v11 := []int{-335, -2147483648, 0}
                _ = v11
            case 607:
                v12 := GenS1{F0: 0, F1: 1}
                _ = v12
            }
        default:
            v13 := []int{128, -879}
            _ = v13
        }
        v14 := GenS2{F0: 232}
        _ = v14
    }
    v15 := []int{651}
    v16 := GenS3{F0: "hello", F1: 530, F2: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 376
}
func genFn1(p0 int16) int8 {
    fmt.Println(p0)
    p0 += p0
    switch 2147483647 {
    case 2147483647:
        p0 = p0
    case 4294967295:
        v17 := []int{}
        _ = v17
    }
    fmt.Println(p0)
    if false {
        v18 := GenS3{F0: ("foo bar" + "\n\t"), F1: 2, F2: 42}
        v19 := []int{1, 32767, 42, 241}
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := []int{(fi20 - fi20), genFn0(1, ("" + "hello"), int32(p0)), (fi20 & (fi20 + (^(fi20)))), (fi20 + genFn0(-9223372036854775808, "foo bar", 1))}
            var v22 float64 = float64((((1e10 * 1e10) + 0.1) * 1e10))
            v22 = v22
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 uint32 = 1
        v24 := GenS4{F0: 65535, F1: 137, F2: p0}
        _ = v23
        _ = v24
    }
    _ = p0
    return -57
}

func main() {
    var v1 float32 = (0.1 + ((float32(0.1) * (0.0 + 0.1)) / 1e10))
    var v2 int8 = 0
    var v3 int8 = -23
    var v4 string = "a"
    var v5 float32 = v1
    var v6 float64 = 1.5
    var v25 string = (v4 + (v4 + v4))
    for fi26 := 0; fi26 < 10; fi26++ {
        v27 := GenS5{F0: "hello", F1: 42, F2: 1}
        if false {
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 uint16 = 255
                v25 = v4
                var v30 uint64 = 18446744073709551615
                if false {
                    v31 := []int{(fi26 | (fi28 / ((fi26) | 1))), 702, 32767}
                    var v32 uint = 255
                    v32 -= (v32 / (((v32 >> ((65535) & 63))) | 1))
                    _ = v31
                    _ = v32
                } else {
                    var v33 float32 = v1
                    v34 := GenS6{F0: 65535, F1: 0}
                    _ = v33
                    _ = v34
                }
                if false {
                    if ((^(v29)) != v29) {
                        if false {
                            v5 = float32(v6)
                        }
                        var v35 float64 = v6
                        var v36 int16 = 1000
                        _ = v35
                        _ = v36
                    }
                }
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            var v37 int16 = 1
            var v38 uint16 = 250
            switch v25 {
            case "a":
                fmt.Println(v4, v38, v2)
            }
            _ = v37
            _ = v38
        }
        _ = fi26
        _ = v27
    }
    v5 -= 0.1
    v39 := []int{9223372036854775807}
    if len(v39) >= 2 {
        v39 = v39[1:len(v39)]
    }
    for fi40 := 0; fi40 < 9; fi40++ {
        for fi41 := 0; fi41 < 6; fi41++ {
            v42 := GenS7{F0: ((!((((!(0 > 42)) || false) || ((2 != 127) && (209 <= 128))) && true)) && (false && false)), F1: 42}
            switch 216 {
            case 117:
                if (2147483648 > 834) {
                    v6 -= (float64(float32(((((v6 - v6) - 3.14159) / v6) / (v6 - v6)))) / -2.25)
                    fmt.Println(v5, v6, v2)
                    if ((((v1 / ((float32(v6) * v5) * v1)) + 3.14159) + float32((v6 * ((v6 / v6) - v6)))) <= 0.0) {
                        v43 := GenS3{F0: v4, F1: 42, F2: 4294967295}
                        _ = v43
                    } else {
                        v44 := GenS8{F0: 42}
                        v45 := GenS3{F0: v25, F1: 287, F2: 1}
                        var v46 uint32 = 412
                        var v47 uint64 = 1
                        _ = v44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                } else {
                    switch fi40 {
                    case 1:
                        var v48 int8 = (v3 << ((455) & 63))
                        _ = v48
                    case 128:
                        switch (v25 + (v25 + ("foo bar" + (("" + "") + ("a" + (v25 + v4)))))) {
                        case "hello":
                            if len(v39) >= 2 {
                                v39 = v39[1:len(v39)]
                            }
                        default:
                            var v49 uint16 = 0
                            _ = v49
                        }
                    default:
                        _ = len(v39)
                        var v50 uint16 = 1
                        _ = v50
                    }
                }
            case -2147483648:
                var v51 uint64 = 2
                _ = v51
            }
            v52 := GenS6{F0: 0, F1: genFn0(774, (v4 + "hello"), -1)}
            var v53 int32 = -1
            _ = fi41
            _ = v42
            _ = v52
            _ = v53
        }
        var v54 int64 = 473
        v3 = (^(v3))
        _ = fi40
        _ = v54
    }
    var v55 int16 = 959
    if (2147483648 < 4294967295) {
        for fi56 := 0; fi56 < 4; fi56++ {
            var v57 uint = 420
            v55 += int16(v3)
            _ = fi56
            _ = v57
        }
        if false {
            _ = len(v39)
            var v58 int64 = -9223372036854775808
            _ = v58
        }
    } else {
        for fi59 := 0; fi59 < 4; fi59++ {
            v6 = float64(v1)
            _ = fi59
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
