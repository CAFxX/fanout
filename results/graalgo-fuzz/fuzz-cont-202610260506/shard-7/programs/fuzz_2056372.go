package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint32; F2 int }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 string; F1 uint32; F2 int8 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 uint32; F1 int; F2 int16 }
func genFn0(p0 uint) uint16 {
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: 4294967295}
        if true {
            for fi6 := 0; fi6 < 4; fi6++ {
                fmt.Println(p0)
                p0 += (p0 * p0)
                var v7 uint32 = 0
                _ = fi6
                _ = v7
            }
            if false {
                switch fi4 {
                case -822:
                    if (!false) {
                        switch fi4 {
                        case 0:
                            v8 := GenS2{F0: 2, F1: ("" + ""), F2: 1}
                            _ = v8
                        }
                        v9 := GenS2{F0: 803, F1: "foo bar", F2: 188}
                        v10 := GenS3{F0: 42, F1: 89}
                        _ = v10
                        _ = v9
                    } else {
                        v11 := []int{(fi4 * 376), (fi4 / ((128) | 1)), (fi4 ^ (^(fi4))), fi4}
                        _ = len(v11)
                        _ = v11
                    }
                case -9223372036854775808:
                    p0 = 1
                }
            } else {
                fmt.Println(p0, p0)
                if (!(!(2 < 96))) {
                    v12 := GenS3{F0: 444, F1: 211}
                    v13 := GenS4{F0: p0, F1: 2147483647, F2: fi4}
                    var v14 int16 = 42
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        p0 = p0
                        p0 = p0
                        v16 := GenS2{F0: 42, F1: ("foo bar" + ("\n\t" + "a")), F2: 206}
                        _ = fi15
                        _ = v16
                    }
                }
                fmt.Println(fi4, fi4)
            }
        }
        p0 += p0
        _ = fi4
        _ = v5
    }
    p0 = p0
    p0 += (p0 - p0)
    for fi17 := 0; fi17 < 19; fi17++ {
        switch "a" {
        case "":
            var v18 string = ""
            _ = v18
        case "foo bar":
            v19 := GenS2{F0: 65535, F1: "a", F2: 122}
            _ = v19
        case "a":
            v20 := GenS2{F0: 42, F1: ("foo bar" + (("hello" + ("a" + "hello")) + ("" + "\n\t"))), F2: 97}
            _ = v20
        }
        fmt.Println(p0, p0)
        p0 += 1
        _ = fi17
    }
    _ = p0
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -816
    var v3 float32 = 3.14159
    switch "\n\t" {
    case "":
        v3 = (1.5 / float32(1e10))
    case "hello":
        var v21 int16 = 42
        _ = v21
    }
    v3 -= float32(float64(1.5))
    fmt.Println(v3, v3)
    v3 = (1e10 + v3)
    v2 += (v2 % (((v2 - v2)) | 1))
    var v22 int64 = 9223372036854775807
    switch (v2 << ((255) & 63)) {
    case -183:
        _ = len(v1)
    }
    if ((0.1 * 1e10) > float64((v3 * ((0.0 + (-2.25 + 1.5)) / -2.25)))) {
        if (v22 == v22) {
            v23 := []int{}
            var v24 float32 = v3
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 8; fi25++ {
                for fi26 := 0; fi26 < 6; fi26++ {
                    v2 = (v2 & 9223372036854775807)
                    if (!false) {
                        if false {
                            v27 := GenS3{F0: 2147483647, F1: 241}
                            _ = v27
                        }
                        var v28 uint64 = 42
                        _ = v28
                    }
                    _ = fi26
                }
                v22 = (v22 / ((v22) | 1))
                _ = fi25
            }
        }
        var v29 int8 = 42
        _ = v29
    }
    v22 -= -127
    if (!(921 > 712)) {
        switch v2 {
        case 128:
            v2 = 2147483647
        case 824:
            for fi30 := 0; fi30 < 15; fi30++ {
                v3 = v3
                fmt.Println(v2, v22)
                _ = fi30
            }
        case 9223372036854775807:
            v31 := GenS4{F0: 255, F1: 4294967295, F2: v2}
            _ = v31
        default:
            if true {
                for fi32 := 0; fi32 < 18; fi32++ {
                    v33 := GenS5{F0: (v22 / (((v22 * 80)) | 1)), F1: fi32}
                    if (42 < 243) {
                        for fi34 := 0; fi34 < 12; fi34++ {
                            v35 := GenS6{F0: "foo bar", F1: 2147483648, F2: -128}
                            if false {
                                switch (("" + "a") + ((("\n\t" + "foo bar") + "\n\t") + "hello")) {
                                case "foo bar":
                                    var v36 int = -109
                                    _ = v36
                                case "hello":
                                    v2 -= (fi34 % ((fi34) | 1))
                                case "\n\t":
                                    var v37 int64 = v22
                                    _ = v37
                                default:
                                    if false {
                                        delete(v1, 0)
                                    } else {
                                        if ("\n\t" > ("" + (("\n\t" + "a") + "a"))) {
                                            v3 = v3
                                            switch "foo bar" {
                                            case "\n\t":
                                                v3 = 0.1
                                            case "":
                                                v38 := GenS3{F0: 972, F1: 11}
                                                _ = v38
                                            }
                                            fmt.Println(v2, v2, v22)
                                        }
                                        v39 := GenS6{F0: ("a" + ""), F1: 0, F2: -1}
                                        _ = v39
                                    }
                                }
                                var v40 uint32 = 42
                                switch 4294967295 {
                                case -278:
                                    v41 := GenS7{F0: 0, F1: v2}
                                    _ = v41
                                default:
                                    if true {
                                        if ((((v22 + v22) == 1) || (((!(333 >= 42)) && (483 != v22)) && true)) && false) {
                                            var v42 int8 = -104
                                            fmt.Println(fi34, v2)
                                            _ = v42
                                        } else {
                                            fmt.Println(fi32)
                                        }
                                        v43 := GenS4{F0: 0, F1: v40, F2: fi34}
                                        _ = v43
                                    }
                                }
                                _ = v40
                            }
                            fmt.Println(v22, v3, fi32)
                            v44 := GenS8{F0: 2147483647, F1: (fi34 % ((fi32) | 1)), F2: 1000}
                            v45 := v1[54]
                            _ = fi34
                            _ = v35
                            _ = v44
                            _ = v45
                        }
                        var v46 int64 = int64(v2)
                        var v47 string = "foo bar"
                        _ = v46
                        _ = v47
                    }
                    var v48 uint64 = 42
                    _ = fi32
                    _ = v33
                    _ = v48
                }
                v22 -= v22
            }
        }
        for fi49 := 0; fi49 < 4; fi49++ {
            fmt.Println(v3)
            if (1000 < 42) {
                v50 := GenS5{F0: -926, F1: (fi49 + v2)}
                var v51 int32 = 1
                _ = v50
                _ = v51
            } else {
                v52 := []int{(-(fi49))}
                if (255 >= 32767) {
                    if len(v52) >= 2 {
                        v52 = v52[1:len(v52)]
                    }
                } else {
                    var v53 uint8 = 127
                    _ = len(v52)
                    var v54 int32 = 1
                    _ = v53
                    _ = v54
                }
                _ = v52
            }
            _ = fi49
        }
    } else {
        var v55 int64 = (v22 - (v22 % ((v22) | 1)))
        v56 := GenS1{F0: 65535}
        _ = v55
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
}
