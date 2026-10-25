package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint8; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 int64 }

func main() {
    v1 := []int{42}
    var v2 int8 = 0
    v3 := []int{9223372036854775807, 524, 128, 32767}
    var v4 int64 = -589
    v5 := make(map[string]string)
    v4 = 0
    v2 = (v2 << ((17) & 63))
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := GenS1{F0: true, F1: 1}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS2{F0: 611, F1: 1}
    v3 = append(v3, -2147483648)
    switch 32767 {
    case 32767:
        for fi10 := 0; fi10 < 7; fi10++ {
            if true {
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 int = (fi10 & fi10)
                    v2 -= v2
                    _ = fi11
                    _ = v12
                }
            }
            _ = fi10
        }
    case 1:
        v2 -= v2
    case -1:
        if false {
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := GenS3{F0: v2, F1: true}
                var v15 int = fi13
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            if (true && true) {
                var v16 uint = 255
                for fi17 := 0; fi17 < 3; fi17++ {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        fmt.Println(fi18, v2)
                        v19 := GenS4{F0: 163, F1: v2, F2: 0}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi17
                }
                _ = v16
            } else {
                v2 += (v2 | v2)
            }
        }
    default:
        v5[("" + "")] = ("" + ("foo bar" + "a"))
    }
    switch "a" {
    case "a":
        if false {
            v20 := GenS5{F0: 65535}
            v21 := GenS4{F0: 127, F1: v2, F2: 32767}
            var v22 bool = ((v4 * (v4 & v4)) <= v4)
            _ = v20
            _ = v21
            _ = v22
        } else {
            if len(v3) > 0 {
                v3[0] = 32768
            }
            _ = len(v5)
            switch ("" + ("a" + "a")) {
            case "foo bar":
                v4 = v4
            case "":
                v2 = v2
            case "hello":
                fmt.Println(v4)
            }
        }
    default:
        v2 = v2
    }
    if true {
        v23 := GenS4{F0: 209, F1: v2, F2: 645}
        _ = v23
    } else {
        v4 = v4
        fmt.Println(v4, v2)
    }
    fmt.Println(v4, v2)
    var v24 string = "foo bar"
    var v25 bool = true
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 uint64 = 2
        var v28 int8 = (v2 - 127)
        fmt.Println(v4, v27)
        _ = fi26
        _ = v27
        _ = v28
    }
    if (false && (!(v4 != -771))) {
        if (2 < 2) {
            if true {
                v29 := GenS3{F0: (v2 << ((18446744073709551615) & 63)), F1: (!true)}
                var v30 uint = 42
                _ = v29
                _ = v30
            } else {
                var v31 float64 = ((-2.25 / -2.25) + (((1.5 + 1e10) * -2.25) / 1.0))
                v32 := v5[v24]
                var v33 uint64 = 2
                _ = v31
                _ = v32
                _ = v33
            }
            switch 42 {
            case 290:
                for fi34 := 0; fi34 < 9; fi34++ {
                    v35 := GenS6{F0: (^(fi34))}
                    v4 = v4
                    for fi36 := 0; fi36 < 16; fi36++ {
                        var v37 int64 = (-(v4))
                        _ = fi36
                        _ = v37
                    }
                    _ = fi34
                    _ = v35
                }
            case -338:
                if v25 {
                    var v38 string = v24
                    _ = v38
                } else {
                    var v39 string = "hello"
                    fmt.Println(v24, v2)
                    _ = v39
                }
            case 888:
                v40 := GenS7{F0: -9223372036854775808, F1: (^(v4))}
                _ = v40
            default:
                var v41 bool = (!v25)
                _ = v41
            }
        } else {
            var v42 int8 = (v2 % ((v2) | 1))
            _ = v42
        }
    } else {
        var v43 int = 32768
        v44 := v5["\n\t"]
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        var v46 float32 = 3.14159
        _ = len(v5)
        for fi47 := 0; fi47 < 19; fi47++ {
            if (v25 && (-499 <= 1)) {
                v25 = v25
            }
            fmt.Println(fi45, fi45, fi45)
            var v48 uint16 = 17
            _ = fi47
            _ = v48
        }
        _ = fi45
        _ = v46
    }
    v1 = append(v1, 1)
    fmt.Println(v25, v25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
