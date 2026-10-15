package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int64; F1 string; F2 int32 }
type GenS3 struct { F0 string; F1 uint8; F2 int }
type GenS4 struct { F0 int16; F1 bool; F2 int8 }
type GenS5 struct { F0 int8; F1 int16 }

func main() {
    v1 := []int{-251, -804}
    var v2 uint32 = 2147483647
    var v3 uint64 = 42
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    for fi6 := 0; fi6 < 18; fi6++ {
        switch "a" {
        case "hello":
            var v7 float64 = (((0.0 - float64((1.5 + -2.25))) / 1e10) / 1.5)
            _ = v7
        case "\n\t":
            v2 = v2
        }
        _ = fi6
    }
    if true {
        fmt.Println(v3)
        if ((^(v2)) < (v2 / ((v2) | 1))) {
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int16 = -32768
                v2 = (v2 ^ (v2 % ((v2) | 1)))
                _ = fi8
                _ = v9
            }
        }
    } else {
        v2 += v2
        if false {
            if false {
                fmt.Println(v3)
                fmt.Println(v2, v3)
            } else {
                switch 943 {
                case 32767:
                    v10 := GenS1{F0: -1, F1: 0}
                    _ = v10
                case -672:
                    for fi11 := 0; fi11 < 10; fi11++ {
                        for fi12 := 0; fi12 < 10; fi12++ {
                            v2 -= (v2 | (v2 - (^(v2))))
                            _ = fi12
                        }
                        for fi13 := 0; fi13 < 2; fi13++ {
                            v2 = v2
                            _ = fi13
                        }
                        _ = fi11
                    }
                case 437:
                    v14 := GenS2{F0: -971, F1: "hello", F2: 0}
                    _ = v14
                default:
                    var v15 float32 = float32((0.1 + ((1e10 / 0.1) - 0.0)))
                    _ = v15
                }
                switch -334 {
                case 9223372036854775807:
                    v2 = (^(v2))
                case 127:
                    v2 += 263
                case 731:
                    fmt.Println(v3, v2)
                default:
                    var v16 int8 = 127
                    _ = v16
                }
                switch "\n\t" {
                case "":
                    v3 += (^(v3))
                case "a":
                    if true {
                        fmt.Println(v2, v3)
                        fmt.Println(v2, v2)
                        v17 := GenS2{F0: -9223372036854775808, F1: ("\n\t" + "hello"), F2: -2147483648}
                        _ = v17
                    }
                }
            }
            if ("foo bar" == ("foo bar" + "\n\t")) {
                for fi18 := 0; fi18 < 9; fi18++ {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        fmt.Println(fi19)
                        _ = fi19
                    }
                    fmt.Println(v2, v3, v2)
                    fmt.Println(v2, v2)
                    _ = fi18
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    var v21 int32 = 372
                    if len(v1) > 0 {
                        v1[0] = fi20
                    }
                    _ = fi20
                    _ = v21
                }
            } else {
                var v22 int8 = -122
                _ = v22
            }
            v23 := GenS2{F0: -9223372036854775808, F1: ("a" + ""), F2: 2147483647}
            _ = v23
        } else {
            var v24 float64 = -2.25
            if false {
                var v25 uint32 = v2
                v2 = v25
                v24 -= 1.5
                _ = v25
            } else {
                v26 := GenS3{F0: (("foo bar" + ("" + ("foo bar" + ""))) + ("\n\t" + "\n\t")), F1: 128, F2: 32768}
                _ = len(v4)
                for fi27 := 0; fi27 < 20; fi27++ {
                    v28 := GenS1{F0: 1, F1: -9223372036854775808}
                    _ = fi27
                    _ = v28
                }
                _ = v26
            }
            if false {
                v29 := GenS4{F0: 1, F1: (false && (!(32767 == 0))), F2: 100}
                fmt.Println(v2, v24)
                _ = v29
            }
            _ = v24
        }
        var v30 int64 = -442
        _ = v30
    }
    v3 += v3
    if false {
        v2 = (v2 | (v2 & v2))
        var v31 float64 = 1.5
        var v32 uint32 = v2
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        var v34 float64 = 1.5
        _ = fi33
        _ = v34
    }
    switch "foo bar" {
    case "\n\t":
        for fi35 := 0; fi35 < 2; fi35++ {
            if false {
                var v36 uint = 2
                for fi37 := 0; fi37 < 9; fi37++ {
                    switch "a" {
                    case "foo bar":
                        var v38 float32 = (((float32(3.14159) * -2.25) + 0.0) / -2.25)
                        _ = v38
                    }
                    _ = fi37
                }
                _ = v36
            } else {
                var v39 uint64 = (v3 + (^(v3)))
                _ = v39
            }
            _ = fi35
        }
    default:
        for fi40 := 0; fi40 < 5; fi40++ {
            if true {
                switch (("a" + "") + "foo bar") {
                case "":
                    v3 = v3
                }
            } else {
                fmt.Println(v3)
            }
            fmt.Println(v2, v2)
            switch ("\n\t" + "foo bar") {
            case "":
                v41 := GenS5{F0: 1, F1: -1}
                _ = v41
            }
            _ = fi40
        }
    }
    fmt.Println(v2)
    var v42 uint8 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
