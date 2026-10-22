package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    v4 := GenS1{F0: 702, F1: 0}
    var v5 uint = 502
    for fi6 := 0; fi6 < 16; fi6++ {
        v5 += (^(v5))
        v7 := []int{(fi6 >> (((v5 % ((4294967295) | 1))) & 63)), (fi6 / ((128) | 1)), 9223372036854775807}
        _ = fi6
        _ = v7
    }
    v5 += (v5 % ((255) | 1))
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{fi8, fi8, 9223372036854775807, (fi8 ^ (fi8 + (^(fi8))))}
        switch "a" {
        case "foo bar":
            fmt.Println(v5)
        case "\n\t":
            var v10 string = "a"
            _ = v10
        case "hello":
            var v11 int64 = -707
            _ = v11
        default:
            v12 := GenS2{F0: 42}
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v5)
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS2{F0: 2}
        _ = fi13
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 7; fi16++ {
                switch "foo bar" {
                case "foo bar":
                    var v17 int64 = 556
                    _ = v17
                case "hello":
                    fmt.Println(fi15)
                default:
                    v5 = (v5 % (((v5 / ((v5) | 1))) | 1))
                }
                _ = fi16
            }
            if false {
                var v18 float32 = 0.0
                _ = v18
            } else {
                v5 += v5
            }
            _ = fi15
        }
    } else {
        switch ("\n\t" + "a") {
        case "a":
            var v19 float64 = 0.1
            _ = v19
        case "foo bar":
            if (42 <= 527) {
                if (!((((false && (!true)) || (-74 != -128)) || false) || (!((-55 <= 100) && (!false))))) {
                    v5 += (v5 >> (((v5 / ((v5) | 1))) & 63))
                    switch -1 {
                    case 2147483647:
                        var v20 float64 = 1e10
                        _ = v20
                    default:
                        v21 := v3[-1]
                        _ = v21
                    }
                }
                if (!(2147483647 != -1)) {
                    v5 += v5
                } else {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        delete(v2, 65535)
                        for fi23 := 0; fi23 < 3; fi23++ {
                            v24 := GenS2{F0: 265}
                            if false {
                                var v25 uint16 = 515
                                if (!false) {
                                    var v26 int = 32768
                                    v27 := []int{}
                                    _ = v26
                                    _ = v27
                                }
                                _ = v25
                            } else {
                                _ = len(v1)
                                if false {
                                    if (916 <= 668) {
                                        var v28 uint8 = 42
                                        var v29 uint64 = 103
                                        v5 = (^(v5))
                                        v5 -= 328
                                        v5 = uint(v29)
                                        var v30 bool = (!false)
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    } else {
                                        var v31 int32 = -719
                                        _ = v31
                                    }
                                    v1[498] = ((0.0 < 1.5) && (0 == 42))
                                } else {
                                    var v32 int32 = 1
                                    var v33 int = (^(fi23))
                                    _ = v32
                                    _ = v33
                                }
                            }
                            v34 := GenS1{F0: 0, F1: -9223372036854775808}
                            _ = fi23
                            _ = v24
                            _ = v34
                        }
                        _ = fi22
                    }
                    if false {
                        v5 -= v5
                    }
                    switch "hello" {
                    case "hello":
                        fmt.Println(v5)
                    default:
                        v35 := GenS3{F0: 207, F1: 483}
                        _ = v35
                    }
                }
            }
        }
        for fi36 := 0; fi36 < 14; fi36++ {
            switch "" {
            case "hello":
                if false {
                    switch ("\n\t" + ("a" + ("foo bar" + ((("foo bar" + "\n\t") + "a") + ("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))))))) {
                    case "foo bar":
                        v37 := GenS3{F0: 128, F1: 32767}
                        _ = v37
                    case "\n\t":
                        switch "hello" {
                        case "a":
                            v2[0] = fi36
                        case "":
                            var v38 int32 = -658
                            _ = v38
                        default:
                            for fi39 := 0; fi39 < 15; fi39++ {
                                v5 = (v5 - v5)
                                var v40 float64 = -2.25
                                v40 = float64(1e10)
                                _ = fi39
                                _ = v40
                            }
                        }
                    default:
                        var v41 int32 = 42
                        _ = v41
                    }
                    var v42 uint64 = 1
                    switch 32767 {
                    case 32768:
                        v43 := GenS4{F0: 1, F1: -1}
                        _ = v43
                    default:
                        v44 := GenS3{F0: 0, F1: 1}
                        _ = v44
                    }
                    _ = v42
                } else {
                    for fi45 := 0; fi45 < 3; fi45++ {
                        fmt.Println(fi45)
                        _ = fi45
                    }
                }
            case "foo bar":
                var v46 bool = true
                _ = v46
            }
            _ = fi36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
