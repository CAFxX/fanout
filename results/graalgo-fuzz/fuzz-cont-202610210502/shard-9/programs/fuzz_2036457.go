package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint8; F2 string }
type GenS5 struct { F0 uint32; F1 bool; F2 int64 }

func main() {
    var v1 uint64 = 7
    var v2 uint16 = 98
    v3 := make(map[string]string)
    var v4 int32 = -842
    var v5 uint32 = 1
    v6 := make(map[int]bool)
    var v7 uint = 0
    for fi8 := 0; fi8 < 5; fi8++ {
        v1 -= (^(v1))
        v9 := []int{2147483647}
        _ = fi8
        _ = v9
    }
    var v10 int = 1
    switch -9223372036854775808 {
    case 127:
        switch "" {
        case "":
            switch (v10 + (v10 / (((-(v10))) | 1))) {
            case 0:
                var v11 int8 = 92
                _ = v11
            }
        case "\n\t":
            v1 = v1
        }
    case 42:
        v12 := GenS1{F0: v2}
        _ = v12
    default:
        var v13 uint16 = uint16(v5)
        _ = v13
    }
    switch "foo bar" {
    case "":
        v5 += 0
    case "a":
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint64 = 18446744073709551615
            _ = fi14
            _ = v15
        }
    }
    switch "hello" {
    case "foo bar":
        fmt.Println(v5, v5)
    case "hello":
        var v16 int16 = -8
        _ = v16
    case "a":
        v7 -= 255
    }
    var v17 int64 = 42
    if ((!false) && true) {
        var v18 int64 = 9223372036854775807
        _ = v18
    }
    if (v17 == -1) {
        v7 -= v7
        fmt.Println(v2, v5, v10)
        for fi19 := 0; fi19 < 1; fi19++ {
            switch fi19 {
            case 553:
                switch fi19 {
                case 2147483647:
                    var v20 uint32 = 65535
                    _ = v20
                case -9223372036854775808:
                    switch "foo bar" {
                    case "":
                        fmt.Println(v10, v4, v2)
                    case "foo bar":
                        for fi21 := 0; fi21 < 19; fi21++ {
                            var v22 uint32 = v5
                            switch 353 {
                            case 127:
                                switch fi21 {
                                case 2147483647:
                                    v17 = v17
                                case -821:
                                    var v23 int32 = (-(v4))
                                    _ = v23
                                }
                            default:
                                if true {
                                    v24 := GenS1{F0: 2}
                                    _ = v24
                                } else {
                                    _ = len(v6)
                                    var v25 int8 = -15
                                    v26 := []int{}
                                    _ = v25
                                    _ = v26
                                }
                            }
                            _ = len(v6)
                            _ = fi21
                            _ = v22
                        }
                    }
                case 32767:
                    if ((v7 & v7) != v7) {
                        var v27 uint = (v7 / (((v7 << ((v7) & 63))) | 1))
                        _ = v27
                    }
                default:
                    var v28 uint64 = 4294967295
                    _ = v28
                }
            case -1:
                v4 -= (v4 % ((v4) | 1))
            default:
                v1 = (v1 / ((v1) | 1))
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                v30 := v3["a"]
                v31 := GenS1{F0: (v2 << (((^(v7))) & 63))}
                var v32 int8 = 74
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            fmt.Println(fi19)
            _ = fi19
        }
    } else {
        v33 := GenS2{F0: (v1 * (v1 % (((v1 / (((v1 / ((v1) | 1))) | 1))) | 1))), F1: "foo bar", F2: 100}
        _ = v33
    }
    if false {
        switch "" {
        case "foo bar":
            v2 -= (v2 >> (((v7 << (((v7 << ((800) & 63))) & 63))) & 63))
        case "\n\t":
            fmt.Println(v2, v1)
        }
        v34 := GenS3{F0: v7}
        _ = v34
    } else {
        fmt.Println(v2, v2, v17)
        switch "\n\t" {
        case "foo bar":
            if (!true) {
                v2 += (v2 << ((v7) & 63))
                delete(v6, (v10 % (((v10 << (((v7 * (v7 % ((v7) | 1)))) & 63))) | 1)))
            } else {
                v35 := []int{(v10 + (-(v10)))}
                var v36 int = v10
                for fi37 := 0; fi37 < 12; fi37++ {
                    var v38 uint8 = 127
                    for fi39 := 0; fi39 < 10; fi39++ {
                        v40 := v3[(("hello" + ("" + "foo bar")) + "")]
                        _ = fi39
                        _ = v40
                    }
                    fmt.Println(v5, v7)
                    _ = fi37
                    _ = v38
                }
                _ = v35
                _ = v36
            }
        case "\n\t":
            var v41 uint8 = 128
            _ = v41
        case "a":
            switch "\n\t" {
            case "hello":
                fmt.Println(v10, v2, v5)
            case "a":
                switch "\n\t" {
                case "":
                    var v42 int64 = (v17 ^ int64(v4))
                    _ = v42
                }
            case "":
                v4 = 1
            }
        }
    }
    v43 := GenS1{F0: (v2 - v2)}
    for fi44 := 0; fi44 < 2; fi44++ {
        fmt.Println(v1, v7)
        _ = fi44
    }
    if ((v17 | int64(v1)) > v17) {
        var v45 uint64 = (v1 << (((^(v7))) & 63))
        switch (v10 * v10) {
        case 4294967295:
            v46 := GenS4{F0: 0, F1: 255, F2: "a"}
            _ = v46
        case -9223372036854775808:
            v3["\n\t"] = "foo bar"
        }
        v47 := GenS5{F0: 4294967295, F1: (255 < (v7 | 350)), F2: (v17 ^ v17)}
        _ = v45
        _ = v47
    } else {
        v48 := GenS3{F0: (v7 ^ 2)}
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
