package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint16; F1 int64 }

func main() {
    var v1 float64 = (1.5 - (-2.25 + 0.1))
    var v2 uint16 = 32767
    var v3 uint = 763
    var v4 float64 = v1
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 32767, F1: 42, F2: -74}
        for fi7 := 0; fi7 < 4; fi7++ {
            if (2147483648 <= uint32(v2)) {
                v2 = v2
            }
            v8 := []int{fi5, fi7, 32768}
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := GenS2{F0: ("a" + "foo bar"), F1: 0}
        var v11 float32 = (1.5 / 0.1)
        var v12 bool = (100 <= 38)
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 float64 = (3.14159 + 0.0)
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            v2 += (^(v2))
            v16 := GenS3{F0: 523}
            _ = fi15
            _ = v16
        }
        if true {
            v17 := GenS3{F0: 255}
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(fi18, v13)
                v19 := GenS4{F0: 4294967295, F1: -1}
                if (0 > 1) {
                    v20 := GenS2{F0: "", F1: -543}
                    _ = v20
                } else {
                    var v21 int = fi18
                    var v22 int = (fi18 / ((fi14) | 1))
                    _ = v21
                    _ = v22
                }
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        _ = fi14
    }
    v3 = (v3 ^ (v3 & (v3 - (v3 & (^(v3))))))
    switch "foo bar" {
    case "\n\t":
        if (((v2 * (v2 << (((v3 ^ v3)) & 63))) > 0) && false) {
            switch -723 {
            case 32767:
                var v23 int16 = -673
                _ = v23
            case -2147483648:
                for fi24 := 0; fi24 < 12; fi24++ {
                    fmt.Println(v3, v3, v3)
                    if true {
                        if true {
                            var v25 uint16 = 530
                            for fi26 := 0; fi26 < 9; fi26++ {
                                v27 := GenS5{F0: 18446744073709551615, F1: "hello"}
                                v3 = (^(v3))
                                _ = fi26
                                _ = v27
                            }
                            fmt.Println(v3)
                            _ = v25
                        }
                        var v28 float32 = (0.1 * 0.0)
                        v29 := []int{-760, fi24}
                        _ = v28
                        _ = v29
                    }
                    _ = fi24
                }
            }
            v30 := []int{}
            _ = v30
        } else {
            v31 := GenS6{F0: v2, F1: -1}
            fmt.Println(v2, v3)
            _ = v31
        }
    case "hello":
        if false {
            for fi32 := 0; fi32 < 18; fi32++ {
                v1 -= 0.1
                _ = fi32
            }
        } else {
            if (207 == 32768) {
                var v33 float32 = (-2.25 / (0.0 / float32(v13)))
                _ = v33
            } else {
                v34 := []int{4294967295}
                var v35 uint = (v3 >> (((v3 / ((v3) | 1))) & 63))
                fmt.Println(v13, v3)
                _ = v34
                _ = v35
            }
            v3 -= 996
            var v36 int = 662
            _ = v36
        }
    case "a":
        v3 = v3
    }
    var v37 int64 = 42
    if (v37 != -1) {
        if true {
            for fi38 := 0; fi38 < 6; fi38++ {
                switch ("foo bar" + "hello") {
                case "foo bar":
                    switch fi38 {
                    case 32768:
                        var v39 string = "\n\t"
                        _ = v39
                    case 4294967295:
                        var v40 uint = 821
                        _ = v40
                    default:
                        var v41 float32 = 1e10
                        _ = v41
                    }
                case "\n\t":
                    v3 += 2
                case "hello":
                    var v42 int64 = (v37 << ((512) & 63))
                    _ = v42
                }
                if true {
                    if (-9 <= 1) {
                        for fi43 := 0; fi43 < 7; fi43++ {
                            v44 := []int{int(v37), -9223372036854775808, fi43}
                            var v45 float64 = v4
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        fmt.Println(v2)
                    } else {
                        var v46 int8 = 42
                        v47 := []int{fi38, (-(fi38)), (fi38 & int(v37)), 42}
                        _ = v46
                        _ = v47
                    }
                    switch "a" {
                    case "a":
                        v3 = v3
                    case "foo bar":
                        for fi48 := 0; fi48 < 18; fi48++ {
                            var v49 uint64 = 2
                            v1 -= 0.0
                            _ = fi48
                            _ = v49
                        }
                    }
                    v50 := []int{fi38, fi38, (fi38 / (((fi38 * 1)) | 1))}
                    _ = v50
                } else {
                    var v51 int16 = 42
                    v13 = v4
                    var v52 string = ("hello" + ("foo bar" + ("" + ("hello" + ""))))
                    _ = v51
                    _ = v52
                }
                _ = fi38
            }
            v1 = ((v4 + (v4 * 3.14159)) + v4)
        }
    }
    var v53 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v53
}
