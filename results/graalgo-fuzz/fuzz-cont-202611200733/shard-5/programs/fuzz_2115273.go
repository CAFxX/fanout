package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint64; F2 uint }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int16; F1 int32 }
func genFn0(p0 int32, p1 uint16, p2 uint32) bool {
    p0 = (p0 / (((^(p0))) | 1))
    v4 := []int{1, -2147483648, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (0 > p2)
}
func genFn1(p0 int) uint32 {
    p0 -= 9223372036854775807
    v5 := GenS1{F0: 859}
    var v6 int16 = 42
    v7 := GenS1{F0: 32767}
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 767
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 int32 = -1
    switch 2147483647 {
    case -9223372036854775808:
        v9 := []int{4294967295, 4294967295}
        _ = v9
    case -1:
        v3 -= v3
    case -947:
        switch 324 {
        case 0:
            v10 := GenS1{F0: 957}
            _ = v10
        }
    default:
        var v11 uint16 = 666
        _ = v11
    }
    switch 2147483647 {
    case -76:
        switch ("hello" + "a") {
        case "foo bar":
            v3 = (v3 ^ v3)
        case "\n\t":
            if true {
                var v12 string = ("a" + "\n\t")
                var v13 bool = (false || false)
                v14 := GenS2{F0: 396}
                _ = v12
                _ = v13
                _ = v14
            } else {
                var v15 uint64 = 2
                if false {
                    var v16 int8 = -128
                    switch -2147483648 {
                    case 9223372036854775807:
                        fmt.Println(v16)
                    case -9223372036854775808:
                        for fi17 := 0; fi17 < 9; fi17++ {
                            switch "a" {
                            case "hello":
                                v16 = (v16 << ((351) & 63))
                            case "a":
                                v18 := GenS3{F0: 2}
                                _ = v18
                            default:
                                var v19 uint16 = 255
                                _ = v19
                            }
                            for fi20 := 0; fi20 < 16; fi20++ {
                                fmt.Println(fi20, v3)
                                _ = fi20
                            }
                            for fi21 := 0; fi21 < 20; fi21++ {
                                var v22 int = (fi21 | fi17)
                                if true {
                                    if (false || (2 == genFn1((fi21 * 0)))) {
                                        if true {
                                            fmt.Println(fi17, v15)
                                        } else {
                                            var v23 int = fi21
                                            _ = v23
                                        }
                                        v24 := GenS4{F0: (fi21 ^ 32768), F1: (v15 * 0), F2: 946}
                                        v25 := v1[v22]
                                        _ = v24
                                        _ = v25
                                    } else {
                                        if (-32768 < 826) {
                                            v26 := []int{9223372036854775807, v22, 42}
                                            fmt.Println(v22)
                                            _ = v26
                                        }
                                        v27 := GenS3{F0: 0}
                                        _ = v27
                                    }
                                    var v28 uint16 = 65535
                                    delete(v2, genFn1((fi21 - fi21)))
                                    _ = v28
                                }
                                if ((^(v3)) == (v3 << ((42) & 63))) {
                                    fmt.Println(fi17)
                                } else {
                                    if false {
                                        v3 = (v3 & 516)
                                    }
                                    v29 := GenS5{F0: true, F1: 1}
                                    _ = v29
                                }
                                _ = fi21
                                _ = v22
                            }
                            _ = fi17
                        }
                    case 128:
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            var v30 uint32 = uint32(v15)
                            _ = v30
                        case "":
                            var v31 uint = 416
                            _ = v31
                        default:
                            fmt.Println(v15, v15)
                        }
                    default:
                        v15 = 4294967295
                    }
                    for fi32 := 0; fi32 < 5; fi32++ {
                        var v33 uint = 437
                        v16 = v16
                        _ = fi32
                        _ = v33
                    }
                    _ = v16
                }
                _ = v15
            }
        case "":
            v34 := GenS1{F0: 1}
            _ = v34
        default:
            v3 = (v3 % ((v3) | 1))
        }
    default:
        var v35 int64 = 0
        _ = v35
    }
    v36 := GenS6{F0: 438, F1: (v3 & (v3 << ((255) & 63)))}
    var v37 float64 = -2.25
    v37 = v37
    var v38 uint32 = 65535
    var v39 float64 = 1.5
    var v40 int = 165
    switch "hello" {
    case "":
        for fi41 := 0; fi41 < 1; fi41++ {
            var v42 int64 = 9223372036854775807
            for fi43 := 0; fi43 < 18; fi43++ {
                v44 := []int{fi43, (fi43 - fi41), fi43}
                if (!false) {
                    v44 = append(v44, (^(fi43)))
                    v39 = ((0.1 * ((v37 / -2.25) + ((0.1 * v37) / float64(0.1)))) / (((v39 + v37) - (-2.25 * v39)) / 1.0))
                    fmt.Println(fi41, v39)
                } else {
                    var v45 uint = 65535
                    _ = v45
                }
                var v46 uint16 = 437
                _ = fi43
                _ = v44
                _ = v46
            }
            _ = fi41
            _ = v42
        }
    case "\n\t":
        v38 = v38
    case "foo bar":
        var v47 bool = true
        _ = v47
    default:
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
