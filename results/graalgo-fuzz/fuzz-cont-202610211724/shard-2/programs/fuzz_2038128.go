package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int; F2 uint64 }
type GenS5 struct { F0 uint16; F1 string }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 int = fi4
        _ = fi4
        _ = v5
    }
    var v6 int16 = 42
    v6 -= (v6 % ((v6) | 1))
    for fi7 := 0; fi7 < 19; fi7++ {
        v6 = v6
        switch "a" {
        case "foo bar":
            if (-1 != -2147483648) {
                if true {
                    if true {
                        if true {
                            for fi8 := 0; fi8 < 13; fi8++ {
                                v9 := []int{}
                                v6 -= v6
                                _ = fi8
                                _ = v9
                            }
                        }
                        if true {
                            var v10 bool = false
                            if ((!false) || false) {
                                v10 = ("foo bar" <= "a")
                                switch "" {
                                case "foo bar":
                                    if (2 > 18446744073709551615) {
                                        if v10 {
                                            var v11 uint32 = 2147483647
                                            v11 = (v11 << ((255) & 63))
                                            var v12 uint8 = 2
                                            _ = v11
                                            _ = v12
                                        } else {
                                            var v13 string = "hello"
                                            for fi14 := 0; fi14 < 18; fi14++ {
                                                if v10 {
                                                    v15 := []int{}
                                                    if false {
                                                        _ = len(v15)
                                                        switch (fi14 % ((128) | 1)) {
                                                        case -1:
                                                            v10 = v10
                                                        case -9223372036854775808:
                                                            v13 = v13
                                                        case 9223372036854775807:
                                                            v15 = append(v15, fi7)
                                                        default:
                                                            v6 = (v6 - (v6 & (v6 | 32767)))
                                                        }
                                                    }
                                                    _ = v15
                                                }
                                                _ = fi14
                                            }
                                            _ = v13
                                        }
                                        if (-2.25 < -2.25) {
                                            var v16 float64 = 0.1
                                            var v17 int16 = 42
                                            v18 := GenS1{F0: -2147483648, F1: (v17 << ((183) & 63))}
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        }
                                        v19 := []int{-1}
                                        _ = v19
                                    }
                                default:
                                    v20 := []int{-950, (fi7 << ((255) & 63)), -1, (fi7 / ((692) | 1))}
                                    _ = v20
                                }
                            }
                            _ = v10
                        } else {
                            var v21 uint = 819
                            v22 := []int{fi7, (-(fi7)), fi7, 4294967295}
                            _ = v21
                            _ = v22
                        }
                    } else {
                        v6 = -1
                    }
                    if true {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            var v24 int8 = -75
                            v24 += 13
                            _ = fi23
                            _ = v24
                        }
                        v25 := []int{fi7, 9223372036854775807, -1, fi7}
                        _ = v25
                    } else {
                        var v26 float32 = 1.5
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 9; fi27++ {
                        var v28 uint8 = 128
                        _ = fi27
                        _ = v28
                    }
                } else {
                    v6 = 204
                }
                v6 += v6
            } else {
                fmt.Println(fi7, fi7)
            }
        default:
            var v29 uint = 332
            _ = v29
        }
        _ = fi7
    }
    v6 = v6
    _ = v6
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{2, 4294967295, 2}
    v3 := make(map[int]uint32)
    fmt.Println(v1)
    fmt.Println(v1)
    v30 := GenS1{F0: -774, F1: 42}
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 string = ("\n\t" + "")
        v33 := GenS2{F0: -1, F1: 2, F2: 4294967295}
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        for fi35 := 0; fi35 < 1; fi35++ {
            switch "" {
            case "hello":
                if (((fi34 / ((fi35) | 1)) != (^(fi34))) || false) {
                    switch (fi34 % (((fi34 * fi34)) | 1)) {
                    case 32767:
                        var v36 uint8 = 1
                        _ = v36
                    case 128:
                        v37 := GenS3{F0: 32767}
                        _ = v37
                    }
                    if true {
                        if (!true) {
                            switch "foo bar" {
                            case "\n\t":
                                var v38 uint64 = 1
                                _ = v38
                            case "":
                                var v39 uint = 18446744073709551615
                                _ = v39
                            default:
                                var v40 int8 = 91
                                _ = v40
                            }
                            v1 = v1
                            var v41 bool = ((0.0 * 3.14159) < float32(-2.25))
                            _ = v41
                        }
                        v42 := GenS4{F0: 42, F1: (fi34 & int(v1)), F2: 397}
                        fmt.Println(fi35, fi35)
                        _ = v42
                    } else {
                        if len(v2) > 0 {
                            v2[0] = 619
                        }
                    }
                }
            }
            _ = fi35
        }
        v1 = v1
        var v43 uint16 = 93
        _ = fi34
        _ = v43
    }
    delete(v3, -910)
    switch "" {
    case "hello":
        v44 := GenS5{F0: 32767, F1: (("a" + "") + "foo bar")}
        _ = v44
    case "\n\t":
        v45 := v3[-2147483648]
        _ = v45
    case "a":
        v1 = v1
    }
    for fi46 := 0; fi46 < 17; fi46++ {
        switch fi46 {
        case 32768:
            v1 = (v1 / (((-(v1))) | 1))
        case 128:
            var v47 int64 = -1
            _ = v47
        case 4294967295:
            v1 += 329
        }
        _ = fi46
    }
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v48 uint = 513
    v49 := v3[449]
    switch (("hello" + "") + "foo bar") {
    case "a":
        fmt.Println(v48)
    case "foo bar":
        v49 += 65535
    }
    for fi50 := 0; fi50 < 17; fi50++ {
        for fi51 := 0; fi51 < 16; fi51++ {
            var v52 bool = (!true)
            _ = fi51
            _ = v52
        }
        var v53 float64 = 0.1
        fmt.Println(v1)
        _ = fi50
        _ = v53
    }
    var v54 int = int(v1)
    var v55 uint64 = 594
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v48
    _ = v49
    _ = v54
    _ = v55
}
