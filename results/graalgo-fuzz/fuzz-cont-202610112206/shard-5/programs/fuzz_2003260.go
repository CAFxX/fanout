package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    var v1 string = "foo bar"
    var v2 float64 = -2.25
    var v3 uint32 = 2147483648
    for fi4 := 0; fi4 < 10; fi4++ {
        if (v1 > ((("\n\t" + ((("foo bar" + "\n\t") + (v1 + v1)) + v1)) + (((("hello" + v1) + (v1 + "a")) + v1) + ((((v1 + (v1 + v1)) + v1) + ("a" + "a")) + v1))) + v1)) {
            v1 = (v1 + "\n\t")
            var v5 bool = (v1 < v1)
            _ = v5
        }
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := []int{}
            fmt.Println(fi6, v2, fi6)
            fmt.Println(v3, v2, fi4)
            _ = fi6
            _ = v7
        }
        _ = fi4
    }
    v2 -= v2
    v8 := []int{32768}
    var v9 int = 32768
    v10 := GenS1{F0: 42, F1: true}
    v11 := GenS1{F0: 2, F1: false}
    v12 := GenS2{F0: ("foo bar" + v1), F1: -1}
    fmt.Println(v9, v2)
    v1 = ("a" + "")
    v9 = v9
    var v13 uint32 = 2
    if (-931 != -32768) {
        v13 -= (v13 - (v13 - (v3 ^ (^(v13)))))
        var v14 uint16 = 150
        if (127 == -34) {
            v8 = append(v8, v9)
            v3 = (v3 ^ v3)
        }
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v9 = (v9 % ((-9223372036854775808) | 1))
        v8 = append(v8, (v9 >> ((18446744073709551615) & 63)))
        _ = fi15
    }
    if false {
        switch v9 {
        case 4294967295:
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        default:
            var v16 int16 = 746
            _ = v16
        }
    } else {
        switch "" {
        case "foo bar":
            if false {
                if (0 != 803) {
                    v13 += v3
                    var v17 float32 = (-2.25 - -2.25)
                    v18 := GenS3{F0: v1, F1: 65535}
                    _ = v17
                    _ = v18
                }
                if ((-800 > 1000) || (!(1 > 65535))) {
                    v2 = v2
                    for fi19 := 0; fi19 < 6; fi19++ {
                        v1 = v1
                        for fi20 := 0; fi20 < 14; fi20++ {
                            fmt.Println(v13, v1)
                            _ = fi20
                        }
                        v1 = "foo bar"
                        _ = fi19
                    }
                }
                if len(v8) > 0 {
                    v8[0] = v9
                }
            } else {
                if true {
                    switch (-(v9)) {
                    case 0:
                        fmt.Println(v2, v9)
                    case -9223372036854775808:
                        var v21 string = v1
                        _ = v21
                    default:
                        fmt.Println(v3)
                    }
                    v13 = v3
                } else {
                    var v22 bool = (-1 < 42)
                    var v23 int64 = -9223372036854775808
                    _ = v22
                    _ = v23
                }
                if true {
                    if true {
                        switch (v1 + v1) {
                        case "foo bar":
                            fmt.Println(v9)
                        }
                    } else {
                        for fi24 := 0; fi24 < 15; fi24++ {
                            fmt.Println(v9, v1, v9)
                            _ = fi24
                        }
                    }
                    v3 -= 456
                } else {
                    fmt.Println(v9)
                    var v25 uint8 = 249
                    _ = v25
                }
            }
        }
        fmt.Println(v13)
        fmt.Println(v3, v3, v3)
    }
    var v26 uint16 = 643
    for fi27 := 0; fi27 < 18; fi27++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        v13 = v13
        _ = fi27
    }
    if (!false) {
        var v28 string = v1
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v8
    _ = v9
}
