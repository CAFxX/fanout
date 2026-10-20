package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 uint8 }

func main() {
    var v1 float64 = (3.14159 - (0.1 / 3.14159))
    var v2 uint32 = 234
    var v3 uint32 = v2
    var v4 int64 = 1
    var v5 float64 = ((((v1 + v1) - (3.14159 + (v1 + v1))) - v1) / (3.14159 - ((0.1 - (v1 + (v1 / (v1 + (float64(-2.25) + (v1 * ((v1 + v1) * 0.1))))))) / (v1 - v1))))
    fmt.Println(v5, v5, v2)
    if (v5 < ((1e10 * (float64(1.5) - v1)) * 1.5)) {
        var v6 uint8 = 127
        v7 := []int{241, 9223372036854775807, 2147483647}
        _ = v6
        _ = v7
    } else {
        v8 := []int{877, 854}
        _ = v8
    }
    var v9 int = -306
    switch -2147483648 {
    case -264:
        for fi10 := 0; fi10 < 12; fi10++ {
            fmt.Println(v3)
            v5 = (v1 / (v1 / v5))
            _ = fi10
        }
    case -447:
        fmt.Println(v5, v4, v9)
    }
    var v11 uint32 = (^(v2))
    fmt.Println(v9, v4, v9)
    if ((("\n\t" + ("" + "a")) + ("hello" + "\n\t")) >= ("\n\t" + (("" + "") + "hello"))) {
        if false {
            v12 := []int{127, v9}
            _ = v12
        } else {
            v13 := GenS1{F0: 1, F1: 1}
            if true {
                v14 := GenS1{F0: 121, F1: 255}
                _ = v14
            } else {
                v15 := []int{v9}
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = v15
            }
            _ = v13
        }
        v16 := []int{v9, (v9 & v9)}
        v17 := GenS2{F0: 65535, F1: 2147483647}
        _ = v16
        _ = v17
    }
    switch (("" + (("foo bar" + "foo bar") + ((("" + "\n\t") + (("foo bar" + ("a" + "")) + "hello")) + ((("\n\t" + "") + "\n\t") + "foo bar")))) + "foo bar") {
    case "foo bar":
        if false {
            v18 := []int{360, v9, 9223372036854775807, 2147483647}
            fmt.Println(v5)
            v18 = append(v18, v9)
            _ = v18
        }
    case "\n\t":
        v1 = v5
    default:
        v5 -= 3.14159
    }
    var v19 uint32 = 2147483647
    var v20 uint16 = 239
    v1 += float64(1.5)
    for fi21 := 0; fi21 < 6; fi21++ {
        switch (v9 % ((v9) | 1)) {
        case 2147483647:
            v3 += v19
        case 42:
            fmt.Println(v2)
        case -1:
            v22 := []int{}
            _ = v22
        default:
            if (true && ((v4 - -1) >= v4)) {
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    if (!((fi21 >> ((630) & 63)) < (fi21 ^ -9223372036854775808))) {
                        for fi23 := 0; fi23 < 11; fi23++ {
                            v24 := GenS3{F0: -1, F1: uint(v20), F2: 2147483648}
                            v25 := GenS4{F0: (v20 ^ 42)}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    } else {
                        var v26 uint8 = 127
                        v27 := GenS5{F0: fi21}
                        var v28 int32 = 0
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                }
                for fi29 := 0; fi29 < 19; fi29++ {
                    v1 += 1.5
                    v30 := GenS6{F0: uint8(v20)}
                    fmt.Println(v5)
                    if (!(true && (false && true))) {
                        if (!true) {
                            v9 += (v9 << ((65535) & 63))
                        }
                        if (!("" != "")) {
                            if (!true) {
                                v31 := GenS1{F0: 1, F1: 0}
                                _ = v31
                            }
                        }
                        v32 := []int{-2147483648, v9, (v9 & (fi21 ^ fi21)), -2147483648}
                        var v33 uint32 = v19
                        switch v9 {
                        case 42:
                            fmt.Println(v4, v9, v20)
                            v34 := GenS7{F0: "foo bar"}
                            _ = v34
                        case 9223372036854775807:
                            v35 := GenS8{F0: true}
                            _ = v35
                        case 1:
                            switch "foo bar" {
                            case "a":
                                var v36 int32 = 1
                                _ = v36
                            }
                        default:
                            switch "foo bar" {
                            case "":
                                v20 -= (v20 & (v20 << ((2) & 63)))
                            case "\n\t":
                                _ = len(v32)
                            }
                        }
                        v2 += v11
                        _ = v32
                        _ = v33
                    } else {
                        v11 = v3
                        var v37 uint32 = 957
                        _ = v37
                    }
                    v9 = (v9 << ((255) & 63))
                    _ = fi29
                    _ = v30
                }
                if false {
                    var v38 float64 = v1
                    v20 = (v20 + v20)
                    fmt.Println(v5, v3)
                    _ = v38
                } else {
                    v39 := GenS9{F0: 18446744073709551615, F1: 42}
                    for fi40 := 0; fi40 < 3; fi40++ {
                        var v41 int = v9
                        if (!(937 != 156)) {
                            v42 := GenS6{F0: 0}
                            _ = v42
                        } else {
                            v2 -= 1
                        }
                        _ = fi40
                        _ = v41
                    }
                    _ = v39
                }
            }
        }
        v43 := []int{(-(fi21))}
        var v44 uint = 308
        _ = fi21
        _ = v43
        _ = v44
    }
    v45 := []int{1, v9}
    var v46 int16 = 1000
    var v47 float64 = (v1 / v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v9
}
