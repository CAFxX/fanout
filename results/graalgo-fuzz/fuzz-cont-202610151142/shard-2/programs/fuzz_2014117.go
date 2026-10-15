package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint64; F2 string }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int8; F1 bool; F2 uint64 }
func genFn0(p0 int, p1 uint, p2 int64) uint {
    var v4 int16 = 42
    if (42 != 2147483648) {
        var v5 int32 = 42
        v6 := GenS1{F0: 1, F1: (-(p0)), F2: (p2 % (((p2 << (((^(p1))) & 63))) | 1))}
        for fi7 := 0; fi7 < 14; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 uint64 = 42
                switch "hello" {
                case "":
                    v10 := []int{}
                    _ = v10
                default:
                    p1 = (p1 >> (((p1 | 1)) & 63))
                }
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        _ = v5
        _ = v6
    } else {
        v11 := GenS1{F0: 1, F1: p0, F2: p2}
        _ = v11
    }
    if true {
        if false {
            p2 -= p2
        }
        var v12 int64 = -9223372036854775808
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p1
}
func genFn1(p0 int64, p1 int, p2 uint64) uint64 {
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint8 = 2
        _ = fi13
        _ = v14
    }
    v15 := []int{819, int(p0), (p1 * p1)}
    if true {
        var v16 string = "hello"
        _ = v16
    } else {
        v17 := GenS2{F0: p2}
        v18 := GenS3{F0: p0, F1: 0}
        v19 := GenS4{F0: 164, F1: 2, F2: ("\n\t" + ("" + ("foo bar" + "hello")))}
        _ = v17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return (p2 & (^(p2)))
}

func main() {
    v1 := []string{"\n\t", ("hello" + ("hello" + (("\n\t" + "foo bar") + "hello"))), (("a" + "") + "")}
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    if false {
        v20 := GenS4{F0: 127, F1: 42, F2: ("a" + ("foo bar" + ("\n\t" + "")))}
        var v21 bool = (true && false)
        for fi22 := 0; fi22 < 7; fi22++ {
            if ((!v21) || (!true)) {
                fmt.Println(v21)
                var v23 int32 = 0
                _ = v23
            }
            if (5 < 2) {
                switch (("\n\t" + ("foo bar" + (("" + (("\n\t" + "\n\t") + "\n\t")) + "foo bar"))) + "\n\t") {
                case "foo bar":
                    if v21 {
                        fmt.Println(fi22)
                        var v24 int32 = -2147483648
                        v25 := GenS5{F0: 1, F1: 4294967295}
                        _ = v24
                        _ = v25
                    } else {
                        fmt.Println(fi22, fi22)
                        if v21 {
                            if len(v1) > 0 {
                                v1[0] = ("hello" + "\n\t")
                            }
                            var v26 uint32 = 65535
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 13; fi27++ {
                                var v28 bool = (-721 >= 2147483647)
                                fmt.Println(v28, fi27, v21)
                                var v29 float64 = 0.0
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                            var v30 int64 = 9223372036854775807
                            _ = v30
                        }
                    }
                default:
                    var v31 uint = 2
                    _ = v31
                }
                var v32 int16 = -32768
                _ = v32
            } else {
                v33 := GenS3{F0: 1, F1: 0}
                _ = v33
            }
            _ = fi22
        }
        _ = v20
        _ = v21
    } else {
        var v34 uint16 = 32767
        _ = v34
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        fmt.Println(fi35)
        _ = fi35
    }
    v36 := GenS6{F0: -118, F1: true, F2: 4294967295}
    var v37 int16 = 1
    switch 32768 {
    case 2147483647:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -236:
        v38 := GenS1{F0: 0, F1: 42, F2: 9223372036854775807}
        _ = v38
    }
    var v39 int8 = -83
    if (778 != 910) {
        var v40 float32 = float32(0.1)
        _ = v40
    }
    v41 := v3[682]
    for fi42 := 0; fi42 < 17; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            v44 := GenS5{F0: 434, F1: (v41 ^ v41)}
            _ = fi43
            _ = v44
        }
        v45 := GenS5{F0: 18446744073709551615, F1: v41}
        _ = fi42
        _ = v45
    }
    var v46 float32 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v41
    _ = v46
}
