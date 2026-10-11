package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint16; F1 int16 }
func genFn0() int {
    v4 := GenS1{F0: 65535, F1: -102}
    var v5 uint64 = 1
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS2{F0: (fi6 % ((4294967295) | 1)), F1: 1}
            v5 = (v5 | v5)
            fmt.Println(fi7, fi6, fi7)
            var v9 uint32 = 248
            _ = fi7
            _ = v8
            _ = v9
        }
        switch "foo bar" {
        case "\n\t":
            if false {
                v5 = v5
            }
        case "hello":
            var v10 int32 = 2147483647
            _ = v10
        case "a":
            for fi11 := 0; fi11 < 6; fi11++ {
                if (float64(0.0) <= (float64((1e10 + 1e10)) - 1.5)) {
                    v5 = (v5 + (v5 | v5))
                }
                v12 := []int{(fi6 % ((fi6) | 1))}
                v5 -= 0
                fmt.Println(fi11)
                _ = fi11
                _ = v12
            }
        }
        _ = fi6
    }
    _ = v4
    _ = v5
    return 32768
}
func genFn1(p0 int16, p1 uint8, p2 int) uint8 {
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := []int{genFn0()}
        _ = fi13
        _ = v14
    }
    fmt.Println(p2, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return (^(p1))
}

func main() {
    var v1 string = ("" + "a")
    v2 := make(map[string]int)
    v3 := []uint32{399, 1}
    v1 = "hello"
    var v15 uint16 = 32767
    v15 = v15
    var v16 float64 = 0.1
    if len(v3) > 0 {
        v3[0] = 1
    }
    v16 += float64((-2.25 / ((-2.25 - ((3.14159 * 1e10) + 3.14159)) * 0.1)))
    v17 := GenS1{F0: 348, F1: -1}
    var v18 string = v1
    v15 = v15
    fmt.Println(v1)
    v2[v1] = -581
    v1 = "foo bar"
    v2[v18] = 1
    v15 = v15
    var v19 uint16 = (v15 ^ (v15 | (v15 << ((0) & 63))))
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 int8 = 1
        var v22 float32 = float32(v16)
        v23 := GenS2{F0: fi20, F1: 2}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 int16 = -1
    if true {
        var v25 float32 = 3.14159
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 int64 = -740
            _ = fi26
            _ = v27
        }
        _ = v25
    } else {
        for fi28 := 0; fi28 < 8; fi28++ {
            switch (v1 + v1) {
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "foo bar":
                switch v18 {
                case "foo bar":
                    v3 = append(v3, 81)
                case "hello":
                    if false {
                        if (!true) {
                            var v29 int16 = v24
                            _ = v29
                        }
                        var v30 string = v1
                        _ = v30
                    }
                case "a":
                    switch v1 {
                    case "foo bar":
                        for fi31 := 0; fi31 < 6; fi31++ {
                            v3 = append(v3, 2)
                            _ = fi31
                        }
                    case "":
                        for fi32 := 0; fi32 < 16; fi32++ {
                            v18 = ("" + "a")
                            _ = fi32
                        }
                    case "\n\t":
                        switch genFn0() {
                        case 127:
                            if (true || false) {
                                var v33 uint = 0
                                _ = v33
                            }
                        }
                    default:
                        if false {
                            v34 := v2[v18]
                            fmt.Println(v34)
                            _ = v34
                        } else {
                            _ = len(v2)
                            var v35 uint32 = 97
                            v16 -= v16
                            _ = v35
                        }
                    }
                default:
                    if false {
                        for fi36 := 0; fi36 < 16; fi36++ {
                            var v37 bool = true
                            _ = fi36
                            _ = v37
                        }
                    }
                }
            default:
                for fi38 := 0; fi38 < 14; fi38++ {
                    switch "hello" {
                    case "\n\t":
                        if true {
                            if false {
                                v16 += (3.14159 + (v16 + v16))
                                v3 = append(v3, 42)
                                switch fi28 {
                                case 9223372036854775807:
                                    var v39 float64 = v16
                                    _ = v39
                                }
                            } else {
                                v24 += v24
                                switch int(v24) {
                                case -9223372036854775808:
                                    if false {
                                        var v40 uint16 = 0
                                        _ = v40
                                    } else {
                                        var v41 uint32 = 2
                                        _ = v41
                                    }
                                case 958:
                                    var v42 float64 = -2.25
                                    _ = v42
                                default:
                                    v43 := GenS2{F0: int(v19), F1: 965}
                                    _ = v43
                                }
                                if true {
                                    v44 := GenS2{F0: fi28, F1: 65535}
                                    _ = v44
                                }
                            }
                        } else {
                            v45 := GenS3{F0: 1}
                            switch (fi28 - (fi38 >> ((2) & 63))) {
                            case 575:
                                _ = len(v2)
                                if true {
                                    if ((-(fi38)) > (fi28 + (fi38 | 316))) {
                                        v1 = ("foo bar" + ("a" + v18))
                                        fmt.Println(fi38)
                                        switch v1 {
                                        case "a":
                                            if (!(9223372036854775807 == 9223372036854775807)) {
                                                v19 += (v15 | v15)
                                                v24 = 475
                                            } else {
                                                var v46 uint8 = 2
                                                var v47 string = v1
                                                v19 += v15
                                                _ = v46
                                                _ = v47
                                            }
                                        default:
                                            var v48 int8 = -30
                                            _ = v48
                                        }
                                    }
                                    if false {
                                        fmt.Println(v18, v18, fi28)
                                    } else {
                                        var v49 int16 = (v24 / ((v24) | 1))
                                        var v50 uint = 822
                                        switch "foo bar" {
                                        case "hello":
                                            var v51 uint = v50
                                            _ = v51
                                        case "":
                                            v52 := GenS4{F0: fi38, F1: v24, F2: 42}
                                            _ = v52
                                        case "\n\t":
                                            if false {
                                                switch (("\n\t" + "a") + v18) {
                                                case "foo bar":
                                                    if false {
                                                        v49 = v24
                                                        var v53 int16 = (-(v24))
                                                        var v54 float64 = v16
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                }
                                                fmt.Println(v24, v15, v49)
                                                var v55 uint64 = 568
                                                _ = v55
                                            } else {
                                                var v56 uint = (^(v50))
                                                _ = len(v3)
                                                _ = v56
                                            }
                                        default:
                                            switch (fi38 & (fi38 << ((2) & 63))) {
                                            case -1:
                                                if false {
                                                    v57 := v2[(v1 + ("\n\t" + ((v1 + v1) + v1)))]
                                                    _ = v57
                                                }
                                                v1 = ((v18 + (v18 + v1)) + v18)
                                            case 4294967295:
                                                v49 -= (v49 | (v49 >> ((441) & 63)))
                                            }
                                        }
                                        _ = v49
                                        _ = v50
                                    }
                                } else {
                                    v18 = v1
                                }
                            default:
                                var v58 uint16 = 73
                                _ = v58
                            }
                            v16 = v16
                            _ = v45
                        }
                    case "":
                        v59 := GenS4{F0: fi38, F1: (v24 % (((v24 / ((v24) | 1))) | 1)), F2: 568}
                        _ = v59
                    case "hello":
                        var v60 bool = true
                        _ = v60
                    }
                    if false {
                        var v61 uint8 = genFn1(v24, 0, 2147483647)
                        _ = v61
                    } else {
                        fmt.Println(v16, v16)
                        if ((!false) || (-67 >= 51)) {
                            fmt.Println(fi38, v16)
                        } else {
                            var v62 string = v18
                            v24 = (v24 / (((v24 + (v24 + (v24 / (((v24 | v24)) | 1))))) | 1))
                            _ = v62
                        }
                        var v63 uint8 = 222
                        _ = v63
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi38
                }
            }
            _ = fi28
        }
        v64 := GenS5{F0: v19, F1: (v24 - (v24 + 42))}
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
}
