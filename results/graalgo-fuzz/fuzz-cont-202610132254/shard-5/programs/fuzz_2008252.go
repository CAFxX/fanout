package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int16; F2 int16 }
func genFn0(p0 int16) int {
    p0 -= 42
    var v7 string = "hello"
    v7 = v7
    _ = p0
    _ = v7
    return 1
}
func genFn1() int64 {
    var v8 uint8 = 42
    if ("\n\t" != "a") {
        switch "hello" {
        case "foo bar":
            v9 := GenS1{F0: false}
            _ = v9
        case "a":
            if (false && (("" + "\n\t") > "\n\t")) {
                v10 := []int{}
                var v11 float32 = 1e10
                _ = v10
                _ = v11
            }
        case "":
            v8 = (v8 | (v8 ^ v8))
        default:
            v12 := []int{-1}
            _ = v12
        }
        fmt.Println(v8)
        if false {
            var v13 float32 = ((-2.25 - 1e10) - 0.0)
            fmt.Println(v13, v13)
            _ = v13
        }
    }
    fmt.Println(v8)
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 1
    v2 := []string{"", "a", ""}
    var v3 float32 = 0.1
    var v4 int64 = 9223372036854775807
    v5 := []uint32{2147483647, 2147483647, 4294967295}
    var v6 int32 = 0
    var v14 uint32 = 1
    var v15 uint32 = 1
    var v16 bool = ((^(v1)) <= 42)
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint64 = 2
        var v19 uint64 = v1
        fmt.Println(v4)
        _ = fi17
        _ = v18
        _ = v19
    }
    switch genFn0(280) {
    case -1:
        if false {
            if v16 {
                v20 := GenS1{F0: (-2.25 < 0.1)}
                _ = v20
            } else {
                var v21 bool = ((78 == 198) || true)
                _ = v21
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                v6 = v6
                v23 := GenS1{F0: true}
                _ = fi22
                _ = v23
            }
        }
    }
    fmt.Println(v15, v4, v15)
    for fi24 := 0; fi24 < 8; fi24++ {
        v6 = (v6 - (v6 << ((4294967295) & 63)))
        var v25 uint64 = 42
        v26 := GenS1{F0: ((v25 * 42) == (v1 * 4294967295))}
        _ = fi24
        _ = v25
        _ = v26
    }
    if (32767 == 900) {
        switch 2147483647 {
        case 32768:
            var v27 uint32 = (v14 % ((0) | 1))
            _ = v27
        case 128:
            fmt.Println(v6, v15)
        default:
            v2 = append(v2, "foo bar")
        }
    } else {
        for fi28 := 0; fi28 < 2; fi28++ {
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 float64 = 3.14159
                var v31 uint64 = (v1 / (((v1 ^ (v1 | v1))) | 1))
                switch fi28 {
                case -9223372036854775808:
                    v32 := GenS1{F0: v16}
                    _ = v32
                case 384:
                    if v16 {
                        var v33 int32 = (^(v6))
                        fmt.Println(v1)
                        switch ("a" + "") {
                        case "foo bar":
                            v34 := GenS2{F0: uint16(v14)}
                            _ = v34
                        case "":
                            v30 += (v30 / v30)
                        }
                        _ = v33
                    } else {
                        if len(v5) > 0 {
                            v5[0] = v14
                        }
                        v14 = v14
                        if v16 {
                            v1 = 4294967295
                        } else {
                            var v35 bool = v16
                            var v36 string = "foo bar"
                            _ = v35
                            _ = v36
                        }
                    }
                case 127:
                    var v37 uint32 = uint32(v31)
                    _ = v37
                default:
                    switch ("\n\t" + "") {
                    case "a":
                        v14 -= (v15 * (v14 & v14))
                    case "\n\t":
                        var v38 int64 = (-(v4))
                        _ = v38
                    default:
                        var v39 float32 = ((0.1 + v3) * v3)
                        _ = v39
                    }
                }
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
        fmt.Println(v16)
        v40 := GenS1{F0: false}
        _ = v40
    }
    var v41 uint16 = 566
    for fi42 := 0; fi42 < 15; fi42++ {
        switch "\n\t" {
        case "foo bar":
            switch "\n\t" {
            case "foo bar":
                v2 = append(v2, "\n\t")
            case "":
                switch "a" {
                case "a":
                    for fi43 := 0; fi43 < 16; fi43++ {
                        v4 = (v4 - 9223372036854775807)
                        var v44 uint = 554
                        var v45 string = "hello"
                        if (-1 > 32767) {
                            v46 := GenS3{F0: 0, F1: v14}
                            v1 += (v1 % (((v1 + v1)) | 1))
                            if v16 {
                                _ = len(v2)
                                v45 = ""
                                v44 += 2
                            }
                            _ = v46
                        } else {
                            v2 = append(v2, v45)
                            var v47 int64 = 42
                            _ = v47
                        }
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                case "":
                    var v48 uint16 = (v41 / (((v41 * v41)) | 1))
                    _ = v48
                case "foo bar":
                    for fi49 := 0; fi49 < 1; fi49++ {
                        v5 = append(v5, (v15 >> ((22) & 63)))
                        fmt.Println(v16, v15)
                        fmt.Println(v6, v14)
                        v1 -= 18446744073709551615
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v2 = append(v2, "a")
                        _ = fi49
                    }
                default:
                    v50 := GenS3{F0: v41, F1: (^(v15))}
                    _ = v50
                }
            case "a":
                fmt.Println(v6, v14, fi42)
            }
        }
        _ = fi42
    }
    if true {
        switch 42 {
        case 128:
            var v51 uint8 = 0
            _ = v51
        case 2147483647:
            if false {
                v14 -= v15
            } else {
                v1 = v1
                switch -1 {
                case 383:
                    var v52 uint16 = v41
                    _ = v52
                default:
                    v53 := GenS4{F0: 40}
                    _ = v53
                }
                for fi54 := 0; fi54 < 3; fi54++ {
                    fmt.Println(v41)
                    for fi55 := 0; fi55 < 4; fi55++ {
                        var v56 uint32 = 363
                        var v57 int8 = int8(v14)
                        v58 := GenS2{F0: (v41 | 965)}
                        _ = len(v2)
                        if ((32767 > -1) || (430 > 0)) {
                            fmt.Println(v6)
                            fmt.Println(v3, fi54, v16)
                        } else {
                            if len(v5) > 0 {
                                v5[0] = (v56 - 42)
                            }
                            v59 := GenS5{F0: "", F1: 42, F2: int16(fi54)}
                            var v60 uint8 = 48
                            _ = v59
                            _ = v60
                        }
                        _ = fi55
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    _ = fi54
                }
            }
        case -766:
            v1 = 18446744073709551615
        }
        fmt.Println(v4)
    }
    switch 42 {
    case 32768:
        var v61 int16 = 1000
        _ = v61
    default:
        v16 = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
