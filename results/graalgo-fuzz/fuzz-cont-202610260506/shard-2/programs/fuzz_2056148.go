package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS4 struct { F0 int; F1 int8 }
func genFn0(p0 string, p1 uint) uint {
    switch (p0 + "a") {
    case "\n\t":
        if true {
            fmt.Println(p1)
            var v6 uint32 = 2147483647
            _ = v6
        }
    case "":
        switch 0 {
        case 0:
            var v7 uint = 2
            _ = v7
        }
    case "foo bar":
        if ((1 != 74) && (true && (48 <= 107))) {
            switch p0 {
            case "":
                fmt.Println(p0, p0)
            case "foo bar":
                p1 += 255
            default:
                var v8 int64 = 229
                _ = v8
            }
        }
    }
    p0 = "hello"
    var v9 int16 = 1000
    v10 := []int{4294967295, 4294967295, 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (p1 ^ (p1 + p1))
}
func genFn1(p0 uint64, p1 bool, p2 uint16) bool {
    switch 4294967295 {
    case -2147483648:
        p2 = (p2 >> ((65535) & 63))
    case -1:
        var v11 uint64 = (^(p0))
        _ = v11
    default:
        if (p2 < p2) {
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := []int{0, (fi12 & fi12)}
                _ = fi12
                _ = v13
            }
            p2 = (^(p2))
            fmt.Println(p1, p0, p2)
        } else {
            if (0 >= 128) {
                p1 = (-1 < -1)
            } else {
                p2 += p2
            }
        }
    }
    p1 = (69 != 546)
    fmt.Println(p0, p2)
    fmt.Println(p0, p2, p0)
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return (p1 && (p0 == 42))
}

func main() {
    v1 := []uint32{2147483647}
    var v2 int = -2147483648
    var v3 int16 = 1
    v4 := make(map[uint32]bool)
    var v5 uint16 = 42
    var v14 int8 = 45
    v14 = v14
    delete(v4, 0)
    if false {
        v15 := v4[841]
        _ = v15
    }
    if ((872 != v2) && false) {
        if true {
            var v16 int64 = -9223372036854775808
            v5 = v5
            _ = v16
        }
        v17 := GenS1{F0: v3, F1: 255, F2: 0}
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS2{F0: 65535}
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS2{F0: 0}
            _ = fi20
            _ = v21
        }
        if len(v1) > 0 {
            v1[0] = 536
        }
        _ = fi18
        _ = v19
    }
    if (725 <= 0) {
        fmt.Println(v3, v3)
    } else {
        for fi22 := 0; fi22 < 3; fi22++ {
            switch "a" {
            case "foo bar":
                v2 = v2
            case "\n\t":
                _ = len(v4)
            case "a":
                if false {
                    var v23 uint = 42
                    _ = v23
                }
            default:
                fmt.Println(v5)
            }
            if (int16(v14) < v3) {
                v2 += -595
                var v24 int = (fi22 + 32768)
                _ = v24
            } else {
                v25 := GenS1{F0: 32767, F1: 128, F2: 65535}
                if (genFn0(("\n\t" + ("hello" + ("foo bar" + (("\n\t" + (("a" + "a") + "\n\t")) + "")))), genFn0((("\n\t" + "hello") + ("\n\t" + "")), 42)) > 0) {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        _ = len(v4)
                        v27 := GenS3{F0: 2147483647, F1: -1, F2: 127}
                        fmt.Println(v5)
                        v2 -= (fi26 & 42)
                        _ = fi26
                        _ = v27
                    }
                    fmt.Println(fi22, v14, v2)
                } else {
                    fmt.Println(v3)
                    v28 := GenS3{F0: 2147483648, F1: 42, F2: 128}
                    _ = v28
                }
                _ = v25
            }
            delete(v4, 536)
            _ = fi22
        }
        var v29 int = v2
        v30 := GenS4{F0: 4294967295, F1: 127}
        _ = v29
        _ = v30
    }
    fmt.Println(v2)
    v2 += v2
    switch "a" {
    case "foo bar":
        var v31 int64 = 608
        _ = v31
    default:
        v2 -= (v2 & 32768)
    }
    if true {
        _ = len(v1)
    } else {
        v32 := GenS3{F0: 297, F1: v14, F2: 0}
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
