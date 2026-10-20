package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    var v4 bool = true
    var v5 uint8 = 127
    v6 := make(map[int]string)
    var v7 uint32 = 2147483648
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := v1[fi8]
        if (!(v4 || (fi8 != fi8))) {
            v10 := v2[732]
            _ = v10
        } else {
            if true {
                fmt.Println(v7)
            }
            if v4 {
                v11 := GenS1{F0: 9223372036854775807}
                switch int(v3) {
                case 945:
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v9 = ""
                        _ = fi12
                    }
                case 642:
                    if v4 {
                        v13 := GenS1{F0: (fi8 + (fi8 + (^(fi8))))}
                        switch v9 {
                        case "hello":
                            v5 = (v5 | uint8(v7))
                        case "foo bar":
                            var v14 uint8 = 154
                            _ = v14
                        case "":
                            v15 := GenS1{F0: 127}
                            _ = v15
                        default:
                            var v16 int32 = 1
                            _ = v16
                        }
                        _ = v13
                    }
                case -9223372036854775808:
                    var v17 bool = ((v3 % (((v3 % ((v3) | 1))) | 1)) == -1)
                    _ = v17
                default:
                    var v18 bool = true
                    _ = v18
                }
                _ = v11
            } else {
                v19 := GenS1{F0: -1}
                if v4 {
                    if (((3.14159 - ((0.1 * 0.0) - 3.14159)) / 0.1) > 0.0) {
                        for fi20 := 0; fi20 < 15; fi20++ {
                            var v21 uint64 = 0
                            switch "foo bar" {
                            case "a":
                                var v22 float32 = 1.5
                                _ = v22
                            case "hello":
                                v23 := []int{}
                                _ = v23
                            }
                            switch ("" + ("a" + v9)) {
                            case "":
                                v24 := GenS1{F0: 1}
                                _ = v24
                            case "hello":
                                v25 := []int{(fi8 >> ((368) & 63)), -1}
                                _ = v25
                            case "\n\t":
                                var v26 int32 = 0
                                _ = v26
                            default:
                                var v27 uint16 = 0
                                _ = v27
                            }
                            _ = fi20
                            _ = v21
                        }
                        var v28 uint8 = (v5 - (v5 % ((139) | 1)))
                        _ = v28
                    }
                } else {
                    switch fi8 {
                    case -439:
                        switch v9 {
                        case "\n\t":
                            v29 := GenS1{F0: fi8}
                            _ = v29
                        case "a":
                            v30 := GenS1{F0: -81}
                            _ = v30
                        default:
                            for fi31 := 0; fi31 < 1; fi31++ {
                                v4 = (v4 && true)
                                v32 := []int{}
                                _ = fi31
                                _ = v32
                            }
                        }
                    case -9223372036854775808:
                        v4 = v4
                    default:
                        v9 = "\n\t"
                    }
                    switch ((((((v9 + v9) + v9) + ("" + v9)) + v9) + v9) + "\n\t") {
                    case "\n\t":
                        fmt.Println(v7)
                    case "":
                        if v4 {
                            v33 := []int{}
                            switch fi8 {
                            case 0:
                                fmt.Println(v5, v5, v5)
                            case 4294967295:
                                v3 = 138
                            default:
                                fmt.Println(fi8, v3, v9)
                            }
                            for fi34 := 0; fi34 < 7; fi34++ {
                                var v35 uint64 = 42
                                delete(v6, fi34)
                                _ = fi34
                                _ = v35
                            }
                            _ = v33
                        }
                    case "hello":
                        var v36 uint16 = 42
                        _ = v36
                    }
                }
                _ = v19
            }
        }
        _ = fi8
        _ = v9
    }
    v37 := GenS1{F0: 0}
    v4 = true
    var v38 int64 = (v3 * (v3 & v3))
    if (242 == 971) {
        v39 := GenS2{F0: 0, F1: 2}
        if (2 != 7) {
            fmt.Println(v5, v7, v5)
            var v40 uint16 = 0
            _ = v40
        }
        v41 := GenS1{F0: 42}
        _ = v39
        _ = v41
    } else {
        for fi42 := 0; fi42 < 18; fi42++ {
            fmt.Println(v38, v5, fi42)
            for fi43 := 0; fi43 < 18; fi43++ {
                var v44 int8 = 127
                v4 = v4
                _ = fi43
                _ = v44
            }
            v45 := []int{fi42, fi42, (fi42 >> ((4294967295) & 63)), 128}
            _ = fi42
            _ = v45
        }
    }
    fmt.Println(v7, v4)
    var v46 bool = v4
    for fi47 := 0; fi47 < 10; fi47++ {
        var v48 uint16 = 42
        _ = fi47
        _ = v48
    }
    v46 = v4
    var v49 float32 = 3.14159
    v50 := []int{-461, -9223372036854775808, -132}
    v6[-373] = "\n\t"
    var v51 uint32 = (v7 << ((255) & 63))
    for fi52 := 0; fi52 < 18; fi52++ {
        var v53 int16 = -379
        fmt.Println(v4)
        v7 -= v51
        _ = fi52
        _ = v53
    }
    v54 := GenS3{F0: -2147483648}
    v51 += v51
    var v55 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v46
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v54
    _ = v55
    _ = v6
    _ = v7
}
