package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 int }
func genFn0(p0 int16, p1 int, p2 uint) uint {
    p0 = p0
    var v6 uint16 = uint16(p2)
    fmt.Println(p2, v6)
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint8 = 242
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}

func main() {
    var v1 uint8 = 0
    var v2 string = ""
    var v3 string = v2
    var v4 uint32 = 0
    v5 := make(map[string]int)
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: fi9}
        _ = fi9
        _ = v10
    }
    if true {
        v3 = "a"
    }
    var v11 int8 = 127
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int64 = 1
        delete(v5, ("a" + v2))
        v14 := []int{}
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v3, v3)
    v11 += v11
    if ((v4 ^ (v4 | 666)) > v4) {
        switch -357 {
        case -230:
            for fi15 := 0; fi15 < 3; fi15++ {
                v2 = v3
                _ = fi15
            }
        case 32767:
            var v16 int = 711
            _ = v16
        case 86:
            var v17 uint16 = 650
            _ = v17
        }
    }
    var v18 uint = 0
    if true {
        v2 = "hello"
        for fi19 := 0; fi19 < 8; fi19++ {
            if false {
                switch fi19 {
                case 530:
                    v18 -= uint(v1)
                }
            } else {
                v3 = v3
                var v20 int64 = 9223372036854775807
                _ = v20
            }
            _ = fi19
        }
        v21 := []int{4294967295, 32768, 399}
        _ = v21
    } else {
        var v22 float64 = 1e10
        if false {
            var v23 uint32 = v4
            var v24 int16 = 0
            var v25 float64 = v22
            _ = v23
            _ = v24
            _ = v25
        } else {
            v2 = ((((v3 + v2) + "foo bar") + "hello") + "a")
            fmt.Println(v2, v1)
            var v26 uint = 65535
            _ = v26
        }
        if true {
            fmt.Println(v11, v3)
        } else {
            v27 := []int{}
            switch "foo bar" {
            case "":
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
            case "foo bar":
                var v28 uint = genFn0(0, 32768, 1)
                _ = v28
            }
            for fi29 := 0; fi29 < 6; fi29++ {
                if false {
                    v4 += (v4 | (^(v4)))
                }
                v30 := GenS2{F0: -175}
                if true {
                    v31 := GenS3{F0: 43}
                    switch "hello" {
                    case "hello":
                        v4 = 1
                    case "foo bar":
                        for fi32 := 0; fi32 < 4; fi32++ {
                            if true {
                                fmt.Println(fi32, v18)
                                v33 := GenS4{F0: 65535}
                                v34 := GenS1{F0: (fi29 & (fi32 / ((fi29) | 1)))}
                                v11 = int8(v4)
                                switch ("" + "hello") {
                                case "\n\t":
                                    v2 = (v2 + ("foo bar" + (v3 + "a")))
                                case "hello":
                                    switch "" {
                                    case "a":
                                        var v35 uint8 = v1
                                        _ = v35
                                    case "\n\t":
                                        if false {
                                            if false {
                                                _ = len(v27)
                                                fmt.Println(v4, fi32, v22)
                                                v18 = v18
                                            }
                                        } else {
                                            if (fi29 < (fi32 % (((fi32 ^ fi32)) | 1))) {
                                                fmt.Println(v4, v1, fi32)
                                                var v36 int16 = 766
                                                switch fi32 {
                                                case 2147483647:
                                                    v4 = (v4 ^ 42)
                                                case -9223372036854775808:
                                                    if true {
                                                        v1 = (v1 >> ((genFn0(v36, 450, v18)) & 63))
                                                        v1 -= (v1 / ((97) | 1))
                                                    }
                                                case 4294967295:
                                                    if ((false && (v4 > 2)) && (true || false)) {
                                                        var v37 uint8 = v1
                                                        var v38 uint = v18
                                                        v4 = v4
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        fmt.Println(v2, v22, v18)
                                                        var v39 uint16 = 2
                                                        if (42 <= -9223372036854775808) {
                                                            if ("foo bar" < "") {
                                                                var v40 uint32 = (v4 | 707)
                                                                var v41 string = v3
                                                                fmt.Println(v3)
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                            v1 = (v1 / ((uint8(v18)) | 1))
                                                            var v42 uint = genFn0((v36 % (((-(v36))) | 1)), 626, 755)
                                                            _ = v42
                                                        } else {
                                                            var v43 uint64 = 617
                                                            v22 = 0.0
                                                            _ = v43
                                                        }
                                                        var v44 uint32 = 2147483647
                                                        v1 = uint8(v4)
                                                        _ = v39
                                                        _ = v44
                                                    }
                                                default:
                                                    v27 = append(v27, 4294967295)
                                                }
                                                _ = v36
                                            } else {
                                                _ = len(v27)
                                                if false {
                                                    v11 = 100
                                                    var v45 float64 = v22
                                                    if len(v27) > 0 {
                                                        v27[0] = fi32
                                                    }
                                                    _ = v45
                                                }
                                            }
                                            if (0 > 354) {
                                                v27 = append(v27, (fi29 * (fi29 + (fi32 ^ 127))))
                                                if true {
                                                    var v46 int64 = -9223372036854775808
                                                    if true {
                                                        switch v2 {
                                                        case "":
                                                            var v47 int16 = 225
                                                            _ = v47
                                                        case "foo bar":
                                                            v46 += v46
                                                        }
                                                        fmt.Println(v46, v46, fi32)
                                                        var v48 uint32 = (^(v4))
                                                        _ = v48
                                                    }
                                                    _ = v46
                                                } else {
                                                    switch -334 {
                                                    case 42:
                                                        if true {
                                                            var v49 float64 = v22
                                                            switch (v2 + v3) {
                                                            case "":
                                                                var v50 bool = true
                                                                _ = v50
                                                            default:
                                                                if (2147483647 >= 355) {
                                                                    var v51 float32 = 0.1
                                                                    fmt.Println(v2)
                                                                    var v52 uint32 = (v4 | 909)
                                                                    _ = v51
                                                                    _ = v52
                                                                } else {
                                                                    switch (fi29 % ((-184) | 1)) {
                                                                    case -151:
                                                                        v11 = v11
                                                                    }
                                                                }
                                                            }
                                                            v53 := GenS5{F0: ""}
                                                            fmt.Println(v11)
                                                            _ = v49
                                                            _ = v53
                                                        } else {
                                                            v54 := GenS3{F0: int32(v4)}
                                                            if false {
                                                                v22 += (1e10 + float64((0.0 / 1.0)))
                                                                var v55 int64 = 9223372036854775807
                                                                if false {
                                                                    v22 += v22
                                                                    fmt.Println(v18, fi32)
                                                                    v1 = (v1 >> (((v18 >> ((v18) & 63))) & 63))
                                                                }
                                                                _ = v55
                                                            }
                                                            _ = v54
                                                        }
                                                    }
                                                    v22 += v22
                                                    _ = len(v5)
                                                }
                                                fmt.Println(v1, v18, fi29)
                                            }
                                            v5[("\n\t" + "")] = -2147483648
                                        }
                                    case "foo bar":
                                        fmt.Println(fi32)
                                    }
                                case "":
                                    delete(v5, ("foo bar" + (("hello" + v3) + v3)))
                                    var v56 uint = uint(v4)
                                    _ = v56
                                default:
                                    switch v2 {
                                    case "foo bar":
                                        v22 += 0.0
                                    }
                                }
                                _ = v33
                                _ = v34
                            } else {
                                if ((v3 >= v3) || false) {
                                    v4 = v4
                                    var v57 float64 = v22
                                    _ = v57
                                }
                            }
                            var v58 uint16 = 894
                            fmt.Println(v2, v11)
                            _ = fi32
                            _ = v58
                        }
                    default:
                        v59 := GenS5{F0: "hello"}
                        _ = v59
                    }
                    _ = v31
                }
                _ = fi29
                _ = v30
            }
            _ = v27
        }
        _ = v22
    }
    switch "" {
    case "hello":
        var v60 uint16 = 0
        _ = v60
    case "\n\t":
        if (!(false && false)) {
            if (((42 == 1) && true) || false) {
                if true {
                    var v61 uint64 = 2
                    _ = v61
                } else {
                    delete(v5, v3)
                    v62 := GenS3{F0: -2147483648}
                    _ = v62
                }
                v2 = v3
                var v63 int32 = -227
                _ = v63
            }
            switch -960 {
            case 32768:
                var v64 int = 1
                _ = v64
            case 32767:
                switch "hello" {
                case "a":
                    var v65 string = v3
                    _ = v65
                case "foo bar":
                    var v66 uint = v18
                    _ = v66
                case "":
                    v5[v3] = 4294967295
                default:
                    for fi67 := 0; fi67 < 2; fi67++ {
                        v68 := []int{}
                        _ = fi67
                        _ = v68
                    }
                }
            }
            v69 := GenS6{F0: 16, F1: 0}
            _ = v69
        } else {
            var v70 bool = false
            if v70 {
                v71 := GenS3{F0: 42}
                v72 := v5["\n\t"]
                switch (^(v72)) {
                case 32767:
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v4, v4)
                    case "":
                        var v73 int64 = 540
                        _ = v73
                    case "a":
                        v74 := GenS1{F0: v72}
                        _ = v74
                    }
                case 0:
                    v18 = (v18 - 1)
                case 2147483647:
                    delete(v5, v2)
                }
                _ = v71
                _ = v72
            } else {
                var v75 int64 = 662
                _ = v75
            }
            _ = v70
        }
    case "foo bar":
        for fi76 := 0; fi76 < 1; fi76++ {
            v3 = (("\n\t" + "a") + ("" + v3))
            fmt.Println(v4)
            v4 += v4
            _ = fi76
        }
    default:
        switch "hello" {
        case "":
            v77 := GenS6{F0: int(v1), F1: 9223372036854775807}
            _ = v77
        }
    }
    var v78 int8 = v11
    v11 = (v11 ^ (v11 ^ (^(v78))))
    var v79 uint = (v18 << (((v18 ^ v18)) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v79
}
