package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint }
func genFn0(p0 uint8, p1 int) uint32 {
    switch (("hello" + "foo bar") + "foo bar") {
    case "a":
        if ((float32(1.5) + 1e10) == 1.5) {
            v5 := GenS1{F0: 2}
            v6 := GenS1{F0: 2}
            _ = v5
            _ = v6
        }
    case "foo bar":
        var v7 uint64 = 4294967295
        _ = v7
    default:
        var v8 uint = 4294967295
        _ = v8
    }
    if false {
        if true {
            v9 := GenS2{F0: 42}
            var v10 uint64 = 4294967295
            p1 += p1
            _ = v10
            _ = v9
        } else {
            if ((^(p0)) == (^(p0))) {
                v11 := []int{-782, p1}
                _ = v11
            } else {
                for fi12 := 0; fi12 < 18; fi12++ {
                    p0 += (p0 & p0)
                    _ = fi12
                }
                p0 = (p0 ^ p0)
            }
        }
    }
    switch "\n\t" {
    case "foo bar":
        p1 += (p1 << ((65535) & 63))
    case "a":
        var v13 float64 = -2.25
        _ = v13
    case "":
        var v14 uint8 = (^(p0))
        _ = v14
    default:
        p0 += (p0 + (p0 * p0))
    }
    p1 -= 42
    _ = p0
    _ = p1
    return 529
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 uint16 = 32767
    if (false || false) {
        v1 -= 176
        v4 = v4
        var v15 int8 = -44
        _ = v15
    }
    if true {
        v16 := GenS3{F0: (100 > -39)}
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        fmt.Println(fi17)
        _ = fi17
    }
    v18 := []int{-2147483648, int(v4), 32767, 128}
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := GenS3{F0: (false && false)}
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    fmt.Println(v4, v4)
    v22 := v3[genFn0(2, 42)]
    switch 2147483647 {
    case -1:
        var v23 float64 = -2.25
        _ = v23
    case -9223372036854775808:
        v4 -= 0
    case 127:
        for fi24 := 0; fi24 < 5; fi24++ {
            switch (fi24 - fi24) {
            case 9223372036854775807:
                v22 = false
            case 127:
                v25 := GenS4{F0: (("" != ((("" + "hello") + ("hello" + "foo bar")) + "hello")) || (0.0 < 0.0)), F1: (v4 >> ((255) & 63))}
                _ = v25
            case -9223372036854775808:
                switch fi24 {
                case -9223372036854775808:
                    v26 := GenS2{F0: 491}
                    _ = v26
                case 736:
                    for fi27 := 0; fi27 < 9; fi27++ {
                        fmt.Println(v4, v4)
                        _ = fi27
                    }
                }
            default:
                v22 = ((^(v1)) == v1)
            }
            _ = fi24
        }
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 int16 = 47
        v30 := GenS5{F0: (^(v1)), F1: 0, F2: 487}
        v1 = 2147483647
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 float64 = 1.5
    for fi32 := 0; fi32 < 12; fi32++ {
        v31 = v31
        v4 += 0
        for fi33 := 0; fi33 < 15; fi33++ {
            if true {
                switch (fi33 | (fi32 >> ((65535) & 63))) {
                case 0:
                    switch (fi32 & 128) {
                    case 2147483647:
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    case 337:
                        var v34 uint = 42
                        _ = v34
                    case 4294967295:
                        var v35 uint64 = 238
                        _ = v35
                    default:
                        _ = len(v3)
                    }
                case 32768:
                    fmt.Println(v4)
                }
                fmt.Println(fi33, v1)
                var v36 uint8 = 138
                _ = v36
            } else {
                v37 := GenS1{F0: (v4 / (((v4 & v4)) | 1))}
                fmt.Println(v31)
                _ = v37
            }
            _ = fi33
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v4
}
