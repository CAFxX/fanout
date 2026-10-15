package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int16; F2 string }
type GenS4 struct { F0 uint32; F1 bool }

func main() {
    var v1 bool = false
    var v2 bool = v1
    var v3 int8 = -1
    v3 = v3
    for fi4 := 0; fi4 < 13; fi4++ {
        switch (fi4 & fi4) {
        case -874:
            if (76 < 1) {
                switch "a" {
                case "\n\t":
                    v5 := GenS1{F0: v3, F1: fi4}
                    _ = v5
                case "foo bar":
                    var v6 float64 = 1.5
                    _ = v6
                default:
                    fmt.Println(v3, fi4, v3)
                }
            } else {
                var v7 float64 = (0.0 / 1e10)
                v2 = false
                fmt.Println(fi4, fi4, fi4)
                _ = v7
            }
        }
        if (32767 == int16(fi4)) {
            var v8 string = "foo bar"
            fmt.Println(fi4, v8, fi4)
            v9 := []int{}
            _ = v8
            _ = v9
        } else {
            var v10 int = (fi4 ^ fi4)
            v2 = (!(0 <= -2147483648))
            _ = v10
        }
        _ = fi4
    }
    switch "hello" {
    case "foo bar":
        v11 := GenS2{F0: (((0.0 + (3.14159 + 0.0)) >= 0.0) && (!(0 == 1))), F1: 813}
        _ = v11
    case "a":
        fmt.Println(v2)
    case "\n\t":
        v2 = (v1 && (true || true))
    default:
        if true {
            var v12 uint8 = 42
            if v1 {
                v13 := GenS3{F0: 207, F1: -502, F2: ""}
                for fi14 := 0; fi14 < 8; fi14++ {
                    v15 := GenS2{F0: true, F1: 18446744073709551615}
                    fmt.Println(v1, v2)
                    _ = fi14
                    _ = v15
                }
                switch "a" {
                case "":
                    var v16 float64 = -2.25
                    _ = v16
                case "a":
                    v1 = true
                }
                _ = v13
            } else {
                v1 = (v1 && ((v3 & (v3 << ((18446744073709551615) & 63))) >= (v3 << ((988) & 63))))
                v3 -= (v3 % ((v3) | 1))
                if (!v2) {
                    v1 = v1
                }
            }
            v12 += 128
            _ = v12
        } else {
            v3 = (v3 + (^(v3)))
            switch -1 {
            case 4294967295:
                if true {
                    v2 = (255 >= 0)
                } else {
                    if true {
                        var v17 int32 = 1
                        _ = v17
                    }
                }
            default:
                fmt.Println(v2, v3)
            }
            var v18 bool = ((894 <= 1) && ((v3 << ((755) & 63)) == 1))
            _ = v18
        }
    }
    v19 := []int{9223372036854775807, -9223372036854775808, -54, 118}
    if true {
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := GenS2{F0: v2, F1: 810}
            v1 = true
            v3 = (v3 >> ((4294967295) & 63))
            _ = fi20
            _ = v21
        }
        if (-1 != 9223372036854775807) {
            fmt.Println(v2)
            switch 1 {
            case -405:
                for fi22 := 0; fi22 < 20; fi22++ {
                    v3 += 84
                    fmt.Println(v1, v2)
                    fmt.Println(v2, v1, v1)
                    _ = fi22
                }
            }
        } else {
            if true {
                v1 = false
            } else {
                fmt.Println(v2, v1)
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 uint32 = 738
                    v1 = (1 < 1)
                    for fi25 := 0; fi25 < 17; fi25++ {
                        v19 = append(v19, fi25)
                        v3 -= v3
                        _ = fi25
                    }
                    _ = fi23
                    _ = v24
                }
                if len(v19) > 0 {
                    v19[0] = 2147483647
                }
            }
            v26 := GenS2{F0: v2, F1: 4294967295}
            var v27 int64 = 0
            _ = v26
            _ = v27
        }
    }
    v2 = (!v1)
    v28 := GenS3{F0: 42, F1: 32767, F2: "\n\t"}
    _ = len(v19)
    var v29 int8 = -128
    var v30 int64 = -1
    v31 := GenS4{F0: 1, F1: v2}
    switch "foo bar" {
    case "hello":
        if v2 {
            var v32 int32 = 2147483647
            _ = v32
        } else {
            if (1 > 290) {
                if true {
                    v29 -= v29
                    v1 = (v2 || v2)
                }
            }
        }
    default:
        if len(v19) > 0 {
            v19[0] = 2147483647
        }
    }
    var v33 int = 32767
    var v34 uint32 = 65535
    var v35 uint64 = 121
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v35
}
