package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
func genFn0(p0 int64, p1 string, p2 uint16) string {
    fmt.Println(p2)
    p0 = p0
    v7 := GenS1{F0: -126}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1() int16 {
    for fi8 := 0; fi8 < 1; fi8++ {
        switch fi8 {
        case -1000:
            var v9 uint64 = 2
            _ = v9
        case 32767:
            v10 := GenS2{F0: 1000}
            _ = v10
        case 10:
            v11 := []int{fi8, fi8, 698, fi8}
            _ = v11
        default:
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 uint16 = 476
                if true {
                    var v14 uint64 = 2
                    _ = v14
                }
                var v15 int32 = 550
                fmt.Println(fi8, fi12, fi12)
                switch 32768 {
                case 42:
                    fmt.Println(v15)
                case 32768:
                    if false {
                        v13 = (v13 - v13)
                    }
                }
                _ = fi12
                _ = v13
                _ = v15
            }
        }
        var v16 bool = true
        _ = fi8
        _ = v16
    }
    var v17 bool = (-2147483648 > 2147483647)
    var v18 int = 0
    v19 := []int{}
    v17 = (127 >= -1)
    _ = v17
    _ = v18
    _ = v19
    return 1
}

func main() {
    var v1 uint16 = 2
    var v2 bool = true
    var v3 int8 = 100
    v4 := make(map[uint32]string)
    var v5 uint32 = 895
    v6 := []string{"\n\t", "", "foo bar", "\n\t"}
    if v2 {
        if ("foo bar" != ("a" + "a")) {
            fmt.Println(v1)
            _ = len(v6)
        } else {
            fmt.Println(v1, v5, v3)
            var v20 float64 = (((0.0 * -2.25) + 3.14159) - 0.0)
            _ = v20
        }
        var v21 string = "foo bar"
        _ = v21
    }
    if true {
        var v22 bool = ((^(v5)) <= (v5 ^ 2))
        var v23 bool = (false || v22)
        _ = v22
        _ = v23
    }
    v6 = append(v6, "hello")
    var v24 uint = 4294967295
    if v2 {
        switch "\n\t" {
        case "hello":
            var v25 uint32 = (v5 >> ((uint(v1)) & 63))
            _ = v25
        case "a":
            v5 = uint32(v1)
        case "\n\t":
            var v26 uint8 = 0
            _ = v26
        default:
            var v27 uint8 = 0
            _ = v27
        }
        var v28 string = ""
        _ = v28
    } else {
        v5 -= v5
    }
    v5 = (v5 % ((v5) | 1))
    v29 := v4[513]
    switch v29 {
    case "foo bar":
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 float64 = 1e10
            for fi32 := 0; fi32 < 17; fi32++ {
                if (true && (((!false) && (42 <= v1)) || (((v31 != 1.5) && false) && v2))) {
                    fmt.Println(fi30, v29)
                    v31 = 1e10
                    switch v29 {
                    case "a":
                        v5 = v5
                    case "foo bar":
                        if v2 {
                            var v33 int = fi32
                            v34 := GenS3{F0: v29}
                            _ = v33
                            _ = v34
                        } else {
                            delete(v4, (v5 | 2))
                            v1 = v1
                            var v35 float64 = (v31 / (1.5 - ((((((-2.25 - v31) * (float64(0.1) / v31)) + ((v31 - (((-2.25 * (float64(1e10) / (float64((3.14159 / 1e10)) - (v31 - ((v31 + v31) + ((3.14159 + v31) + v31)))))) / (((v31 / -2.25) / 1.0) - 3.14159)) - (v31 * (0.0 * v31)))) / v31)) + 0.0) / v31) / v31)))
                            if (0 != 1) {
                                fmt.Println(v1, v3)
                                var v36 int8 = (v3 * v3)
                                _ = v36
                            }
                            _ = v35
                        }
                    case "":
                        var v37 float64 = 3.14159
                        _ = v37
                    default:
                        var v38 bool = (!v2)
                        _ = v38
                    }
                    if v2 {
                        v39 := v4[(^(v5))]
                        fmt.Println(v1)
                        var v40 float64 = v31
                        _ = v39
                        _ = v40
                    }
                    var v41 int = fi30
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v41
                } else {
                    var v42 float64 = v31
                    var v43 uint8 = 255
                    _ = v42
                    _ = v43
                }
                _ = fi32
            }
            var v44 bool = ((v31 * v31) == (v31 / (float64(0.1) + v31)))
            _ = fi30
            _ = v31
            _ = v44
        }
    case "\n\t":
        v45 := v4[(v5 & 309)]
        _ = v45
    default:
        v46 := GenS1{F0: 71}
        _ = v46
    }
    v47 := GenS2{F0: -487}
    switch "\n\t" {
    case "":
        fmt.Println(v5, v2, v2)
    case "a":
        var v48 uint64 = 155
        _ = v48
    case "foo bar":
        var v49 uint = (v24 % (((v24 ^ v24)) | 1))
        _ = v49
    }
    v24 = v24
    if false {
        if v2 {
            fmt.Println(v24, v3)
            var v50 string = "\n\t"
            _ = v50
        } else {
            for fi51 := 0; fi51 < 5; fi51++ {
                for fi52 := 0; fi52 < 13; fi52++ {
                    v2 = v2
                    switch "a" {
                    case "\n\t":
                        if (0 == 127) {
                            var v53 string = v29
                            v3 = v3
                            _ = v53
                        }
                    default:
                        v2 = (65535 != 407)
                    }
                    _ = fi52
                }
                if (!(false || (v2 || v2))) {
                    if v2 {
                        if true {
                            var v54 uint16 = (v1 & (v1 << ((v24) & 63)))
                            _ = v54
                        } else {
                            for fi55 := 0; fi55 < 14; fi55++ {
                                if (((((v29 + v29) + ((v29 + (v29 + v29)) + (v29 + "hello"))) + v29) + (v29 + "foo bar")) == (v29 + v29)) {
                                    v5 -= 0
                                    switch "foo bar" {
                                    case "a":
                                        var v56 float32 = 1.5
                                        _ = v56
                                    case "":
                                        v1 -= (v1 * v1)
                                    case "foo bar":
                                        v1 -= (v1 | 1)
                                    }
                                    v57 := GenS2{F0: 192}
                                    _ = v57
                                } else {
                                    if (!((!(v2 || v2)) || ((fi55 >> ((4294967295) & 63)) < (fi51 ^ (fi51 + (fi51 % ((fi55) | 1))))))) {
                                        fmt.Println(v3)
                                    } else {
                                        v3 = v3
                                        if (!(v3 <= -52)) {
                                            var v58 int64 = 9223372036854775807
                                            var v59 int32 = 2147483647
                                            _ = v58
                                            _ = v59
                                        } else {
                                            v4[1] = v29
                                            v1 += v1
                                            var v60 int16 = 1000
                                            _ = v60
                                        }
                                        if (v2 && (!((!false) && v2))) {
                                            v29 = "hello"
                                        }
                                    }
                                    var v61 uint16 = 32767
                                    _ = v61
                                }
                                switch "" {
                                case "\n\t":
                                    v29 = ("a" + ((v29 + ("foo bar" + "\n\t")) + (v29 + v29)))
                                case "a":
                                    var v62 int8 = 108
                                    _ = v62
                                default:
                                    fmt.Println(fi51, v2)
                                }
                                _ = fi55
                            }
                            var v63 int32 = -2147483648
                            _ = v63
                        }
                        v2 = v2
                        var v64 int8 = v3
                        _ = v64
                    }
                } else {
                    switch 2147483647 {
                    case -795:
                        if len(v6) > 0 {
                            v6[0] = ((v29 + (v29 + ("\n\t" + v29))) + (v29 + "foo bar"))
                        }
                    default:
                        if v2 {
                            v1 -= v1
                            var v65 uint64 = 467
                            _ = v65
                        }
                    }
                }
                _ = fi51
            }
        }
    }
    for fi66 := 0; fi66 < 5; fi66++ {
        fmt.Println(fi66, v2)
        for fi67 := 0; fi67 < 3; fi67++ {
            fmt.Println(v24)
            _ = fi67
        }
        _ = fi66
    }
    _ = len(v6)
    v5 = 2147483647
    var v68 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
    _ = v68
}
