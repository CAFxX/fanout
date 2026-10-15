package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int; F1 int64; F2 string }
type GenS6 struct { F0 string; F1 uint }
type GenS7 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    var v3 int8 = -23
    var v4 bool = true
    v5 := v1[482]
    v6 := GenS1{F0: 1}
    v7 := v1[4294967295]
    v8 := GenS1{F0: 65535}
    v7 = (false || true)
    fmt.Println(v4, v3, v4)
    if v7 {
        var v9 float64 = (0.0 * 1.5)
        _ = v9
    } else {
        var v10 int64 = -851
        switch -2147483648 {
        case -338:
            v5 = v7
        }
        switch "hello" {
        case "hello":
            var v11 uint16 = 1
            _ = v11
        default:
            var v12 uint64 = 18446744073709551615
            _ = v12
        }
        _ = v10
    }
    if v7 {
        v4 = (0.1 >= 1.5)
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := GenS2{F0: 0, F1: fi13, F2: "foo bar"}
            _ = fi13
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 16; fi15++ {
            if v7 {
                v16 := []int{fi15, (^(fi15)), -278, fi15}
                _ = v16
            } else {
                fmt.Println(fi15)
            }
            _ = fi15
        }
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            if (!v7) {
                if true {
                    v19 := v2[(fi17 % ((32768) | 1))]
                    var v20 bool = v5
                    _ = v19
                    _ = v20
                }
                var v21 float64 = 1e10
                _ = v21
            }
            _ = fi18
        }
        _ = fi17
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            switch 127 {
            case 9223372036854775807:
                v4 = (126 == 1)
            default:
                v24 := GenS3{F0: 42}
                _ = v24
            }
            v25 := v2[-982]
            fmt.Println(fi23, v4)
            _ = fi23
            _ = v25
        }
        if v5 {
            for fi26 := 0; fi26 < 6; fi26++ {
                v4 = ((8 > 175) && v5)
                if (1000 <= 42) {
                    v5 = ((42 <= 728) || ((1 != 42) || v4))
                } else {
                    v27 := v2[(fi22 << ((4294967295) & 63))]
                    _ = v27
                }
                v28 := []int{fi26, 127, fi22, (fi22 ^ fi26)}
                _ = fi26
                _ = v28
            }
            switch fi22 {
            case 32767:
                fmt.Println(v3)
            default:
                if v7 {
                    if false {
                        fmt.Println(v7, v7)
                    } else {
                        if (true && v5) {
                            fmt.Println(fi22, v3, v4)
                        }
                        if false {
                            for fi29 := 0; fi29 < 7; fi29++ {
                                v30 := GenS4{F0: 9223372036854775807, F1: 65535, F2: 0}
                                var v31 bool = v4
                                var v32 int64 = -613
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            var v33 uint16 = 2
                            var v34 int = fi22
                            _ = v33
                            _ = v34
                        }
                        var v35 float64 = 0.1
                        _ = v35
                    }
                    if (!v5) {
                        switch "foo bar" {
                        case "hello":
                            v36 := v1[418]
                            _ = v36
                        case "foo bar":
                            v37 := v2[fi22]
                            _ = v37
                        }
                        var v38 int = (fi22 ^ (fi22 << ((4294967295) & 63)))
                        var v39 float32 = 1e10
                        _ = v38
                        _ = v39
                    }
                    for fi40 := 0; fi40 < 16; fi40++ {
                        v41 := GenS5{F0: -271, F1: int64(fi22), F2: "foo bar"}
                        _ = fi40
                        _ = v41
                    }
                }
            }
        } else {
            v42 := GenS6{F0: "", F1: 893}
            _ = v42
        }
        switch ("\n\t" + "") {
        case "foo bar":
            for fi43 := 0; fi43 < 20; fi43++ {
                v44 := GenS1{F0: 283}
                v45 := []int{fi43, (fi43 >> ((65535) & 63))}
                var v46 int = (fi22 + -1)
                fmt.Println(fi43, fi43, v7)
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
        case "\n\t":
            v47 := []int{(fi22 - (-(fi22)))}
            _ = v47
        case "a":
            v48 := GenS7{F0: 33}
            _ = v48
        default:
            v7 = v4
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
