package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 string = "\n\t"
    var v2 int = -9223372036854775808
    var v3 uint8 = 1
    for fi4 := 0; fi4 < 1; fi4++ {
        v3 = v3
        v5 := []int{42}
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 float64 = 1e10
            if false {
                var v8 uint8 = 255
                if true {
                    v2 = (v2 | (-(fi6)))
                    var v9 int64 = 0
                    if true {
                        v10 := GenS1{F0: v1, F1: 2147483647}
                        if len(v5) > 0 {
                            v5[0] = fi4
                        }
                        _ = v10
                    } else {
                        v1 = v1
                        switch ((("hello" + (v1 + (v1 + ((v1 + v1) + "foo bar")))) + v1) + v1) {
                        case "hello":
                            fmt.Println(v7, v8)
                        case "foo bar":
                            if ((1.5 * 0.0) < 1e10) {
                                var v11 uint8 = 255
                                fmt.Println(v8)
                                v12 := GenS1{F0: v1, F1: 167}
                                if (true || (false || (!true))) {
                                    var v13 uint8 = v11
                                    fmt.Println(v3)
                                    v14 := GenS2{F0: (fi6 & (v2 | (v2 ^ fi4))), F1: 460}
                                    var v15 int = 32767
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                } else {
                                    if true {
                                        fmt.Println(v9)
                                        v5 = append(v5, fi6)
                                    } else {
                                        v16 := GenS1{F0: "\n\t", F1: 781}
                                        var v17 int64 = (-(v9))
                                        _ = v16
                                        _ = v17
                                    }
                                    var v18 bool = (false && (false && (-2147483648 == -742)))
                                    switch ("foo bar" + "\n\t") {
                                    case "hello":
                                        var v19 uint64 = 1
                                        v20 := GenS3{F0: 0}
                                        _ = v19
                                        _ = v20
                                    }
                                    switch 4294967295 {
                                    case -856:
                                        var v21 int64 = -9223372036854775808
                                        _ = v21
                                    }
                                    _ = v18
                                }
                                _ = v11
                                _ = v12
                            }
                        case "a":
                            v22 := GenS4{F0: false, F1: -1}
                            _ = v22
                        default:
                            v5 = append(v5, fi6)
                        }
                        var v23 int8 = 0
                        v5 = append(v5, -9223372036854775808)
                        _ = v23
                    }
                    _ = v9
                } else {
                    switch "a" {
                    case "foo bar":
                        v5 = append(v5, (-(fi6)))
                    }
                    switch fi4 {
                    case -9223372036854775808:
                        if (479 >= 0) {
                            v7 = ((1e10 / v7) + ((v7 - v7) + 3.14159))
                            v5 = append(v5, v2)
                        }
                    }
                }
                var v24 int = fi6
                _ = v24
                _ = v8
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v3 -= 128
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    var v25 float32 = 1.5
    switch v2 {
    case 32768:
        switch "foo bar" {
        case "\n\t":
            var v26 uint8 = 255
            _ = v26
        }
    default:
        if (-975 <= 200) {
            var v27 string = v1
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        var v29 int = (v2 % ((v2) | 1))
        v30 := []int{v2, v29, (^(v2))}
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        v32 := []int{127}
        _ = fi31
        _ = v32
    }
    var v33 uint32 = 42
    if (42 >= -2147483648) {
        v34 := []int{}
        if (88 == -85) {
            fmt.Println(v1)
        }
        fmt.Println(v2)
        _ = v34
    }
    var v35 int8 = 127
    for fi36 := 0; fi36 < 7; fi36++ {
        if (1 > 317) {
            v1 = ""
            v25 -= (v25 / v25)
        }
        var v37 uint64 = 0
        _ = fi36
        _ = v37
    }
    var v38 string = (v1 + (("" + v1) + "\n\t"))
    if true {
        fmt.Println(v25)
        v39 := []int{v2, (v2 | (v2 << ((uint(v3)) & 63)))}
        var v40 string = (v1 + v1)
        _ = v39
        _ = v40
    } else {
        if true {
            v41 := GenS4{F0: true, F1: 793}
            _ = v41
        }
    }
    var v42 int16 = 192
    if false {
        fmt.Println(v38, v25)
    } else {
        v43 := GenS5{F0: 18446744073709551615}
        var v44 int32 = 586
        _ = v43
        _ = v44
    }
    fmt.Println(v38, v3, v38)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v35
    _ = v38
    _ = v42
}
