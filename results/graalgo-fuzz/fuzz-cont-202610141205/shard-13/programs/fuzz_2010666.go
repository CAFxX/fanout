package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int8; F2 bool }
type GenS6 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{-318, -2147483648, 127}
    v3 := make(map[string]uint32)
    var v4 int32 = 2147483647
    var v5 int16 = 0
    var v6 string = "foo bar"
    v7 := GenS1{F0: (v6 + ("foo bar" + v6)), F1: (v6 + v6), F2: v6}
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 bool = (2 >= 370)
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 int16 = 0
            if (!(v9 && (fi10 == 197))) {
                _ = len(v2)
                v4 -= (v4 * v4)
                var v12 int8 = 76
                _ = v12
            } else {
                fmt.Println(v11, v11)
                fmt.Println(v5, fi8, v4)
                switch "\n\t" {
                case "hello":
                    v11 = (^(v5))
                default:
                    var v13 uint = 1
                    v5 -= v5
                    _ = v13
                }
            }
            _ = fi10
            _ = v11
        }
        var v14 string = ((v6 + ((v6 + "a") + (v6 + ("a" + (v6 + v6))))) + "foo bar")
        _ = fi8
        _ = v14
        _ = v9
    }
    if ((!false) && (42 >= -317)) {
        fmt.Println(v4, v5)
        v15 := GenS2{F0: (false || true)}
        _ = v15
    }
    var v16 uint64 = 2
    for fi17 := 0; fi17 < 2; fi17++ {
        if false {
            var v18 uint32 = 2147483648
            var v19 int = fi17
            var v20 string = v6
            _ = v18
            _ = v19
            _ = v20
        } else {
            if false {
                v21 := GenS1{F0: "a", F1: "\n\t", F2: v6}
                v5 += v5
                switch "a" {
                case "\n\t":
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v1[(^(fi17))] = (^(fi22))
                        var v23 float64 = 0.1
                        var v24 int8 = 105
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                case "a":
                    for fi25 := 0; fi25 < 20; fi25++ {
                        fmt.Println(v4, fi25)
                        v26 := GenS1{F0: v6, F1: v6, F2: (v6 + v6)}
                        _ = fi25
                        _ = v26
                    }
                case "foo bar":
                    fmt.Println(v4, fi17, fi17)
                default:
                    v3[("\n\t" + v6)] = 65535
                }
                _ = v21
            } else {
                v27 := GenS2{F0: true}
                v28 := GenS3{F0: 1}
                if true {
                    if len(v2) > 0 {
                        v2[0] = fi17
                    }
                    var v29 int16 = (v5 ^ 544)
                    v5 += v29
                    _ = v29
                }
                _ = v27
                _ = v28
            }
            v16 += (v16 - (v16 * v16))
        }
        v30 := GenS3{F0: 637}
        _ = fi17
        _ = v30
    }
    v31 := GenS3{F0: 2147483647}
    switch "foo bar" {
    case "hello":
        for fi32 := 0; fi32 < 7; fi32++ {
            for fi33 := 0; fi33 < 13; fi33++ {
                if len(v2) > 0 {
                    v2[0] = (fi33 - fi33)
                }
                var v34 int32 = 2147483647
                var v35 int64 = 42
                _ = fi33
                _ = v34
                _ = v35
            }
            for fi36 := 0; fi36 < 18; fi36++ {
                fmt.Println(v6, v6, v6)
                _ = len(v3)
                _ = fi36
            }
            _ = fi32
        }
    case "\n\t":
        v37 := GenS4{F0: 0}
        _ = v37
    case "":
        var v38 uint8 = 2
        _ = v38
    default:
        _ = len(v2)
    }
    v39 := GenS3{F0: 2147483647}
    var v40 string = (v6 + v6)
    v41 := GenS5{F0: 42, F1: -128, F2: (false && (42 == 1))}
    v40 = ((v6 + (("\n\t" + (v6 + ((v40 + (v6 + "\n\t")) + (v40 + ((("\n\t" + "hello") + "hello") + ("foo bar" + (v6 + v6))))))) + ("a" + (v40 + v40)))) + (v6 + ((v40 + ((v6 + "foo bar") + (v40 + ""))) + (v40 + ((v40 + v6) + "a")))))
    fmt.Println(v40, v5, v16)
    if ((0.0 * 1.5) != (((0.1 / ((float64((0.1 * 0.0)) / 1.5) + (1.5 / (1.5 - 1e10)))) * 1e10) / 1.5)) {
        _ = len(v2)
        for fi42 := 0; fi42 < 15; fi42++ {
            v43 := GenS6{F0: 0}
            v5 = v5
            _ = fi42
            _ = v43
        }
    }
    v4 = (v4 & (-(v4)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
