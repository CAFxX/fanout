package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 int; F1 uint16; F2 bool }
func genFn0() bool {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{(fi6 / (((fi6 | (fi6 - 1))) | 1))}
        fmt.Println(fi6)
        _ = fi6
        _ = v7
    }
    if false {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: (fi8 % (((fi8 / ((9223372036854775807) | 1))) | 1)), F1: -1}
            fmt.Println(fi8)
            _ = fi8
            _ = v9
        }
        var v10 string = "hello"
        var v11 string = ((v10 + (v10 + (v10 + v10))) + v10)
        _ = v10
        _ = v11
    }
    return (-9223372036854775808 != 127)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"foo bar", "a", "hello"}
    var v3 uint64 = 1
    var v4 int8 = 127
    var v5 float64 = 1e10
    if true {
        _ = len(v1)
        v3 += (^(v3))
    } else {
        switch 2147483647 {
        case 32767:
            if (4294967295 < 42) {
                _ = len(v2)
            } else {
                if (-1 != 9223372036854775807) {
                    var v12 int32 = 2147483647
                    _ = v12
                }
                var v13 uint16 = 65535
                _ = v13
            }
        case 9223372036854775807:
            switch ("" + "\n\t") {
            case "foo bar":
                v14 := GenS1{F0: 0, F1: v4}
                _ = v14
            case "\n\t":
                if true {
                    if true {
                        var v15 uint8 = 105
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        for fi16 := 0; fi16 < 14; fi16++ {
                            v1[2147483648] = 2
                            v17 := GenS2{F0: 255, F1: (!false)}
                            fmt.Println(v5)
                            _ = fi16
                            _ = v17
                        }
                        _ = v15
                    }
                    if false {
                        if true {
                            v5 += -2.25
                            v18 := GenS2{F0: 65535, F1: (false && (-1 != 0))}
                            _ = v18
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                v4 -= -105
                            case "hello":
                                v3 = 0
                            }
                        }
                        _ = len(v1)
                        v3 = (v3 | 42)
                    } else {
                        switch 2147483647 {
                        case 1:
                            var v19 int32 = -1
                            _ = v19
                        case -625:
                            fmt.Println(v5, v5, v3)
                        case -2147483648:
                            fmt.Println(v3)
                        }
                        fmt.Println(v3, v3)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
            }
        }
    }
    var v20 uint32 = 2147483648
    v20 -= v20
    var v21 uint64 = v3
    fmt.Println(v21)
    v22 := GenS3{F0: 42, F1: -2147483648, F2: -98}
    if true {
        v20 = 42
        var v23 string = (("hello" + "") + "\n\t")
        _ = v23
    } else {
        v5 = v5
        v20 += (v20 * (v20 % (((^(v20))) | 1)))
        switch 2147483647 {
        case 2147483647:
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 uint = 123
                v26 := GenS4{F0: "a", F1: 1}
                var v27 int = fi24
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        case -641:
            var v28 uint16 = 0
            _ = v28
        }
    }
    if true {
        var v29 float64 = ((v5 * (1.5 * float64((0.0 * 0.0)))) / v5)
        if ((0 <= 556) && ((!false) || ((^(v21)) == (^(v21))))) {
            v30 := GenS5{F0: 9223372036854775807, F1: 733, F2: (false && true)}
            _ = v30
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 int16 = 32767
            _ = fi31
            _ = v32
        }
        _ = v29
    }
    var v33 uint64 = v3
    v5 += ((v5 + v5) * v5)
    for fi34 := 0; fi34 < 10; fi34++ {
        switch (fi34 & 128) {
        case -2147483648:
            var v35 int = (fi34 % (((-(fi34))) | 1))
            _ = v35
        default:
            var v36 uint16 = 182
            _ = v36
        }
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
