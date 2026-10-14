package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]bool)
    var v4 int16 = 32767
    v5 := GenS1{F0: v4}
    var v6 float32 = 0.0
    var v7 bool = (1 <= 0)
    if ((0.0 != 3.14159) && (42 <= -37)) {
        switch "hello" {
        case "":
            for fi8 := 0; fi8 < 14; fi8++ {
                v4 = v4
                v9 := []int{}
                _ = fi8
                _ = v9
            }
        case "hello":
            v6 = (((v6 * v6) * (1.5 + v6)) + 1e10)
        }
        var v10 bool = false
        _ = v10
    } else {
        v11 := []int{570, -1, -2147483648, 356}
        if ((!v7) || (v7 || v7)) {
            if v7 {
                if (v7 || false) {
                    if (-2147483648 >= 777) {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    }
                    var v12 int64 = 42
                    _ = v12
                } else {
                    v7 = (v7 || (v7 || v7))
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v4 -= v4
                        _ = fi13
                    }
                    v14 := GenS2{F0: 2147483648}
                    _ = v14
                }
            } else {
                var v15 uint = 18446744073709551615
                delete(v3, 265)
                _ = len(v3)
                _ = v15
            }
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS1{F0: v4}
                if ((!false) && (3.14159 < (1e10 + 1e10))) {
                    var v18 string = "a"
                    var v19 int64 = 42
                    _ = v18
                    _ = v19
                }
                switch fi16 {
                case -9223372036854775808:
                    v2 = 2
                default:
                    for fi20 := 0; fi20 < 14; fi20++ {
                        var v21 int = fi20
                        _ = fi20
                        _ = v21
                    }
                }
                _ = fi16
                _ = v17
            }
        }
        v7 = (float64(3.14159) >= (3.14159 * (-2.25 + -2.25)))
        _ = v11
    }
    var v22 bool = false
    if ((v7 && v22) || true) {
        v7 = (v2 == v2)
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        v24 := v3[-1]
        _ = fi23
        _ = v24
    }
    var v25 float32 = -2.25
    if (v4 < v4) {
        if false {
            for fi26 := 0; fi26 < 9; fi26++ {
                var v27 uint = 18446744073709551615
                _ = fi26
                _ = v27
            }
            if (65535 > 1) {
                for fi28 := 0; fi28 < 3; fi28++ {
                    if v7 {
                        v7 = (!(1 != (-(v4))))
                    }
                    for fi29 := 0; fi29 < 11; fi29++ {
                        var v30 uint16 = 541
                        _ = fi29
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 13; fi31++ {
                        var v32 uint8 = 128
                        v3[(fi31 << ((1) & 63))] = false
                        var v33 uint = 767
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    _ = fi28
                }
                var v34 bool = v22
                for fi35 := 0; fi35 < 9; fi35++ {
                    fmt.Println(v6)
                    _ = fi35
                }
                _ = v34
            }
            fmt.Println(v2, v2)
        } else {
            fmt.Println(v22)
            if v7 {
                fmt.Println(v7)
            }
        }
    }
    v2 += v2
    if (!((1 > -546) || v7)) {
        v36 := []int{-984}
        _ = v36
    }
    var v37 float32 = -2.25
    switch "\n\t" {
    case "":
        for fi38 := 0; fi38 < 11; fi38++ {
            if v22 {
                switch 574 {
                case 42:
                    fmt.Println(v6)
                case 1:
                    if (!("a" < "foo bar")) {
                        fmt.Println(v25, v22, v25)
                        var v39 uint32 = 2147483648
                        v3[(fi38 % ((1) | 1))] = v7
                        _ = v39
                    } else {
                        v6 = 0.0
                        v6 += v25
                        v40 := []int{}
                        _ = v40
                    }
                }
                v41 := []int{}
                if v7 {
                    if false {
                        var v42 int16 = 32767
                        _ = v42
                    } else {
                        for fi43 := 0; fi43 < 6; fi43++ {
                            var v44 bool = ((fi38 << ((1) & 63)) <= (-(fi38)))
                            v45 := GenS1{F0: v4}
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        fmt.Println(v25, v7, v6)
                        v2 -= v2
                    }
                    v37 = (0.0 + v6)
                } else {
                    v4 -= (v4 - 42)
                    if len(v41) >= 2 {
                        v41 = v41[1:len(v41)]
                    }
                }
                _ = v41
            } else {
                v4 += (v4 / ((42) | 1))
            }
            _ = fi38
        }
    case "foo bar":
        v46 := v1[2147483648]
        _ = v46
    }
    var v47 int32 = -1
    switch "a" {
    case "hello":
        v47 += (v47 - v47)
    case "\n\t":
        v48 := GenS1{F0: 42}
        _ = v48
    case "a":
        for fi49 := 0; fi49 < 15; fi49++ {
            switch "a" {
            case "a":
                var v50 int32 = v47
                _ = v50
            case "\n\t":
                switch "hello" {
                case "foo bar":
                    if (65535 < 214) {
                        var v51 uint32 = 1
                        for fi52 := 0; fi52 < 6; fi52++ {
                            if false {
                                v53 := GenS3{F0: 619, F1: 109}
                                _ = v53
                            }
                            _ = len(v3)
                            _ = fi52
                        }
                        _ = v51
                    } else {
                        var v54 uint16 = 2
                        var v55 float32 = v25
                        _ = v54
                        _ = v55
                    }
                case "hello":
                    fmt.Println(v47, fi49, v6)
                case "\n\t":
                    v56 := GenS4{F0: 827, F1: v47}
                    _ = v56
                default:
                    var v57 uint8 = 52
                    _ = v57
                }
            default:
                var v58 int64 = 9223372036854775807
                _ = v58
            }
            switch (fi49 - (-(fi49))) {
            case 42:
                fmt.Println(v2, fi49)
            case -226:
                for fi59 := 0; fi59 < 18; fi59++ {
                    delete(v1, 1)
                    _ = fi59
                }
            case -1:
                for fi60 := 0; fi60 < 18; fi60++ {
                    v61 := GenS4{F0: 133, F1: (v47 ^ v47)}
                    _ = fi60
                    _ = v61
                }
            }
            _ = fi49
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
